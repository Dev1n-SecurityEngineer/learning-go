#!/bin/bash

# Setup script for Go learning environment

echo "🚀 Setting up Go learning environment..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first."
    exit 1
fi

echo "✅ Go version: $(go version)"

# Initialize workspace if not exists
if [ ! -f "go.work" ]; then
    echo "📝 Initializing Go workspace..."
    go work init
fi

# Create common project templates
echo "📂 Creating project templates..."

# Create a basic CLI tool template
mkdir -p projects/cli-tools/template
cat > projects/cli-tools/template/main.go << 'EOF'
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: program <command>")
		os.Exit(1)
	}
	
	command := os.Args[1]
	fmt.Printf("Executing command: %s\n", command)
}
EOF

# Create a basic web server template
mkdir -p projects/web-servers/template
cat > projects/web-servers/template/main.go << 'EOF'
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
EOF

echo "✅ Setup complete! Happy learning! 🎉"
