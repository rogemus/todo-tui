package views

import (
	"fmt"
	"todo-tui/internal/consts"
	"todo-tui/internal/models"
	"todo-tui/internal/storage"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DetailsViewModel struct {
	sel    models.Item
	repo   storage.TasksRepository
	Width  int
	Height int
}

func NewDetailsModel(repo storage.TasksRepository) DetailsViewModel {
	return DetailsViewModel{
		repo: repo,
	}
}

func (m DetailsViewModel) Init() tea.Cmd {
	return nil
}

var detailsViewStyles = lipgloss.NewStyle()

var titleStyle = lipgloss.NewStyle().
	Bold(true).
	PaddingLeft(1)

var titleContainer = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderTop(false).
	BorderLeft(false).
	BorderRight(false).
	BorderForeground(consts.ColSubtle).
	PaddingLeft(1)

var descStyles = lipgloss.NewStyle().
	PaddingTop(1)

var checkStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("106"))

var placeholderStyles = lipgloss.NewStyle().
	Faint(true)

func (m DetailsViewModel) RenderTitle() string {
	status := "[ ]"

	if m.sel.Status == models.DONE {
		status = fmt.Sprintf("[%s]", checkStyle.Render("✓"))
	}

	titleContainer = titleContainer.Width(m.Width)
	titleStyle = titleStyle.Width(m.Width)

	return titleContainer.Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			status,
			titleStyle.Render(m.sel.Title),
		),
	)
}

func (m DetailsViewModel) RenderDesc() string {
	desc := m.sel.Description

	if len(desc) == 0 {
		return descStyles.Render(
			placeholderStyles.Render("No description for this task ..."),
		)
	}

	return descStyles.Render(m.sel.Description)
}

func (m *DetailsViewModel) SetSize(width, height int) {
	m.Width = width
	m.Height = height
	detailsViewStyles.Height(m.Height).Width(m.Width)
}

func (m DetailsViewModel) View() string {
	return detailsViewStyles.Render(
		lipgloss.JoinVertical(lipgloss.Top,
			m.RenderTitle(),
			m.RenderDesc(),
		),
	)
}

func (m DetailsViewModel) Update(msg tea.Msg) (DetailsViewModel, tea.Cmd) {
	switch msg := msg.(type) {
	case SelectedItemMsg:
		m.sel = msg.Item
	}
	return m, nil
}
