package main

import (
	"errors"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor      int
	lists       map[string]List
	currentList int
	listsKeys   []string
}

func initialModel() model {
	return model{
		lists: map[string]List{
			"inprogress": {
				Title:       "In Progress",
				Placeholder: "[No item on the list. Select any item and press `s` to start work on the item.]",
				Items: []Item{
					{Title: "Test", Desc: "Lorem", Done: false},
				},
			},
			"todo": {
				Title:       "Todo",
				Placeholder: "[No item on the list. Press `a` to add new item.]",
				Items: []Item{
					{Title: "Test", Desc: "Lorem", Done: false},
				},
			},
			"done": {
				Title:       "Done",
				Placeholder: "[No item on the list. Select any item and press `c` to complete it.]",
				Items: []Item{
					{Title: "Test", Desc: "Lorem", Done: true},
				},
			},
		},
		currentList: 0,
		listsKeys:   []string{"inprogress", "todo", "done"},
	}
}

type Item struct {
	Title string
	Desc  string
	Done  bool
}

type List struct {
	Title       string
	Placeholder string
	Items       []Item
}

func (m model) Init() tea.Cmd {
	return nil
}


var ErrEmptyList = errors.New("Empty list")

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			listKey := m.listsKeys[m.currentList]

			if m.cursor < len(m.lists[listKey].Items)-1 {
				m.cursor++
			}
		case "tab":
			if m.currentList == len(m.listsKeys)-1 {
				m.currentList = 0
			} else {
				m.currentList += 1
			}
		case "s":
			listKey := m.listsKeys[m.currentList]
			list := m.lists[listKey]
			item := list.Items[m.cursor]
			nextListKey := m.listsKeys[0]

			// Move item to Done list
			if list, ok := m.lists[listKey]; ok {
				list.Items = append(list.Items[:m.cursor], list.Items[m.cursor+1:]...)
				m.lists[listKey] = list
			}

			if list, ok := m.lists[nextListKey]; ok {
				list.Items = append(list.Items, item)
				m.lists[nextListKey] = list
			}
		case "a":
			item := Item{Title: "New item", Done: false, Desc: ""}

			if list, ok := m.lists[m.listsKeys[1]]; ok {
				list.Items = append(list.Items, item)
				m.lists[m.listsKeys[1]] = list
			}
		case "c":
			listKey := m.listsKeys[m.currentList]
			list := m.lists[listKey]
			item := list.Items[m.cursor]
			nextListKey := m.listsKeys[2]

			// Move item to Done list
			if list, ok := m.lists[listKey]; ok {
				list.Items = append(list.Items[:m.cursor], list.Items[m.cursor+1:]...)
				m.lists[listKey] = list
			}

			if list, ok := m.lists[nextListKey]; ok {
				item.Done = !item.Done
				list.Items = append(list.Items, item)
				m.lists[nextListKey] = list
			}

		case "d":
			listKey := m.listsKeys[m.currentList]

			if list, ok := m.lists[listKey]; ok {
				list.Items = append(list.Items[:m.cursor], list.Items[m.cursor+1:]...)
				m.lists[listKey] = list
			}
		case "enter":
			listKey := m.listsKeys[m.currentList]
			nextListKey := m.listsKeys[2]

			if m.cursor < len(m.lists[listKey].Items) {
				item := m.lists[listKey].Items[m.cursor]

				// Move item to Todo list
				if item.Done {
					nextListKey = m.listsKeys[1]
				}

				// Move item to Done list
				if list, ok := m.lists[listKey]; ok {
					list.Items = append(list.Items[:m.cursor], list.Items[m.cursor+1:]...)
					m.lists[listKey] = list
				}

				if list, ok := m.lists[nextListKey]; ok {
					item.Done = !item.Done
					list.Items = append(list.Items, item)
					m.lists[nextListKey] = list
				}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What should we buy at the market?\n\n"

	for listIndex, key := range m.listsKeys {
		list := m.lists[key]
		s += fmt.Sprintf("%s:\n", list.Title)

		if len(list.Items) == 0 {
			s += list.Placeholder
		}

		for i, item := range list.Items {
			cursor := " "
			checked := " "

			if item.Done {
				checked = "x"
			}

			if listIndex == m.currentList {
				if m.cursor == i {
					cursor = ">"
				}
			}
			s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, item.Title)
		}

		s += "\n"
	}

	s += "\nPress q to quit.\n"
	return s
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
