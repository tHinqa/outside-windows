package winnls

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	"testing"
)

func init() {
	outside.AddApis(WinNlsApis)
	outside.AddApis(WinNlsANSIApis)
	//outside.AddApis(WinNlsUnicodeApis)
}

var tT *testing.T

func callback2(s PVString, l LONG_PTR) BOOL {
	if l != 123 {
		return 0
	}
	tT.Log(*s)
	return 1
}

func callback1(s PVString) BOOL {
	tT.Log(*s)
	return 1
}

func Test(t *testing.T) {
	tT = t
	t.Log("=== UIanguages ===")
	EnumUILanguages(callback2, 0, 123)
	t.Log("=== Locales ===")
	EnumSystemLocales(callback1, 0)
	t.Log("=== CodePages ===")
	EnumSystemCodePages(callback1, 0)
	t.Log("=== CalendarInfo ===")
	for i := 1; i <= 0x30; i++ {
		EnumCalendarInfo(callback1, 0, 0xFFFFFFFF, CALTYPE(i))
	}
	t.Log("=== TimeFormat ===")
	EnumTimeFormats(callback1, 0, 0)
	t.Log("=== DateFormat ===")
	EnumDateFormats(callback1, 0, 0)
}
