#!/bin/sh
set -e
curl -fsSL -o tool.tar.gz https://github.com/acme/tool/releases/download/v1.2.3/tool.tar.gz
curl -fsSL -o checksums.txt https://github.com/acme/tool/releases/download/v1.2.3/checksums.txt
sha256sum --ignore-missing -c checksums.txt
tar -xzf tool.tar.gz
