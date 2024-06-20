package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
)

type focusedInput int

const (
	TITLE_INPUT focusedInput = iota
	DESC_INPUT
)

type createModel struct {
	repo         TasksRepository
	titleInput   textinput.Model
	descInput    textarea.Model
	focusedInput focusedInput
}

func NewCreateModal(repo TasksRepository) createModel {
	titleInput := textinput.New()
	titleInput.Placeholder = "Task title..."
	titleInput.Focus()
	titleInput.SetValue("")

	descInput := textarea.New()
	descInput.Placeholder = "Task description..."
	descInput.SetValue("")

	return createModel{
		repo:         repo,
		titleInput:   titleInput,
		descInput:    descInput,
		focusedInput: TITLE_INPUT,
	}
}

func (m createModel) Init() tea.Cmd {
	return nil
}

func (m createModel) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.titleInput.View(),
		m.descInput.View(),
	)
}

func (m createModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case changeToCreateViewMsg:
		m.titleInput.SetValue("")
		m.descInput.SetValue("")
		m.titleInput.Focus()
		m.descInput.Blur()

	case tea.KeyMsg:
		switch msg.String() {
		case "tab":

			switch m.focusedInput {
			case TITLE_INPUT:
				m.titleInput.Blur()
				m.descInput.Focus()
				m.focusedInput = DESC_INPUT

			case DESC_INPUT:
				m.descInput.Blur()
				m.titleInput.Focus()
				m.focusedInput = TITLE_INPUT
			}

		case "!":
			task := NewTask(uuid.New().String(), m.titleInput.Value(), m.descInput.Value(), false)
			m.repo.AddTask(task)
			return m, changeViewToListCmd

		case "esc":
			return m, changeViewToListCmd
		}
	}

	switch m.focusedInput {
	case TITLE_INPUT:
		m.titleInput, cmd = m.titleInput.Update(msg)

	case DESC_INPUT:
		m.descInput, cmd = m.descInput.Update(msg)
	}

	return m, cmd
}
