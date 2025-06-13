package commandpalette

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNew(t *testing.T) {
	config := DefaultCommandPaletteConfig()
	model := New(config)

	if model.visible {
		t.Error("Expected model to start hidden")
	}

	if len(model.filtered) != len(config.Commands) {
		t.Errorf("Expected %d filtered commands, got %d", len(config.Commands), len(model.filtered))
	}
}

func TestShowHide(t *testing.T) {
	config := DefaultCommandPaletteConfig()
	model := New(config)

	// Test show
	model.Show()
	if !model.IsVisible() {
		t.Error("Expected model to be visible after Show()")
	}

	// Test hide
	model.Hide()
	if model.IsVisible() {
		t.Error("Expected model to be hidden after Hide()")
	}
}

func TestSetCommands(t *testing.T) {
	config := DefaultCommandPaletteConfig()
	model := New(config)

	newCommands := []Command{
		{Name: "/test", Description: "Test command", Usage: "/test", NeedsParameters: false},
	}

	model.SetCommands(newCommands)

	if len(model.GetCommands()) != 1 {
		t.Errorf("Expected 1 command after SetCommands, got %d", len(model.GetCommands()))
	}

	if model.GetCommands()[0].Name != "/test" {
		t.Errorf("Expected command name '/test', got '%s'", model.GetCommands()[0].Name)
	}
}

func TestFilterCommands(t *testing.T) {
	config := DefaultCommandPaletteConfig()
	model := New(config)
	model.Show()

	// Test filtering by name
	model.input.SetValue("help")
	model.filterCommands()

	found := false
	for _, cmd := range model.filtered {
		if cmd.Name == "/help" {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected to find '/help' command when filtering by 'help'")
	}
}

func TestGetSelectedCommand(t *testing.T) {
	config := DefaultCommandPaletteConfig()
	model := New(config)
	model.Show()

	// Test with valid selection
	if model.selected < len(model.filtered) {
		cmd := model.GetSelectedCommand()
		if cmd == nil {
			t.Error("Expected to get a selected command")
		}
	}

	// Test with no commands
	model.filtered = []Command{}
	cmd := model.GetSelectedCommand()
	if cmd != nil {
		t.Error("Expected nil when no commands are available")
	}
}

func TestUpdate(t *testing.T) {
	config := DefaultCommandPaletteConfig()
	model := New(config)

	// Test that hidden model doesn't process messages
	updatedModel, cmd := model.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if cmd != nil {
		t.Error("Expected no command when model is hidden")
	}

	// Test showing and hiding with escape
	model.Show()
	updatedModel, cmd = model.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if updatedModel.IsVisible() {
		t.Error("Expected model to be hidden after escape key")
	}

	// Test selection with enter
	model.Show()
	if len(model.filtered) > 0 {
		updatedModel, cmd = model.Update(tea.KeyMsg{Type: tea.KeyEnter})
		if cmd == nil {
			t.Error("Expected command to be returned when selecting with enter")
		}
		if updatedModel.IsVisible() {
			t.Error("Expected model to be hidden after selection")
		}
	}
}

func TestNavigation(t *testing.T) {
	config := DefaultCommandPaletteConfig()
	model := New(config)
	model.Show()

	originalSelected := model.selected

	// Test down navigation
	if len(model.filtered) > 1 {
		model, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
		if model.selected != originalSelected+1 {
			t.Errorf("Expected selected to increase by 1, got %d", model.selected)
		}

		// Test up navigation
		model, _ = model.Update(tea.KeyMsg{Type: tea.KeyUp})
		if model.selected != originalSelected {
			t.Errorf("Expected selected to return to original value %d, got %d", originalSelected, model.selected)
		}
	}
}

func TestDefaultCommands(t *testing.T) {
	commands := DefaultCommands()
	if len(commands) == 0 {
		t.Error("Expected default commands to not be empty")
	}

	// Check for some expected commands
	expectedCommands := []string{"/help", "/clear", "/quit"}
	for _, expected := range expectedCommands {
		found := false
		for _, cmd := range commands {
			if cmd.Name == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected to find command '%s' in default commands", expected)
		}
	}
}

func TestDefaultConfig(t *testing.T) {
	config := DefaultCommandPaletteConfig()

	if config.Width <= 0 {
		t.Error("Expected positive width in default config")
	}

	if config.Height <= 0 {
		t.Error("Expected positive height in default config")
	}

	if config.Title == "" {
		t.Error("Expected non-empty title in default config")
	}

	if config.Placeholder == "" {
		t.Error("Expected non-empty placeholder in default config")
	}

	if len(config.Commands) == 0 {
		t.Error("Expected commands in default config")
	}
}

func TestUpdateConfig(t *testing.T) {
	config := DefaultCommandPaletteConfig()
	model := New(config)

	newConfig := CommandPaletteConfig{
		Commands:    []Command{{Name: "/new", Description: "New command"}},
		Width:       100,
		Height:      20,
		Title:       "New Title",
		Placeholder: "New placeholder",
	}

	model.UpdateConfig(newConfig)

	if model.GetConfig().Title != "New Title" {
		t.Errorf("Expected title 'New Title', got '%s'", model.GetConfig().Title)
	}

	if model.input.Placeholder != "New placeholder" {
		t.Errorf("Expected placeholder 'New placeholder', got '%s'", model.input.Placeholder)
	}
}
