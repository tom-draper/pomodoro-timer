package cmd

import (
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
	return tea.Tick(time.Duration(0), func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m *model) initRestPeriod() {
	m.rest = true
	m.timer.timeRemaining = m.timer.restPeriod
}

func (m *model) initWorkPeriod() {
	m.rest = false
	m.timer.timeRemaining = m.timer.workPeriod
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case TickMsg:
		if m.cancelNextTick {
			m.cancelNextTick = false
			return m, nil
		} else {
			m.timer.timeRemaining--
			if m.timer.timeRemaining < 1 {
				if m.rest {
					m.initWorkPeriod()
				} else {
					m.initRestPeriod()
				}
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
			m.timer.paused = !m.timer.paused
			if m.timer.paused {
				m.cancelNextTick = true // Cancel effects from the next tickEvery()
				return m, nil
			} else {
				return m, doTick()
			}
		}
	}
	return m, nil
}
