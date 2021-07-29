package memory

import (
	"bytes"
	"errors"
	"unsafe"
)

func GetProcessByName(process string) (Process, bool) {
	var snap HANDLE
	var pe32 PROCESSENTRY32
	snap = CreateToolhelp32Snapshot(TH32CS_SNAPALL, 0)
	defer CloseHandle(snap)
	pe32.DwSize = uint32(unsafe.Sizeof(pe32))
	exit := Process32First(snap, &pe32)
	if !exit {
		return Process{}, false
	} else {
		for Process32Next(snap, &pe32) {
			proc, _ := GetProcess(pe32.Th32ProcessID)
			if proc.Name == process {
				proc.OpenProcess()

				return proc, true
			}
		}
		return Process{}, false
	}

}

func GetProcess(PID uint32) (Process, error) {
	var pe32 MODULEENTRY32
	pe32.DwSize = uint32(unsafe.Sizeof(pe32))
	snap := CreateToolhelp32Snapshot(TH32CS_SNAPMODULE|TH32CS_SNAPMODULE32, PID)
	defer CloseHandle(snap)
	if Module32First(snap, &pe32) {
		proc := Process{
			Name:        parseint8(pe32.SzModule[:]),
			Id:          PID,
			BaseSize:    pe32.ModBaseSize,
			BaseAddress: uintptr(unsafe.Pointer(pe32.ModBaseAddr)),
			Modules:     map[string]Module{},
		}
		for Module32Next(snap, &pe32) {
			proc.Modules[parseint8(pe32.SzModule[:])] = Module{
				Name:        parseint8(pe32.SzModule[:]),
				BaseSize:    pe32.DwSize,
				BaseAddress: uintptr(unsafe.Pointer(pe32.ModBaseAddr)),
			}
		}

		return proc, nil
	} else {
		return Process{}, errors.New("cannot open module snap")
	}
}

func GetModule(module string, PID uint32) (MODULEENTRY32, bool, unsafe.Pointer) {
	var me32 MODULEENTRY32

	snap := CreateToolhelp32Snapshot(TH32CS_SNAPMODULE|TH32CS_SNAPMODULE32, PID)

	me32.DwSize = uint32(unsafe.Sizeof(me32))

	exit := Module32First(snap, &me32)
	if !exit {
		CloseHandle(snap)
		return me32, false, unsafe.Pointer(me32.ModBaseAddr)
	} else {
		for i := true; i; i = Module32Next(snap, &me32) {
			parsed := parseint8(me32.SzModule[:])
			if parsed == module {
				return me32, true, unsafe.Pointer(me32.ModBaseAddr)
			}
		}
	}
	return me32, false, unsafe.Pointer(me32.ModBaseAddr)
}

func parseint8(arr []uint8) string {
	n := bytes.Index(arr, []uint8{0})
	return string(arr[:n])
}

func OffsetAddr(hProcess HANDLE, baseAddr uintptr, offAddrs []uintptr) uintptr {
	var finalAddr uintptr

	return finalAddr
}
