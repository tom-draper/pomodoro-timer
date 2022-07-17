package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func formatTimer(timeRemaining int) string {
	var timer strings.Builder
	minutes := fmt.Sprint(timeRemaining / 60)
	if len(minutes) == 1 {
		timer.WriteString("0")
	}
	timer.WriteString(minutes)
	timer.WriteString(":")
	seconds := fmt.Sprint(timeRemaining % 60)
	if len(seconds) == 1 {
		timer.WriteString("0")
	}
	timer.WriteString(seconds)
	return timer.String()
}

func (m model) progressBar() string {
	var progress strings.Builder
	nBars := int(float32(m.width) * 0.7)
	totalPeriod := m.timer.workPeriod
	timerStyle := m.styles.workTimer
	if m.rest {
		totalPeriod = m.timer.restPeriod
		timerStyle = m.styles.restTimer
	}
	timeSpent := totalPeriod - m.timer.timeRemaining
	bars := int((float64(timeSpent) / float64(totalPeriod)) * float64(nBars))

	progress.WriteString(style(strings.Repeat("|", bars), timerStyle))
	progress.WriteString(style(strings.Repeat("|", (nBars-bars)), m.styles.faint))
	return progress.String()
}

func (m model) View() string {
	var sb strings.Builder

	timerStyle := m.styles.workTimer
	if m.timer.paused {
		sb.WriteString("Paused")
	} else if m.rest {
		sb.WriteString("Rest Period")
		timerStyle = m.styles.restTimer
	}
	sb.WriteString("\n\n")

	timer := formatTimer(m.timer.timeRemaining)
	sb.WriteString(style(timer, timerStyle))
	sb.WriteString(" ")

	progress := m.progressBar()
	sb.WriteString(progress)
	sb.WriteString("\n")

	s := lipgloss.NewStyle().Align(lipgloss.Center).Render(sb.String())

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, s)
}

func style(s string, style Style) string {
	return style(s).String()
}
