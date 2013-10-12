// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package winnls provides API definitions for accessing
//nls in kernel32.dll.
package winnls

import (
	. "github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/kernel32"
)

var (
	IsValidCodePage func(CodePage T.UINT) T.BOOL

	GetACP func() T.UINT

	GetOEMCP func() T.UINT

	GetCPInfo func(CodePage T.UINT, cpInfo *T.CPINFO) T.BOOL

	GetCPInfoExA func(
		CodePage T.UINT, Flags T.DWORD, cpInfoEx *T.CPINFOEXA) T.BOOL

	GetCPInfoExW func(
		CodePage T.UINT, Flags T.DWORD, cpInfoEx *T.CPINFOEXW) T.BOOL

	IsDBCSLeadByte func(TestChar T.BYTE) T.BOOL

	IsDBCSLeadByteEx func(CodePage T.UINT, TestChar T.BYTE) T.BOOL

	MultiByteToWideChar func(
		CodePage T.UINT,
		Flags T.DWORD,
		MultiByteStr T.AString,
		cbMultiByte int,
		WideCharStr T.OWString,
		WideChar int) int

	WideCharToMultiByte func(
		CodePage T.UINT,
		Flags T.DWORD,
		WideCharStr T.WString,
		WideChar int,
		MultiByteStr T.OAString,
		cbMultiByte int,
		DefaultChar T.AString,
		UsedDefaultChar *T.BOOL) int

	CompareString func(
		Locale T.LCID,
		CmpFlags T.DWORD,
		String1 VString,
		Count1 int,
		String2 VString,
		Count2 int) int

	LCMapString func(
		Locale T.LCID,
		MapFlags T.DWORD,
		SrcStr VString,
		Src int,
		DestStr OVString,
		Dest int) int

	GetLocaleInfo func(
		Locale T.LCID,
		LCType T.LCTYPE,
		LCData OVString,
		Data int) int

	SetLocaleInfo func(
		Locale T.LCID,
		LCType T.LCTYPE,
		LCData VString) T.BOOL

	GetCalendarInfo func(
		Locale T.LCID,
		Calendar T.CALID,
		CalType T.CALTYPE,
		CalData OVString,
		Data int,
		Value *T.DWORD) int

	SetCalendarInfo func(
		Locale T.LCID,
		Calendar T.CALID,
		CalType T.CALTYPE,
		CalData VString) T.BOOL

	GetTimeFormat func(
		Locale T.LCID,
		Flags T.DWORD,
		Time *T.SYSTEMTIME,
		Format VString,
		TimeStr OVString,
		cTime int) int

	GetDateFormat func(
		Locale T.LCID,
		Flags T.DWORD,
		Date *T.SYSTEMTIME,
		Format VString,
		DateStr OVString,
		cDate int) int

	GetNumberFormat func(
		Locale T.LCID,
		Flags T.DWORD,
		Value VString,
		Format *T.NUMBERFMT,
		NumberStr OVString,
		Number int) int

	GetCurrencyFormat func(
		Locale T.LCID,
		Flags T.DWORD,
		Value VString,
		Format *T.CURRENCYFMT,
		CurrencyStr OVString,
		Currency int) int

	EnumCalendarInfo func(
		CalInfoEnumProc T.CALINFO_ENUMPROC,
		Locale T.LCID,
		Calendar T.CALID,
		CalType T.CALTYPE) T.BOOL

	EnumCalendarInfoEx func(
		CalInfoEnumProcEx T.CALINFO_ENUMPROCEX,
		Locale T.LCID,
		Calendar T.CALID,
		CalType T.CALTYPE) T.BOOL

	EnumTimeFormats func(
		TimeFmtEnumProc T.TIMEFMT_ENUMPROC,
		Locale T.LCID,
		Flags T.DWORD) T.BOOL

	EnumDateFormats func(
		DateFmtEnumProc T.DATEFMT_ENUMPROC,
		Locale T.LCID,
		Flags T.DWORD) T.BOOL

	EnumDateFormatsEx func(
		DateFmtEnumProcEx T.DATEFMT_ENUMPROCEX,
		Locale T.LCID,
		Flags T.DWORD) T.BOOL

	IsValidLanguageGroup func(
		LanguageGroup T.LGRPID,
		Flags T.DWORD) T.BOOL

	GetNLSVersion func(
		T.NLS_FUNCTION, T.LCID, *T.NLSVERSIONINFO) T.BOOL

	IsNLSDefinedString func(
		Function T.NLS_FUNCTION,
		Flags T.DWORD,
		VersionInformation *T.NLSVERSIONINFO,
		String T.WString,
		Str T.INT) T.BOOL

	IsValidLocale func(
		Locale T.LCID,
		Flags T.DWORD) T.BOOL

	GetGeoInfo func(
		Location T.GEOID,
		GeoType T.GEOTYPE,
		GeoData OVString,
		Data int,
		LangId T.LANGID) int

	EnumSystemGeoID func(
		GeoClass T.GEOCLASS,
		ParentGeoId T.GEOID,
		GeoEnumProc T.GEO_ENUMPROC) T.BOOL

	GetUserGeoID func(GeoClass T.GEOCLASS) T.GEOID

	SetUserGeoID func(GeoId T.GEOID) T.BOOL

	ConvertDefaultLocale func(Locale T.LCID) T.LCID

	GetThreadLocale func() T.LCID

	SetThreadLocale func(Locale T.LCID) T.BOOL

	GetSystemDefaultUILanguage func() T.LANGID

	GetUserDefaultUILanguage func() T.LANGID

	GetSystemDefaultLangID func() T.LANGID

	GetUserDefaultLangID func() T.LANGID

	GetSystemDefaultLCID func() T.LCID

	GetUserDefaultLCID func() T.LCID

	GetStringTypeEx func(
		Locale T.LCID,
		InfoType T.DWORD,
		SrcStr VString,
		Src int,
		CharType *T.WORD) T.BOOL

	GetStringType func(
		Locale T.LCID,
		InfoType T.DWORD,
		SrcStr VString,
		Src int,
		CharType *T.WORD) T.BOOL

	FoldString func(
		MapFlags T.DWORD,
		SrcStr VString,
		Src int,
		DestStr OVString,
		Dest int) int

	EnumSystemLanguageGroups func(
		LanguageGroupEnumProc T.LANGUAGEGROUP_ENUMPROC,
		Flags T.DWORD,
		lParam T.LONG_PTR) T.BOOL

	EnumLanguageGroupLocales func(
		LangGroupLocaleEnumProc T.LANGGROUPLOCALE_ENUMPROC,
		LanguageGroup T.LGRPID,
		Flags T.DWORD,
		lParam T.LONG_PTR) T.BOOL

	EnumUILanguages func(
		UILanguageEnumProc T.UILANGUAGE_ENUMPROC,
		Flags T.DWORD,
		lParam T.LONG_PTR) T.BOOL

	EnumSystemLocales func(
		LocaleEnumProc T.LOCALE_ENUMPROC,
		Flags T.DWORD) T.BOOL

	EnumSystemCodePages func(
		CodePageEnumProc T.CODEPAGE_ENUMPROC,
		Flags T.DWORD) T.BOOL
)

