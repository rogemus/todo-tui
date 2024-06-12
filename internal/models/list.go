package models

import (
	"fmt"
	"time"
	"todo-tui/internal/consts"

	"github.com/charmbracelet/lipgloss"
)

/* Styles */

var listContainer = lipgloss.NewStyle()

var listTitleStyle = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderTop(false).
	BorderLeft(false).
	BorderRight(false).
	BorderForeground(consts.ColSubtle).
	PaddingLeft(1).
	Bold(true)

var itemContainerStyle = lipgloss.NewStyle()

var enumeratorStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.HiddenBorder()).
	BorderRight(false).
	BorderTop(false).
	BorderBottom(false).
	BorderLeft(true).
	PaddingRight(1)

var enumeratorFocusedStyle = enumeratorStyle.
	BorderStyle(lipgloss.ThickBorder()).
	BorderForeground(lipgloss.Color("111"))

var checkStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("106"))

var itemTitleStyle = lipgloss.NewStyle().Padding(0, 1)

var doneItemTitleStyle = itemTitleStyle.
	Faint(true).
	Strikethrough(true)

var descStyle = lipgloss.NewStyle().
	Faint(true)

var placeholderStyle = lipgloss.NewStyle().
	PaddingLeft(1).
	PaddingBottom(1).
	Faint(true)

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
	Width       int
	Heigth      int
	viewLimit   int
}

func NewList(title string, placeholder string, items []Item) *List {
	return &List{
		title:       title,
		placeholder: placeholder,
		items:       items,
		selected:    -9999,
		viewLimit:   5,
	}
}

func (l *List) SetSelected(index int) {
	l.selected = index
}

func (l List) Selected() int {
	return l.selected
}

func (l List) Items() []Item {
	return l.items
}

func (l *List) AddItem(item Item) {
	l.items = append([]Item{item}, l.items...)
}

func (l *List) UpdateItem(index int, item Item) {
	l.items[index] = item
}

func (l List) Title() string {
	return listTitleStyle.Render(
		l.title,
		descStyle.Render(fmt.Sprintf("(%d)", len(l.items))),
	)
}

func (l *List) RemoveItem(index int) {
	l.items = append(l.items[:index], l.items[index+1:]...)
}

func (l *List) SetSize(width, heigth int) {
	l.Width = width
	l.Heigth = heigth
	listContainer = listContainer.Height(l.Heigth)
	// listTitleStyle = listTitleStyle.Width(l.Width)
	// itemTitleStyle = itemTitleStyle.Width(l.Width)
}

func (l *List) View() string {
	view := ""
	view += fmt.Sprintf("%s\n", l.Title())

	if len(l.items) == 0 {
		view += placeholderStyle.Render(l.placeholder)
	}

	for index, item := range l.items {
		if index == l.viewLimit {
			break
		}

		part_enumarator := enumeratorStyle.Render()
		part_status := "[ ]"
		part_title := itemTitleStyle.Render(item.Title)
		part_desc := ""

		if index == l.selected {
			part_enumarator = enumeratorFocusedStyle.Render()
		}

		if item.Status == DONE {
			part_status = fmt.Sprintf("[%s]", checkStyle.Render("✓"))
			part_title = doneItemTitleStyle.Render(item.Title)
		}

		if len(item.Description) > 0 {
			part_desc = descStyle.Render("...")
		}

		itemStr := fmt.Sprintf("%s%s%s%s", part_enumarator, part_status, part_title, part_desc)
		view += itemContainerStyle.Render(itemStr)
		view += "\n"
	}

	return listContainer.Render(view)
}
