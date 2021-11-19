/**
* Copyright (C) 2021  Maximiliano Morel (pisanvs) <maxmorel@pisanvs.cl>
*
* This file is part of Noteer, a note taking application.
*
* Noteer is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License v3 as
* published by the Free Software Foundation
*
* Noteer is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with Noteer.  If not, see <https://www.gnu.org/licenses/>.
*
*
* @license GPL-3.0 <https://www.gnu.org/licenses/gpl-3.0.txt>
 */

package main

import (
	"context"
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

	"github.com/go-playground/validator/v10"

	"golang.org/x/crypto/bcrypt"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// CONSTANTS

const apiRoot string = "/api/v1"

// VARIABLES

var ctx = context.Background()

var sessionsDB *mongo.Collection
var usersDB *mongo.Collection
var userDataDB *mongo.Collection
var documentsDB *mongo.Collection

var dmp *diffmatchpatch.DiffMatchPatch

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

	// create validator instance
	validate = validator.New()

	// create diff match patch instance
	dmp = diffmatchpatch.New()

	// create securecookie instance
	hashKey := os.Getenv("SECURECOOKIE_HASH")
	blockKey := os.Getenv("SECURECOOKIE_BLOCK")

	sc = securecookie.New([]byte(hashKey), []byte(blockKey))

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://172.16.150.17:27017/?readPreference=primary&connect=direct&sslInsecure=true"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer ctxCancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	sessionsDB = client.Database("auth").Collection("sessions")
	usersDB = client.Database("auth").Collection("users")

	userDataDB = client.Database("users").Collection("userData")
	documentsDB = client.Database("users").Collection("documents")

	fmt.Println(client.ListDatabaseNames(ctx, bson.M{}))

	mainMux := mux.NewRouter().StrictSlash(true)
	mainMux.HandleFunc(apiRoot+"/auth/session", session).Methods("GET", "OPTIONS", "POST")
	mainMux.HandleFunc(apiRoot+"/auth/login", login)
	mainMux.HandleFunc(apiRoot+"/auth/register", register)

	mainMux.HandleFunc(apiRoot+"/data/user", user)
	mainMux.HandleFunc(apiRoot+"/data/docs/new", newDoc)
	mainMux.HandleFunc(apiRoot+"/data/docs/get", getDoc)
	mainMux.HandleFunc(apiRoot+"/data/docs/edit", editDoc)

	mainMux.Use(mux.CORSMethodMiddleware(mainMux))

	log.Fatal(http.ListenAndServe(":8888", mainMux))
}
