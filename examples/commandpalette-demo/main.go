package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/go-go-golems/bobatea/pkg/commandpalette"
)

// Custom commands for the demo
func demoCommands() []commandpalette.Command {
	return []commandpalette.Command{
		{Name: "/hello", Description: "Say hello", Usage: "/hello", NeedsParameters: false},
		{Name: "/echo", Description: "Echo a message", Usage: "/echo <message>", NeedsParameters: true},
		{Name: "/time", Description: "Show current time", Usage: "/time", NeedsParameters: false},
		{Name: "/clear", Description: "Clear the screen", Usage: "/clear", NeedsParameters: false},
		{Name: "/theme", Description: "Change theme", Usage: "/theme <light|dark>", NeedsParameters: true},
		{Name: "/help", Description: "Show help", Usage: "/help", NeedsParameters: false},
		{Name: "/quit", Description: "Exit the application", Usage: "/quit", NeedsParameters: false},
		{Name: "/random", Description: "Generate random number", Usage: "/random", NeedsParameters: false},
		{Name: "/status", Description: "Show app status", Usage: "/status", NeedsParameters: false},
		{Name: "/version", Description: "Show version info", Usage: "/version", NeedsParameters: false},
	}
}

type model struct {
	palette  commandpalette.Model
	messages []string
	width    int
	height   int
	quitting bool
}

func initialModel() model {
	config := commandpalette.CommandPaletteConfig{
		Commands:    demoCommands(),
		Width:       70,
		Height:      20,
		Title:       "🎯 Command Palette Demo",
		Placeholder: "Type to search commands...",
	}

	return model{
		palette:  commandpalette.New(config),
		messages: []string{"Welcome to the Command Palette Demo!", "Press Ctrl+K to open the command palette"},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.palette.IsVisible() {
			m.palette, cmd = m.palette.Update(msg)
			return m, cmd
		}

		switch msg.String() {
		case "ctrl+c", "q":
			if !m.palette.IsVisible() {
				m.quitting = true
				return m, tea.Quit
			}
		case "ctrl+k":
			m.palette.Show()
			return m, nil
		}

	case commandpalette.CommandSelectedMsg:
		return m.handleCommand(msg.Command, []string{})

	case commandpalette.CommandWithParamsMsg:
		// For demo purposes, we'll simulate getting parameters
		params := []string{"demo parameter"}
		return m.handleCommand(msg.Command, params)

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.palette, cmd = m.palette.Update(msg)
		return m, cmd
	}

	if m.palette.IsVisible() {
		m.palette, cmd = m.palette.Update(msg)
	}

	return m, cmd
}

func (m model) handleCommand(command string, params []string) (tea.Model, tea.Cmd) {
	log.Info().Str("command", command).Strs("params", params).Msg("Executing command")

	switch command {
	case "/hello":
		m.messages = append(m.messages, "👋 Hello there!")
	case "/echo":
		if len(params) > 0 {
			m.messages = append(m.messages, fmt.Sprintf("📢 Echo: %s", strings.Join(params, " ")))
		} else {
			m.messages = append(m.messages, "📢 Echo: (no message provided)")
		}
	case "/time":
		m.messages = append(m.messages, fmt.Sprintf("🕒 Current time: %s", time.Now().Format("15:04:05")))
	case "/clear":
		m.messages = []string{"Screen cleared! ✨"}
	case "/theme":
		if len(params) > 0 {
			m.messages = append(m.messages, fmt.Sprintf("🎨 Theme changed to: %s", params[0]))
		} else {
			m.messages = append(m.messages, "🎨 Theme: (no theme specified)")
		}
	case "/help":
		m.messages = append(m.messages, "💡 Available commands:")
		for _, cmd := range m.palette.GetCommands() {
			m.messages = append(m.messages, fmt.Sprintf("  %s - %s", cmd.Name, cmd.Description))
		}
	case "/quit":
		m.quitting = true
		return m, tea.Quit
	case "/random":
		m.messages = append(m.messages, fmt.Sprintf("🎲 Random number: %d", time.Now().UnixNano()%100))
	case "/status":
		m.messages = append(m.messages, "✅ Application status: Running")
		m.messages = append(m.messages, fmt.Sprintf("📊 Commands available: %d", len(m.palette.GetCommands())))
	case "/version":
		m.messages = append(m.messages, "🏷️  Command Palette Demo v1.0.0")
		m.messages = append(m.messages, "🔧 Built with Bubble Tea & Bobatea")
	default:
		m.messages = append(m.messages, fmt.Sprintf("❓ Unknown command: %s", command))
	}

	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return "Goodbye! 👋\n"
	}

	// Show command palette overlay if visible
	if m.palette.IsVisible() {
		return m.palette.View()
	}

	// Main content styles
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7C3AED")).
		Background(lipgloss.Color("#1F2937")).
		Padding(0, 2).
		Width(m.width)

	messageStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#F9FAFB")).
		Background(lipgloss.Color("#374151")).
		Padding(1, 2).
		Width(m.width - 4).
		Height(m.height - 8)

	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6B7280")).
		Italic(true).
		Padding(0, 2)

	// Title
	title := titleStyle.Render("🎯 Command Palette Demo")

	// Messages area
	var messageLines []string
	maxMessages := (m.height - 8) - 2 // Account for padding and borders
	
	startIdx := 0
	if len(m.messages) > maxMessages {
		startIdx = len(m.messages) - maxMessages
	}
	
	for i := startIdx; i < len(m.messages); i++ {
		messageLines = append(messageLines, m.messages[i])
	}
	
	messagesContent := strings.Join(messageLines, "\n")
	messages := messageStyle.Render(messagesContent)

	// Help text
	help := helpStyle.Render("Press Ctrl+K to open command palette • Ctrl+C or 'q' to quit")

	// Combine everything
	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		messages,
		"",
		help,
	)

	return content
}

func main() {
	// Setup logging
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().Msg("Starting Command Palette Demo")

	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal().Err(err).Msg("Failed to run program")
	}
}
