package views

import (
	"todo-tui/internal/consts"
	"todo-tui/internal/storage"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CreateViewModel struct {
	repo  storage.TasksRepository
	title textinput.Model
	desc  textarea.Model
	keys  consts.KeyMap
}

func NewCreateView(repo storage.TasksRepository) CreateViewModel {
	title := textinput.New()
	title.Placeholder = "Task title ..."
	desc := textarea.New()
	desc.Placeholder = "Task description ..."
	title.Focus()

	return CreateViewModel{
		repo:  repo,
		title: title,
		desc:  desc,
		keys:  consts.Keys,
	}
}

func (m CreateViewModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CreateViewModel) Update(msg tea.Msg) (CreateViewModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.ChangeFocus):
      /* TODO Add submit and make it more generic */
			if m.title.Focused() {
				m.title.Blur()
				m.desc.Focus()
			} else {
				m.desc.Blur()
				m.title.Focus()
			}
		}
	}

	if m.title.Focused() {
		m.title, cmd = m.title.Update(msg)
	} else {
		m.desc, cmd = m.desc.Update(msg)
	}

	return m, cmd
}

var dialogStyles = lipgloss.NewStyle().
	Background(lipgloss.Color("10")).
	Padding(2)

var frameStyles = lipgloss.NewStyle().
	Align(lipgloss.Position(0.5)).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(consts.ColSubtle)

var fieldStyles = lipgloss.NewStyle().
	MarginBottom(1)

var labelStyles = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(consts.ColSubtle).
	BorderTop(false).
	BorderLeft(false).
	BorderRight(false).
	PaddingLeft(1).
	Width(50).
	Bold(true).
	Foreground(consts.Highlight)

func (m CreateViewModel) View() string {
	return frameStyles.Render(lipgloss.JoinVertical(
		lipgloss.Top,
		fieldStyles.Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				labelStyles.Render("Task Title"),
				m.title.View(),
			),
		),
		fieldStyles.Render(
			lipgloss.JoinVertical(
				lipgloss.Top,
				labelStyles.Render("Task Description"),
				m.desc.View(),
			),
		),
	))
}
