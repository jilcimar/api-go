package models

import "time"

type User struct {
	ID         uint64    `json:"id,omitempty"`
	name       string    `json:"name,omitempty"`
	username   string    `json:"username,omitempty"`
	email      string    `json:"email,omitempty"`
	password   string    `json:"password,omitempty"`
	created_at time.Time `json:"created_at,omitempty"`
}
