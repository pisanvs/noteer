package main

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
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
	Username     string `validate:"required,printascii "`
	Password     string `validate:"required,printascii"`
	Name         string `validate:"required,alpha"`
	Email        string `validate:"required,email"`
	CreationDate time.Time
	LastLogin    time.Time
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
	SessionID string
	Username  string
	Created   time.Time
	Expires   time.Time
}

// VARIABLES

var client *mongo.Client
var ctx = context.Background()

var sessionsDB *mongo.Collection
var usersDB *mongo.Collection

var validate *validator.Validate

var sc *securecookie.SecureCookie

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

	sc = securecookie.New([]byte(hashKey), []byte(blockKey))

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
	mainMux.HandleFunc(apiRoot+"/auth/session", session).Methods("GET", "OPTIONS", "POST")
	mainMux.HandleFunc(apiRoot+"/auth/login", login)
	mainMux.HandleFunc(apiRoot+"/auth/register", register)

	mainMux.Use(mux.CORSMethodMiddleware(mainMux))

	log.Fatal(http.ListenAndServe(":8000", mainMux))
}

// START AUTH ROUTES

func session(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")

	if r.Method == http.MethodOptions {
		return
	}

	// decode session cookie
	var session [64]byte
	auth := r.Header.Get("Authorization")
	err := sc.Decode("session", auth, &session)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	// get session from db
	var result bson.M
	err = sessionsDB.FindOne(ctx, bson.M{"sessionid": hex.EncodeToString(session[:])}).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		fmt.Print(err)
	}

	response, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("JSON marshal error: " + err.Error()))
		return
	}
	w.Write([]byte(string(response)))
}

func login(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Wrong method, use POST for login endpoint"))
		return
	}

	// check if body is empty
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Validation of data failed" + err.Error()))
	}

	// check if user exists
	var user DBUser
	err = usersDB.FindOne(ctx, bson.M{"username": form.username}).Decode(&user)
	fmt.Println(err)
	fmt.Print("User: ")
	fmt.Println(user)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("User does not exist. Error: " + err.Error()))
		return
	}

	// check if password is correct
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(form.password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Wrong password. Error:" + err.Error()))
		return
	}

	// create sessionID
	su := uuid.New().String()
	sessionID := sha512.Sum512([]byte(su + " " + form.username))

	// create session document
	session := DBSession{
		SessionID: hex.EncodeToString(sessionID[:]),
		Username:  form.username,
		Created:   time.Now(),
		Expires:   time.Now().Add(time.Hour * 24 * 7),
	}

	// insert session document
	_, err = sessionsDB.InsertOne(ctx, session)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating session"))
		return
	}

	// send encrypted session cookie
	encoded, err := sc.Encode("session", sessionID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error encrypting session, try again"))
		return
	}

	// set cookie
	cookie := http.Cookie{
		Name:     "sesh",
		Value:    encoded,
		Path:     "/",
		HttpOnly: false,
		Secure:   true,
		Domain:   "." + os.Getenv("MAIN_DOMAIN"),
		Expires:  time.Now().Add(time.Hour * 24 * 7),
	}
	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "https://"+os.Getenv("MAIN_DOMAIN")+"/app", http.StatusFound)
}

func register(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Wrong method, use POST for register endpoint"))
		return
	}

	// check if body is empty
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusBadRequest)
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username already exists"))
		return
	}

	// Hash password
	hashedPass, err := hashPassword(form.password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error hashing password. Please try again, if it still doesn't work, submit an issue on Github"))
		return
	}

	// Create new user
	newUser := DBUser{
		Username:     form.username,
		Password:     hashedPass,
		Name:         form.name,
		Email:        form.email,
		CreationDate: time.Now(),
		LastLogin:    time.Now(),
	}

	usersDB.InsertOne(ctx, bson.M{"username": newUser.Username, "password": newUser.Password, "name": newUser.Name, "email": newUser.Email, "creationDate": newUser.CreationDate, "lastLogin": newUser.LastLogin})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created"))

}
