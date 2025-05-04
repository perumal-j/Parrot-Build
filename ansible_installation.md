sudo apt install pipx
pipx install --include-deps ansible

# Install locales package if missing
sudo apt update && sudo apt install locales

# Generate UTF-8 locale
sudo locale-gen en_US.UTF-8

# Update system-wide locale settings
sudo update-locale LANG=en_US.UTF-8 LC_ALL=en_US.UTF-8

# Apply to current session
export LANG=en_US.UTF-8
export LC_ALL=en_US.UTF-8

ansible-galaxy install -r requirements.yml
