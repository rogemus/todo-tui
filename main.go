package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const (
	todo status = iota
	inprogress
	done
)

/* STYLING */
var (
	windowStyle = lipgloss.NewStyle().
			Padding(1, 1)
	columnStyle = lipgloss.NewStyle().
			Padding(1, 1).
			Border(lipgloss.HiddenBorder()).
			BorderTop(false).
			BorderBottom(false).
			BorderRight(false).
			BorderForeground(lipgloss.Color("67"))
	focusedStyle = lipgloss.NewStyle().
			Padding(1, 1).
			Border(lipgloss.NormalBorder()).
			BorderTop(false).
			BorderBottom(false).
			BorderRight(false).
			BorderForeground(lipgloss.Color("67"))
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

/* CUSTOM ITEM (TASK) */
type Task struct {
	title       string
	description string
	status      status
}

// list.Item interface
func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

/* MAIN MODEL */
type Model struct {
	lists   []list.Model
	focused status
	loaded  bool
}

func NewModel() *Model {
	return &Model{}
}

func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height/4)
	defaultList.SetShowHelp(false)

	m.lists = []list.Model{defaultList, defaultList, defaultList}

	// InProgress List
	m.lists[inprogress].Title = "In Progress"
	m.lists[inprogress].SetItems([]list.Item{
		Task{status: inprogress, title: "Do something", description: "desc"},
	})

	// Todo List
	m.lists[todo].Title = "To Do"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "Do something", description: "desc"},
		Task{status: todo, title: "Do test something", description: "desc"},
	})

	// Done List
	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		Task{status: done, title: "Do something", description: "desc"},
		Task{status: done, title: "Do test something", description: "desc"},
		Task{status: done, title: "Do test something", description: "desc"},
	})
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initLists(msg.Width, msg.Height)
			m.loaded = true
		}
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if m.focused == 2 {
				m.focused = 0
			} else {
				m.focused += 1
			}
		}
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)
	return m, cmd
}

func (m *Model) View() string {
	if !m.loaded {
		return "Loading..."
	}

	todoView := m.lists[todo].View()
	inprogView := m.lists[inprogress].View()
	doneView := m.lists[done].View()

	sidepanel := lipgloss.JoinVertical(
		lipgloss.Top,
		columnStyle.Render(todoView),
	)

	lists := lipgloss.JoinVertical(
		lipgloss.Top,
		columnStyle.Render(inprogView),
		focusedStyle.Render(todoView),
		columnStyle.Render(doneView),
	)

	view := lipgloss.JoinHorizontal(
		lipgloss.Top,
    lipgloss.Width(400).Render(lists),
		sidepanel,
	)

	return view

	// switch m.focused {
	// case todo:
	// 	return lipgloss.JoinVertical(
	// 		lipgloss.Left,
	// 		columnStyle.Render(inprogView),
	// 		focusedStyle.Render(todoView),
	// 		columnStyle.Render(doneView),
	// 	)
	// case inprogress:
	// 	return lipgloss.JoinVertical(
	// 		lipgloss.Left,
	// 		focusedStyle.Render(inprogView),
	// 		columnStyle.Render(todoView),
	// 		columnStyle.Render(doneView),
	// 	)
	// case done:
	// 	return lipgloss.JoinVertical(
	// 		lipgloss.Left,
	// 		columnStyle.Render(inprogView),
	// 		columnStyle.Render(todoView),
	// 		focusedStyle.Render(doneView),
	// 	)
	// }
	//
	// return lipgloss.JoinVertical(
	// 	lipgloss.Left,
	// 	columnStyle.Render(todoView),
	// 	columnStyle.Render(inprogView),
	// 	columnStyle.Render(doneView),
	// )
}

/* MAIN */
func main() {
	m := NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
