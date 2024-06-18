package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo-tui/tui"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

const dbName = "tasks.db/?parseTime=true"

func main() {
	db, _ := sql.Open("sqlite3", dbName)
	tasksRepo := tui.NewTasksRepository(db)
	err := tasksRepo.CreateNewDb()

	if err != nil {
		log.Printf("Error while creating DB")
		os.Exit(1)
	}

	m := tui.NewTuiModel(tasksRepo)
	p := tea.NewProgram(&m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Ups ...")
		os.Exit(1)
	}
}
