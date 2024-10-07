package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func stateDB(db *sql.DB) *sql.Stmt {
	state, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (username,login,password) ")
	if err != nil {
		log.Fatalf("err preparing statement %v", err)
	}
	return state
}
func userLogExists(db *sql.DB, login string) (bool, error) {
	query := `
    SELECT CASE 
               WHEN COUNT(*) > 0 THEN TRUE 
               ELSE FALSE 
           END AS exists
    FROM users
    WHERE login = ?;`

	var exists bool
	err := db.QueryRow(query, login).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	return exists, nil
}
func userNameExists(db *sql.DB, name string) (bool, error) {
	query := `
    SELECT CASE 
               WHEN COUNT(*) > 0 THEN TRUE 
               ELSE FALSE 
           END AS exists
    FROM users
    WHERE name = ?;`

	var exists bool
	err := db.QueryRow(query, name).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	return exists, nil
}
func getPasswordByLogin(db *sql.DB, login string) (string, error) {
	query := "SELECT password FROM users WHERE login = ?;"

	var password string
	err := db.QueryRow(query, login).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("пользователь с логином '%s' не найден", login)
		}
		return "", fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	return password, nil
}
