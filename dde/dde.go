// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package dde provides API definitions for accessing dde
//functions of user32.dll.
package dde

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/user32"
)

var (
	DdeSetQualityOfService func(
		Client HWND,
		QosNew *SECURITY_QUALITY_OF_SERVICE,
		QosPrev *SECURITY_QUALITY_OF_SERVICE) BOOL

	ImpersonateDdeClientWindow func(
		Client HWND,
		Server HWND) BOOL

	PackDDElParam func(
		Msg UINT,
		Lo UINT_PTR,
		Hi UINT_PTR) LPARAM

	UnpackDDElParam func(
		Msg UINT,
		LParam LPARAM,
		Lo *UINT_PTR,
		Hi *UINT_PTR) BOOL

	FreeDDElParam func(
		Msg UINT,
		LParam LPARAM) BOOL

	ReuseDDElParam func(
		LParam LPARAM,
		In UINT,
		Out UINT,
		Lo UINT_PTR,
		Hi UINT_PTR) LPARAM
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
