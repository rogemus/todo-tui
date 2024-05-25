package internal

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	tasks []Task
}

func NewModel() model {
	return model{}
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, cmd
}

var container = lipgloss.NewStyle().Padding(1)

var lists = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("45"))

var sidebar = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("3"))

var title = lipgloss.NewStyle().
	Border(lipgloss.NormalBorder()).
	BorderTop(false).
	BorderLeft(false).
	BorderRight(false).
	BorderForeground(lipgloss.Color("8")).
	PaddingLeft(1).
	Bold(true).
	Width(40)

func (n *model) View() string {
	TaskList.SetShowTitle(false)
	TaskList.SetShowHelp(false)
	TaskList.SetShowStatusBar(false)

	Textarea.SetWidth(40)
  Textarea.ShowLineNumbers = false
  Textarea.SetHeight(30)
	Textarea.SetValue(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec a mauris rhoncus nunc vehicula faucibus non auctor neque. 
      * Lorem ipsum
      * Lorem ipsum 1 

    Quisque eget lacus a ex sodales accumsan. Quisque at sagittis ipsum. Morbi consequat non est quis aliquam. Morbi ac nisl sed lacus varius aliquet sit amet vitae felis. Aenean vitae nunc ut ligula fringilla rutrum. Praesent rhoncus, ligula eget iaculis accumsan, turpis odio viverra orci, a faucibus nisl risus non nunc. Integer rutrum lorem nec ex gravida bibendum.
  `)

	return container.Render(
		lipgloss.JoinHorizontal(lipgloss.Top,
			lists.Render(
				lipgloss.JoinVertical(lipgloss.Top,
					lipgloss.JoinVertical(lipgloss.Top,
						title.Render("IN PROGRESS"),
						TaskList.View(),
					),
					lipgloss.JoinVertical(lipgloss.Top,
						title.Render("TODO"),
						TaskList.View(),
					),
					lipgloss.JoinVertical(lipgloss.Top,
						title.Render("DONE"),
						TaskList.View(),
					),
				),
			),
			sidebar.Render(
				lipgloss.JoinVertical(lipgloss.Top,
					Textarea.View(),
				),
			),
		),
	)
}
