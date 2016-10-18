package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strings"
)

func publish(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stream := strings.Split(r.FormValue("name"), "-")
	nom_user := stream[0]
	db_mu.RLock()
	query, err := db.Query("SELECT username, password, status FROM admin WHERE username = ?", nom_user)
	db_mu.RUnlock()
	if err != nil {
		Warning.Println(err)
	}
	for query.Next() {
		var user, pass string
		var status int
		err = query.Scan(&user, &pass, &status)
		if err != nil {
			Warning.Println(err)
		}
		if user == r.FormValue("username") && pass == r.FormValue("password") && r.FormValue("call") == "publish" && status == 1 {
			fmt.Fprintf(w, "Server OK")
		} else {
			http.Error(w, "Internal Server Error", 500)
		}
	}
}

func onplay(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Internal Server Error", 500)
}