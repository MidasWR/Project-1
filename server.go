package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/main", HomePage)
	http.HandleFunc("/main/reg", RegPage)
	http.HandleFunc("/main/auth", AuthPage)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home page")

}
func RegPage(w http.ResponseWriter, r *http.Request) {
	var rUser RegUser
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Println(1)
		panic(err)
	}
	username := r.URL.Query().Get("username")
	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")
	rUser.username = username
	rUser.login = login
	rUser.password = password
	q, err := userLogExists(db, login)
	if err != nil {
		log.Println("2.1")
		RegPage(w, r)
	}
	if q {
		fmt.Fprintln(w, "This login already used now")
		RegPage(w, r)
	}
	f, err := userNameExists(db, username)
	if err != nil {
		log.Println("2.2")
		RegPage(w, r)
	}
	if f {
		fmt.Fprintln(w, "This username already used now")
		RegPage(w, r)
	}
	state := stateDB(db)
	RegNewUser(state, rUser)
	http.HandleFunc("/private", privatePage)
	http.Redirect(w, r, "/private", http.StatusFound)
}
func AuthPage(w http.ResponseWriter, r *http.Request) {
	login := r.URL.Query().Get("login")
	password := r.URL.Query().Get("password")
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Println("err")
	}
	q, err := getPasswordByLogin(db, login)
	if err != nil {
		log.Println("err")
	}
	if password != q {
		fmt.Fprintln(w, "Login password wrong")
		AuthPage(w, r)
	}
	http.HandleFunc("/private", privatePage)
	http.Redirect(w, r, "/private", http.StatusFound)
}
func privatePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this page is private")
}
