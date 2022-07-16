package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
	termbox "github.com/nsf/termbox-go"
)

func terminalDimensions() (int, int) {
	defer termbox.Close()
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	w, h := termbox.Size()
	return w, h
}

func InitialModel() model {
	profile := termenv.ColorProfile()
	w, h := terminalDimensions()
	return model{
		width:         w,
		height:        h,
		timeRemaining: 25 * 60,
		styles: Styles{
			runningTimer: func(str string) termenv.Style {
				return termenv.String(str).Foreground(profile.Color("10")).Faint()
			},
			stoppedTimer: func(str string) termenv.Style {
				return termenv.String(str).Foreground(profile.Color("6")).Faint()
			},
		},
	}
}

func (m model) Init() tea.Cmd {
	return tickEvery()
}
