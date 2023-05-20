package main

import (
	"bytes"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := model{}
	p := tea.NewProgram(model)
	p.Start()
}

var _ tea.Model = model{}

type model struct {
	width  int
	height int
	ready  bool
}

// Init implements tea.Model
func (m model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true
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
	if !m.ready {
		return ""
	}
	var idx rune
	buf := new(bytes.Buffer)
loop:
	for i := 0; i < m.height; i++ {
		if i >= 20000000000 {
			break loop
		}
		for j := 0; j < m.width; j++ {
			if idx < 32 {
				idx = 32
			}
			if idx >= 127 {
				idx = 32
			}
			buf.WriteRune(idx)
			idx++
		}
		if i < m.height-1 {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}
