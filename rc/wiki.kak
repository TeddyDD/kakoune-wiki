declare-option -docstring "Path to wiki directory" \
    str wiki_path

set-option global wiki_path /tmp/wiki

# python_relpath is loaded from wiki.sh
declare-option -hidden \
    -docstring "command that returns relative path between two files" \
    str wiki_relative_path python_relpath

declare-option -hidden \
    str wiki_source %val{source}

declare-option -hidden \
	completions wiki_completers

# select [[mediawiki style link|mediawiki]]
define-command -hidden wiki-select-mediawiki-link %{
    execute-keys <a-a>c\[\[,\]\]<ret>
}

define-command -hidden wiki-select-markdown-link %{
    execute-keys <a-a>c\[,\)<ret><a-k>\[.+?\]\(.+?\)<ret>
}

define-command -hidden -params 1 wiki-touch-page %{
    nop %sh{
        cd "$kak_opt_wiki_path" || echo fail "Wiki path not exists"
        dir="${1%/*}"
        [ ! -d "$dir" -a ! -n "$dir" ] && mkdir -p "$dir"
        touch "${1%.md}.md"
    }
}

define-command -params 1.. -docstring 'grep wiki path' wiki-grep %{
}

define-command -shell-script-candidates %{
    . "${kak_opt_wiki_source%/*}/wiki.sh"
    list_all_pages "$kak_opt_wiki_path" | clean_find_result
} \
-docstring 'Open or create a wiki page' \
-params 1.. \
wiki-edit %{
    evaluate-commands %{
        wiki-touch-page "%arg{@}"
        wiki-edit-impl "%arg{@}"
    }
}

define-command -hidden -params 1.. wiki-edit-impl %{
	evaluate-commands %sh{
    . "${kak_opt_wiki_source%/*}/wiki.sh"
	get_edit_command "$kak_opt_wiki_path" "$@"
	}
}


