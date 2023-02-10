package config

// 代码来自:https://zhwt.github.io/yaml-to-go/

type Config struct {
	Title             string   `yaml:"title"`
	Outputfile        string   `yaml:"outputfile"`
	Root              string   `yaml:"root"`
	Postfix           string   `yaml:"postfix"`
	Ignores           []string `yaml:"ignores"`
	IsSort            bool     `yaml:"isSort"`
	SortBy            string   `yaml:"sortBy"`
	IsFileNameToTitle bool     `yaml:"isFileNameToTitle"`
}
