declare-option -docstring %{ Path to wiki directory } str wiki_path

declare-option -hidden str wiki_relative_patch_program %{ perl -e 'use File::Spec; print File::Spec->abs2rel(@ARGV) . "\n"' }
define-command -hidden -params 1 wiki_setup %{
	%sh{
		echo "set-option global wiki_path $1"
		echo "hook global BufCreate $1/.+\.md %{ wiki_enable }"
	}
}

define-command wiki -params 1  \
-docstring %{ wiki [file]: Edit existing wiki page. } \
-shell-candidates %{ find $kak_opt_wiki_path -type f -name '*.md' }  \
%{
        edit %arg{1}
}

define-command wiki_enable %{
    add-highlighter buffer group wiki
    add-highlighter buffer/wiki regex '\B@\w+' 0:link
    add-highlighter buffer/wiki regex '\[\w+\]' 0:link
    hook buffer InsertChar \n -group wiki %{
        evaluate-commands %{ try %{ 
    		execute-keys -draft %{
				2h<a-b><a-k>\B@\w+<ret>
				:wiki_expand_tag<ret>
    		}
    		execute-keys <esc>hi
		} }
    }
    hook buffer NormalKey <ret> -group wiki %{
		wiki_follow_link
		wiki_toggle_checkbox
    }
}

define-command wiki_disable %{
    remove-highlighter buffer/wiki
	remove-hooks buffer wiki
}

define-command wiki_expand_tag \
-docstring %{ Expands tag from @filename form to [filename](filename.md)
Creates empty markdown file in wiki_path if not exist
selection must be somewhere on @tag } %{
	execute-keys -draft %{
        <a-i>Ws[^@]+<ret>yi<backspace>[<esc>a]<esc>a(<esc>pA.md)<esc>
        :wiki_new_page <c-r>"<ret>
	}
	execute-keys 'f);'
}

define-command -params 1 \
-docstring %{ wiki_new_page [name]: create new wiki page in wiki_path if not exists } \
wiki_new_page %{
    %sh{
        dir="$(dirname $kak_opt_wiki_path/$1.md)"
        mkdir -p "$dir"
        touch "$kak_opt_wiki_path/$1.md"
    }
}

define-command wiki_follow_link \
-docstring %{ Follow markdown link and open file if exists } %{
    evaluate-commands %{ try %{
    	execute-keys %{
    		<esc><a-a>c\[,\)<ret><a-:>
			<a-i>b
    	}
        evaluate-commands -try-client %opt{jumpclient} edit -existing %sh{ echo $kak_selection }
        focus %opt{jumpclient}
	}}
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
