// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package wincon provides API definitions for accessing the console
//functions in kernel32.dll.
package wincon

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/kernel32"
)

var (
	PeekConsoleInput func(
		ConsoleInput HANDLE,
		Buffer *INPUT_RECORD,
		Length DWORD,
		NumberOfEventsRead *DWORD) BOOL

	ReadConsoleInput func(
		ConsoleInput HANDLE,
		Buffer *INPUT_RECORD,
		Length DWORD,
		NumberOfEventsRead *DWORD) BOOL

	WriteConsoleInput func(
		ConsoleInput HANDLE,
		Buffer *INPUT_RECORD,
		Length DWORD,
		NumberOfEventsWritten *DWORD) BOOL

	ReadConsoleOutput func(
		ConsoleOutput HANDLE,
		Buffer *CHAR_INFO,
		BufferSize COORD,
		BufferCoord COORD,
		ReadRegion *SMALL_RECT) BOOL

	WriteConsoleOutput func(
		ConsoleOutput HANDLE,
		Buffer *CHAR_INFO,
		BufferSize COORD,
		BufferCoord COORD,
		WriteRegion *SMALL_RECT) BOOL

	ReadConsoleOutputCharacter func(
		hConsoleOutput HANDLE,
		Character OVString,
		Length DWORD,
		ReadCoord COORD,
		NumberOfCharsRead *DWORD) BOOL

	ReadConsoleOutputAttribute func(
		ConsoleOutput HANDLE,
		Attribute *WORD,
		Length DWORD,
		ReadCoord COORD,
		NumberOfAttrsRead *DWORD) BOOL

	WriteConsoleOutputCharacter func(
		ConsoleOutput HANDLE,
		Character VString,
		Length DWORD,
		WriteCoord COORD,
		NumberOfCharsWritten *DWORD) BOOL

	WriteConsoleOutputAttribute func(
		ConsoleOutput HANDLE,
		Attribute *WORD,
		Length DWORD,
		WriteCoord COORD,
		NumberOfAttrsWritten *DWORD) BOOL

	FillConsoleOutputCharacter func(
		ConsoleOutput HANDLE,
		Character VChar,
		Length DWORD,
		WriteCoord COORD,
		NumberOfCharsWritten *DWORD) BOOL

	FillConsoleOutputAttribute func(
		ConsoleOutput HANDLE,
		Attribute WORD,
		Length DWORD,
		WriteCoord COORD,
		NumberOfAttrsWritten *DWORD) BOOL

	GetConsoleMode func(
		ConsoleHandle HANDLE,
		Mode *DWORD) BOOL

	GetNumberOfConsoleInputEvents func(
		ConsoleInput HANDLE,
		NumberOfEvents *DWORD) BOOL

	GetConsoleScreenBufferInfo func(
		ConsoleOutput HANDLE,
		ConsoleScreenBufferInfo *CONSOLE_SCREEN_BUFFER_INFO) BOOL

	GetLargestConsoleWindowSize func(
		ConsoleOutput HANDLE) COORD

	GetConsoleCursorInfo func(
		ConsoleOutput HANDLE,
		ConsoleCursorInfo *CONSOLE_CURSOR_INFO) BOOL

	GetCurrentConsoleFont func(
		ConsoleOutput HANDLE,
		MaximumWindow BOOL,
		ConsoleCurrentFont *CONSOLE_FONT_INFO) BOOL

	GetConsoleFontSize func(
		ConsoleOutput HANDLE,
		Font DWORD) COORD

	GetConsoleSelectionInfo func(
		ConsoleSelectionInfo *CONSOLE_SELECTION_INFO) BOOL

	GetNumberOfConsoleMouseButtons func(
		NumberOfMouseButtons *DWORD) BOOL

	SetConsoleMode func(
		ConsoleHandle HANDLE,
		Mode DWORD) BOOL

	SetConsoleActiveScreenBuffer func(
		ConsoleOutput HANDLE) BOOL

	FlushConsoleInputBuffer func(
		ConsoleInput HANDLE) BOOL

	SetConsoleScreenBufferSize func(
		ConsoleOutput HANDLE,
		Size COORD) BOOL

	SetConsoleCursorPosition func(
		ConsoleOutput HANDLE,
		CursorPosition COORD) BOOL

	SetConsoleCursorInfo func(
		ConsoleOutput HANDLE,
		ConsoleCursorInfo *CONSOLE_CURSOR_INFO) BOOL

	ScrollConsoleScreenBuffer func(
		ConsoleOutput HANDLE,
		ScrollRectangle *SMALL_RECT,
		ClipRectangle *SMALL_RECT,
		DestinationOrigin COORD,
		Fill *CHAR_INFO) BOOL

	SetConsoleWindowInfo func(
		ConsoleOutput HANDLE,
		Absolute BOOL,
		ConsoleWindow *SMALL_RECT) BOOL

	SetConsoleTextAttribute func(
		ConsoleOutput HANDLE, Attributes WORD) BOOL

	SetConsoleCtrlHandler func(
		HandlerRoutine *HANDLER_ROUTINE, Add BOOL) BOOL

	GenerateConsoleCtrlEvent func(
		CtrlEvent DWORD, ProcessGroupId DWORD) BOOL

	AllocConsole func() BOOL

	FreeConsole func() BOOL

	AttachConsole func(ProcessId DWORD) BOOL

	GetConsoleTitle func(
		ConsoleTitle OVString, Size DWORD) DWORD

	SetConsoleTitle func(
		ConsoleTitle VString) BOOL

	ReadConsole func(
		ConsoleInput HANDLE,
		Buffer *VOID,
		NumberOfCharsToRead DWORD,
		NumberOfCharsRead *DWORD,
		Reserved *VOID) BOOL

	WriteConsole func(
		ConsoleOutput HANDLE,
		Buffer *VOID,
		NumberOfCharsToWrite DWORD,
		NumberOfCharsWritten *DWORD,
		Reserved *VOID) BOOL

	CreateConsoleScreenBuffer func(
		DesiredAccess DWORD,
		ShareMode DWORD,
		SecurityAttributes *SECURITY_ATTRIBUTES,
		Flags DWORD,
		ScreenBufferData *VOID) HANDLE

	GetConsoleCP func() UINT

	SetConsoleCP func(CodePageID UINT) BOOL

	GetConsoleOutputCP func() UINT

	SetConsoleOutputCP func(wCodePageID UINT) BOOL

	GetConsoleDisplayMode func(lpModeFlags *DWORD) BOOL

	GetConsoleWindow func() HWND

	GetConsoleProcessList func(
		ProcessList *DWORD, ProcessCount DWORD) DWORD

	AddConsoleAlias func(
		Source, Target, ExeName VString) BOOL

	GetConsoleAlias func(
		Source VString,
		TargetBuffer OVString,
		TargetBufferLength DWORD,
		ExeName VString) DWORD

	GetConsoleAliasesLength func(ExeName VString) DWORD

	GetConsoleAliasExesLength func() DWORD

	GetConsoleAliases func(
		AliasBuffer OVString,
		AliasBufferLength DWORD,
		ExeName VString) DWORD

	GetConsoleAliasExes func(
		ExeNameBuffer OVString,
		ExeNameBufferLength DWORD) DWORD
)

var WinConANSIApis = Apis{
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

var WinConUnicodeApis = Apis{
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

var WinConApis = Apis{
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
