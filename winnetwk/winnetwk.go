// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package winnetwk provides API definitions for accessing
//mpr.dll.
package winnetwk

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	_ "github.com/tHinqa/outside/win32/mpr"
)

var (
	WNetAddConnection func(
		RemoteName, Password, LocalName VString) T.DWORD

	WNetAddConnection2 func(
		NetResource *T.NETRESOURCE,
		Password, UserName VString,
		Flags T.DWORD) T.DWORD

	WNetAddConnection3 func(
		hwndOwner T.HWND,
		NetResource *T.NETRESOURCE,
		Password, UserName VString,
		Flags T.DWORD) T.DWORD

	WNetCancelConnection func(Name VString, fForce T.BOOL) T.DWORD

	WNetCancelConnection2 func(
		Name VString, Flags T.DWORD, Force T.BOOL) T.DWORD

	WNetGetConnection func(
		LocalName VString,
		RemoteName OVString,
		Length *T.DWORD) T.DWORD

	WNetRestoreConnection func(
		Parent T.HWND,
		Device VString) T.DWORD

	WNetUseConnection func(
		hwndOwner T.HWND,
		NetResource *T.NETRESOURCE,
		Password, UserID VString,
		Flags T.DWORD,
		AccessName OVString, BufferSize,
		Result *T.DWORD) T.DWORD

	WNetConnectionDialog func(Wnd T.HWND, Type T.DWORD) T.DWORD

	WNetDisconnectDialog func(Wnd T.HWND, Type T.DWORD) T.DWORD

	WNetOpenEnum func(
		Scope, Type, Usage T.DWORD,
		NetResource *T.NETRESOURCE,
		Enum *T.HANDLE) T.DWORD

	WNetEnumResource func(
		Enum T.HANDLE,
		Count *T.DWORD,
		Buffer *T.VOID,
		BufferSize *T.DWORD) T.DWORD

	WNetCloseEnum func(
		hEnum T.HANDLE) T.DWORD

	WNetGetResourceParent func(
		NetResource *T.NETRESOURCE,
		Buffer *T.VOID,
		BufferSize *T.DWORD) T.DWORD

	WNetGetResourceInformation func(
		NetResource *T.NETRESOURCE,
		Buffer *T.VOID,
		BufferSize *T.DWORD,
		System *OVString) T.DWORD

	WNetGetUser func(
		Name VString,
		UserName OVString,
		Length *T.DWORD) T.DWORD

	WNetGetProviderName func(
		NetType T.DWORD,
		ProviderName OVString,
		BufferSize *T.DWORD) T.DWORD

	WNetGetNetworkInformation func(
		Provider VString,
		NetInfoStruct *T.NETINFOSTRUCT) T.DWORD

	WNetGetLastError func(
		Error *T.DWORD,
		ErrorBuf OVString,
		ErrorBufSize T.DWORD,
		NameBuf OVString,
		NameBufSize T.DWORD) T.DWORD

	MultinetGetConnectionPerformance func(
		*T.NETRESOURCE, *T.NETCONNECTINFOSTRUCT) T.DWORD
)

var WinNetwkANSIApis = outside.Apis{
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

var WinNetwkUnicodeApis = outside.Apis{
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

var WinNetwkApis = outside.Apis{
	{"WNetCloseEnum", &WNetCloseEnum},
	{"WNetConnectionDialog", &WNetConnectionDialog},
	{"WNetDisconnectDialog", &WNetDisconnectDialog},
	{"WNetRestoreConnectionW", &WNetRestoreConnection},
}
