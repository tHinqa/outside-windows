package mmsystem

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(MMSystemApis)
	AddApis(MMSystemANSIApis)
	AddApis(MMSystemUnicodeApis)
}
