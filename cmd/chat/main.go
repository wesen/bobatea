package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-go-golems/bobatea/pkg/chat"
	conversation2 "github.com/go-go-golems/bobatea/pkg/conversation"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chat",
	Short: "A chat application with different backend options",
	Long:  `This chat application allows you to choose between a fake backend and an HTTP backend.`,
}

var fakeCmd = &cobra.Command{
	Use:   "fake",
	Short: "Run the chat application with a fake backend",
	Run:   runFakeBackend,
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Run the chat application with an HTTP backend",
	Run:   runHTTPBackend,
}

var httpAddr string

func init() {
	rootCmd.AddCommand(fakeCmd)
	rootCmd.AddCommand(httpCmd)

	httpCmd.Flags().StringVarP(&httpAddr, "addr", "a", ":8080", "HTTP server address")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runFakeBackend(cmd *cobra.Command, args []string) {
	runChat(func() chat.Backend {
		return NewFakeBackend()
	})
}

func runHTTPBackend(cmd *cobra.Command, args []string) {
	runChat(func() chat.Backend {
		return NewHTTPBackend(httpAddr)
	})
}

func runChat(backendFactory func() chat.Backend) {
	manager := conversation2.NewManager(conversation2.WithMessages(
		conversation2.NewChatMessage(conversation2.RoleSystem, "Welcome to the chat application!"),
	))

	backend := backendFactory()

	options := []tea.ProgramOption{
		tea.WithMouseCellMotion(),
		tea.WithAltScreen(),
	}

	p := tea.NewProgram(chat.InitialModel(manager, backend), options...)

	// Set the program for the backend after initialization
	if setterBackend, ok := backend.(interface{ SetProgram(*tea.Program) }); ok {
		setterBackend.SetProgram(p)
	}

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
