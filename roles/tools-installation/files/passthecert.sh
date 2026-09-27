#!/bin/bash

uv run --with ldap3 --with ldapdomaindump --with impacket /opt/PassTheCert/Python/passthecert.py "$@"
