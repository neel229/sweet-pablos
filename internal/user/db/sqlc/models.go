// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"time"
)

type Uaccount struct {
	ID        int64          `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  sql.NullString `json:"last_name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Ts        time.Time      `json:"ts"`
}