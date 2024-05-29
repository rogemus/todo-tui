package views

import (
	"todo-tui/internal/models"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DetailsViewModel struct {
	sel models.Item
}

func NewDetailsModel() DetailsViewModel {
	return DetailsViewModel{}
}

func (m DetailsViewModel) Init() tea.Cmd {
	return nil
}

func (m DetailsViewModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Top, "Details", m.sel.Title, m.sel.Description)
}

func (m DetailsViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case SelectedItemMsg:
		m.sel = msg.Item
	}
	return m, nil
}
