.PHONY: run build demos clean install-vhs

# Build the demo application
build:
	go build -o demo .

# Run the demo application
run:
	go run .

# Install VHS if not already installed
install-vhs:
	@which vhs > /dev/null || go install github.com/charmbracelet/vhs@latest

# Generate all demo GIFs
demos: install-vhs
	@echo "🎬 Generating demo GIFs..."
	@mkdir -p demos
	vhs < demos/basic-usage.tape
	vhs < demos/command-filtering.tape
	vhs < demos/navigation.tape
	@echo "✅ Demo GIFs generated in demos/ directory"

# Generate a specific demo
demo-basic: install-vhs
	vhs < demos/basic-usage.tape

demo-filtering: install-vhs
	vhs < demos/command-filtering.tape

demo-navigation: install-vhs
	vhs < demos/navigation.tape

# Clean generated files
clean:
	rm -f demo
	rm -f demos/*.gif

# Test the application
test:
	go test -v

# Show help
help:
	@echo "Available commands:"
	@echo "  make run           - Run the demo application"
	@echo "  make build         - Build the demo application"
	@echo "  make demos         - Generate all demo GIFs"
	@echo "  make demo-basic    - Generate basic usage demo"
	@echo "  make demo-filtering - Generate command filtering demo"
	@echo "  make demo-navigation - Generate navigation demo"
	@echo "  make install-vhs   - Install VHS tool"
	@echo "  make clean         - Clean generated files"
	@echo "  make test          - Run tests"
	@echo "  make help          - Show this help"
