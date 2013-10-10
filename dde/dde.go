// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package dde provides API definitions for accessing dde
//functions of user32.dll.
package dde

import (
	. "github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/user32"
)

var (
	DdeSetQualityOfService func(
		Client T.HWND,
		QosNew *T.SECURITY_QUALITY_OF_SERVICE,
		QosPrev *T.SECURITY_QUALITY_OF_SERVICE) T.BOOL

	ImpersonateDdeClientWindow func(
		Client T.HWND,
		Server T.HWND) T.BOOL

	PackDDElParam func(
		Msg T.UINT,
		Lo T.UINT_PTR,
		Hi T.UINT_PTR) T.LPARAM

	UnpackDDElParam func(
		Msg T.UINT,
		LParam T.LPARAM,
		Lo *T.UINT_PTR,
		Hi *T.UINT_PTR) T.BOOL

	FreeDDElParam func(
		Msg T.UINT,
		LParam T.LPARAM) T.BOOL

	ReuseDDElParam func(
		LParam T.LPARAM,
		In T.UINT,
		Out T.UINT,
		Lo T.UINT_PTR,
		Hi T.UINT_PTR) T.LPARAM
)

var DdeANSIApis = Apis{
	{"DdeCreateStringHandleA", &DdeCreateStringHandle},
	{"DdeInitializeA", &DdeInitialize},
	{"DdeQueryStringA", &DdeQueryString},
}

var DdeUnicodeApis = Apis{
	{"DdeCreateStringHandleW", &DdeCreateStringHandle},
	{"DdeInitializeW", &DdeInitialize},
	{"DdeQueryStringW", &DdeQueryString},
}

var DdeApis = Apis{
	{"DdeAbandonTransaction", &DdeAbandonTransaction},
	{"DdeAccessData", &DdeAccessData},
	{"DdeAddData", &DdeAddData},
	{"DdeClientTransaction", &DdeClientTransaction},
	{"DdeCmpStringHandles", &DdeCmpStringHandles},
	{"DdeConnect", &DdeConnect},
	{"DdeConnectList", &DdeConnectList},
	{"DdeCreateDataHandle", &DdeCreateDataHandle},
	{"DdeDisconnect", &DdeDisconnect},
	{"DdeDisconnectList", &DdeDisconnectList},
	{"DdeEnableCallback", &DdeEnableCallback},
	{"DdeFreeDataHandle", &DdeFreeDataHandle},
	{"DdeFreeStringHandle", &DdeFreeStringHandle},
	{"DdeGetData", &DdeGetData},
	{"DdeGetLastError", &DdeGetLastError},
	{"DdeImpersonateClient", &DdeImpersonateClient},
	{"DdeKeepStringHandle", &DdeKeepStringHandle},
	{"DdeNameService", &DdeNameService},
	{"DdePostAdvise", &DdePostAdvise},
	{"DdeQueryConvInfo", &DdeQueryConvInfo},
	{"DdeQueryNextServer", &DdeQueryNextServer},
	{"DdeReconnect", &DdeReconnect},
	{"DdeSetQualityOfService", &DdeSetQualityOfService},
	{"DdeSetUserHandle", &DdeSetUserHandle},
	{"DdeUnaccessData", &DdeUnaccessData},
	{"DdeUninitialize", &DdeUninitialize},
	{"FreeDDElParam", &FreeDDElParam},
	{"ImpersonateDdeClientWindow", &ImpersonateDdeClientWindow},
	{"PackDDElParam", &PackDDElParam},
	{"ReuseDDElParam", &ReuseDDElParam},
	{"UnpackDDElParam", &UnpackDDElParam},
}
