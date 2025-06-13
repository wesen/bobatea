Output demos/basic-usage.gif

Set FontSize 14
Set Width 1200
Set Height 800
Set Theme "Catppuccin Mocha"

Type "cd bobatea/examples/commandpalette-demo && go run ."
Enter
Sleep 2s

Type "This demo shows the command palette in action!"
Sleep 1s

# Open command palette
Ctrl+K
Sleep 1s

# Show some typing to filter commands
Type "hel"
Sleep 1s

# Navigate with arrows
Down
Sleep 500ms
Up
Sleep 500ms

# Select the hello command
Enter
Sleep 2s

# Open palette again
Ctrl+K
Sleep 500ms

# Try a different command
Type "time"
Sleep 1s
Enter
Sleep 2s

# Open palette one more time
Ctrl+K
Sleep 500ms

# Try echo command (needs parameters)
Type "echo"
Sleep 1s
Enter
Sleep 2s

# Exit
Ctrl+C
Sleep 1s
