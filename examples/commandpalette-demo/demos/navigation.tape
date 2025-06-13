Output demos/navigation.gif

Set FontSize 14
Set Width 1200
Set Height 800
Set Theme "Catppuccin Mocha"

Type "cd bobatea/examples/commandpalette-demo && go run ."
Enter
Sleep 2s

# Open command palette
Ctrl+K
Sleep 1s

# Show navigation with arrow keys
Down
Sleep 500ms
Down
Sleep 500ms
Down
Sleep 500ms
Down
Sleep 500ms

# Go back up
Up
Sleep 500ms
Up
Sleep 500ms

# Select a command
Enter
Sleep 2s

# Open again and show escape to close
Ctrl+K
Sleep 1s
Escape
Sleep 1s

# Open again
Ctrl+K
Sleep 500ms

# Navigate to quit command
Type "quit"
Sleep 1s
Enter
Sleep 1s
