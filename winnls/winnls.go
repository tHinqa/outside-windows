// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package winnls provides API definitions for accessing
//nls in kernel32.dll.
package winnls

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	_ "github.com/tHinqa/outside/win32/kernel32"
)

var (
	IsValidCodePage func(CodePage T.UINT) (T.BOOL, error)

	GetACP func() (T.UINT, error)

	GetOEMCP func() (T.UINT, error)

	GetCPInfo func(CodePage T.UINT, cpInfo *T.CPINFO) (T.BOOL, error)

	GetCPInfoExA func(
		CodePage T.UINT, Flags T.DWORD, cpInfoEx *T.CPINFOEXA) (T.BOOL, error)

	GetCPInfoExW func(
		CodePage T.UINT, Flags T.DWORD, cpInfoEx *T.CPINFOEXW) (T.BOOL, error)

	IsDBCSLeadByte func(TestChar T.BYTE) (T.BOOL, error)

	IsDBCSLeadByteEx func(CodePage T.UINT, TestChar T.BYTE) (T.BOOL, error)

	MultiByteToWideChar func(
		CodePage T.UINT,
		Flags T.DWORD,
		MultiByteStr T.AString,
		cbMultiByte int,
		WideCharStr T.OWString,
		WideChar int) (int, error)

	WideCharToMultiByte func(
		CodePage T.UINT,
		Flags T.DWORD,
		WideCharStr T.WString,
		WideChar int,
		MultiByteStr T.OAString,
		cbMultiByte int,
		DefaultChar T.AString,
		UsedDefaultChar *T.BOOL) (int, error)

	CompareString func(
		Locale T.LCID,
		CmpFlags T.DWORD,
		String1 VString,
		Count1 int,
		String2 VString,
		Count2 int) (int, error)

	LCMapString func(
		Locale T.LCID,
		MapFlags T.DWORD,
		SrcStr VString,
		Src int,
		DestStr OVString,
		Dest int) (int, error)

	GetLocaleInfo func(
		Locale T.LCID,
		LCType T.LCTYPE,
		LCData OVString,
		Data int) (int, error)

	SetLocaleInfo func(
		Locale T.LCID,
		LCType T.LCTYPE,
		LCData VString) (T.BOOL, error)

	GetCalendarInfo func(
		Locale T.LCID,
		Calendar T.CALID,
		CalType T.CALTYPE,
		CalData OVString,
		Data int,
		Value *T.DWORD) (int, error)

	SetCalendarInfo func(
		Locale T.LCID,
		Calendar T.CALID,
		CalType T.CALTYPE,
		CalData VString) (T.BOOL, error)

	GetTimeFormat func(
		Locale T.LCID,
		Flags T.DWORD,
		Time *T.SYSTEMTIME,
		Format VString,
		TimeStr OVString,
		cTime int) (int, error)

	GetDateFormat func(
		Locale T.LCID,
		Flags T.DWORD,
		Date *T.SYSTEMTIME,
		Format VString,
		DateStr OVString,
		cDate int) (int, error)

	GetNumberFormat func(
		Locale T.LCID,
		Flags T.DWORD,
		Value VString,
		Format *T.NUMBERFMT,
		NumberStr OVString,
		Number int) (int, error)

	GetCurrencyFormat func(
		Locale T.LCID,
		Flags T.DWORD,
		Value VString,
		Format *T.CURRENCYFMT,
		CurrencyStr OVString,
		Currency int) (int, error)

	EnumCalendarInfo func(
		CalInfoEnumProc T.CALINFO_ENUMPROC,
		Locale T.LCID,
		Calendar T.CALID,
		CalType T.CALTYPE) (T.BOOL, error)

	EnumCalendarInfoEx func(
		CalInfoEnumProcEx T.CALINFO_ENUMPROCEX,
		Locale T.LCID,
		Calendar T.CALID,
		CalType T.CALTYPE) (T.BOOL, error)

	EnumTimeFormats func(
		TimeFmtEnumProc T.TIMEFMT_ENUMPROC,
		Locale T.LCID,
		Flags T.DWORD) (T.BOOL, error)

	EnumDateFormats func(
		DateFmtEnumProc T.DATEFMT_ENUMPROC,
		Locale T.LCID,
		Flags T.DWORD) (T.BOOL, error)

	EnumDateFormatsEx func(
		DateFmtEnumProcEx T.DATEFMT_ENUMPROCEX,
		Locale T.LCID,
		Flags T.DWORD) (T.BOOL, error)

	IsValidLanguageGroup func(
		LanguageGroup T.LGRPID,
		Flags T.DWORD) (T.BOOL, error)

	GetNLSVersion func(
		T.NLS_FUNCTION, T.LCID, *T.NLSVERSIONINFO) (T.BOOL, error)

	IsNLSDefinedString func(
		Function T.NLS_FUNCTION,
		Flags T.DWORD,
		VersionInformation *T.NLSVERSIONINFO,
		String T.WString,
		Str T.INT) (T.BOOL, error)

	IsValidLocale func(
		Locale T.LCID,
		Flags T.DWORD) (T.BOOL, error)

	GetGeoInfo func(
		Location T.GEOID,
		GeoType T.GEOTYPE,
		GeoData OVString,
		Data int,
		LangId T.LANGID) (int, error)

	EnumSystemGeoID func(
		GeoClass T.GEOCLASS,
		ParentGeoId T.GEOID,
		GeoEnumProc T.GEO_ENUMPROC) (T.BOOL, error)

	GetUserGeoID func(GeoClass T.GEOCLASS) (T.GEOID, error)

	SetUserGeoID func(GeoId T.GEOID) (T.BOOL, error)

	ConvertDefaultLocale func(Locale T.LCID) (T.LCID, error)

	GetThreadLocale func() (T.LCID, error)

	SetThreadLocale func(Locale T.LCID) (T.BOOL, error)

	GetSystemDefaultUILanguage func() (T.LANGID, error)

	GetUserDefaultUILanguage func() (T.LANGID, error)

	GetSystemDefaultLangID func() (T.LANGID, error)

	GetUserDefaultLangID func() (T.LANGID, error)

	GetSystemDefaultLCID func() (T.LCID, error)

	GetUserDefaultLCID func() (T.LCID, error)

	GetStringTypeEx func(
		Locale T.LCID,
		InfoType T.DWORD,
		SrcStr VString,
		Src int,
		CharType *T.WORD) (T.BOOL, error)

	GetStringType func(
		Locale T.LCID,
		InfoType T.DWORD,
		SrcStr VString,
		Src int,
		CharType *T.WORD) (T.BOOL, error)

	FoldString func(
		MapFlags T.DWORD,
		SrcStr VString,
		Src int,
		DestStr OVString,
		Dest int) (int, error)

	EnumSystemLanguageGroups func(
		LanguageGroupEnumProc T.LANGUAGEGROUP_ENUMPROC,
		Flags T.DWORD,
		lParam T.LONG_PTR) (T.BOOL, error)

	EnumLanguageGroupLocales func(
		LangGroupLocaleEnumProc T.LANGGROUPLOCALE_ENUMPROC,
		LanguageGroup T.LGRPID,
		Flags T.DWORD,
		lParam T.LONG_PTR) (T.BOOL, error)

	EnumUILanguages func(
		UILanguageEnumProc T.UILANGUAGE_ENUMPROC,
		Flags T.DWORD,
		lParam T.LONG_PTR) (T.BOOL, error)

	EnumSystemLocales func(
		LocaleEnumProc T.LOCALE_ENUMPROC,
		Flags T.DWORD) (T.BOOL, error)

	EnumSystemCodePages func(
		CodePageEnumProc T.CODEPAGE_ENUMPROC,
		Flags T.DWORD) (T.BOOL, error)
)

//TODO(t): not on XP
//{"GetNLSVersion", &GetNLSVersion},
//{"IsNLSDefinedString", &IsNLSDefinedString},

var WinNlsANSIApis = outside.Apis{
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

var WinNlsUnicodeApis = outside.Apis{
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

var WinNlsApis = outside.Apis{
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
