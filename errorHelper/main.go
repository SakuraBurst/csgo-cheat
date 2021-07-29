package errorhelper

import "github.com/SakuraBurst/csgo-cheat/logger"

func CheckErrorAndLog(err error) {
	if err != nil {
		logger.ErrorLogger.Println(err)
	}
}

func CheckErrorAndLogWithPrefix(err error, prefix string) {
	if err != nil {
		logger.ErrorLogger.Println(err, prefix)
	}
}

func CheckErrorAndFatal(err error) {
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
	}
}

func CheckErrorAndFatalWithPrefix(err error, prefix string) {
	if err != nil {
		logger.ErrorLogger.Fatalln(err, prefix)
	}
}
