package logger

import (
	"io"
	"log"
	"os"

	"github.com/ttacon/chalk"
)

var (
	InfoLogger *log.Logger
	SuccessLogger *log.Logger
	WarningLogger *log.Logger
	ErrorLogger *log.Logger
	EventLogger *log.Logger
	DebugLogger *log.Logger
	Logger *log.Logger
)

const (
	info_icon = "ℹ"
	success_icon = "✔"
	warning_icon = "⚠"
	error_icon = "✖"
	event_icon = "☄"
	debug_icon = "☢"
	log_icon = "✎"
)

func init() {
    infoFile, err := os.OpenFile("logger/logs/INFO.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	infoWriter := io.MultiWriter(os.Stdout, infoFile)
	InfoLogger = log.New(infoWriter, chalk.Blue.Color(info_icon) + " - ", log.Ldate|log.Ltime|log.Lshortfile)

	successFile, err := os.OpenFile("logger/logs/SUCCESS.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	successWriter := io.MultiWriter(os.Stdout, successFile)
	SuccessLogger = log.New(successWriter, chalk.Green.Color(success_icon) + " - ", log.Ldate|log.Ltime|log.Lshortfile)

	warningFile, err := os.OpenFile("logger/logs/WARNING.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	warningWriter := io.MultiWriter(os.Stdout, warningFile)
	WarningLogger = log.New(warningWriter, chalk.Yellow.Color(warning_icon) + " - ", log.Ldate|log.Ltime|log.Lshortfile)

	errorFile, err := os.OpenFile("logger/logs/ERROR.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	errorWriter := io.MultiWriter(os.Stdout, errorFile)
	ErrorLogger = log.New(errorWriter, chalk.Red.Color(warning_icon) + " - ", log.Ldate|log.Ltime|log.Lshortfile)

	eventFile, err := os.OpenFile("logger/logs/EVENT.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	eventWriter := io.MultiWriter(os.Stdout, eventFile)
	EventLogger = log.New(eventWriter, chalk.Magenta.Color(event_icon) + " - ", log.Ldate|log.Ltime|log.Lshortfile)

	debugFile, err := os.OpenFile("logger/logs/DEBUG.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	debugWriter := io.MultiWriter(os.Stdout, debugFile)
	DebugLogger = log.New(debugWriter, chalk.Yellow.Color(debug_icon) + " - ", log.Ldate|log.Ltime|log.Lshortfile)

	logFile, err := os.OpenFile("logger/logs/LOG.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }
	logWriter := io.MultiWriter(os.Stdout, logFile)
	Logger = log.New(logWriter, chalk.White.Color(log_icon) + " - ", log.Ldate|log.Ltime|log.Lshortfile)
}
