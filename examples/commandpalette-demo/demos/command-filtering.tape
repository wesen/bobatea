Output demos/command-filtering.gif

Set FontSize 14
Set Width 1200
Set Height 800
Set Theme "Catppuccin Mocha"

Type "cd bobatea/examples/commandpalette-demo && go run ."
Enter
Sleep 2s

# Open command palette to show all commands
Ctrl+K
Sleep 1s

# Show filtering by typing
Type "th"
Sleep 1s

# Clear and try another filter
Backspace
Backspace
Type "ran"
Sleep 1s

# Clear and try partial match
Backspace
Backspace
Backspace
Type "sho"
Sleep 1s

# Select one
Down
Sleep 500ms
Enter
Sleep 2s

# Try filtering by description
Ctrl+K
Sleep 500ms
Type "current"
Sleep 1s
Enter
Sleep 2s

# Exit
Ctrl+C
Sleep 1s
