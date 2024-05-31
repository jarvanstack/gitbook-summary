package config

// 代码来自:https://zhwt.github.io/yaml-to-go/

type Config struct {
	Outputfile        string   `yaml:"outputfile"`
	Ignores           []string `yaml:"ignores"`
	SortOrder         string   `yaml:"sortOrder"` // 排序规则
	IsSort            bool     `yaml:"isSort"`
	IsisSortDesc      bool     `yaml:"isisSortDesc"`
	SortBy            string   `yaml:"sortBy"`
	IsFileNameToTitle bool     `yaml:"isFileNameToTitle"`
}
