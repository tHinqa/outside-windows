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
		RemoteName, Password, LocalName VString) (T.DWORD, error)

	WNetAddConnection2 func(
		NetResource *T.NETRESOURCE,
		Password, UserName VString,
		Flags T.DWORD) (T.DWORD, error)

	WNetAddConnection3 func(
		hwndOwner T.HWND,
		NetResource *T.NETRESOURCE,
		Password, UserName VString,
		Flags T.DWORD) (T.DWORD, error)

	WNetCancelConnection func(Name VString, fForce T.BOOL) (T.DWORD, error)

	WNetCancelConnection2 func(
		Name VString, Flags T.DWORD, Force T.BOOL) (T.DWORD, error)

	WNetGetConnection func(
		LocalName VString,
		RemoteName OVString,
		Length *T.DWORD) (T.DWORD, error)

	WNetRestoreConnection func(
		Parent T.HWND,
		Device VString) (T.DWORD, error)

	WNetUseConnection func(
		hwndOwner T.HWND,
		NetResource *T.NETRESOURCE,
		Password, UserID VString,
		Flags T.DWORD,
		AccessName OVString, BufferSize,
		Result *T.DWORD) (T.DWORD, error)

	WNetConnectionDialog func(Wnd T.HWND, Type T.DWORD) (T.DWORD, error)

	WNetDisconnectDialog func(Wnd T.HWND, Type T.DWORD) (T.DWORD, error)

	WNetOpenEnum func(
		Scope, Type, Usage T.DWORD,
		NetResource *T.NETRESOURCE,
		Enum *T.HANDLE) (T.DWORD, error)

	WNetEnumResource func(
		Enum T.HANDLE,
		Count *T.DWORD,
		Buffer *T.VOID,
		BufferSize *T.DWORD) (T.DWORD, error)

	WNetCloseEnum func(
		hEnum T.HANDLE) (T.DWORD, error)

	WNetGetResourceParent func(
		NetResource *T.NETRESOURCE,
		Buffer *T.VOID,
		BufferSize *T.DWORD) (T.DWORD, error)

	WNetGetResourceInformation func(
		NetResource *T.NETRESOURCE,
		Buffer *T.VOID,
		BufferSize *T.DWORD,
		System *OVString) (T.DWORD, error)

	WNetGetUser func(
		Name VString,
		UserName OVString,
		Length *T.DWORD) (T.DWORD, error)

	WNetGetProviderName func(
		NetType T.DWORD,
		ProviderName OVString,
		BufferSize *T.DWORD) (T.DWORD, error)

	WNetGetNetworkInformation func(
		Provider VString,
		NetInfoStruct *T.NETINFOSTRUCT) (T.DWORD, error)

	WNetGetLastError func(
		Error *T.DWORD,
		ErrorBuf OVString,
		ErrorBufSize T.DWORD,
		NameBuf OVString,
		NameBufSize T.DWORD) (T.DWORD, error)

	MultinetGetConnectionPerformance func(
		*T.NETRESOURCE, *T.NETCONNECTINFOSTRUCT) (T.DWORD, error)
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
