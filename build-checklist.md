1) Configure ip tables, firewall rules // Done
2) Configure SUDO no password // Done
3) Configure Obsidian Automation
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
certipy
sshuttle // Done
impacket (reinstall)  // Done
pywhisker // Done
bloodyad // Done
Install latest NetExec // Done
targetedkerberoast.py // Done
pip install pygments  ---> for GDB syntax highlighting

Watch 0xdf video and Automate Python tools installation

Write Script to install github tools: // Done
chisel (linux/windows)
PEASS-ng (linux/windows)
Chainsaw
BloodhoundAD/bloodhound
SharpCollection
SecLists
SharpHound


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