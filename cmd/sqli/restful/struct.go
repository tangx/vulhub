package main

type User struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}
