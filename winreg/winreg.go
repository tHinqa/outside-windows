package winreg

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/advapi32"
)

var (
	RegCloseKey func(Key HKEY) LONG

	RegOverridePredefKey func(Key HKEY, NewHKey HKEY) LONG

	RegOpenUserClassesRoot func(
		Token HANDLE,
		Options DWORD,
		SamDesired REGSAM,
		Result *HKEY) LONG

	RegOpenCurrentUser func(
		SamDesired REGSAM, Result *HKEY) LONG

	RegDisablePredefinedCache func() LONG

	RegConnectRegistry func(
		MachineName VString,
		Key HKEY,
		Result *HKEY) LONG

	RegConnectRegistryEx func(
		MachineName VString,
		Key HKEY,
		Flags ULONG,
		Result *HKEY) LONG

	RegCreateKey func(
		Key HKEY, SubKey VString, Result *HKEY) LONG

	RegCreateKeyEx func(
		Key HKEY,
		SubKey VString,
		Reserved DWORD,
		Class VString,
		Options DWORD,
		SamDesired REGSAM,
		SecurityAttributes *SECURITY_ATTRIBUTES,
		Result *HKEY,
		Disposition *DWORD) LONG

	RegDeleteKey func(Key HKEY, SubKey VString) LONG

	RegDeleteKeyEx func(
		Key HKEY,
		SubKey VString,
		SamDesired REGSAM,
		Reserved DWORD) LONG

	RegDisableReflectionKey func(Base HKEY) LONG

	RegEnableReflectionKey func(Base HKEY) LONG

	RegQueryReflectionKey func(
		Base HKEY, IsReflectionDisabled *BOOL) LONG

	RegDeleteValue func(Key HKEY, ValueName VString) LONG

	RegEnumKey func(
		Key HKEY,
		Index DWORD,
		Name OVString,
		cName DWORD) LONG

	RegEnumKeyEx func(
		Key HKEY,
		Index DWORD,
		Name OVString,
		cName *DWORD,
		Reserved *DWORD,
		Class OVString,
		cClass *DWORD,
		LastWriteTime *FILETIME) LONG

	RegEnumValue func(
		Key HKEY,
		Index DWORD,
		ValueName OVString,
		cValueName *DWORD,
		Reserved *DWORD,
		Type *DWORD,
		Data *BYTE,
		cData *DWORD) LONG

	RegFlushKey func(Key HKEY) LONG

	RegGetKeySecurity func(
		Key HKEY,
		SecurityInformation SECURITY_INFORMATION,
		SecurityDescriptor *SECURITY_DESCRIPTOR,
		cSecurityDescriptor *DWORD) LONG

	RegLoadKey func(
		Key HKEY, SubKey, File VString) LONG

	RegNotifyChangeKeyValue func(
		Key HKEY,
		WatchSubtree BOOL,
		NotifyFilter DWORD,
		Event HANDLE,
		Asynchronous BOOL) LONG

	RegOpenKey func(
		Key HKEY, SubKey VString, Result *HKEY) LONG

	RegOpenKeyEx func(
		Key HKEY,
		SubKey VString,
		Options DWORD,
		SamDesired REGSAM,
		Result *HKEY) LONG

	RegQueryInfoKey func(
		Key HKEY,
		Class OVString,
		cClass,
		Reserved,
		SubKeys,
		MaxSubKeyLen,
		MaxClassLen,
		Values,
		MaxValueNameLen,
		MaxValueLen,
		SecurityDescriptor *DWORD,
		LastWriteTime *FILETIME) LONG

	RegQueryValue func(
		Key HKEY,
		SubKey VString,
		Data OVString,
		cData *LONG) LONG

	RegQueryMultipleValues func(
		Key HKEY,
		ValList *VALENT,
		NumVals DWORD,
		ValueBuf VString,
		TotSize *DWORD) LONG

	RegQueryValueEx func(
		Key HKEY,
		ValueName VString,
		Reserved,
		Type *DWORD,
		Data *BYTE,
		cData *DWORD) LONG

	RegReplaceKey func(
		Key HKEY, SubKey, NewFile, OldFile VString) LONG

	RegRestoreKey func(
		Key HKEY, File VString, Flags DWORD) LONG

	RegSaveKey func(
		Key HKEY,
		File VString,
		SecurityAttributes *SECURITY_ATTRIBUTES) LONG

	RegSetKeySecurity func(
		Key HKEY,
		SecurityInformation SECURITY_INFORMATION,
		pSecurityDescriptor *SECURITY_DESCRIPTOR) LONG

	RegSetValue func(
		Key HKEY,
		SubKey VString,
		Type DWORD,
		Data VString,
		cbData DWORD) LONG

	RegSetValueEx func(
		Key HKEY,
		ValueName VString,
		Reserved,
		Type DWORD,
		Data *BYTE,
		cData DWORD) LONG

	RegUnLoadKey func(Key HKEY, SubKey VString) LONG

	RegGetValue func(
		Key HKEY,
		SubKey, Value VString,
		Flags DWORD,
		Type *DWORD,
		Data *VOID,
		cData *DWORD) LONG

	InitiateSystemShutdown func(
		MachineName, Message VString,
		Timeout DWORD,
		ForceAppsClosed, RebootAfterShutdown BOOL) BOOL

	AbortSystemShutdown func(
		MachineName VString) BOOL

	InitiateSystemShutdownEx func(
		MachineName, Message VString,
		Timeout DWORD,
		ForceAppsClosed, RebootAfterShutdown BOOL,
		Reason DWORD) BOOL

	RegSaveKeyEx func(
		Key HKEY,
		File VString,
		SecurityAttributes *SECURITY_ATTRIBUTES,
		Flags DWORD) LONG

	Wow64Win32ApiEntry func(
		FuncNumber DWORD, Flag DWORD, Res DWORD) LONG
)

var WinRegANSIApis = Apis{
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

var WinRegUnicodeApis = Apis{
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

var WinRegApis = Apis{
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
