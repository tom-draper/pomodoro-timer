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

func doTick() tea.Cmd {
	return tea.Tick(time.Second*0, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m *model) initRestPeriod() {
	m.rest = true
	m.timeRemaining = 5 * 60
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case TickMsg:
		if m.cancelNextTick {
			m.cancelNextTick = false
			return m, nil
		} else {
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
				return m, doTick()
			}
			m.cancelNextTick = true
			return m, nil
		}
	}
	return m, nil
}
