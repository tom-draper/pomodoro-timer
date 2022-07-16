package cmd

import (
	"time"

	"github.com/muesli/termenv"
)

type model struct {
	width         int
	height        int
	styles        Styles
	paused        bool
	rest          bool
	prevTick      time.Time
	owedTime      time.Duration
	timeRemaining int
}

type Style func(string) termenv.Style

type Styles struct {
	runningTimer Style
	stoppedTimer Style
}
