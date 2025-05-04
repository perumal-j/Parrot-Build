1) Configure ip tables, firewall rules
2) Configure SUDO no password
2) Configure Obsidian Automation
5) Install Sublime Text
6) install Rust
6) Tools:
flameshot
rlwrap
feroxbuster

AD Tools:
certipy
impacket (reinstall) 
pywhisker
bloodyad
Install latest NetExec

Watch 0xdf video and Automate Python tools installation

Write Script to install github tools:
chisel (linux/windows)
PEASS-ng (linux/windows)
Chainsaw
BloodhoundAD/bloodhound
SharpCollection
SecLists
SharpHound


Install Gem Tools:
    - logger
    - stringio
    - winrm
    - builder
    - erubi
    - gssapi
    - gyoku
    - httpclient
    - logging
    - little-plugger
    - nori
    - rubyntlm
    - winrm-fs
    - evil-winrm


Remove Unofficial Docker Images & reinstall:

- name: "Add Docker keyring to apt"
  apt_key:
    url: "https://download.docker.com/linux/debian/gpg"
    state: present
  become: true
  become_method: sudo

APT tools:
      - jq
      - pipx
      - ntpdate
      - flameshot
      - exiftool
      - rsyslog
      - ca-certificates
      - curl
      - gh


3) Bloodhound-CE
4) Bloodhound Injestor

Configure Task Bar, change desktop wallpaper