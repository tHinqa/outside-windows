package ole2

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(Ole2Apis)
	AddApis(ObjBaseApis)
	AddApis(OleAutoApis)
}
