package winnls

import . "github.com/tHinqa/outside"
import . "github.com/tHinqa/outside-windows/types"
import "testing"

func init() {
	AddApis(WinNlsApis)
	AddApis(WinNlsANSIApis)
	//AddApis(WinNlsUnicodeApis)
}

func callback2(s PVString, l LONG_PTR) BOOL {
	if l != 123 {
		return 0
	}
	println(*s)
	return 1
}

func callback1(s PVString) BOOL {
	println(*s)
	return 1
}

func aTest(*testing.T) {
	println("=== UIanguages ===")
	EnumUILanguages(callback2, 0, 123)
	println("=== Locales ===")
	EnumSystemLocales(callback1, 0)
	println("=== CodePages ===")
	EnumSystemCodePages(callback1, 0)
	println("=== CalendarInfo ===")
	for i := 1; i <= 0x30; i++ {
		EnumCalendarInfo(callback1, 0, 0xFFFFFFFF, CALTYPE(i))
	}
	println("=== TimeFormat ===")
	EnumTimeFormats(callback1, 0, 0)
	println("=== DateFormat ===")
	EnumDateFormats(callback1, 0, 0)
}
