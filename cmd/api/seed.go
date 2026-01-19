package main

import (
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func SeedUser() {
	email := "admin@test.com"
	username := "admin"
	password := "1234"
	firstName := "Admin"
	lastName := "User"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(
		context.Background(),
		`INSERT INTO users (email, username, password_hash, first_name, last_name)
		 VALUES ($1, $2, $3, $4, $5)
		 ON CONFLICT (email) DO NOTHING`,
		email,
		username,
		string(hash),
		firstName,
		lastName,
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ user created: admin@test.com / 1234")
}
