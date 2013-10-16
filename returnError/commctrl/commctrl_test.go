package commctrl

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(CommCtrlApis)
	AddApis(CommCtrlANSIApis)
	AddApis(CommCtrlUnicodeApis)
}
