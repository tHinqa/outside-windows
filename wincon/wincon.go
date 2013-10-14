// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package wincon provides API definitions for accessing the console
//functions in kernel32.dll.
package wincon

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	_ "github.com/tHinqa/outside/win32/kernel32"
)

var (
	PeekConsoleInput func(
		consoleInput T.HANDLE,
		buffer *T.INPUT_RECORD,
		length T.DWORD,
		numberOfEventsRead *T.DWORD) T.BOOL

	ReadConsoleInput func(
		consoleInput T.HANDLE,
		buffer *T.INPUT_RECORD,
		length T.DWORD,
		numberOfEventsRead *T.DWORD) T.BOOL

	WriteConsoleInput func(
		consoleInput T.HANDLE,
		buffer *T.INPUT_RECORD,
		length T.DWORD,
		numberOfEventsWritten *T.DWORD) T.BOOL

	ReadConsoleOutput func(
		consoleOutput T.HANDLE,
		buffer *T.CHAR_INFO,
		bufferSize T.COORD,
		bufferCoord T.COORD,
		readRegion *T.SMALL_RECT) T.BOOL

	WriteConsoleOutput func(
		consoleOutput T.HANDLE,
		buffer *T.CHAR_INFO,
		bufferSize T.COORD,
		bufferCoord T.COORD,
		writeRegion *T.SMALL_RECT) T.BOOL

	ReadConsoleOutputCharacter func(
		hConsoleOutput T.HANDLE,
		character OVString,
		length T.DWORD,
		readCoord T.COORD,
		numberOfCharsRead *T.DWORD) T.BOOL

	ReadConsoleOutputAttribute func(
		consoleOutput T.HANDLE,
		attribute *T.WORD,
		length T.DWORD,
		readCoord T.COORD,
		numberOfAttrsRead *T.DWORD) T.BOOL

	WriteConsoleOutputCharacter func(
		consoleOutput T.HANDLE,
		character VString,
		length T.DWORD,
		writeCoord T.COORD,
		numberOfCharsWritten *T.DWORD) T.BOOL

	WriteConsoleOutputAttribute func(
		consoleOutput T.HANDLE,
		attribute *T.WORD,
		length T.DWORD,
		writeCoord T.COORD,
		numberOfAttrsWritten *T.DWORD) T.BOOL

	FillConsoleOutputCharacter func(
		consoleOutput T.HANDLE,
		character T.VChar,
		length T.DWORD,
		writeCoord T.COORD,
		numberOfCharsWritten *T.DWORD) T.BOOL

	FillConsoleOutputAttribute func(
		consoleOutput T.HANDLE,
		attribute T.WORD,
		length T.DWORD,
		writeCoord T.COORD,
		numberOfAttrsWritten *T.DWORD) T.BOOL

	GetConsoleMode func(
		consoleHandle T.HANDLE,
		mode *T.DWORD) T.BOOL

	GetNumberOfConsoleInputEvents func(
		consoleInput T.HANDLE,
		numberOfEvents *T.DWORD) T.BOOL

	GetConsoleScreenBufferInfo func(
		consoleOutput T.HANDLE,
		consoleScreenBufferInfo *T.CONSOLE_SCREEN_BUFFER_INFO) T.BOOL

	GetLargestConsoleWindowSize func(
		consoleOutput T.HANDLE) T.COORD

	GetConsoleCursorInfo func(
		consoleOutput T.HANDLE,
		consoleCursorInfo *T.CONSOLE_CURSOR_INFO) T.BOOL

	GetCurrentConsoleFont func(
		consoleOutput T.HANDLE,
		maximumWindow T.BOOL,
		consoleCurrentFont *T.CONSOLE_FONT_INFO) T.BOOL

	GetConsoleFontSize func(
		consoleOutput T.HANDLE,
		font T.DWORD) T.COORD

	GetConsoleSelectionInfo func(
		consoleSelectionInfo *T.CONSOLE_SELECTION_INFO) T.BOOL

	GetNumberOfConsoleMouseButtons func(
		numberOfMouseButtons *T.DWORD) T.BOOL

	SetConsoleMode func(
		consoleHandle T.HANDLE,
		mode T.DWORD) T.BOOL

	SetConsoleActiveScreenBuffer func(
		consoleOutput T.HANDLE) T.BOOL

	FlushConsoleInputBuffer func(
		consoleInput T.HANDLE) T.BOOL

	SetConsoleScreenBufferSize func(
		consoleOutput T.HANDLE,
		size T.COORD) T.BOOL

	SetConsoleCursorPosition func(
		consoleOutput T.HANDLE,
		cursorPosition T.COORD) T.BOOL

	SetConsoleCursorInfo func(
		consoleOutput T.HANDLE,
		consoleCursorInfo *T.CONSOLE_CURSOR_INFO) T.BOOL

	ScrollConsoleScreenBuffer func(
		consoleOutput T.HANDLE,
		scrollRectangle *T.SMALL_RECT,
		clipRectangle *T.SMALL_RECT,
		destinationOrigin T.COORD,
		fill *T.CHAR_INFO) T.BOOL

	SetConsoleWindowInfo func(
		consoleOutput T.HANDLE,
		absolute T.BOOL,
		consoleWindow *T.SMALL_RECT) T.BOOL

	SetConsoleTextAttribute func(
		consoleOutput T.HANDLE, attributes T.WORD) T.BOOL

	SetConsoleCtrlHandler func(
		handlerRoutine *T.HANDLER_ROUTINE, add T.BOOL) T.BOOL

	GenerateConsoleCtrlEvent func(
		ctrlEvent T.DWORD, processGroupId T.DWORD) T.BOOL

	AllocConsole func() T.BOOL

	FreeConsole func() T.BOOL

	AttachConsole func(processId T.DWORD) T.BOOL

	GetConsoleTitle func(
		consoleTitle OVString, size T.DWORD) T.DWORD

	SetConsoleTitle func(
		consoleTitle VString) T.BOOL

	ReadConsole func(
		consoleInputh T.HANDLE,
		buffer *T.VOID,
		numberOfCharsToRead T.DWORD,
		numberOfCharsRead *T.DWORD,
		reserved *T.VOID) T.BOOL

	WriteConsole func(
		consoleOutput T.HANDLE,
		buffer *T.VOID,
		numberOfCharsToWrite T.DWORD,
		numberOfCharsWritten *T.DWORD,
		reserved *T.VOID) T.BOOL

	CreateConsoleScreenBuffer func(
		desiredAccess T.DWORD,
		shareMode T.DWORD,
		securityAttributes *T.SECURITY_ATTRIBUTES,
		flags T.DWORD,
		screenBufferData *T.VOID) T.HANDLE

	GetConsoleCP func() T.UINT

	SetConsoleCP func(codePageID T.UINT) T.BOOL

	GetConsoleOutputCP func() T.UINT

	SetConsoleOutputCP func(wCodePageID T.UINT) T.BOOL

	GetConsoleDisplayMode func(lpModeFlags *T.DWORD) T.BOOL

	GetConsoleWindow func() T.HWND

	GetConsoleProcessList func(
		processList *T.DWORD, processCount T.DWORD) T.DWORD

	AddConsoleAlias func(
		source, Target, exeName VString) T.BOOL

	GetConsoleAlias func(
		source VString,
		targetBuffer OVString,
		targetBufferLength T.DWORD,
		exeName VString) T.DWORD

	GetConsoleAliasesLength func(exeName VString) T.DWORD

	GetConsoleAliasExesLength func() T.DWORD

	GetConsoleAliases func(
		aliasBuffer OVString,
		aliasBufferLength T.DWORD,
		exeName VString) T.DWORD

	GetConsoleAliasExes func(
		exeNameBuffer OVString,
		exeNameBufferLength T.DWORD) T.DWORD
)

