package main

import (
	"github.com/charmbracelet/bubbletea"
)
import "github.com/go-go-golems/bobatea/pkg/filepicker"

type Model struct {
	fp filepicker.Model
}

func NewModel() Model {
	fp := filepicker.NewModel()
	fp.Filepicker.DirAllowed = false
	fp.Filepicker.FileAllowed = true
	fp.Filepicker.CurrentDirectory = "/home/manuel"
	fp.Filepicker.Height = 10

	return Model{
		fp: fp,
	}
}

func (m Model) Init() tea.Cmd {
	return m.fp.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case filepicker.SelectFileMsg:
		return m, tea.Quit

	case filepicker.CancelFilePickerMsg:
		return m, tea.Quit

	case tea.KeyMsg:
		switch {
		case msg.Type == tea.KeyCtrlC:
			return m, tea.Quit
		default:
			m.fp, cmd = m.fp.Update(msg)
			return m, cmd
		}

	default:
		m.fp, cmd = m.fp.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	return m.fp.View()
}

func main() {
	b := NewModel()

	p := tea.NewProgram(b)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
