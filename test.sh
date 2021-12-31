#!/bin/sh
set -x

make
export wiki_debug=true
kak -e "
set-option global debug hooks
source rc/wiki.kak
set-option global wiki_helper_cli '$(pwd)/kakoune-wiki'
set-option global wiki_path ./testdata/
edit testdata/os/linux.md
set-option buffer debug shell|commands"
