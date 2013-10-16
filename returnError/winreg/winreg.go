// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package winreg provides API definitions for accessing
//advapi32.dll.
package winreg

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	_ "github.com/tHinqa/outside/win32/advapi32"
)

var (
	RegCloseKey func(Key T.HKEY) (T.LONG, error)

	RegOverridePredefKey func(Key T.HKEY, NewHKey T.HKEY) (T.LONG, error)

	RegOpenUserClassesRoot func(
		Token T.HANDLE,
		Options T.DWORD,
		SamDesired T.REGSAM,
		Result *T.HKEY) (T.LONG, error)

	RegOpenCurrentUser func(
		SamDesired T.REGSAM, Result *T.HKEY) (T.LONG, error)

	RegDisablePredefinedCache func() (T.LONG, error)

	RegConnectRegistry func(
		MachineName VString,
		Key T.HKEY,
		Result *T.HKEY) (T.LONG, error)

	RegConnectRegistryEx func(
		MachineName VString,
		Key T.HKEY,
		Flags T.ULONG,
		Result *T.HKEY) (T.LONG, error)

	RegCreateKey func(
		Key T.HKEY, SubKey VString, Result *T.HKEY) (T.LONG, error)

	RegCreateKeyEx func(
		Key T.HKEY,
		SubKey VString,
		Reserved T.DWORD,
		Class VString,
		Options T.DWORD,
		SamDesired T.REGSAM,
		SecurityAttributes *T.SECURITY_ATTRIBUTES,
		Result *T.HKEY,
		Disposition *T.DWORD) (T.LONG, error)

	RegDeleteKey func(Key T.HKEY, SubKey VString) (T.LONG, error)

	RegDeleteKeyEx func(
		Key T.HKEY,
		SubKey VString,
		SamDesired T.REGSAM,
		Reserved T.DWORD) (T.LONG, error)

	RegDisableReflectionKey func(Base T.HKEY) (T.LONG, error)

	RegEnableReflectionKey func(Base T.HKEY) (T.LONG, error)

	RegQueryReflectionKey func(
		Base T.HKEY, IsReflectionDisabled *T.BOOL) (T.LONG, error)

	RegDeleteValue func(Key T.HKEY, ValueName VString) (T.LONG, error)

	RegEnumKey func(
		Key T.HKEY,
		Index T.DWORD,
		Name OVString,
		cName T.DWORD) (T.LONG, error)

	RegEnumKeyEx func(
		Key T.HKEY,
		Index T.DWORD,
		Name OVString,
		cName *T.DWORD,
		Reserved *T.DWORD,
		Class OVString,
		cClass *T.DWORD,
		LastWriteTime *T.FILETIME) (T.LONG, error)

	RegEnumValue func(
		Key T.HKEY,
		Index T.DWORD,
		ValueName OVString,
		cValueName *T.DWORD,
		Reserved *T.DWORD,
		Type *T.DWORD,
		Data *T.BYTE,
		cData *T.DWORD) (T.LONG, error)

	RegFlushKey func(Key T.HKEY) (T.LONG, error)

	RegGetKeySecurity func(
		Key T.HKEY,
		SecurityInformation T.SECURITY_INFORMATION,
		SecurityDescriptor *T.SECURITY_DESCRIPTOR,
		cSecurityDescriptor *T.DWORD) (T.LONG, error)

	RegLoadKey func(
		Key T.HKEY, SubKey, File VString) (T.LONG, error)

	RegNotifyChangeKeyValue func(
		Key T.HKEY,
		WatchSubtree T.BOOL,
		NotifyFilter T.DWORD,
		Event T.HANDLE,
		Asynchronous T.BOOL) (T.LONG, error)

	RegOpenKey func(
		Key T.HKEY, SubKey VString, Result *T.HKEY) (T.LONG, error)

	RegOpenKeyEx func(
		Key T.HKEY,
		SubKey VString,
		Options T.DWORD,
		SamDesired T.REGSAM,
		Result *T.HKEY) (T.LONG, error)

	RegQueryInfoKey func(
		Key T.HKEY,
		Class OVString,
		cClass,
		Reserved,
		SubKeys,
		MaxSubKeyLen,
		MaxClassLen,
		Values,
		MaxValueNameLen,
		MaxValueLen,
		SecurityDescriptor *T.DWORD,
		LastWriteTime *T.FILETIME) (T.LONG, error)

	RegQueryValue func(
		Key T.HKEY,
		SubKey VString,
		Data OVString,
		cData *T.LONG) (T.LONG, error)

	RegQueryMultipleValues func(
		Key T.HKEY,
		ValList *T.VALENT,
		NumVals T.DWORD,
		ValueBuf VString,
		TotSize *T.DWORD) (T.LONG, error)

	RegQueryValueEx func(
		Key T.HKEY,
		ValueName VString,
		Reserved,
		Type *T.DWORD,
		Data *T.BYTE,
		cData *T.DWORD) (T.LONG, error)

	RegReplaceKey func(
		Key T.HKEY, SubKey, NewFile, OldFile VString) (T.LONG, error)

	RegRestoreKey func(
		Key T.HKEY, File VString, Flags T.DWORD) (T.LONG, error)

	RegSaveKey func(
		Key T.HKEY,
		File VString,
		SecurityAttributes *T.SECURITY_ATTRIBUTES) (T.LONG, error)

	RegSetKeySecurity func(
		Key T.HKEY,
		SecurityInformation T.SECURITY_INFORMATION,
		pSecurityDescriptor *T.SECURITY_DESCRIPTOR) (T.LONG, error)

	RegSetValue func(
		Key T.HKEY,
		SubKey VString,
		Type T.DWORD,
		Data VString,
		cbData T.DWORD) (T.LONG, error)

	RegSetValueEx func(
		Key T.HKEY,
		ValueName VString,
		Reserved,
		Type T.DWORD,
		Data *T.BYTE,
		cData T.DWORD) (T.LONG, error)

	RegUnLoadKey func(Key T.HKEY, SubKey VString) (T.LONG, error)

	RegGetValue func(
		Key T.HKEY,
		SubKey, Value VString,
		Flags T.DWORD,
		Type *T.DWORD,
		Data *T.VOID,
		cData *T.DWORD) (T.LONG, error)

	InitiateSystemShutdown func(
		MachineName, Message VString,
		Timeout T.DWORD,
		ForceAppsClosed, RebootAfterShutdown T.BOOL) (T.BOOL, error)

	AbortSystemShutdown func(
		MachineName VString) (T.BOOL, error)

	InitiateSystemShutdownEx func(
		MachineName, Message VString,
		Timeout T.DWORD,
		ForceAppsClosed, RebootAfterShutdown T.BOOL,
		Reason T.DWORD) (T.BOOL, error)

	RegSaveKeyEx func(
		Key T.HKEY,
		File VString,
		SecurityAttributes *T.SECURITY_ATTRIBUTES,
		Flags T.DWORD) (T.LONG, error)

	Wow64Win32ApiEntry func(
		FuncNumber T.DWORD, Flag T.DWORD, Res T.DWORD) (T.LONG, error)
)

