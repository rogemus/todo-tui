package main

import (
	"fmt"
	"os"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	cursor      int
	currentList string
	lists       map[string]List
}

func initialModel() model {
	return model{
		lists: map[string]List{
			"todo": {
				Title: "Todo",
				Items: []Item{
					{Title: "Test", Desc: "Lorem", Done: false},
				},
			},
			"done": {
				Title: "Done",
				Items: []Item{
					{Title: "Test", Desc: "Lorem", Done: true},
				},
			},
		},
		currentList: "todo",
	}
}

type Item struct {
	Title  string
	Desc   string
	Done   bool
	ListId string
}

type List struct {
	Title string
	Items []Item
}

func (m model) Init() tea.Cmd {
	return nil
}

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
			if m.cursor < len(m.lists[m.currentList].Items)-1 {
				m.cursor++
			}
		case "tab":
			if m.currentList == "todo" {
				m.currentList = "done"
			} else {
				m.currentList = "todo"
			}

			m.cursor = 0
		case "enter":
			item := m.lists[m.currentList].Items[m.cursor]
			nextList := ""

			if m.currentList == "todo" {
				nextList = "done"
			} else {
				nextList = "todo"
			}

			if list, ok := m.lists[m.currentList]; ok {
				list.Items = append(list.Items[:m.cursor], list.Items[m.cursor+1:]...)
				m.lists[m.currentList] = list
			}

			if list, ok := m.lists[nextList]; ok {
				item.Done = !item.Done
				list.Items = append(list.Items, item)
				m.lists[nextList] = list
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What should we buy at the market?\n\n"

	keys := make([]string, 0)
	for k := range m.lists {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		list := m.lists[key]
		s += fmt.Sprintf("%s:\n", list.Title)

		for i, item := range list.Items {
			cursor := " "
			checked := " "

			if item.Done {
				checked = "x"
			}

			if key == m.currentList {
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
