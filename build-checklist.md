1) Configure ip tables, firewall rules // Done
2) Configure SUDO no password // Done
3) Configure Obsidian Automation + Git // Partially Done. Obsidian Config to be automated
4) Install Sublime Text // Done
5) install Rust  // Done
6) feroxbuster // Done


APT tools: // Done
      - jq
      - pipx
      - ntpdate
      - flameshot
      - rlwrap
      - exiftool
      - rsyslog
      - ca-certificates
      - curl
      - gh

AD Tools:
certipy // Done
sshuttle // Done
impacket (reinstall)  // Done
pywhisker // Done
bloodyad // Done
Install latest NetExec // Done
targetedkerberoast.py // Done

Watch 0xdf video and Automate Python tools installation // Done

Write Script to install github tools: // Done
chisel (linux/windows) // Done
PEASS-ng (linux/windows) // Done
Chainsaw // Done
BloodhoundAD/bloodhound
SharpCollection // Done
SecLists // Done
SharpHound // Done


Install Gem Tools: // Done
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


====================
To Be Done
====================

3) Configure Obsidian Automation + Git // Partially Done. Obsidian Config to be automated

Remove Unofficial Docker Images & reinstall:

- name: "Add Docker keyring to apt"
  apt_key:
    url: "https://download.docker.com/linux/debian/gpg"
    state: present
  become: true
  become_method: sudo


3) Bloodhound-CE
4) Bloodhound Injestor

Configure Task Bar, change desktop wallpaper