package internal

import "github.com/charmbracelet/bubbles/list"
// import "github.com/charmbracelet/lipgloss/list"

type Task struct {
	title       string
	description string
}

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

const (
	TodoListTitle       string = "Todo"
	InProgressListTitle string = "In Progress"
	DoneListTitle       string = "Done"
)

var TodoTasks = []list.Item{
	Task{title: "Raspberry Pi’s", description: "I have ’em all over my house"},
	Task{title: "Nutella", description: "It's good on toast"},
	Task{title: "Bitter melon", description: "It cools you down"},
	Task{title: "Nice socks", description: "And by that I mean socks without holes"},
	Task{title: "Eight hours of sleep", description: "I had this once"},
	Task{title: "Cats", description: "Usually"},
}

var InProgressTasks = []list.Item{
	Task{title: "20° Weather", description: "Celsius, not Fahrenheit"},
	Task{title: "Warm light", description: "Like around 2700 Kelvin"},
}

var DoneTasks = []list.Item{
	Task{title: "Plantasia, the album", description: "My plants love it too"},
	Task{title: "Pour over coffee", description: "It takes forever to make though"},
	Task{title: "VR", description: "Virtual reality...what is there to say?"},
	Task{title: "Noguchi Lamps", description: "Such pleasing organic forms"},
	Task{title: "Linux", description: "Pretty much the best OS"},
	Task{title: "Business school", description: "Just kidding"},
	Task{title: "Pottery", description: "Wet clay is a great feeling"},
	Task{title: "Shampoo", description: "Nothing like clean hair"},
	Task{title: "Table tennis", description: "It’s surprisingly exhausting"},
	Task{title: "Milk crates", description: "Great for packing in your extra stuff"},
	Task{title: "Afternoon tea", description: "Especially the tea sandwich part"},
	Task{title: "Stickers", description: "The thicker the vinyl the better"},
	Task{title: "The vernal equinox", description: "The autumnal equinox is pretty good too"},
	Task{title: "Gaffer’s tape", description: "Basically sticky fabric"},
	Task{title: "Terrycloth", description: "In other words, towel fabric"},
}

var listDeleage = list.NewDefaultDelegate()

func NewList(items []list.Item, width, height int) list.Model {
  listDeleage.ShowDescription = false
	list := list.New(items, listDeleage, width, height)

  list.SetFilteringEnabled(false)
	list.SetShowTitle(false)
	list.SetShowHelp(false)
	list.SetShowStatusBar(false)
	return list
}
