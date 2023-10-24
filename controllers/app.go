package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aksentijevicd1/go-postgres/models"
	"github.com/aksentijevicd1/go-postgres/utils"
	"github.com/gorilla/mux"
)

const (
	DateTime = "2006-01-02 15:04:05"
)

type Tasks struct {
	l *log.Logger
}

func NewTasks(l *log.Logger) *Tasks {
	return &Tasks{l}
}

type Categories struct {
	l *log.Logger
}

func NewCategories(l *log.Logger) *Categories {
	return &Categories{l}
}

type Statuses struct {
	l *log.Logger
}

func NewStatuses(l *log.Logger) *Statuses {
	return &Statuses{l}
}

func (t *Tasks) CreateTask(w http.ResponseWriter, r *http.Request) {
	newTask := &models.Task{}
	//cita telo u newTask
	err := utils.ParseBody(r, newTask)

	if err != nil {
		t.l.Printf("Error while parsing %s", err)
		return
	}
	//iako se ne prosledjuje neophodno je
	newTask.CreatedOn = time.Now()
	newTask.EditedOn = time.Now()

	task := newTask.CreateTask()
	res, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (t *Tasks) GetTasks(w http.ResponseWriter, r *http.Request) {
	allTasks := models.GetTasks()
	res, err := json.Marshal(allTasks)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (t *Tasks) GetTaskById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		t.l.Println("Error while converting to int")
		return
	}
	task := models.GetTaskById(id)
	res, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (t *Tasks) GetTasksByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	tasks := models.GetTasksByCategory(category)
	res, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (t *Tasks) GetTasksByStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	status := vars["status"]
	tasks := models.GetTasksByStatus(status)
	res, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (t *Tasks) GetTasksByCategoryAndStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]
	status := vars["status"]
	tasks := models.GetTasksByCategoryAndStatus(status, category)
	res, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (t *Tasks) UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		t.l.Println("Error while converting to int")
		return
	}

	existingTask := models.GetTaskById(id)
	updatedTask := &models.Task{}
	err = utils.ParseBody(r, updatedTask)
	updatedTask.CreatedOn = time.Now()
	updatedTask.EditedOn = time.Now()

	if updatedTask.Owner != "" {
		existingTask.Owner = updatedTask.Owner
	}
	if updatedTask.Realm != "" {
		existingTask.Realm = updatedTask.Realm
	}
	if updatedTask.Name != "" {
		existingTask.Name = updatedTask.Name
	}

	existingTask.Enabled = updatedTask.Enabled

	if updatedTask.StatusID != 0 {
		existingTask.StatusID = updatedTask.StatusID
	}
	if updatedTask.CategoryID != 0 {
		existingTask.CategoryID = updatedTask.CategoryID
	}
	if updatedTask.AlertID != 0 {
		existingTask.AlertID = updatedTask.AlertID
	}
	if updatedTask.AssignedTo != "" {
		existingTask.AssignedTo = updatedTask.AssignedTo
	}
	if updatedTask.ExternalID != 0 {
		existingTask.ExternalID = updatedTask.ExternalID
	}
	if updatedTask.MetadataSchema != "" {
		existingTask.MetadataSchema = updatedTask.MetadataSchema
	}
	if updatedTask.Metadata != "" {
		existingTask.Metadata = updatedTask.Metadata
	}
	if updatedTask.CreatedBy != "" {
		existingTask.CreatedBy = updatedTask.CreatedBy
	}
	if updatedTask.EditedByName != "" {
		existingTask.EditedByName = updatedTask.EditedByName
	}
	if updatedTask.EditedBy != "" {
		existingTask.EditedBy = updatedTask.EditedBy
	}
	existingTask.EditedOn = time.Now()

	if err != nil {
		t.l.Printf("Error while parsing %s", err)
		return
	}

	task, err := models.UpdateTask(id, existingTask)
	if err != nil {
		t.l.Printf("Error while updating task!, %s", err)
	}

	res, err := json.Marshal(task)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (t *Tasks) RemoveTaskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)

	if err != nil {

		http.Error(w, "Unable to parse to int", http.StatusBadRequest)
		return
	}
	task := models.RemoveTaskByID(id)
	res, err := json.Marshal(task)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (c *Categories) CreateCategory(w http.ResponseWriter, r *http.Request) {
	newCategory := &models.Category{}
	err := utils.ParseBody(r, newCategory)

	if err != nil {
		c.l.Printf("Error while parsing %s", err)
		return
	}

	newCategory.CreatedOn = time.Now()
	newCategory.EditedOn = time.Now()
	category := newCategory.CreateCategory()
	res, err := json.Marshal(category)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (s *Statuses) CreateStatus(w http.ResponseWriter, r *http.Request) {
	newStatus := &models.Status{}
	err := utils.ParseBody(r, newStatus)

	if err != nil {
		s.l.Printf("Error while parsing %s", err)
		return
	}

	newStatus.CreatedOn = time.Now()
	newStatus.EditedOn = time.Now()
	task := newStatus.CreateStatus()
	res, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *Categories) GetCategories(w http.ResponseWriter, r *http.Request) {
	allCategories := models.GetCategories()
	res, err := json.Marshal(allCategories)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (s *Statuses) GetStatuses(w http.ResponseWriter, r *http.Request) {
	allStatuses := models.GetStatuses()
	res, err := json.Marshal(allStatuses)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *Categories) GetCategoryByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.l.Println("Error while converting to int")
		return
	}
	category := models.GetCategoryByID(id)
	res, err := json.Marshal(category)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (s *Statuses) GetStatusByID(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		s.l.Println("Error while converting to int")
		return
	}
	status := models.GetStatusByID(id)
	res, err := json.Marshal(status)
	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (c *Categories) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.l.Println("Error while converting to int")
		return
	}
	existingCategory := models.GetCategoryByID(id)
	newCategory := &models.Category{}
	err = utils.ParseBody(r, newCategory)

	if err != nil {
		c.l.Printf("Error while parsing %s", err)
		return
	}
	newCategory.CreatedOn = time.Now()
	newCategory.EditedOn = time.Now()

	//update samo ono sto je popunjeno

	if newCategory.Owner != "" {
		existingCategory.Owner = newCategory.Owner
	}
	if newCategory.Realm != "" {
		existingCategory.Realm = newCategory.Realm
	}

	if newCategory.Category != "" {
		existingCategory.Category = newCategory.Category
	}

	if newCategory.MetadataSchema != "" {
		existingCategory.MetadataSchema = newCategory.MetadataSchema
	}
	if newCategory.Metadata != "" {
		existingCategory.Metadata = newCategory.Metadata
	}
	if newCategory.CreatedByName != "" {
		existingCategory.CreatedByName = newCategory.CreatedByName
	}
	if newCategory.CreatedBy != "" {
		existingCategory.CreatedBy = newCategory.CreatedBy
	}
	if newCategory.EditedByName != "" {
		existingCategory.EditedByName = newCategory.EditedByName
	}
	if newCategory.EditedBy != "" {
		existingCategory.EditedBy = newCategory.EditedBy
	}
	existingCategory.EditedOn = time.Now()

	category, err := models.UpdateCategory(id, existingCategory)
	if err != nil {
		c.l.Printf("Error while updating category!, %s", err)
	}

	res, err := json.Marshal(category)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (s *Statuses) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)

	if err != nil {
		s.l.Println("Error while converting to int")
		return
	}
	existingStatus := models.GetStatusByID(id)
	newStatus := &models.Status{}
	err = utils.ParseBody(r, newStatus)

	if err != nil {
		s.l.Printf("Error while parsing %s", err)
		return
	}
	newStatus.CreatedOn = time.Now()
	newStatus.EditedOn = time.Now()

	if newStatus.Owner != "" {
		existingStatus.Owner = newStatus.Owner
	}
	if newStatus.Realm != "" {
		existingStatus.Realm = newStatus.Realm
	}

	if newStatus.Status != "" {
		existingStatus.Status = newStatus.Status
	}

	if newStatus.MetadataSchema != "" {
		existingStatus.MetadataSchema = newStatus.MetadataSchema
	}
	if newStatus.Metadata != "" {
		existingStatus.Metadata = newStatus.Metadata
	}
	if newStatus.CreatedByName != "" {
		existingStatus.CreatedByName = newStatus.CreatedByName
	}
	if newStatus.CreatedBy != "" {
		existingStatus.CreatedBy = newStatus.CreatedBy
	}
	if newStatus.EditedByName != "" {
		existingStatus.EditedByName = newStatus.EditedByName
	}
	if newStatus.EditedBy != "" {
		existingStatus.EditedBy = newStatus.EditedBy
	}
	existingStatus.EditedOn = time.Now()

	status, err := models.UpdateStatus(id, existingStatus)
	if err != nil {
		s.l.Printf("Error while updating status!, %s", err)
	}

	res, err := json.Marshal(status)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (c *Categories) RemoveCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)

	if err != nil {

		http.Error(w, "Unable to parse to int", http.StatusBadRequest)
		return
	}
	RemoveCategory := models.RemoveCategory(id)
	res, err := json.Marshal(RemoveCategory)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func (s *Statuses) RemoveStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)

	if err != nil {

		http.Error(w, "Unable to parse to int", http.StatusBadRequest)
		return
	}
	status := models.RemoveStatus(id)
	res, err := json.Marshal(status)

	if err != nil {
		http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
