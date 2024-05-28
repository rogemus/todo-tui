package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"todo-tui/internal/models"
)

type ListsViewModel struct {
	cursor      int
	currentList int
	listKeys    [3]string
	lists       map[string]*models.List
}

func NewListsModel() ListsViewModel {
	//TODO: store and load items from file/db
	var inprogressTasks = []models.Item{
		{Title: "Warm light", Description: "Like around 2700 Kelvin"},
		{Title: "Warm light", Description: ""},
	}

	var todoTasks = []models.Item{
		{Title: "20° Weather", Description: "Celsius, not Fahrenheit"},
		{Title: "Warm light", Description: ""},
	}
	var doneTasks = []models.Item{
		{Title: "20° Weather", Description: "Celsius, not Fahrenheit"},
		{Title: "Table tennis", Description: "It’s surprisingly exhausting"},
		{Title: "Milk crates", Description: "Great for packing in your extra stuff"},
		{Title: "Afternoon tea", Description: "Especially the tea sandwich part"},
		{Title: "Warm light", Description: ""},
	}

	lists := make(map[string]*models.List)
	lists["inprogress"] = models.NewList("In Progress", inprogressTasks)
	lists["todo"] = models.NewList("To Do", todoTasks)
	lists["done"] = models.NewList("Done", doneTasks)

	lists["inprogress"].SetSelected(0)
	return ListsViewModel{
		cursor:      0,
		currentList: 0,
		lists:       lists,
		listKeys:    [3]string{"inprogress", "todo", "done"},
	}
}

func (m ListsViewModel) Init() tea.Cmd {
	return nil
}

func (m ListsViewModel) View() string {
	return lipgloss.JoinVertical(lipgloss.Top,
		m.lists["inprogress"].View(),
		m.lists["todo"].View(),
		m.lists["done"].View(),
	)
}

func (m ListsViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "up", "k":
			listKey := m.listKeys[m.currentList]
			list := m.lists[listKey]
			sel := list.Selected()

			if sel > 0 {
				m.cursor -= 1
				list.SetSelected(m.cursor)
			}
		case "tab":
			m.lists[m.listKeys[m.currentList]].SetSelected(-9999)

			if m.currentList == len(m.listKeys)-1 {
				m.currentList = 0
			} else {
				m.currentList += 1
			}

			m.cursor = 0
			m.lists[m.listKeys[m.currentList]].SetSelected(0)
		case "down", "j":
			listKey := m.listKeys[m.currentList]
			sel := m.lists[listKey].Selected()

			if sel < len(m.lists[listKey].Items())-1 {
				m.cursor += 1
				m.lists[listKey].SetSelected(m.cursor)
			}
		}
	}

	return m, cmd
}
