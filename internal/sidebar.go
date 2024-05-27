package internal

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
)

func NewInput(value, placeholder string, width int) textinput.Model {
	input := textinput.New()
	input.Placeholder = placeholder
	input.SetValue(value)
	input.Width = width

	return input
}

func NewTextarea(value, placeholder string, width, heigth int) textarea.Model {
	textarea := textarea.New()
	textarea.SetValue(value)
	textarea.Placeholder = placeholder
	textarea.SetWidth(width)
	textarea.SetHeight(heigth)

	return textarea
}
