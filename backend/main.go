package main

import (
	"log"
	"net/http"

	"github.com/CrossStack-Q/Vacation-App/api"
	"github.com/CrossStack-Q/Vacation-App/db"
)

func main() {
	// Connect to database

	dbConn := db.Connect()
	defer dbConn.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET, PUT, DELETE")
		if r.Method == http.MethodOptions {
			return
		}
		http.DefaultServeMux.ServeHTTP(w, r)
	})

	// Register routes
	http.HandleFunc("/signup", api.SignupHandler)
	http.HandleFunc("/login", api.LoginHandler)
	http.HandleFunc("/logout", api.LogoutHandler)
	http.HandleFunc("/data", api.DataHandler)
	http.HandleFunc("/hotels", api.HotelsLocateHandler)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
