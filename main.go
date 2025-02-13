package main

import (
	"fmt"
	"game/internal/game"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func run() error {
	g := game.New()
	p := tea.NewProgram(g)

	_, err := p.Run()
	return err
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
