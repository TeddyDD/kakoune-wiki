#!/bin/sh

# kak_escape vendored from prelude
kak_escape() {
	for text
	do
		printf "'"
		while true
		do
			case "$text" in
			*"'"*)
				head=${text%%"'"*}
				tail=${text#*"'"}
				printf "%s''" "$head"
				text=$tail
				;;
			*)
				printf "%s' " "$text"
				break
				;;
			esac
		done
	done
	printf "${KAK_ESCAPE_EOF:-\n}"
}

kak_escape_partial() {
	KAK_ESCAPE_EOF=' ' kak_escape "$@"
}

trim() {
	sed -E 's![.]{2}/!!'
}

python_relpath() {
	python - "$1" "$2" <<-'EOF' | trim
		import os, sys
		print(os.path.relpath(sys.argv[1], sys.argv[2]))
	EOF
}

perl_relpath() {
	perl -e 'use File::Spec; print File::Spec->abs2rel(@ARGV) . "\n"' "$1" "$2" |
		trim
}

clean_find_result() {
	awk '{ gsub(/^\.\//, "", $0); gsub(/\.md$/, "", $0); print }'
}

list_all_pages() (
	cd "$1" || :
	find . -iname '*.md' -type f
)

list_all_pages_abs() {
	find "$1" -iname '*.md' -type f
}

normalize() {
	clean_find_result
}

get_edit_command() {
	wiki_path="$1"
	page_name="$2"
	page_name=$(echo "$page_name" | normalize)
	printf "edit '%s/%s.md'\n" "$wiki_path" "$page_name"
}

mediawiki_completions() {
	wiki_path="$1"
	selection="$2"
	list_all_pages "$wiki_path" |
		normalize |
		grep -iF "$selection" |
		awk '{print $0"||"$0 }' |
		while read -r l
		do
			kak_escape "$l"
		done |
		tr -d "\n"
}

# format_md_completion() {
# 	awk -F '/' '
# 	{
#     	pp=$NF;
#     	gsub(/\.md$/, "", pp);
#     	print $0"||"pp;
# 	}
# 	'
# }

md_completions() {
	wiki_path="$1"
	selection="$2"
	kak_buffile="$3"
	list_all_pages_abs "$wiki_path" |
		grep -iF "$selection" |
		while read -r page
		do
			kak_escape "$(
				perl_relpath "$page" "$kak_buffile" |
					awk '{print $0"||"$0 }'
			)"
		done |
		tr -d "\n"
}
