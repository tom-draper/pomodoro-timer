package cmd

import (
	"github.com/muesli/termenv"
)

type timer struct {
	paused        bool
	workPeriod    int
	restPeriod    int
	timeRemaining int
}

type model struct {
	width          *int
	height         *int
	styles         Styles
	rest           bool
	cancelNextTick bool
	timer          timer
}

type Style func(string) termenv.Style

type Styles struct {
	restTimer Style
	workTimer Style
	faint     Style
}
