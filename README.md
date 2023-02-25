## gitbook-summary

[English](./README.md) | [简体中文](README.zh-CN.md)

A Gitbook Summary Generator implemented by Golang

## Example

your gitbook directory structure:

```bash
├── docs
│   ├── 1-FirstChapter
│   │   ├── 1-FirstDocument.md
│   │   ├── 1a-SecondDocument.md
│   │   ├── 2-ThirdDocument.md
│   │   └── README.md
│   ├── 1a-SencondChapter
│   │   ├── 1-FirstDocument.md
│   │   ├── 1a-SecondDocument.md
│   │   └── 2-ThirdDocument.md
│   ├── 2-ThirdChapter
│   │   └── 1-FirstDocument.md
│   └── README.md
├── gitbook-summary.yaml
```

gitbook-summary.yaml

```yaml
# Title of summary
title: doc2
# Output file name
outputfile: _sidebar.md
# Root directory
root: "docs"
# File suffix, default .md
postfix: ".md"
# Ignore files, default ignore .git and _
ignores:
  - _
# Is sort, Will sort by name
isSort: true
# Split by "-" and sort by name, default "-"
sortBy: "-"
# Convert the file name to a title, remove the separator and sorting and suffix, for example: 10a-How to use.md, "How to use" as the title, the first letter is capitalized
isFileNameToTitle: true
```

run

```bash
$ gitbook-summary
Summary generate success, output file:  _sidebar.md 
```

output _sidebar.md

```markdown
# doc2

* [Docs](docs/README.md)
- [FirstChapter](1-FirstChapter/README.md)
    * [FirstDocument](1-FirstChapter/1-FirstDocument.md)
    * [SecondDocument](1-FirstChapter/1a-SecondDocument.md)
    * [ThirdDocument](1-FirstChapter/2-ThirdDocument.md)
- SencondChapter
    * [FirstDocument](1a-SencondChapter/1-FirstDocument.md)
    * [SecondDocument](1a-SencondChapter/1a-SecondDocument.md)
    * [ThirdDocument](1a-SencondChapter/2-ThirdDocument.md)
- ThirdChapter
    * [FirstDocument](2-ThirdChapter/1-FirstDocument.md)


```

## install

### 1. Use Golang Install

```bash
$ go install github.com/dengjiawen8955/gitbook-summary@latest
```

### 2. Download Binary

[Windows](release/gitbook-summary.exe) | [Linux](release/gitbook-summary) | [MacOS](release/gitbook-summary.darwin)

## license

[MIT](./LICENSE)
