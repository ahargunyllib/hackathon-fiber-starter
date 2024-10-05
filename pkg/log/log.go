package log

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type LogInfo map[string]interface{}

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "15:04:05.000",
		FullTimestamp: true,
		ForceColors:  true,
	})

	timeStr := time.Now().Format("2006-01-02")

	file, err := os.OpenFile("./log/logrus."+timeStr+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.SetFormatter(&logrus.JSONFormatter{})
		log.Out = file
		return
	}

	log.Info("Failed to log to file, using default stderr")
}

func Info(fields LogInfo, info string) {
	var convFields map[string]interface{} = fields

	log.WithFields(convFields).Info(info)
}

func Warn(fields LogInfo, info string) {
	var convFields map[string]interface{} = fields

	log.WithFields(convFields).Warn(info)
}

func Fatal(fields LogInfo, info string) {
	var convFields map[string]interface{} = fields

	log.WithFields(convFields).Fatal(info)
}

func Error(fields LogInfo, info string) {
	var convFields map[string]interface{} = fields

	log.WithFields(convFields).Error(info)
}
