package packages

import (
	errorhelper "github.com/SakuraBurst/csgo-cheat/errorHelper"
	"github.com/SakuraBurst/csgo-cheat/memory"
	"github.com/SakuraBurst/csgo-cheat/offset"
	"github.com/SakuraBurst/csgo-cheat/utils"
)

var whiteColor = []float32{255.0, 255.0, 255.0, 255.0}

func Glow(proc memory.Process) {
	client := utils.GetClient(proc)
	glowObject, err := proc.ReadIntPtr(client + uintptr(offset.Signatures.DwGlowObjectManager))
	errorhelper.CheckErrorAndLog(err)
	for i := 0; i < 64; i++ {
		entity, err := proc.ReadIntPtr(client + +uintptr(offset.Signatures.DwEntityList) + uintptr(i)*0x10)
		errorhelper.CheckErrorAndLog(err)
		if entity > 0 {
			glowIndex, err := proc.ReadIntPtr(entity + uintptr(offset.Netvars.MIGlowIndex))
			errorhelper.CheckErrorAndLog(err)
			glowIndexPointer := glowObject + glowIndex*0x38
			isDoormat, err := proc.ReadInt(entity + uintptr(offset.Signatures.MBDormant))
			errorhelper.CheckErrorAndLog(err)
			if isDoormat == 0 {
				proc.WriteFloats32(glowIndexPointer+0x8, whiteColor)
				proc.WriteBytes(glowIndexPointer+0x28, []byte{1, 0})

			}

		}
	}
}
