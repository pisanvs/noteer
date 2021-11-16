package main

import (
	"time"
)

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

type DocumentRef struct {
	ID           string `validate:"required,printascii"`
	Title        string `validate:"required,printascii"`
	CreationDate time.Time
	LastModified time.Time
}

type Document struct {
	ID           string `validate:"required,printascii"`
	Title        string `validate:"required,printascii"`
	Owner        string `validate:"required,printascii"`
	Content      []byte
	CreationDate time.Time
	LastModified time.Time
}

type UserData struct {
	Username     string `validate:"required,printascii "`
	ProfileImage string `validate:"url"`
	Documents    []DocumentRef
}

type DBSession struct {
	SessionID string
	Username  string
	Created   time.Time
	Expires   time.Time
}
