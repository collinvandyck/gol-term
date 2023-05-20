package main

import (
	"bytes"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	model := model{}
	p := tea.NewProgram(model)
	p.Start()
}

type model struct {
	board [][]bool
	ready bool
}

type TickMsg time.Time

// Init implements tea.Model
func (m model) Init() tea.Cmd {
	return doTick()
}

func doTick() tea.Cmd {
	return tea.Tick(time.Second/60, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Update implements tea.Model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.board = make([][]bool, msg.Height)
		for i := range m.board {
			m.board[i] = make([]bool, msg.Width)
		}
		for i := range m.board {
			for j := range m.board[i] {
				if rand.Float64() < 0.5 {
					m.board[i][j] = true
				}
			}
		}
		m.ready = true
	case TickMsg:
		return m, doTick()
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
	buf := new(bytes.Buffer)
	for i := range m.board {
		for j := range m.board[i] {
			if m.board[i][j] {
				buf.WriteString("x")
			} else {
				buf.WriteString(" ")
			}
		}
		if i != len(m.board)-1 {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}
