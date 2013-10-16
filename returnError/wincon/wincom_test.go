package wincon

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(WinConApis)
	AddApis(WinConANSIApis)
	AddApis(WinConUnicodeApis)
}
