package models

import (
	"fmt"
	"time"

	"github.com/aksentijevicd1/go-postgres/database"
	"github.com/jinzhu/gorm"
)

type Status struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Owner          string    `json:"owner" gorm:"column:owner"`
	Realm          string    `json:"realm" gorm:"column:realm"`
	Status         string    `json:"status" gorm:"column:status"`
	MetadataSchema string    `json:"metadataschema" gorm:"column:metadataschema"`
	Metadata       string    `json:"metadata" gorm:"column:metadata"`
	CreatedByName  string    `json:"createdbyname" gorm:"column:createdbyname"`
	CreatedBy      string    `json:"createdby" gorm:"column:createdby"`
	CreatedOn      time.Time `json:"createdon" gorm:"column:createdon"`
	EditedByName   string    `json:"editedbyname" gorm:"column:editedbyname"`
	EditedBy       string    `json:"editedby" gorm:"column:editedby"`
	EditedOn       time.Time `json:"editedon" gorm:"column:editedon"`
}

type Task struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Owner          string    `json:"owner" gorm:"column:owner"`
	Realm          string    `json:"realm" gorm:"column:realm"`
	Name           string    `json:"name" gorm:"column:name"`
	Enabled        bool      `json:"enabled" gorm:"column:enabled"`
	StatusID       int       `json:"statusid" gorm:"column:statusid"`
	CategoryID     int       `json:"categoryid" gorm:"column:categoryid"`
	AlertID        int       `json:"alertid" gorm:"column:alertid"`
	AssignedTo     string    `json:"assignedto" gorm:"column:assignedto"`
	ExternalID     int       `json:"externalid" gorm:"column:externalid"`
	MetadataSchema string    `json:"metadataschema" gorm:"column:metadataschema"`
	Metadata       string    `json:"metadata" gorm:"column:metadata"`
	CreatedBy      string    `json:"createdby" gorm:"column:createdby"`
	CreatedOn      time.Time `json:"createdon" gorm:"column:createdon"`
	EditedByName   string    `json:"editedbyname" gorm:"column:editedbyname"`
	EditedBy       string    `json:"editedby" gorm:"column:editedby"`
	EditedOn       time.Time `json:"editedon" gorm:"column:editedon"`
}

type Category struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	Owner          string    `json:"owner" gorm:"column:owner"`
	Realm          string    `json:"realm" gorm:"column:realm"`
	Category       string    `json:"category" gorm:"column:category"`
	MetadataSchema string    `json:"metadataschema" gorm:"column:metadataschema"`
	Metadata       string    `json:"metadata" gorm:"column:metadata"`
	CreatedByName  string    `json:"createdbyname" gorm:"column:createdbyname"`
	CreatedBy      string    `json:"createdby" gorm:"column:createdby"`
	CreatedOn      time.Time `json:"createdon" gorm:"column:createdon"`
	EditedByName   string    `json:"editedbyname" gorm:"column:editedbyname"`
	EditedBy       string    `json:"editedby" gorm:"column:editedby"`
	EditedOn       time.Time `json:"editedon" gorm:"column:editedon"`
}

var db *gorm.DB

func init() {
	database.Connect()
	db = database.GetDB()
	db.AutoMigrate(&Task{}, &Category{}, &Status{})
	addValues()
}

