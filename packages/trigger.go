package packages

import (
	errorhelper "github.com/SakuraBurst/csgo-cheat/errorHelper"
	"github.com/SakuraBurst/csgo-cheat/memory"
	"github.com/SakuraBurst/csgo-cheat/offset"
	"github.com/SakuraBurst/csgo-cheat/utils"
)

func Trigger(procces memory.Process) {
	player := utils.GetPlayer(procces)
	client := utils.GetClient(procces)
	crosshairId, err := procces.ReadInt(player.BaseAddress + uintptr(offset.Netvars.MICrosshairID))
	errorhelper.CheckErrorAndLog(err)
	playerTeam, err := procces.ReadInt(player.BaseAddress + uintptr(offset.Netvars.MITeamNum))
	errorhelper.CheckErrorAndLog(err)
	crosshairTarget, err := procces.ReadIntPtr(client + uintptr(offset.Signatures.DwEntityList) + (uintptr(crosshairId-1))*0x10)
	errorhelper.CheckErrorAndLog(err)
	crosshairTeam, err := procces.ReadInt(crosshairTarget + uintptr(offset.Netvars.MITeamNum))
	errorhelper.CheckErrorAndLog(err)
	if crosshairId > 0 && playerTeam != crosshairTeam {
		err := procces.WriteInt(client+uintptr(offset.Signatures.DwForceAttack), 5)
		errorhelper.CheckErrorAndLog(err)
	} else {
		err := procces.WriteInt(client+uintptr(offset.Signatures.DwForceAttack), 4)
		errorhelper.CheckErrorAndLog(err)
	}
}
