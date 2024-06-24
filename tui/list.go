package tui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ListModel list.Model

type itemDelegate struct {
	cursorVisible bool
}

func (d itemDelegate) SetCursorVisible(visible bool) {
	print("visible", visible)
	d.cursorVisible = visible
}

func (d itemDelegate) Height() int { return 1 }

func (d itemDelegate) Spacing() int { return 0 }

func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

var itemStyles = lipgloss.NewStyle().PaddingLeft(2)
var titleStyles = lipgloss.NewStyle()

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(task)

	if !ok {
		return
	}

	statusStr := " "
	cursorStr := " "
	titleStr := i.title

	if i.done {
		statusStr = "✓"
	}

	if index == m.Index() && d.cursorVisible {
		cursorStr = ">"
	}

	str := fmt.Sprintf("%s [%s] %s", cursorStr, statusStr, titleStr)
	fmt.Fprint(w, str)
}
