package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://admin:secretpassword@localhost:5433/mydatabase2")
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func QueryData(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "SELECT userid, username FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User ID: %d, Name: %s\n", id, name)
	}
}

func verifyUsername(conn *pgx.Conn, username string) bool {
	rows, err := conn.Query(context.Background(), "SELECT username FROM users WHERE username = $1", username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	return rows.Next()
}

func saveUser(conn *pgx.Conn, username string, password uint64) error {
	_, err := conn.Exec(context.Background(), "INSERT INTO users (username, passwordhash) VALUES ($1, $2)", username, string(rune(password)))
	if err != nil {
		fmt.Printf("Error executing query: %v\n", password)
		fmt.Printf("Error saving user: %v\n", err)
		return err
	}
	return nil
}

func getHash(conn *pgx.Conn, username string) (string, error) {
	var hash string
	err := conn.QueryRow(context.Background(), "SELECT passwordhash FROM users WHERE username = $1", username).Scan(&hash)
	if err != nil {
		log.Printf("Error executing query: %v\n", err)
		return "", err
	}

	return hash, nil
}
