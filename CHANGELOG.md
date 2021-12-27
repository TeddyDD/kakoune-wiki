# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Types of changes:

- *Added* for new features.
- *Changed* for changes in existing functionality.
- *Deprecated* for soon-to-be removed features.
- *Removed* for now removed features.
- *Fixed* for any bug fixes.
- *Security* in case of vulnerabilities.

## [Unreleased]

### Added
- link conversion between Markdown and Media Wiki links
- autocompletion for links

### Changed
- Changelog in separate file
- Kakoune v2021.11.08 compatibility
- Plugin uses CLI helper written in Go

### Removed
- @tag syntax for links

---

## Legacy (pre 1.0.0)

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
  - _CHANGE_ new directory layout (**breaking**: update path in source
  command in `kakrc`)
  - _CHANGE_ Kakoune v2018.10.27 compatibility **breaking**
  - _CHANGE_ Changelog formatting
  - _FIX_ update README, fix spelling mistakes
- 0.7 2019-01-04:
  - _CHANGE_ update README
  - _CHANGE_ small refactoring of wiki command
  - _FIX_ following links when pwd is not in wiki_path
  - _FIX_ following links from wiki_path subdirectories
  - _FIX_ expanding tags won't create new line anymore
  - _ADD_ wiki_expand_pic and corresponding syntax `@!path/to/pic.jpg`
  (based on [PR #2])
- 0.8 2020-02-03:
  - __CHANGE__ Kakoune v2020.01.16 compatibility **breaking**
  - __CHANGE__ `wiki_setup` rename to `wiki-setup`

[PR #2]: https://github.com/TeddyDD/kakoune-wiki/pull/2
