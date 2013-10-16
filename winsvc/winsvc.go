// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package winsvc provides API definitions for accessing
//advapi32.dll.
package winsvc

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	_ "github.com/tHinqa/outside/win32/advapi32"
)

var (
	ChangeServiceConfig func(
		Service T.SC_HANDLE,
		ServiceType,
		StartType,
		ErrorControl T.DWORD,
		BinaryPathName,
		LoadOrderGroup VString,
		TagId *T.DWORD,
		Dependencies,
		ServiceStartName,
		Password,
		DisplayName VString) (T.BOOL, error)

	ChangeServiceConfig2 func(
		Service T.SC_HANDLE,
		InfoLevel T.DWORD,
		Info *T.VOID) (T.BOOL, error)

	CloseServiceHandle func(SCObject T.SC_HANDLE) (T.BOOL, error)

	ControlService func(
		Service T.SC_HANDLE,
		Control T.DWORD,
		ServiceStatus *T.SERVICE_STATUS) (T.BOOL, error)

	CreateService func(
		SCManager T.SC_HANDLE,
		ServiceName,
		DisplayName VString,
		DesiredAccess,
		ServiceType,
		StartType,
		ErrorControl T.DWORD,
		BinaryPathName,
		LoadOrderGroup VString,
		TagId *T.DWORD,
		Dependencies,
		ServiceStartName,
		Password VString) (T.SC_HANDLE, error)

	DeleteService func(Service T.SC_HANDLE) (T.BOOL, error)

	EnumDependentServices func(
		Service T.SC_HANDLE,
		ServiceState T.DWORD,
		Services *T.ENUM_SERVICE_STATUS,
		BufSize T.DWORD,
		BytesNeeded,
		ServicesReturned *T.DWORD) (T.BOOL, error)

	EnumServicesStatus func(
		SCManager T.SC_HANDLE,
		ServiceType,
		ServiceState T.DWORD,
		Services *T.ENUM_SERVICE_STATUS,
		BufSize T.DWORD,
		BytesNeeded,
		ServicesReturned,
		ResumeHandle *T.DWORD) (T.BOOL, error)

	EnumServicesStatusEx func(
		SCManager T.SC_HANDLE,
		InfoLevel T.SC_ENUM_TYPE,
		ServiceType,
		ServiceState T.DWORD,
		Services *T.BYTE,
		BufSize T.DWORD,
		BytesNeeded,
		ServicesReturned,
		ResumeHandle *T.DWORD,
		pszGroupName VString) (T.BOOL, error)

	GetServiceKeyName func(
		SCManager T.SC_HANDLE,
		DisplayName VString,
		ServiceName OVString,
		cBuffer *T.DWORD) (T.BOOL, error)

	GetServiceDisplayName func(
		SCManager T.SC_HANDLE,
		ServiceName VString,
		DisplayName OVString,
		cBuffer *T.DWORD) (T.BOOL, error)

	LockServiceDatabase func(SCManager T.SC_HANDLE) (T.SC_LOCK, error)

	NotifyBootConfigStatus func(BootAcceptable T.BOOL) (T.BOOL, error)

	OpenSCManager func(
		MachineName,
		DatabaseName VString,
		DesiredAccess T.DWORD) (T.SC_HANDLE, error)

	OpenService func(
		SCManager T.SC_HANDLE,
		ServiceName VString,
		DesiredAccess T.DWORD) (T.SC_HANDLE, error)

	QueryServiceConfig func(
		Service T.SC_HANDLE,
		ServiceConfig *T.QUERY_SERVICE_CONFIG,
		BufSize T.DWORD,
		BytesNeeded *T.DWORD) (T.BOOL, error)

	QueryServiceConfig2 func(
		Service T.SC_HANDLE,
		InfoLevel T.DWORD,
		Buffer *T.BYTE,
		BufSize T.DWORD,
		BytesNeeded *T.DWORD) (T.BOOL, error)

	QueryServiceLockStatus func(
		SCManager T.SC_HANDLE,
		LockStatus *T.QUERY_SERVICE_LOCK_STATUS,
		BufSize T.DWORD,
		BytesNeeded *T.DWORD) (T.BOOL, error)

	QueryServiceObjectSecurity func(
		Service T.SC_HANDLE,
		SecurityInformation T.SECURITY_INFORMATION,
		SecurityDescriptor *T.SECURITY_DESCRIPTOR,
		BufSize T.DWORD,
		BytesNeeded *T.DWORD) (T.BOOL, error)

	QueryServiceStatus func(
		Service T.SC_HANDLE,
		ServiceStatus *T.SERVICE_STATUS) (T.BOOL, error)

	QueryServiceStatusEx func(
		Service T.SC_HANDLE,
		InfoLevel T.SC_STATUS_TYPE,
		Buffer *T.BYTE,
		BufSize T.DWORD,
		BytesNeeded *T.DWORD) (T.BOOL, error)

	RegisterServiceCtrlHandler func(
		ServiceName VString,
		HandlerProc *T.HANDLER_FUNCTION) (T.SERVICE_STATUS_HANDLE, error)

	RegisterServiceCtrlHandlerEx func(
		ServiceName VString,
		HandlerProc *T.HANDLER_FUNCTION_EX,
		Context *T.VOID) (T.SERVICE_STATUS_HANDLE, error)

	SetServiceObjectSecurity func(
		Service T.SC_HANDLE,
		SecurityInformation T.SECURITY_INFORMATION,
		SecurityDescriptor *T.SECURITY_DESCRIPTOR) (T.BOOL, error)

	SetServiceStatus func(
		ServiceStatusH T.SERVICE_STATUS_HANDLE,
		ServiceStatus *T.SERVICE_STATUS) (T.BOOL, error)

	StartServiceCtrlDispatcher func(
		ServiceStartTable *T.SERVICE_TABLE_ENTRY) (T.BOOL, error)

	StartService func(
		Service T.SC_HANDLE,
		NumServiceArgs T.DWORD,
		ServiceArgVectors *VString) (T.BOOL, error)

	UnlockServiceDatabase func(ScLock T.SC_LOCK) (T.BOOL, error)
)

