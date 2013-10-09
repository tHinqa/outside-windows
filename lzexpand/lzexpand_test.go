package lzexpand

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(LZExpandApis)
	AddApis(LZExpandANSIApis)
	AddApis(LZExpandUnicodeApis)
}
