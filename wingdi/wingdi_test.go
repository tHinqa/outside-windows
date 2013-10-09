package wingdi

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(WinGdiApis)
	AddApis(WinGdiANSIApis)
	AddApis(WinGdiUnicodeApis)
	//MessageBox(0, "Hello World (A)\n\nNote: The caption contains a unicode  \ncharacter that MessageBoxA cannot\ndisplay correctly", "• Go says...", 0)
}
