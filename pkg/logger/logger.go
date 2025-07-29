package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

// Logger 日志管理器
type Logger struct {
	accessLogger *logrus.Logger
	errorLogger  *logrus.Logger
	infoLogger   *logrus.Logger
	debugLogger  *logrus.Logger
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level    string // 日志级别: debug, info, warn, error
	Dir      string // 日志目录
	MaxSize  int    // 单个日志文件最大大小(MB)
	MaxAge   int    // 日志文件保留天数
	Compress bool   // 是否压缩旧日志文件
}

var globalLogger *Logger

// Init 初始化日志系统
func Init(config LoggerConfig) error {
	// 创建日志目录
	if err := os.MkdirAll(config.Dir, 0755); err != nil {
		return fmt.Errorf("创建日志目录失败: %v", err)
	}

	// 解析日志级别
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		level = logrus.InfoLevel
	}

	// 创建各类型日志器
	accessLogger := createLogger(filepath.Join(config.Dir, "access.log"), level)
	errorLogger := createLogger(filepath.Join(config.Dir, "error.log"), logrus.ErrorLevel)
	infoLogger := createLogger(filepath.Join(config.Dir, "info.log"), logrus.InfoLevel)
	debugLogger := createLogger(filepath.Join(config.Dir, "debug.log"), logrus.DebugLevel)

	globalLogger = &Logger{
		accessLogger: accessLogger,
		errorLogger:  errorLogger,
		infoLogger:   infoLogger,
		debugLogger:  debugLogger,
	}

	return nil
}

// createLogger 创建单个日志器
func createLogger(filename string, level logrus.Level) *logrus.Logger {
	logger := logrus.New()

	// 创建文件
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// 如果文件创建失败，使用标准输出
		logger.SetOutput(os.Stdout)
	} else {
		// 同时输出到文件和控制台
		multiWriter := io.MultiWriter(os.Stdout, file)
		logger.SetOutput(multiWriter)
	}

	// 设置日志级别
	logger.SetLevel(level)

	// 设置日志格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
		},
	})

	return logger
}

// GetLogger 获取全局日志器
func GetLogger() *Logger {
	if globalLogger == nil {
		// 如果没有初始化，使用默认配置
		_ = Init(LoggerConfig{
			Level: "info",
			Dir:   "logs",
		})
	}
	return globalLogger
}

// Access 记录访问日志
func (l *Logger) Access(message string, fields logrus.Fields) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	fields["type"] = "access"
	l.accessLogger.WithFields(fields).Info(message)
}

// Error 记录错误日志
func (l *Logger) Error(message string, fields logrus.Fields) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	fields["type"] = "error"
	l.errorLogger.WithFields(fields).Error(message)
}

// Info 记录信息日志
func (l *Logger) Info(message string, fields logrus.Fields) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	fields["type"] = "info"
	l.infoLogger.WithFields(fields).Info(message)
}

// Debug 记录调试日志
func (l *Logger) Debug(message string, fields logrus.Fields) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	fields["type"] = "debug"
	l.debugLogger.WithFields(fields).Debug(message)
}

// Warn 记录警告日志
func (l *Logger) Warn(message string, fields logrus.Fields) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	fields["type"] = "warn"
	l.infoLogger.WithFields(fields).Warn(message)
}

// 便捷函数，直接使用全局日志器
func Access(message string, fields logrus.Fields) {
	GetLogger().Access(message, fields)
}

func Error(message string, fields logrus.Fields) {
	GetLogger().Error(message, fields)
}

func Info(message string, fields logrus.Fields) {
	GetLogger().Info(message, fields)
}

func Debug(message string, fields logrus.Fields) {
	GetLogger().Debug(message, fields)
}

func Warn(message string, fields logrus.Fields) {
	GetLogger().Warn(message, fields)
}

// Errorf 格式化错误日志
func Errorf(format string, args ...interface{}) {
	Error(fmt.Sprintf(format, args...), nil)
}

// Infof 格式化信息日志
func Infof(format string, args ...interface{}) {
	Info(fmt.Sprintf(format, args...), nil)
}

// Debugf 格式化调试日志
func Debugf(format string, args ...interface{}) {
	Debug(fmt.Sprintf(format, args...), nil)
}

// Warnf 格式化警告日志
func Warnf(format string, args ...interface{}) {
	Warn(fmt.Sprintf(format, args...), nil)
}

// GetLogLevel 获取日志级别字符串对应的logrus级别
func GetLogLevel(level string) logrus.Level {
	switch strings.ToLower(level) {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}