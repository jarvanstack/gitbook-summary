package main

import (
	"fmt"
	"testing"

	"github.com/dengjiawen8955/gitbook-summary/matcher"
)

func Test_main(t *testing.T) {
	main()

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
