package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var sb strings.Builder

	title := style(fmt.Sprint(m.timeRemaining), m.styles.runningTimer)
	sb.WriteString(title)

	s := lipgloss.NewStyle().Align(lipgloss.Center).Render(sb.String())

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, s)
}

func style(s string, style Style) string {
	return style(s).String()
}
