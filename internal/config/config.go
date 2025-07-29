package config

import (
	"fmt"
	"os"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Redis  RedisConfig  `yaml:"redis"`
	Log    LogConfig    `yaml:"log"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level    string `yaml:"level"`    // 日志级别: debug, info, warn, error
	Dir      string `yaml:"dir"`      // 日志目录
	MaxSize  int    `yaml:"max_size"` // 单个日志文件最大大小(MB)
	MaxAge   int    `yaml:"max_age"`  // 日志文件保留天数
	Compress bool   `yaml:"compress"` // 是否压缩旧日志文件
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnvInt("SERVER_PORT", 11779),
		},
		Redis: RedisConfig{
			Host:     getEnv("REDIS_HOST", "localhost"),
			Port:     getEnvInt("REDIS_PORT", 6379),
			Password: getEnv("REDIS_PASSWORD", ""),
			DB:       getEnvInt("REDIS_DB", 0),
		},
		Log: LogConfig{
			Level:    getEnv("LOG_LEVEL", "info"),
			Dir:      getEnv("LOG_DIR", "logs"),
			MaxSize:  getEnvInt("LOG_MAX_SIZE", 100),
			MaxAge:   getEnvInt("LOG_MAX_AGE", 30),
			Compress: getEnvBool("LOG_COMPRESS", true),
		},
	}
}

func (c *Config) GetServerAddr() string {
	return fmt.Sprintf("%s:%d", c.Server.Host, c.Server.Port)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue := parseInt(value); intValue != 0 {
			return intValue
		}
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return value == "true" || value == "1" || value == "yes"
	}
	return defaultValue
}

func parseInt(s string) int {
	var result int
	for _, char := range s {
		if char >= '0' && char <= '9' {
			result = result*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return result
}
