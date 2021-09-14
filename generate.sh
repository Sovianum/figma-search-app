#!/bin/bash -e

export GOENV=$(dirname "$0")/go.env
export GO=$( [ -z "$GOROOT" ] && echo "go" || echo "${GOROOT}/bin/go" )

cd `dirname "$0"`
export GOPATH=$("$GO" env GOPATH)

wire ./src/... && echo Success!
