package routes

import (
	"net/http"
	"to-do-api/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.Landing)
	http.HandleFunc("/tasks", handlers.GetTasks)
	http.HandleFunc("/tasks/create", handlers.CreateTask)
	http.HandleFunc("/tasks/updateTaskName", handlers.UpdateTaskName)
	http.HandleFunc("/tasks/updateStatus", handlers.UpdateStatusOfTask)
	http.HandleFunc("/tasks/deleteTask", handlers.DeleteTask)
}

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
