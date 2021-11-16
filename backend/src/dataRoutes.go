package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func user(w http.ResponseWriter, r *http.Request) {
	// TODO: CORS CONFIGURATION -- when ready for prod**
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")

	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed, use GET"))
		return
	}

	// decode session cookie
	var session [64]byte
	auth := r.Header.Get("Authorization")
	err := sc.Decode("session", auth, &session)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error decoding session"))
		return
	}

	// get session from db
	var dbsesh DBSession
	err = sessionsDB.FindOne(ctx, bson.M{"sessionid": session}).Decode(&dbsesh)
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

	// get user data from database
	var userData UserData
	userDataDB.FindOne(ctx, bson.M{"username": dbsesh.Username}).Decode(&userData)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	// return user data
	w.WriteHeader(http.StatusOK)
	response, err := json.Marshal(userData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(response)
}

func newDoc(w http.ResponseWriter, r *http.Request) {
	// TODO: CORS CONFIGURATION -- when ready for prod**
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")

	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed, use POST"))
		return
	}

	// decode session cookie
	var session [64]byte
	auth := r.Header.Get("Authorization")
	err := sc.Decode("session", auth, &session)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error decoding session"))
		return
	}

	// get session from db
	var dbsesh DBSession
	err = sessionsDB.FindOne(ctx, bson.M{"sessionid": session}).Decode(&dbsesh)
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

	// create new doc
	result, err := documentsDB.InsertOne(ctx, bson.M{"title": r.FormValue("title"), "content": []byte(""), "owner": dbsesh.Username})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating document, please try again"))
		return
	}

	// return success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Document created succesfully, id: " + result.InsertedID.(primitive.ObjectID).Hex()))
}

func getDoc(w http.ResponseWriter, r *http.Request) {
	// TODO: CORS CONFIGURATION -- when ready for prod**
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Authorization")

	if r.Method == http.MethodOptions {
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed, use GET"))
		return
	}

	// decode session cookie
	var session [64]byte
	auth := r.Header.Get("Authorization")
	err := sc.Decode("session", auth, &session)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error decoding session"))
		return
	}

	// get session from db
	var dbsesh DBSession
	err = sessionsDB.FindOne(ctx, bson.M{"sessionid": session}).Decode(&dbsesh)
	fmt.Println("session cookie " + auth)
	fmt.Println("session username " + dbsesh.Username)
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

	// get doc from db
	var doc Document
	err = documentsDB.FindOne(ctx, bson.M{"_id": r.FormValue("id")}).Decode(&doc)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error getting document"))
		return
	}

	// check if user is owner
	if doc.Owner != dbsesh.Username {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("You are not the owner of this document"))
		return
	}

	// return doc

}

func editDoc(w http.ResponseWriter, r *http.Request) {

}
