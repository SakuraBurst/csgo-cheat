package utils

import (
	"unsafe"

	"github.com/barbarbar338/csgo-cheat-go/logger"
	"github.com/barbarbar338/csgo-cheat-go/memory"
	"github.com/barbarbar338/csgo-cheat-go/offset"
	"github.com/ttacon/chalk"
)

var (
	ClientDLL DLL
	EngineDLL DLL
	PID uint32
	Process memory.HANDLE
	Dummy uint32
)

type DLL struct {
	Dll memory.MODULEENTRY32
	Addr unsafe.Pointer
}

func init() {
	PID = GetPID()

	proc, err := memory.OpenProcess(memory.PROCESS_ALL_ACCESS, false, PID)
	if err == nil {
		logger.ErrorLogger.Fatalln("Cannot open process:", chalk.Red.Color(err.Error()))
	}

	Process = proc

	clientDLL, clientAdrr := GetDLL("client.dll")
	ClientDLL = DLL {
		Addr: clientAdrr,
		Dll: clientDLL,
	}
	
	engineDLL, engineAddr := GetDLL("engine.dll")
	EngineDLL = DLL {
		Addr: engineAddr,
		Dll: engineDLL,
	}
}

func GetPID() (uint32) {
	PID, success := memory.GetProcessID("csgo.exe")
	if !success {
		logger.ErrorLogger.Fatalln("Cannot find csgo.exe Run CS:GO first.")
	}

	logger.SuccessLogger.Println("csgo.exe found, PID:", PID)

	return PID
}

func GetDLL(dllName string) (memory.MODULEENTRY32, unsafe.Pointer){
	dll, success, addr := memory.GetModule(dllName, PID)
	if !success {
		logger.ErrorLogger.Fatalln("Cannot read client.dll")
	}

	return dll, addr
}

func GetPlayer() (uintptr) {
	base := uintptr(unsafe.Pointer(ClientDLL.Addr))

	var player uintptr
	memory.ReadProcessMemory(Process, memory.LPCVOID(base + uintptr(offset.Signatures.DwLocalPlayer)), &player, unsafe.Sizeof(Dummy))
	
	return player
}
