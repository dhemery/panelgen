#!/usr/bin/env zsh

script_dir=$(dirname "$0")

if [[ "$#" -lt 2 ]]; then
    echo "usage: $(basename "${0}") from-dir to-dir"
    exit 1
fi

from_dir="${1}"
to_dir="${2}"

for file in ${from_dir}/*
do
    frame_name="$(basename ${file})"
    module_slug="$(basename $(dirname ${file}))"
    to_file="${to_dir}/${module_slug}/${frame_name}"
    "${script_dir}/install-svg.sh" "${1}/${frame_name}" "${to_file}"
done
