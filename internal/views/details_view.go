package views

import tea "github.com/charmbracelet/bubbletea"

type DetailsViewModel struct{}

func NewDetailsModel() DetailsViewModel {
	return DetailsViewModel{}
}

func (m DetailsViewModel) Init() tea.Cmd {
	return nil
}

func (m DetailsViewModel) View() string {
	return "Details"
}

func (m DetailsViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}
