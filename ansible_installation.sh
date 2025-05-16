#!/bin/bash
sudo apt-get update
sudo apt-get install pipx -y
pipx install --include-deps ansible

# Install locales package if missing
sudo apt-get install locales -y

# Generate UTF-8 locale
sudo locale-gen en_US.UTF-8

# Update system-wide locale settings
sudo update-locale LANG=en_US.UTF-8 LC_ALL=en_US.UTF-8

# Apply to current session
export LANG=en_US.UTF-8
export LC_ALL=en_US.UTF-8

ansible-galaxy install -r requirements.yml
