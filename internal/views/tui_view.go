package views

import (
	"todo-tui/internal/consts"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
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
	help    help.Model
	keys    consts.KeyMap
	lastKey string
}

func NewTuiModel() TuiModel {
	return TuiModel{
		state:   listsView,
		lists:   NewListsModel(),
		details: NewDetailsModel(),
		keys:    consts.Keys,
		help:    help.New(),
	}
}

func (m TuiModel) Init() tea.Cmd {
	return nil
}

// Styles
var containerStyles = lipgloss.NewStyle().Padding(1)

func (m TuiModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Top,
		containerStyles.Render(
			lipgloss.JoinHorizontal(lipgloss.Top,
				m.lists.View(),
				m.details.View(),
			)),
		m.help.View(m.keys),
	)
}

func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.help.Width = msg.Width
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
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
