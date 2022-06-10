package bar

import (
	"github.com/charmbracelet/lipgloss"
)

const height = 1

// ColorConfig
type ColorConfig struct {
	Foreground lipgloss.AdaptiveColor
	Background lipgloss.AdaptiveColor
}

// Bubble represents the properties of the statusbar.
type Bubble struct {
	Width        int
	Height       int
	FirstColumn  string
	SecondColumn string
	ThirdColumn  string
	FourthColumn string
}
