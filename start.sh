#!usr/bin/bash
output="build/go.ui.html.template.out"
set -x
go build -o "${output}" && "./${output}"