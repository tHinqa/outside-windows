// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package winnls provides API definitions for accessing
//nls in kernel32.dll.
package winnls

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/kernel32"
)

var (
	IsValidCodePage func(CodePage UINT) BOOL

	GetACP func() UINT

	GetOEMCP func() UINT

	GetCPInfo func(CodePage UINT, CPInfo *CPINFO) BOOL

	GetCPInfoExA func(
		CodePage UINT, Flags DWORD, CPInfoEx *CPINFOEXA) BOOL

	GetCPInfoExW func(
		CodePage UINT, Flags DWORD, CPInfoEx *CPINFOEXW) BOOL

	IsDBCSLeadByte func(TestChar BYTE) BOOL

	IsDBCSLeadByteEx func(CodePage UINT, TestChar BYTE) BOOL

	MultiByteToWideChar func(
		CodePage UINT,
		Flags DWORD,
		MultiByteStr AString,
		cbMultiByte int,
		WideCharStr OWString,
		WideChar int) int

	WideCharToMultiByte func(
		CodePage UINT,
		Flags DWORD,
		WideCharStr WString,
		WideChar int,
		MultiByteStr OAString,
		cbMultiByte int,
		DefaultChar AString,
		UsedDefaultChar *BOOL) int

	CompareString func(
		Locale LCID,
		CmpFlags DWORD,
		String1 VString,
		Count1 int,
		String2 VString,
		Count2 int) int

	LCMapString func(
		Locale LCID,
		MapFlags DWORD,
		SrcStr VString,
		Src int,
		DestStr OVString,
		Dest int) int

	GetLocaleInfo func(
		Locale LCID,
		LCType LCTYPE,
		LCData OVString,
		Data int) int

	SetLocaleInfo func(
		Locale LCID,
		LCType LCTYPE,
		LCData VString) BOOL

	GetCalendarInfo func(
		Locale LCID,
		Calendar CALID,
		CalType CALTYPE,
		CalData OVString,
		Data int,
		Value *DWORD) int

	SetCalendarInfo func(
		Locale LCID,
		Calendar CALID,
		CalType CALTYPE,
		CalData VString) BOOL

	GetTimeFormat func(
		Locale LCID,
		Flags DWORD,
		Time *SYSTEMTIME,
		Format VString,
		TimeStr OVString,
		cTime int) int

	GetDateFormat func(
		Locale LCID,
		Flags DWORD,
		Date *SYSTEMTIME,
		Format VString,
		DateStr OVString,
		cDate int) int

	GetNumberFormat func(
		Locale LCID,
		Flags DWORD,
		Value VString,
		Format *NUMBERFMT,
		NumberStr OVString,
		Number int) int

	GetCurrencyFormat func(
		Locale LCID,
		Flags DWORD,
		Value VString,
		Format *CURRENCYFMT,
		CurrencyStr OVString,
		Currency int) int

	EnumCalendarInfo func(
		CalInfoEnumProc CALINFO_ENUMPROC,
		Locale LCID,
		Calendar CALID,
		CalType CALTYPE) BOOL

	EnumCalendarInfoEx func(
		CalInfoEnumProcEx CALINFO_ENUMPROCEX,
		Locale LCID,
		Calendar CALID,
		CalType CALTYPE) BOOL

	EnumTimeFormats func(
		TimeFmtEnumProc TIMEFMT_ENUMPROC,
		Locale LCID,
		Flags DWORD) BOOL

	EnumDateFormats func(
		DateFmtEnumProc DATEFMT_ENUMPROC,
		Locale LCID,
		Flags DWORD) BOOL

	EnumDateFormatsEx func(
		DateFmtEnumProcEx DATEFMT_ENUMPROCEX,
		Locale LCID,
		Flags DWORD) BOOL

	IsValidLanguageGroup func(
		LanguageGroup LGRPID,
		Flags DWORD) BOOL

	GetNLSVersion func(
		NLS_FUNCTION, LCID, *NLSVERSIONINFO) BOOL

	IsNLSDefinedString func(
		Function NLS_FUNCTION,
		Flags DWORD,
		VersionInformation *NLSVERSIONINFO,
		String WString,
		Str INT) BOOL

	IsValidLocale func(
		Locale LCID,
		Flags DWORD) BOOL

	GetGeoInfo func(
		Location GEOID,
		GeoType GEOTYPE,
		GeoData OVString,
		Data int,
		LangId LANGID) int

	EnumSystemGeoID func(
		GeoClass GEOCLASS,
		ParentGeoId GEOID,
		GeoEnumProc GEO_ENUMPROC) BOOL

	GetUserGeoID func(GeoClass GEOCLASS) GEOID

	SetUserGeoID func(GeoId GEOID) BOOL

	ConvertDefaultLocale func(Locale LCID) LCID

	GetThreadLocale func() LCID

	SetThreadLocale func(Locale LCID) BOOL

	GetSystemDefaultUILanguage func() LANGID

	GetUserDefaultUILanguage func() LANGID

	GetSystemDefaultLangID func() LANGID

	GetUserDefaultLangID func() LANGID

	GetSystemDefaultLCID func() LCID

	GetUserDefaultLCID func() LCID

	GetStringTypeEx func(
		Locale LCID,
		InfoType DWORD,
		SrcStr VString,
		Src int,
		CharType *WORD) BOOL

	GetStringType func(
		Locale LCID,
		InfoType DWORD,
		SrcStr VString,
		Src int,
		CharType *WORD) BOOL

	FoldString func(
		MapFlags DWORD,
		SrcStr VString,
		Src int,
		DestStr OVString,
		Dest int) int

	EnumSystemLanguageGroups func(
		LanguageGroupEnumProc LANGUAGEGROUP_ENUMPROC,
		Flags DWORD,
		lParam LONG_PTR) BOOL

	EnumLanguageGroupLocales func(
		LangGroupLocaleEnumProc LANGGROUPLOCALE_ENUMPROC,
		LanguageGroup LGRPID,
		Flags DWORD,
		lParam LONG_PTR) BOOL

	EnumUILanguages func(
		UILanguageEnumProc UILANGUAGE_ENUMPROC,
		Flags DWORD,
		lParam LONG_PTR) BOOL

	EnumSystemLocales func(
		LocaleEnumProc LOCALE_ENUMPROC,
		Flags DWORD) BOOL

	EnumSystemCodePages func(
		CodePageEnumProc CODEPAGE_ENUMPROC,
		Flags DWORD) BOOL
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
