package wingdi

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(WinGdiApis)
	AddApis(WinGdiANSIApis)
	AddApis(WinGdiUnicodeApis)
}
