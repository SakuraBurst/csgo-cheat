package utils

import (
	"fmt"

	entitytypes "github.com/SakuraBurst/csgo-cheat/entityTypes"
	"github.com/SakuraBurst/csgo-cheat/memory"
	"github.com/SakuraBurst/csgo-cheat/offset"
)

func GetPlayer(proc memory.Process) entitytypes.Player {
	var player entitytypes.Player
	var err error
	player.BaseAddress, err = proc.ReadIntPtr(GetClient(proc) + uintptr(offset.Signatures.DwLocalPlayer))
	if err != nil {
		fmt.Errorf("cannot open process. reason: %s", err.Error())
	}
	return player
}

func GetClient(proc memory.Process) uintptr {
	return proc.Modules["client.dll"].BaseAddress
}

func GetPlayers() {
	// return players
}