//TODO(t): not on XP
//{"GetNLSVersion", &GetNLSVersion},
//{"IsNLSDefinedString", &IsNLSDefinedString},

var WinNlsANSIApis = Apis{
	{"GetCPInfoExA", &GetCPInfoExA},
	{"LCMapStringA", &LCMapString},
	{"GetLocaleInfoA", &GetLocaleInfo},
	{"SetLocaleInfoA", &SetLocaleInfo},
	{"GetCalendarInfoA", &GetCalendarInfo},
	{"SetCalendarInfoA", &SetCalendarInfo},
	{"GetTimeFormatA", &GetTimeFormat},
	{"GetDateFormatA", &GetDateFormat},
	{"CompareStringA", &CompareString},
	{"GetNumberFormatA", &GetNumberFormat},
	{"GetCurrencyFormatA", &GetCurrencyFormat},
	{"EnumCalendarInfoA", &EnumCalendarInfo},
	{"EnumCalendarInfoExA", &EnumCalendarInfoEx},
	{"EnumTimeFormatsA", &EnumTimeFormats},
	{"EnumDateFormatsA", &EnumDateFormats},
	{"EnumDateFormatsExA", &EnumDateFormatsEx},
	{"GetGeoInfoA", &GetGeoInfo},
	{"GetStringTypeExA", &GetStringTypeEx},
	{"GetStringTypeA", &GetStringType},
	{"FoldStringA", &FoldString},
	{"EnumSystemLanguageGroupsA", &EnumSystemLanguageGroups},
	{"EnumLanguageGroupLocalesA", &EnumLanguageGroupLocales},
	{"EnumUILanguagesA", &EnumUILanguages},
	{"EnumSystemLocalesA", &EnumSystemLocales},
	{"EnumSystemCodePagesA", &EnumSystemCodePages},
}

