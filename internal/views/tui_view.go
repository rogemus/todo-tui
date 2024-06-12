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

type focusedView int

const (
	LIST_VIEW focusedView = iota
	DETAILS_VIEW
	CREATE_VIEW
)

type SelectedItemMsg struct {
	Item models.Item
}

type TuiModel struct {
	listsViewModel   ListsViewModel
	detailsViewModel DetailsViewModel
	createViewModel  CreateViewModel
	help             help.Model
	keys             consts.KeyMap
	detailsVisible   bool
	focusedView      focusedView
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
		focusedView:      LIST_VIEW,
		listsViewModel:   NewListsModel(tasksRepo),
		detailsViewModel: NewDetailsModel(tasksRepo),
		createViewModel:  NewCreateView(tasksRepo),
		keys:             consts.Keys,
		help:             help.New(),
		detailsVisible:   false,
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

	if m.focusedView == CREATE_VIEW {
		return m.createViewModel.View()
	}

	if m.detailsVisible {
		divider = dividerStyles.Render()
		details = m.detailsViewModel.View()
	}

	return lipgloss.JoinVertical(lipgloss.Top,
		containerStyles.Render(
			lipgloss.JoinHorizontal(lipgloss.Top,
				m.listsViewModel.View(),
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
		m.detailsViewModel, cmd = m.detailsViewModel.Update(msg)
	case tea.WindowSizeMsg:
		msg.Height = msg.Height - 3
		m.help.Width = msg.Width

		dividerStyles = dividerStyles.Height(msg.Height)
		m.detailsViewModel.SetSize(msg.Width, msg.Height)
		m.listsViewModel.SetSize(msg.Width, msg.Height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.ToggleDetails):
			m.detailsVisible = !m.detailsVisible
		case key.Matches(msg, m.keys.AddTask):
			m.focusedView = CREATE_VIEW
		default:
			switch m.focusedView {
			case CREATE_VIEW:
				m.createViewModel, cmd = m.createViewModel.Update(msg)
			case LIST_VIEW:
				m.listsViewModel, cmd = m.listsViewModel.Update(msg)
			case DETAILS_VIEW:
				m.detailsViewModel, cmd = m.detailsViewModel.Update(msg)
			}
		}
	}

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
