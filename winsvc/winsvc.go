package winsvc

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/advapi32"
)

var (
	ChangeServiceConfig func(
		Service SC_HANDLE,
		ServiceType,
		StartType,
		ErrorControl DWORD,
		BinaryPathName,
		LoadOrderGroup VString,
		TagId *DWORD,
		Dependencies,
		ServiceStartName,
		Password,
		DisplayName VString) BOOL

	ChangeServiceConfig2 func(
		Service SC_HANDLE,
		InfoLevel DWORD,
		Info *VOID) BOOL

	CloseServiceHandle func(SCObject SC_HANDLE) BOOL

	ControlService func(
		Service SC_HANDLE,
		Control DWORD,
		ServiceStatus *SERVICE_STATUS) BOOL

	CreateService func(
		SCManager SC_HANDLE,
		ServiceName,
		DisplayName VString,
		DesiredAccess,
		ServiceType,
		StartType,
		ErrorControl DWORD,
		BinaryPathName,
		LoadOrderGroup VString,
		TagId *DWORD,
		Dependencies,
		ServiceStartName,
		Password VString) SC_HANDLE

	DeleteService func(Service SC_HANDLE) BOOL

	EnumDependentServices func(
		Service SC_HANDLE,
		ServiceState DWORD,
		Services *ENUM_SERVICE_STATUS,
		BufSize DWORD,
		BytesNeeded,
		ServicesReturned *DWORD) BOOL

	EnumServicesStatus func(
		SCManager SC_HANDLE,
		ServiceType,
		ServiceState DWORD,
		Services *ENUM_SERVICE_STATUS,
		BufSize DWORD,
		BytesNeeded,
		ServicesReturned,
		ResumeHandle *DWORD) BOOL

	EnumServicesStatusEx func(
		SCManager SC_HANDLE,
		InfoLevel SC_ENUM_TYPE,
		ServiceType,
		ServiceState DWORD,
		Services *BYTE,
		BufSize DWORD,
		BytesNeeded,
		ServicesReturned,
		ResumeHandle *DWORD,
		pszGroupName VString) BOOL

	GetServiceKeyName func(
		SCManager SC_HANDLE,
		DisplayName VString,
		ServiceName OVString,
		cBuffer *DWORD) BOOL

	GetServiceDisplayName func(
		SCManager SC_HANDLE,
		ServiceName VString,
		DisplayName OVString,
		cBuffer *DWORD) BOOL

	LockServiceDatabase func(SCManager SC_HANDLE) SC_LOCK

	NotifyBootConfigStatus func(BootAcceptable BOOL) BOOL

	OpenSCManager func(
		MachineName,
		DatabaseName VString,
		DesiredAccess DWORD) SC_HANDLE

	OpenService func(
		SCManager SC_HANDLE,
		ServiceName VString,
		DesiredAccess DWORD) SC_HANDLE

	QueryServiceConfig func(
		Service SC_HANDLE,
		ServiceConfig *QUERY_SERVICE_CONFIG,
		BufSize DWORD,
		BytesNeeded *DWORD) BOOL

	QueryServiceConfig2 func(
		Service SC_HANDLE,
		InfoLevel DWORD,
		Buffer *BYTE,
		BufSize DWORD,
		BytesNeeded *DWORD) BOOL

	QueryServiceLockStatus func(
		SCManager SC_HANDLE,
		LockStatus *QUERY_SERVICE_LOCK_STATUS,
		BufSize DWORD,
		BytesNeeded *DWORD) BOOL

	QueryServiceObjectSecurity func(
		Service SC_HANDLE,
		SecurityInformation SECURITY_INFORMATION,
		SecurityDescriptor *SECURITY_DESCRIPTOR,
		BufSize DWORD,
		BytesNeeded *DWORD) BOOL

	QueryServiceStatus func(
		Service SC_HANDLE,
		ServiceStatus *SERVICE_STATUS) BOOL

	QueryServiceStatusEx func(
		Service SC_HANDLE,
		InfoLevel SC_STATUS_TYPE,
		Buffer *BYTE,
		BufSize DWORD,
		BytesNeeded *DWORD) BOOL

	RegisterServiceCtrlHandler func(
		ServiceName VString,
		HandlerProc *HANDLER_FUNCTION) SERVICE_STATUS_HANDLE

	RegisterServiceCtrlHandlerEx func(
		ServiceName VString,
		HandlerProc *HANDLER_FUNCTION_EX,
		Context *VOID) SERVICE_STATUS_HANDLE

	SetServiceObjectSecurity func(
		Service SC_HANDLE,
		SecurityInformation SECURITY_INFORMATION,
		SecurityDescriptor *SECURITY_DESCRIPTOR) BOOL

	SetServiceStatus func(
		ServiceStatusH SERVICE_STATUS_HANDLE,
		ServiceStatus *SERVICE_STATUS) BOOL

	StartServiceCtrlDispatcher func(
		ServiceStartTable *SERVICE_TABLE_ENTRY) BOOL

	StartService func(
		Service SC_HANDLE,
		NumServiceArgs DWORD,
		ServiceArgVectors *VString) BOOL

	UnlockServiceDatabase func(ScLock SC_LOCK) BOOL
)

var WinSvcANSIApis = Apis{
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

var WinSvcUnicodeApis = Apis{
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

var WinSvcApis = Apis{
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
