package internal

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	tasks     []Task
	taskTitle textinput.Model
	taskDesc  textarea.Model
}

func NewModel() model {
	return model{}
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
		}
	}

	// Init fields

	// 1. Sidebar
	// 1.1 Init TaskTitle
	// 1.2 Init TaskDesc

	// 2. Lists
	// 2.1 Init InProgress
	// 2.2 Init ToDo
	// 2.3 Init Done

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
	inprogressList := NewList(InProgressTasks, 40, 10)
	todoList := NewList(TodoTasks, 40, 10)
	doneList := NewList(DoneTasks, 40, 10)

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
					lipgloss.JoinVertical(lipgloss.Top,
						title.Render(InProgressListTitle),
						inprogressList.View(),
					),
					lipgloss.JoinVertical(lipgloss.Top,
						title.Render(TodoListTitle),
						todoList.View(),
					),
					lipgloss.JoinVertical(lipgloss.Top,
						title.Render(DoneListTitle),
						doneList.View(),
					),
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
