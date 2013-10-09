package winsock2

import . "github.com/tHinqa/outside"
import "testing"

func Test(*testing.T) {
	AddApis(WinSock2Apis)
	AddApis(WinSock2ANSIApis)
	AddApis(WinSock2UnicodeApis)
}
