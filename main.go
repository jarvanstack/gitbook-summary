package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/jarvanstack/gitbook-summary/config"
	"github.com/jarvanstack/gitbook-summary/matcher"
	"github.com/spf13/pflag"
)

var (
	IgnoreMatcher matcher.Marcher
	Readme        = "README.md"
	Postfix       = ".md"
)

var (
	gcfg *config.SugaredConfig
)

func main() {
	// 加载配置文件
	config.Init()
	Summary(pflag.CommandLine.Lookup("root").Value.String(), config.Global)
}

func Summary(root string, cfg *config.SugaredConfig) {
	gcfg = cfg
	// 忽略的文件
	IgnoreMatcher = matcher.NewRegexMatcher(gcfg.Ignores)

	rootNode, err := ScanDir(root)
	if err != nil {
		panic(err)
	}

	// 获取 summary 内容
	summary := GenerateSummary(rootNode)

	// 替换路径中的空格
	summary = ReplaceSpaceInPath(summary)

	// 写入文件 outputfile
	f, err := os.Create(gcfg.Outputfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(summary)

	fmt.Printf("Summary generate success, output file:  %s \n\n", gcfg.Outputfile)
}

func ReplaceSpaceInPath(s string) string {
	var buf bytes.Buffer
	inBrackets := false

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			inBrackets = true
		} else if s[i] == ')' {
			inBrackets = false
		}

		if inBrackets && s[i] == ' ' {
			buf.WriteString("%20")
		} else {
			buf.WriteByte(s[i])
		}
	}

	return buf.String()
}

type TreeNode struct {
	Name     string      // 名称
	IsDir    bool        // 是否是目录
	Level    int         // 目录层级
	Children []*TreeNode // 子目录
}

func ScanDir(root string) (*TreeNode, error) {
	rootNode := &TreeNode{Name: root, Level: 0}

	err := scan(root, rootNode, 0)
	if err != nil {
		return nil, err
	}

	return rootNode, nil
}

func scan(path string, parent *TreeNode, level int) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	fileInfos, err := f.Readdir(-1)
	if err != nil {
		return err
	}

	for _, fileInfo := range fileInfos {
		// 忽略文件
		if IgnoreMatcher != nil && IgnoreMatcher.Match(fileInfo.Name()) {
			continue
		}

		// 不包含后缀而且不是目录
		if !strings.HasSuffix(fileInfo.Name(), Postfix) && !fileInfo.IsDir() {
			continue
		}

		child := &TreeNode{
			Name:  fileInfo.Name(),
			Level: level + 1,
		}
		parent.Children = append(parent.Children, child)

		if fileInfo.IsDir() {
			child.IsDir = true
			err := scan(filepath.Join(path, fileInfo.Name()), child, level+1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// FileNameToTitle 将文件名转换为标题
// 删除 - 分隔符和前面的排序内容，将 - 分隔符后的首字母大写
func FileNameToTitle(fileName string) string {
	if !gcfg.IsFileNameToTitle {
		return fileName
	}

	// 删除后缀
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	// 删除 - 分隔符和前面的排序内容
	fileName = strings.TrimPrefix(fileName, strings.Split(fileName, gcfg.SortBy)[0]+gcfg.SortBy)
	// 将 - 分隔符后的首字母大写
	fileName = strings.Title(fileName)
	return fileName
}

func GenerateSummary(root *TreeNode) string {
	var buffer bytes.Buffer

	sort.Slice(root.Children, func(i, j int) bool {
		return root.Children[i].Name < root.Children[j].Name
	})

	// 根目录下如果有 README.md 文件，直接使用 README.md 的标题放到最前面
	if hasReadme := hasReadme(root); hasReadme {
		buffer.WriteString(fmt.Sprintf("* [%s](/%s)\n", "README", Readme))
	}

	for _, child := range root.Children {
		generateSummaryNode(child, &buffer, 0, "")
	}

	// 最后写入换行
	buffer.WriteString("\n")

	return buffer.String()
}

func generateSummaryNode(node *TreeNode, buffer *bytes.Buffer, level int, pathPrefix string) {
	prefix := "- "
	if !node.IsDir {
		prefix = "* "
	}

	indent := strings.Repeat("    ", level)
	if node.IsDir {
		if hasReadme := hasReadme(node); hasReadme {
			// 有 README.md 文件，直接使用 README.md 的标题
			buffer.WriteString(fmt.Sprintf("%s%s[%s](/%s)\n", indent, prefix, FileNameToTitle(node.Name), filepath.Join(pathPrefix, node.Name, Readme)))
		} else {
			// 没有 README.md 文件，使用目录名作为标题
			buffer.WriteString(fmt.Sprintf("%s%s%s\n", indent, prefix, FileNameToTitle(node.Name)))
		}
	} else if node.Name != Readme {
		// 不是目录，不是 README.md 文件，使用文件名作为标题
		buffer.WriteString(fmt.Sprintf("%s%s[%s](/%s)\n", indent, prefix, FileNameToTitle(node.Name), filepath.Join(pathPrefix, node.Name)))
	}

	sort.Slice(node.Children, func(i, j int) bool {
		return node.Children[i].Name < node.Children[j].Name
	})
	for _, child := range node.Children {
		generateSummaryNode(child, buffer, level+1, filepath.Join(pathPrefix, node.Name))
	}
}

func hasReadme(node *TreeNode) bool {
	for _, child := range node.Children {
		if child.Name == Readme {
			return true
		}
	}
	return false
}

func Test_Scan(t *testing.T) {
	root, err := ScanDir(".")
	if err != nil {
		t.Error(err)
		return
	}
	printTree(root, "")
}

func printTree(node *TreeNode, prefix string) {
	fmt.Printf("%s%s\n", prefix, node.Name)
	for _, child := range node.Children {
		newPrefix := prefix
		if child.IsDir {
			newPrefix += "│  "
			if child == node.Children[len(node.Children)-1] {
				newPrefix = prefix + "    "
			}
		} else {
			newPrefix += "├─ "
		}
		printTree(child, newPrefix)
	}
}

func readIgnoreFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening ignore file:", err)
		return nil
	}
	defer file.Close()

	var ignores []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ignores = append(ignores, scanner.Text())
	}
	return ignores
}
