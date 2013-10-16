// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package lzexpand provides API definitions for accessing
//the lz functions in kernel32.dll.
package lzexpand

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	_ "github.com/tHinqa/outside/win32/kernel32"
)

var (
	LZStart func() (T.INT, error)

	LZDone func() (T.VOID, error)

	CopyLZFile func(T.INT, T.INT) (T.LONG, error)

	LZCopy func(T.INT, T.INT) (T.LONG, error)

	LZInit func(T.INT) (T.INT, error)

	GetExpandedName func(VString, OVString) (T.INT, error)

	LZOpenFile func(OVString, *T.OFSTRUCT, T.WORD) (T.INT, error)

	LZSeek func(T.INT, T.LONG, T.INT) (T.LONG, error)

	LZRead func(T.INT, T.AString, T.INT) (T.INT, error)

	LZClose func(T.INT) (T.VOID, error)
)

var LZExpandANSIApis = outside.Apis{
	{"GetExpandedNameA", &GetExpandedName},
	{"LZOpenFileA", &LZOpenFile},
}

var LZExpandUnicodeApis = outside.Apis{
	{"GetExpandedNameW", &GetExpandedName},
	{"LZOpenFileW", &LZOpenFile},
}

var LZExpandApis = outside.Apis{
	{"CopyLZFile", &CopyLZFile},
	{"LZClose", &LZClose},
	{"LZCopy", &LZCopy},
	{"LZDone", &LZDone},
	{"LZInit", &LZInit},
	{"LZRead", &LZRead},
	{"LZSeek", &LZSeek},
	{"LZStart", &LZStart},
}
