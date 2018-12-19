package config

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/gomail.v2"
)

var config *Config

func Conf() *Config {
	if config == nil {
		config = new(Config)
	}
	return config
}

type Config struct {
	PrivateKeyPath string `json:"private_key_path"`
	PublicKeyPath  string `json:"public_key_path"`

	ApiPort string `json:"api_port"` //httpServer
	WebPort string `json:"web_port"` //httpServer

	StaticPath   string     `json:"static_path"`
	DBConfig     DBConfig   `json:"db_config"`     //数据库配置
	RouterPrefix string     `json:"router_prefix"` //api前缀
	AuthPrefix   string     `json:"auth_prefix"`   //白名单
	MailConfig   MailConfig `json:"mail_config"`   //邮箱
	AppConfig    AppConfig  `json:"app_config"`
}

type AppConfig struct {
	ApiHost    string `json:"api_host"`    //api请求host
	DomainName string `json:"domain_name"` //域名
	Version    string `json:"version"`     //版本
}

/*
数据库配置结构体
*/
type DBConfig struct {
	Url       string `json:"url"`     //连接地址
	DBName    string `json:"db_name"` //用户名
	ForceSync bool   `json:"force_sync"`
}

type MailConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

/*
读取配置文件
*/
func ReadConfig(path string) error {
	config = new(Config)
	err := config.Parse(path)
	return err
}

/*
解析配置文件
*/
func (c *Config) Parse(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &c)
	if err != nil {
		return err
	}
	return nil
}

func (c *MailConfig) Dialer() (*gomail.Dialer, error) {
	d := gomail.NewDialer(c.Host, c.Port, c.Username, c.Password)
	return d, nil
}
