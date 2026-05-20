package database

import (
	"context"
	"golab/components/hash"
	"log"
)

func UserNameExists(username string) bool {
	conn, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	userNameExists := verifyUsername(conn, username)
	defer conn.Close(context.Background())

	return userNameExists
}

func SaveNewUser(username string, password string) error {
	conn, err := Connect()
	if err != nil {
		return err
	}
	err = saveUser(conn, username, hash.HashPassword(password))
	defer conn.Close(context.Background())
	return err
}

func VerifyLogin(username, password string) bool {
	conn, err := Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())
	hashPassword, err := getHash(conn, username)
	if err != nil {
		log.Printf("Error fetching hash for user %s: %v", username, err)
		return false
	}
	return hashPassword != "" && hashPassword == string(rune(hash.HashPassword(password)))
}
