#!/bin/bash

set -e

# Get rid of existing binaries
rm -f dist/secretsmanager*

# Check if VERSION variable set and not empty, otherwise set to default value
if [ -z "$VERSION" ]; then
  VERSION="0.0.1-dev"
fi
echo "Building application version $VERSION"

# Build binaries for different platforms/architectures
# Set OS to darwin to build for OSX
OS=linux
ARCH=amd64
echo "Building binaries for $OS/$ARCH..."
GOARCH=$ARCH GOOS=$OS CGO_ENABLED=0 go build -ldflags "-s -w" -ldflags "-X main.version=${VERSION}" -o "dist/secretsmanager" secretsmanager.go
