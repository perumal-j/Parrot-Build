#!/bin/bash

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then
	ARCH="x64"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
	ARCH="aarch64"
else
	echo "Unsupported Architecture: $ARCH"
	exit 1
fi

DOWNLOAD_URL="https://api.adoptium.net/v3/binary/latest/21/ga/linux/${ARCH}/jdk/hotspot/normal/eclipse"

curl -sL "$DOWNLOAD_URL" -o "/tmp/openjdk-latest.tar.gz"
