package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// GameState represents different states of the game
type GameState int

const (
	StateNormal GameState = iota
	StateDrawOption
)

type Console struct {
	messages     []string
	input        string
	prompt       string
	history      []string
	historyIndex int
	gameState    GameState
	options      []string
	selected     int
}

func NewConsole() *Console {
	return &Console{
		messages:     make([]string, 0),
		prompt:       "Enter command (draw/quit): ",
		history:      make([]string, 0),
		historyIndex: -1,
		gameState:    StateNormal,
		options:      []string{"Draw", "Show"},
		selected:     0,
	}
}

func (c *Console) SetGameState(state GameState) {
	c.gameState = state
}

func (c *Console) GetGameState() GameState {
	return c.gameState
}

func (c *Console) Display(message string) {
	c.messages = append(c.messages, message)
	if len(c.messages) > 20 {
		c.messages = c.messages[len(c.messages)-20:]
	}
}

// Style definitions
var (
	borderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63"))

	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205")).
			Bold(true).
			Padding(0, 1)

	promptStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("205"))
)

func (c *Console) View(title string) string {
	header := borderStyle.Render(titleStyle.Render(title))

	messageArea := strings.Join(c.messages, "\n")
	if messageArea == "" {
		messageArea = "\n\n\n"
	}
	content := borderStyle.Render(messageArea)

	var input string
	if c.gameState == StateDrawOption {
		// Show options when in draw option state
		options := make([]string, len(c.options))
		for i, opt := range c.options {
			if i == c.selected {
				options[i] = "> " + opt
			} else {
				options[i] = "  " + opt
			}
		}
		input = strings.Join(options, "\n")
	} else {
		input = promptStyle.Render(c.prompt) + c.input
	}

	return fmt.Sprintf("%s\n\n%s\n\n%s", header, content, input)
}

func (c *Console) UpdateInput(s string) {
	c.input = s
}

func (c *Console) ClearInput() {
	c.input = ""
}

func (c *Console) AddToHistory(command string) {
	if command != "" {
		c.history = append(c.history, command)
		c.historyIndex = len(c.history)
	}
}

// Navigate history
func (c *Console) NavigateHistory(direction int) {
	if len(c.history) == 0 {
		return
	}

	c.historyIndex += direction

	// Bounds checking
	if c.historyIndex >= len(c.history) {
		c.historyIndex = len(c.history)
		c.input = ""
		return
	}
	if c.historyIndex < 0 {
		c.historyIndex = 0
	}

	if c.historyIndex < len(c.history) {
		c.input = c.history[c.historyIndex]
	}
}

func (c *Console) NavigateOptions(direction int) {
	c.selected += direction
	if c.selected < 0 {
		c.selected = len(c.options) - 1
	}
	if c.selected >= len(c.options) {
		c.selected = 0
	}
}

// GetInput returns the current input
func (c *Console) GetInput() string {
	return c.input
}

func (c *Console) GetSelected() int {
	return c.selected
}

func (c *Console) ClearMessages() {
	c.messages = make([]string, 0)
}
