package main

import (
	"fmt"
	"os"
	"todo-tui/internal/views"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
  m := views.NewTuiModel()
  p := tea.NewProgram(&m, tea.WithAltScreen())

  if _, err := p.Run(); err != nil {
    fmt.Println("Ups ...")
    os.Exit(1)
  }
}
