package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func DeleteTask(taskService TasksService) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		strId := mux.Vars(r)["id"]

		id, err := strconv.Atoi(strId)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
			return
		}

		if err = taskService.DeleteTask(id); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)
	}
}
