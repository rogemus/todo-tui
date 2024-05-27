package internal

type List struct {
	title       string
	placeholder string
	items       []Item
}

func NewList2(title string, items []Item) List {
	return List{title: title, items: items}
}

func (l *List) Items() []Item {
	return l.items
}

func (l *List) AddItem(item Item) {
	l.items = append(l.items, item)
}

func (l *List) RemoveItem(index int) {
	l.items = append(l.items[:index], l.items[index+1:]...)
}

func (l *List) Enumerator() string {
	return ">>>"
}

func (l *List) View() string {
	return ""
}
