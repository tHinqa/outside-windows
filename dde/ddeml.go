// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package dde

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/user32"
)

var (
	DdeInitialize func(
		Inst *DWORD,
		Callback *FNCALLBACK,
		Cmd DWORD,
		Res DWORD) UINT

	DdeUninitialize func(Inst DWORD) BOOL

	DdeConnectList func(
		Inst DWORD,
		Service HSZ,
		Topic HSZ,
		ConvList HCONVLIST,
		CC *CONVCONTEXT) HCONVLIST

	DdeQueryNextServer func(
		ConvList HCONVLIST, ConvPrev HCONV) HCONV

	DdeDisconnectList func(ConvList HCONVLIST) BOOL

	DdeConnect func(
		Inst DWORD,
		Service HSZ,
		Topic HSZ,
		CC *CONVCONTEXT) HCONV

	DdeDisconnect func(Conv HCONV) BOOL

	DdeReconnect func(Conv HCONV) HCONV

	DdeQueryConvInfo func(
		Conv HCONV,
		Transaction DWORD,
		ConvInfo *CONVINFO) UINT

	DdeSetUserHandle func(
		Conv HCONV, Id DWORD, User DWORD_PTR) BOOL

	DdeAbandonTransaction func(
		Inst DWORD, Conv HCONV, Transaction DWORD) BOOL

	DdePostAdvise func(
		Inst DWORD, Topic HSZ, Item HSZ) BOOL

	DdeEnableCallback func(
		Inst DWORD, Conv HCONV, Cmd UINT) BOOL

	DdeImpersonateClient func(Conv HCONV) BOOL

	DdeNameService func(
		Inst DWORD, S1 HSZ, S2 HSZ, Cmd UINT) HDDEDATA

	DdeClientTransaction func(
		Data *BYTE,
		cData DWORD,
		Conv HCONV,
		Item HSZ,
		Fmt UINT,
		Type UINT,
		Timeout DWORD,
		Result *DWORD) HDDEDATA

	DdeCreateDataHandle func(
		Inst DWORD,
		Src *BYTE,
		C DWORD,
		Off DWORD,
		Item HSZ,
		Fmt UINT,
		Cmd UINT) HDDEDATA

	DdeAddData func(
		Data HDDEDATA,
		Src *BYTE,
		C DWORD,
		Off DWORD) HDDEDATA

	DdeGetData func(
		Data HDDEDATA,
		Dst *BYTE,
		Max DWORD,
		Off DWORD) DWORD

	DdeAccessData func(
		Data HDDEDATA, DataSize *DWORD) *BYTE

	DdeUnaccessData func(Data HDDEDATA) BOOL

	DdeFreeDataHandle func(Data HDDEDATA) BOOL

	DdeGetLastError func(Inst DWORD) UINT

	DdeCreateStringHandle func(
		Inst DWORD, S VString, CodePage int) HSZ

	DdeQueryString func(
		Inst DWORD,
		S HSZ,
		OS OVString,
		hMax DWORD,
		CodePage int) DWORD

	DdeFreeStringHandle func(Inst DWORD, S HSZ) BOOL

	DdeKeepStringHandle func(Inst DWORD, S HSZ) BOOL

	DdeCmpStringHandles func(S1 HSZ, S2 HSZ) int
)