func addValues() {

	if isTableEmpty(&Task{}) {
		db := db.Exec(`
		INSERT INTO tasks (Owner, Realm, Name, Enabled, StatusID, CategoryID, AlertID, AssignedTo, ExternalID, MetadataSchema, Metadata, CreatedBy, CreatedOn, EditedByName, EditedBy, EditedOn)
		VALUES
			('John Doe', 'Realm1', 'Task 1', true, 1, 1, 101, 'User A', 1001, 'Schema1', 'Metadata1', 'John Doe', CURRENT_TIMESTAMP, 'Jane Smith', 'User B', CURRENT_TIMESTAMP),
			('Alice Johnson', 'Realm2', 'Task 2', true, 2, 2, 102, 'User C', 1002, 'Schema2', 'Metadata2', 'Alice Johnson', CURRENT_TIMESTAMP, 'Bob Wilson', 'User D', CURRENT_TIMESTAMP),
			('Ella Davis', 'Realm3', 'Task 3', false, 3, 3, 103, 'User E', 1003, 'Schema3', 'Metadata3', 'Ella Davis', CURRENT_TIMESTAMP, 'David Brown', 'User F', CURRENT_TIMESTAMP),
			('Michael Lee', 'Realm4', 'Task 4', true, 4, 4, 104, 'User G', 1004, 'Schema4', 'Metadata4', 'Michael Lee', CURRENT_TIMESTAMP, 'Sarah Taylor', 'User H', CURRENT_TIMESTAMP),
			('William White', 'Realm5', 'Task 5', true, 5, 5, 105, 'User I', 1005, 'Schema5', 'Metadata5', 'William White', CURRENT_TIMESTAMP, 'Olivia Turner', 'User J', CURRENT_TIMESTAMP),
			('Sophia Johnson', 'Realm6', 'Task 6', false, 6, 6, 106, 'User K', 1006, 'Schema6', 'Metadata6', 'Sophia Johnson', CURRENT_TIMESTAMP, 'James Smith', 'User L', CURRENT_TIMESTAMP),
			('Oliver Wilson', 'Realm7', 'Task 7', true, 7, 7, 107, 'User M', 1007, 'Schema7', 'Metadata7', 'Oliver Wilson', CURRENT_TIMESTAMP, 'Emma Davis', 'User N', CURRENT_TIMESTAMP),
			('Ava Martinez', 'Realm8', 'Task 8', true, 8, 8, 108, 'User O', 1008, 'Schema8', 'Metadata8', 'Ava Martinez', CURRENT_TIMESTAMP, 'Daniel Garcia', 'User P', CURRENT_TIMESTAMP),
			('Liam Anderson', 'Realm9', 'Task 9', false, 9, 9, 109, 'User Q', 1009, 'Schema9', 'Metadata9', 'Liam Anderson', CURRENT_TIMESTAMP, 'Mia Hernandez', 'User R', CURRENT_TIMESTAMP),
			('Noah Smith', 'Realm10', 'Task 10', true, 10, 10, 110, 'User S', 1010, 'Schema10', 'Metadata10', 'Noah Smith', CURRENT_TIMESTAMP, 'Sophia Johnson', 'User T', CURRENT_TIMESTAMP)
		`)
		if db.Error != nil {
			fmt.Println("Error in executing sql")
			return
		}
	}

	if isTableEmpty(&Category{}) {

		db := db.Exec(`
		INSERT INTO categories (Owner, Realm, Category, MetadataSchema, Metadata, CreatedByName, CreatedBy, CreatedOn, EditedByName, EditedBy, EditedOn)
		VALUES
			('weee Doe', 'Realm1', 'category1', 'Schema1', 'Metadata1', 'John Doe', 'user1', CURRENT_TIMESTAMP, 'Jane Smith', 'user2', CURRENT_TIMESTAMP),
			('dada Johnson', 'Realm2', 'category2', 'Schema2', 'Metadata2', 'Alice Johnson', 'user3', CURRENT_TIMESTAMP, 'Bob Wilson', 'user4', CURRENT_TIMESTAMP),
			('Ellsssa Davis', 'Realm3', 'category3', 'Schema3', 'Metadata3', 'Ella Davis', 'user5', CURRENT_TIMESTAMP, 'David Brown', 'user6', CURRENT_TIMESTAMP),
			('Mfieechaael Lee', 'Realm4', 'category4', 'Schema4', 'Metadata4', 'Michael Lee', 'user7', CURRENT_TIMESTAMP, 'Sarah Taylor', 'user8', CURRENT_TIMESTAMP),
			('Waasdiam White', 'Realm5', 'category5', 'Schema5', 'Metadata5', 'William White', 'user9', CURRENT_TIMESTAMP, 'Olivia Turner', 'user10', CURRENT_TIMESTAMP),
			('Soraja Johnson', 'Realm6', 'category6', 'Schema6', 'Metadata6', 'Sophia Johnson', 'user11', CURRENT_TIMESTAMP, 'James Smith', 'user12', CURRENT_TIMESTAMP),
			('Mariahg Wilson', 'Realm7', 'category7', 'Schema7', 'Metadata7', 'Oliver Wilson', 'user13', CURRENT_TIMESTAMP, 'Emma Davis', 'user14', CURRENT_TIMESTAMP),
			('adas Martinez', 'Realm8', 'category8', 'Schema8', 'Metadata8', 'Ava Martinez', 'user15', CURRENT_TIMESTAMP, 'Daniel Garcia', 'user16', CURRENT_TIMESTAMP),
			('Lifaeem Anderson', 'Realm9', 'category9', 'Schema9', 'Metadata9', 'Liam Anderson', 'user17', CURRENT_TIMESTAMP, 'Mia Hernandez', 'user18', CURRENT_TIMESTAMP),
			('uioos Smith', 'Realm10', 'category10', 'Schema10', 'Metadata10', 'Noah Smith', 'user19', CURRENT_TIMESTAMP, 'Sophia Johnson', 'user20', CURRENT_TIMESTAMP)
		`)
		if db.Error != nil {
			fmt.Println("Error in executing sql")
			return
		}
	}

	if isTableEmpty(&Status{}) {
		db := db.Exec(`
		INSERT INTO statuses (Owner, Realm, Status, MetadataSchema, Metadata, CreatedByName, CreatedBy, CreatedOn, EditedByName, EditedBy, EditedOn)
		VALUES
			('John Doe', 'Realm1', 'status1', 'Schema1', 'Metadata1', 'John Doe', 'user1', CURRENT_TIMESTAMP, 'Jane Smith', 'user2', CURRENT_TIMESTAMP),
			('Alice Johnson', 'Realm2', 'status2', 'Schema2', 'Metadata2', 'Alice Johnson', 'user3', CURRENT_TIMESTAMP, 'Bob Wilson', 'user4', CURRENT_TIMESTAMP),
			('Ella Davis', 'Realm3', 'status3', 'Schema3', 'Metadata3', 'Ella Davis', 'user5', CURRENT_TIMESTAMP, 'David Brown', 'user6', CURRENT_TIMESTAMP),
			('Michael Lee', 'Realm4', 'status4', 'Schema4', 'Metadata4', 'Michael Lee', 'user7', CURRENT_TIMESTAMP, 'Sarah Taylor', 'user8', CURRENT_TIMESTAMP),
			('William White', 'Realm5', 'status5', 'Schema5', 'Metadata5', 'William White', 'user9', CURRENT_TIMESTAMP, 'Olivia Turner', 'user10', CURRENT_TIMESTAMP),
			('Sophia Johnson', 'Realm6', 'status6', 'Schema6', 'Metadata6', 'Sophia Johnson', 'user11', CURRENT_TIMESTAMP, 'James Smith', 'user12', CURRENT_TIMESTAMP),
			('Oliver Wilson', 'Realm7', 'status7', 'Schema7', 'Metadata7', 'Oliver Wilson', 'user13', CURRENT_TIMESTAMP, 'Emma Davis', 'user14', CURRENT_TIMESTAMP),
			('Ava Martinez', 'Realm8', 'status8', 'Schema8', 'Metadata8', 'Ava Martinez', 'user15', CURRENT_TIMESTAMP, 'Daniel Garcia', 'user16', CURRENT_TIMESTAMP),
			('Liam Anderson', 'Realm9', 'status9', 'Schema9', 'Metadata9', 'Liam Anderson', 'user17', CURRENT_TIMESTAMP, 'Mia Hernandez', 'user18', CURRENT_TIMESTAMP),
			('Noah Smith', 'Realm10', 'status10', 'Schema10', 'Metadata10', 'Noah Smith', 'user19', CURRENT_TIMESTAMP, 'Sophia Johnson', 'user20', CURRENT_TIMESTAMP)
		`)
		if db.Error != nil {
			fmt.Println("Error in executing sql")
			return
		}
	}

}

