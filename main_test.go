package main

import (
	"fmt"
	"testing"

	"github.com/jarvanstack/gitbook-summary/config"
	"github.com/jarvanstack/gitbook-summary/matcher"
)

// func TestMain1(m *testing.M) {
// 	config.Init(config.WithConfigPath("docs/gitbook-summary.yaml"))
// 	m.Run()
// 	os.Exit(0)
// }

func TestMainFunc(t *testing.T) {
	config.Init(config.WithConfigPath("/Users/hcm-b0487/workspace/work2/study1/doc2/docs/gitbook-summary.yaml"), config.WithRootDir("/Users/hcm-b0487/workspace/work2/study1/doc2/docs"))
	Summary(config.Global.RootDir, config.Global)
}

func Test_Scan(t *testing.T) {
	root, err := ScanDir(".")
	if err != nil {
		t.Error(err)
		return
	}
	printTree(root, "")
}

func TestScanAndSort(t *testing.T) {
	// config.Init("gitbook-summary.yaml")
	IgnoreMatcher = matcher.NewRegexMatcher([]string{"git"})
	root, err := ScanDir("data")
	if err != nil {
		t.Error(err)
		return
	}

	// 获取 summary 内容
	summary := GenerateSummary(root)
	fmt.Println(summary)
}

func TestFileNameToTitle(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				fileName: "test.md",
			},
			want: "Test",
		},
		{
			args: args{
				fileName: "09234a-test.md",
			},
			want: "Test",
		},
		{
			args: args{
				fileName: "09234a-test Test2.md",
			},
			want: "Test Test2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileNameToTitle(tt.args.fileName); got != tt.want {
				t.Errorf("FileNameToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSummary(t *testing.T) {
	Summary("./docs", config.Global)
}
