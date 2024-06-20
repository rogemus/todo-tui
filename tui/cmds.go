package tui

import tea "github.com/charmbracelet/bubbletea"

type changeToListViewMsg string
type changeToCreateViewMsg string

func changeViewToListCmd() tea.Msg {
	return changeToListViewMsg("list")
}

func changeViewToCreateCmd() tea.Msg {
	return changeToCreateViewMsg("create")
}
