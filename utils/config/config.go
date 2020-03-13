package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// LoadGlobalConfig 加载全局配置
func LoadGlobalConfig(path string) (c *Config, err error) {
	c, err = ParseConfig(path)
	if err != nil {
		return
	}
	return
}

// ParseConfig 解析配置文件
func ParseConfig(path string) (*Config, error) {
	var c Config
	_, err := toml.DecodeFile(path, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

// Config 配置参数
type Config struct {
	Project    Project    `toml:"project"`
	Log        Log        `toml:"log"`
	Gorm       Gorm       `toml:"gorm"`
	MySQL      MySQL      `toml:"mysql"`
	HTTP       HTTP       `toml:"http"`
	Redis      Redis      `toml:"redis"`
	FileUpload FileUpload `toml:"fileUpload"`
}

// Project 项目
type Project struct {
	Name string `toml:"name"`
}

// HTTP http配置参数
type HTTP struct {
	Host            string `toml:"host"`
	Port            int    `toml:"port"`
	ShutdownTimeout int    `toml:"shutdown_timeout"`
}

// Log 日志配置参数
type Log struct {
	Level    int    `toml:"level"`
	Path     string `toml:"path"`
	MaxDays  int    `toml:"maxdays"`
	Separate string `toml:"separate"`
}

// Gorm gorm配置参数
type Gorm struct {
	Debug        bool   `toml:"debug"`
	DBType       string `toml:"db_type"`
	MaxLifetime  int    `toml:"max_lifetime"`
	MaxOpenConns int    `toml:"max_open_conns"`
	MaxIdleConns int    `toml:"max_idle_conns"`
	TablePrefix  string `toml:"table_prefix"`
}

// MySQL mysql配置参数
type MySQL struct {
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	User       string `toml:"user"`
	Password   string `toml:"password"`
	DBName     string `toml:"db_name"`
	Parameters string `toml:"parameters"`
}

// Redis redis配置参数
type Redis struct {
	Addr         string `toml:"addr"`
	Password     string `toml:"password"`
	Database     int    `toml:"database"`
	MaxOpenConns int    `toml:"maxOpenConns"`
	MaxIdleConns int    `toml:"maxIdleConns"`
}

// FileUpload 文件上传配置参数
type FileUpload struct {
	BasePath    string `toml:"basePath"`
	Path        string `toml:"path"`
	DoMain      string `toml:"doMain"`
	MaxFileSize int64  `toml:"maxFileSize"`
	ExtFilter   string `toml:"extFilter"`
}

// DSN 数据库连接串
func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}
