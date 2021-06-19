package packages

import (
	"unsafe"

	"github.com/barbarbar338/csgo-cheat-go/memory"
	"github.com/barbarbar338/csgo-cheat-go/offset"
	"github.com/barbarbar338/csgo-cheat-go/utils"
)

var flash = uintptr(0x0)

func NoFlash() {
	player := utils.GetPlayer()
	if player != 0 {
		var flashValue uintptr
		memory.ReadProcessMemory(utils.Process, memory.LPCVOID(player + uintptr(offset.Netvars.MFlFlashMaxAlpha)) , &flashValue, unsafe.Sizeof(flashValue))
		
		if flash != 0 {
			memory.WriteProcessMemory(utils.Process, player + uintptr(offset.Netvars.MFlFlashMaxAlpha), unsafe.Pointer(&flash), unsafe.Sizeof(flash))
		}
	}
}