var WinSvcANSIApis = outside.Apis{
	{"ChangeServiceConfigA", &ChangeServiceConfig},
	{"ChangeServiceConfig2A", &ChangeServiceConfig2},
	{"CreateServiceA", &CreateService},
	{"EnumDependentServicesA", &EnumDependentServices},
	{"EnumServicesStatusA", &EnumServicesStatus},
	{"EnumServicesStatusExA", &EnumServicesStatusEx},
	{"GetServiceKeyNameA", &GetServiceKeyName},
	{"GetServiceDisplayNameA", &GetServiceDisplayName},
	{"OpenSCManagerA", &OpenSCManager},
	{"OpenServiceA", &OpenService},
	{"QueryServiceConfigA", &QueryServiceConfig},
	{"QueryServiceConfig2A", &QueryServiceConfig2},
	{"QueryServiceLockStatusA", &QueryServiceLockStatus},
	{"RegisterServiceCtrlHandlerA", &RegisterServiceCtrlHandler},
	{"RegisterServiceCtrlHandlerExA", &RegisterServiceCtrlHandlerEx},
	{"StartServiceCtrlDispatcherA", &StartServiceCtrlDispatcher},
	{"StartServiceA", &StartService},
}

var WinSvcUnicodeApis = outside.Apis{
	{"ChangeServiceConfigW", &ChangeServiceConfig},
	{"ChangeServiceConfig2W", &ChangeServiceConfig2},
	{"CreateServiceW", &CreateService},
	{"EnumDependentServicesW", &EnumDependentServices},
	{"EnumServicesStatusW", &EnumServicesStatus},
	{"EnumServicesStatusExW", &EnumServicesStatusEx},
	{"GetServiceKeyNameW", &GetServiceKeyName},
	{"GetServiceDisplayNameW", &GetServiceDisplayName},
	{"OpenSCManagerW", &OpenSCManager},
	{"OpenServiceW", &OpenService},
	{"QueryServiceConfigW", &QueryServiceConfig},
	{"QueryServiceConfig2W", &QueryServiceConfig2},
	{"QueryServiceLockStatusW", &QueryServiceLockStatus},
	{"RegisterServiceCtrlHandlerW", &RegisterServiceCtrlHandler},
	{"RegisterServiceCtrlHandlerExW", &RegisterServiceCtrlHandlerEx},
	{"StartServiceCtrlDispatcherW", &StartServiceCtrlDispatcher},
	{"StartServiceW", &StartService},
}

var WinSvcApis = outside.Apis{
	{"CloseServiceHandle", &CloseServiceHandle},
	{"ControlService", &ControlService},
	{"DeleteService", &DeleteService},
	{"LockServiceDatabase", &LockServiceDatabase},
	{"NotifyBootConfigStatus", &NotifyBootConfigStatus},
	{"QueryServiceObjectSecurity", &QueryServiceObjectSecurity},
	{"QueryServiceStatus", &QueryServiceStatus},
	{"QueryServiceStatusEx", &QueryServiceStatusEx},
	{"SetServiceObjectSecurity", &SetServiceObjectSecurity},
	{"SetServiceStatus", &SetServiceStatus},
	{"UnlockServiceDatabase", &UnlockServiceDatabase},
}
