#!/bin/bash
set -xe
# --no-bufferを付けなくても機能する
curl --no-buffer -v http://localhost:8080/call