package main

import (
	"fmt"
	"net/http"
	"to-do-api/routes"
)

func main() {
	fmt.Println("Server Started")
	routes.RegisterRoutes()
	fmt.Println("Server is listening in port :8080")
	http.ListenAndServe(":8080", nil)

}
