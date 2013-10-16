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
		numberOfEventsRead *T.DWORD) (T.BOOL, error)

	ReadConsoleInput func(
		consoleInput T.HANDLE,
		buffer *T.INPUT_RECORD,
		length T.DWORD,
		numberOfEventsRead *T.DWORD) (T.BOOL, error)

	WriteConsoleInput func(
		consoleInput T.HANDLE,
		buffer *T.INPUT_RECORD,
		length T.DWORD,
		numberOfEventsWritten *T.DWORD) (T.BOOL, error)

	ReadConsoleOutput func(
		consoleOutput T.HANDLE,
		buffer *T.CHAR_INFO,
		bufferSize T.COORD,
		bufferCoord T.COORD,
		readRegion *T.SMALL_RECT) (T.BOOL, error)

	WriteConsoleOutput func(
		consoleOutput T.HANDLE,
		buffer *T.CHAR_INFO,
		bufferSize T.COORD,
		bufferCoord T.COORD,
		writeRegion *T.SMALL_RECT) (T.BOOL, error)

	ReadConsoleOutputCharacter func(
		hConsoleOutput T.HANDLE,
		character OVString,
		length T.DWORD,
		readCoord T.COORD,
		numberOfCharsRead *T.DWORD) (T.BOOL, error)

	ReadConsoleOutputAttribute func(
		consoleOutput T.HANDLE,
		attribute *T.WORD,
		length T.DWORD,
		readCoord T.COORD,
		numberOfAttrsRead *T.DWORD) (T.BOOL, error)

	WriteConsoleOutputCharacter func(
		consoleOutput T.HANDLE,
		character VString,
		length T.DWORD,
		writeCoord T.COORD,
		numberOfCharsWritten *T.DWORD) (T.BOOL, error)

	WriteConsoleOutputAttribute func(
		consoleOutput T.HANDLE,
		attribute *T.WORD,
		length T.DWORD,
		writeCoord T.COORD,
		numberOfAttrsWritten *T.DWORD) (T.BOOL, error)

	FillConsoleOutputCharacter func(
		consoleOutput T.HANDLE,
		character T.VChar,
		length T.DWORD,
		writeCoord T.COORD,
		numberOfCharsWritten *T.DWORD) (T.BOOL, error)

	FillConsoleOutputAttribute func(
		consoleOutput T.HANDLE,
		attribute T.WORD,
		length T.DWORD,
		writeCoord T.COORD,
		numberOfAttrsWritten *T.DWORD) (T.BOOL, error)

	GetConsoleMode func(
		consoleHandle T.HANDLE,
		mode *T.DWORD) (T.BOOL, error)

	GetNumberOfConsoleInputEvents func(
		consoleInput T.HANDLE,
		numberOfEvents *T.DWORD) (T.BOOL, error)

	GetConsoleScreenBufferInfo func(
		consoleOutput T.HANDLE,
		consoleScreenBufferInfo *T.CONSOLE_SCREEN_BUFFER_INFO) (T.BOOL, error)

	GetLargestConsoleWindowSize func(
		consoleOutput T.HANDLE) (T.COORD, error)

	GetConsoleCursorInfo func(
		consoleOutput T.HANDLE,
		consoleCursorInfo *T.CONSOLE_CURSOR_INFO) (T.BOOL, error)

	GetCurrentConsoleFont func(
		consoleOutput T.HANDLE,
		maximumWindow T.BOOL,
		consoleCurrentFont *T.CONSOLE_FONT_INFO) (T.BOOL, error)

	GetConsoleFontSize func(
		consoleOutput T.HANDLE,
		font T.DWORD) (T.COORD, error)

	GetConsoleSelectionInfo func(
		consoleSelectionInfo *T.CONSOLE_SELECTION_INFO) (T.BOOL, error)

	GetNumberOfConsoleMouseButtons func(
		numberOfMouseButtons *T.DWORD) (T.BOOL, error)

	SetConsoleMode func(
		consoleHandle T.HANDLE,
		mode T.DWORD) (T.BOOL, error)

	SetConsoleActiveScreenBuffer func(
		consoleOutput T.HANDLE) (T.BOOL, error)

	FlushConsoleInputBuffer func(
		consoleInput T.HANDLE) (T.BOOL, error)

	SetConsoleScreenBufferSize func(
		consoleOutput T.HANDLE,
		size T.COORD) (T.BOOL, error)

	SetConsoleCursorPosition func(
		consoleOutput T.HANDLE,
		cursorPosition T.COORD) (T.BOOL, error)

	SetConsoleCursorInfo func(
		consoleOutput T.HANDLE,
		consoleCursorInfo *T.CONSOLE_CURSOR_INFO) (T.BOOL, error)

	ScrollConsoleScreenBuffer func(
		consoleOutput T.HANDLE,
		scrollRectangle *T.SMALL_RECT,
		clipRectangle *T.SMALL_RECT,
		destinationOrigin T.COORD,
		fill *T.CHAR_INFO) (T.BOOL, error)

	SetConsoleWindowInfo func(
		consoleOutput T.HANDLE,
		absolute T.BOOL,
		consoleWindow *T.SMALL_RECT) (T.BOOL, error)

	SetConsoleTextAttribute func(
		consoleOutput T.HANDLE, attributes T.WORD) (T.BOOL, error)

	SetConsoleCtrlHandler func(
		handlerRoutine *T.HANDLER_ROUTINE, add T.BOOL) (T.BOOL, error)

	GenerateConsoleCtrlEvent func(
		ctrlEvent T.DWORD, processGroupId T.DWORD) (T.BOOL, error)

	AllocConsole func() (T.BOOL, error)

	FreeConsole func() (T.BOOL, error)

	AttachConsole func(processId T.DWORD) (T.BOOL, error)

	GetConsoleTitle func(
		consoleTitle OVString, size T.DWORD) (T.DWORD, error)

	SetConsoleTitle func(
		consoleTitle VString) (T.BOOL, error)

	ReadConsole func(
		consoleInputh T.HANDLE,
		buffer *T.VOID,
		numberOfCharsToRead T.DWORD,
		numberOfCharsRead *T.DWORD,
		reserved *T.VOID) (T.BOOL, error)

	WriteConsole func(
		consoleOutput T.HANDLE,
		buffer *T.VOID,
		numberOfCharsToWrite T.DWORD,
		numberOfCharsWritten *T.DWORD,
		reserved *T.VOID) (T.BOOL, error)

	CreateConsoleScreenBuffer func(
		desiredAccess T.DWORD,
		shareMode T.DWORD,
		securityAttributes *T.SECURITY_ATTRIBUTES,
		flags T.DWORD,
		screenBufferData *T.VOID) (T.HANDLE, error)

	GetConsoleCP func() (T.UINT, error)

	SetConsoleCP func(codePageID T.UINT) (T.BOOL, error)

	GetConsoleOutputCP func() (T.UINT, error)

	SetConsoleOutputCP func(wCodePageID T.UINT) (T.BOOL, error)

	GetConsoleDisplayMode func(lpModeFlags *T.DWORD) (T.BOOL, error)

	GetConsoleWindow func() (T.HWND, error)

	GetConsoleProcessList func(
		processList *T.DWORD, processCount T.DWORD) (T.DWORD, error)

	AddConsoleAlias func(
		source, Target, exeName VString) (T.BOOL, error)

	GetConsoleAlias func(
		source VString,
		targetBuffer OVString,
		targetBufferLength T.DWORD,
		exeName VString) (T.DWORD, error)

	GetConsoleAliasesLength func(exeName VString) (T.DWORD, error)

	GetConsoleAliasExesLength func() (T.DWORD, error)

	GetConsoleAliases func(
		aliasBuffer OVString,
		aliasBufferLength T.DWORD,
		exeName VString) (T.DWORD, error)

	GetConsoleAliasExes func(
		exeNameBuffer OVString,
		exeNameBufferLength T.DWORD) (T.DWORD, error)
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
