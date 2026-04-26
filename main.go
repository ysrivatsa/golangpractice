package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"to-do-api/handlers"
	"to-do-api/routes"
)

var db *sql.DB

func main() {
	fmt.Println("Server Started")
	handlers.ConnectDB()
	routes.RegisterRoutes()
	fmt.Println("Server is listening in port :8080")
	http.ListenAndServe(":8080", nil)

}
