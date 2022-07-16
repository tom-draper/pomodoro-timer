package cmd

import "github.com/muesli/termenv"

type model struct {
	width  int
	height int
	styles Styles
}

type Style func(string) termenv.Style

type Styles struct {
	runningTimer Style
	stoppedTimer Style
}
