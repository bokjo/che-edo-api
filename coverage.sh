#!/bin/bash

COVER_DIR=cover
mkdir -p ${COVER_DIR}

PKG_LIST=$(go list ./... | grep -v /vendor/)

for package in ${PKG_LIST}; do
    go test -covermode=count -coverprofile "${COVER_DIR}/${package##*/}.cov" "$package" ;
done

tail -q -n +2 ${COVER_DIR}/*.cov >> ${COVER_DIR}/coverage.cov

go tool cover -func=${COVER_DIR}/coverage.cov
go tool cover -html=${COVER_DIR}/coverage.cov -o ${COVER_DIR}/coverage.html