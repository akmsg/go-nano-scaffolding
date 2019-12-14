#!/usr/bin/env bash

cwd=$(pwd)
dir="$cwd"

# source pkg to compile
mainPkg="${dir}/cmd/"

# destination for output
execDir="${dir}/bin"
mkdir -p "${execDir}"
exec="${execDir}/nanokit.bin"

# build the binary
go build -o "${exec}" "${mainPkg}" || (echo "compilation failed :(" && exit 1)


cd "${cwd}" || exit 2
"${exec}" "${@}"