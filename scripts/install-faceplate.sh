#!/usr/bin/env zsh

script_dir=$(dirname "$0")

if [[ "$#" -lt 2 ]]; then
    echo "usage: $(basename "${0}") from-file to-file"
    exit 1
fi

${script_dir}/install-svg.sh ${1} ${2} --export-id=faceplate --export-id-only
