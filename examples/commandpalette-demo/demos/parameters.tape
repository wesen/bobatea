Output demos/parameters.gif

Require go

Set FontSize 14
Set Width 1200
Set Height 800
Set Padding 20
Set Theme "Dracula"

Type "cd bobatea/examples/commandpalette-demo"
Enter
Sleep 1s

Type "go run ."
Enter
Sleep 2s

# Show parameter commands
Key ctrl+k
Sleep 1s

# Filter to echo command
Type "echo"
Sleep 1s
Key Enter
Sleep 1s

# Type parameters
Type "Hello from VHS demo!"
Sleep 1s
Key Enter
Sleep 2s

# Try theme command
Key ctrl+k
Sleep 1s
Type "theme"
Sleep 1s
Key Enter
Sleep 1s
Type "dark"
Sleep 1s
Key Enter
Sleep 2s

# Try nick command
Key ctrl+k
Sleep 1s
Type "nick"
Sleep 1s
Key Enter
Sleep 1s
Type "DemoUser"
Sleep 1s
Key Enter
Sleep 2s

# Show status
Key ctrl+k
Sleep 1s
Type "status"
Sleep 1s
Key Enter
Sleep 3s

# Exit
Type "q"
Sleep 1s
