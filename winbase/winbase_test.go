package winbase

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside/windows/types"
	"syscall"
	"testing"
)

import . "fmt"
import . "unsafe"

func init() {
	AddApis(WinBaseApis)
	AddApis(WinBaseANSIApis)
	//AddApis(WinBaseUnicodeApis)
	SI.Len = DWORD(Sizeof(SI))
}

var SI STARTUPINFO
var ST SYSTEMTIME

func Test(t *testing.T) {
	Buffer := new(MEMORYSTATUSEX)
	SetStructSize(Buffer)
	ret := GlobalMemoryStatusEx(Buffer)
	Printf("%v %+v\n", ret, *Buffer)

	var a, i DWORD
	var b BOOL
	Printf("%v %v %v %v\n", GetSystemTimeAdjustment(&a, &i, &b), a, i, b)

	GetStartupInfo(&SI)
	GetStartupInfo(&SI)
	Printf("%+v\n%s\n%s\n", SI, *SI.Desktop, *SI.Title)

	GetLocalTime(&ST)
	Printf("%+v\n", ST)
}

func BenchmarkSyscall(b *testing.B) {
	var d = syscall.MustLoadDLL("kernel32.dll")
	defer d.Release()
	var G = d.MustFindProc("GetStartupInfoA")
	for i := 0; i < b.N; i++ {
		syscall.Syscall(G.Addr(), 1, (uintptr)(Pointer(&SI)), 0, 0)
	}
}

func BenchmarkApi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Api(EP("GetStartupInfoA"), (uintptr)(Pointer(&SI)))
	}
}

func BenchmarkReflectStartupInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetStartupInfo(&SI)
	}
}

func BenchmarkVanillaStartupInfo(b *testing.B) {
	var d = syscall.MustLoadDLL("kernel32.dll")
	var G = d.MustFindProc("GetStartupInfoA")
	defer d.Release()
	for i := 0; i < b.N; i++ {
		syscall.Syscall(G.Addr(), 1, (uintptr)(Pointer(&SI)), 0, 0)
		s := CStrToString((uintptr)(Pointer(SI.Desktop)))
		SI.Desktop = &s
		s = CStrToString((uintptr)(Pointer(SI.Title)))
		SI.Title = &s
	}
}

func BenchmarkReflectSystemTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetLocalTime(&ST)
	}
}

func BenchmarkVanillaSystemTime(b *testing.B) {
	var d = syscall.MustLoadDLL("kernel32.dll")
	var G = d.MustFindProc("GetLocalTime")
	defer d.Release()
	for i := 0; i < b.N; i++ {
		syscall.Syscall(G.Addr(), 1, (uintptr)(Pointer(&ST)), 0, 0)
	}
}
