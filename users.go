package main

import "database/sql"

type User struct {
	id       int
	username string
	login    string
	password string
}
type RegUser struct {
	username string
	login    string
	password string
}

func RegNewUser(state *sql.Stmt, r RegUser) User {
	var user User
	state.Exec(`INSERT INTO users (username,login,password) VALUES (?,?,?)`, r.username, r.login, r.password)
	user.username = r.username
	user.login = r.login
	user.password = r.password
	return user
}