var WinConANSIApis = outside.Apis{
	{"AddConsoleAliasA", &AddConsoleAlias},
	{"FillConsoleOutputCharacterA", &FillConsoleOutputCharacter},
	{"GetConsoleAliasA", &GetConsoleAlias},
	{"GetConsoleAliasesA", &GetConsoleAliases},
	{"GetConsoleAliasesLengthA", &GetConsoleAliasesLength},
	{"GetConsoleAliasExesA", &GetConsoleAliasExes},
	{"GetConsoleAliasExesLengthA", &GetConsoleAliasExesLength},
	{"GetConsoleTitleA", &GetConsoleTitle},
	{"PeekConsoleInputA", &PeekConsoleInput},
	{"ReadConsoleA", &ReadConsole},
	{"ReadConsoleInputA", &ReadConsoleInput},
	{"ReadConsoleOutputA", &ReadConsoleOutput},
	{"ReadConsoleOutputCharacterA", &ReadConsoleOutputCharacter},
	{"ScrollConsoleScreenBufferA", &ScrollConsoleScreenBuffer},
	{"SetConsoleTitleA", &SetConsoleTitle},
	{"WriteConsoleA", &WriteConsole},
	{"WriteConsoleInputA", &WriteConsoleInput},
	{"WriteConsoleOutputA", &WriteConsoleOutput},
	{"WriteConsoleOutputCharacterA", &WriteConsoleOutputCharacter},
}

