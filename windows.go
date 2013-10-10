// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package windows.
package windows

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/commctrl"
	. "github.com/tHinqa/outside-windows/dde"
	. "github.com/tHinqa/outside-windows/lzexpand"
	. "github.com/tHinqa/outside-windows/mmsystem"
	. "github.com/tHinqa/outside-windows/nb30"
	. "github.com/tHinqa/outside-windows/ole2"
	. "github.com/tHinqa/outside-windows/shellapi"
	. "github.com/tHinqa/outside-windows/winbase"
	. "github.com/tHinqa/outside-windows/wincon"
	. "github.com/tHinqa/outside-windows/wingdi"
	. "github.com/tHinqa/outside-windows/winnetwk"
	. "github.com/tHinqa/outside-windows/winnls"
	. "github.com/tHinqa/outside-windows/winreg"
	// . "github.com/tHinqa/outside-windows/winsock"
	. "github.com/tHinqa/outside-windows/winsock2"
	. "github.com/tHinqa/outside-windows/winsvc"
	. "github.com/tHinqa/outside-windows/winuser"
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
