package lzexpand

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/kernel32"
)

var (
	LZStart func() INT

	LZDone func() VOID

	CopyLZFile func(INT, INT) LONG

	LZCopy func(INT, INT) LONG

	LZInit func(INT) INT

	GetExpandedName func(VString, OVString) INT

	LZOpenFile func(OVString, *OFSTRUCT, WORD) INT

	LZSeek func(INT, LONG, INT) LONG

	LZRead func(INT, AString, INT) INT

	LZClose func(INT) VOID
)

var LZExpandANSIApis = Apis{
	{"GetExpandedNameA", &GetExpandedName},
	{"LZOpenFileA", &LZOpenFile},
}

var LZExpandUnicodeApis = Apis{
	{"GetExpandedNameW", &GetExpandedName},
	{"LZOpenFileW", &LZOpenFile},
}

var LZExpandApis = Apis{
	{"CopyLZFile", &CopyLZFile},
	{"LZClose", &LZClose},
	{"LZCopy", &LZCopy},
	{"LZDone", &LZDone},
	{"LZInit", &LZInit},
	{"LZRead", &LZRead},
	{"LZSeek", &LZSeek},
	{"LZStart", &LZStart},
}
