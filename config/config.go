package config

// 代码来自:https://zhwt.github.io/yaml-to-go/

type Config struct {
	Outputfile        string   `yaml:"outputfile"`
	Ignores           []string `yaml:"ignores"`
	IsSort            bool     `yaml:"isSort"`
	SortBy            string   `yaml:"sortBy"`
	IsFileNameToTitle bool     `yaml:"isFileNameToTitle"`
}
