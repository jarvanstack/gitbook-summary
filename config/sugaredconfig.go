package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var version = "v0.0.1"

// SugaredConfig 将配置文件的参数解析,比如解析时间为 time.Ticker
type SugaredConfig struct {
	*Config
	*OptionCfg
}

var Global *SugaredConfig

type OptionCfg struct {
	ConfigPath string
	RootDir    string
}

type Option func(*OptionCfg)

func WithConfigPath(path string) Option {
	return func(o *OptionCfg) {
		o.ConfigPath = path
	}
}

func WithRootDir(dir string) Option {
	return func(o *OptionCfg) {
		o.RootDir = dir
	}
}

func Init(opts ...Option) *SugaredConfig {

	var o OptionCfg
	for _, opt := range opts {
		opt(&o)
	}

	configPath := o.ConfigPath
	if configPath == "" {
		configPath = "gitbook-summary.yaml"
	}

	// 初始化配置文件
	pflag.StringP("config", "c", configPath, "config file")
	pflag.BoolP("version", "v", false, "show version")
	pflag.StringP("root", "r", ".", "root dir")
	pflag.Parse()

	// 初始化 root config
	if o.RootDir == "" {
		o.RootDir = pflag.CommandLine.Lookup("root").Value.String()
	}

	// Print version
	if b, _ := pflag.CommandLine.GetBool("version"); b {
		fmt.Println("gitbook-summary version: ", version)
		os.Exit(0)
	}

	viper.SetConfigType("yaml")
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		panic(err)
	}
	conf := viper.GetString("config")
	viper.SetConfigFile(conf)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("load config %s fail: %v", conf, err))
	}

	// 解析初始配置
	baseConf := &Config{}
	if err := viper.Unmarshal(baseConf); err != nil {
		panic(err)
	}

	// 构造 SugaredConfig
	Global = &SugaredConfig{
		Config:    baseConf,
		OptionCfg: &o,
	}

	// 默认 Ignores
	Global.Ignores = append(Global.Ignores, "_")
	Global.Ignores = append(Global.Ignores, "\\.git")

	// 默认输出文件名 _sidebar.md
	if len(Global.Outputfile) == 0 {
		Global.Outputfile = "_sidebar.md"
	}

	// 默认 排序分隔符 -
	if len(Global.SortBy) == 0 {
		Global.SortBy = "-"
	}

	return Global
}
