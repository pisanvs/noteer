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

// VARIABLES

var client *mongo.Client
var ctx context.Context

var sessionsDB *mongo.Collection
var usersDB *mongo.Collection

var validate *validator.Validate

func main() {

	validate = validator.New()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://172.16.150.17:27017"))
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
	mainMux.PathPrefix(apiRoot)
	mainMux.HandleFunc("/auth/session", session)
	mainMux.HandleFunc("/auth/login", login)
	mainMux.HandleFunc("/auth/register", register)

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
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Wrong method, use POST for register endpoint"))
		return
	}

	form := RegisterForm{
		name:     r.Form.Get("name"),
		username: r.Form.Get("username"),
		password: r.Form.Get("password"),
		email:    r.Form.Get("email"),
	}

	err := validate.Struct(form)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Validation of data Failed"))
		return
	}

}
