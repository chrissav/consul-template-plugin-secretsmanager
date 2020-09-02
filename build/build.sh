#!/bin/bash

set -e

# Check if VERSION variable set and not empty, otherwise set to default value
if [ -z "$VERSION" ]; then
  VERSION="0.0.1-dev"
fi
echo "Building application version $VERSION"

# Build binaries for different platforms/architectures
ARCH=amd64
for OS in linux darwin; do
  echo "Building binaries for $OS/$ARCH..."
  GOARCH=$ARCH GOOS=$OS CGO_ENABLED=0 go build -ldflags "-s -w" -ldflags "-X main.version=${VERSION}" -o "dist/secretsmanager_${OS}_${ARCH}" secretsmanager.go
done
