## gitbook-summary

[简体中文](README.zh-CN.md)

A Gitbook Summary Generator implemented by Golang

## Example

your gitbook directory structure:

```bash
├── docs
│   ├── 1-ChapterOne
│   │   ├── 1a-SectionOne.md
│   │   └── 1b-SectionTwo.md
│   ├── 2-ChapterTwo
│   │   ├── 1a-SectionOne.md
│   │   └── 1b-SectionTwo.md
│   ├── README.md
│   ├── _coverpage.md
│   ├── _media
│   │   └── icon.svg
│   ├── _sidebar.md
│   ├── gitbook-summary.yaml
│   └── index.html
```

gitbook-summary.yaml

```yaml
# 输出的文件 
outputfile: _sidebar.md
# 忽略的文件 默认是 .git 和 _
ignores:
  - _
# 排序分割符
# 例如: 10a-如何使用.md, "10a" 为排序将会忽略, "如何使用" 为标题
# 排序标题组成为 [数字][字母][排序分隔符][标题].md
isSort: true
# 排序分隔符
sortBy: "-"
# 将文件名转换为标题, 去掉分割符和排序和后缀, 例如: 10a-如何使用.md, "如何使用" 为标题, 首字母大写
isFileNameToTitle: true

```

run

```bash
$ cd docs
$ gitbook-summary 
Summary generate success, output file:  _sidebar.md 
```

output _sidebar.md

```markdown
* [README](README.md)
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
$ go install github.com/jarvanstack/gitbook-summary@latest
```

### 2. Download Binary

[Windows](release/gitbook-summary.exe) | [Linux](release/gitbook-summary) | [MacOS](release/gitbook-summary.darwin)

## license

[MIT](./LICENSE)
