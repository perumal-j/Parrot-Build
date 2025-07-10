#!/bin/bash

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then
	ARCH="x86_64"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
	ARCH="aarch64"
else
	echo "Unsupported Architecture: $ARCH"
	exit 1
fi

jq -r '.body' /tmp/caido-latest.json | grep -oP 'https://[^"]+\.AppImage' | grep $ARCH