package packages

import (
	errorhelper "github.com/SakuraBurst/csgo-cheat/errorHelper"
	"github.com/SakuraBurst/csgo-cheat/memory"
	"github.com/SakuraBurst/csgo-cheat/offset"
	"github.com/SakuraBurst/csgo-cheat/utils"
)

var flash = []byte{0}

func NoFlash(proc memory.Process) {
	player := utils.GetPlayer(proc)
	if player.BaseAddress != 0 {
		flashValue, err := proc.ReadInt(player.BaseAddress + uintptr(offset.Netvars.MFlFlashMaxAlpha))
		errorhelper.CheckErrorAndLog(err)
		if flashValue != 0 {
			err = proc.WriteInt(player.BaseAddress+uintptr(offset.Netvars.MFlFlashMaxAlpha), 0)
			errorhelper.CheckErrorAndLog(err)
		}
	}
}
