package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	levels []int
	idx    int
}

var p *tea.Program

func (m model) Init() tea.Cmd {
	return nil
}

type Refresh struct{}

func triggerRefresh() {
	if p != nil {
		go func() {
			p.Send(Refresh{})
		}()
		time.Sleep(20 * time.Millisecond)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "m":
			go func() {
				MergeSort(m.levels)
			}()
		case "u":
			go func() {
				QuickSort(m.levels)
			}()
		case "b":
			go func() {
				BubbleSort(m.levels)
			}()
		case "s":
			go func() {
				SelectionSort(m.levels)
			}()
		case "r":
			m.reset()
		}
	}
	return m, nil
}

func (m *model) reset() {
	m.levels = []int{42, 7, 68, 35, 16, 54, 27, 3, 49, 21, 62, 10, 39, 58, 5, 31, 67, 12, 45, 26, 70, 18, 34, 55, 1, 48, 29, 14, 53, 23, 66, 9, 37, 61, 6, 47, 20, 33, 56, 2, 50, 25, 11, 44, 30, 69, 19, 36, 59, 8, 41, 17, 64, 4, 51, 28, 13, 46, 24, 65, 15, 40, 57, 22, 60, 32, 63, 38, 52, 43}
	go func() {
		p.Send(Refresh{})
	}()
}

func (m model) View() string {
	s := "                 "

	for i := 0; i < len(m.levels); i++ {

		line := ""
		for j := 70; j >= 1; j-- {
			if m.levels[i] >= j {
				line += "█\n"
			} else {
				line += " \n"
			}
		}
		s = lipgloss.JoinHorizontal(lipgloss.Left, s, line)
	}

	return s
}

func main() {
	A := []int{42, 7, 68, 35, 16, 54, 27, 3, 49, 21, 62, 10, 39, 58, 5, 31, 67, 12, 45, 26, 70, 18, 34, 55, 1, 48, 29, 14, 53, 23, 66, 9, 37, 61, 6, 47, 20, 33, 56, 2, 50, 25, 11, 44, 30, 69, 19, 36, 59, 8, 41, 17, 64, 4, 51, 28, 13, 46, 24, 65, 15, 40, 57, 22, 60, 32, 63, 38, 52, 43}
	p = tea.NewProgram(model{levels: A, idx: 0}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app: %v", err)
		os.Exit(1)
	}
}
