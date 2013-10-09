// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package shellapi provides API definitions for accessing
//shell32.dll.
package shellapi

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/shell32"
)

var (
	DragQueryFile func(HDROP, UINT, VString, UINT) UINT

	DragQueryPoint func(HDROP, *POINT) BOOL

	DragFinish func(HDROP)

	DragAcceptFiles func(HWND, BOOL)

	ShellExecute func(
		Wnd HWND,
		Operation,
		File, Parameters,
		Directory VString,
		ShowCmd INT) HINSTANCE

	FindExecutable func(
		File, Directory VString,
		Result OVString) HINSTANCE

	CommandLineToArgvW func(
		CmdLine WString, NumArgs *int) *WString
	//TODO(t):*WString was *LPWSTR

	ShellAbout func(
		Wnd HWND,
		App, OtherStuff VString,
		Icon HICON) INT

	DuplicateIcon func(Inst HINSTANCE, Icon HICON) HICON

	ExtractAssociatedIcon func(
		Inst HINSTANCE, IconPath VString, Icon *WORD) HICON

	ExtractAssociatedIconEx func(
		Inst HINSTANCE,
		IconPath VString,
		IconIndex *WORD,
		IconId *WORD) HICON

	ExtractIcon func(
		Inst HINSTANCE,
		ExeFileName VString,
		IconIndex UINT) HICON

	ExtractIconEx func(
		File VString,
		IconIndex int,
		Large, Small *HICON,
		Icons UINT) UINT

	SHFreeNameMappings func(hNameMappings HANDLE)

	WinExecError func(
		Wnd HWND,
		Error int,
		FileName,
		Title VString)

	SHCreateProcessAsUserW func(
		Scpi *SHCREATEPROCESSINFOW) BOOL

	SHGetFileInfoA func(
		Path AString,
		FileAttributes DWORD,
		Sfi *SHFILEINFOA,
		FileInfo, Flags UINT) DWORD_PTR

	SHGetFileInfoW func(
		Path WString,
		FileAttributes DWORD,
		Sfi *SHFILEINFOW,
		FileInfo, Flags UINT) DWORD_PTR

	SHGetDiskFreeSpaceEx func(
		DirectoryName VString,
		FreeBytesAvailableToCaller,
		TotalNumberOfBytes,
		TotalNumberOfFreeBytes *ULARGE_INTEGER) BOOL

	SHGetNewLinkInfo func(
		LinkTo,
		Dir VString,
		Name OVString,
		MustCopy *BOOL,
		Flags UINT) BOOL

	SHInvokePrinterCommand func(
		Wnd HWND,
		Action UINT,
		Buf1, Buf2 VString,
		Modal BOOL) BOOL

	IsLFNDrive func(Path VString) BOOL

	SHEnumerateUnreadMailAccounts func(
		KeyUser HKEY,
		Index DWORD,
		MailAddress OVString,
		cMailAddress int)

	SHGetUnreadMailCount func(
		KeyUser HKEY,
		MailAddress VString,
		Count *DWORD,
		FileTime *FILETIME,
		ShellExecuteCommand OVString,
		cShellExecuteCommand int)

	SHSetUnreadMailCount func(
		MailAddress VString,
		Count DWORD,
		ShellExecuteCommand AString)

	SHTestTokenMembership func(
		Token HANDLE, RID ULONG) BOOL

	SHGetImageList func(
		ImageList int, Riid REFIID, Obj **VOID)
)

var ShellApiANSIApis = Apis{
	{"ExtractAssociatedIconA", &ExtractAssociatedIcon},
	{"ExtractAssociatedIconExA", &ExtractAssociatedIconEx},
	{"ExtractIconA", &ExtractIcon},
	{"ExtractIconExA", &ExtractIconEx},
	{"FindExecutableA", &FindExecutable},
	{"ShellAboutA", &ShellAbout},
	{"ShellExecuteA", &ShellExecute},
	{"SHGetDiskFreeSpaceExA", &SHGetDiskFreeSpaceEx},
	{"SHGetFileInfoA", &SHGetFileInfoA},
	{"SHInvokePrinterCommandA", &SHInvokePrinterCommand},
}

var ShellApiUnicodeApis = Apis{
	{"ExtractAssociatedIconW", &ExtractAssociatedIcon},
	{"ExtractAssociatedIconExW", &ExtractAssociatedIconEx},
	{"ExtractIconW", &ExtractIcon},
	{"ExtractIconExW", &ExtractIconEx},
	{"FindExecutableW", &FindExecutable},
	{"ShellAboutW", &ShellAbout},
	{"ShellExecuteW", &ShellExecute},
	{"SHGetDiskFreeSpaceExW", &SHGetDiskFreeSpaceEx},
	{"SHGetFileInfoW", &SHGetFileInfoW},
	{"SHInvokePrinterCommandW", &SHInvokePrinterCommand},
}

//NOTR(t):Not implemented
//{"WinExecErrorA", &WinExecError},
//{"WinExecErrorW", &WinExecError},

//TODO(t): Not on XP
//{"SHEnumerateUnreadMailAccounts", &SHEnumerateUnreadMailAccounts},
//{"SHGetUnreadMailCount", &SHGetUnreadMailCount},
//{"SHSetUnreadMailCount", &SHSetUnreadMailCount},

var ShellApiApis = Apis{
	{"CommandLineToArgvW", &CommandLineToArgvW},
	{"DragAcceptFiles", &DragAcceptFiles},
	{"DragFinish", &DragFinish},
	{"DragQueryFile", &DragQueryFile},
	{"DragQueryPoint", &DragQueryPoint},
	{"DuplicateIcon", &DuplicateIcon},
	{"IsLFNDrive", &IsLFNDrive},
	{"SHCreateProcessAsUserW", &SHCreateProcessAsUserW},
	{"SHFreeNameMappings", &SHFreeNameMappings},
	{"SHGetImageList", &SHGetImageList},
	{"SHGetNewLinkInfo", &SHGetNewLinkInfo},
	{"SHTestTokenMembership", &SHTestTokenMembership},
}
