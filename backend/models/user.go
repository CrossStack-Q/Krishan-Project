package models

import (
	"database/sql"
	"fmt"

	"github.com/CrossStack-Q/Vacation-App/db"
	"github.com/CrossStack-Q/Vacation-App/types"
)

var database *sql.DB = db.Connect()

func CreateUser(user types.User) error {
	query := `INSERT INTO users (email, password) VALUES ($1, $2)`
	_, err := database.Exec(query, user.Email, user.Password)
	if err != nil {
		fmt.Printf("Error executing query: %v\n", err) // Log detailed error
	}
	return err
}

func GetUserByEmail(email string) (types.User, error) {
	var user types.User
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := database.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}
