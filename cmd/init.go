package cmd

import tea "github.com/charmbracelet/bubbletea"

func InitialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