var WinRegANSIApis = outside.Apis{
	{"AbortSystemShutdownA", &AbortSystemShutdown},
	{"InitiateSystemShutdownA", &InitiateSystemShutdown},
	{"InitiateSystemShutdownExA", &InitiateSystemShutdownEx},
	{"RegConnectRegistryA", &RegConnectRegistry},
	{"RegCreateKeyA", &RegCreateKey},
	{"RegCreateKeyExA", &RegCreateKeyEx},
	{"RegDeleteKeyA", &RegDeleteKey},
	{"RegDeleteValueA", &RegDeleteValue},
	{"RegEnumKeyA", &RegEnumKey},
	{"RegEnumKeyExA", &RegEnumKeyEx},
	{"RegEnumValueA", &RegEnumValue},
	{"RegLoadKeyA", &RegLoadKey},
	{"RegOpenKeyA", &RegOpenKey},
	{"RegOpenKeyExA", &RegOpenKeyEx},
	{"RegQueryInfoKeyA", &RegQueryInfoKey},
	{"RegQueryMultipleValuesA", &RegQueryMultipleValues},
	{"RegQueryValueA", &RegQueryValue},
	{"RegQueryValueExA", &RegQueryValueEx},
	{"RegReplaceKeyA", &RegReplaceKey},
	{"RegRestoreKeyA", &RegRestoreKey},
	{"RegSaveKeyA", &RegSaveKey},
	{"RegSaveKeyExA", &RegSaveKeyEx},
	{"RegSetValueA", &RegSetValue},
	{"RegSetValueExA", &RegSetValueEx},
	{"RegUnLoadKeyA", &RegUnLoadKey},
}

