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
	CONFIRM_BTN
	CANCEL_BTN
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

// Colors
var subtleColor = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
var highlightColor = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}

// TODO accent colors
var accentColorBg = lipgloss.Color("12")
var accentColorFg = lipgloss.Color("11")
var secondaryAccentColorBg = lipgloss.Color("12")
var secondaryAccentColorFg = lipgloss.Color("12")

// Styles
var boxStyles = lipgloss.NewStyle().
	Padding(1).
	Border(lipgloss.RoundedBorder()).
	BorderForeground(subtleColor)

var fieldWrapper = lipgloss.NewStyle().
	MarginBottom(1)

var fieldStyle = lipgloss.NewStyle()

var fieldLabalStyle = lipgloss.NewStyle()

var dialogTitleStyle = lipgloss.NewStyle().
	Bold(true).
	MarginTop(2)

var dialogBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#874BFD")).
	Padding(1, 0).
	BorderTop(true).
	BorderLeft(true).
	BorderRight(true).
	BorderBottom(true)

var buttonStyle = lipgloss.NewStyle().
	Foreground(secondaryAccentColorFg).
	Background(secondaryAccentColorBg).
	Padding(0, 3).
	MarginTop(1)

var accentBtnStyle = buttonStyle.
	Foreground(accentColorFg).
	Background(accentColorBg).
	MarginLeft(2)

var cancelBtnStyles = buttonStyle
var confirmBtnStyles = accentBtnStyle

func (m createModel) View() string {
	dialogContent := lipgloss.JoinVertical(lipgloss.Top,
		fieldWrapper.Render(
			lipgloss.JoinVertical(lipgloss.Top,
				fieldLabalStyle.Render("Task Title"),
				m.titleInput.View(),
			),
		),

		fieldWrapper.Render(
			lipgloss.JoinVertical(lipgloss.Top,
				fieldLabalStyle.Render("Task Description"),
				m.descInput.View(),
			),
		),

		lipgloss.JoinHorizontal(lipgloss.Top,
			cancelBtnStyles.Render("Cancel"),
			confirmBtnStyles.Render("Create"),
		),
	)

	dialogWrapper := lipgloss.JoinVertical(lipgloss.Top,
		dialogTitleStyle.Render("Create Task"),
		boxStyles.Render(dialogContent),
	)

	dialog := lipgloss.Place(100, 15,
		lipgloss.Center, lipgloss.Center,
		dialogWrapper,
	)

	return dialog
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
		case "enter":

			switch m.focusedInput {
			case CANCEL_BTN:
				return m, changeViewToListCmd

			case CONFIRM_BTN:
				// TODO display error if title is empty
				if len(m.titleInput.Value()) > 3 {
					task := NewTask(uuid.New().String(), m.titleInput.Value(), m.descInput.Value(), false)
					m.repo.AddTask(task)
					return m, changeViewToListCmd
				}
			}

		case "tab":
			switch m.focusedInput {
			case TITLE_INPUT:
				m.titleInput.Blur()
				m.descInput.Focus()
				m.focusedInput = DESC_INPUT

			case DESC_INPUT:
				m.descInput.Blur()
				m.titleInput.Blur()
				m.focusedInput = CONFIRM_BTN
				confirmBtnStyles = confirmBtnStyles.Underline(true)

			case CONFIRM_BTN:
				confirmBtnStyles = confirmBtnStyles.Underline(false)
				cancelBtnStyles = cancelBtnStyles.Underline(true)
				m.focusedInput = CANCEL_BTN

			case CANCEL_BTN:
				cancelBtnStyles = cancelBtnStyles.Underline(false)
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
