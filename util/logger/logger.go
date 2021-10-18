package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log Logger

type LogConfig struct {
	LOG_MAXSIZE   int
	LOG_MAXBACKUP int
	LOG_MAXAGE    int
	LOG_COMPRESS  bool
}

type Logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

func setLogger(newLogger Logger) {
	Log = newLogger
}

type writerHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
	Formatter logrus.Formatter
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	// line, err := entry.String()
	// if err != nil {
	// 	return err
	// }
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to format log entry: %v", err)
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func LoadLogger(path string, logConfig LogConfig) error {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceColors:     true,
	})

	log.SetLevel(logrus.DebugLevel)
	log.SetOutput(os.Stderr)
	log.SetReportCaller(true)

	lumberjackLoggerRoot := &lumberjack.Logger{
		// Log file abbsolute path, os agnostic
		Filename:   filepath.ToSlash("/usr/local/go/logs/" + path + "/root.log"),
		MaxSize:    500, // MB
		MaxBackups: 0,
		MaxAge:     0,    // days
		Compress:   true, // disabled by default
	}
	writerRoot := io.Writer(lumberjackLoggerRoot)

	log.AddHook(&writerHook{ // Send logs with level higher than warning to stderr
		Writer: writerRoot,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
		Formatter: &logrus.JSONFormatter{
			PrettyPrint: true,
		},
	})

	lumberjackLoggerDebug := &lumberjack.Logger{
		// Log file abbsolute path, os agnostic
		Filename:   filepath.ToSlash("/usr/local/go/logs/" + path + "/" + path + ".log"),
		MaxSize:    logConfig.LOG_MAXSIZE, // MB
		MaxBackups: logConfig.LOG_MAXBACKUP,
		MaxAge:     logConfig.LOG_MAXAGE,   // days
		Compress:   logConfig.LOG_COMPRESS, // disabled by default
	}
	writerDebug := io.Writer(lumberjackLoggerDebug)

	log.AddHook(&writerHook{ // Send info and debug logs to stdout
		Writer: writerDebug,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
		Formatter: &logrus.JSONFormatter{
			PrettyPrint: true,
		},
	})

	setLogger(log)

	return nil
}
