#!/bin/bash

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

if [ "$ARCH" = "x86_64" ]; then
	ARCH="amd64"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
	ARCH="arm64"
else
	echo "Unsupported Architecture: $ARCH"
	exit 1
fi

LATEST_VERSION=$(curl -s https://go.dev/VERSION?m=text | head -n 1)

TAR_BALL="${LATEST_VERSION}.${OS}-${ARCH}.tar.gz"

DOWNLOAD_URL="https://go.dev/dl/${TAR_BALL}"

curl -sL "$DOWNLOAD_URL" -o "/tmp/${TAR_BALL}"

sudo rm -rf /usr/local/go

sudo tar -C /usr/local -xzf "/tmp/${TAR_BALL}"
