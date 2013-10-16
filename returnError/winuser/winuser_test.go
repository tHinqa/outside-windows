package winuser

import (
	. "fmt"
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	"testing"
	"unsafe"
)

// import "unsafe"

func init() {
	outside.AddApis(WinUserApis)
	// outside.AddApis(WinUserANSIApis)
	outside.AddApis(WinUserUnicodeApis)
}

//TODO(t): Handle output string buffers

func callback(h HWND, l LPARAM) BOOL {
	var t [256]WChar
	GetClassName(h, &t[0], 255)
	// c := CStrToString((uintptr)(unsafe.Pointer(&t[0])))
	// c := UniStrToString((uintptr)(unsafe.Pointer(&t[0])))
	// s := GetWindowTextLength(h)
	// for i := LPARAM(0); i < l; i++ {
	// 	Print("\t")
	// }

	// Print(c)
	// if s != 0 {
	// 	GetWindowText(h, &t[0], 255)
	// 	// Print(" ", CStrToString((uintptr)(unsafe.Pointer(&t[0]))))
	// 	// Print(" ", UniStrToString((uintptr)(unsafe.Pointer(&t[0]))))
	// }
	// Println()
	EnumChildWindows(h, callback, l+1)
	W++
	return 1
}

var W uint32

func wscallback(c PVString, l LPARAM) BOOL {
	if l != 123 {
		return 0
	}
	Println(*c)
	return 1
}

func TestEnumWindows(t *testing.T) {
	ew, e := EnumWindows(callback, 0)
	if ew == 0 {
		t.Fatal(e)
	}
	Println(W, "Windows")
}
func TestEnumDesktops(t *testing.T) {
	h, _ := GetProcessWindowStation()
	ed, e := EnumDesktops(h, wscallback, 123)
	if ed == 0 {
		t.Fatal(e)
	}
}
func TestEnumWindowStations(t *testing.T) {
	ew, e := EnumWindowStations(wscallback, 123)
	if ew == 0 {
		t.Fatal(e)
	}
}

func TestWsprintf(t *testing.T) {
	var o [1000]WChar
	Wsprintf(&o[0], "%d %d %d", 123, 456, 789)
	t.Log(outside.UniStrToString((uintptr)(unsafe.Pointer(&o[0]))))
}
