package internal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor      int
	currentList int
	listKeys    [3]string
	lists       map[string]*List
}

func NewModel() model {
	// Init fields

	// 1. Sidebar
	// 1.1 Init TaskTitle
	// 1.2 Init TaskDesc


  //TODO: store and load items from file/db 
	var inprogressTasks = []Item{
		{Title: "Warm light", Description: "Like around 2700 Kelvin"},
		{Title: "Warm light", Description: ""},
	}

	var todoTasks = []Item{
		{Title: "20° Weather", Description: "Celsius, not Fahrenheit"},
		{Title: "Warm light", Description: ""},
	}
	var doneTasks = []Item{
		{Title: "20° Weather", Description: "Celsius, not Fahrenheit"},
		{Title: "Table tennis", Description: "It’s surprisingly exhausting"},
		{Title: "Milk crates", Description: "Great for packing in your extra stuff"},
		{Title: "Afternoon tea", Description: "Especially the tea sandwich part"},
		{Title: "Warm light", Description: ""},
	}

	lists := make(map[string]*List)
	lists["inprogress"] = NewList("In Progress", inprogressTasks)
	lists["todo"] = NewList("To Do", todoTasks)
	lists["done"] = NewList("Done", doneTasks)

	lists["inprogress"].SetSelected(0)
	return model{
		cursor:      0,
		currentList: 0,
		lists:       lists,
		listKeys:    [3]string{"inprogress", "todo", "done"},
	}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
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

var container = lipgloss.NewStyle().Padding(1)

var lists = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("45"))

var sidebar = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("3"))

var title = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderTop(false).
	BorderLeft(false).
	BorderRight(false).
	BorderForeground(lipgloss.Color("8")).
	PaddingLeft(1).
	Bold(true).
	Width(40)

func (m *model) View() string {
	// inprogressList := NewList(InProgressTasks, 40, 10)
	// todoList := NewList(TodoTasks, 40, 10)
	// doneList := NewList(DoneTasks, 40, 10)

	taskTitle := NewInput("Title Lorem", "Title ...", 40)
	taskDesc := NewTextarea(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a mauris rhoncus nunc vehicula faucibus non auctor neque. 
      * Lorem ipsum
      * Lorem ipsum 1 

    Quisque eget lacus a ex sodales accumsan. Quisque at sagittis ipsum. Morbi consequat non est quis aliquam. Morbi ac nisl sed lacus varius aliquet sit amet vitae felis. Aenean vitae nunc ut ligula fringilla rutrum. Praesent rhoncus, ligula eget iaculis accumsan, turpis odio viverra orci, a faucibus nisl risus non nunc. Integer rutrum lorem nec ex gravida bibendum.
  `, "Description ...", 40, 30)

	return container.Render(
		lipgloss.JoinHorizontal(lipgloss.Top,
			lists.Render(
				lipgloss.JoinVertical(lipgloss.Top,
					m.lists["inprogress"].View(),
					m.lists["todo"].View(),
					m.lists["done"].View(),
				),
			),
			sidebar.Render(
				lipgloss.JoinVertical(lipgloss.Top,
					title.Render(
						taskTitle.View(),
					),
					taskDesc.View(),
				),
			),
		),
	)
}
