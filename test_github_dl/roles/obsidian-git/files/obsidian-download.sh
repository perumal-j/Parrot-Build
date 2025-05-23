#!/bin/bash

ARCH=$(uname -m)
VERSION=$(curl -s https://api.github.com/repos/obsidianmd/obsidian-releases/releases/latest | jq -r .name)

if [ "$ARCH" = "x86_64" ]; then
	ARCH="Obsidian-${VERSION}.AppImage"
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
	ARCH="Obsidian-${VERSION}-arm64.AppImage"
else
	echo "Unsupported Architecture: $ARCH"
	exit 1
fi

echo $ARCH

curl -s https://api.github.com/repos/obsidianmd/obsidian-releases/releases/latest -o /tmp/obsidian-latest.json

jq -r '.assets[].browser_download_url' /tmp/obsidian-latest.json | grep "$ARCH" | xargs curl -sL -o /tmp/obsidian-latest.AppImage

chmod +x /tmp/obsidian-latest.AppImage

sudo mv /tmp/obsidian-latest.AppImage /usr/local/bin/obsidian