var WinRegUnicodeApis = outside.Apis{
	{"AbortSystemShutdownW", &AbortSystemShutdown},
	{"InitiateSystemShutdownW", &InitiateSystemShutdown},
	{"InitiateSystemShutdownExW", &InitiateSystemShutdownEx},
	{"RegConnectRegistryW", &RegConnectRegistry},
	{"RegCreateKeyW", &RegCreateKey},
	{"RegCreateKeyExW", &RegCreateKeyEx},
	{"RegDeleteKeyW", &RegDeleteKey},
	{"RegDeleteValueW", &RegDeleteValue},
	{"RegEnumKeyW", &RegEnumKey},
	{"RegEnumKeyExW", &RegEnumKeyEx},
	{"RegEnumValueW", &RegEnumValue},
	{"RegLoadKeyW", &RegLoadKey},
	{"RegOpenKeyW", &RegOpenKey},
	{"RegOpenKeyExW", &RegOpenKeyEx},
	{"RegQueryInfoKeyW", &RegQueryInfoKey},
	{"RegQueryMultipleValuesW", &RegQueryMultipleValues},
	{"RegQueryValueW", &RegQueryValue},
	{"RegQueryValueExW", &RegQueryValueEx},
	{"RegReplaceKeyW", &RegReplaceKey},
	{"RegRestoreKeyW", &RegRestoreKey},
	{"RegSaveKeyW", &RegSaveKey},
	{"RegSaveKeyExW", &RegSaveKeyEx},
	{"RegSetValueW", &RegSetValue},
	{"RegSetValueExW", &RegSetValueEx},
	{"RegUnLoadKeyW", &RegUnLoadKey},
}

//TODO(t):Not on XP.
//{"RegConnectRegistryExA", &RegConnectRegistryEx},
//{"RegConnectRegistryExW", &RegConnectRegistryEx},
//{"RegDeleteKeyExA", &RegDeleteKeyEx},
//{"RegDeleteKeyExW", &RegDeleteKeyEx},
//{"RegDisableReflectionKey", &RegDisableReflectionKey},
//{"RegEnableReflectionKey", &RegEnableReflectionKey},
//{"RegGetValueA", &RegGetValue},
//{"RegGetValueW", &RegGetValue},
//{"RegQueryReflectionKey", &RegQueryReflectionKey},

var WinRegApis = outside.Apis{
	{"RegCloseKey", &RegCloseKey},
	{"RegDisablePredefinedCache", &RegDisablePredefinedCache},
	{"RegFlushKey", &RegFlushKey},
	{"RegGetKeySecurity", &RegGetKeySecurity},
	{"RegNotifyChangeKeyValue", &RegNotifyChangeKeyValue},
	{"RegOpenCurrentUser", &RegOpenCurrentUser},
	{"RegOpenUserClassesRoot", &RegOpenUserClassesRoot},
	{"RegOverridePredefKey", &RegOverridePredefKey},
	{"RegSetKeySecurity", &RegSetKeySecurity},
	{"Wow64Win32ApiEntry", &Wow64Win32ApiEntry},
}
