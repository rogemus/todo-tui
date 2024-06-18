package tui

type task struct {
	id, title, description string
	done                   bool
}

func NewTask(id, title, description string, done bool) task {
	return task{id, title, description, done}
}

func (t task) Done() bool { return t.done }

func (t task) Title() string { return t.title }

func (t task) Description() string { return t.description }

func (t task) Id() string { return t.id }

func (t task) FilterValue() string { return t.title }
