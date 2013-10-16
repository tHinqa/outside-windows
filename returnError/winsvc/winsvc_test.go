package winsvc

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(WinSvcApis)
	AddApis(WinSvcANSIApis)
	AddApis(WinSvcUnicodeApis)
}
