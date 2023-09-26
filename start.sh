#!usr/bin/bash
pwd=$(pwd)
output="${pwd}/build/go.ui.html.template.out"
set -x
go build -C src -o "${output}" && "${output}"