#!/usr/bin/env bash

cwd=$(pwd)
dir="$cwd"

# installing dep
dep=$(which dep)
 if [[ -f "${dep}" ]]; then
    echo "dep found at ${dep}"
 else
    curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
    dep ensure
 fi
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

# todo: move to new service dir and run go generate