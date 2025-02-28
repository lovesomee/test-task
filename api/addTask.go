package api

import (
	"encoding/json"
	"log"
	"net/http"
	"skillsrock-test/models"
)

func AddTask(taskService TasksService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var tasks models.Tasks
		if err := json.NewDecoder(r.Body).Decode(&tasks); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}

		if err := taskService.AddTask(tasks); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}

		w.WriteHeader(http.StatusCreated)
	}
}