# Test if we are in unfinished mediawiki style link
define-command -hidden wiki-test-unfinished-mediawiki-link %{
    execute-keys -draft [[H<a-k>\[\[<ret> _ <a-K>\n<ret>
}

# Test if we are in unfinished markdown style link
define-command -hidden wiki-test-unfinished-md-link %{
    execute-keys -draft [(HM<a-k>\[.+\]\(.*<ret> _ <a-K>\n<ret>
}

define-command -hidden wiki-select-unfinished-md-link %{
	execute-keys <a-[>(_
}

define-command -hidden wiki-select-unfinished-mediawiki-link %{
	execute-keys <a-[>[H<a-K>\A\[\[<ret>L_
	try %{ execute-keys s\|\K.+\z<ret>_ }
}

define-command -hidden wiki-populate-completion-header %{
    set-option window wiki_completers \
                    "%val{cursor_line}.%val{cursor_column}+%val{selection_length}@%val{timestamp}"
}

define-command -hidden wiki-populate-mediawiki-completion %{
	evaluate-commands %sh{
        . "${kak_opt_wiki_source%/*}/wiki.sh"
        printf "set-option -add window wiki_completers %s\n" \
            "$(mediawiki_completions "$kak_opt_wiki_path" "$kak_selection")"
	}
}

define-command -hidden wiki-populate-md-completion %{
	evaluate-commands %sh{
        . "${kak_opt_wiki_source%/*}/wiki.sh"
        [ -e "$kak_buffile" ] || exit 0
        printf "set-option -add window wiki_completers %s\n" \
            "$(md_completions "$kak_opt_wiki_path" "$kak_selection" "$kak_buffile")"
	}
}

define-command -hidden wiki-populate-md-completion-async %{
    nop %sh{
        (
            . "${kak_opt_wiki_source%/*}/wiki.sh"
            header="${kak_cursor_line}.${kak_cursor_column}+${kak_selection_length}@${kak_timestamp}"
            compl=$(md_completions "$kak_opt_wiki_path" "$kak_selection" "$kak_buffile")
            kak_escape evaluate-commands -client ${kak_client} set-option window wiki_completers ${header} ${compl} |
                kak -p "${kak_session}"
        ) > /dev/null 2>&1 < /dev/null &
    }
}

provide-module markdown-wiki %{

}

# Hooks

hook global WinSetOption filetype=markdown-wiki %{
    require-module markdown-wiki
    set-option window completers option=wiki_completers %opt{completers}
    hook window -group wiki-autocomplete InsertIdle .* %{
        try %{
            wiki-test-unfinished-mediawiki-link
            evaluate-commands -draft %{
                wiki-select-unfinished-mediawiki-link
                wiki-populate-completion-header
                wiki-populate-mediawiki-completion
            }
        } catch %{
			wiki-test-unfinished-md-link
			evaluate-commands -draft %{
    			wiki-select-unfinished-md-link
    			wiki-populate-md-completion-async
			}
        } catch nop
    }

    hook -once -always window WinSetOption filetype=.* %{
        # remove-hooks window awk-.+
        evaluate-commands %sh{
            printf "set-option window completers %s\n" \
                $(printf %s "${kak_opt_completers}" | sed -e "s/'option=wiki_completers'//g")
        }
    }
}

#### OLD

define-command -hidden -params 1 wiki-setup %{
    evaluate-commands %sh{
        echo "set-option global wiki_path $1"
        echo "hook global BufCreate $1/.+\.md %{ wiki_enable }"
    }
}

define-command wiki -params 1  \
-docstring %{ wiki [file.md]: Edit or create wiki page } \
-shell-script-candidates %{ cd $kak_opt_wiki_path; find . -type f -name '*.md' | sed -e 's/^\.\///' }  \
%{ evaluate-commands %sh{
    if [ ! -e "$kak_opt_wiki_path/$1" ]; then
        echo "wiki_new_page %{$1}"
    fi
    echo "edit %{$kak_opt_wiki_path/$1}"
}}

define-command wiki_enable %{
    add-highlighter buffer/wiki group
    add-highlighter buffer/wiki/tag regex '\B@\S+' 0:link
    add-highlighter buffer/wiki/link regex '\[\w+\]' 0:link
    hook buffer InsertKey '<ret>' -group wiki %{
        evaluate-commands %{ try %{
            execute-keys -draft %{
                <a-b><a-k>\A@!\w+<ret>
                :wiki_expand_pic<ret>
        }} catch %{
            try %{ execute-keys -draft %{
                    <a-b><a-k>\A@\w+<ret>
                    :wiki_expand_tag<ret>
                }
                execute-keys <backspace>
            }
        }}
    }
    hook buffer NormalKey '<ret>' -group wiki %{
        try %{ wiki_follow_link }
        try %{ wiki_toggle_checkbox }
    }
}

define-command wiki_disable %{
    remove-highlighter buffer/wiki
    remove-hooks buffer wiki
}

define-command wiki_expand_tag \
    -docstring %{ Expands tag from @filename form to [filename](filename.md)
    Creates empty markdown file in wiki_path if not exist. Selection must be
    somewhere on @tag and @tag should not contain extension } %{
    evaluate-commands %sh{
        this="$kak_buffile"
        tag=$(echo $kak_selection | sed -e 's/^\@//')
        other="$kak_opt_wiki_path/$tag.md"
        relative=$(eval "$kak_opt_wiki_relative_path_program" "$other" $(dirname "$this"))
        # sanity chceck
        echo "execute-keys  -draft '<a-k>\A@[^@]+<ret>'"
        echo "execute-keys \"c[$tag]($relative)<esc>\""
        echo "wiki_new_page \"$tag\""
    }
}

define-command wiki_expand_pic \
-docstring %{ Expands images from @!filename.png form to ![filename.png](filename.png)} %{
    evaluate-commands %sh{
        this="$kak_buffile"
        tag=$(echo $kak_selection | sed -e 's/^\@!//')
        other="$kak_opt_wiki_path/$tag"
        relative=$(eval "$kak_opt_wiki_relative_path_program" "$other" $(dirname "$this"))
        # sanity check
        echo execute-keys -draft '<a-k>^@\+[^@!]+'
        echo execute-keys "c![$tag]($relative)<esc>"
    }
}

define-command -params 1 -hidden \
-docstring %{ wiki_new_page [name]: create new wiki page in wiki_path if not exists } \
wiki_new_page %{
    nop %sh{
        file="$kak_opt_wiki_path/${1%.md}"
        mkdir -p "$(dirname $file)"
        touch "${file}.md"
    }
}

define-command wiki_follow_link \
-docstring %{ Follow markdown link and open file if exists } %{
    evaluate-commands %{
        execute-keys %{
            <esc><a-a>c\[,\)<ret><a-:>
            <a-i>b
        }
        evaluate-commands -try-client %opt{jumpclient} edit -existing %sh{
            echo "'${kak_buffile%/*.md}/$kak_selection'"
        }
        try %{ focus %opt{jumpclient} }
    }
}

define-command wiki_toggle_checkbox \
-docstring "Toggle markdown checkbox in current line" %{
    try %{
        try %{
            execute-keys -draft %{
                <esc><space>;xs-\s\[\s\]<ret><a-i>[rX
        }} catch %{
            execute-keys -draft %{
                <esc><space>;xs-\s\[X\]<ret><a-i>[r<space>
    }}}
}
