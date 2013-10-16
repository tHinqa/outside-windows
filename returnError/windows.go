// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package windows.
package windows

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/returnError/commctrl"
	. "github.com/tHinqa/outside-windows/returnError/dde"
	. "github.com/tHinqa/outside-windows/returnError/lzexpand"
	. "github.com/tHinqa/outside-windows/returnError/mmsystem"
	. "github.com/tHinqa/outside-windows/returnError/nb30"
	. "github.com/tHinqa/outside-windows/returnError/ole2"
	. "github.com/tHinqa/outside-windows/returnError/shellapi"
	. "github.com/tHinqa/outside-windows/returnError/winbase"
	. "github.com/tHinqa/outside-windows/returnError/wincon"
	. "github.com/tHinqa/outside-windows/returnError/wingdi"
	. "github.com/tHinqa/outside-windows/returnError/winnetwk"
	. "github.com/tHinqa/outside-windows/returnError/winnls"
	. "github.com/tHinqa/outside-windows/returnError/winreg"
	// . "github.com/tHinqa/outside-windows/winsock"
	. "github.com/tHinqa/outside-windows/returnError/winsock2"
	. "github.com/tHinqa/outside-windows/returnError/winsvc"
	. "github.com/tHinqa/outside-windows/returnError/winuser"
)

//TODO(t): Check for underscore and lowercase prefixes
// New struct func filling in structsize
// use s rather than c for size fields
// Check all xxx xxxA dups

func AllANSIApis() {
	AddApis(WinBaseApis)
	AddApis(WinBaseANSIApis)
	AddApis(WinUserApis)
	AddApis(WinUserANSIApis)
	AddApis(WinGdiApis)
	AddApis(WinGdiANSIApis)
	AddApis(WinConApis)
	AddApis(WinConANSIApis)
	AddApis(WinRegApis)
	AddApis(WinRegANSIApis)
	AddApis(ShellApiApis)
	AddApis(ShellApiANSIApis)
	AddApis(DdeApis)
	AddApis(DdeANSIApis)
	AddApis(WinSvcApis)
	AddApis(WinSvcANSIApis)
	AddApis(Ole2Apis)
	AddApis(CommCtrlApis)
	AddApis(CommCtrlANSIApis)
	AddApis(WinNlsApis)
	AddApis(WinNlsANSIApis)
	AddApis(LZExpandApis)
	AddApis(LZExpandANSIApis)
	AddApis(MMSystemApis)
	AddApis(MMSystemANSIApis)
	AddApis(Nb30Apis)
	//AddApis(WinSockApis)
	AddApis(WinSock2Apis)
	AddApis(WinSock2ANSIApis)
	AddApis(WinNetwkApis)
	AddApis(WinNetwkANSIApis)
}

func AllUnicodeApis() {
	AddApis(WinBaseApis)
	AddApis(WinBaseUnicodeApis)
	AddApis(WinUserApis)
	AddApis(WinUserUnicodeApis)
	AddApis(WinGdiApis)
	AddApis(WinGdiUnicodeApis)
	AddApis(WinConApis)
	AddApis(WinConUnicodeApis)
	AddApis(WinRegApis)
	AddApis(WinRegUnicodeApis)
	AddApis(ShellApiApis)
	AddApis(ShellApiUnicodeApis)
	AddApis(WinSvcApis)
	AddApis(WinSvcUnicodeApis)
	AddApis(Ole2Apis)
	AddApis(CommCtrlApis)
	AddApis(CommCtrlUnicodeApis)
	AddApis(WinNlsApis)
	AddApis(WinNlsUnicodeApis)
	AddApis(LZExpandApis)
	AddApis(LZExpandUnicodeApis)
	AddApis(MMSystemApis)
	AddApis(MMSystemUnicodeApis)
	AddApis(Nb30Apis)
	//AddApis(WinSockApis)
	AddApis(WinSock2Apis)
	AddApis(WinSock2UnicodeApis)
	AddApis(WinNetwkApis)
	AddApis(WinNetwkUnicodeApis)
}
