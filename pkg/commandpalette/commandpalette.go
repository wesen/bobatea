// Package commandpalette provides a reusable command palette component for Bubble Tea applications.
//
// The command palette allows users to search and select commands with fuzzy filtering,
// keyboard navigation, and parameter support.
//
// Example usage:
//
//	config := DefaultCommandPaletteConfig()
//	palette := commandpalette.New(config)
//
//	// In your model's Update method:
//	if palette.IsVisible() {
//		palette, cmd = palette.Update(msg)
//		return model, cmd
//	}
//
//	// To show the palette:
//	palette.Show()
//
//	// In your View method:
//	if palette.IsVisible() {
//		return palette.View()
//	}
package commandpalette

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Command represents a command that can be executed
type Command struct {
	Name            string
	Description     string
	Usage           string
	NeedsParameters bool
}

// CommandPaletteConfig holds configuration for the command palette
type CommandPaletteConfig struct {
	Commands    []Command
	Width       int
	Height      int
	Title       string
	Placeholder string
}

// Event types for communication between components
type (
	// CommandSelectedMsg is sent when a command is selected (no parameters needed)
	CommandSelectedMsg struct {
		Command string
	}

	// CommandWithParamsMsg is sent when a command needing parameters is selected
	CommandWithParamsMsg struct {
		Command string
	}
)

// DefaultCommands returns a default set of commands
func DefaultCommands() []Command {
	return []Command{
		{Name: "/help", Description: "Show available commands", Usage: "/help", NeedsParameters: false},
		{Name: "/clear", Description: "Clear history", Usage: "/clear", NeedsParameters: false},
		{Name: "/users", Description: "List users", Usage: "/users", NeedsParameters: false},
		{Name: "/nick", Description: "Change nickname", Usage: "/nick <new_name>", NeedsParameters: true},
		{Name: "/theme", Description: "Change color theme", Usage: "/theme [light|dark]", NeedsParameters: true},
		{Name: "/status", Description: "Show status", Usage: "/status", NeedsParameters: false},
		{Name: "/time", Description: "Show current time", Usage: "/time", NeedsParameters: false},
		{Name: "/quit", Description: "Exit", Usage: "/quit", NeedsParameters: false},
	}
}

// DefaultCommandPaletteConfig returns a default command palette configuration
func DefaultCommandPaletteConfig() CommandPaletteConfig {
	return CommandPaletteConfig{
		Commands:    DefaultCommands(),
		Width:       60,
		Height:      15,
		Title:       "Command Palette",
		Placeholder: "Search commands...",
	}
}

// Styles for the command palette
var (
	primary   = lipgloss.Color("#7C3AED") // Purple
	secondary = lipgloss.Color("#06B6D4") // Cyan
	accent    = lipgloss.Color("#10B981") // Emerald
	muted     = lipgloss.Color("#6B7280") // Gray
	text      = lipgloss.Color("#F9FAFB") // Almost white
	surface   = lipgloss.Color("#1F2937") // Lighter dark
)

// Model represents the command palette component
type Model struct {
	config   CommandPaletteConfig
	input    textinput.Model
	filtered []Command
	selected int
	visible  bool
	width    int
	height   int
}

// New creates a new command palette with the given configuration
func New(config CommandPaletteConfig) Model {
	input := textinput.New()
	input.Placeholder = config.Placeholder
	input.Width = config.Width - 10 // Account for padding and borders

	return Model{
		config:   config,
		input:    input,
		filtered: config.Commands,
		visible:  false,
	}
}

// Show displays the command palette
func (m *Model) Show() {
	m.visible = true
	m.input.Focus()
	m.input.SetValue("")
	m.filtered = m.config.Commands
	m.selected = 0
}

// Hide hides the command palette
func (m *Model) Hide() {
	m.visible = false
	m.input.Blur()
}

// IsVisible returns whether the palette is visible
func (m Model) IsVisible() bool {
	return m.visible
}

// SetCommands updates the available commands
func (m *Model) SetCommands(commands []Command) {
	m.config.Commands = commands
	m.filterCommands()
}

// GetSelectedCommand returns the currently selected command, if any
func (m Model) GetSelectedCommand() *Command {
	if len(m.filtered) > 0 && m.selected < len(m.filtered) {
		return &m.filtered[m.selected]
	}
	return nil
}

// filterCommands filters commands based on input
func (m *Model) filterCommands() {
	query := strings.ToLower(m.input.Value())
	m.filtered = []Command{}

	for _, cmd := range m.config.Commands {
		if strings.Contains(strings.ToLower(cmd.Name), query) ||
			strings.Contains(strings.ToLower(cmd.Description), query) {
			m.filtered = append(m.filtered, cmd)
		}
	}

	// Reset selection if out of bounds
	if m.selected >= len(m.filtered) {
		m.selected = 0
	}
}

// Update handles command palette updates
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.visible {
		return m, nil
	}

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.Hide()
			return m, nil
		case "enter":
			if len(m.filtered) > 0 && m.selected < len(m.filtered) {
				selectedCmd := m.filtered[m.selected]
				m.Hide()
				if selectedCmd.NeedsParameters {
					return m, func() tea.Msg {
						return CommandWithParamsMsg{Command: selectedCmd.Name}
					}
				} else {
					return m, func() tea.Msg {
						return CommandSelectedMsg{Command: selectedCmd.Name}
					}
				}
			}
		case "up":
			if m.selected > 0 {
				m.selected--
			}
		case "down":
			if m.selected < len(m.filtered)-1 {
				m.selected++
			}
		default:
			m.input, cmd = m.input.Update(msg)
			m.filterCommands()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		// Update input width based on container size
		m.input.Width = min(m.config.Width-10, msg.Width-10)
	}

	return m, cmd
}

// View renders the command palette
func (m Model) View() string {
	if !m.visible {
		return ""
	}

	// Create the container style
	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(primary).
		Background(surface).
		Padding(1, 2).
		Width(min(m.config.Width, m.width-4)).
		Height(min(m.config.Height, m.height-4))

	// Title
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(primary).
		Render(m.config.Title)

	// Input
	input := m.input.View()

	// Commands list
	var commandsList []string
	maxVisible := min(8, len(m.filtered)) // Show max 8 commands

	for i := 0; i < maxVisible; i++ {
		if i >= len(m.filtered) {
			break
		}

		cmd := m.filtered[i]
		style := lipgloss.NewStyle().Foreground(text)
		if i == m.selected {
			style = style.Background(primary).Bold(true)
		}

		nameStyle := lipgloss.NewStyle().Foreground(accent).Bold(true)
		descStyle := lipgloss.NewStyle().Foreground(muted).Italic(true)

		line := nameStyle.Render(cmd.Name) + descStyle.Render(" - "+cmd.Description)
		commandsList = append(commandsList, style.Render(line))
	}

	// Show scroll indicator if there are more commands
	if len(m.filtered) > maxVisible {
		scrollInfo := lipgloss.NewStyle().
			Foreground(muted).
			Italic(true).
			Render("... and more (use arrow keys)")
		commandsList = append(commandsList, scrollInfo)
	}

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		input,
		"",
		strings.Join(commandsList, "\n"),
	)

	return containerStyle.Render(content)
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GetCommands returns the current list of commands
func (m Model) GetCommands() []Command {
	return m.config.Commands
}

// GetConfig returns the current configuration
func (m Model) GetConfig() CommandPaletteConfig {
	return m.config
}

// UpdateConfig updates the configuration
func (m *Model) UpdateConfig(config CommandPaletteConfig) {
	m.config = config
	m.input.Placeholder = config.Placeholder
	m.input.Width = config.Width - 10
}
