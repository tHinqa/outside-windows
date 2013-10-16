package dde

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(DdeApis)
	AddApis(DdeANSIApis)
	AddApis(DdeUnicodeApis)
}
