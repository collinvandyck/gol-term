package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := model{}
	p := tea.NewProgram(model)
	p.Start()
}

var _ tea.Model = model{}

type model struct {
}

// Init implements tea.Model
func (m model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View implements tea.Model
func (m model) View() string {
	return "view"
}
