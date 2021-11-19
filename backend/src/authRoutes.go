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
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func session(w http.ResponseWriter, r *http.Request) {
	// TODO: CORS CONFIGURATION -- when ready for prod**
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")

	if r.Method == http.MethodOptions {
		return
	}

	// decode session cookie
	var seshBytes [64]byte
	err := sc.Decode("session", r.Header.Get("Authorization"), &seshBytes)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error decoding session"))
		fmt.Println(err)
		return
	}

	// get session from db
	var dbsesh DBSession
	err = sessionsDB.FindOne(ctx, bson.M{"sessionid": hex.EncodeToString(seshBytes[:])}).Decode(&dbsesh)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error getting session, try logging in again"))
		return
	}

	// check if session is expired
	if dbsesh.Expires.Before(time.Now()) {
		sessionsDB.FindOneAndDelete(ctx, bson.M{"sessionid": session})
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Session expired, try logging in again"))
		return
	}

	response, err := json.Marshal(dbsesh)
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
