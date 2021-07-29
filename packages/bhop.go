package packages

import (
	errorhelper "github.com/SakuraBurst/csgo-cheat/errorHelper"
	"github.com/SakuraBurst/csgo-cheat/memory"
	"github.com/SakuraBurst/csgo-cheat/offset"
	"github.com/SakuraBurst/csgo-cheat/utils"
)

func Bhop(proc memory.Process) {
	if memory.GetAsyncKeyState(32) > 0 {
		player := utils.GetPlayer(proc)
		client := utils.GetClient(proc)
		onGround, err := proc.ReadBytes(player.BaseAddress+uintptr(offset.Netvars.MFFlags), 1)
		errorhelper.CheckErrorAndLog(err)
		if onGround[0] == 1 || onGround[0] == 7 {
			err = proc.WriteInt(client+uintptr(offset.Signatures.DwForceJump), 6)
			errorhelper.CheckErrorAndLog(err)
		}
	}

}