func isTableEmpty(table interface{}) bool {
	var count int
	db.Model(table).Count(&count)
	return count == 0
}

func (t *Task) CreateTask() *Task {
	db.Table("tasks").Create(&t)
	return t
}

func GetTasks() []Task {
	var tasks []Task
	db.Table("tasks").Find(&tasks)
	return tasks
}

func GetTaskById(Id int) Task {
	var newTask Task
	db.Table("tasks").Where("ID=?", Id).Find(&newTask)
	return newTask
}

func GetTasksByCategory(CategoryName string) []Task {
	var newTasks []Task
	var category Category
	db.Table("categories").Where("category = ?", CategoryName).First(&category)
	db.Table("tasks").Where("categoryid = ?", category.ID).Find(&newTasks)
	return newTasks

}

func GetTasksByStatus(StatusName string) []Task {
	var newTasks []Task
	var status Status
	db.Table("statuses").Where("status = ?", StatusName).First(&status)
	db.Table("tasks").Where("statusid = ?", status.ID).Find(&newTasks)
	return newTasks

}

func GetTasksByCategoryAndStatus(StatusName string, CategoryName string) []Task {
	var newTasks []Task
	var status Status
	var category Category
	db.Table("statuses").Where("status = ?", StatusName).First(&status)
	db.Table("categories").Where("category = ?", CategoryName).First(&category)
	db.Table("tasks").Where("statusid = ? AND categoryid = ?", status.ID, category.ID).Find(&newTasks)
	return newTasks
}

