package winnetwk

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/mpr"
)

var (
	WNetAddConnection func(
		RemoteName, Password, LocalName VString) DWORD

	WNetAddConnection2 func(
		NetResource *NETRESOURCE,
		Password, UserName VString,
		Flags DWORD) DWORD

	WNetAddConnection3 func(
		hwndOwner HWND,
		NetResource *NETRESOURCE,
		Password, UserName VString,
		Flags DWORD) DWORD

	WNetCancelConnection func(Name VString, fForce BOOL) DWORD

	WNetCancelConnection2 func(
		Name VString, Flags DWORD, Force BOOL) DWORD

	WNetGetConnection func(
		LocalName VString,
		RemoteName OVString,
		Length *DWORD) DWORD

	WNetRestoreConnection func(
		Parent HWND,
		Device VString) DWORD

	WNetUseConnection func(
		hwndOwner HWND,
		NetResource *NETRESOURCE,
		Password, UserID VString,
		Flags DWORD,
		AccessName OVString, BufferSize,
		Result *DWORD) DWORD

	WNetConnectionDialog func(Wnd HWND, Type DWORD) DWORD

	WNetDisconnectDialog func(Wnd HWND, Type DWORD) DWORD

	WNetOpenEnum func(
		Scope, Type, Usage DWORD,
		NetResource *NETRESOURCE,
		Enum *HANDLE) DWORD

	WNetEnumResource func(
		Enum HANDLE,
		Count *DWORD,
		Buffer *VOID,
		BufferSize *DWORD) DWORD

	WNetCloseEnum func(
		hEnum HANDLE) DWORD

	WNetGetResourceParent func(
		NetResource *NETRESOURCE,
		Buffer *VOID,
		BufferSize *DWORD) DWORD

	WNetGetResourceInformation func(
		NetResource *NETRESOURCE,
		Buffer *VOID,
		BufferSize *DWORD,
		System *OVString) DWORD

	WNetGetUser func(
		Name VString,
		UserName OVString,
		Length *DWORD) DWORD

	WNetGetProviderName func(
		NetType DWORD,
		ProviderName OVString,
		BufferSize *DWORD) DWORD

	WNetGetNetworkInformation func(
		Provider VString,
		NetInfoStruct *NETINFOSTRUCT) DWORD

	WNetGetLastError func(
		Error *DWORD,
		ErrorBuf OVString,
		ErrorBufSize DWORD,
		NameBuf OVString,
		NameBufSize DWORD) DWORD

	MultinetGetConnectionPerformance func(
		*NETRESOURCE, *NETCONNECTINFOSTRUCT) DWORD
)

var WinNetwkANSIApis = Apis{
	{"MultinetGetConnectionPerformanceA", &MultinetGetConnectionPerformance},
	{"WNetAddConnectionA", &WNetAddConnection},
	{"WNetAddConnection2A", &WNetAddConnection2},
	{"WNetAddConnection3A", &WNetAddConnection3},
	{"WNetCancelConnectionA", &WNetCancelConnection},
	{"WNetCancelConnection2A", &WNetCancelConnection2},
	{"WNetEnumResourceA", &WNetEnumResource},
	{"WNetGetConnectionA", &WNetGetConnection},
	{"WNetGetLastErrorA", &WNetGetLastError},
	{"WNetGetNetworkInformationA", &WNetGetNetworkInformation},
	{"WNetGetProviderNameA", &WNetGetProviderName},
	{"WNetGetResourceInformationA", &WNetGetResourceInformation},
	{"WNetGetResourceParentA", &WNetGetResourceParent},
	{"WNetGetUserA", &WNetGetUser},
	{"WNetOpenEnumA", &WNetOpenEnum},
	{"WNetUseConnectionA", &WNetUseConnection},
}

var WinNetwkUnicodeApis = Apis{
	{"MultinetGetConnectionPerformanceW", &MultinetGetConnectionPerformance},
	{"WNetAddConnectionW", &WNetAddConnection},
	{"WNetAddConnection2W", &WNetAddConnection2},
	{"WNetAddConnection3W", &WNetAddConnection3},
	{"WNetCancelConnectionW", &WNetCancelConnection},
	{"WNetCancelConnection2W", &WNetCancelConnection2},
	{"WNetEnumResourceW", &WNetEnumResource},
	{"WNetGetConnectionW", &WNetGetConnection},
	{"WNetGetLastErrorW", &WNetGetLastError},
	{"WNetGetNetworkInformationW", &WNetGetNetworkInformation},
	{"WNetGetProviderNameW", &WNetGetProviderName},
	{"WNetGetResourceInformationW", &WNetGetResourceInformation},
	{"WNetGetResourceParentW", &WNetGetResourceParent},
	{"WNetGetUserW", &WNetGetUser},
	{"WNetOpenEnumW", &WNetOpenEnum},
	{"WNetUseConnectionW", &WNetUseConnection},
}

var WinNetwkApis = Apis{
	{"WNetCloseEnum", &WNetCloseEnum},
	{"WNetConnectionDialog", &WNetConnectionDialog},
	{"WNetDisconnectDialog", &WNetDisconnectDialog},
	{"WNetRestoreConnectionW", &WNetRestoreConnection},
}
