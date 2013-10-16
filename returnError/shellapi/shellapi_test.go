package shellapi

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(ShellApiApis)
	AddApis(ShellApiANSIApis)
	AddApis(ShellApiUnicodeApis)
}
