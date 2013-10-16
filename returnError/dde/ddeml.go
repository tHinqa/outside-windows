// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package dde

import (
	. "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	_ "github.com/tHinqa/outside/win32/user32"
)

var (
	DdeInitialize func(
		Inst *DWORD,
		Callback *FNCALLBACK,
		Cmd DWORD,
		Res DWORD) (UINT, error)

	DdeUninitialize func(Inst DWORD) (BOOL, error)

	DdeConnectList func(
		Inst DWORD,
		Service HSZ,
		Topic HSZ,
		ConvList HCONVLIST,
		CC *CONVCONTEXT) (HCONVLIST, error)

	DdeQueryNextServer func(
		ConvList HCONVLIST, ConvPrev HCONV) (HCONV, error)

	DdeDisconnectList func(ConvList HCONVLIST) (BOOL, error)

	DdeConnect func(
		Inst DWORD,
		Service HSZ,
		Topic HSZ,
		CC *CONVCONTEXT) (HCONV, error)

	DdeDisconnect func(Conv HCONV) (BOOL, error)

	DdeReconnect func(Conv HCONV) (HCONV, error)

	DdeQueryConvInfo func(
		Conv HCONV,
		Transaction DWORD,
		ConvInfo *CONVINFO) (UINT, error)

	DdeSetUserHandle func(
		Conv HCONV, Id DWORD, User DWORD_PTR) (BOOL, error)

	DdeAbandonTransaction func(
		Inst DWORD, Conv HCONV, Transaction DWORD) (BOOL, error)

	DdePostAdvise func(
		Inst DWORD, Topic HSZ, Item HSZ) (BOOL, error)

	DdeEnableCallback func(
		Inst DWORD, Conv HCONV, Cmd UINT) (BOOL, error)

	DdeImpersonateClient func(Conv HCONV) (BOOL, error)

	DdeNameService func(
		Inst DWORD, S1 HSZ, S2 HSZ, Cmd UINT) (HDDEDATA, error)

	DdeClientTransaction func(
		Data *BYTE,
		cData DWORD,
		Conv HCONV,
		Item HSZ,
		Fmt UINT,
		Type UINT,
		Timeout DWORD,
		Result *DWORD) (HDDEDATA, error)

	DdeCreateDataHandle func(
		Inst DWORD,
		Src *BYTE,
		C DWORD,
		Off DWORD,
		Item HSZ,
		Fmt UINT,
		Cmd UINT) (HDDEDATA, error)

	DdeAddData func(
		Data HDDEDATA,
		Src *BYTE,
		C DWORD,
		Off DWORD) (HDDEDATA, error)

	DdeGetData func(
		Data HDDEDATA,
		Dst *BYTE,
		Max DWORD,
		Off DWORD) (DWORD, error)

	DdeAccessData func(
		Data HDDEDATA, DataSize *DWORD) (*BYTE, error)

	DdeUnaccessData func(Data HDDEDATA) (BOOL, error)

	DdeFreeDataHandle func(Data HDDEDATA) (BOOL, error)

	DdeGetLastError func(Inst DWORD) (UINT, error)

	DdeCreateStringHandle func(
		Inst DWORD, S VString, CodePage int) (HSZ, error)

	DdeQueryString func(
		Inst DWORD,
		S HSZ,
		OS OVString,
		hMax DWORD,
		CodePage int) (DWORD, error)

	DdeFreeStringHandle func(Inst DWORD, S HSZ) (BOOL, error)

	DdeKeepStringHandle func(Inst DWORD, S HSZ) (BOOL, error)

	DdeCmpStringHandles func(S1 HSZ, S2 HSZ) (int, error)
)
