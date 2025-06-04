#!/bin/bash

ARCH=$(uname -m)
VERSION=$(curl -s https://api.github.com/repos/obsidianmd/obsidian-releases/releases/latest | jq -r .name)

if [ "$ARCH" = "x86_64" ]; then
	ARCH="Obsidian-${VERSION}.AppImage"
	curl -s https://api.github.com/repos/obsidianmd/obsidian-releases/releases/latest -o /tmp/obsidian-latest.json
	jq -r '.assets[].browser_download_url' /tmp/obsidian-latest.json | grep "$ARCH" | xargs curl -sL -o /tmp/obsidian-latest.AppImage
	chmod +x /tmp/obsidian-latest.AppImage
	sudo mv /tmp/obsidian-latest.AppImage /usr/local/bin/obsidian
elif [ "$ARCH" = "aarch64" ] || [ "$ARCH" = "arm64" ]; then
	ARCH="Obsidian-${VERSION}-arm64.AppImage"
	curl -s https://api.github.com/repos/obsidianmd/obsidian-releases/releases/latest -o /tmp/obsidian-latest.json
	jq -r '.assets[].browser_download_url' /tmp/obsidian-latest.json | grep "$ARCH" | xargs curl -sL -o /tmp/obsidian-latest.AppImage
	chmod +x /tmp/obsidian-latest.AppImage
	sudo mv /tmp/obsidian-latest.AppImage /opt/obsidian
	sudo bash -c "echo -e '#!/bin/bash\n/opt/obsidian --no-sandbox' > /usr/local/bin/obsidian"
	sudo chmod +x /usr/local/bin/obsidian
else
	echo "Unsupported Architecture: $ARCH"
	exit 1
fi


# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/debian \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update


# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update