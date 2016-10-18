package main

import ()

var username string
var password string

// This function could be used to access to a Database for user/pass authentication procedure
func authentication(user, pass string) bool {
	var tipo int
	db_mu.RLock()
	query2, err := db.Query("SELECT username, password, type FROM admin WHERE username = ?", user)
	db_mu.RUnlock()
	if err != nil {
		Error.Println(err)
	}
	for query2.Next() {
		err = query2.Scan(&username, &password, &tipo)
		if err != nil {
			Error.Println(err)
		}
	}
	if user == username && pass == password && tipo == 0 {
		return true
	} else {
		return false
	}
}

// This function could be used to access to a Database for user/pass authentication procedure
func authentication_admin(user, pass string) bool {
	var tipo int
	db_mu.RLock()
	query2, err := db.Query("SELECT username, password, type FROM admin WHERE username = ?", user)
	db_mu.RUnlock()
	if err != nil {
		Error.Println(err)
	}
	for query2.Next() {
		err = query2.Scan(&username, &password, &tipo)
		if err != nil {
			Error.Println(err)
		}
	}
	if user == username && pass == password && tipo == 1 {
		return true
	} else {
		return false
	}
}
