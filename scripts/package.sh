#!/bin/bash
set -eu

DIR=$(cd $(dirname ${0})/.. && pwd)
cd ${DIR}

NAME=${1}
VERSION=${2}
COMMIT=${3}

# Run Compile
./scripts/compile.sh ${COMMIT}

if [ -d pkg ];then
    rm -rf ./pkg/dist
fi

# Package all binary as .zip
mkdir -p ./pkg/dist/${VERSION}
for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
    PLATFORM_NAME=$(basename ${PLATFORM})
    ARCHIVE_NAME=${NAME}_${VERSION}_${PLATFORM_NAME}

    if [ $PLATFORM_NAME = "dist" ]; then
        continue
    fi

    pushd ${PLATFORM}
    zip ${DIR}/pkg/dist/${VERSION}/${ARCHIVE_NAME}.zip ./*
    popd
done

# Generate shasum
pushd ./pkg/dist/${VERSION}
shasum -a 256 * > ./${VERSION}_SHASUMS
popd
