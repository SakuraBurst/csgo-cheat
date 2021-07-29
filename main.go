package main

import (
	"github.com/SakuraBurst/csgo-cheat/logger"
	"github.com/SakuraBurst/csgo-cheat/memory"
	"github.com/SakuraBurst/csgo-cheat/packages"
)

func main() {
	process, success := memory.GetProcessByName("csgo.exe")
	if !success {
		logger.ErrorLogger.Fatalln("Cannot find csgo.exe Run CS:GO first.")
	}
	logger.SuccessLogger.Println("csgo.exe found, PID:", process.Id)
	for {
		packages.Glow(process)
		packages.Trigger(process)
		packages.Bhop(process)
		packages.NoFlash(process)

	}
}
