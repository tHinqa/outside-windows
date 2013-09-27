package winuser

import . "github.com/tHinqa/outside"
import . "github.com/tHinqa/outside/windows/types"
import "testing"
import . "fmt"

// import "unsafe"

func init() {
	AddApis(WinUserApis)
	// AddApis(WinUserANSIApis)
	AddApis(WinUserUnicodeApis)
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

func Test1(t *testing.T) {
	if EnumWindows(callback, 0) == 0 {
		t.Fail()
	}
	Println(W, "Windows")
}
func Test2(t *testing.T) {
	h := GetProcessWindowStation()
	if EnumDesktops(h, wscallback, 123) == 0 {
		t.Fail()
	}
}
func Test3(t *testing.T) {
	if EnumWindowStations(wscallback, 123) == 0 {
		t.Fail()
	}
}
