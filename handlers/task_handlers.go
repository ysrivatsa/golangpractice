package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"to-do-api/models"
)

var Tasks []models.Task

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	}
	var task []models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(task)

	Tasks = append(Tasks, task...)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Success")

}

func UpdateTaskName(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Incorrect Method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var task models.New_Task

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i := 0; i < len(Tasks); i++ {
		if Tasks[i].TaskName == task.OldTaskName {
			Tasks[i].TaskName = task.NewTaskName

		}
	}

	json.NewEncoder(w).Encode(Tasks)

}

func UpdateStatusOfTask(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invaild Method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var _json models.Update_Status
	err := json.NewDecoder(r.Body).Decode(&_json)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(Tasks); i++ {
		if Tasks[i].TaskName == _json.TaskName {
			Tasks[i].Status = _json.Status
		}
	}
	json.NewEncoder(w).Encode(Tasks)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var _json models.Update_Status
	err := json.NewDecoder(r.Body).Decode(&_json)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := 0; i < len(Tasks); i++ {
		if Tasks[i].TaskName == _json.TaskName {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(Tasks)
}

func Landing(w http.ResponseWriter, r *http.Request) {
	enableCors(w)
	json.NewEncoder(w).Encode("Hellow WOrld ")
}
