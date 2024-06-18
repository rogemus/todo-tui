package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type focusedList int

const (
	TODO_LIST focusedList = iota
	DONE_LIST
)

type tuiModel struct {
	repo        TasksRepository
	todo        list.Model
	done        list.Model
	showDone    bool
	focusedList focusedList
}

func NewTuiModel(repo TasksRepository) tuiModel {
	doneTasks, _ := repo.GetTasksByStatus(true)
	todoTasks, _ := repo.GetTasksByStatus(false)
	doneItems := convertToListitem(doneTasks)
	todoItems := convertToListitem(todoTasks)

	doneList := list.New(doneItems, itemDelegate{}, 0, 0)
	todoList := list.New(todoItems, itemDelegate{}, 0, 0)

	doneList.SetShowHelp(false)
	doneList.SetShowStatusBar(false)
	doneList.Title = "Done"

	todoList.SetShowHelp(false)
	todoList.SetShowStatusBar(false)
	todoList.Title = "Todo"

	return tuiModel{
		repo:        repo,
		todo:        todoList,
		done:        doneList,
		showDone:    true,
		focusedList: TODO_LIST,
	}
}

func convertToListitem(tasks []task) []list.Item {
	var items []list.Item

	for _, t := range tasks {
		items = append(items, list.Item(t))
	}

	return items
}

func (m tuiModel) Init() tea.Cmd {
	return nil
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

func (m tuiModel) View() string {
	if !m.showDone {
		return docStyle.Render(m.todo.View())
	}

	return docStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			m.todo.View(),
			m.done.View(),
		))
}

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "ctrl+d":
			m.showDone = !m.showDone

		case "tab":
			if m.focusedList == TODO_LIST {
				m.focusedList = DONE_LIST
			} else {
				m.focusedList = TODO_LIST
			}

		case "D":
			switch m.focusedList {
			case DONE_LIST:
				i := m.done.Index()
				t, ok := m.done.SelectedItem().(task)
				if ok {
					m.done.RemoveItem(i)
					m.repo.RemoveTask(t.Id())
				}

			case TODO_LIST:
				i := m.todo.Index()
				t, ok := m.todo.SelectedItem().(task)
				if ok {
					m.todo.RemoveItem(i)
					m.repo.RemoveTask(t.Id())
				}
			}

		case "enter":
			var t task
			var i int
			var ok bool

			switch m.focusedList {
			case DONE_LIST:
				i = m.done.Index()
				t, ok = m.done.SelectedItem().(task)

				if ok {
					t.done = false
					m.done.RemoveItem(i)
					m.todo.InsertItem(0, t)
				}

			case TODO_LIST:
				i = m.todo.Index()
				t, ok = m.todo.SelectedItem().(task)

				if ok {
					t.done = true
					m.todo.RemoveItem(i)
					m.done.InsertItem(0, t)
				}
			}
		}

	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()

		m.todo.SetSize(msg.Width-h, msg.Height/2-v)
		m.done.SetSize(msg.Width-h, msg.Height/2-v)
	}

	switch m.focusedList {
	case TODO_LIST:
		m.todo, cmd = m.todo.Update(msg)

	case DONE_LIST:
		m.done, cmd = m.done.Update(msg)
	}

	return m, cmd
}