#!/bin/bash

REPO="yanklio/courls"
BINARY="courls"
DEST="/usr/local/bin"

# Get the latest tag
TAG=$(curl -s "https://api.github.com/repos/$REPO/releases/latest" | grep -o '"tag_name": ".*"' | sed 's/"tag_name": "//;s/"//')

# Direct link to the raw binary (not the tar.gz)
URL="https://github.com/$REPO/releases/download/$TAG/$BINARY"

echo "Downloading $BINARY version $TAG..."
curl -L $URL -o /tmp/$BINARY

# Install
chmod +x /tmp/$BINARY
echo "Installing $BINARY to $DEST..."
sudo mv /tmp/$BINARY $DEST/$BINARY

echo "Done!"
