#!/bin/bash

uv run --with impacket --with "pyOpenSSL<23.3.0" /usr/local/bin/ntlmrelayx.py "$@"
