package packages

import (
	"time"
	"unsafe"

	"github.com/barbarbar338/csgo-cheat-go/memory"
	"github.com/barbarbar338/csgo-cheat-go/offset"
	"github.com/barbarbar338/csgo-cheat-go/utils"
)

var jump = uintptr(0x6)

func Bhop() {
	if memory.GetAsyncKeyState(32) > 0 {
		player := utils.GetPlayer()

		var onGround uintptr
		memory.ReadProcessMemory(utils.Process, memory.LPCVOID(player+uintptr(offset.Netvars.MFFlags)), &onGround, 1)

		if onGround == 1 || onGround == 7 {
			memory.WriteProcessMemory(utils.Process, uintptr(utils.ClientDLL.Addr)+uintptr(offset.Signatures.DwForceJump), unsafe.Pointer(&jump), unsafe.Sizeof(jump))
		}
	}

	time.Sleep(time.Millisecond * 3)
}
