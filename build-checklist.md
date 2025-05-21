1) Configure ip tables, firewall rules // Done
2) Configure SUDO no password
3) Configure Obsidian Automation
4) Install Sublime Text // Done
5) install Rust  // Done
6) feroxbuster


APT tools:
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
certipy
impacket (reinstall) 
pywhisker
bloodyad
Install latest NetExec
targetedkerberoast
pip install pygments  ---> for GDB syntax highlighting

Watch 0xdf video and Automate Python tools installation

Write Script to install github tools:
chisel (linux/windows)
PEASS-ng (linux/windows)
Chainsaw
BloodhoundAD/bloodhound
sshuttle
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


3) Bloodhound-CE
4) Bloodhound Injestor

Configure Task Bar, change desktop wallpaper