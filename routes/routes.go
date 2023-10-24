package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/aksentijevicd1/go-postgres/controllers"
	"github.com/gorilla/mux"
)

var RegisterRoutes = func(Router *mux.Router) {
	l1 := log.New(os.Stdout, "api-task", log.LstdFlags)
	t := controllers.NewTasks(l1)

	l2 := log.New(os.Stdout, "api-category", log.LstdFlags)
	c := controllers.NewCategories(l2)

	l3 := log.New(os.Stdout, "api-status", log.LstdFlags)
	s := controllers.NewStatuses(l3)

	Router.HandleFunc("/task", t.CreateTask).Methods(http.MethodPost)
	Router.HandleFunc("/task", t.GetTasks).Methods(http.MethodGet)
	Router.HandleFunc("/task/{id:[0-9]+}", t.GetTaskById).Methods(http.MethodGet)
	Router.HandleFunc("/task/category/{category}", t.GetTasksByCategory).Methods(http.MethodGet)
	Router.HandleFunc("/task/status/{status}", t.GetTasksByStatus).Methods(http.MethodGet)
	Router.HandleFunc("/task/category/{category}/status/{status}", t.GetTasksByCategoryAndStatus).Methods(http.MethodGet)
	Router.HandleFunc("/task/{id:[0-9]+}", t.UpdateTask).Methods(http.MethodPut)
	Router.HandleFunc("/task/{id:[0-9]+}", t.RemoveTaskByID).Methods(http.MethodDelete)

	Router.HandleFunc("/category", c.CreateCategory).Methods(http.MethodPost)
	Router.HandleFunc("/category", c.GetCategories).Methods(http.MethodGet)
	Router.HandleFunc("/category/{id:[0-9]+}", c.GetCategoryByID).Methods(http.MethodGet)
	Router.HandleFunc("/category/{id:[0-9]+}", c.UpdateCategory).Methods(http.MethodPut)
	Router.HandleFunc("/category/{id:[0-9]+}", c.RemoveCategory).Methods(http.MethodDelete)

	Router.HandleFunc("/status", s.CreateStatus).Methods(http.MethodPost)
	Router.HandleFunc("/status", s.GetStatuses).Methods(http.MethodGet)
	Router.HandleFunc("/status/{id:[0-9]+}", s.GetStatusByID).Methods(http.MethodGet)
	Router.HandleFunc("/status/{id:[0-9]+}", s.UpdateStatus).Methods(http.MethodPut)
	Router.HandleFunc("/status/{id:[0-9]+}", s.RemoveStatus).Methods(http.MethodDelete)
}
