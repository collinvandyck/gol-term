package main

import (
	"bytes"
	"math/rand"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const fps = 10
const alive = "😂"
const dead = "🌊"

func main() {
	model := model{}
	p := tea.NewProgram(model,
		tea.WithAltScreen(),
	)
	p.Start()
}

type board [][]bool

type model struct {
	board board
	ready bool
}

type TickMsg time.Time

// Init implements tea.Model
func (m model) Init() tea.Cmd {
	return doTick()
}

func doTick() tea.Cmd {
	return tea.Tick(time.Second/fps, func(t time.Time) tea.Msg {
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
		m.seed()
		m.ready = true
	case TickMsg:
		if m.ready {
			m.board = m.tick()
		}
		return m, doTick()
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		default:
			m.seed()
		}
	}
	return m, nil
}

func (m model) seed() {
	const minRatio = 0.25
	const maxRatio = 0.65
	liveRatio := rand.Float64()
	if liveRatio < minRatio {
		liveRatio = minRatio
	}
	if liveRatio > maxRatio {
		liveRatio = maxRatio
	}
	for i := range m.board {
		for j := range m.board[i] {
			if rand.Float64() < liveRatio {
				m.board[i][j] = true
			} else {
				m.board[i][j] = false
			}
		}
	}
}

func (m model) tick() board {
	res := make([][]bool, len(m.board))
	for i := range res {
		res[i] = make([]bool, len(m.board[i]))
	}
	for i := range m.board {
		for j := range m.board[i] {
			alive := m.board[i][j]
			res[i][j] = m.board[i][j]
			switch {
			case alive:
				n := m.countLiveNeighbors(i, j)
				if n < 2 || n > 3 {
					res[i][j] = false
				}
			default:
				n := m.countLiveNeighbors(i, j)
				if n == 3 {
					res[i][j] = true
				}
			}
		}
	}
	return res
}

func (m model) countLiveNeighbors(i, j int) int {
	var count int
	for x := i - 1; x <= i+1; x++ {
		for y := j - 1; y <= j+1; y++ {
			if x == i && y == j {
				continue
			}
			cx, cy := x, y
			if cx < 0 {
				cx = len(m.board) - 1
			}
			if cx >= len(m.board) {
				cx = 0
			}
			if cy < 0 {
				cy = len(m.board[cx]) - 1
			}
			if cy >= len(m.board[cx]) {
				cy = 0
			}
			if m.board[cx][cy] {
				count++
			}
		}
	}
	return count
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
				buf.WriteString(alive)
			} else {
				buf.WriteString(dead)
			}
		}
		if i != len(m.board)-1 {
			buf.WriteString("\n")
		}
	}
	return buf.String()
}
