package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"taskmanager/models"

	"github.com/gorilla/mux"
)

var tasks = make(map[int]models.Task)
var taskIDCounter = 1


func GetTasks(w http.ResponseWriter, r *http.Request) {
	taskList := make([]models.Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	json.NewEncoder(w).Encode(taskList)
}


func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)
	task.ID = taskIDCounter
	taskIDCounter++
	tasks[task.ID] = task
	json.NewEncoder(w).Encode(task)
}


func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if task, exists := tasks[id]; exists {
		json.NewEncoder(w).Encode(task)
		return
	}
	http.NotFound(w, r)
}


func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if _, exists := tasks[id]; exists {
		var updatedTask models.Task
		json.NewDecoder(r.Body).Decode(&updatedTask)
		updatedTask.ID = id
		tasks[id] = updatedTask
		json.NewEncoder(w).Encode(updatedTask)
		return
	}
	http.NotFound(w, r)
}


func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	if _, exists := tasks[id]; exists {
		delete(tasks, id)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	http.NotFound(w, r)
}
