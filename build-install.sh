#!/bin/bash
set -euo pipefail

# Build script for todoist CLI
# Source: https://github.com/sachaos/todoist
# Installs to: ~/prj/util/bin/todoist

INSTALL_DIR="$HOME/prj/util/bin"
BINARY_NAME="todoist"

echo "==> Building todoist CLI..."

# Check for Go
if ! command -v go &> /dev/null; then
    echo "Error: go is not installed"
    exit 1
fi

# Ensure install directory exists
mkdir -p "$INSTALL_DIR"

# Build the binary (not install - we want control over destination)
echo "==> Running: go build"
go build -o "$BINARY_NAME"

# Move to install directory
echo "==> Installing to: $INSTALL_DIR/$BINARY_NAME"
mv "$BINARY_NAME" "$INSTALL_DIR/$BINARY_NAME"
chmod +x "$INSTALL_DIR/$BINARY_NAME"

# Verify installation
if [ -f "$INSTALL_DIR/$BINARY_NAME" ]; then
    echo "==> Success! Installed to: $INSTALL_DIR/$BINARY_NAME"
    echo "==> Size: $(ls -lh "$INSTALL_DIR/$BINARY_NAME" | awk '{print $5}')"
else
    echo "==> Error: Installation failed"
    exit 1
fi

echo ""
echo "==> Next steps:"
echo "   1. Ensure ~/prj/util/bin is in your PATH"
echo "   2. Run 'todoist' to configure your API token"
echo "   3. Get your token from: https://todoist.com/prefs/integrations"
echo "   4. Run 'todoist sync' to fetch your tasks"
