package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"skillsrock-test/config"
)

func NewServer(cfg config.Settings, tasks TasksService) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/ping", Ping()).Methods(http.MethodGet)
	router.HandleFunc("/tasks", AddTask(tasks)).Methods(http.MethodPost)
	router.HandleFunc("/tasks", UpdateTask(tasks)).Methods(http.MethodPut)
	router.HandleFunc("/tasks/{id}", DeleteTask(tasks)).Methods(http.MethodDelete)

	return &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", cfg.Port),
	}
}
