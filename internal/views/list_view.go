package views

import (
	"todo-tui/internal/consts"
	"todo-tui/internal/models"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListsViewModel struct {
	cursor      int
	currentList int
	listKeys    [3]string
	lists       map[string]*models.List
	keys        consts.KeyMap
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
		keys:        consts.Keys,
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
		switch {
		case key.Matches(msg, m.keys.Up):
			listKey := m.listKeys[m.currentList]
			list := m.lists[listKey]
			sel := list.Selected()

			if sel > 0 {
				m.cursor -= 1
				list.SetSelected(m.cursor)
			}
		case key.Matches(msg, m.keys.ChangeFocus):
			m.lists[m.listKeys[m.currentList]].SetSelected(-9999)

			if m.currentList == len(m.listKeys)-1 {
				m.currentList = 0
			} else {
				m.currentList += 1
			}

			m.cursor = 0
			m.lists[m.listKeys[m.currentList]].SetSelected(0)
		case key.Matches(msg, m.keys.Down):
			listKey := m.listKeys[m.currentList]
			sel := m.lists[listKey].Selected()

			if sel < len(m.lists[listKey].Items())-1 {
				m.cursor += 1
				m.lists[listKey].SetSelected(m.cursor)
			}
		case key.Matches(msg, m.keys.DeleteTask):
			listKey := m.listKeys[m.currentList]
			m.lists[listKey].RemoveItem(m.cursor)
		case key.Matches(msg, m.keys.MarkAsDone):
			listKey := m.listKeys[m.currentList]
			items := m.lists[listKey].Items()

			if listKey != "done" && len(items) > 0 {
				item := items[m.cursor]
				m.lists[listKey].RemoveItem(m.cursor)
				m.lists["done"].AddItem(item)
			}
		case key.Matches(msg, m.keys.AddTask):
			item := models.Item{Title: "New Item", Description: ""}
			m.lists["todo"].AddItem(item)
		case key.Matches(msg, m.keys.StartTask):
			listKey := m.listKeys[m.currentList]
			items := m.lists[listKey].Items()

			if listKey != "inprogress" && len(items) > 0 {
				item := items[m.cursor]
				m.lists[listKey].RemoveItem(m.cursor)
				m.lists["inprogress"].AddItem(item)
			}
		case key.Matches(msg, m.keys.MoveToTodo):
			listKey := m.listKeys[m.currentList]
			items := m.lists[listKey].Items()

			if listKey != "todo" && len(items) > 0 {
				item := items[m.cursor]
				m.lists[listKey].RemoveItem(m.cursor)
				m.lists["todo"].AddItem(item)
			}
		}
	}

	return m, cmd
}
