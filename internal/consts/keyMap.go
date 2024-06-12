package consts

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up              key.Binding
	Down            key.Binding
	Quit            key.Binding
	Help            key.Binding
	ChangeFocus     key.Binding
	AddTask         key.Binding
	MarkAsDone      key.Binding
	DeleteTask      key.Binding
	StartTask       key.Binding
	MoveToTodo      key.Binding
	ToggleDetails   key.Binding
	ChangeViewFocus key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.ChangeFocus, k.AddTask, k.ToggleDetails, k.ChangeViewFocus, k.Help}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up},
		{k.Down},
		{k.AddTask},
		{k.DeleteTask},
		{k.MarkAsDone},
		{k.MoveToTodo},
		{k.Help},
		{k.ChangeFocus},
		{k.ChangeViewFocus},
		{k.Quit},
	}
}

var Keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	ChangeFocus: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "change focus"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	AddTask: key.NewBinding(
		key.WithKeys("a"),
		key.WithHelp("a", "add task"),
	),
	MarkAsDone: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "mark as done"),
	),
	DeleteTask: key.NewBinding(
		key.WithKeys("D"),
		key.WithHelp("D", "delete task"),
	),
	MoveToTodo: key.NewBinding(
		key.WithKeys("E"),
		key.WithHelp("E", "stop work/move to todo"),
	),
	StartTask: key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "start work"),
	),
	ToggleDetails: key.NewBinding(
		key.WithKeys("B"),
		key.WithHelp("B", "show/hide details"),
	),
}
