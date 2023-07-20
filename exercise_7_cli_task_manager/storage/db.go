package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDB() {
	dbConnection, err := gorm.Open(sqlite.Open("test.sqlite3"), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to connect database: %s", err.Error())
	}
	db = dbConnection
}

func SaveTask(description string) {
	task := &Task{
		Description: description,
		Done:        false,
	}
	db.Create(task)
	fmt.Println("Added " + "\"" + task.Description + "\"" + " to your task list.")

}

func GetList() {
	tasks := make([]Task, 0)
	result := db.Where("done = ?", false).Find(&tasks)

	if result.Error != nil {
		log.Panicf("failed to get list: %v", result.Error.Error())
	}

	if len(tasks) > 0 {
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Description)
		}
	} else {
		fmt.Println("You have no tasks")
	}
}

func CompleteTask(description string) {
	tasks := make([]Task, 0)
	result := db.Where("description = ?", description).Find(&tasks)

	if result.Error != nil {
		log.Panicf("failed to complete task: %v", result.Error.Error())
	}

	if len(tasks) > 0 {
		db.Delete(tasks)
		fmt.Println("You have completed the " + "\"" + description + "\"" + " task.")

	} else {
		fmt.Printf("No tasks to complete: %s\n", description)
	}

}
