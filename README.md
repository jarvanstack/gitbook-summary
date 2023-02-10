## gitbook-summary


[English](./README.md) | 中文

Golang 实现的 Gitbook 摘要生成器

## 快速开始

目录结构

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

gitbook-summary.yaml 配置文件

```yaml
# 标题
title: doc2
# 输出文件名
outputfile: _sidebar.md
# 扫描根目录
root: "docs"
# 匹配文件后缀
postfix: ".md"
# 忽略的文件或者目录, 默认是 .git and _
ignores:
  - _
# 是否排序, 建议开启通过文件名排序
isSort: true
# 排序分隔符, 搭配 isFileNameToTitle 使用, 默认 "-"
sortBy: "-"
# 将文档名转换为标题，去除分隔符和排序及后缀，例如：10a-How to use.md，“How to use”为标题，首字母大写
isFileNameToTitle: true
```

启动

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

## 安装

### 方式1: go 安装

```bash
$ go install -u github.com/zhengxiaoyao0716/gitbook-summary@latest
```

### 方式2: 下载二进制文件

[Windows](bin/gitbook-summary.exe) | [Linux](gitbook-summary) | [MacOS](gitbook-summary.darwin)

## license

[MIT](./LICENSE)