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
