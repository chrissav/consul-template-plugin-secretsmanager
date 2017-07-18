#!/bin/sh

set -e

NO_COLOR='\033[0m'
OK_COLOR='\033[32;01m'
ERROR_COLOR='\033[31;01m'
WARN_COLOR='\033[33;01m'
PASS="${OK_COLOR}PASS ${NO_COLOR}"
FAIL="${ERROR_COLOR}FAIL ${NO_COLOR}"


echo "${OK_COLOR}Running tests: ${NO_COLOR}"
go test -cover -v

echo "${OK_COLOR}Formatting: ${NO_COLOR}"
ERRS=$(find . -type f -name \*.go | xargs gofmt -l 2>&1 || true)
if [ -n "${ERRS}" ]; then
    echo "${ERROR_COLOR}FAIL - the following files need to be gofmt'ed: ${NO_COLOR}"
    for e in ${ERRS}; do
        echo "    $e"
    done
    exit 1
fi
echo ${PASS}
