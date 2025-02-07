package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	levels  []int
	autoIdx int
}

var nums = []int{42, 7, 68, 35, 16, 54, 27, 3, 49, 21, 62, 10, 39, 58, 5, 31, 67, 12, 45, 26, 70, 18, 34, 55, 1, 48, 29, 14, 53, 23, 66, 9, 37, 61, 6, 47, 20, 33, 56, 2, 50, 25, 11, 44, 30, 69, 19, 36, 59, 8, 41, 17, 64, 4, 51, 28, 13, 46, 24, 65, 15, 40, 57, 22, 60, 32, 63, 38, 52, 43}

var sorts []func([]int) = []func([]int){
	MergeSort,
	QuickSort,
	SelectionSort,
	BubbleSort,
	InsertionSort,
}

var p *tea.Program

func (m model) Init() tea.Cmd {
	return nil
}

type Refresh struct{}
type Next struct{}

func triggerNext() {
	p.Send(Next{})
}

func triggerRefresh() {
	p.Send(Refresh{})
	time.Sleep(50 * time.Millisecond)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "m":
			go MergeSort(m.levels)
		case "u":
			go QuickSort(m.levels)
		case "b":
			go BubbleSort(m.levels)
		case "s":
			go SelectionSort(m.levels)
		case "i":
			go InsertionSort(m.levels)
		case "r":
			m.reset()
		case "a":
			go triggerNext()
		}
	case Next:
		m.reset()

		m.autoIdx++

		if m.autoIdx == 5 {
			m.autoIdx = 0
		}

		go func() {
			(sorts[m.autoIdx])(m.levels)
			triggerNext()
		}()
	}
	return m, nil
}

func (m *model) reset() {
	m.levels = make([]int, len(nums))
	copy(m.levels, nums)
}

var s lipgloss.Style = lipgloss.NewStyle()

func (m model) View() string {
	out := ""
	sec := int(time.Now().Unix())

	for i := 0; i < len(m.levels); i++ {
		var line strings.Builder

		for j := 70; j >= 1; j-- {
			if m.levels[i] >= j {
				line.WriteString("█\n")
			} else {
				line.WriteString(" \n")
			}
		}

		out = lipgloss.JoinHorizontal(lipgloss.Left, out,
			s.Foreground(
				lipgloss.Color(
					fmt.Sprintf("%d", (i+sec)%255))).
				Render(line.String()),
		)
	}

	return out
}

func main() {
	levels := make([]int, len(nums))
	copy(levels, nums)
	p = tea.NewProgram(model{levels: levels, autoIdx: 0}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app: %v", err)
		os.Exit(1)
	}
}
