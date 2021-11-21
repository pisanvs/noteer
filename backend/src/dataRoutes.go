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
	"encoding/hex"
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

	// create new doc
	result, err := documentsDB.InsertOne(ctx, bson.M{"title": r.FormValue("title"), "content": []byte(""), "owner": dbsesh.Username})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating document, please try again"))
		return
	}

	// make document reference object
	docRef := DocumentRef{
		ID:           result.InsertedID.(primitive.ObjectID).String(),
		Title:        r.FormValue("title"),
		CreationDate: time.Now(),
		LastModified: time.Now(),
	}

	// add document reference to user data
	userDataDB.FindOneAndUpdate(ctx, bson.M{"username": dbsesh.Username}, bson.M{"$push": bson.M{"documents": docRef}})

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
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed, use POST"))
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

	// get doc from db
	var doc Document
	docid, err := primitive.ObjectIDFromHex(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error parsing id, please try again"))
		return
	}
	err = documentsDB.FindOne(ctx, bson.M{"_id": docid}).Decode(&doc)
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

	w.WriteHeader(http.StatusOK)
	w.Write(doc.Content)

}

func editDoc(w http.ResponseWriter, r *http.Request) {
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

	// get doc from db
	var doc Document
	docid, err := primitive.ObjectIDFromHex(r.FormValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error parsing id, please try again"))
		return
	}
	err = documentsDB.FindOne(ctx, bson.M{"_id": docid}).Decode(&doc)
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

	// edit doc
	patches, err := dmp.PatchFromText(r.FormValue("content"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error parsing content, please try again"))
		return
	}
	patched, app := dmp.PatchApply(patches, string(doc.Content))

	fmt.Println(patched)
	fmt.Println(patches)

	for _, p := range app {
		if !p {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error editing document, please try again"))
			return
		}
	}

	// update doc
	_, err = documentsDB.UpdateOne(ctx, bson.M{"_id": docid}, bson.M{"$set": bson.M{"content": []byte(patched)}})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error editing document, please try again"))
		return
	}

	// return success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Document edited successfully"))

}
