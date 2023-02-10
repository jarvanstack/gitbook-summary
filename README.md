# sm

A Gitbook Summary Generator

## Example

your gitbook directory structure:


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
main.go
```

use

```go
ScanDir("./docs")
```
