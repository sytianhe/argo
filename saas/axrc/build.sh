#!/bin/bash

set -xe

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
. $DIR/../build_env.sh

gofmt -w $SRCROOT/saas/axrc/src
go tool vet -composites=false $SRCROOT/saas/axrc/src/

go install applatix.io/axrc/server
