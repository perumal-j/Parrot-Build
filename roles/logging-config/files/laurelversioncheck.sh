#!/bin/bash

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then
	ARCH="x86_64-glibc"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
	ARCH="aarch64-glibc"
else
	echo "Unsupported Architecture: $ARCH"
	exit 1
fi

jq -r --arg ARCH "$ARCH" '.assets[].browser_download_url | select(contains($ARCH))' /tmp/latest.json
