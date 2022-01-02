declare-option -docstring %{ Path to wiki directory } str wiki_path
declare-option -docstring %{ Wiki helper executable path } str wiki_helper_cli kakoune-wiki

# Public interface

define-command -shell-script-candidates %{
    export kak_opt_wiki_path=$kak_opt_wiki_path
    "$kak_opt_wiki_helper_cli" -complete -all-markdown-files | sed 's/.md$//'
} -params 1 wiki-edit %{
    edit "%sh{ echo ""${kak_opt_wiki_path%/}/${1%.md}.md"" }"
}

define-command wiki-jump %{
    evaluate-commands -save-regs '^' %{
        execute-keys -save-regs '' Z
        try %{
            execute-keys -save-regs '' z
            wiki-select-md-link
            execute-keys <a-i>(
            evaluate-commands %{
                wiki-helper "-edit -edit-markdown %val{selection}"
            }
        } catch %{
            execute-keys -save-regs '' z
            wiki-select-mediawiki-link
            execute-keys <a-i>c\[\[,\]\]<ret>
            wiki-select-unfinished-mediawiki-link
            evaluate-commands %{ wiki-edit "%val{selection}" }
        } catch %{
            execute-keys -save-regs '' z
            fail %val{error}
        }
    }
}

map global normal <ret> ': wiki-jump'

# define-command -shell-script-completion %{
#     export kak_opt_wiki_path=$kak_opt_wiki_path
#     export kak_session=$kak_session
#     export kak_buffile=$kak_buffile
#     export kak_token_to_complete=$kak_token_to_complete
#     export kak_pos_in_token=$kak_pos_in_token
#     export wiki_debug=true
#     "$kak_opt_wiki_helper_cli" -complete -wiki-cmd $@
# } -params 1.. wiki %{
#     nop
# }

# Completion
############

declare-option -hidden completions wiki_completions

# Test if we are in unfinished mediawiki style link
define-command -hidden wiki-test-unfinished-mediawiki-link %{
    execute-keys -draft '<a-[>c\[\[,\]<ret>_<a-K>\n<ret>'
}

# Test if we are in unfinished markdown style link
define-command -hidden wiki-test-unfinished-md-link %!
    execute-keys -draft '{(HM<a-k>\[[^\n]*?\]\(.*<ret>_<a-K>\n<ret>'
!

define-command -hidden wiki-select-unfinished-md-link %{
    execute-keys <a-[>(_
    try %{ execute-keys <a-i>) }
    execute-keys <a-:><a-semicolon>
}

define-command -hidden wiki-select-unfinished-mediawiki-link %{
    execute-keys <a-[>c\[\[,\]<ret>_
    try %{ execute-keys <a-i>c\[\[,\]\]<ret> }
    try %{ execute-keys s\|\K.*\z<ret> }
    execute-keys <a-:><a-semicolon>
}

define-command -hidden wiki-populate-completion-header %{
    set-option buffer wiki_completions \
        "%val{cursor_line}.%val{cursor_column}+%val{selection_length}@%val{timestamp}"
    # echo -debug compl %opt{wiki_completions}
}

define-command wiki-enable-autocomplete %{
   set-option window completers option=wiki_completions
   hook -group wiki-autocomplete window InsertIdle .* %{
       wiki-update-completion
   }
   alias window complete wiki-update-completion
}

define-command wiki-disable-autocomplete -docstring "Disable wiki completion" %{
    set-option window completers %sh{ printf %s\\n "'${kak_opt_completers}'" | sed -e 's/option=wiki_completions://g' }
    remove-hooks window wiki-autocomplete
    unalias window complete wiki-complete
}


# Convert links

define-command -docstring %{
    Select all [[mediawiki]] style links
    Works only if selection_length == 1, otherwise it's nop
} wiki-select-mediawiki-link %{
    evaluate-commands -itersel %sh{
        [ "$kak_selection_length" -eq 1 ] &&
            printf "execute-keys '%s'\n" '<a-a>c\[\[,\]\]<ret><a-K>\n<ret>'
    }
}

define-command -docstring %{
	Convert [[link]] to [md](link)
} wiki-convert-link-to-md %{
    wiki-select-mediawiki-link
    evaluate-commands -save-regs '|' %{
        set-register '|' %{
            : $kak_session
            : $kak_opt_wiki_path
            : $kak_buffile
            "$kak_opt_wiki_helper_cli" -convert-link -to-markdown
        }
        execute-keys '|<ret>'
    }
}

define-command -docstring %{
    Select [md](link.md) style links
    Works only if selection_length == 1, otherwise it's nop
} wiki-select-md-link %{
    evaluate-commands -itersel %sh{
        [ "$kak_selection_length" -eq 1 ] &&
            printf "execute-keys '%s'\n" '<a-a>c\[,\)<ret><a-k>\[[^\n]*\]\([^\n]+\)<ret>'
    }
}

define-command -docstring %{
	Convert [md](link) to [[mediawiki]]
} wiki-convert-link-to-mediawiki %{
    wiki-select-md-link
    evaluate-commands -save-regs '|' %{
        set-register '|' %{
            : $kak_session
            : $kak_opt_wiki_path
            : $kak_buffile
            "$kak_opt_wiki_helper_cli" -convert-link -to-mediawiki
        }
        execute-keys '|<ret>'
    }
}


# Helper
########

define-command -hidden -params 1.. wiki-helper %{
    nop %sh{
    (
       # set -x
       export kak_opt_wiki_path="$kak_opt_wiki_path"
       export kak_session="$kak_session"
       export kak_client="$kak_client"
       export kak_buffile="$kak_buffile"

       # shellcheck disable=SC2086
       "$kak_opt_wiki_helper_cli" $@ | kak -p "$kak_session"
    ) > /dev/null &
    }
}

define-command -hidden wiki-update-completion %{
    evaluate-commands -draft %{
        try %{
            wiki-test-unfinished-md-link
            wiki-select-unfinished-md-link
            wiki-populate-completion-header
            wiki-helper "-complete -markdown-prefix %val{selection}"
        } catch %{
            wiki-test-unfinished-mediawiki-link
            wiki-select-unfinished-mediawiki-link
            wiki-populate-completion-header
            wiki-helper "-complete -mediawiki-prefix %val{selection}"
        } catch nop
    }
}

# Detection
###########

declare-option -hidden bool wiki_mode

hook -group wiki global BufCreate "%opt{wiki_path}.*\.md$" %{
    set-option buffer wiki_mode true
}

# TODO: remove

hook -group wiki global WinSetOption wiki_mode=true %{
    wiki-enable-autocomplete
}
