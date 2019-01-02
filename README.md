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
- quickly jump between wiki pages. Just point on link and press enter
- create interconnected pages with `@pagename` syntax that expands to standard Markdown links
- insert images with `@path/to/image.png` syntax
- toggle Markdown check-boxes with `<ret>` key in normal mode
- wiki is just bunch of Markdown files, you can process/edit them further
with tools like [pandoc](https://pandoc.org/),
[MdWiki](http://dynalon.github.io/mdwiki/),
[markor](https://github.com/gsantner/markor) or any text editor. No vendor lock-ins
- minimal and simple — only essential features, script around 100 LOC

## Installation

You can either:

- load `rc/wiki.kak` from your `kakrc`: `source path/to/rc/wiki.kak`
- put `rc/wiki.kak` in your autoloads directory `~/.config/kak/autoload/`
- use [plug.kak] - plugin manager

Then you have to choose directory for your wiki. Call following command from
your kakrc:

```
wiki_setup `/home/user/my/wiki/directory`
```

**only last stable releases of Kakoune is supported**

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

To edit wiki page use (you guessed it) `wiki` command. You can press `TAB` key for autocompletion:

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

### Images

To insert image from your wiki directory use `@!image` syntax. Type
`@!image.jpg<ret>` to insert `![image.jpg](image.jpg)`. There is also
`wiki_expand_pic` command (`@!image` tag must be selected). You can use
subdirectories like in `@tag`.

### Checkboxes

You can toggle Markdown checkboxes on and off using `<ret>` key in normal mode or `wiki_toggle_checkbox` command:

```
# TODO

- [ ] foo
- [ ] bar # press <ret>
- [X] bar # result

```

## Changelog 

- 0.1:
	- initial release
- 0.2:
	- _ADD_ toggle checkbox feature
- 0.3 2018-07-15:
	- _ADD_ support for nested directories	
	- _REMOVE_ hide wiki_new_page command, use wiki instead
	- _CHANGE_ wiki command use relative paths now
- 0.4 2018-09-06:
	- _CHANGE_ update to Kakoune v2018.09.04 **breaking**
- 0.5 2018-09-11:
	- _FIX_ tag expansion in middle of the line
	- _FIX_ new line causing unwanted tag expansion
	- _FIX_ refactoring of try statements in NormalMode hooks and commands
- 0.6 2018-10-27:
    - _CHANGE_ new directory layout (**breaking**: update path in source command in `kakrc`)
    - _CHANGE_ Kakoune v2018.10.27 compatibility **breaking**
    - _CHANGE_ Changelog formatting
    - _FIX_ update README, fix spelling mistakes
- 0.7 2019-01-04:
    - _CHANGE_ update README
    - _CHANGE_ small refactoring of wiki command
    - _FIX_ following links when pwd is not in wiki_path
    - _FIX_ following links from wiki_path subdirectories
    - _FIX_ expanding tags won't create new line anymore
    - _ADD_ wiki_expand_pic and corresponding syntax `@!path/to/pic.jpg` (based on [PR #2])

[plug.kak]: https://github.com/andreyorst/plug.kak
[kakoune]: http://kakoune.org/
[PR #2]: https://github.com/TeddyDD/kakoune-wiki/pull/2
