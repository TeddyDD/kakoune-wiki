# kakoune-wiki

![icon](kakoune-wiki.png)

Personal wiki plugin for [**Kakoune**][kakoune]

## Motivation

Personal wiki is collection of interconnected notes you can store your
knowledge in. I wanted to create plugin for Kakoune that would make creating
such wiki as hassle-free as possible. My main issue was manual creation of
new pages and writing Markdown links by hand. When I've got an idea I want
to write it down right now, without navigating file system.  I want to be
able to create and link other notes during writing as well.

## Features

- find and edit any page using wiki command, no matter where you are in
  the file system
- quickly jump between wiki pages. Just point at link and press enter
- create interconnected pages with `@pagename` syntax that expands to standard
  Markdown links
- insert images with `@path/to/image.png` syntax
- toggle Markdown check-boxes with `<ret>` key in normal mode
- wiki is just bunch of Markdown files, you can process/edit them further with
  tools like [pandoc], [MdWiki], [markor] or any text editor. No vendor lock-ins
- minimal and simple — only essential features, script around 100 LOC

## Installation

**Only last stable releases of Kakoune is supported**

### Dependencies

To build helper from source you will need Go compiler 1.17.5 or newer and Make.

### [plug.kak]

```kak
plug "TeddyDD/kakoune-wiki" do %{
	# Optional, only if you have Go compiler and $GOPATH/bin in $PATH
	make go-install
} config %{
	# Your configuration here
}
```

### Manual

Source `rc/wiki.kak` or put it into autoloads directory. Ensure that
helper binary is in `PATH` or set `wiki_helper_cli` option to path to the
`kakoune-wiki` program.

- load `rc/wiki.kak` from your `kakrc`: `source path/to/rc/wiki.kak`
- or put `rc/wiki.kak` in your autoloads directory `~/.config/kak/autoload/`

## Setup

Then you have to choose directory for your wiki. Call following command from
your kakrc:

```
wiki-setup `/home/user/my/wiki/directory`
# or
wiki-setup %sh{ echo $HOME/wiki }
```

## Usage

### Create new page

To create wiki page use `wiki` command. Provide file name as parameter:

```
wiki cookies.md
```

This command creates file `cookies.md` in your wiki directory. You can also use
subdirectories for organization purpose. To create new page in subdirectory:

```
wiki recipes/cookies.md
```

Note that path is always relative to the wiki root directory.

### Edit existing page

To edit wiki page use (you guessed it) `wiki` command. You can press `TAB`
key for autocompletion:

```
wiki recipes/<TAB> # cycle through available pages
```

### Link pages

To reference other wiki page use `@tag` syntax. Type `@cookies<ret>` in insert
mode to create standard Markdown link to wiki page `cookies.md` in your wiki
directory. As alternative you can use `wiki_expand_tag` command in normal
mode when whole `@tag` is selected.  You can use subdirectories as well,
path is always relative to **wiki root directory**, however expanded link
**will be relative to currently edited wiki page**:

```
# editing recipes/cookies.md
@chocolate<ret>
expands to
[chocolate](../chocolate.md)
```

If page referenced by `@tag` does not exist it will be created. Directories
will be created as well.

If you press `<ret>` with cursor on link, Kakoune will follow link.

### Images

To insert image from your wiki directory use `@!image` syntax. Type
`@!image.jpg<ret>` to insert `![image.jpg](image.jpg)`. There is also
`wiki_expand_pic` command (`@!image` tag must be selected). You can use
subdirectories like in `@tag`.

### Checkboxes

You can toggle Markdown checkboxes on and off using `<ret>` key in normal
mode or `wiki_toggle_checkbox` command:

```
# TODO

- [ ] foo
- [ ] bar # press <ret>
- [X] bar # result

```

[plug.kak]: https://github.com/andreyorst/plug.kak
[kakoune]: http://kakoune.org/
[PR #2]: https://github.com/TeddyDD/kakoune-wiki/pull/2
[pandoc]: https://pandoc.org/
[MdWiki]: http://dynalon.github.io/mdwiki/
[markor]: https://github.com/gsantner/markor
