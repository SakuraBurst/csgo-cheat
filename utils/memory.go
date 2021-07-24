package utils

import (
	"unsafe"

	"github.com/barbarbar338/csgo-cheat-go/logger"
	"github.com/barbarbar338/csgo-cheat-go/memory"
	"github.com/barbarbar338/csgo-cheat-go/offset"
	"github.com/ttacon/chalk"
)

var (
	ClientDLL  DLL
	EngineDLL  DLL
	ClientAdrr uintptr
	EngineAdrr uintptr
	PID        uint32
	Process    memory.HANDLE
	Dummy      uint32
)

type DLL struct {
	Dll  memory.MODULEENTRY32
	Addr unsafe.Pointer
}

type Entity struct {
	Entity uintptr
}

type EntInfo struct {
	M_pEntity      *Entity
	M_SerialNumber int
	M_pPrev        *EntInfo
	M_pNext        *EntInfo
}

func init() {
	PID = GetPID()

	proc, err := memory.OpenProcess(memory.PROCESS_ALL_ACCESS, false, PID)
	if err == nil {
		logger.ErrorLogger.Fatalln("Cannot open process:", chalk.Red.Color(err.Error()))
	}

	Process = proc

	clientDLL, clientAdrr := GetDLL("client.dll")
	ClientDLL = DLL{
		Addr: clientAdrr,
		Dll:  clientDLL,
	}
	ClientAdrr = uintptr(unsafe.Pointer(clientAdrr))
	engineDLL, engineAddr := GetDLL("engine.dll")
	EngineDLL = DLL{
		Addr: engineAddr,
		Dll:  engineDLL,
	}
	EngineAdrr = uintptr(unsafe.Pointer(engineAddr))
}

func GetPID() uint32 {
	PID, success := memory.GetProcessID("csgo.exe")
	if !success {
		logger.ErrorLogger.Fatalln("Cannot find csgo.exe Run CS:GO first.")
	}

	logger.SuccessLogger.Println("csgo.exe found, PID:", PID)

	return PID
}

func GetDLL(dllName string) (memory.MODULEENTRY32, unsafe.Pointer) {
	dll, success, addr := memory.GetModule(dllName, PID)
	if !success {
		logger.ErrorLogger.Fatalln("Cannot read client.dll")
	}

	return dll, addr
}

func GetPlayer() uintptr {
	var player uintptr
	memory.ReadProcessMemory(Process, memory.LPCVOID(ClientAdrr+uintptr(offset.Signatures.DwLocalPlayer)), &player, unsafe.Sizeof(Dummy))

	return player
}

func GetPlayers() {
	// return players
}

// type GlowEntrya struct {
// 	m_nNextFreeSlot                    int32
// 	entity                             uint32
// 	m_flRed                            float32
// 	m_flGreen                          float32
// 	m_flBlue                           float32
// 	m_flAlpha                          float32
// 	m_bGlowAlphaCappedByRenderAlpha    bool
// 	m_flGlowAlphaFunctionOfMaxVelocity float32
// 	m_flGlowAlphaMax                   float32
// 	m_flGlowPulseOverdrive             float32
// 	m_bRenderWhenOccluded              bool
// 	m_bRenderWhenUnoccluded            bool
// 	m_bFullBloomRender                 bool
// 	m_nFullBloomStencilTestValue       int32
// 	m_nGlowStyle                       int32
// 	m_nSplitScreenSlot                 int32
// }
// gl := GlowEntrya{
// 	entity:                             uint32(entity),
// 	m_flRed:                            glowR,
// 	m_flGreen:                          glowG,
// 	m_flBlue:                           glowB,
// 	m_flAlpha:                          glowA,
// 	m_bGlowAlphaCappedByRenderAlpha:    true,
// 	m_flGlowAlphaFunctionOfMaxVelocity: 2,
// 	m_flGlowAlphaMax:                   2,
// 	m_flGlowPulseOverdrive:             2,
// 	m_bRenderWhenOccluded:              true,
// 	m_bRenderWhenUnoccluded:            false,
// 	m_bFullBloomRender:                 true,
// 	m_nFullBloomStencilTestValue:       2,
// 	m_nGlowStyle:                       2,
// 	m_nSplitScreenSlot:                 0,
// }
