#!/bin/bash

# Build script for all Go projects

echo "🔨 Building all Go projects..."

# Find all directories with go.mod files
for dir in $(find . -name "go.mod" -not -path "./vendor/*" | xargs dirname); do
    echo "Building project in $dir..."
    cd "$dir"
    
    if go build .; then
        echo "✅ Successfully built $dir"
    else
        echo "❌ Failed to build $dir"
    fi
    
    cd - > /dev/null
done

echo "🎉 Build process complete!"