var WinConUnicodeApis = outside.Apis{
	{"AddConsoleAliasW", &AddConsoleAlias},
	{"FillConsoleOutputCharacterW", &FillConsoleOutputCharacter},
	{"GetConsoleAliasW", &GetConsoleAlias},
	{"GetConsoleAliasesW", &GetConsoleAliases},
	{"GetConsoleAliasesLengthW", &GetConsoleAliasesLength},
	{"GetConsoleAliasExesW", &GetConsoleAliasExes},
	{"GetConsoleAliasExesLengthW", &GetConsoleAliasExesLength},
	{"GetConsoleTitleW", &GetConsoleTitle},
	{"PeekConsoleInputW", &PeekConsoleInput},
	{"ReadConsoleW", &ReadConsole},
	{"ReadConsoleInputW", &ReadConsoleInput},
	{"ReadConsoleOutputW", &ReadConsoleOutput},
	{"ReadConsoleOutputCharacterW", &ReadConsoleOutputCharacter},
	{"ScrollConsoleScreenBufferW", &ScrollConsoleScreenBuffer},
	{"SetConsoleTitleW", &SetConsoleTitle},
	{"WriteConsoleW", &WriteConsole},
	{"WriteConsoleInputW", &WriteConsoleInput},
	{"WriteConsoleOutputW", &WriteConsoleOutput},
	{"WriteConsoleOutputCharacterW", &WriteConsoleOutputCharacter},
}

var WinConApis = outside.Apis{
	{"AllocConsole", &AllocConsole},
	{"AttachConsole", &AttachConsole},
	{"CreateConsoleScreenBuffer", &CreateConsoleScreenBuffer},
	{"FillConsoleOutputAttribute", &FillConsoleOutputAttribute},
	{"FlushConsoleInputBuffer", &FlushConsoleInputBuffer},
	{"FreeConsole", &FreeConsole},
	{"GenerateConsoleCtrlEvent", &GenerateConsoleCtrlEvent},
	{"GetConsoleCP", &GetConsoleCP},
	{"GetConsoleCursorInfo", &GetConsoleCursorInfo},
	{"GetConsoleDisplayMode", &GetConsoleDisplayMode},
	{"GetConsoleFontSize", &GetConsoleFontSize},
	{"GetConsoleMode", &GetConsoleMode},
	{"GetConsoleOutputCP", &GetConsoleOutputCP},
	{"GetConsoleProcessList", &GetConsoleProcessList},
	{"GetConsoleScreenBufferInfo", &GetConsoleScreenBufferInfo},
	{"GetConsoleSelectionInfo", &GetConsoleSelectionInfo},
	{"GetConsoleWindow", &GetConsoleWindow},
	{"GetCurrentConsoleFont", &GetCurrentConsoleFont},
	{"GetLargestConsoleWindowSize", &GetLargestConsoleWindowSize},
	{"GetNumberOfConsoleInputEvents", &GetNumberOfConsoleInputEvents},
	{"GetNumberOfConsoleMouseButtons", &GetNumberOfConsoleMouseButtons},
	{"ReadConsoleOutputAttribute", &ReadConsoleOutputAttribute},
	{"SetConsoleActiveScreenBuffer", &SetConsoleActiveScreenBuffer},
	{"SetConsoleCP", &SetConsoleCP},
	{"SetConsoleCtrlHandler", &SetConsoleCtrlHandler},
	{"SetConsoleCursorInfo", &SetConsoleCursorInfo},
	{"SetConsoleCursorPosition", &SetConsoleCursorPosition},
	{"SetConsoleMode", &SetConsoleMode},
	{"SetConsoleOutputCP", &SetConsoleOutputCP},
	{"SetConsoleScreenBufferSize", &SetConsoleScreenBufferSize},
	{"SetConsoleTextAttribute", &SetConsoleTextAttribute},
	{"SetConsoleWindowInfo", &SetConsoleWindowInfo},
	{"WriteConsoleOutputAttribute", &WriteConsoleOutputAttribute},
}
