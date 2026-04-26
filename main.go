package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"to-do-api/routes"

	_ "github.com/lib/pq"
)

var db *sql.DB

func connectDB() {
	connStr := "host=todoapp.c9sgquaek4s9.ap-southeast-2.rds.amazonaws.com port=5432 user=postgres password=Srivatsa123 dbname=postgres sslmode=require"

	fmt.Println("Opening DB connection...")
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Before Ping...")
	err = db.Ping()
	fmt.Println("After Ping...")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected Successfully WITH THE DB")
}

func main() {
	fmt.Println("Server Started")
	connectDB()
	routes.RegisterRoutes()
	fmt.Println("Server is listening in port :8080")
	http.ListenAndServe(":8080", nil)

}
