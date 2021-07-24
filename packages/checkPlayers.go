package packages

import (
	"syscall"
	"unsafe"

	"github.com/barbarbar338/csgo-cheat-go/memory"
	"github.com/barbarbar338/csgo-cheat-go/offset"

	"github.com/barbarbar338/csgo-cheat-go/utils"
)

// type Mm struct {
// 	Red      float32
// 	Green    float32
// 	Blue     float32
// 	Alpha    float32
// 	Padding  uint8
// 	Unknown  float32
// 	Padding2 uint8
// 	RenderO  bool
// 	RenderU  bool
// 	FullB    bool
// }

var glowR float32 = 2.0
var glowG float32 = 1.0
var glowB float32 = 1.9
var glowA float32 = 2.0
var glowEn = true
var glowEj = false

var (
	kernel32               = syscall.MustLoadDLL("kernel32.dll")
	procWriteProcessMemory = kernel32.MustFindProc("WriteProcessMemory")
)

func CheckPlayers() {
	var glowObject uintptr
	memory.ReadProcessMemory(utils.Process, memory.LPCVOID(utils.ClientAdrr+uintptr(offset.Signatures.DwGlowObjectManager)), &glowObject, unsafe.Sizeof(utils.Dummy))
	// meh := Mm{
	// 	Red:      2.0,
	// 	Blue:     2.0,
	// 	Green:    2.0,
	// 	Alpha:    2.0,
	// 	Padding:  8,
	// 	Unknown:  1.0,
	// 	Padding2: 4,
	// 	RenderO:  true,
	// 	RenderU:  false,
	// 	FullB:    false,
	// }
	for {

		for i := 0; i < 32; i++ {
			var glowIndex uintptr
			var hp uintptr
			var entity uintptr
			memory.ReadProcessMemory(utils.Process, memory.LPCVOID(utils.ClientAdrr+uintptr(offset.Signatures.DwEntityList)+uintptr(i)*0x10), &entity, unsafe.Sizeof(utils.Dummy))

			if entity > 0 {
				// memory.WriteProcessMemory(utils.Process, (entity + uintptr(offset.Netvars.MBSpotted)), unsafe.Pointer(&glowEn), unsafe.Sizeof(glowEn))

				memory.ReadProcessMemory(utils.Process, memory.LPCVOID(entity+uintptr(offset.Netvars.MIGlowIndex)), &glowIndex, unsafe.Sizeof(glowIndex))
				memory.ReadProcessMemory(utils.Process, memory.LPCVOID((glowObject + glowIndex*0x38 + 0x28)), &hp, 1)
				memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0x8), unsafe.Pointer(&glowR), unsafe.Sizeof(glowR))
				memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0xC), unsafe.Pointer(&glowG), unsafe.Sizeof(glowG))
				memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0x10), unsafe.Pointer(&glowB), unsafe.Sizeof(glowB))
				memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0x14), unsafe.Pointer(&glowA), unsafe.Sizeof(glowA))
				memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0x28), unsafe.Pointer(&glowEn), unsafe.Sizeof(glowEn))
				memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0x29), unsafe.Pointer(&glowEj), unsafe.Sizeof(glowEj))

			}
		}

	}

	// time.Sleep(time.Millisecond)
}

// m, err := memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0x8), unsafe.Pointer(&glowR), unsafe.Sizeof(glowR))
// 				if err.Error() == "The operation completed successfully." {
// 					fmt.Println(m)
// 					m, err = memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0xC), unsafe.Pointer(&glowG), unsafe.Sizeof(glowG))
// 					time.Sleep(time.Microsecond * 10)
// 					if err.Error() == "The operation completed successfully." {
// 						fmt.Println(m)
// 						m, err = memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0x10), unsafe.Pointer(&glowB), unsafe.Sizeof(glowB))
// 						time.Sleep(time.Microsecond * 10)
// 						if err.Error() == "The operation completed successfully." {
// 							fmt.Println(m)
// 							m, err = memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0x14), unsafe.Pointer(&glowA), unsafe.Sizeof(glowA))
// 							time.Sleep(time.Microsecond * 10)
// 							if err.Error() == "The operation completed successfully." {
// 								fmt.Println(m)
// 								m, err = memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0x28), unsafe.Pointer(&glowEn), unsafe.Sizeof(glowEn))
// 								time.Sleep(time.Microsecond * 10)
// 								if err.Error() == "The operation completed successfully." {
// 									fmt.Println(m)
// 									memory.WriteProcessMemory(utils.Process, (glowObject + glowIndex*0x38 + 0x29), unsafe.Pointer(&glowEj), unsafe.Sizeof(glowEj))
// 									time.Sleep(time.Microsecond * 300)
// 								}
// 							}
// 						}
// 					}

// 				}
// procWriteProcessMemory.Call(
// 	uintptr(utils.Process),
// 	uintptr((glowObject + glowIndex*0x38 + 0x8)),
// 	uintptr(unsafe.Pointer(&glowR)),
// 	uintptr(unsafe.Sizeof(glowR)),
// )
// procWriteProcessMemory.Call(
// 	uintptr(utils.Process),
// 	uintptr((glowObject + glowIndex*0x38 + 0xC)),
// 	uintptr(unsafe.Pointer(&glowG)),
// 	uintptr(unsafe.Sizeof(glowG)),
// )
// procWriteProcessMemory.Call(
// 	uintptr(utils.Process),
// 	uintptr((glowObject + glowIndex*0x38 + 0x10)),
// 	uintptr(unsafe.Pointer(&glowB)),
// 	uintptr(unsafe.Sizeof(glowB)),
// )
// procWriteProcessMemory.Call(
// 	uintptr(utils.Process),
// 	uintptr((glowObject + glowIndex*0x38 + 0x14)),
// 	uintptr(unsafe.Pointer(&glowA)),
// 	uintptr(unsafe.Sizeof(glowA)),
// )
// procWriteProcessMemory.Call(
// 	uintptr(utils.Process),
// 	uintptr((glowObject + glowIndex*0x38 + 0x28)),
// 	uintptr(unsafe.Pointer(&glowEn)),
// 	uintptr(unsafe.Sizeof(glowEn)),
// )
// procWriteProcessMemory.Call(
// 	uintptr(utils.Process),
// 	uintptr((glowObject + glowIndex*0x38 + 0x29)),
// 	uintptr(unsafe.Pointer(&glowEj)),
// 	uintptr(unsafe.Sizeof(glowEj)),
// )
