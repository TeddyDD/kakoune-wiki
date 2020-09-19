declare-option -docstring %{ Path to wiki directory } str wiki_path

# program that outputs relative path given two absolute as params
declare-option -hidden str wiki_relative_path_program %{ perl -e 'use File::Spec; print File::Spec->abs2rel(@ARGV) . "\n"' }


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
