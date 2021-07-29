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

// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/SakuraBurst/csgo-cheat/offset"
// 	"github.com/SakuraBurst/csgo-cheat/readandwritememory"
// )

// const (
// 	dwLocalPlayer       = 14197468
// 	dwEntityList        = 81408476
// 	dwGlowObjectManager = 86947408
// 	iGlowIndex          = 42040
// 	iTeamNum            = 244
// 	bDormant            = 237
// )

// var (
// 	tColor = []float32{255.0, 0.0, 0.0, 1.3}
// 	cColor = []float32{0.0, 0.0, 255.0, 1.3}
// )

// func main() {
// 	process, err := readandwritememory.ProcessByName("csgo.exe")
// 	if err != nil {
// 		log.Panicf("cs go running? Error: %s", err.Error())
// 	}
// 	log.Printf("Base: 0x%06X", process.ModBaseAddr)
// 	// log.Printf("Modules load: %s", reflect.ValueOf(process.Modules).MapKeys())

// 	gameModule := process.Modules["client.dll"].ModBaseAddr
// 	log.Printf("Client Module Base: 0x%06X", gameModule)

// 	localPlayer, err := process.ReadIntPtr(gameModule + dwLocalPlayer)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	localPlayercroshair, err := process.ReadIntPtr(dwLocalPlayer + uintptr(offset.Netvars.MICrosshairID))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	log.Println(localPlayer, localPlayercroshair, offset.Netvars.MICrosshairID)

// 	// Mainloop
// 	for {
// 		for i := 0; i < 64; i++ {
// 			entAddr, _ := process.ReadIntPtr(gameModule + dwEntityList + uintptr(i)*0x10)
// 			if entAddr != 0 && entAddr != localPlayer {
// 				teamNum, _ := process.ReadInt(entAddr + iTeamNum)
// 				isDormant, _ := process.ReadInt(entAddr + bDormant)

// 				if isDormant == 0 {
// 					GlowObject, _ := process.ReadInt(gameModule + dwGlowObjectManager)
// 					GlowIndex, _ := process.ReadInt(entAddr + iGlowIndex)
// 					GlowIndexPointer := uintptr(GlowObject + GlowIndex*0x38)

// 					if teamNum == 2 {
// 						_ = process.WriteFloats(GlowIndexPointer+8, tColor)
// 					} else {
// 						_ = process.WriteFloats(GlowIndexPointer+8, cColor)
// 					}

// 					_ = process.WriteBytes(GlowIndexPointer+0x28, []byte{1, 0})
// 				}
// 			}
// 		}

// 		time.Sleep(1 * time.Millisecond)
// 	}
// }
