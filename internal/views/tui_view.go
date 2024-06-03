package views

import (
	"database/sql"
	"log"
	"os"
	"todo-tui/internal/consts"
	"todo-tui/internal/models"
	"todo-tui/internal/storage"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	_ "github.com/mattn/go-sqlite3"
)

type state int

const (
	listsView state = iota
	detailsView
)

type SelectedItemMsg struct {
	Item models.Item
}

type TuiModel struct {
	state          state
	lists          ListsViewModel
	details        DetailsViewModel
	help           help.Model
	keys           consts.KeyMap
	detailsVisible bool
}

const dbName = "tasks.db/?parseTime=true"

func NewTuiModel() TuiModel {
	db, _ := sql.Open("sqlite3", dbName)
	tasksRepo := storage.NewTasksRepository(db)
	err := tasksRepo.CreateNewDb()

	if err != nil {
		log.Printf("Error while creating DB")
		os.Exit(1)
	}

	return TuiModel{
		state:          listsView,
		lists:          NewListsModel(tasksRepo),
		details:        NewDetailsModel(tasksRepo),
		keys:           consts.Keys,
		help:           help.New(),
		detailsVisible: false,
	}
}

func (m TuiModel) Init() tea.Cmd {
	return nil
}

// Styles
var containerStyles = lipgloss.NewStyle().PaddingTop(1)

var dividerStyles = lipgloss.NewStyle().
	Margin(0, 1).
	Faint(true).
	BorderStyle(lipgloss.NormalBorder()).
	BorderTop(false).
	BorderBottom(false).
	BorderLeft(false).
	BorderRight(true).
	BorderForeground(consts.ColSubtle)

func (m TuiModel) View() string {
	divider := ""
	details := ""

	if m.detailsVisible {
		divider = dividerStyles.Render()
		details = m.details.View()
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		containerStyles.Render(
			lipgloss.JoinHorizontal(lipgloss.Top,
				m.lists.View(),
				divider,
				details,
			)),
		m.help.View(m.keys),
	)
}

func (m TuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case SelectedItemMsg:
		m.details, cmd = m.details.Update(msg)
	case tea.WindowSizeMsg:
		msg.Height = msg.Height - 2
		m.help.Width = msg.Width

		dividerStyles = dividerStyles.Height(msg.Height)
    m.details.SetSize(msg.Width, msg.Height)
		m.lists.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.ToggleDetails):
			m.detailsVisible = !m.detailsVisible
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