var WinNlsUnicodeApis = Apis{
	{"GetCPInfoExW", &GetCPInfoExW},
	{"LCMapStringW", &LCMapString},
	{"GetLocaleInfoW", &GetLocaleInfo},
	{"SetLocaleInfoW", &SetLocaleInfo},
	{"GetCalendarInfoW", &GetCalendarInfo},
	{"SetCalendarInfoW", &SetCalendarInfo},
	{"GetTimeFormatW", &GetTimeFormat},
	{"GetDateFormatW", &GetDateFormat},
	{"CompareStringW", &CompareString},
	{"GetNumberFormatW", &GetNumberFormat},
	{"GetCurrencyFormatW", &GetCurrencyFormat},
	{"EnumCalendarInfoW", &EnumCalendarInfo},
	{"EnumCalendarInfoExW", &EnumCalendarInfoEx},
	{"EnumTimeFormatsW", &EnumTimeFormats},
	{"EnumDateFormatsW", &EnumDateFormats},
	{"EnumDateFormatsExW", &EnumDateFormatsEx},
	{"GetGeoInfoW", &GetGeoInfo},
	{"GetStringTypeExW", &GetStringTypeEx},
	{"GetStringTypeW", &GetStringType},
	{"FoldStringW", &FoldString},
	{"EnumSystemLanguageGroupsW", &EnumSystemLanguageGroups},
	{"EnumLanguageGroupLocalesW", &EnumLanguageGroupLocales},
	{"EnumUILanguagesW", &EnumUILanguages},
	{"EnumSystemLocalesW", &EnumSystemLocales},
	{"EnumSystemCodePagesW", &EnumSystemCodePages},
}

var WinNlsApis = Apis{
	{"IsValidCodePage", &IsValidCodePage},
	{"GetACP", &GetACP},
	{"GetOEMCP", &GetOEMCP},
	{"GetCPInfo", &GetCPInfo},
	{"IsDBCSLeadByte", &IsDBCSLeadByte},
	{"IsDBCSLeadByteEx", &IsDBCSLeadByteEx},
	{"MultiByteToWideChar", &MultiByteToWideChar},
	{"WideCharToMultiByte", &WideCharToMultiByte},
	{"IsValidLanguageGroup", &IsValidLanguageGroup},
	{"IsValidLocale", &IsValidLocale},
	{"EnumSystemGeoID", &EnumSystemGeoID},
	{"GetUserGeoID", &GetUserGeoID},
	{"SetUserGeoID", &SetUserGeoID},
	{"ConvertDefaultLocale", &ConvertDefaultLocale},
	{"GetThreadLocale", &GetThreadLocale},
	{"SetThreadLocale", &SetThreadLocale},
	{"GetSystemDefaultUILanguage", &GetSystemDefaultUILanguage},
	{"GetUserDefaultUILanguage", &GetUserDefaultUILanguage},
	{"GetSystemDefaultLangID", &GetSystemDefaultLangID},
	{"GetUserDefaultLangID", &GetUserDefaultLangID},
	{"GetSystemDefaultLCID", &GetSystemDefaultLCID},
	{"GetUserDefaultLCID", &GetUserDefaultLCID},
}
