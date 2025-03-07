package lab

import log "github.com/sirupsen/logrus"

func Lab5() {
	log.SetLevel(log.DebugLevel)
	logMessage("debug", "This is a debug message")
	logMessage("info", "This is an info message")
	logMessage("warning", "This is a warning message")
	logMessage("error", "This is an error message")
	logMessage("invalid", "This is an invalid message")
}

func logMessage(s string, s2 string) {
	switch s {
	case "debug":
		log.Debug(s2)
	case "info":
		log.Info(s2)
	case "warning":
		log.Warning(s2)
	case "error":
		log.Error(s2)
	case "fatal":
		log.Fatal(s2)
	case "panic":
		log.Panic(s2)
	default:
		log.Error("This is an invalid: " + s)
	}
}