func UpdateTask(ID int, newTask Task) (Task, error) {
	var decoyTask Task
	// Trazi ukoliko postoji task.
	if err := db.Table("tasks").Where("ID = ?", ID).First(&decoyTask).Error; err != nil {
		return Task{}, err
	}
	if err := db.Save(&newTask).Error; err != nil {
		return Task{}, err
	}

	return newTask, nil
}

func RemoveTaskByID(Id int) Task {
	var newTask Task
	db.Table("tasks").Where("ID=?", Id).Delete(&newTask)
	return newTask
}

func (c *Category) CreateCategory() *Category {
	db.Table("categories").Create(&c)
	return c
}

func GetCategories() []Category {
	var categories []Category
	db.Table("categories").Find(&categories)
	return categories
}

func GetCategoryByID(Id int) Category {
	var newCategory Category
	db.Table("categories").Where("ID=?", Id).Find(&newCategory)
	return newCategory
}

func UpdateCategory(Id int, newCategory Category) (Category, error) {
	var decoyCategory Category

	if err := db.Table("categories").Where("ID=?", Id).First(&decoyCategory).Error; err != nil {
		return Category{}, err
	}
	if err := db.Save(&newCategory).Error; err != nil {
		return Category{}, err
	}

	return newCategory, nil
}

func RemoveCategory(Id int) Category {
	var category Category
	db.Table("categories").Where("ID=?", Id).Delete(&category)
	return category
}

func (s *Status) CreateStatus() *Status {
	db.Table("statuses").Create(&s)
	return s
}

func GetStatuses() []Status {
	var statuses []Status
	db.Table("statuses").Find(&statuses)
	return statuses
}

func GetStatusByID(Id int) Status {
	var status Status
	db.Table("statuses").Where("ID=?", Id).Find(&status)
	return status
}

func UpdateStatus(Id int, newStatus Status) (Status, error) {
	var decoyStatus Status
	// Trazi ukoliko postoji task.
	if err := db.Table("statuses").Where("ID = ?", Id).First(&decoyStatus).Error; err != nil {
		return Status{}, err
	}
	if err := db.Save(&newStatus).Error; err != nil {
		return Status{}, err
	}

	return newStatus, nil
}

func RemoveStatus(Id int) Status {
	var status Status
	db.Table("statuses").Where("ID=?", Id).Delete(&status)
	return status
}
