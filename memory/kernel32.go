package memory

import "C"
import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/SakuraBurst/csgo-cheat/logger"
	"github.com/Xustyx/w32"
)

var (
	kernel32                     = syscall.NewLazyDLL("kernel32.dll")
	procCloseHandle              = kernel32.NewProc("CloseHandle")
	procCreateToolhelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
	procGetLastError             = kernel32.NewProc("GetLastError")
	procGetModuleHandle          = kernel32.NewProc("GetModuleHandleW")
	procProcess32First           = kernel32.NewProc("Process32First")
	procProcess32Next            = kernel32.NewProc("Process32Next")
	procModule32First            = kernel32.NewProc("Module32First")
	procModule32Next             = kernel32.NewProc("Module32Next")
	procOpenProcess              = kernel32.NewProc("OpenProcess")
	procReadProcessMemory        = kernel32.NewProc("ReadProcessMemory")
	procWriteProcessMemory       = kernel32.NewProc("WriteProcessMemory")
	psapi                        = syscall.NewLazyDLL("psapi.dll")
	procEnumProcessModules       = psapi.NewProc("EnumProcessModules")
	modadvapi32                  = syscall.NewLazyDLL("advapi32.dll")
	procOpenProcessToken         = modadvapi32.NewProc("OpenProcessToken")
	procAdjustTokenPrivileges    = modadvapi32.NewProc("AdjustTokenPrivileges")
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

func WriteProcessMemory(hProcess HANDLE, lpBaseAddress uintptr, data []byte) error {
	var numBytesRead uintptr
	// fmt.Printf("%#x,%#x,%#x,%#x,%#x,%#x\n", procWriteProcessMemory.Addr(), 0x4, hProcess, lpBaseAddress, data,)
	// fmt.Println(hProcess, lpBaseAddress, data)
	ret, _, err := procWriteProcessMemory.Call(
		uintptr(hProcess),
		uintptr(lpBaseAddress),
		uintptr(unsafe.Pointer(&data[0])),
		uintptr(len(data)),
		uintptr(unsafe.Pointer(&numBytesRead)),
	)
	if ret == 0 {
		return err
	}
	err = nil
	return err
}

func ReadProcessMemory(hProcess HANDLE, lpBaseAddress LPCVOID, lpBuffer unsafe.Pointer, nSize uint) error {
	var numBytesRead uintptr
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess),
		uintptr(unsafe.Pointer(lpBaseAddress)),
		uintptr(lpBuffer),
		uintptr(nSize),
		uintptr(unsafe.Pointer(&numBytesRead)),
	)
	if ret == 0 {
		return err
	}
	err = nil
	return err
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
	fmt.Println(dwDesiredAccess)
	inHandle := 0
	if bInheritHandle {
		inHandle = 1
	}
	setDebugPrivilege()
	mm := 2035711
	fmt.Println(mm)
	ret, _, err := procOpenProcess.Call(
		uintptr(mm),
		uintptr(inHandle),
		uintptr(dwProcessId),
	)
	if ret == 0 {
		return HANDLE(ret), err
	}
	err = nil
	return HANDLE(ret), err
}
func setDebugPrivilege() bool {

	// Thanks to Xustyx' Goxymemory code for setting privileges and the w32 fork.
	_handle, err := syscall.GetCurrentProcess()
	pseudoHandle := HANDLE(_handle)
	if err != nil {
		return false
	}

	hToken := HANDLE(0)
	if !OpenProcessToken(pseudoHandle, w32.TOKEN_ADJUST_PRIVILEGES|w32.TOKEN_QUERY, &hToken) {
		return false
	}

	return setPrivilege(hToken, w32.SE_DEBUG_NAME, true)
}
func setPrivilege(hToken HANDLE, lpszPrivilege string, bEnablePrivilege bool) bool {
	tPrivs := w32.TOKEN_PRIVILEGES{}
	luid := w32.LUID{}

	if !w32.LookupPrivilegeValue("", lpszPrivilege, &luid) {
		return false
	}

	tPrivs.PrivilegeCount = 1
	tPrivs.Privileges[0].Luid = luid

	if bEnablePrivilege {
		tPrivs.Privileges[0].Attributes = w32.SE_PRIVILEGE_ENABLED
	} else {
		tPrivs.Privileges[0].Attributes = 0
	}
	return AdjustTokenPrivileges(hToken, 0, &tPrivs, uint32(unsafe.Sizeof(tPrivs)), nil, nil)
}

func OpenProcessToken(processHandle HANDLE, desiredAccess uint32, tokenHandle *HANDLE) bool {

	ret, _, _ := procOpenProcessToken.Call(
		uintptr(processHandle),
		uintptr(desiredAccess),
		uintptr(unsafe.Pointer(tokenHandle)))

	return ret != 0
}

type BOOL = int32

func AdjustTokenPrivileges(tokenHandle HANDLE, disableAllPrivileges BOOL, newState *w32.TOKEN_PRIVILEGES, bufferLength uint32, previousState *w32.TOKEN_PRIVILEGES, returnLength *uint32) bool {

	ret, _, _ := procAdjustTokenPrivileges.Call(
		uintptr(tokenHandle),
		uintptr(disableAllPrivileges),
		uintptr(unsafe.Pointer(newState)),
		uintptr(bufferLength),
		uintptr(unsafe.Pointer(previousState)),
		uintptr(unsafe.Pointer(returnLength)))

	return ret != 0
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
