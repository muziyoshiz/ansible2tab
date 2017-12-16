#!/bin/bash
set -eu

DIR=$(cd $(dirname ${0})/.. && pwd)
cd ${DIR}

COMMIT=${1}

XC_ARCH=${XC_ARCH:-386 amd64}
XC_OS=${XC_OS:-darwin linux windows}

rm -rf pkg/
gox \
    -ldflags "-X main.commit=${COMMIT}" \
    -parallel=5 \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
-output "pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"
