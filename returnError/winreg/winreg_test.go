package winreg

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(WinRegApis)
	AddApis(WinRegANSIApis)
	AddApis(WinRegUnicodeApis)
}
