package models

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

/* Styles */

var titleStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderTop(false).
	BorderLeft(false).
	BorderRight(false).
	BorderForeground(lipgloss.Color("8")).
	PaddingLeft(1).
	Bold(true).
	Width(40)

type Status int

const (
	IN_PROGRESS Status = iota
	TODO
	DONE
)

type Item struct {
	Id          int
	Title       string
	Description string
	CreatedAt   time.Time
	Status      Status
}

// TODO: Pagination. View Limit
type List struct {
	title       string
	placeholder string
	items       []Item
	selected    int
}

func NewList(title string, items []Item) *List {
	return &List{title: title, items: items, selected: -9999}
}

func (l *List) SetSelected(index int) {
	l.selected = index
}

func (l *List) Selected() int {
	return l.selected
}

func (l *List) Items() []Item {
	return l.items
}

func (l *List) AddItem(item Item) {
	l.items = append(l.items, item)
}

func (l *List) UpdateItem(index int, item Item) {
	l.items[index] = item
}

func (l *List) Title() string {
	return titleStyle.Render(l.title)
}

func (l *List) RemoveItem(index int) {
	l.items = append(l.items[:index], l.items[index+1:]...)
}

func (l *List) Enumerator() string {
	return ">>> "
}

func (l *List) View() string {
	view := ""
	view += fmt.Sprintf("%s\n", l.Title())

	for index, item := range l.items {
		prefix := "    "
		sufix := ""

		if l.selected == index {
			prefix = l.Enumerator()
		}

		if len(item.Description) > 0 {
			sufix = " 📝"
		}

		view += fmt.Sprintf("%s%s%s\n", prefix, item.Title, sufix)
	}

	return view
}
