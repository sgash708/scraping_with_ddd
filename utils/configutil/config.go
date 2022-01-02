package configutil

import (
	"gopkg.in/ini.v1"
)

// ScrapeList クレイピングに使う要素
type ScrapeList struct {
	LAWURL  string
	OFFURL  string
	CWToken string
	CWURL   string
	CWMes   string
	DBUser  string
	DBPass  string
	DBName  string
	DBHost  string
	DBPort  int
	SSHHost string
	SSHPort int
	SSHUser string
	SSHPass string
}

type Config struct {
	filePath string
}

// IConfig Configインターフェース
type IConfig interface {
	Setup() (*ScrapeList, error)
}

// NewChatWork 初期化
func NewConfig(filePath string) IConfig {
	return &Config{filePath: filePath}
}

// Setup 初期化
func (c *Config) Setup() (*ScrapeList, error) {
	config, err := ini.Load(c.filePath)
	if err != nil {
		return nil, err
	}
	Scraping := ScrapeList{
		LAWURL:  config.Section("web").Key("LAWURL").MustString(""),
		OFFURL:  config.Section("web").Key("OFFURL").MustString(""),
		CWToken: config.Section("chatwork").Key("TOKEN").MustString(""),
		CWURL:   config.Section("chatwork").Key("URL").MustString(""),
		CWMes:   config.Section("chatwork").Key("Message").MustString(""),
		DBUser:  config.Section("pro").Key("USER").MustString(""),
		DBPass:  config.Section("pro").Key("PASS").MustString(""),
		DBName:  config.Section("pro").Key("NAME").MustString(""),
		DBHost:  config.Section("pro").Key("HOST").MustString(""),
		DBPort:  config.Section("pro").Key("PORT").MustInt(),
		SSHHost: config.Section("ssh").Key("HOST").MustString(""),
		SSHPort: config.Section("ssh").Key("PORT").MustInt(),
		SSHUser: config.Section("ssh").Key("USER").MustString(""),
		SSHPass: config.Section("ssh").Key("PASS").MustString(""),
	}

	return &Scraping, nil
}
