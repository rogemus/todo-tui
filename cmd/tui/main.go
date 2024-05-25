package main

import (
	"fmt"
	"os"
	"todo-tui/internal"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  m := internal.NewModel()
  p := tea.NewProgram(&m, tea.WithAltScreen())

  if _, err := p.Run(); err != nil {
    fmt.Println("Ups ...")
    os.Exit(1)
  }
}
