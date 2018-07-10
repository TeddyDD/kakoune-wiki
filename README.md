# kakoune-wiki

![icon](kakoune-wiki.png)

Personal wiki plugin for [**Kakoune**](http://kakoune.org/)

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
- toggle Markdown check-boxes with `<ret>` key in normal mode
- wiki is just bunch of Markdown files, you can process/edit them further
with tools like [pandoc](https://pandoc.org/),
[MdWiki](http://dynalon.github.io/mdwiki/),
[markor](https://github.com/gsantner/markor) or any text editor. No vendor lock-ins
- minimal and simple — only essential features, script around 60 LOC

## Installation

You can either:

- load `wiki.kak` from your `kakrc`: `source path/to/wiki.kak`
- put `wiki.kak` in your autoloads directory `~/.config/kak/autoload/`

Then you have to choose directory for your wiki. Call following command from
your kakrc:

```
wiki_setup `~/my/wiki/directory`
```

## Usage

You probably want some kind of index page. Create it manually using wiki_new_page command:

```
:wiki_new_page index
```

Now you can edit index page using `wiki` command. You can press `TAB` for completion:

```
:wiki index<TAB>
```

Let's add some content. You can create new page using `@pagename` syntax.
Just type @ character and desired filename, then press enter. If file doesn't
exist it will be created in wiki directory. `@filename` will be expanded to
`[filename](filename.md)`

```
# Wiki

## Cooking
- @cookies
```

To edit new page about :cookie: put selection anywhere in the link and press enter.
