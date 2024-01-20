package logger

import (
	"log"
	"os"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	i, iErr := os.OpenFile("../../pkg/logger/logs/infoLogs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	w, wErr := os.OpenFile("../../pkg/logger/logs/warningLogs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	e, eErr := os.OpenFile("../../pkg/logger/logs/errorLogs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if iErr != nil || wErr != nil || eErr != nil {
		log.Println(iErr, wErr, eErr)
		log.Fatal("Unable to open log file")
	}

	Info = log.New(i, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(w, "WARNING: ", log.Ldate|log.Ltime)
	Error = log.New(e, "ERROR: ", log.Ldate|log.Ltime)
}
