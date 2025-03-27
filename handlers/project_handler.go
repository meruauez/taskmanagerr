package handlers

import (
	"encoding/json"
	"net/http"
	"taskmanager/models"
)

var projects = make(map[int]models.Project)
var projectIDCounter = 1

func GetProjects(w http.ResponseWriter, r *http.Request) {
	projectList := make([]models.Project, 0, len(projects))
	for _, project := range projects {
		projectList = append(projectList, project)
	}
	json.NewEncoder(w).Encode(projectList)
}


func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project models.Project
	json.NewDecoder(r.Body).Decode(&project)
	project.ID = projectIDCounter
	projectIDCounter++
	projects[project.ID] = project
	json.NewEncoder(w).Encode(project)
}
