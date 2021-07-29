package memory

import (
	"fmt"
	"unsafe"

	errorhelper "github.com/SakuraBurst/csgo-cheat/errorHelper"
)

type (
	HANDLE  uintptr
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

type Process struct {
	Name        string
	Id          uint32
	BaseSize    uint32
	BaseAddress uintptr
	Handle      HANDLE
	Modules     map[string]Module
}

type Module struct {
	Name        string
	BaseSize    uint32
	BaseAddress uintptr
}

func (p *Process) OpenProcess() {
	proc, err := OpenProcess(PROCESS_ALL_ACCESS, false, p.Id)
	errorhelper.CheckErrorAndFatalWithPrefix(err, "Open Proccess Error: ")
	fmt.Println(proc)
	p.Handle = proc

}
func (p Process) ReadBytes(address uintptr, size uint) ([]byte, error) {
	buffer := make([]byte, size)
	err := ReadProcessMemory(p.Handle, LPCVOID(address), unsafe.Pointer(&buffer[0]), size)
	if err != nil {
		return []byte{}, err
	}
	return buffer, nil

}

func (p Process) ReadInt(address uintptr) (int, error) {
	data, err := p.ReadBytes(address, 4)
	if err != nil {
		return 0, err
	}
	return ByteToInt(data), nil

}
func (p Process) ReadIntPtr(address uintptr) (uintptr, error) {
	data, err := p.ReadBytes(address, 4)
	if err != nil {
		return 0, err
	}
	return ByteToIntPtr(data), nil

}
func (p Process) ReadInts(address uintptr, howMuchInts int) ([]int, error) {
	data, err := p.ReadBytes(address, uint(4*howMuchInts))
	if err != nil {
		return nil, err
	}
	return ByteToInts(data), nil
}

func (p Process) ReadFloat32(address uintptr) (float32, error) {
	data, err := p.ReadBytes(address, 4)
	if err != nil {
		return 0, err
	}
	fmt.Println(data)
	return ByteToFloat32(data), nil
}

func (p Process) ReadFloats32(address uintptr, howMuchFloats int) ([]float32, error) {
	data, err := p.ReadBytes(address, uint(4*howMuchFloats))
	if err != nil {
		return nil, err
	}
	return ByteToFloats32(data), nil
}
func (p Process) WriteBytes(address uintptr, data []byte) error {
	err := WriteProcessMemory(p.Handle, address, data)
	return err

}

func (p Process) WriteInt(address uintptr, integer int) error {
	err := p.WriteBytes(address, IntToByte(integer))
	if err != nil {
		return err
	}
	return nil
}
func (p Process) WriteInts(address uintptr, integers []int) error {
	err := p.WriteBytes(address, IntsToByte(integers))
	if err != nil {
		return err
	}
	return nil
}

func (p Process) WriteFloat32(address uintptr, float float32) error {
	err := p.WriteBytes(address, Float32ToByte(float))
	if err != nil {
		return err
	}
	return nil
}

func (p Process) WriteFloats32(address uintptr, floats []float32) error {
	err := p.WriteBytes(address, Floats32ToByte(floats))
	if err != nil {
		return err
	}
	return nil
}
