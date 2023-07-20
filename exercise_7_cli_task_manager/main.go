/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/feynmaz/gophercises/exercise_7_cli_task_manager/cmd"
	"github.com/feynmaz/gophercises/exercise_7_cli_task_manager/storage"
)

func main() {
	storage.InitDB()
	storage.DoMigrations()
	cmd.Execute()
}
