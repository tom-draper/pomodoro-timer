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
	foreground := termenv.ForegroundColor()
	w, h := terminalDimensions()
	workPeriod := 25 * 60
	restPeriod := 5 * 60
	return model{
		width:  w,
		height: h,
		timer: timer{
			timeRemaining: workPeriod,
			workPeriod:    workPeriod,
			restPeriod:    restPeriod,
		},
		rest: false,
		styles: Styles{
			workTimer: func(str string) termenv.Style {
				return termenv.String(str).Foreground(profile.Color("11"))
			},
			restTimer: func(str string) termenv.Style {
				return termenv.String(str).Foreground(profile.Color("10"))
			},
			faint: func(str string) termenv.Style {
				return termenv.String(str).Foreground(foreground).Faint()
			},
		},
	}
}

func (m model) Init() tea.Cmd {
	return tickEvery()
}
