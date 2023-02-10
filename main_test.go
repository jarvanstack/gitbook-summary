package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"
)

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

func Test_Matcher(t *testing.T) {
	ignoredFiles := readIgnoreFile(".gitignore")
	ignoredFiles = append(ignoredFiles, ".git")
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		for _, ignore := range ignoredFiles {
			matched, _ := regexp.MatchString(ignore, path)
			if matched {
				return nil
			}
		}

		fmt.Println(path)
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
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
