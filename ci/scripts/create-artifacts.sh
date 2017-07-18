#!/bin/sh

# Ensure this script fails if anything errors
set -e

CWD=$(pwd)
mkdir -p artifacts

#Create release body
cd ./source-code
git log -n1 --oneline > $CWD/artifacts/release-body

# Create tarball
cd $CWD/artifact

echo "Creating release"
tar -czf ssm.tar.gz ssm
mv ssm.tar.gz ${CWD}/artifacts/
