package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state int

const (
	listsView state = iota
	detailsView
)

type TuiModel struct {
	state   state
	lists   tea.Model
	details tea.Model
}

func NewTuiModel() TuiModel {
	return TuiModel{
		state:   listsView,
		lists:   NewListsModel(),
		details: NewDetailsModel(),
	}
}

func (m TuiModel) Init() tea.Cmd {
	return nil
}

// Styles
var containerStyles = lipgloss.NewStyle().Padding(1)

func (m TuiModel) View() string {
	return containerStyles.Render(
		lipgloss.JoinHorizontal(lipgloss.Top,
			m.lists.View(),
			m.details.View(),
		))
}

func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			return m, tea.Quit
		default:
			switch m.state {
			case listsView:
				m.lists, cmd = m.lists.Update(msg)
			case detailsView:
				m.details, cmd = m.details.Update(msg)
			}
		}
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
