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

	"github.com/dengjiawen8955/gitbook-summary/config"
	"github.com/dengjiawen8955/gitbook-summary/matcher"
)

var (
	ignoreMatcher matcher.Marcher
	readme        = "README.md"
)

func main() {
	// 加载配置文件
	config.Init("gitbook-summary.yaml")

	// 忽略的文件
	ignoreMatcher = matcher.NewRegexMatcher(config.Global.Ignores)

	root, err := ScanDir(config.Global.Root)
	if err != nil {
		panic(err)
	}

	// 获取 summary 内容
	summary := GenerateSummary(root)

	// 写入文件 outputfile
	f, err := os.Create(config.Global.Outputfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.WriteString(summary)

	fmt.Printf("Summary generate success, output file:  %s \n\n", config.Global.Outputfile)
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
		if ignoreMatcher != nil && ignoreMatcher.Match(fileInfo.Name()) {
			continue
		}

		// 不包含后缀而且不是目录
		if !strings.HasSuffix(fileInfo.Name(), config.Global.Postfix) && !fileInfo.IsDir() {
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
	if !config.Global.IsFileNameToTitle {
		return fileName
	}

	// 删除后缀
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))
	// 删除 - 分隔符和前面的排序内容
	fileName = strings.TrimPrefix(fileName, strings.Split(fileName, config.Global.SortBy)[0]+config.Global.SortBy)
	// 将 - 分隔符后的首字母大写
	fileName = strings.Title(strings.ReplaceAll(fileName, config.Global.SortBy, " "))
	return fileName
}

func GenerateSummary(root *TreeNode) string {
	var buffer bytes.Buffer

	if config.Global.Title != "" {
		buffer.WriteString(fmt.Sprintf("# %s\n\n", config.Global.Title))
	}

	sort.Slice(root.Children, func(i, j int) bool {
		return root.Children[i].Name < root.Children[j].Name
	})

	// 根目录下如果有 README.md 文件，直接使用 README.md 的标题放到最前面
	if hasReadme := hasReadme(root); hasReadme {
		buffer.WriteString(fmt.Sprintf("* [%s](%s)\n", FileNameToTitle(root.Name), filepath.Join(root.Name, readme)))
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
			buffer.WriteString(fmt.Sprintf("%s%s[%s](%s)\n", indent, prefix, FileNameToTitle(node.Name), filepath.Join(config.Global.Root, pathPrefix, node.Name, readme)))
		} else {
			// 没有 README.md 文件，使用目录名作为标题
			buffer.WriteString(fmt.Sprintf("%s%s%s\n", indent, prefix, FileNameToTitle(node.Name)))
		}
	} else if node.Name != readme {
		// 不是目录，不是 README.md 文件，使用文件名作为标题
		buffer.WriteString(fmt.Sprintf("%s%s[%s](%s)\n", indent, prefix, FileNameToTitle(node.Name), filepath.Join(config.Global.Root, pathPrefix, node.Name)))
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
		if child.Name == readme {
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
