package cmd

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type TickMsg time.Time

func tickEvery() tea.Cmd {
	// Send a message every second
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func doTick(waitTime time.Duration, paused *bool) tea.Cmd {
	fmt.Println("Started")
	return tea.Tick(waitTime, func(t time.Time) tea.Msg {
		fmt.Println("Doing tick")
		if !*paused {
			fmt.Println("Returning tick")
			return TickMsg(t)
		}
		fmt.Println("Retuning nil")
		return nil
	})
}

func (m *model) initRestPeriod() {
	m.rest = true
	m.timeRemaining = 5 * 60
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case TickMsg:
		if m.paused {
			fmt.Println("Paused")
			return m, nil
		} else {
			fmt.Println("Tick")
			m.prevTick = time.Now()
			m.owedTime = 0
			m.timeRemaining--
			if m.timeRemaining < 1 {
				m.initRestPeriod()
				return m, nil
			}
			return m, tickEvery()
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case " ":
			m.paused = !m.paused
			if !m.paused {
				// Wait the owed time from the second which was paused inside of
				return m, doTick(m.owedTime, &m.paused)
			} else {
				m.owedTime = time.Since(m.prevTick)
			}
			return m, nil
		}
	}
	return m, nil
}
