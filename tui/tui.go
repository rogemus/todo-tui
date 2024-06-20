package tui

import tea "github.com/charmbracelet/bubbletea"

type selectedView int

const (
	LIST_VIEW selectedView = iota
	CREATE_VIEW
)

type tuiModel struct {
	listsModel   tea.Model
	createModel  tea.Model
	selectedView selectedView
}

func NewTuiModel(repo TasksRepository) tuiModel {
	listsView := NewListsModel(repo)
	createView := NewCreateModal(repo)

	return tuiModel{
		listsModel:   listsView,
		createModel:  createView,
		selectedView: LIST_VIEW,
	}
}

func (m tuiModel) Init() tea.Cmd {
	return nil
}

func (m tuiModel) View() string {
	if m.selectedView == CREATE_VIEW {
		return m.createModel.View()
	}

	return m.listsModel.View()
}

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case changeToListViewMsg:
		m.selectedView = LIST_VIEW

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "ctrl+a", "A":
			m.selectedView = CREATE_VIEW
			return m, changeViewToCreateCmd
		}
	}

	switch m.selectedView {
	case LIST_VIEW:
		m.listsModel, cmd = m.listsModel.Update(msg)

	case CREATE_VIEW:
		m.createModel, cmd = m.createModel.Update(msg)
	}

	return m, cmd
}
