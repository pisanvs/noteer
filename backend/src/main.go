package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-playground/validator/v10"

	"golang.org/x/crypto/bcrypt"
)

// CONSTANTS

const apiRoot string = "/api/v1"

// REQUEST DEFINITIONS

type BasicAuthBody struct {
	Auth string `json:"Auth"`
}

type RegisterForm struct {
	username string `validate:"required,printascii"`
	password string `validate:"required,printascii"`
	name     string `validate:"required,alpha"`
	email    string `validate:"required,email"`
}

type DBUser struct {
	username     string `validate:"required,printascii "`
	password     string `validate:"required,printascii"`
	name         string `validate:"required,alpha"`
	email        string `validate:"required,email"`
	creationDate time.Time
	lastLogin    time.Time
}

// VARIABLES

var client *mongo.Client
var ctx context.Context

var sessionsDB *mongo.Collection
var usersDB *mongo.Collection

var validate *validator.Validate

// hash password using bcrypt
func hashPassword(password string) (string, error) {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		return "", err
	}
	return string(bcryptPassword), nil
}

func main() {

	validate = validator.New()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://172.16.150.17:27017/?readPreference=primary&connect=direct&sslInsecure=true"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	sessionsDB = client.Database("auth").Collection("sessions")
	usersDB = client.Database("auth").Collection("users")
	fmt.Println(client.ListDatabaseNames(ctx, bson.M{}))

	mainMux := mux.NewRouter().StrictSlash(true)
	mainMux.HandleFunc(apiRoot+"/auth/session", session)
	mainMux.HandleFunc(apiRoot+"/auth/login", login)
	mainMux.HandleFunc(apiRoot+"/auth/register", register)

	log.Fatal(http.ListenAndServe(":8000", mainMux))
}

// START AUTH ROUTES

func session(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header["Authorization"][0]

	var result bson.M
	err := sessionsDB.FindOne(context.TODO(), bson.D{{"sessionID", authToken}}).Decode(&result)
	fmt.Print(result)
	if err != nil {
		w.WriteHeader(401)
		fmt.Print(err)
	}

	response, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write(response)
}

func login(w http.ResponseWriter, r *http.Request) {

}

func register(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Wrong method, use POST for register endpoint"))
		return
	}

	// decode request body
	form := RegisterForm{
		name:     r.FormValue("name"),
		username: r.FormValue("username"),
		password: r.FormValue("password"),
		email:    r.FormValue("email"),
	}

	// Validate form
	err := validate.Struct(form)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Validation of data Failed"))
		return
	}

	// Check if username already exists
	var result bson.M
	err = usersDB.FindOne(context.TODO(), bson.D{{"username", form.username}}).Decode(&result)
	if err != nil {
		fmt.Print(err)
	}
	if result != nil {
		w.WriteHeader(400)
		w.Write([]byte("Username already exists"))
		return
	}

	// Hash password
	hashedPass, err := hashPassword(form.password)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error hashing password. Please try again, if it still doesn't work, submit an issue on Github"))
		return
	}

	// Create new user
	newUser := DBUser{
		username:     form.username,
		password:     hashedPass,
		name:         form.name,
		email:        form.email,
		creationDate: time.Now(),
		lastLogin:    time.Now(),
	}

	usersDB.InsertOne(ctx, bson.D{{"username", newUser.username}, {"password", newUser.password}, {"name", newUser.name}, {"email", newUser.email}, {"creationDate", newUser.creationDate}, {"lastLogin", newUser.lastLogin}})

	w.WriteHeader(200)
	w.Write([]byte("User created"))

}
