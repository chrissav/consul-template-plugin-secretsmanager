#!/bin/sh

# Ensure this script fails if anything errors
set -e

make deps
make build

mv dist/* ../artifact
