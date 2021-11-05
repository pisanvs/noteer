package main

import (
	"context"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/google/uuid"

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

type LoginForm struct {
	username string `validate:"required,printascii"`
	password string `validate:"required,printascii"`
}

type DBUser struct {
	username     string `validate:"required,printascii "`
	password     string `validate:"required,printascii"`
	name         string `validate:"required,alpha"`
	email        string `validate:"required,email"`
	creationDate time.Time
	lastLogin    time.Time
}

// type User struct {
// 	username     string
// 	password     string
// 	name         string
// 	email        string
// 	creationDate time.Time
// 	lastLogin    time.Time
// }

type DBSession struct {
	sessionID string
	username  string
	created   time.Time
	expires   time.Time
}

// VARIABLES

var client *mongo.Client
var ctx = context.Background()

var sessionsDB *mongo.Collection
var usersDB *mongo.Collection

var validate *validator.Validate

var secureCookie *securecookie.SecureCookie

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

	// create securecookie instance
	hashKey := os.Getenv("SECURECOOKIE_HASH")
	blockKey := os.Getenv("SECURECOOKIE_BLOCK")

	secureCookie = securecookie.New([]byte(hashKey), []byte(blockKey))

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
	err := sessionsDB.FindOne(context.TODO(), bson.M{"sessionID": authToken}).Decode(&result)
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
	// check if method is POST
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Wrong method, use POST for login endpoint"))
		return
	}

	// check if body is empty
	if r.Body == nil {
		w.WriteHeader(400)
		w.Write([]byte("Empty request, try again"))
		return
	}

	// decode request body
	form := LoginForm{
		username: r.FormValue("username"),
		password: r.FormValue("password"),
	}

	// validate form
	err := validate.Struct(form)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Validation of data failed" + err.Error()))
	}

	// check if user exists
	var user DBUser
	err = usersDB.FindOne(ctx, bson.M{"username": form.username}).Decode(&user)
	fmt.Println(err)
	fmt.Print("User: ")
	fmt.Println(user)
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("User does not exist. Error: " + err.Error()))
		return
	}

	// check if password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.password), []byte(form.password))
	if err != nil {
		w.WriteHeader(401)
		w.Write([]byte("Wrong password. Error:" + err.Error()))
		return
	}

	// create sessionID
	su := uuid.New().String()
	sessionID := sha512.Sum512([]byte(su + " " + form.username))

	// create session document
	session := DBSession{
		sessionID: string(sessionID[:]),
		username:  form.username,
		created:   time.Now(),
		expires:   time.Now().Add(time.Hour * 24 * 7),
	}

	// insert session document
	_, err = sessionsDB.InsertOne(context.TODO(), session)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error creating session"))
		return
	}

	// send encrypted session cookie
	encoded, err := secureCookie.Encode("session", sessionID)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error encrypting session, try again"))
		return
	}

	// set cookie
	cookie := http.Cookie{
		Name:     "sesh",
		Value:    encoded,
		Path:     "/",
		HttpOnly: false,
		Domain:   "." + os.Getenv("MAIN_DOMAIN"),
		Expires:  time.Now().Add(time.Hour * 24 * 7),
	}
	http.SetCookie(w, &cookie)

}

func register(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Wrong method, use POST for register endpoint"))
		return
	}

	// check if body is empty
	if r.Body == nil {
		w.WriteHeader(400)
		w.Write([]byte("Empty request, try again"))
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
	err = usersDB.FindOne(context.TODO(), bson.M{"username": form.username}).Decode(&result)
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

	usersDB.InsertOne(ctx, bson.M{"username": newUser.username, "password": newUser.password, "name": newUser.name, "email": newUser.email, "creationDate": newUser.creationDate, "lastLogin": newUser.lastLogin})

	w.WriteHeader(200)
	w.Write([]byte("User created"))

}
