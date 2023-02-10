
## Dev

Scan directory and get tree file, use golang

```go
type TreeNode struct {
    Name    string  // 名称
    IsDir   bool // 是否是目录
    Level   int  // 目录层级
    Children []*TreeNode // 子目录
}
func ScanDir(root string) (*TreeNode, error)  {
    // ...
}
```

Example directory structure:

```bash
.
docs
    ├── README.md
    ├── SUMMARY.md
    ├── chapter1
    │   ├── README.md'
    │   ├── chapter1-1
    │   │   ├── README.md
    │   │   ├── how to use.md
    │   │   ├── how to unuse.md
main.go
```

use

```go
ScanDir("./docs")
```

Now I need to use root node to generate SUMMARY.md file.

implement this function:

```go
func GenerateSummary(root *TreeNode) string {
    // ...
}
```

oupout:

Behind the brackets is the complete Markdown file path, which needs to be concatenated. The path is relative to the SUMMARY.md file.

```markdown
# Summary

* [README](docs/README.md)
* [chapter1](docs/chapter1/README.md)
    * [chapter1-1](docs/chapter1/chapter1-1/README.md)
        * [how to use](docs/chapter1/chapter1-1/how to use.md)
        * [how to unuse](docs/chapter1/chapter1-1/how to unuse.md)
```


it is not every dictory have README.md file

if dictory child files have README.md file

it will be

```markdown
* [chapter1](docs/chapter1/README.md)
```

if dictory child files have not README.md file

it will be

```markdown
* chapter1
```

if tree node is file use * as prefix, if tree node is dictory use - as prefix


example

```markdown
- [chapter1](docs/chapter1/README.md)
    - [chapter1-1](docs/chapter1/chapter1-1/README.md)
        * [how to use](docs/chapter1/chapter1-1/how to use.md)
        * [how to unuse](docs/chapter1/chapter1-1/how to unuse.md)
```

Title should be capitalized, and the first letter of each word should be capitalized.

example

```markdown
- [Chapter1](docs/chapter1/README.md)
    - [Chapter1-1](docs/chapter1/chapter1-1/README.md)
        * [How to use](docs/chapter1/chapter1-1/how to use.md)
        * [How to unuse](docs/chapter1/chapter1-1/how to unuse.md)
```