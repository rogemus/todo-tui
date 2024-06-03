package views

import (
	"todo-tui/internal/consts"
	"todo-tui/internal/models"
	"todo-tui/internal/storage"

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
	repo        storage.TasksRepository
	Height      int
	Width       int
}

func NewListsModel(repo storage.TasksRepository) ListsViewModel {
	inprogressTasks, _ := repo.GetTasks(models.IN_PROGRESS)
	todoTasks, _ := repo.GetTasks(models.TODO)
	doneTasks, _ := repo.GetTasks(models.DONE)

	lists := make(map[string]*models.List)
	lists["inprogress"] = models.NewList("In Progress", "No item in In Progress list ...", inprogressTasks)
	lists["todo"] = models.NewList("To Do", "No item in To Do list ...", todoTasks)
	lists["done"] = models.NewList("Done", "No item in Done list ...", doneTasks)

	lists["inprogress"].SetSelected(0)

	return ListsViewModel{
		cursor:      0,
		currentList: 0,
		lists:       lists,
		listKeys:    [3]string{"inprogress", "todo", "done"},
		keys:        consts.Keys,
		repo:        repo,
	}
}

func (m ListsViewModel) Init() tea.Cmd {
	return nil
}

func (m ListsViewModel) SetSize(width, height int) {
	m.Width = width
	m.Height = height

	singleListHeight := (height / 3) - 3
	m.lists["inprogress"].SetSize(m.Width, singleListHeight)
	m.lists["todo"].SetSize(m.Width, singleListHeight)
	m.lists["done"].SetSize(m.Width, singleListHeight)
	listViewStyle = listViewStyle.Height(m.Height)
}

var listViewStyle = lipgloss.NewStyle()

func (m ListsViewModel) View() string {
	return listViewStyle.Render(
		lipgloss.JoinVertical(lipgloss.Top,
			m.lists["inprogress"].View(),
			m.lists["todo"].View(),
			m.lists["done"].View(),
		),
	)
}

func handleItemChange(item models.Item) tea.Cmd {
	return func() tea.Msg {
		return SelectedItemMsg{
			Item: item,
		}
	}
}

func (m ListsViewModel) Update(msg tea.Msg) (ListsViewModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Up):
			listKey := m.listKeys[m.currentList]

			if len(m.lists[listKey].Items()) > 0 {
				if m.cursor > 0 {
					m.cursor -= 1
					m.lists[listKey].SetSelected(m.cursor)
				}
				cmd = handleItemChange(m.lists[listKey].Items()[m.cursor])
			}
		case key.Matches(msg, m.keys.ChangeFocus):
			m.lists[m.listKeys[m.currentList]].SetSelected(-9999)

			if m.currentList == len(m.listKeys)-1 {
				m.currentList = 0
			} else {
				m.currentList += 1
			}

			m.cursor = 0
			if len(m.lists[m.listKeys[m.currentList]].Items()) > 0 {
				m.lists[m.listKeys[m.currentList]].SetSelected(0)
				cmd = handleItemChange(m.lists[m.listKeys[m.currentList]].Items()[m.cursor])
			}
		case key.Matches(msg, m.keys.Down):
			listKey := m.listKeys[m.currentList]
			sel := m.lists[listKey].Selected()

			if len(m.lists[listKey].Items()) > 0 {
				if sel < len(m.lists[listKey].Items())-1 {
					m.cursor += 1
					m.lists[listKey].SetSelected(m.cursor)
				}
				cmd = handleItemChange(m.lists[listKey].Items()[m.cursor])
			}

		case key.Matches(msg, m.keys.DeleteTask):
			listKey := m.listKeys[m.currentList]
			item := m.lists[listKey].Items()[m.cursor]
			m.lists[listKey].RemoveItem(m.cursor)
			m.repo.RemoveTask(item.Id)
		case key.Matches(msg, m.keys.MarkAsDone):
			listKey := m.listKeys[m.currentList]
			items := m.lists[listKey].Items()

			if listKey != "done" && len(items) > 0 {
				item := items[m.cursor]
				m.lists[listKey].RemoveItem(m.cursor)
				item.Status = models.DONE
				m.lists["done"].AddItem(item)
				m.repo.UpdateTask(item)
			}
		case key.Matches(msg, m.keys.AddTask):
			item := models.Item{Title: "New Item", Description: "", Status: models.TODO}
			m.lists["todo"].AddItem(item)
			m.repo.AddTask(item)
		case key.Matches(msg, m.keys.StartTask):
			listKey := m.listKeys[m.currentList]
			items := m.lists[listKey].Items()

			if listKey != "inprogress" && len(items) > 0 {
				item := items[m.cursor]
				m.lists[listKey].RemoveItem(m.cursor)
				item.Status = models.IN_PROGRESS
				m.lists["inprogress"].AddItem(item)
				m.repo.UpdateTask(item)
			}
		case key.Matches(msg, m.keys.MoveToTodo):
			listKey := m.listKeys[m.currentList]
			items := m.lists[listKey].Items()

			if listKey != "todo" && len(items) > 0 {
				item := items[m.cursor]
				m.lists[listKey].RemoveItem(m.cursor)
				item.Status = models.TODO
				m.lists["todo"].AddItem(item)
				m.repo.UpdateTask(item)
			}
		}
	}

	return m, cmd
}
