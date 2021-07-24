package memory

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/barbarbar338/csgo-cheat-go/logger"
)

var (
	kernel32                     = syscall.MustLoadDLL("kernel32.dll")
	procCloseHandle              = kernel32.MustFindProc("CloseHandle")
	procCreateToolhelp32Snapshot = kernel32.MustFindProc("CreateToolhelp32Snapshot")
	procGetLastError             = kernel32.MustFindProc("GetLastError")
	procGetModuleHandle          = kernel32.MustFindProc("GetModuleHandleW")
	procProcess32First           = kernel32.MustFindProc("Process32First")
	procProcess32Next            = kernel32.MustFindProc("Process32Next")
	procModule32First            = kernel32.MustFindProc("Module32First")
	procModule32Next             = kernel32.MustFindProc("Module32Next")
	procOpenProcess              = kernel32.MustFindProc("OpenProcess")
	procReadProcessMemory        = kernel32.MustFindProc("ReadProcessMemory")
	procWriteProcessMemory       = kernel32.MustFindProc("WriteProcessMemory")
	psapi                        = syscall.MustLoadDLL("psapi.dll")
	procEnumProcessModules       = psapi.MustFindProc("EnumProcessModules")
)

func EnumProcessModules(hProcess HANDLE, cb uintptr, lpcbNeeded uintptr) (uintptr, []uint16, error) {

	defer func() {
		logger.DebugLogger.Println("Done")
		if x := recover(); x != nil {
			logger.ErrorLogger.Fatalln("Runtime panic: ", x)
		}
	}()

	lphModuleBuffer := make([]uint16, uintptr(lpcbNeeded))

	ret, _, err := procEnumProcessModules.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(&lphModuleBuffer[0])),
		uintptr(cb),
		uintptr(lpcbNeeded),
	)

	return ret, lphModuleBuffer, err
}

func WriteProcessMemory(hProcess HANDLE, lpBaseAddress uintptr, lpBuffer unsafe.Pointer, nSize uintptr) (uintptr, error) {
	fmt.Println(hProcess, lpBaseAddress, lpBuffer, nSize)
	ret, _, err := procWriteProcessMemory.Call(
		uintptr(hProcess),
		uintptr(lpBaseAddress),
		uintptr(lpBuffer),
		uintptr(nSize),
	)
	return ret, err
}

func ReadProcessMemory(hProcess HANDLE, lpBaseAddress LPCVOID, lpBuffer *uintptr, nSize uintptr) (uintptr, error) {
	ret, _, str := procReadProcessMemory.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpBaseAddress)),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nSize),
		0,
	)
	return ret, str
}

func CreateToolhelp32Snapshot(dwFlags uintptr, th32ProcessID uint32) HANDLE {
	ret, _, _ := procCreateToolhelp32Snapshot.Call(
		uintptr(dwFlags),
		uintptr(th32ProcessID),
	)
	return HANDLE(ret)
}

func Process32First(hSnapshot HANDLE, pe *PROCESSENTRY32) bool {
	ret, _, _ := procProcess32First.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(pe)),
	)
	return ret != 0
}

func Process32Next(hSnapshot HANDLE, pe *PROCESSENTRY32) bool {
	ret, _, _ := procProcess32Next.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(pe)),
	)
	return ret != 0
}

func Module32First(hSnapshot HANDLE, me *MODULEENTRY32) bool {
	ret, _, _ := procModule32First.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(me)),
	)
	return ret != 0
}

func Module32Next(hSnapshot HANDLE, me *MODULEENTRY32) bool {
	ret, _, _ := procModule32Next.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(me)),
	)
	return ret != 0
}

func GetModuleHandle(lpModuleName string) HMODULE {
	buff, err := syscall.UTF16PtrFromString(lpModuleName)

	if err != nil {
		panic(err)
	}

	ret, _, _ := procGetModuleHandle.Call(
		uintptr(unsafe.Pointer(buff)),
	)
	return HMODULE(ret)
}

func OpenProcess(dwDesiredAccess uint32, bInheritHandle bool, dwProcessId uint32) (HANDLE, error) {
	inHandle := 0
	if bInheritHandle {
		inHandle = 1
	}

	ret, _, err := procOpenProcess.Call(
		uintptr(dwDesiredAccess),
		uintptr(inHandle),
		uintptr(dwProcessId),
	)
	return HANDLE(ret), err
}

func CloseHandle(hObject HANDLE) bool {
	ret, _, _ := procCloseHandle.Call(
		uintptr(hObject),
	)
	return ret != 0
}

func GetLastError() uint32 {
	ret, _, _ := procGetLastError.Call()
	return uint32(ret)
}
