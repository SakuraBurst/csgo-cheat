package memory

import "unsafe"

type (
	HANDLE uintptr
	LPCVOID unsafe.Pointer
	HMODULE HANDLE
)

type PROCESSENTRY32 struct {
	DwSize              uint32
	CntUsage            uint32
	Th32ProcessID       uint32
	Th32DefaultHeapID   uintptr
	Th32ModuleID        uint32
	CntThreads          uint32
	Th32ParentProcessID uint32
	PcPriClassBase      uint32
	DwFlags             uint32
	SzExeFile           [MAX_PATH]uint8
}

type MODULEENTRY32 struct {
	DwSize        uint32
	Th32ModuleID  uint32
	Th32ProcessID uint32
	GlblcntUsage  uint32
	ProccntUsage  uint32
	ModBaseAddr   *uintptr
	ModBaseSize   uint32
	HModule       HMODULE
	SzModule      [MAX_MODULE_NAME32 + 1]uint8
	SzExePath     [MAX_PATH]uint8
}
