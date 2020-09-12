package internal

import (
	"os"
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
)

func InitialiseLogger() {
	var logLevel = logrus.InfoLevel

	logFileDirectory := os.Getenv("LOGS_PATH")
	filename := logFileDirectory + "/console.log"
	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   filename,
		MaxSize:    50,
		MaxBackups: 3,
		MaxAge:     28,
		Level:      logLevel,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: time.RFC822,
		},
	})

	if err != nil {
		logrus.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	logrus.SetLevel(logLevel)
	logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	logrus.AddHook(rotateFileHook)
}
