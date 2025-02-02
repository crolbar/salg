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
	autoIdx    int
}

var nums = []int{42, 7, 68, 35, 16, 54, 27, 3, 49, 21, 62, 10, 39, 58, 5, 31, 67, 12, 45, 26, 70, 18, 34, 55, 1, 48, 29, 14, 53, 23, 66, 9, 37, 61, 6, 47, 20, 33, 56, 2, 50, 25, 11, 44, 30, 69, 19, 36, 59, 8, 41, 17, 64, 4, 51, 28, 13, 46, 24, 65, 15, 40, 57, 22, 60, 32, 63, 38, 52, 43}

var p *tea.Program

func (m model) Init() tea.Cmd {
	return nil
}

type Refresh struct{}
type Next struct{}

func triggerNext() {
	go func() {
		p.Send(Next{})
	}()
}

func triggerRefresh() {
	go func() {
		p.Send(Refresh{})
	}()
	time.Sleep(50 * time.Millisecond)
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
		case "a":
			triggerNext()
		}
	case Next:
		m.reset()

		switch m.autoIdx {
		case 4:
			m.autoIdx = 0
			fallthrough
		case 0:
			go func() {
				MergeSort(m.levels)
				triggerNext()
			}()
		case 1:
			go func() {
				QuickSort(m.levels)
				triggerNext()
			}()
		case 2:
			go func() {
				SelectionSort(m.levels)
				triggerNext()
			}()
		case 3:
			go func() {
				BubbleSort(m.levels)
				triggerNext()
			}()

		}

		m.autoIdx++
	}
	return m, nil
}

func (m *model) reset() {
	copy(m.levels, nums)
}

func (m model) View() string {
	out := ""
	sec := int(time.Now().Unix())

	var (
		i_start int
		i_end   func(j int) bool
		i_inc   int
	)

	// if (sec & (sec % 1)) > 0 {
	i_start = 0
	i_end = func(j int) bool { return j < len(m.levels) }
	i_inc = 1
	// } else {
	// 	i_start = len(m.levels) - 1
	// 	i_end = func(j int) bool { return j >= 0 }
	// 	i_inc = -1
	// }

	for i := i_start; i_end(i); i += i_inc {

		line := ""

		var (
			j_start int
			j_end   func(j int) bool
			j_inc   int
		)

		// if (sec & 1) > 0 {
		// j_start = 1
		// j_end = func(j int) bool { return j <= 70 }
		// j_inc = 1
		// } else {
		j_start = 70
		j_end = func(j int) bool { return j >= 1 }
		j_inc = -1
		// }

		for j := j_start; j_end(j); j += j_inc {
			if m.levels[i] >= j {
				line += "█\n"
			} else {
				line += " \n"
			}
		}

		s := lipgloss.NewStyle().
			Foreground(lipgloss.Color(fmt.Sprintf("%d", (i+sec)%255)))
		out = lipgloss.JoinHorizontal(lipgloss.Left, out, s.Render(line))
	}

	return out
}

func main() {
	p = tea.NewProgram(model{levels: nums, autoIdx: 0}, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting app: %v", err)
		os.Exit(1)
	}
}
