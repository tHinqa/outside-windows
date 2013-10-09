// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package types provides API structs, enums and other types.
package types

//TODO(t): bytes(bool) cannot be returned by callbacks so
//	BOOLS/BOOLEANS cannot easily be transformed to bool
//TODO(t): Check all callback use for indirections (structs and funcs)

import (
	. "github.com/tHinqa/outside"
	"unsafe"
)

type (
	Fake_type_Fix_me uintptr
	enum             int

	ACCESS_MASK                    DWORD
	AChar                          Char
	ALG_ID                         Unsigned_int
	AString                        *string
	ATOM                           WORD
	BOOL                           int
	BOOLEAN                        byte
	BSTR                           *OLECHAR
	BYTE                           byte
	CALID                          DWORD
	CALTYPE                        DWORD
	CH                             AChar
	CLIPFORMAT                     WORD
	CLSID                          GUID
	COLOR16                        USHORT
	COLORREF                       DWORD
	Char                           uint8
	DATE                           DOUBLE
	DISPID                         int32
	DOUBLE                         float64
	DWORD                          uint32 //TODO(t):True w32/w64?
	DWORDLONG                      uint64 //TODO(t):True w32/w64?
	DWORD_PTR                      ULONG_PTR
	EXECUTION_STATE                DWORD
	FARPROC                        uintptr //TODO(t):unsafe.Pointer?
	FLOAT                          float32
	FOURCC                         DWORD
	FXPT2DOT30                     LONG
	GEOCLASS                       DWORD
	GEOID                          LONG
	GEOTYPE                        DWORD
	GROUP                          uint
	HACCEL                         HANDLE
	HANDLE                         DWORD //TODO(t):Actually strict -> struct but go does not do coercion
	HBITMAP                        int   //TODO(t):HANDLE but go does not do coercion
	HBRUSH                         HANDLE
	HCOLORSPACE                    HANDLE
	HCONV                          HANDLE
	HCONVLIST                      HANDLE
	HCRYPTHASH                     ULONG_PTR
	HCRYPTKEY                      ULONG_PTR
	HCRYPTPROV                     ULONG_PTR
	HCURSOR                        HANDLE
	HDC                            HANDLE
	HDDEDATA                       HANDLE
	HDESK                          HANDLE
	HDEVNOTIFY                     HANDLE
	HDPA                           HANDLE
	HDROP                          HANDLE
	HDRVR                          HANDLE
	HDSA                           HANDLE
	HDWP                           HANDLE
	HENHMETAFILE                   HANDLE
	HFILE                          HANDLE
	HFONT                          HANDLE
	HGDIOBJ                        HANDLE
	HGLOBAL                        HANDLE
	HGLRC                          HANDLE
	HHOOK                          HANDLE
	HICON                          HANDLE
	HIMAGELIST                     HANDLE
	HINSTANCE                      HANDLE
	HKEY                           HANDLE
	HKL                            HANDLE
	HLOCAL                         HANDLE
	HMENU                          HANDLE
	HMETAFILE                      HANDLE
	HMIDI                          HANDLE
	HMIDIIN                        HANDLE
	HMIDIOUT                       HANDLE
	HMIDISTRM                      HANDLE
	HMIXER                         HANDLE
	HMIXEROBJ                      HANDLE
	HMMIO                          HANDLE
	HMODULE                        HINSTANCE
	HMONITOR                       HANDLE
	HOLEMENU                       HANDLE
	HPALETTE                       HANDLE
	HPEN                           HANDLE
	HRAWINPUT                      HANDLE
	HRESULT                        HANDLE
	HRGN                           HANDLE
	HRSRC                          HANDLE
	HSZ                            HANDLE
	HTASK                          HANDLE
	HWAVEIN                        HANDLE
	HWAVEOUT                       HANDLE
	HWINEVENTHOOK                  HANDLE
	HWINSTA                        HANDLE
	HWND                           HANDLE
	IAdviseSink                    struct{}
	IBindCtx                       struct{}
	IBindStatusCallback            struct{}
	IChannelHook                   struct{}
	IClassFactory                  struct{}
	ICreateErrorInfo               struct{}
	ICreateTypeLib                 struct{}
	ICreateTypeLib2                struct{}
	IDataAdviseHolder              struct{}
	IDataObject                    struct{}
	IDispatch                      struct{}
	IDropSource                    struct{}
	IDropTarget                    struct{}
	IEnumFORMATETC                 struct{}
	IEnumOLEVERB                   struct{}
	IErrorInfo                     struct{}
	IFillLockBytes                 struct{}
	IID                            GUID
	IInitializeSpy                 struct{}
	ILockBytes                     struct{}
	IMalloc                        struct{}
	IMallocSpy                     struct{}
	IMarshal                       struct{}
	IMessageFilter                 struct{}
	IMoniker                       struct{}
	INT                            int
	INT_PTR                        INT
	IOAString                      *string
	IOleAdviseHolder               struct{}
	IOleClientSite                 struct{}
	IOleInPlaceActiveObject        struct{}
	IOleInPlaceFrame               struct{}
	IOleObject                     struct{}
	IPersistStorage                struct{}
	IPersistStream                 struct{}
	IRecordInfo                    struct{}
	IRunningObjectTable            struct{}
	IStorage                       struct{}
	IStream                        struct{}
	ISurrogate                     struct{}
	ITypeInfo                      struct{}
	ITypeLib                       struct{}
	IUnknown                       struct{}
	LANGID                         USHORT
	LCID                           DWORD
	LCSCSTYPE                      LONG
	LCSGAMUTMATCH                  LONG
	LCTYPE                         DWORD
	LGRPID                         DWORD
	LONG                           int32 //TODO(t):?SIZE __3264
	LONG64                         int64
	LONGLONG                       int64 //TODO(t):Win64=128
	LONG_PTR                       LONG
	LPARAM                         LONG_PTR
	LRESULT                        LONG_PTR
	MCIDEVICEID                    UINT
	MCIERROR                       DWORD
	MENUTEMPLATE                   *VOID
	MMRESULT                       UINT
	MMVERSION                      UINT
	MVString                       *string //TODO(t):*[]string?
	NLS_FUNCTION                   DWORD
	OAString                       *string
	OLECHAR                        WChar
	OLESTR                         WString
	OLESTREAM                      struct{}
	OMVString                      *string //TODO(t):*[]string ?
	OWString                       *string
	REFCLSID                       *IID
	REFGUID                        *GUID
	REFIID                         *IID
	REGSAM                         ACCESS_MASK
	RPC_AUTHZ_HANDLE               *VOID
	RPC_AUTH_IDENTITY_HANDLE       *VOID
	SCODE                          int32
	SC_HANDLE                      HANDLE
	SC_LOCK                        *VOID
	SECURITY_CONTEXT_TRACKING_MODE BOOLEAN
	SECURITY_DESCRIPTOR_CONTROL    USHORT
	SECURITY_INFORMATION           DWORD
	SERVICETYPE                    ULONG
	SERVICE_STATUS_HANDLE          HANDLE
	SHORT                          int16
	SIZE_T                         ULONG_PTR
	SNB                            **OLECHAR
	SOCKET                         UINT_PTR
	STREAM                         IStream
	UCHAR                          byte
	UINT                           uint
	UINT_PTR                       UINT
	ULONG                          DWORD //TODO(t):size
	ULONG64                        uint64
	ULONGLONG                      uint64  //TODO(t):Win64=128?
	ULONG_PTR                      uintptr //TODO(t):true def
	USHORT                         uint16  //TODO(t):size
	U_int                          uint
	U_long                         uint32
	U_short                        uint16
	Unsigned_int                   uint
	Unsigned_long                  uint32
	Unsigned_short                 uint16
	VARIANTARG                     VARIANT
	VARIANT_BOOL                   int16
	VARTYPE                        uint16
	VChar                          *uint16  //TODO(t): uint8/uint16
	VOID                           struct{} // //TODO(t):Go does not do coercion; Any side-effects?
	VOID64                         struct{} //TODO(t):__ptr64 //TODO(t):Go does not do coercion; Any side-effects?
	WChar                          uint16
	WORD                           uint16 //TODO(t):True 32/64? signed?
	WPARAM                         LONG_PTR
	WSAEVENT                       HANDLE
	WSAOVERLAPPED                  OVERLAPPED
	WString                        *string
)

type FILETIME struct {
	LowDateTime  DWORD
	HighDateTime DWORD
}

type STARTUPINFO struct {
	Len           DWORD
	_             *VOID // RESERVED POVString
	Desktop       POVString
	Title         POVString
	X             DWORD
	Y             DWORD
	XSize         DWORD
	YSize         DWORD
	XCountChars   DWORD
	YCountChars   DWORD
	FillAttribute DWORD
	Flags         DWORD
	ShowWindow    WORD
	_             WORD  // wReserved2
	_             *BYTE // lpReserved2
	StdInput      HANDLE
	StdOutput     HANDLE
	StdError      HANDLE
}

type OSVERSIONINFOA struct {
	OSVersionInfoSize DWORD
	MajorVersion      DWORD
	MinorVersion      DWORD
	BuildNumber       DWORD
	PlatformId        DWORD
	CSDVersion        [128]AChar
}

type OSVERSIONINFOW struct {
	OSVersionInfoSize DWORD
	MajorVersion      DWORD
	MinorVersion      DWORD
	BuildNumber       DWORD
	PlatformId        DWORD
	CSDVersion        [128]WChar
}

type OSVERSIONINFOEXA struct {
	OSVersionInfoSize DWORD
	MajorVersion      DWORD
	MinorVersion      DWORD
	BuildNumber       DWORD
	PlatformId        DWORD
	CSDVersion        [128]AChar
	ServicePackMajor  WORD
	ServicePackMinor  WORD
	SuiteMask         WORD
	ProductType       BYTE
	Reserved          BYTE
}

type OSVERSIONINFOEXW struct {
	OSVersionInfoSize DWORD
	MajorVersion      DWORD
	MinorVersion      DWORD
	BuildNumber       DWORD
	PlatformId        DWORD
	CSDVersion        [128]WChar
	ServicePackMajor  WORD
	ServicePackMinor  WORD
	SuiteMask         WORD
	ProductType       BYTE
	Reserved          BYTE
}

type SECURITY_ATTRIBUTES struct {
	Length             DWORD
	SecurityDescriptor *VOID
	InheritHandle      BOOL
}

type PROCESS_INFORMATION struct {
	Process   HANDLE
	Thread    HANDLE
	ProcessId DWORD
	ThreadId  DWORD
}

type SYSTEMTIME struct {
	Year         WORD
	Month        WORD
	DayOfWeek    WORD
	Day          WORD
	Hour         WORD
	Minute       WORD
	Second       WORD
	Milliseconds WORD
}

type MEMORYSTATUS struct {
	Length        DWORD
	MemoryLoad    DWORD
	TotalPhys     SIZE_T
	AvailPhys     SIZE_T
	TotalPageFile SIZE_T
	AvailPageFile SIZE_T
	TotalVirtual  SIZE_T
	AvailVirtual  SIZE_T
}

type EXCEPTION_DEBUG_INFO struct {
	ExceptionRecord EXCEPTION_RECORD
	FirstChance     DWORD
}

type CREATE_THREAD_DEBUG_INFO struct {
	Thread          HANDLE
	ThreadLocalBase *VOID
	StartAddress    THREAD_START_ROUTINE
}

type CREATE_PROCESS_DEBUG_INFO struct {
	File                HANDLE
	Process             HANDLE
	Thread              HANDLE
	BaseOfImage         *VOID
	DebugInfoFileOffset DWORD
	DebugInfoSize       DWORD
	ThreadLocalBase     *VOID
	StartAddress        THREAD_START_ROUTINE
	ImageName           *VOID
	Unicode             WORD //TODO(t): flag
}

type EXIT_THREAD_DEBUG_INFO struct {
	ExitCode DWORD
}

type EXIT_PROCESS_DEBUG_INFO struct {
	ExitCode DWORD
}

type LOAD_DLL_DEBUG_INFO struct {
	File                HANDLE
	BaseOfDll           *VOID
	DebugInfoFileOffset DWORD
	DebugInfoSize       DWORD
	ImageName           *VOID
	Unicode             WORD
}

type UNLOAD_DLL_DEBUG_INFO struct {
	BaseOfDll *VOID
}

type OUTPUT_DEBUG_STRING_INFO struct {
	DebugStringData   AString
	Unicode           WORD //TODO(t):use bool
	DebugStringLength WORD
}

type RIP_INFO struct {
	Error DWORD
	Type  DWORD
}

const OFS_MAXPATHNAME = 128

type OFSTRUCT struct {
	Bytes     BYTE
	FixedDisk BYTE
	ErrCode   WORD
	_, _      WORD
	PathName  [OFS_MAXPATHNAME]AChar
}

type MEMORYSTATUSEX struct {
	Length               DWORD
	MemoryLoad           DWORD
	TotalPhys            DWORDLONG
	AvailPhys            DWORDLONG
	TotalPageFile        DWORDLONG
	AvailPageFile        DWORDLONG
	TotalVirtual         DWORDLONG
	AvailVirtual         DWORDLONG
	AvailExtendedVirtual DWORDLONG
}

//type EXCEPTION_RECORD Fix

type BY_HANDLE_FILE_INFORMATION struct {
	FileAttributes     DWORD
	CreationTime       FILETIME
	LastAccessTime     FILETIME
	LastWriteTime      FILETIME
	VolumeSerialNumber DWORD
	FileSizeHigh       DWORD
	FileSizeLow        DWORD
	NumberOfLinks      DWORD
	FileIndexHigh      DWORD
	FileIndexLow       DWORD
}

type TIME_ZONE_INFORMATION struct {
	Bias         LONG
	StandardName [32]WChar
	StandardDate SYSTEMTIME
	StandardBias LONG
	DaylightName [32]WChar
	DaylightDate SYSTEMTIME
	DaylightBias LONG
}

const ANYSIZE_ARRAY = 1 //TODO(t):?

type WIN32_STREAM_ID struct {
	StreamId         DWORD
	StreamAttributes DWORD
	Size             LARGE_INTEGER
	StreamNameSize   DWORD
	StreamName       [ANYSIZE_ARRAY]WChar
}

const MAX_PATH = 1 //TODO(t):?

type WIN32_FIND_DATAA struct {
	FileAttributes    DWORD
	CreationTime      FILETIME
	LastAccessTime    FILETIME
	LastWriteTime     FILETIME
	FileSizeHigh      DWORD
	FileSizeLow       DWORD
	_, _              DWORD
	FileName          [MAX_PATH]AChar
	AlternateFileName [14]AChar
}

type WIN32_FIND_DATAW struct {
	FileAttributes    DWORD
	CreationTime      FILETIME
	LastAccessTime    FILETIME
	LastWriteTime     FILETIME
	FileSizeHigh      DWORD
	FileSizeLow       DWORD
	_, _              DWORD
	FileName          [MAX_PATH]WChar
	AlternateFileName [14]WChar
}

type WIN32_FILE_ATTRIBUTE_DATA struct {
	FileAttributes DWORD
	CreationTime   FILETIME
	LastAccessTime FILETIME
	LastWriteTime  FILETIME
	FileSizeHigh   DWORD
	FileSizeLow    DWORD
}

type COMMTIMEOUTS struct {
	ReadIntervalTimeout         DWORD
	ReadTotalTimeoutMultiplier  DWORD
	ReadTotalTimeoutConstant    DWORD
	WriteTotalTimeoutMultiplier DWORD
	WriteTotalTimeoutConstant   DWORD
}

type WIN32_FIND_STREAM_DATA struct {
	StreamSize LARGE_INTEGER
	StreamName [MAX_PATH + 36]WChar
}

type STREAM_INFO_LEVELS enum

const (
	FindStreamInfoStandard STREAM_INFO_LEVELS = iota
	FindStreamInfoMaxInfoLevel
)

type EVENTLOG_FULL_INFORMATION struct {
	dwFull DWORD
}

type SYSTEM_POWER_STATUS struct {
	ACLineStatus        BYTE
	BatteryFlag         BYTE
	BatteryLifePercent  BYTE
	Reserved1           BYTE
	BatteryLifeTime     DWORD
	BatteryFullLifeTime DWORD
}

type ACTCTX struct {
	Size                  ULONG
	Flags                 DWORD
	Source                PVString
	ProcessorArchitecture USHORT
	LangId                LANGID
	AssemblyDirectory     PVString
	ResourceName          PVString
	ApplicationName       PVString
	Module                HMODULE
}

const (
	HW_PROFILE_GUIDLEN = 39
	MAX_PROFILE_LEN    = 80
)

type HW_PROFILE_INFOA struct {
	DockInfo      DWORD
	HwProfileGuid [HW_PROFILE_GUIDLEN]AChar
	HwProfileName [MAX_PROFILE_LEN]AChar
}

type HW_PROFILE_INFOW struct {
	DockInfo      DWORD
	HwProfileGuid [HW_PROFILE_GUIDLEN]WChar
	HwProfileName [MAX_PROFILE_LEN]WChar
}

type ACTCTX_SECTION_KEYED_DATA_2600 struct {
	Size                    ULONG
	DataFormatVersion       ULONG
	Data                    *VOID
	Length                  ULONG
	SectionGlobalData       *VOID
	SectionGlobalDataLength ULONG
	SectionBase             *VOID
	SectionTotalLength      ULONG
	ActCtx                  HANDLE
	AssemblyRosterIndex     ULONG
}

type ACTCTX_SECTION_KEYED_DATA_ASSEMBLY_METADATA struct {
	Information             *VOID
	SectionBase             *VOID
	SectionLength           ULONG
	SectionGlobalDataBase   *VOID
	SectionGlobalDataLength ULONG
}

type ACTCTX_SECTION_KEYED_DATA struct {
	Size                    ULONG
	DataFormatVersion       ULONG
	Data                    *VOID
	Length                  ULONG
	SectionGlobalData       *VOID
	SectionGlobalDataLength ULONG
	SectionBase             *VOID
	SectionTotalLength      ULONG
	ActCtx                  HANDLE
	AssemblyRosterIndex     ULONG
	// 2600 stops here
	Flags            ULONG
	AssemblyMetadata ACTCTX_SECTION_KEYED_DATA_ASSEMBLY_METADATA
}

type ACTIVATION_CONTEXT_BASIC_INFORMATION struct {
	ActCtx HANDLE
	Flags  DWORD
}

type LUID struct {
	LowPart  DWORD
	HighPart LONG
}

type SECURITY_DESCRIPTOR struct {
	Revision BYTE
	Sbz1     BYTE
	Control  SECURITY_DESCRIPTOR_CONTROL
	Owner    *SID
	Group    *SID
	Sacl     *ACL
	Dacl     *ACL
}

type ACL struct {
	AclRevision BYTE
	Sbz1        BYTE
	AclSize     WORD
	AceCount    WORD
	Sbz2        WORD
}

type GENERIC_MAPPING struct {
	GenericRead    ACCESS_MASK
	GenericWrite   ACCESS_MASK
	GenericExecute ACCESS_MASK
	GenericAll     ACCESS_MASK
}

type PRIVILEGE_SET struct {
	PrivilegeCount DWORD
	Control        DWORD
	Privilege      [ANYSIZE_ARRAY]LUID_AND_ATTRIBUTES
}

type LUID_AND_ATTRIBUTES struct {
	Luid       LUID
	Attributes DWORD
}

type OBJECT_TYPE_LIST struct {
	Level      WORD
	Sbz        WORD
	ObjectType *GUID
}

type AUDIT_EVENT_TYPE enum

const (
	AuditEventObjectAccess AUDIT_EVENT_TYPE = iota
	AuditEventDirectoryServiceAccess
)

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type VECTORED_EXCEPTION_HANDLER func(
	ExceptionInfo *EXCEPTION_POINTERS) LONG

type EXCEPTION_POINTERS struct {
	ExceptionRecord *EXCEPTION_RECORD
	ContextRecord   *CONTEXT
}

type TOKEN_GROUPS struct {
	GroupCount DWORD
	Groups     [ANYSIZE_ARRAY]SID_AND_ATTRIBUTES
}

type SID_AND_ATTRIBUTES struct {
	Sid        *SID
	Attributes DWORD
}

type TOKEN_PRIVILEGES struct {
	PrivilegeCount DWORD
	Privileges     [ANYSIZE_ARRAY]LUID_AND_ATTRIBUTES
}

type SID_IDENTIFIER_AUTHORITY struct {
	Value [6]BYTE
}

type SID struct {
	Revision            BYTE
	SubAuthorityCount   BYTE
	IdentifierAuthority SID_IDENTIFIER_AUTHORITY
	SubAuthority        [ANYSIZE_ARRAY]DWORD
}

type OVERLAPPED_COMPLETION_ROUTINE func(
	ErrorCode, NumberOfBytesTransfered DWORD,
	Overlapped *OVERLAPPED)

type DCB struct {
	DCBlength, BaudRate, Bitfields DWORD
	//TODO(t):way of access
	// Binary: 1
	// Parity: 1
	// OutxCtsFlow:1
	// OutxDsrFlow:1
	// DtrControl:2
	// DsrSensitivity:1
	// TXContinueOnXoff: 1
	// OutX: 1
	// InX: 1
	// ErrorChar: 1
	// Null: 1
	// RtsControl:2
	// AbortOnError:1
	// Dummy2:17
	Reserved, XonLim, XoffLim                      WORD
	ByteSize, Parity, StopBits                     BYTE
	XonChar, XoffChar, ErrorChar, EofChar, EvtChar AChar
	Reserved1                                      WORD
}

type COMSTAT struct {
	Bitfields,
	//TODO(t):way of access
	// CtsHold : 1
	// DsrHold : 1
	// RlsdHold : 1
	// XoffHold : 1
	// XoffSent : 1
	// Eof : 1
	// Txim : 1
	// Reserved : 25
	InQue, OutQue DWORD
}

type COMMCONFIG struct {
	Size            DWORD
	Version         WORD
	Reserved        WORD
	Dcb             DCB
	ProviderSubType DWORD
	ProviderOffset  DWORD
	ProviderSize    DWORD
	ProviderData    [1]WChar
}

type OVERLAPPED struct {
	Internal     ULONG_PTR
	InternalHigh ULONG_PTR
	Offset       DWORD //TODO(t): ALSO  Pointer *VOID
	OffsetHigh   DWORD
	Event        HANDLE
}

type PROGRESS_ROUTINE func(
	TotalFileSize, TotalBytesTransferred, StreamSize,
	StreamBytesTransferred LARGE_INTEGER,
	StreamNumber,
	CallbackReason DWORD,
	SourceFile,
	DestinationFile HANDLE,
	Data *VOID) DWORD

type FIBER_START_ROUTINE func(FiberParameter *VOID)

type JOB_SET_ARRAY struct {
	JobHandle   HANDLE
	MemberLevel DWORD
	Flags       DWORD
}

type MEMORY_RESOURCE_NOTIFICATION_TYPE enum

const (
	LowMemoryResourceNotification MEMORY_RESOURCE_NOTIFICATION_TYPE = iota
	HighMemoryResourceNotification
)

type THREAD_START_ROUTINE func(
	ThreadParameter *VOID) DWORD

type WAITORTIMERCALLBACK func(*VOID, BOOLEAN)

type WELL_KNOWN_SID_TYPE enum

const (
	WinNullSid = iota
	WinWorldSid
	WinLocalSid
	WinCreatorOwnerSid
	WinCreatorGroupSid
	WinCreatorOwnerServerSid
	WinCreatorGroupServerSid
	WinNtAuthoritySid
	WinDialupSid
	WinNetworkSid
	WinBatchSid
	WinInteractiveSid
	WinServiceSid
	WinAnonymousSid
	WinProxySid
	WinEnterpriseControllersSid
	WinSelfSid
	WinAuthenticatedUserSid
	WinRestrictedCodeSid
	WinTerminalServerSid
	WinRemoteLogonIdSid
	WinLogonIdsSid
	WinLocalSystemSid
	WinLocalServiceSid
	WinNetworkServiceSid
	WinBuiltinDomainSid
	WinBuiltinAdministratorsSid
	WinBuiltinUsersSid
	WinBuiltinGuestsSid
	WinBuiltinPowerUsersSid
	WinBuiltinAccountOperatorsSid
	WinBuiltinSystemOperatorsSid
	WinBuiltinPrintOperatorsSid
	WinBuiltinBackupOperatorsSid
	WinBuiltinReplicatorSid
	WinBuiltinPreWindows2000CompatibleAccessSid
	WinBuiltinRemoteDesktopUsersSid
	WinBuiltinNetworkConfigurationOperatorsSid
	WinAccountAdministratorSid
	WinAccountGuestSid
	WinAccountKrbtgtSid
	WinAccountDomainAdminsSid
	WinAccountDomainUsersSid
	WinAccountDomainGuestsSid
	WinAccountComputersSid
	WinAccountControllersSid
	WinAccountCertAdminsSid
	WinAccountSchemaAdminsSid
	WinAccountEnterpriseAdminsSid
	WinAccountPolicyAdminsSid
	WinAccountRasAndIasServersSid
	WinNTLMAuthenticationSid
	WinDigestAuthenticationSid
	WinSChannelAuthenticationSid
	WinThisOrganizationSid
	WinOtherOrganizationSid
	WinBuiltinIncomingForestTrustBuildersSid
	WinBuiltinPerfMonitoringUsersSid
	WinBuiltinPerfLoggingUsersSid
	WinBuiltinAuthorizationAccessSid
	WinBuiltinTerminalServerLicenseServersSid
	WinBuiltinDCOMUsersSid
)

type CRITICAL_SECTION struct {
	DebugInfo                   *CRITICAL_SECTION_DEBUG
	LockCount, RecursionCount   LONG
	OwningThread, LockSemaphore HANDLE
	SpinCount                   ULONG_PTR
}

type CRITICAL_SECTION_DEBUG struct {
	Type, CreatorBackTraceIndex WORD
	CriticalSection             *CRITICAL_SECTION
	ProcessLocksList            LIST_ENTRY
	EntryCount, ContentionCount DWORD
	Spare                       [2]DWORD
}

type LIST_ENTRY struct {
	Flink, Blink *LIST_ENTRY
}
type SECURITY_IMPERSONATION_LEVEL enum

const (
	SecurityAnonymous SECURITY_IMPERSONATION_LEVEL = iota
	SecurityIdentification
	SecurityImpersonation
	SecurityDelegation
)

type ENUMRESLANGPROC func(
	Module, HMODULE,
	Type, Name PVString,
	Language WORD,
	Param LONG_PTR) BOOL

type ENUMRESTYPEPROC func(
	Module HMODULE, Type PVString, Param LONG_PTR) BOOL

type ENUMRESNAMEPROC func(
	Module HMODULE,
	Type PVString,
	Name PVString, //TODO(t):???? LPSTR
	// lagId WORD missing?
	Param LONG_PTR) BOOL

type FINDEX_INFO_LEVELS enum

const (
	FindExInfoStandard FINDEX_INFO_LEVELS = iota
	FindExInfoMaxInfoLevel
)

type FINDEX_SEARCH_OPS enum

const (
	FindExSearchNameMatch FINDEX_SEARCH_OPS = iota
	FindExSearchLimitToDirectories
	FindExSearchLimitToDevices
	FindExSearchMaxSearchOp
)

type FLS_CALLBACK_FUNCTION func(FlsData *VOID)

type ACL_INFORMATION_CLASS enum

const (
	AclRevisionInformation ACL_INFORMATION_CLASS = iota + 1
	AclSizeInformation
)

type ULARGE_INTEGER struct {
	LowPart  DWORD
	HighPart DWORD
	//TODO(t): UNION QuadPart ULONGLONG
}

type GET_FILEEX_INFO_LEVELS enum

const (
	GetFileExInfoStandard GET_FILEEX_INFO_LEVELS = iota
	GetFileExMaxInfoLevel
)

type SYSTEM_LOGICAL_PROCESSOR_INFORMATION struct {
	ProcessorMask ULONG_PTR
	Relationship  LOGICAL_PROCESSOR_RELATIONSHIP
	//TODO(t): union {BYTE  Flags;} ProcessorCore;
	//TODO(t): union {DWORD NodeNumber;} NumaNode;
	//TODO(t): union
	Cache    CACHE_DESCRIPTOR
	Reserved [2]ULONGLONG
	//TODO(t): end union
}

type LOGICAL_PROCESSOR_RELATIONSHIP enum

const (
	RelationProcessorCore LOGICAL_PROCESSOR_RELATIONSHIP = iota
	RelationNumaNode
	RelationCache
)

type CACHE_DESCRIPTOR struct {
	Level, Associativity BYTE
	LineSize             WORD
	Size                 DWORD
	Type                 PROCESSOR_CACHE_TYPE
}

type PROCESSOR_CACHE_TYPE enum

const (
	CacheUnified PROCESSOR_CACHE_TYPE = iota
	CacheInstruction
	CacheData
	CacheTrace
)

type SYSTEM_INFO struct {
	//TODO(t):UNION OemId DWORD // Obsolete
	//TODO(t):UNION START

	ProcessorArchitecture PROCESSOR_ARCHITECTURE
	_                     WORD
	//TODO(t):UNION END

	PageSize                  DWORD
	MinimumApplicationAddress *VOID
	MaximumApplicationAddress *VOID
	ActiveProcessorMask       DWORD_PTR
	NumberOfProcessors        DWORD
	ProcessorType             PROCESSOR_TYPE
	AllocationGranularity     DWORD
	ProcessorLevel            WORD
	ProcessorRevision         WORD
}

type PROCESSOR_ARCHITECTURE WORD

const (
	PROCESSOR_ARCHITECTURE_INTEL PROCESSOR_ARCHITECTURE = iota
	PROCESSOR_ARCHITECTURE_MIPS
	PROCESSOR_ARCHITECTURE_ALPHA
	PROCESSOR_ARCHITECTURE_PPC
	PROCESSOR_ARCHITECTURE_SHX
	PROCESSOR_ARCHITECTURE_ARM
	PROCESSOR_ARCHITECTURE_IA64
	PROCESSOR_ARCHITECTURE_ALPHA64
	PROCESSOR_ARCHITECTURE_MSIL
	PROCESSOR_ARCHITECTURE_AMD64
	PROCESSOR_ARCHITECTURE_IA32_ON_WIN64
	PROCESSOR_ARCHITECTURE_UNKNOWN PROCESSOR_ARCHITECTURE = 0xFFFF
)

type PROCESSOR_TYPE DWORD

const (
	PROCESSOR_INTEL_386     PROCESSOR_TYPE = 386
	PROCESSOR_INTEL_486     PROCESSOR_TYPE = 486
	PROCESSOR_INTEL_PENTIUM PROCESSOR_TYPE = 586
	PROCESSOR_INTEL_IA64    PROCESSOR_TYPE = 2200
	PROCESSOR_AMD_X8664     PROCESSOR_TYPE = 8664
	PROCESSOR_MIPS_R4000    PROCESSOR_TYPE = 4000
	PROCESSOR_ALPHA_21064   PROCESSOR_TYPE = 21064
	PROCESSOR_PPC_601       PROCESSOR_TYPE = 601
	PROCESSOR_PPC_603       PROCESSOR_TYPE = 603
	PROCESSOR_PPC_604       PROCESSOR_TYPE = 604
	PROCESSOR_PPC_620       PROCESSOR_TYPE = 620
	PROCESSOR_HITACHI_SH3   PROCESSOR_TYPE = 10003
	PROCESSOR_HITACHI_SH3E  PROCESSOR_TYPE = 10004
	PROCESSOR_HITACHI_SH4   PROCESSOR_TYPE = 10005
	PROCESSOR_MOTOROLA_821  PROCESSOR_TYPE = 821
	PROCESSOR_SHx_SH3       PROCESSOR_TYPE = 103
	PROCESSOR_SHx_SH4       PROCESSOR_TYPE = 104
	PROCESSOR_STRONGARM     PROCESSOR_TYPE = 2577
	PROCESSOR_ARM720        PROCESSOR_TYPE = 1824
	PROCESSOR_ARM820        PROCESSOR_TYPE = 2080
	PROCESSOR_ARM920        PROCESSOR_TYPE = 2336
	PROCESSOR_ARM_7TDMI     PROCESSOR_TYPE = 70001
	PROCESSOR_OPTIL         PROCESSOR_TYPE = 0x494f
)

type CONTEXT struct {
	ContextFlags      DWORD
	Dr0               DWORD
	Dr1               DWORD
	Dr2               DWORD
	Dr3               DWORD
	Dr6               DWORD
	Dr7               DWORD
	FloatSave         FLOATING_SAVE_AREA
	SegGs             DWORD
	SegFs             DWORD
	SegEs             DWORD
	SegDs             DWORD
	Edi               DWORD
	Esi               DWORD
	Ebx               DWORD
	Edx               DWORD
	Ecx               DWORD
	Eax               DWORD
	Ebp               DWORD
	Eip               DWORD
	SegCs             DWORD
	EFlags            DWORD
	Esp               DWORD
	SegSs             DWORD
	ExtendedRegisters [MAXIMUM_SUPPORTED_EXTENSION]BYTE
}

const MAXIMUM_SUPPORTED_EXTENSION = 512

type FLOATING_SAVE_AREA struct {
	ControlWord   DWORD
	StatusWord    DWORD
	TagWord       DWORD
	ErrorOffset   DWORD
	ErrorSelector DWORD
	DataOffset    DWORD
	DataSelector  DWORD
	RegisterArea  [SIZE_OF_80387_REGISTERS]BYTE
	Cr0NpxState   DWORD
}

const SIZE_OF_80387_REGISTERS = 80

type LDT_ENTRY struct {
	LimitLow,
	BaseLow WORD
	BaseMid, Flags1, Flags2, BaseHi BYTE
	//TODO(t):Access
	// BaseMid : 8
	// Type : 5
	// Dpl : 2
	// Pres : 1
	// LimitHi : 4
	// Sys : 1
	// Reserved_0 : 1
	// Default_Big : 1
	// Granularity : 1
	// BaseHi : 8
}

type TOKEN_INFORMATION_CLASS enum

const (
	TokenUser TOKEN_INFORMATION_CLASS = iota + 1
	TokenGroups
	TokenPrivileges
	TokenOwner
	TokenPrimaryGroup
	TokenDefaultDacl
	TokenSource
	TokenType
	TokenImpersonationLevel
	TokenStatistics
	TokenRestrictedSids
	TokenSessionId
	TokenGroupsAndPrivileges
	TokenSessionReference
	TokenSandBoxInert
	TokenAuditPolicy
	TokenOrigin
	MaxTokenInfoClass
)

type IO_COUNTERS struct {
	ReadOperationCount  ULONGLONG
	WriteOperationCount ULONGLONG
	OtherOperationCount ULONGLONG
	ReadTransferCount   ULONGLONG
	WriteTransferCount  ULONGLONG
	OtherTransferCount  ULONGLONG
}

type HEAP_INFORMATION_CLASS enum

const (
	HeapCompatibilityInformation HEAP_INFORMATION_CLASS = iota
)

type PROCESS_HEAP_ENTRY struct {
	Data        *VOID
	cData       DWORD
	Overhead    BYTE
	RegionIndex BYTE
	Flags       WORD
	//TODO(t):Union Start "Block"
	// hMem       HANDLE
	// dwReserved [3]DWORD
	//TODO(t):Union End "Block"
	//TODO(t):Union Start "Region"
	CommittedSize   DWORD
	UnCommittedSize DWORD
	FirstBlock      *VOID
	LastBlock       *VOID
	//TODO(t):Union End "Region"
}

type SINGLE_LIST_ENTRY struct {
	Next *SINGLE_LIST_ENTRY
}

// TODO(t):WIN64 ALIGN 16
// type  SLIST_ENTRY struct {
//   Next *SLIST_ENTRY
// }

type SLIST_ENTRY SINGLE_LIST_ENTRY

// TODO(t):WIN64 ALIGN 16
// type SLIST_HEADER struct {
//   Alignment ULONGLONG
//   Region ULONGLONG
// }

type SLIST_HEADER struct {
	// Alignment ULONGLONG
	Next     SLIST_ENTRY
	Depth    WORD
	Sequence WORD
}

//TODO(t):WIN64
//const MAX_NATURAL_ALIGNMENT = 8 sizeof(ULONGLONG)
//const MEMORY_ALLOCATION_ALIGNMENT = 16

const (
	MAX_NATURAL_ALIGNMENT       = 4
	MEMORY_ALLOCATION_ALIGNMENT = 8
)

type QUOTA_LIMITS struct {
	PagedPoolLimit        SIZE_T
	NonPagedPoolLimit     SIZE_T
	MinimumWorkingSetSize SIZE_T
	MaximumWorkingSetSize SIZE_T
	PagefileLimit         SIZE_T
	TimeLimit             LARGE_INTEGER
}

type SID_NAME_USE enum

const (
	SidTypeUser SID_NAME_USE = iota + 1
	SidTypeGroup
	SidTypeDomain
	SidTypeAlias
	SidTypeWellKnownGroup
	SidTypeDeletedAccount
	SidTypeInvalid
	SidTypeUnknown
	SidTypeComputer
)

type JOBOBJECTINFOCLASS enum

const (
	JobObjectBasicAccountingInformation JOBOBJECTINFOCLASS = 1 + iota
	JobObjectBasicLimitInformation
	JobObjectBasicProcessIdList
	JobObjectBasicUIRestrictions
	JobObjectSecurityLimitInformation
	JobObjectEndOfJobTimeInformation
	JobObjectAssociateCompletionPortInformation
	JobObjectBasicAndIoAccountingInformation
	JobObjectExtendedLimitInformation
	JobObjectJobSetInformation
	MaxJobObjectInfoClass
)

type APCFUNC func(Param ULONG_PTR)

type FE_EXPORT_FUNC func(
	Data *BYTE, CallbackContext *VOID, Length ULONG) DWORD

type FE_IMPORT_FUNC func(
	Data *BYTE, CallbackContext *VOID, Length *ULONG) DWORD

type FILE_SEGMENT_ELEMENT struct {
	//TODO(t): Alignment ULONGLONG
	Buffer *VOID64
}

type LATENCY_TIME enum

const (
	LT_DONT_CARE LATENCY_TIME = iota
	LT_LOWEST_LATENCY
)

type COMPUTER_NAME_FORMAT enum

const (
	ComputerNameNetBIOS COMPUTER_NAME_FORMAT = iota
	ComputerNameDnsHostname
	ComputerNameDnsDomain
	ComputerNameDnsFullyQualified
	ComputerNamePhysicalNetBIOS
	ComputerNamePhysicalDnsHostname
	ComputerNamePhysicalDnsDomain
	ComputerNamePhysicalDnsFullyQualified
	ComputerNameMax
)

type MEMORY_BASIC_INFORMATION struct {
	BaseAddress       *VOID
	AllocationBase    *VOID
	AllocationProtect DWORD
	RegionSize        SIZE_T
	State             DWORD
	Protect           DWORD
	Type              DWORD
}

type DLGTEMPLATE struct {
	Style         DWORD
	ExtendedStyle DWORD
	DItems        WORD
	X             SHORT
	Y             SHORT
	nX            SHORT
	nY            SHORT
}

type DLGITEMTEMPLATE struct {
	Style         DWORD
	ExtendedStyle DWORD
	X             SHORT
	Y             SHORT
	nX            SHORT
	nY            SHORT
	Id            WORD
}

type PAINTSTRUCT struct {
	DC        HDC
	Erase     BOOL
	Paint     RECT
	Restore   BOOL
	IncUpdate BOOL
	Reserved  [32]BYTE
}

const (
	CCHDEVICENAME = 32
	CCHFORMNAME   = 32
)

type DEVMODEA struct {
	DeviceName    [CCHDEVICENAME]WChar
	SpecVersion   WORD
	DriverVersion WORD
	Size          WORD
	DriverExtra   WORD
	Fields        DWORD
	// union {
	// struct {
	Orientation   SHORT
	PaperSize     SHORT
	PaperLength   SHORT
	PaperWidth    SHORT
	Scale         SHORT
	Copies        SHORT
	DefaultSource SHORT
	PrintQuality  SHORT
	// } struct
	// struct {
	// Position POINTL
	// DisplayOrientation DWORD
	// DisplayFixedOutput DWORD
	// } stuct
	// } union
	Color       SHORT
	Duplex      SHORT
	YResolution SHORT
	TTOption    SHORT
	Collate     SHORT
	FormName    [CCHFORMNAME]WChar
	LogPixels   WORD
	BitsPerPel  DWORD
	PelsWidth   DWORD
	PelsHeight  DWORD
	//union {
	DisplayFlags DWORD
	//Nup DWORD
	//} union
	DisplayFrequency DWORD
	ICMMethod        DWORD
	ICMIntent        DWORD
	MediaType        DWORD
	DitherType       DWORD
	Reserved1        DWORD
	Reserved2        DWORD
	PanningWidth     DWORD
	PanningHeight    DWORD
}

type DEVMODEW struct {
	DeviceName    [CCHDEVICENAME]WChar
	SpecVersion   WORD
	DriverVersion WORD
	Size          WORD
	DriverExtra   WORD
	Fields        DWORD
	// union {
	// struct {
	Orientation   SHORT
	PaperSize     SHORT
	PaperLength   SHORT
	PaperWidth    SHORT
	Scale         SHORT
	Copies        SHORT
	DefaultSource SHORT
	PrintQuality  SHORT
	// } struct
	// struct {
	// Position POINTL
	// DisplayOrientation DWORD
	// DisplayFixedOutput DWORD
	// } stuct
	// } union
	Color       SHORT
	Duplex      SHORT
	YResolution SHORT
	TTOption    SHORT
	Collate     SHORT
	FormName    [CCHFORMNAME]WChar
	LogPixels   WORD
	BitsPerPel  DWORD
	PelsWidth   DWORD
	PelsHeight  DWORD
	//union {
	DisplayFlags DWORD
	//Nup DWORD
	//} union
	DisplayFrequency DWORD
	ICMMethod        DWORD
	ICMIntent        DWORD
	MediaType        DWORD
	DitherType       DWORD
	Reserved1        DWORD
	Reserved2        DWORD
	PanningWidth     DWORD
	PanningHeight    DWORD
}

type DLGPROC func(HWND, UINT, WPARAM, LPARAM) INT_PTR

type ICONINFO struct {
	Icon     BOOL
	XHotspot DWORD
	YHotspot DWORD
	Mask     HBITMAP
	Color    HBITMAP
}

type BSM_FLAGS DWORD

const (
	BSF_QUERY BSM_FLAGS = 1 << iota
	BSF_IGNORECURRENTTASK
	BSF_FLUSHDISK
	BSF_NOHANG
	BSF_POSTMESSAGE
	BSF_FORCEIFHUNG
	BSF_NOTIMEOUTIFNOTHUNG
	BSF_ALLOWSFW
	BSF_SENDNOTIFYMESSAGE
	BSF_RETURNHDESK
	BSF_LUID
)

type WNDENUMPROC func(HWND, LPARAM) BOOL

type NAMEENUMPROC func(PVString, LPARAM) BOOL

type WINSTAENUMPROC NAMEENUMPROC
type DESKTOPENUMPROC NAMEENUMPROC

type DISPLAY_DEVICEA struct {
	Size         DWORD
	DeviceName   [32]AChar
	DeviceString [128]AChar
	StateFlags   DWORD
	DeviceID     [128]AChar
	DeviceKey    [128]AChar
}

type DISPLAY_DEVICEW struct {
	Size         DWORD
	DeviceName   [32]AChar
	DeviceString [128]AChar
	StateFlags   DWORD
	DeviceID     [128]AChar
	DeviceKey    [128]AChar
}

type DRAWSTATEPROC func(
	hdc HDC, lData LPARAM, wData WPARAM, cx int, cy int) BOOL

type DRAWTEXTPARAMS struct {
	Size        UINT
	TabLength   int
	LeftMargin  int
	RightMargin int
	LengthDrawn UINT
}

type PROPENUMPROCA func(HWND, AString, HANDLE) BOOL
type PROPENUMPROCW func(HWND, WString, HANDLE) BOOL

type PROPENUMPROCEXA func(HWND, AString, HANDLE, ULONG_PTR) BOOL
type PROPENUMPROCEXW func(HWND, WString, HANDLE, ULONG_PTR) BOOL

type FLASHWINFO struct {
	Size    UINT
	Wnd     HWND
	Flags   DWORD
	Count   UINT
	Timeout DWORD
}

type ALTTABINFO struct {
	Size     DWORD
	Items    int
	Columns  int
	Rows     int
	ColFocus int
	RowFocus int
	CXItem   int
	CYItem   int
	Start    POINT
}

type POINT struct {
	X, Y LONG
}

type WNDCLASS struct {
	Style      CLASS_STYLE
	WndProc    WNDPROC
	ClsExtra   int
	WndExtra   int
	Instance   HINSTANCE
	Icon       HICON
	Cursor     HCURSOR
	Background HBRUSH
	MenuName   PVString
	ClassName  PVString
}

type WNDCLASSEX struct {
	Size       UINT
	Style      CLASS_STYLE
	WndProc    WNDPROC
	ClsExtra   int
	WndExtra   int
	Instance   HINSTANCE
	Icon       HICON
	Cursor     HCURSOR
	Background HBRUSH
	MenuName   PVString
	ClassName  PVString
	IconSm     HICON
}

type COMBOBOXINFO struct {
	Size        DWORD
	Item        RECT
	Button      RECT
	StateButton DWORD
	Combo       HWND
	ItemWin     HWND
	List        HWND
}

type CURSORINFO struct {
	Size      DWORD
	Flags     DWORD
	Cursor    HCURSOR
	ScreenPos POINT
}

type GUITHREADINFO struct {
	Size         DWORD
	Flags        DWORD
	ActiveWnd    HWND
	FocusWnd     HWND
	CaptureWnd   HWND
	MenuOwnerWnd HWND
	MoveSizeWnd  HWND
	CaretWnd     HWND
	CaretRect    RECT
}

type LASTINPUTINFO struct {
	Size UINT
	Time DWORD
}

type MENUBARINFO struct {
	Size    DWORD
	Bar     RECT
	Menu    HMENU
	MenuWnd HWND
	Flags   DWORD
	//TODO(t):FIX THIS
	//BarFocused BOOL :1 // in Flags
	//Focused BOOL :1 // in Flags
}

type RECT struct {
	Left   LONG
	Top    LONG
	Right  LONG
	Bottom LONG
}

type MONITORINFO struct {
	Size    DWORD
	Monitor RECT
	Work    RECT
	Flags   DWORD
}

type LARGE_INTEGER struct {
	LowPart  DWORD
	HighPart LONG
	//TODO(t):Union
	//QuadPart LONGLONG
}

const EXCEPTION_MAXIMUM_PARAMETERS = 15

type EXCEPTION_RECORD struct {
	ExceptionCode        DWORD
	ExceptionFlags       DWORD
	ExceptionRecord      *EXCEPTION_RECORD
	ExceptionAddress     *VOID
	NumberParameters     DWORD
	ExceptionInformation [EXCEPTION_MAXIMUM_PARAMETERS]ULONG_PTR
}

type WNDPROC func(HWND, WINDOW_MESSAGE, WPARAM, LPARAM) LRESULT

type MSG struct {
	Wnd     HWND
	Message UINT
	WParam  WPARAM
	LParam  LPARAM
	Time    DWORD
	Pt      POINT
}

type RAWINPUT struct {
	Header RAWINPUTHEADER
	//TODO(t): union { 24
	Mouse RAWMOUSE
	//Keyboard RAWKEYBOARD 16
	//HID      RAWHID 12
	//} union
}

type RAWINPUTHEADER struct {
	Type   RAW_TYPE
	Size   DWORD
	Device HANDLE
	Param  WPARAM
}

type RAW_TYPE DWORD

const (
	RIM_TYPEMOUSE = iota
	RIM_TYPEKEYBOARD
	RIM_TYPEHID
)

type RAWMOUSE struct {
	Flags USHORT
	// union {
	//     ULONG Buttons;
	//     struct  {
	ButtonFlags USHORT
	ButtonData  USHORT
	//     };
	// };
	RawButtons       ULONG
	LastX            LONG
	LastY            LONG
	ExtraInformation ULONG
}

type RAWKEYBOARD struct {
	MakeCode         USHORT
	Flags            USHORT
	Reserved         USHORT
	VKey             USHORT
	Message          UINT
	ExtraInformation ULONG
}

type RAWHID struct {
	SizeHid DWORD
	Count   DWORD
	RawData [1]BYTE
}

type MENUINFO struct {
	Size          DWORD
	Mask          DWORD
	Style         DWORD
	MaxHeight     UINT
	Back          HBRUSH
	ContextHelpID DWORD
	MenuData      ULONG_PTR
}

type MONITORENUMPROC func(HMONITOR, HDC, *RECT, LPARAM) BOOL

type MOUSEMOVEPOINT struct {
	X         int
	Y         int
	Time      DWORD
	ExtraInfo ULONG_PTR
}

type RAWINPUTDEVICELIST struct {
	Device HANDLE
	Type   DWORD
}

type RAWINPUTDEVICE struct {
	UsagePage USHORT
	Usage     USHORT
	Flags     RID_FLAGS
	Target    HWND
}

type RID_FLAGS DWORD

const (
	RIDEV_REMOVE       RID_FLAGS = 1
	RIDEV_CAPTUREMOUSE           = RIDEV_NOHOTKEYS
	RIDEV_EXMODEMASK   RID_FLAGS = 0xF0
)
const (
	RIDEV_EXCLUDE RID_FLAGS = 1<<iota + 4
	RIDEV_PAGEONLY
	RIDEV_NOLEGACY
)
const (
	RIDEV_INPUTSINK RID_FLAGS = 1<<iota + 8
	RIDEV_NOHOTKEYS
	RIDEV_APPKEYS
)

// RIDEV_EXMODE(mode)  ((mode) & RIDEV_EXMODEMASK)

const CCHILDREN_SCROLLBAR = 5

type SCROLLBARINFO struct {
	Size        DWORD
	ScrollBar   RECT
	LineButton  int
	ThumbTop    int
	ThumbBottom int
	reserved    int
	RGState     [CCHILDREN_SCROLLBAR + 1]DWORD
}

//type SB_MESSAGES
//#define SBM_SETPOS                  0x00E0
//#define SBM_GETPOS                  0x00E1
//#define SBM_SETRANGE                0x00E2
//#define SBM_SETRANGEREDRAW          0x00E6
//#define SBM_GETRANGE                0x00E3
//#define SBM_ENABLE_ARROWS           0x00E4
//#define SBM_SETSCROLLINFO           0x00E9
//#define SBM_GETSCROLLINFO           0x00EA
//#define SBM_GETSCROLLBARINFO        0x00EB

type SCROLLINFO struct {
	Size     UINT
	Mask     UINT
	Min      int
	Max      int
	Page     UINT
	Pos      int
	TrackPos int
}

const CCHILDREN_TITLEBAR = 5

type TITLEBARINFO struct {
	Size     DWORD
	TitleBar RECT
	RGState  [CCHILDREN_TITLEBAR + 1]DWORD
}

type WINDOWINFO struct {
	Size           DWORD
	Window         RECT
	Client         RECT
	Style          DWORD
	ExStyle        DWORD
	WindowStatus   DWORD
	XWindowBorders UINT
	YWindowBorders UINT
	WindowType     ATOM
	CreatorVersion WORD
}

type WINDOWPLACEMENT struct {
	Length         UINT
	Flags          WP_FLAGS
	ShowCmd        UINT
	MinPosition    POINT
	MaxPosition    POINT
	NormalPosition RECT
}

type WP_FLAGS UINT

const (
	WPF_SETMINPOSITION WP_FLAGS = 1 << iota
	WPF_RESTORETOMAXIMIZED
	WPF_ASYNCWINDOWPLACEMENT
)

type INPUT_TYPE DWORD

const (
	INPUT_MOUSE INPUT_TYPE = iota
	INPUT_KEYBOARD
	INPUT_HARDWARE
)

type INPUT struct {
	Type INPUT_TYPE
	//TODO(t):union {

	M MOUSEINPUT
	//K KEYBDINPUT
	//H HARDWAREINPUT
	// } union
}

type MOUSEINPUT struct {
	X         LONG
	Y         LONG
	MouseData DWORD
	Flags     DWORD
	Time      DWORD
	ExtraInfo ULONG_PTR
}

type KEYBDINPUT struct {
	Vk        WORD
	Scan      WORD
	Flags     DWORD
	Time      DWORD
	ExtraInfo ULONG_PTR
}

type HARDWAREINPUT struct {
	Msg    DWORD
	ParamL WORD
	ParamH WORD
}

type TOKEN_TYPE enum

const (
	TokenPrimary TOKEN_TYPE = iota + 1
	TokenImpersonation
)

type COMMPROP struct {
	PacketLength       WORD
	PacketVersion      WORD
	ServiceMask        DWORD
	Reserved1          DWORD
	MaxTxQueue         DWORD
	MaxRxQueue         DWORD
	MaxBaud            DWORD
	ProvSubType        DWORD
	ProvCapabilities   DWORD
	SettableParams     DWORD
	SettableBaud       DWORD
	SettableData       WORD
	SettableStopParity WORD
	CurrentTxQueue     DWORD
	CurrentRxQueue     DWORD
	ProvSpec1          DWORD
	ProvSpec2          DWORD
	ProvChar           [1]WChar
}

type TOP_LEVEL_EXCEPTION_FILTER func(
	excepInfo *EXCEPTION_POINTERS) LONG

type TIMERAPCROUTINE func(
	completionRoutine *VOID, howVal DWORD, highVal DWORD)

type DEBUG_EVENT_CODE DWORD

const (
	EXCEPTION_DEBUG_EVENT DEBUG_EVENT_CODE = iota + 1
	CREATE_THREAD_DEBUG_EVENT
	CREATE_PROCESS_DEBUG_EVENT
	EXIT_THREAD_DEBUG_EVENT
	EXIT_PROCESS_DEBUG_EVENT
	LOAD_DLL_DEBUG_EVENT
	UNLOAD_DLL_DEBUG_EVENT
	OUTPUT_DEBUG_STRING_EVENT
	RIP_EVENT
)

type DEBUG_EVENT struct {
	DebugEventCode DEBUG_EVENT_CODE
	ProcessId      DWORD
	ThreadId       DWORD
	//TODO(t):union { 84

	Exception EXCEPTION_DEBUG_INFO
	// CreateThread      CREATE_THREAD_DEBUG_INFO 12
	// CreateProcessInfo CREATE_PROCESS_DEBUG_INFO 40
	// ExitThread        EXIT_THREAD_DEBUG_INFO 4
	// ExitProcess       EXIT_PROCESS_DEBUG_INFO 4
	// LoadDll           LOAD_DLL_DEBUG_INFO 24 4
	// UnloadDll         UNLOAD_DLL_DEBUG_INFO4 4
	// DebugString       OUTPUT_DEBUG_STRING_INFO 8
	// RipInfo           RIP_INFO 8
	// } union
}

type TIMERPROC func(HWND, UINT, UINT_PTR, DWORD)

type BSMINFO struct {
	Size UINT
	Desk HDESK
	Wnd  HWND
	Luid LUID
}

type FVIRT WORD

const (
	FVIRTKEY FVIRT = 1 << iota
	FNOINVERT
	FSHIFT
	FCONTROL
	FALT
)

type ACCEL struct {
	FVirt FVIRT
	Key   WORD
	Cmd   DWORD
}

type MIIM Fake_type_Fix_me

const (
	MIIM_STATE MIIM = 1 << iota
	MIIM_ID
	MIIM_SUBMENU
	MIIM_CHECKMARKS
	MIIM_TYPE
	MIIM_DATA
	MIIM_STRING
	MIIM_BITMAP
	MIIM_FTYPE
)

const (
	HBMMENU_CALLBACK HBITMAP = iota - 1
	HBMMENU_SYSTEM   HBITMAP = iota
	HBMMENU_MBAR_RESTORE
	HBMMENU_MBAR_MINIMIZE
	_
	HBMMENU_MBAR_CLOSE
	HBMMENU_MBAR_CLOSE_D
	HBMMENU_MBAR_MINIMIZE_D
	HBMMENU_POPUP_CLOSE
	HBMMENU_POPUP_RESTORE
	HBMMENU_POPUP_MAXIMIZE
	HBMMENU_POPUP_MINIMIZE
)

type MENUITEMINFO struct {
	Size         UINT
	Mask         UINT
	Type         UINT
	State        UINT
	ID           UINT
	SubMenu      HMENU
	Checked      HBITMAP
	Unchecked    HBITMAP
	ItemData     ULONG_PTR
	TypeData     POVString
	TypeDataSize UINT
	Item         HBITMAP
}

type GRAYSTRINGPROC func(HDC, LPARAM, int) BOOL

type MSGBOXPARAMS struct {
	Size           UINT
	Owner          HWND
	Instance       HINSTANCE
	Text           PVString
	Caption        PVString
	Style          MSGBOX_TYPE
	Icon           PVString
	ContextHelpId  DWORD_PTR
	MsgBoxCallback MSGBOXCALLBACK
	LanguageId     DWORD
}

type MSGBOXCALLBACK func(hi *HELPINFO)

type HELP_CONTEXT int

const (
	HELPINFO_WINDOW HELP_CONTEXT = iota + 1
	HELPINFO_MENUITEM
)

type HELPINFO struct {
	Size        UINT
	ContextType HELP_CONTEXT
	CtrlId      int
	ItemHandle  HANDLE
	ContextId   DWORD_PTR
	MousePos    POINT
}

type SENDASYNCPROC func(HWND, UINT, ULONG_PTR, LRESULT)

type HOOKPROC func(code int, w WPARAM, l LPARAM) LRESULT

type WINEVENTPROC func(
	winEventHook HWINEVENTHOOK,
	event DWORD,
	w HWND,
	object LONG,
	child LONG,
	eventThread DWORD,
	eventTime DWORD)

type WINEVENT_FLAGS DWORD

const (
	WINEVENT_OUTOFCONTEXT  WINEVENT_FLAGS = 0
	WINEVENT_SKIPOWNTHREAD WINEVENT_FLAGS = 1 << iota
	WINEVENT_SKIPOWNPROCESS
	WINEVENT_INCONTEXT
)

type TME_FLAGS DWORD

const (
	TME_HOVER     TME_FLAGS = 0x1
	TME_LEAVE     TME_FLAGS = 0x2
	TME_NONCLIENT TME_FLAGS = 0x10
	TME_QUERY     TME_FLAGS = 0x40000000
	TME_CANCEL    TME_FLAGS = 0x80000000
)

//HOVER_DEFAULT   0xFFFFFFFF

type TRACKMOUSEEVENT struct {
	Size      DWORD
	Flags     DWORD
	Track     HWND
	HoverTime DWORD
}

type TPMPARAMS struct {
	Size    UINT
	Exclude RECT
}

type SIZE struct{ CX, CU INT }

type BLENDFUNCTION struct {
	BlendOp             BYTE
	BlendFlags          BYTE
	SourceConstantAlpha BYTE
	AlphaFormat         BYTE
}

type UPDATELAYEREDWINDOWINFO struct {
	Size   DWORD
	Dst    HDC
	PtDst  *POINT
	PtSize *SIZE
	Src    HDC
	PtSrc  *POINT
	Key    COLORREF
	Func   *BLENDFUNCTION //TODO(t):&?
	Flags  DWORD
	Dirty  *RECT
}

type PALETTEENTRY struct{ Red, Green, Blue, Flags BYTE }

type PIXELFORMATDESCRIPTOR struct {
	Size           WORD
	Version        WORD
	Flags          DWORD
	PixelType      BYTE
	ColorBits      BYTE
	RedBits        BYTE
	RedShift       BYTE
	GreenBits      BYTE
	GreenShift     BYTE
	BlueBits       BYTE
	BlueShift      BYTE
	AlphaBits      BYTE
	AlphaShift     BYTE
	AccumBits      BYTE
	AccumRedBits   BYTE
	AccumGreenBits BYTE
	AccumBlueBits  BYTE
	AccumAlphaBits BYTE
	DepthBits      BYTE
	StencilBits    BYTE
	AuxBuffers     BYTE
	LayerType      BYTE
	Reserved       BYTE
	LayerMask      DWORD
	VisibleMask    DWORD
	DamageMask     DWORD
}

type BITMAP struct {
	Type       LONG
	Width      LONG
	Height     LONG
	WidthBytes LONG
	Planes     WORD
	BitsPixel  WORD
	Bits       *VOID
}

type LOGBRUSH struct {
	Style UINT
	Color COLORREF
	Hatch ULONG_PTR
}

type BITMAPINFOHEADER struct {
	Size          DWORD
	Width         LONG
	Height        LONG
	Planes        WORD
	BitCount      WORD
	Compression   DWORD
	SizeImage     DWORD
	XPelsPerMeter LONG
	YPelsPerMeter LONG
	ClrUsed       DWORD
	ClrImportant  DWORD
}

type BITMAPINFO struct {
	Header BITMAPINFOHEADER
	Colors [1]RGBQUAD
}

type RGBQUAD struct{ Blue, Green, Red, Reserved BYTE }

type LOGPALETTE struct {
	Version    WORD
	NumEntries WORD
	PalEntry   [1]PALETTEENTRY
}

type LOGPEN struct {
	Style UINT
	Width POINT
	Color COLORREF
}

type MFENUMPROC func(
	hdc HDC,
	ht *HANDLETABLE,
	md *METARECORD,
	nObj int,
	param LPARAM) int

type HANDLETABLE struct {
	ObjectHandle [1]HGDIOBJ
}

type METARECORD struct {
	Size     DWORD
	Function WORD
	Parm     [1]WORD
}

type GOBJENUMPROC func(*VOID, LPARAM) int

type XFORM struct {
	M11 FLOAT
	M12 FLOAT
	M21 FLOAT
	M22 FLOAT
	Dx  FLOAT
	Dy  FLOAT
}

type RGNDATAHEADER struct {
	Size    DWORD
	Type    DWORD
	Count   DWORD
	RgnSize DWORD
	Bound   RECT
}

type ABC struct {
	A int
	B UINT
	C int
}

type GLYPHSET struct {
	This            DWORD
	Accel           DWORD
	GlyphsSupported DWORD
	RangesCount     DWORD
	Tanges          [1]WCRANGE
}

type ENHMFENUMPROC func(
	Hdc HDC,
	HT *HANDLETABLE,
	EMR *ENHMETARECORD,
	Handles int,
	Data LPARAM) int

type ENHMETAHEADER struct {
	Type              DWORD
	Size              DWORD
	Bounds            RECTL
	Frame             RECTL
	Signature         DWORD
	Version           DWORD
	Bytes             DWORD
	Records           DWORD
	Handles           WORD
	Reserved          WORD
	DescriptionCount  DWORD
	DescriptionOffset DWORD
	PalEntries        DWORD
	Device            SIZEL
	Millimeters       SIZEL
	PixelFormatSize   DWORD
	PixelFormatOffset DWORD
	OpenGL            DWORD
	Micrometers       SIZEL
}

type EMRSETCOLORADJUSTMENT struct {
	EMR             EMR
	ColorAdjustment COLORADJUSTMENT
}

type COLORADJUSTMENT struct {
	Size            WORD
	Flags           WORD
	IlluminantIndex WORD
	RedGamma        WORD
	GreenGamma      WORD
	BlueGamma       WORD
	ReferenceBlack  WORD
	ReferenceWhite  WORD
	Contrast        SHORT
	Brightness      SHORT
	Colorfulness    SHORT
	RedGreenTint    SHORT
}

/*
#define CA_NEGATIVE                 0x0001
#define CA_LOG_FILTER               0x0002

#define ILLUMINANT_DEVICE_DEFAULT   0
#define ILLUMINANT_A                1
#define ILLUMINANT_B                2
#define ILLUMINANT_C                3
#define ILLUMINANT_D50              4
#define ILLUMINANT_D55              5
#define ILLUMINANT_D65              6
#define ILLUMINANT_D75              7
#define ILLUMINANT_F2               8
#define ILLUMINANT_MAX_INDEX        ILLUMINANT_F2

#define ILLUMINANT_TUNGSTEN         ILLUMINANT_A
#define ILLUMINANT_DAYLIGHT         ILLUMINANT_C
#define ILLUMINANT_FLUORESCENT      ILLUMINANT_F2
#define ILLUMINANT_NTSC             ILLUMINANT_C

#define RGB_GAMMA_MIN               (WORD)02500
#define RGB_GAMMA_MAX               (WORD)65000

#define REFERENCE_WHITE_MIN         (WORD)6000
#define REFERENCE_WHITE_MAX         (WORD)10000
#define REFERENCE_BLACK_MIN         (WORD)0
#define REFERENCE_BLACK_MAX         (WORD)4000

#define COLOR_ADJ_MIN               (SHORT)-100
#define COLOR_ADJ_MAX               (SHORT)100
*/

type KERNINGPAIR struct {
	First      WORD
	Second     WORD
	KernAmount int
}

type WCRANGE struct {
	Low    WChar
	Glyphs USHORT
}

type ENHMETARECORD struct {
	Type DWORD
	Size DWORD
	Parm [1]DWORD
}

type SIZEL struct{ CX, CY LONG }

type RECTL struct {
	Left   LONG
	Top    LONG
	Right  LONG
	Bottom LONG
}

type EMR struct {
	Type DWORD
	Size DWORD
}

type RGNDATA struct {
	RDH    RGNDATAHEADER
	Buffer [1]Char
}

type GLYPHMETRICSFLOAT struct {
	BlackBoxX   FLOAT
	BlackBoxY   FLOAT
	GlyphOrigin POINTFLOAT
	CellIncX    FLOAT
	CellIncY    FLOAT
}

type GLYPHMETRICS struct {
	BlackBoxX   UINT
	BlackBoxY   UINT
	GlyphOrigin POINT
	CellIncX    SHORT
	CellIncY    SHORT
}

type MAT2 struct {
	M11 FIXED
	M12 FIXED
	M21 FIXED
	M22 FIXED
}

type RASTERIZER_STATUS struct {
	Size       SHORT
	Flags      SHORT
	LanguageID SHORT
}

/* bits defined in wFlags of RASTERIZER_STATUS */
//#define TT_AVAILABLE    0x0001
//#define TT_ENABLED      0x0002

type FONTSIGNATURE struct {
	Usb [4]DWORD
	Csb [2]DWORD
}

type CHARSETINFO struct {
	Charset UINT
	ACP     UINT
	FS      FONTSIGNATURE
}

type LINEDDAPROC func(int, int, LPARAM)

type METAFILEPICT struct {
	MM   LONG
	XExt LONG
	YExt LONG
	HMF  HMETAFILE
}

type ABORTPROC func(HDC, int) BOOL

type FIXED struct {
	Fract WORD
	Value SHORT
}

type POINTFLOAT struct{ X, Y FLOAT }

const LF_FACESIZE = 32

type LOGFONTA struct {
	Height         LONG
	Width          LONG
	Escapement     LONG
	Orientation    LONG
	Weight         LONG
	Italic         BYTE
	Underline      BYTE
	StrikeOut      BYTE
	CharSet        BYTE
	OutPrecision   BYTE
	ClipPrecision  BYTE
	Quality        BYTE
	PitchAndFamily BYTE
	FaceName       [LF_FACESIZE]AChar
}

type LOGFONTW struct {
	Height         LONG
	Width          LONG
	Escapement     LONG
	Orientation    LONG
	Weight         LONG
	Italic         BYTE
	Underline      BYTE
	StrikeOut      BYTE
	CharSet        BYTE
	OutPrecision   BYTE
	ClipPrecision  BYTE
	Quality        BYTE
	PitchAndFamily BYTE
	FaceName       [LF_FACESIZE]WChar
}

type ENUMLOGFONTA struct {
	LogFont  LOGFONTA
	FullName [LF_FULLFACESIZE]AChar
	Style    [LF_FACESIZE]AChar
}

type ENUMLOGFONTW struct {
	LogFont  LOGFONTW
	FullName [LF_FULLFACESIZE]WChar
	Style    [LF_FACESIZE]WChar
}

const LF_FULLFACESIZE = 64

type FONTENUMPROCA func(*LOGFONTA, *TEXTMETRICA, DWORD, LPARAM) int

type FONTENUMPROCW func(*LOGFONTW, *TEXTMETRICW, DWORD, LPARAM) int

type TEXTMETRICA struct {
	Height           LONG
	Ascent           LONG
	Descent          LONG
	InternalLeading  LONG
	ExternalLeading  LONG
	AveCharWidth     LONG
	MaxCharWidth     LONG
	Weight           LONG
	Overhang         LONG
	DigitizedAspectX LONG
	DigitizedAspectY LONG
	FirstChar        AChar
	LastChar         AChar
	DefaultChar      AChar
	BreakChar        AChar
	Italic           BYTE
	Underlined       BYTE
	StruckOut        BYTE
	PitchAndFamily   BYTE
	CharSet          BYTE
}

type TEXTMETRICW struct {
	Height           LONG
	Ascent           LONG
	Descent          LONG
	InternalLeading  LONG
	ExternalLeading  LONG
	AveCharWidth     LONG
	MaxCharWidth     LONG
	Weight           LONG
	Overhang         LONG
	DigitizedAspectX LONG
	DigitizedAspectY LONG
	FirstChar        WChar
	LastChar         WChar
	DefaultChar      WChar
	BreakChar        WChar
	Italic           BYTE
	Underlined       BYTE
	StruckOut        BYTE
	PitchAndFamily   BYTE
	CharSet          BYTE
}

type PROC func() int //TODO(t):Win64 INT_PTR

type ABCFLOAT struct{ A, B, C FLOAT }

type OUTLINETEXTMETRICA struct {
	Size               UINT
	TextMetrics        TEXTMETRICA
	Filler             BYTE
	PanoseNumber       PANOSE
	Selection          UINT
	Type               UINT
	CharSlopeRise      int
	CharSlopeRun       int
	ItalicAngle        int
	EMSquare           UINT
	Ascent             int
	Descent            int
	LineGap            UINT
	CapEmHeight        UINT
	XHeight            UINT
	FontBox            RECT
	MacAscent          int
	MacDescent         int
	MacLineGap         UINT
	MinimumPPEM        UINT
	SubscriptSize      POINT
	SubscriptOffset    POINT
	SuperscriptSize    POINT
	SuperscriptOffset  POINT
	StrikeoutSize      UINT
	StrikeoutPosition  int
	UnderscoreSize     int
	UnderscorePosition int
	FamilyName         AString
	FaceName           AString
	StyleName          AString
	FullName           AString
}

type OUTLINETEXTMETRICW struct {
	Size               UINT
	TextMetrics        TEXTMETRICW
	Filler             BYTE
	PanoseNumber       PANOSE
	Selection          UINT
	Type               UINT
	CharSlopeRise      int
	CharSlopeRun       int
	ItalicAngle        int
	EMSquare           UINT
	Ascent             int
	Descent            int
	LineGap            UINT
	CapEmHeight        UINT
	XHeight            UINT
	FontBox            RECT
	MacAscent          int
	MacDescent         int
	MacLineGap         UINT
	MinimumPPEM        UINT
	SubscriptSize      POINT
	SubscriptOffset    POINT
	SuperscriptSize    POINT
	SuperscriptOffset  POINT
	StrikeoutSize      UINT
	StrikeoutPosition  int
	UnderscoreSize     int
	UnderscorePosition int
	FamilyName         WString
	FaceName           WString
	StyleName          WString
	FullName           WString
}

type GCP_RESULTS struct {
	StructSize  DWORD
	OutString   POVString
	Order       *UINT
	Dx          *int
	CaretPos    *int
	Class       AString
	Glyphs      POVString
	GlyphsCount UINT
	MaxFit      int
}

type ENUMLOGFONTEXDVA struct {
	EnumLogfontEx ENUMLOGFONTEXA
	DesignVector  DESIGNVECTOR
}
type ENUMLOGFONTEXDVW struct {
	EnumLogfontEx ENUMLOGFONTEXW
	DesignVector  DESIGNVECTOR
}

type LOGCOLORSPACEA struct {
	Signature  DWORD
	Version    DWORD
	Size       DWORD
	CSType     LCSCSTYPE
	Intent     LCSGAMUTMATCH
	Endpoints  CIEXYZTRIPLE
	GammaRed   DWORD
	GammaGreen DWORD
	GammaBlue  DWORD
	Filename   [MAX_PATH]AChar
}
type LOGCOLORSPACEW struct {
	Signature  DWORD
	Version    DWORD
	Size       DWORD
	CSType     LCSCSTYPE
	Intent     LCSGAMUTMATCH
	Endpoints  CIEXYZTRIPLE
	GammaRed   DWORD
	GammaGreen DWORD
	GammaBlue  DWORD
	Filename   [MAX_PATH]WChar
}

type PANOSE struct {
	FamilyType      BYTE
	SerifStyle      BYTE
	Weight          BYTE
	Proportion      BYTE
	Contrast        BYTE
	StrokeVariation BYTE
	ArmStyle        BYTE
	Letterform      BYTE
	Midline         BYTE
	XHeight         BYTE
}

type ENUMLOGFONTEXA struct {
	LogFont  LOGFONTA
	FullName [LF_FULLFACESIZE]AChar
	Style    [LF_FACESIZE]AChar
	Script   [LF_FACESIZE]AChar
}
type ENUMLOGFONTEXW struct {
	LogFont  LOGFONTW
	FullName [LF_FULLFACESIZE]WChar
	Style    [LF_FACESIZE]WChar
	Script   [LF_FACESIZE]WChar
}

const MM_MAX_NUMAXES = 16

type DESIGNVECTOR struct {
	Reserved DWORD
	NumAxes  DWORD
	Values   [MM_MAX_NUMAXES]LONG
}

type CIEXYZTRIPLE struct{ Red, Green, Blue CIEXYZ }

type CIEXYZ struct{ X, Y, Z FXPT2DOT30 }

type DOCINFO struct {
	Size     int
	DocName  PVString
	Output   PVString
	Datatype PVString
	Type     DWORD
}

type POLYTEXT struct {
	X     int
	Y     int
	n     UINT
	Str   PVString
	Flags UINT
	Rect  RECT
	DX    *int
}

type ICMENUMPROCA func(AString, LPARAM) int

type ICMENUMPROCW func(WString, LPARAM) int

type LAYERPLANEDESCRIPTOR struct {
	Size           WORD
	Version        WORD
	Flags          DWORD
	PixelType      BYTE
	ColorBits      BYTE
	RedBits        BYTE
	RedShift       BYTE
	GreenBits      BYTE
	GreenShift     BYTE
	BlueBits       BYTE
	BlueShift      BYTE
	AlphaBits      BYTE
	AlphaShift     BYTE
	AccumBits      BYTE
	AccumRedBits   BYTE
	AccumGreenBits BYTE
	AccumBlueBits  BYTE
	AccumAlphaBits BYTE
	DepthBits      BYTE
	StencilBits    BYTE
	AuxBuffers     BYTE
	LayerPlane     BYTE
	Reserved       BYTE
	Transparent    COLORREF
}

type WGLSWAP struct {
	DC    HDC
	Flags UINT
}

type TRIVERTEX struct {
	X     LONG
	Y     LONG
	Red   COLOR16
	Green COLOR16
	Blue  COLOR16
	Alpha COLOR16
}

type WGL_SWAP_FLAG UINT

const (
	WGL_SWAP_MAIN_PLANE WGL_SWAP_FLAG = 1 << iota
	WGL_SWAP_OVERLAY1
	WGL_SWAP_OVERLAY2
	WGL_SWAP_OVERLAY3
	WGL_SWAP_OVERLAY4
	WGL_SWAP_OVERLAY5
	WGL_SWAP_OVERLAY6
	WGL_SWAP_OVERLAY7
	WGL_SWAP_OVERLAY8
	WGL_SWAP_OVERLAY9
	WGL_SWAP_OVERLAY10
	WGL_SWAP_OVERLAY11
	WGL_SWAP_OVERLAY12
	WGL_SWAP_OVERLAY13
	WGL_SWAP_OVERLAY14
	WGL_SWAP_OVERLAY15
	WGL_SWAP_UNDERLAY1
	WGL_SWAP_UNDERLAY2
	WGL_SWAP_UNDERLAY3
	WGL_SWAP_UNDERLAY4
	WGL_SWAP_UNDERLAY5
	WGL_SWAP_UNDERLAY6
	WGL_SWAP_UNDERLAY7
	WGL_SWAP_UNDERLAY8
	WGL_SWAP_UNDERLAY9
	WGL_SWAP_UNDERLAY10
	WGL_SWAP_UNDERLAY11
	WGL_SWAP_UNDERLAY12
	WGL_SWAP_UNDERLAY13
	WGL_SWAP_UNDERLAY14
	WGL_SWAP_UNDERLAY15
)

type FNCALLBACK func(
	Type UINT,
	Fmt UINT,
	Conv HCONV,
	Sz1 HSZ,
	Sz2 HSZ,
	Data HDDEDATA,
	Data1 ULONG_PTR,
	Data2 ULONG_PTR) HDDEDATA

type SECURITY_QUALITY_OF_SERVICE struct {
	Length              DWORD
	ImpersonationLevel  SECURITY_IMPERSONATION_LEVEL
	ContextTrackingMode SECURITY_CONTEXT_TRACKING_MODE
	EffectiveOnly       BOOLEAN
}

type CONVCONTEXT struct {
	Size      UINT
	Flags     UINT
	CountryID UINT
	CodePage  int
	LangID    DWORD
	Security  DWORD
	QoS       SECURITY_QUALITY_OF_SERVICE
}

type CONVINFO struct {
	Size        DWORD
	User        DWORD_PTR
	ConvPartner HCONV
	SvcPartner  HSZ
	ServiceReq  HSZ
	Topic       HSZ
	Item        HSZ
	Fmt         UINT
	Type        UINT
	Status      UINT
	Convst      UINT
	LastError   UINT
	ConvList    HCONVLIST
	ConvCtxt    CONVCONTEXT
	Wnd         HWND
	WndPartner  HWND
}

type IMAGELISTDRAWPARAMS struct {
	Size    DWORD
	Iml     HIMAGELIST
	I       int
	Dst     HDC
	X       int
	Y       int
	CX      int
	CY      int
	XBitmap int
	YBitmap int
	Bk      COLORREF
	Fg      COLORREF
	Style   UINT
	Rop     DWORD
	State   DWORD
	Frame   DWORD
	Effect  COLORREF
}

type INITCOMMONCONTROLSEX struct {
	Size DWORD
	ICC  ICC_CLASSES
}

type ICC_CLASSES DWORD

const (
	ICC_LISTVIEW_CLASSES ICC_CLASSES = 1 << iota
	ICC_TREEVIEW_CLASSES
	ICC_BAR_CLASSES
	ICC_TAB_CLASSES
	ICC_UPDOWN_CLASS
	ICC_PROGRESS_CLASS
	ICC_HOTKEY_CLASS
	ICC_ANIMATE_CLASS
	ICC_DATE_CLASSES
	ICC_USEREX_CLASSES
	ICC_COOL_CLASSES
	ICC_INTERNET_CLASSES
	ICC_PAGESCROLLER_CLASS
	ICC_NATIVEFNTCTL_CLASS
	ICC_STANDARD_CLASSES
	ICC_LINK_CLASS
)

type IMAGEINFO struct {
	Image   HBITMAP
	Mask    HBITMAP
	Unused1 int
	Unused2 int
	Rect    RECT
}

type TBBUTTON struct {
	Bitmap  int
	Command int
	State   BYTE
	Style   BYTE
	_       [2]BYTE
	//TODO(t): Win64 _ is 6

	Data   DWORD_PTR
	String INT_PTR
}

type COLORMAP struct {
	From COLORREF
	To   COLORREF
}

type FNDPAENUMCALLBACK func(p *VOID, data *VOID) int

type FNDSAENUMCALLBACK func(p *VOID, data *VOID) int

type FNDPACOMPARE func(p1 *VOID, p2 *VOID, lParam LPARAM) int

type SUBCLASSPROC func(
	win HWND,
	msg UINT,
	w WPARAM,
	l LPARAM,
	subclass UINT_PTR,
	refData DWORD_PTR) LRESULT

const MAXPNAMELEN = 32

type AUXCAPSA struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]AChar
	Technology    WORD
	_             WORD
	Support       DWORD
}

type AUXCAPSW struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]WChar
	Technology    WORD
	_             WORD
	Support       DWORD
}

type MIDIINCAPSA struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]AChar
	Support       DWORD
}

type MIDIINCAPSW struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]WChar
	Support       DWORD
}

type MIXERCONTROLDETAILS struct {
	Struct    DWORD
	ControlID DWORD
	Channels  DWORD
	//TODO(t):union {

	Owner HWND
	//cMultipleItems DWORD
	//   } union

	DetailsCount DWORD
	Details      *VOID
}

type MMIOPROC func(
	mmioinfo AString, Msg UINT, Param1, Param2 LPARAM) LRESULT

type WAVEINCAPSA struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]AChar
	Formats       DWORD
	Channels      WORD
	_             WORD
}
type WAVEINCAPSW struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]WChar
	Formats       DWORD
	Channels      WORD
	_             WORD
}

type MMIOINFO struct {
	Flags       DWORD
	Fourcc      FOURCC
	IOProc      *MMIOPROC
	ErrorRet    UINT
	Task        HTASK
	BufferCount LONG
	Buffer      *Char
	Next        *Char
	EndRead     *Char
	EndWrite    *Char
	BufOffset   LONG
	DiskOffset  LONG
	Info        [3]DWORD
	_           DWORD
	_           DWORD
	Mmio        HMMIO
}

type WAVEOUTCAPSA struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]AChar
	Formats       DWORD
	Channels      WORD
	_             WORD
	Support       DWORD
}

type WAVEOUTCAPSW struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]WChar
	Formats       DWORD
	Channels      WORD
	_             WORD
	Support       DWORD
}

type MIDIOUTCAPSA struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]AChar
	Technology    WORD
	Voices        WORD
	Notes         WORD
	ChannelMask   WORD
	Support       DWORD
}

type MIDIOUTCAPSW struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]WChar
	Technology    WORD
	Voices        WORD
	Notes         WORD
	ChannelMask   WORD
	Support       DWORD
}

type MIXERCAPSA struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]AChar
	Support       DWORD
	Destinations  DWORD
}

type MIXERCAPSW struct {
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]WChar
	Support       DWORD
	Destinations  DWORD
}

const MIXER_SHORT_NAME_CHARS = 16

const MIXER_LONG_NAME_CHARS = 64

type MIXERLINEA struct {
	Struct        DWORD
	Destination   DWORD
	Source        DWORD
	LineID        DWORD
	Line          DWORD
	User          DWORD_PTR
	ComponentType DWORD
	Channels      DWORD
	Connections   DWORD
	Controls      DWORD
	ShortName     [MIXER_SHORT_NAME_CHARS]AChar
	Name          [MIXER_LONG_NAME_CHARS]AChar
	// TODO(t):struct {

	Type          DWORD
	DeviceID      DWORD
	Mid           WORD
	Pid           WORD
	DriverVersion MMVERSION
	Pname         [MAXPNAMELEN]AChar
	// } Target

}

type MIXERLINEW struct {
	cbStruct        DWORD
	dwDestination   DWORD
	dwSource        DWORD
	dwLineID        DWORD
	fdwLine         DWORD
	dwUser          DWORD_PTR
	dwComponentType DWORD
	cChannels       DWORD
	cConnections    DWORD
	cControls       DWORD
	szShortName     [MIXER_SHORT_NAME_CHARS]WChar
	szName          [MIXER_LONG_NAME_CHARS]WChar
	// TODO(t):struct {

	dwType         DWORD
	dwDeviceID     DWORD
	wMid           WORD
	wPid           WORD
	vDriverVersion MMVERSION
	szPname        [MAXPNAMELEN]WChar
	// } Target

}

const MAX_JOYSTICKOEMVXDNAME = 260

type JOYCAPSA struct {
	Mid        WORD
	Pid        WORD
	Pname      [MAXPNAMELEN]AChar
	Xmin       UINT
	Xmax       UINT
	Ymin       UINT
	Ymax       UINT
	Zmin       UINT
	Zmax       UINT
	NumButtons UINT
	PeriodMin  UINT
	PeriodMax  UINT
	Rmin       UINT
	Rmax       UINT
	Umin       UINT
	Umax       UINT
	Vmin       UINT
	Vmax       UINT
	Caps       UINT
	MaxAxes    UINT
	NumAxes    UINT
	MaxButtons UINT
	RegKey     [MAXPNAMELEN]AChar
	OEMVxD     [MAX_JOYSTICKOEMVXDNAME]AChar
}

type JOYCAPSW struct {
	Mid        WORD
	Pid        WORD
	Pname      [MAXPNAMELEN]WChar
	Xmin       UINT
	Xmax       UINT
	Ymin       UINT
	Ymax       UINT
	Zmin       UINT
	Zmax       UINT
	NumButtons UINT
	PeriodMin  UINT
	PeriodMax  UINT
	Rmin       UINT
	Rmax       UINT
	Umin       UINT
	Umax       UINT
	Vmin       UINT
	Vmax       UINT
	Caps       UINT
	MaxAxes    UINT
	NumAxes    UINT
	MaxButtons UINT
	RegKey     [MAXPNAMELEN]WChar
	OEMVxD     [MAX_JOYSTICKOEMVXDNAME]WChar
}

type JOYINFO struct {
	Xpos    UINT
	Ypos    UINT
	Zpos    UINT
	Buttons UINT
}

type JOYINFOEX struct {
	Size         DWORD
	Flags        DWORD
	Xpos         DWORD
	Ypos         DWORD
	Zpos         DWORD
	Rpos         DWORD
	Upos         DWORD
	Vpos         DWORD
	Buttons      DWORD
	ButtonNumber DWORD
	POV          DWORD
	_            DWORD
	_            DWORD
}

type MIDIHDR struct {
	Data          AString
	BufferLength  DWORD
	BytesRecorded DWORD
	User          DWORD_PTR
	Flags         DWORD
	Next          *MIDIHDR
	_             DWORD_PTR
	Offset        DWORD
	_             [8]DWORD_PTR
}

type MMTIME struct {
	Type UINT
	//TODO(t):union {

	Ms     DWORD
	Sample DWORD
	Cb     DWORD
	Ticks  DWORD
	//TODO(t): struct {

	Hour  BYTE
	Min   BYTE
	Sec   BYTE
	Frame BYTE
	Fps   BYTE
	Dummy BYTE
	Pad   [2]BYTE
	//TODO(t): } smpte
	//struct{

	songptrpos DWORD
	//} midi
	//} union
}

type MMCKINFO struct {
	Id         FOURCC
	Size       DWORD
	Type       FOURCC
	DataOffset DWORD
	Flags      DWORD
}

type TIMECAPS struct {
	PeriodMin UINT
	PeriodMax UINT
}

type WAVEHDR struct {
	Data          AString
	BufferLength  DWORD
	BytesRecorded DWORD
	User          DWORD_PTR
	Flags         DWORD
	Loops         DWORD
	Next          *WAVEHDR
	_             DWORD_PTR
}

type WAVEFORMATEX struct {
	FormatTag      WORD
	Channels       WORD
	SamplesPerSec  DWORD
	AvgBytesPerSec DWORD
	BlockAlign     WORD
	BitsPerSample  WORD
	Size           WORD
}

type TIMECALLBACK func(
	timerId, msg UINT, urser, dw1, dw2 DWORD_PTR)

type YIELDPROC func(mciId MCIDEVICEID, yieldData DWORD) UINT

const NCBNAMSZ = 16

type NCB struct {
	Command  UCHAR
	Retcode  UCHAR
	Lsn      UCHAR
	Num      UCHAR
	Buffer   *UCHAR
	Length   WORD
	Callname [NCBNAMSZ]UCHAR
	Name     [NCBNAMSZ]UCHAR
	Rto      UCHAR
	Sto      UCHAR
	Post     func(*NCB)
	Lana_num UCHAR
	Cmd_cplt UCHAR
	_        [10]UCHAR
	Event    HANDLE
}

type SHCREATEPROCESSINFOW struct {
	Size              DWORD
	Mask              ULONG
	Wnd               HWND
	File              WString
	Parameters        WString
	CurrentDirectory  WString
	UserToken         HANDLE
	ProcessAttributes *SECURITY_ATTRIBUTES
	ThreadAttributes  *SECURITY_ATTRIBUTES
	InheritHandles    BOOL
	CreationFlags     DWORD
	StartupInfo       *STARTUPINFO
	//TODO(t):STARTUPINFOW

	ProcessInformation *PROCESS_INFORMATION
}

type SHFILEINFOA struct {
	HIcon       HICON
	Icon        int
	Attributes  DWORD
	DisplayName [MAX_PATH]AChar
	TypeName    [80]AChar
}

type SHFILEINFOW struct {
	HIcon       HICON
	Icon        int
	Attributes  DWORD
	DisplayName [MAX_PATH]WChar
	TypeName    [80]WChar
}

type ENUM_SERVICE_STATUS struct {
	ServiceName   PVString
	DisplayName   PVString
	ServiceStatus SERVICE_STATUS
}

type SERVICE_STATUS struct {
	ServiceType             DWORD
	CurrentState            DWORD
	ControlsAccepted        DWORD
	Win32ExitCode           DWORD
	ServiceSpecificExitCode DWORD
	CheckPoint              DWORD
	WaitHint                DWORD
}

type QUERY_SERVICE_CONFIG struct {
	ServiceType      DWORD
	StartType        DWORD
	ErrorControl     DWORD
	BinaryPathName   PVString
	LoadOrderGroup   PVString
	TagId            DWORD
	Dependencies     PVString
	ServiceStartName PVString
	DisplayName      PVString
}

type QUERY_SERVICE_LOCK_STATUS struct {
	IsLocked     DWORD
	LockOwner    PVString
	LockDuration DWORD
}

type SERVICE_TABLE_ENTRY struct {
	ServiceName PVString
	ServiceProc *SERVICE_MAIN_FUNCTION
	//TODO(t):*SERVICE_MAIN_FUNCTIONW
}

type SERVICE_MAIN_FUNCTION func(
	NumServicesArgs DWORD, ServiceArgVectors *VString)

//TODO(t):*VString

type HANDLER_FUNCTION func(Control DWORD)

type HANDLER_FUNCTION_EX func(
	Control, EventType DWORD, EventData, Context *VOID)

type SC_ENUM_TYPE enum

const SC_ENUM_PROCESS_INFO SC_ENUM_TYPE = 0

type SC_STATUS_TYPE enum

const SC_STATUS_PROCESS_INFO SC_STATUS_TYPE = 0

const WSADESCRIPTION_LEN = 256
const WSASYS_STATUS_LEN = 128

type WSADATA struct {
	Version      WORD
	HighVersion  WORD
	Description  [WSADESCRIPTION_LEN + 1]AChar
	SystemStatus [WSASYS_STATUS_LEN + 1]AChar
	MaxSockets   Unsigned_short
	MaxUdpDg     Unsigned_short
	VendorInfo   AString
	//TODO(t): WIN64
	//MaxSockets   Unsigned_short
	//MaxUdpDg     Unsigned_short
	//VendorInfo   AString
	//Description  [WSADESCRIPTION_LEN + 1]AChar
	//SystemStatus [WSASYS_STATUS_LEN + 1]AChar
}

type TRANSMIT_FILE_BUFFERS struct {
	Head       *VOID
	HeadLength DWORD
	Tail       *VOID
	TailLength DWORD
}

type VALENT struct {
	Valuename PVString
	Valuelen  DWORD
	Valueptr  DWORD_PTR
	Type      DWORD
}

const MAX_LEADBYTES = 12
const MAX_DEFAULTCHAR = 2

type CPINFO struct {
	MaxCharSize UINT
	DefaultChar [MAX_DEFAULTCHAR]BYTE
	LeadByte    [MAX_LEADBYTES]BYTE
}

type CPINFOEXA struct {
	MaxCharSize        UINT
	DefaultChar        [MAX_DEFAULTCHAR]BYTE
	LeadByte           [MAX_LEADBYTES]BYTE
	UnicodeDefaultChar WChar
	CodePage           UINT
	CodePageName       [MAX_PATH]AChar
}

type CPINFOEXW struct {
	MaxCharSize        UINT
	DefaultChar        [MAX_DEFAULTCHAR]BYTE
	LeadByte           [MAX_LEADBYTES]BYTE
	UnicodeDefaultChar WChar
	CodePage           UINT
	CodePageName       [MAX_PATH]WChar
}

type NUMBERFMT struct {
	NumDigits     UINT
	LeadingZero   UINT
	Grouping      UINT
	DecimalSep    PVString
	ThousandSep   PVString
	NegativeOrder UINT
}

type CURRENCYFMT struct {
	NumDigits      UINT
	LeadingZero    UINT
	Grouping       UINT
	DecimalSep     PVString
	ThousandSep    PVString
	NegativeOrder  UINT
	PositiveOrder  UINT
	CurrencySymbol PVString
}

type CALINFO_ENUMPROC func(PVString) BOOL

type CALINFO_ENUMPROCEX func(PVString, CALID) BOOL

type DATEFMT_ENUMPROC func(PVString) BOOL

type DATEFMT_ENUMPROCEX func(PVString, CALID) BOOL

type TIMEFMT_ENUMPROC func(PVString) BOOL

//TODO(t):*PVString not handled yet (multiple)
type LANGUAGEGROUP_ENUMPROC func(LGRPID, *PVString, *PVString, DWORD, LONG_PTR) BOOL

type LANGGROUPLOCALE_ENUMPROC func(LGRPID, LCID, *PVString, LONG_PTR) BOOL

type UILANGUAGE_ENUMPROC func(PVString, LONG_PTR) BOOL

type LOCALE_ENUMPROC func(PVString) BOOL

type CODEPAGE_ENUMPROC func(PVString) BOOL

type GEO_ENUMPROC func(GEOID) BOOL

type NLSVERSIONINFO struct {
	NLSVersionInfoSize DWORD
	NLSVersion         DWORD
	DefinedVersion     DWORD
}

type NETRESOURCE struct {
	Scope       DWORD
	Type        DWORD
	DisplayType DWORD
	Usage       DWORD
	LocalName   PVString
	RemoteName  PVString
	Comment     PVString
	Provider    PVString
}

type NETINFOSTRUCT struct {
	Structure       DWORD
	ProviderVersion DWORD
	Status          DWORD
	Characteristics DWORD
	Handle          ULONG_PTR
	NetType         WORD
	Printers        DWORD
	Drives          DWORD
}

type NETCONNECTINFOSTRUCT struct {
	Structure   DWORD
	Flags       DWORD
	Speed       DWORD
	Delay       DWORD
	OptDataSize DWORD
}

type INPUT_RECORD struct {
	EventType INPUT_EVENT_TYPE
	//TODO(t):union {

	KeyEvent              KEY_EVENT_RECORD
	MouseEvent            MOUSE_EVENT_RECORD
	WindowBufferSizeEvent WINDOW_BUFFER_SIZE_RECORD
	MenuEvent             MENU_EVENT_RECORD
	FocusEvent            FOCUS_EVENT_RECORD
	//} Event
}

type INPUT_EVENT_TYPE WORD

const (
	KEY_EVENT INPUT_EVENT_TYPE = 1 << iota
	MOUSE_EVENT
	WINDOW_BUFFER_SIZE_EVENT
	MENU_EVENT
	FOCUS_EVENT
)

type KEY_EVENT_RECORD struct {
	KeyDown         BOOL
	RepeatCount     WORD
	VirtualKeyCode  WORD
	VirtualScanCode WORD
	//TODO(t):union {

	UnicodeChar WChar
	//AsciiChar   CHAR
	//} uChar;

	ControlKeyState CONTROL_KEY_STATE
}

type CONTROL_KEY_STATE DWORD

const (
	RIGHT_ALT_PRESSED CONTROL_KEY_STATE = 1 << iota
	LEFT_ALT_PRESSED
	RIGHT_CTRL_PRESSED
	LEFT_CTRL_PRESSED
	SHIFT_PRESSED
	NUMLOCK_ON
	SCROLLLOCK_ON
	CAPSLOCK_ON
	ENHANCED_KEY
)
const (
	NLS_DBCSCHAR CONTROL_KEY_STATE = 0x10000 << iota
	NLS_KATAKANA
	NLS_HIRAGANA
	NLS_ROMAN CONTROL_KEY_STATE = 0x400000 << iota
	NLS_IME_CONVERSION
	NLS_ALPHANUMERIC CONTROL_KEY_STATE = 0x00000000
	NLS_IME_DISABLE  CONTROL_KEY_STATE = 0x20000000
)

type MOUSE_EVENT_RECORD struct {
	MousePosition   COORD
	ButtonState     MOUSE_BUTTON_STATE
	ControlKeyState CONTROL_KEY_STATE
	EventFlags      MOUSE_EVENT_FLAGS
}

type MOUSE_BUTTON_STATE DWORD

const (
	FROM_LEFT_1ST_BUTTON_PRESSED MOUSE_BUTTON_STATE = 1 << iota
	RIGHTMOST_BUTTON_PRESSED
	FROM_LEFT_2ND_BUTTON_PRESSED
	FROM_LEFT_3RD_BUTTON_PRESSED
	FROM_LEFT_4TH_BUTTON_PRESSED
)

type MOUSE_EVENT_FLAGS DWORD

const (
	MOUSE_MOVED MOUSE_EVENT_FLAGS = 1 << iota
	DOUBLE_CLICK
	MOUSE_WHEELED
)

type WINDOW_BUFFER_SIZE_RECORD struct {
	Size COORD
}

type MENU_EVENT_RECORD struct {
	CommandId UINT
}

type FOCUS_EVENT_RECORD struct {
	SetFocus BOOL
}

type COORD struct {
	X SHORT
	Y SHORT
}

type CHAR_INFO struct {
	//TODO(t):union {

	UnicodeChar WChar
	// CHAR   AsciiChar;
	//} Char;

	Attributes WORD
}

type SMALL_RECT struct {
	Left   SHORT
	Top    SHORT
	Right  SHORT
	Bottom SHORT
}

type CONSOLE_SCREEN_BUFFER_INFO struct {
	Size              COORD
	CursorPosition    COORD
	Attributes        CONSOLE_SCREEN_ATTRIBUTES
	Window            SMALL_RECT
	MaximumWindowSize COORD
}

type CONSOLE_SCREEN_ATTRIBUTES WORD

const (
	FOREGROUND_BLUE CONSOLE_SCREEN_ATTRIBUTES = 1 << iota
	FOREGROUND_GREEN
	FOREGROUND_RED
	FOREGROUND_INTENSITY
	BACKGROUND_BLUE
	BACKGROUND_GREEN
	BACKGROUND_RED
	BACKGROUND_INTENSITY
	COMMON_LVB_LEADING_BYTE
	COMMON_LVB_TRAILING_BYTE
	COMMON_LVB_GRID_HORIZONTAL
	COMMON_LVB_GRID_LVERTICAL
	COMMON_LVB_GRID_RVERTICAL
	_
	COMMON_LVB_REVERSE_VIDEO
	COMMON_LVB_UNDERSCORE
	COMMON_LVB_SBCSDBCS = COMMON_LVB_LEADING_BYTE | COMMON_LVB_TRAILING_BYTE
)

type CONSOLE_CURSOR_INFO struct {
	Size    DWORD
	Visible BOOL
}

type CONSOLE_FONT_INFO struct {
	Font     DWORD
	FontSize COORD
}

type CONSOLE_SELECTION_INFO struct {
	Flags           CONSOLE_SELECTION_FLAGS
	SelectionAnchor COORD
	Selection       SMALL_RECT
}

type CONSOLE_SELECTION_FLAGS DWORD

const (
	CONSOLE_SELECTION_IN_PROGRESS CONSOLE_SELECTION_FLAGS = 1 << iota
	CONSOLE_SELECTION_NOT_EMPTY
	CONSOLE_MOUSE_SELECTION
	CONSOLE_MOUSE_DOWN
	CONSOLE_NO_SELECTION CONSOLE_SELECTION_FLAGS = 0x0000
)

type HANDLER_ROUTINE func(ctrlType DWORD) BOOL

type COSERVERINFO struct {
	_        DWORD
	Name     WString
	AuthInfo *COAUTHINFO
	_        DWORD
}

type COAUTHIDENTITY struct {
	User           *USHORT
	UserLength     ULONG
	Domain         *USHORT
	DomainLength   ULONG
	Password       *USHORT
	PasswordLength ULONG
	Flags          ULONG
}

type COAUTHINFO struct {
	AuthnSvc           DWORD
	AuthzSvc           DWORD
	ServerPrincName    WString
	AuthnLevel         DWORD
	ImpersonationLevel DWORD
	AuthIdentityData   *COAUTHIDENTITY
	Capabilities       DWORD
}

type MULTI_QI struct {
	IID *IID
	Itf *IUnknown
	Hr  HRESULT
}

type COMSD enum

const (
	SD_LAUNCHPERMISSIONS COMSD = iota
	SD_ACCESSPERMISSIONS
	SD_LAUNCHRESTRICTIONS
	SD_ACCESSRESTRICTIONS
)

type SOLE_AUTHENTICATION_SERVICE struct {
	AuthnSvc      DWORD
	AuthzSvc      DWORD
	PrincipalName *OLECHAR
	//TODO(t): Should probab;y be OLESTR

	Hr HRESULT
}

type UCLSSPEC struct {
	Tyspec DWORD
	//TODO(t):union {
	// CLSID clsid;
	// LPOLESTR pFileExt;
	// LPOLESTR pMimeType;
	// LPOLESTR pProgId;
	// LPOLESTR pFileName;
	// struct             {
	//    LPOLESTR pPackageName;
	//    GUID PolicyId;
	//    } 	ByName;
	// struct             {
	//    GUID ObjectId;
	//    GUID PolicyId;
	//    } 	ByObjectId;
	// } tagged_union
}

type STGOPTIONS struct {
	Version      USHORT
	_            USHORT
	SectorSize   ULONG
	TemplateFile *WChar
	//TODO(t):WString?
}

type QUERYCONTEXT struct {
	Context   DWORD
	Platform  CSPLATFORM
	Locale    LCID
	VersionHi DWORD
	VersionLo DWORD
}

type CSPLATFORM struct {
	PlatformId    DWORD
	VersionHi     DWORD
	VersionLo     DWORD
	ProcessorArch DWORD
}

type BIND_OPTS struct {
	Struct            DWORD
	Flags             DWORD
	Mode              DWORD
	TickCountDeadline DWORD
}

type FORMATETC struct {
	Format CLIPFORMAT
	TD     *DVTARGETDEVICE
	Aspect DWORD
	Index  LONG
	Tymed  DWORD
}

type DVTARGETDEVICE struct {
	Size             DWORD
	DriverNameOffset WORD
	DeviceNameOffset WORD
	PortNameOffset   WORD
	ExtDevmodeOffset WORD
	Data             [1]BYTE
}

type STGMEDIUM struct {
	Tymed DWORD
	//TODO(t):union {

	Bitmap HBITMAP
	// MetaFilePict HMETAFILEPICT
	// EnhMetaFile  HENHMETAFILE
	// Global       HGLOBAL
	// FileName     LPOLESTR
	// Stm          *IStream
	// Stg          *IStorage
	//}
	UnkForRelease *IUnknown
}

type OLEMENUGROUPWIDTHS struct {
	Width [6]LONG
}

type OLEINPLACEFRAMEINFO struct {
	CB           UINT
	MDIApp       BOOL
	Frame        HWND
	Accel        HACCEL
	AccelEntries UINT
}

type SAFEARRAY struct {
	Dims     USHORT
	Features USHORT
	Elements ULONG
	Locks    ULONG
	Data     *VOID
	SABound  [1]SAFEARRAYBOUND
}

type SAFEARRAYBOUND struct {
	Elements ULONG
	Lbound   LONG
}

type SYSKIND enum

const (
	SYS_WIN16 SYSKIND = iota
	SYS_WIN32
	SYS_MAC
	SYS_WIN64
)

type DISPPARAMS struct {
	Varg       *VARIANTARG
	NamedArgs  *DISPID
	CArgs      UINT
	CNamedArgs UINT
}

type VARIANT struct {
	//TODO(t):Variant
	// union
	//     {
	//     struct
	//         {
	//         VARTYPE vt;
	//         WORD wReserved1;
	//         WORD wReserved2;
	//         WORD wReserved3;
	//         union
	//             {
	//             LONGLONG llVal;
	//             LONG lVal;
	//             BYTE bVal;
	//             SHORT iVal;
	//             FLOAT fltVal;
	//             DOUBLE dblVal;
	//             VARIANT_BOOL boolVal;
	//             SCODE scode;
	//             CY cyVal;
	//             DATE date;
	//             BSTR bstrVal;
	//             IUnknown *punkVal;
	//             IDispatch *pdispVal;
	//             SAFEARRAY *parray;
	//             BYTE *pbVal;
	//             SHORT *piVal;
	//             LONG *plVal;
	//             LONGLONG *pllVal;
	//             FLOAT *pfltVal;
	//             DOUBLE *pdblVal;
	//             VARIANT_BOOL *pboolVal;
	//             SCODE *pscode;
	//             CY *pcyVal;
	//             DATE *pdate;
	//             BSTR *pbstrVal;
	//             IUnknown **ppunkVal;
	//             IDispatch **ppdispVal;
	//             SAFEARRAY **pparray;
	//             VARIANT *pvarVal;
	//             PVOID byref;
	//             CHAR cVal;
	//             USHORT uiVal;
	//             ULONG ulVal;
	//             ULONGLONG ullVal;
	//             INT intVal;
	//             UINT uintVal;
	//             DECIMAL *pdecVal;
	//             CHAR *pcVal;
	//             USHORT *puiVal;
	//             ULONG *pulVal;
	//             ULONGLONG *pullVal;
	//             INT *pintVal;
	//             UINT *puintVal;
	//             struct
	//                 {
	//                 PVOID pvRecord;
	//                 IRecordInfo *pRecInfo;
	//                 } 	;
	//             } 	;
	//         } 	;
	//     DECIMAL decVal;
	//     } 	;
}

type CALLCONV enum

const (
	CC_FASTCALL CALLCONV = iota
	CC_CDECL
	CC_MSCPASCAL
	CC_MACPASCAL
	CC_STDCALL
	CC_FPFASTCALL
	CC_SYSCALL
	CC_MPWCDECL
	CC_MPWPASCAL
	CC_MAX
	CC_PASCAL = CC_MSCPASCAL
)

type CY struct {
	//TODO(t):union
	// struct {
	// Lo Unsigned_long
	// Hi LONG
	// }

	Int64 LONGLONG
}

type DECIMAL struct {
	_ USHORT
	//TODO(t):union
	// union {
	//     struct {
	//         BYTE scale;
	//         BYTE sign;
	//     };
	//     USHORT signscale;
	// };
	// ULONG Hi32;
	// union {
	//     struct {
	//         ULONG Lo32;
	//         ULONG Mid32;
	//     };
	//     ULONGLONG Lo64;
	// };
}

type UDATE struct {
	ST        SYSTEMTIME
	DayOfYear USHORT
}

type REGKIND enum

const (
	REGKIND_DEFAULT REGKIND = iota
	REGKIND_REGISTER
	REGKIND_NONE
)

type NUMPARSE struct {
	Dig       INT
	InFlags   ULONG
	OutFlags  ULONG
	Used      INT
	BaseShift INT
	Pwr10     INT
}

type EXCEPINFO struct {
	Code           WORD
	_              WORD
	Source         BSTR
	Description    BSTR
	HelpFile       BSTR
	HelpContext    DWORD
	_              *VOID
	DeferredFillIn func(*EXCEPINFO) HRESULT
	SCode          SCODE
}

type METHODDATA struct {
	Name   *OLECHAR
	Data   *PARAMDATA
	Dispid DISPID
	Meth   UINT
	CC     CALLCONV
	Args   UINT
	Flags  WORD
	Return VARTYPE
}

type INTERFACEDATA struct {
	Methdata *METHODDATA
	Members  UINT
}

type PARAMDATA struct {
	Name *OLECHAR
	VT   VARTYPE
}

type CUSTDATAITEM struct {
	Guid  GUID
	Value VARIANTARG
}
type CUSTDATA struct {
	CustDataSize DWORD
	CustData     *CUSTDATAITEM
}

type TIMEVAL struct {
	Sec  LONG
	USec LONG
}

const WSAPROTOCOL_LEN = 255

type WSAPROTOCOL_INFOA struct {
	ServiceFlags1     DWORD
	ServiceFlags2     DWORD
	ServiceFlags3     DWORD
	ServiceFlags4     DWORD
	ProviderFlags     DWORD
	ProviderId        GUID
	CatalogEntryId    DWORD
	ProtocolChain     WSAPROTOCOLCHAIN
	Version           int
	AddressFamily     int
	MaxSockAddr       int
	MinSockAddr       int
	SocketType        int
	Protocol          int
	ProtocolMaxOffset int
	NetworkByteOrder  int
	SecurityScheme    int
	MessageSize       DWORD
	ProviderReserved  DWORD
	ProtocolString    [WSAPROTOCOL_LEN + 1]AChar
}

type WSAPROTOCOL_INFOW struct {
	ServiceFlags1     DWORD
	ServiceFlags2     DWORD
	ServiceFlags3     DWORD
	ServiceFlags4     DWORD
	ProviderFlags     DWORD
	ProviderId        GUID
	CatalogEntryId    DWORD
	ProtocolChain     WSAPROTOCOLCHAIN
	Version           int
	AddressFamily     int
	MaxSockAddr       int
	MinSockAddr       int
	SocketType        int
	Protocol          int
	ProtocolMaxOffset int
	NetworkByteOrder  int
	SecurityScheme    int
	MessageSize       DWORD
	ProviderReserved  DWORD
	ProtocolString    [WSAPROTOCOL_LEN + 1]WChar
}

const (
	LAYERED_PROTOCOL = iota
	BASE_PROTOCOL
)

const MAX_PROTOCOL_CHAIN = 7

type WSAPROTOCOLCHAIN struct {
	ChainLen     int
	ChainEntries [MAX_PROTOCOL_CHAIN]DWORD
}

type SOCKADDR struct {
	Sa_family U_short
	Sa_data   [14]Char
}

//TODO(t):Check VStrings for i/o
type WSAQUERYSET struct {
	Size                DWORD
	ServiceInstanceName PVString
	ServiceClassId      *GUID
	Version             *WSAVERSION
	Comment             PVString
	NameSpace           DWORD
	NSProviderId        *GUID
	Context             PVString
	NumberOfProtocols   DWORD
	AfpProtocols        *AFPROTOCOLS
	QueryString         PVString
	NumberOfCsAddrs     DWORD
	SaBuffer            *CSADDR_INFO
	OutputFlags         DWORD
	Blob                *BLOB
}

type WSAVERSION struct {
	Version DWORD
	How     WSAECOMPARATOR
}

type WSAECOMPARATOR enum

const (
	COMP_EQUAL WSAECOMPARATOR = iota
	COMP_NOTLESS
)

type AFPROTOCOLS struct {
	AddressFamily INT
	Protocol      INT
}

type CSADDR_INFO struct {
	LocalAddr  SOCKET_ADDRESS
	RemoteAddr SOCKET_ADDRESS
	SocketType INT
	Protocol   INT
}

type BLOB struct {
	Size     ULONG
	BlobData *BYTE
}

type SOCKET_ADDRESS struct {
	Sockaddr       *SOCKADDR
	SockaddrLength INT
}

type WSASERVICECLASSINFO struct {
	ServiceClassId   *GUID
	ServiceClassName POVString
	Count            DWORD
	ClassInfos       *WSANSCLASSINFO
}

type WSANSCLASSINFO struct {
	Name      PVString
	NameSpace DWORD
	ValueType DWORD
	ValueSize DWORD
	Value     *VOID
}

type WSAESETSERVICEOP enum

const (
	RNRSERVICE_REGISTER WSAESETSERVICEOP = iota
	RNRSERVICE_DEREGISTER
	RNRSERVICE_DELETE
)

type HOSTENT struct {
	Name      *Char
	Aliases   **Char
	Addrtype  SHORT
	Length    SHORT
	Addr_list **Char
}

type NETENT struct {
	Name     *Char
	Aliases  **Char
	Addrtype SHORT
	Net      U_long
}

type SERVENT struct {
	Name    *Char
	Aliases **Char
	//TODO(t):Win64
	// Proto *Char
	// Port  SHORT
	//else

	Port  SHORT
	Proto *Char
}

type PROTOENT struct {
	Name    *Char
	Aliases **Char
	Proto   SHORT
}

type IN_ADDR struct {
	//union {
	//struct {
	B1, B2, B3, B4 byte
	//} S_un_b;
	//struct {
	//W1, W2 uint16
	//} S_un_w;
	//Addr uint32
	//} S_un;
}

const FD_SETSIZE = 64

type FD_SET struct {
	Count U_int
	Array [FD_SETSIZE]SOCKET
}

type WSANAMESPACE_INFO struct {
	NSProviderId GUID
	NameSpace    DWORD
	Active       BOOL
	Version      DWORD
	Identifier   PVString
}

type CONDITIONPROC func(
	callerId *WSABUF,
	callerData *WSABUF,
	sQOS *QOS,
	gQOS *QOS,
	calleeId *WSABUF,
	calleeData *WSABUF,
	g *GROUP,
	callbackData DWORD_PTR) int

type WSAOVERLAPPED_COMPLETION_ROUTINE func(
	err DWORD,
	transferred DWORD,
	overlapped *WSAOVERLAPPED,
	flags DWORD)

type WSABUF struct {
	Len U_long
	Buf *Char
}

type QOS struct {
	SendingFlowspec   FLOWSPEC
	ReceivingFlowspec FLOWSPEC
	ProviderSpecific  WSABUF
}

type FLOWSPEC struct {
	TokenRate          ULONG
	TokenBucketSize    ULONG
	PeakBandwidth      ULONG
	Latency            ULONG
	DelayVariation     ULONG
	ServiceType        SERVICETYPE
	MaxSduSize         ULONG
	MinimumPolicedSize ULONG
}

const FD_MAX_EVENTS = 10

type WSANETWORKEVENTS struct {
	NetworkEvents LONG
	ErrorCode     [FD_MAX_EVENTS]int
}

type WSACOMPLETIONTYPE enum

const (
	NSP_NOTIFY_IMMEDIATELY WSACOMPLETIONTYPE = iota
	NSP_NOTIFY_HWND
	NSP_NOTIFY_EVENT
	NSP_NOTIFY_PORT
	NSP_NOTIFY_APC
)

type WSACOMPLETION struct {
	Type WSACOMPLETIONTYPE
	//TODO(t):union {
	//struct {

	Wnd     HWND
	Msg     UINT
	Context WPARAM
	//} WindowMessage;
	//struct {
	//Overlapped WSAOVERLAPPED
	//} Event;
	//struct {
	//Overlapped WSAOVERLAPPED
	//CompletionProc WSAOVERLAPPED_COMPLETION_ROUTINE
	//} Apc;
	//struct {
	//Overlapped WSAOVERLAPPED
	//Port HANDLE
	//Key  ULONG_PTR
	//} Port;
	//} Parameters;
}

const (
	CAL_ICALINTVALUE CALTYPE = iota + 1
	CAL_SCALNAME
	CAL_IYEAROFFSETRANGE
	CAL_SERASTRING
	CAL_SSHORTDATE
	CAL_SLONGDATE
	CAL_SDAYNAME1
	CAL_SDAYNAME2
	CAL_SDAYNAME3
	CAL_SDAYNAME4
	CAL_SDAYNAME5
	CAL_SDAYNAME6
	CAL_SDAYNAME7
	CAL_SABBREVDAYNAME1
	CAL_SABBREVDAYNAME2
	CAL_SABBREVDAYNAME3
	CAL_SABBREVDAYNAME4
	CAL_SABBREVDAYNAME5
	CAL_SABBREVDAYNAME6
	CAL_SABBREVDAYNAME7
	CAL_SMONTHNAME1
	CAL_SMONTHNAME2
	CAL_SMONTHNAME3
	CAL_SMONTHNAME4
	CAL_SMONTHNAME5
	CAL_SMONTHNAME6
	CAL_SMONTHNAME7
	CAL_SMONTHNAME8
	CAL_SMONTHNAME9
	CAL_SMONTHNAME10
	CAL_SMONTHNAME11
	CAL_SMONTHNAME12
	CAL_SMONTHNAME13
	CAL_SABBREVMONTHNAME1
	CAL_SABBREVMONTHNAME2
	CAL_SABBREVMONTHNAME3
	CAL_SABBREVMONTHNAME4
	CAL_SABBREVMONTHNAME5
	CAL_SABBREVMONTHNAME6
	CAL_SABBREVMONTHNAME7
	CAL_SABBREVMONTHNAME8
	CAL_SABBREVMONTHNAME9
	CAL_SABBREVMONTHNAME10
	CAL_SABBREVMONTHNAME11
	CAL_SABBREVMONTHNAME12
	CAL_SABBREVMONTHNAME13
	CAL_SYEARMONTH
	CAL_ITWODIGITYEARMAX
)

func MAKEINTRESOURCE(i WORD) PVString {
	return (PVString)(unsafe.Pointer(uintptr(i)))
}

type WINDOW_STYLE DWORD

const (
	WS_TABSTOP WINDOW_STYLE = 1 << (16 + iota)
	WS_GROUP
	WS_THICKFRAME
	WS_SYSMENU
	WS_HSCROLL
	WS_VSCROLL
	WS_DLGFRAME
	WS_BORDER
	WS_MAXIMIZE
	WS_CLIPCHILDREN
	WS_CLIPSIBLINGS
	WS_DISABLED
	WS_VISIBLE
	WS_MINIMIZE
	WS_CHILD
	WS_POPUP
)
const (
	WS_OVERLAPPED WINDOW_STYLE = 0

	WS_MAXIMIZEBOX      = WS_TABSTOP
	WS_MINIMIZEBOX      = WS_GROUP
	WS_CAPTION          = WS_BORDER | WS_DLGFRAME
	WS_TILED            = WS_OVERLAPPED
	WS_ICONIC           = WS_MINIMIZE
	WS_SIZEBOX          = WS_THICKFRAME
	WS_TILEDWINDOW      = WS_OVERLAPPEDWINDOW
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION |
		WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX |
		WS_MAXIMIZEBOX
	WS_POPUPWINDOW = WS_POPUP | WS_BORDER | WS_SYSMENU
	WS_CHILDWINDOW = WS_CHILD
)

type WINDOW_STYLE_EX DWORD

const (
	WS_EX_DLGMODALFRAME WINDOW_STYLE_EX = 1 << iota
	_
	WS_EX_NOPARENTNOTIFY
	WS_EX_TOPMOST
	WS_EX_ACCEPTFILES
	WS_EX_TRANSPARENT
	WS_EX_MDICHILD
	WS_EX_TOOLWINDOW
	WS_EX_WINDOWEDGE
	WS_EX_CLIENTEDGE
	WS_EX_CONTEXTHELP
	_
	WS_EX_RIGHT
	WS_EX_RTLREADING
	WS_EX_LEFTSCROLLBAR
	_
	WS_EX_CONTROLPARENT
	WS_EX_STATICEDGE
	WS_EX_APPWINDOW
	WS_EX_LAYERED
	WS_EX_NOINHERITLAYOUT
	_
	WS_EX_LAYOUTRTL
	_
	_
	WS_EX_COMPOSITED
	_
	WS_EX_NOACTIVATE
)
const (
	WS_EX_LEFT WINDOW_STYLE_EX = 0
	WS_EX_LTRREADING
	WS_EX_RIGHTSCROLLBAR
	WS_EX_OVERLAPPEDWINDOW WINDOW_STYLE_EX = WS_EX_WINDOWEDGE | WS_EX_CLIENTEDGE
	WS_EX_PALETTEWINDOW    WINDOW_STYLE_EX = WS_EX_WINDOWEDGE | WS_EX_TOOLWINDOW | WS_EX_TOPMOST
)

type CLASS_STYLE DWORD

const (
	CS_VREDRAW CLASS_STYLE = 1 << iota
	CS_HREDRAW
	_
	CS_DBLCLKS
	_
	CS_OWNDC
	CS_CLASSDC
	CS_PARENTDC
	_
	CS_NOCLOSE
	_
	CS_SAVEBITS
	CS_BYTEALIGNCLIENT
	CS_BYTEALIGNWINDOW
	CS_GLOBALCLASS
	_
	CS_IME
	CS_DROPSHADOW
)

const (
	BS_PUSHBUTTON WINDOW_STYLE = iota
	BS_DEFPUSHBUTTON
	BS_CHECKBOX
	BS_AUTOCHECKBOX
	BS_RADIOBUTTON
	BS_3STATE
	BS_AUTO3STATE
	BS_GROUPBOX
	BS_USERBUTTON
	BS_AUTORADIOBUTTON
	BS_PUSHBOX
	BS_OWNERDRAW
	BS_TYPEMASK = 0x0000000F
)

const (
	BS_TEXT     WINDOW_STYLE = 0
	BS_LEFTTEXT              = 0x10 << iota
	BS_ICON
	BS_BITMAP
	BS_LEFT
	BS_RIGHT
	BS_TOP
	BS_BOTTOM
	BS_PUSHLIKE
	BS_MULTILINE
	BS_NOTIFY
	BS_FLAT
	BS_CENTER      = BS_LEFT | BS_RIGHT
	BS_VCENTER     = BS_TOP | BS_BOTTOM
	BS_RIGHTBUTTON = BS_LEFTTEXT
)

const (
	CBS_SIMPLE WINDOW_STYLE = 1 << iota
	CBS_DROPDOWN
	_
	_
	CBS_OWNERDRAWFIXED
	CBS_OWNERDRAWVARIABLE
	CBS_AUTOHSCROLL
	CBS_OEMCONVERT
	CBS_SORT
	CBS_HASSTRINGS
	CBS_NOINTEGRALHEIGHT
	CBS_DISABLENOSCROLL
	_
	CBS_UPPERCASE
	CBS_LOWERCASE
	CBS_DROPDOWNLIST = CBS_SIMPLE | CBS_DROPDOWN
)
const (
	ES_CENTER WINDOW_STYLE = 1 << iota
	ES_RIGHT
	ES_MULTILINE
	ES_UPPERCASE
	ES_LOWERCASE
	ES_PASSWORD
	ES_AUTOVSCROLL
	ES_AUTOHSCROLL
	ES_NOHIDESEL
	_
	ES_OEMCONVERT
	ES_READONLY
	ES_WANTRETURN
	ES_NUMBER
	ES_LEFT = 0
)
const (
	LBS_NOTIFY WINDOW_STYLE = 1 << iota
	LBS_SORT
	LBS_NOREDRAW
	LBS_MULTIPLESEL
	LBS_OWNERDRAWFIXED
	LBS_OWNERDRAWVARIABLE
	LBS_HASSTRINGS
	LBS_USETABSTOPS
	LBS_NOINTEGRALHEIGHT
	LBS_MULTICOLUMN
	LBS_WANTKEYBOARDINPUT
	LBS_EXTENDEDSEL
	LBS_DISABLENOSCROLL
	LBS_NODATA
	LBS_NOSEL
	LBS_COMBOBOX
	LBS_STANDARD = LBS_NOTIFY | LBS_SORT | WS_VSCROLL | WS_BORDER
)

//TODO(t):Rich Edit Styles

const (
	SBS_VERT WINDOW_STYLE = 1 << iota
	SBS_TOPALIGN
	SBS_BOTTOMALIGN
	SBS_SIZEBOX
	SBS_SIZEGRIP
	SBS_HORZ                    = SBS_VERT
	SBS_LEFTALIGN               = SBS_TOPALIGN
	SBS_RIGHTALIGN              = SBS_BOTTOMALIGN
	SBS_SIZEBOXTOPLEFTALIGN     = SBS_TOPALIGN
	SBS_SIZEBOXBOTTOMRIGHTALIGN = SBS_BOTTOMALIGN
)
const (
	SS_LEFT WINDOW_STYLE = iota
	SS_CENTER
	SS_RIGHT
	SS_ICON
	SS_BLACKRECT
	SS_GRAYRECT
	SS_WHITERECT
	SS_BLACKFRAME
	SS_GRAYFRAME
	SS_WHITEFRAME
	SS_USERITEM
	SS_SIMPLE
	SS_LEFTNOWORDWRAP
	SS_OWNERDRAW
	SS_BITMAP
	SS_ENHMETAFILE
	SS_ETCHEDHORZ
	SS_ETCHEDVERT
	SS_ETCHEDFRAME
	SS_TYPEMASK = 0x1F
)
const (
	SS_REALSIZECONTROL WINDOW_STYLE = 0x40 << iota
	SS_NOPREFIX
	SS_NOTIFY
	SS_CENTERIMAGE
	SS_RIGHTJUST
	SS_REALSIZEIMAGE
	SS_SUNKEN
	SS_EDITCONTROL
	SS_ENDELLIPSIS
	SS_PATHELLIPSIS
	SS_WORDELLIPSIS = SS_ENDELLIPSIS | SS_PATHELLIPSIS
	SS_ELLIPSISMASK = SS_ENDELLIPSIS | SS_PATHELLIPSIS
)

const CW_USEDEFAULT int = -0x80000000

type SHELLHOOKINFO struct {
	Hwnd HWND
	Rc   RECT
}

type EVENTMSG struct {
	Message UINT
	ParamL  UINT
	ParamH  UINT
	Time    DWORD
	Hwnd    HWND
}

type CWPSTRUCT struct {
	LParam  LPARAM
	WParam  WPARAM
	Message UINT
	Hwnd    HWND
}

type CWPRETSTRUCT struct {
	LResult LRESULT
	LParam  LPARAM
	WParam  WPARAM
	Message UINT
	Hwnd    HWND
}

type KBDLLHOOKSTRUCT struct {
	VKCode    DWORD
	ScanCode  DWORD
	Flags     DWORD
	Time      DWORD
	ExtraInfo ULONG_PTR
}

type MSLLHOOKSTRUCT struct {
	Pt        POINT
	MouseData DWORD
	Flags     DWORD
	Time      DWORD
	ExtraInfo ULONG_PTR
}

type DEBUGHOOKINFO struct {
	Thread          DWORD
	ThreadInstaller DWORD
	LParam          LPARAM
	WParam          WPARAM
	Code            int
}

type MOUSEHOOKSTRUCT struct {
	Pt          POINT
	Wnd         HWND
	HitTestCode UINT
	ExtraInfo   ULONG_PTR
}

type MOUSEHOOKSTRUCTEX struct {
	MHS       MOUSEHOOKSTRUCT
	MouseData DWORD
}

type HARDWAREHOOKSTRUCT struct {
	Wnd     HWND
	Message UINT
	WParam  WPARAM
	LParam  LPARAM
}

type USEROBJECTFLAGS struct {
	Inherit  BOOL
	Reserved BOOL
	Flags    DWORD
}

type MINMAXINFO struct {
	Reserved     POINT
	MaxSize      POINT
	MaxPosition  POINT
	MinTrackSize POINT
	MaxTrackSize POINT
}

type COPYDATASTRUCT struct {
	Data1     ULONG_PTR
	Data2Size DWORD
	Data2     *VOID
}

type MDINEXTMENU struct {
	MenuIn   HMENU
	MenuNext HMENU
	WndNext  HWND
}

type WINDOWPOS struct {
	Wnd            HWND
	WndInsertAfter HWND
	X              int
	Y              int
	CX             int
	CY             int
	Flags          UINT
}

type NCCALCSIZE_PARAMS struct {
	Rgrc [3]RECT
	Pos  *WINDOWPOS
}

type CREATESTRUCT struct {
	CreateParams *VOID
	Instance     HINSTANCE
	Menu         HMENU
	WndParent    HWND
	CY           int
	CX           int
	Y            int
	X            int
	Style        LONG
	Name         VString
	Class        VString
	ExStyle      DWORD
}

type NMHDR struct {
	WndFrom HWND
	IdFrom  UINT_PTR
	Code    UINT
}

type STYLESTRUCT struct {
	StyleOld DWORD
	StyleNew DWORD
}

type MEASUREITEMSTRUCT struct {
	CtlType UINT
	CtlID   UINT
	ID      UINT
	Width   UINT
	Height  UINT
	Data    ULONG_PTR
}

type DRAWITEMSTRUCT struct {
	CtlType UINT
	CtlID   UINT
	ID      UINT
	Action  UINT
	State   UINT
	Item    HWND
	DC      HDC
	Rect    RECT
	Data    ULONG_PTR
}

type DELETEITEMSTRUCT struct {
	CtlType UINT
	CtlID   UINT
	ID      UINT
	Item    HWND
	Data    ULONG_PTR
}

type COMPAREITEMSTRUCT struct {
	CtlType  UINT
	CtlID    UINT
	Item     HWND
	ID1      UINT
	Data1    ULONG_PTR
	ID2      UINT
	Data2    ULONG_PTR
	LocaleId DWORD
}

type MENUGETOBJECTINFO struct {
	Flags DWORD
	Pos   UINT
	Menu  HMENU
	Riid  *VOID
	Obj   *VOID
}

//TODO(t):Obsolete
type DROPSTRUCT struct {
	Source      HWND
	Sink        HWND
	Fmt         DWORD
	Data        ULONG_PTR
	Drop        POINT
	ControlData DWORD
}

type MENUITEMTEMPLATEHEADER struct {
	VersionNumber WORD
	Offset        WORD
}

type MENUITEMTEMPLATE struct {
	Option WORD
	ID     WORD
	String [1]WChar
}

type CURSORSHAPE struct {
	XHotSpot  int
	UHotSpot  int
	CX        int
	CY        int
	Width     int
	Planes    BYTE
	BitsPixel BYTE
}

type MDICREATESTRUCT struct {
	Class  VString
	Title  VString
	Owner  HANDLE
	X      int
	Y      int
	CX     int
	CY     int
	Style  DWORD
	LParam LPARAM
}

type CLIENTCREATESTRUCT struct {
	WindowMenu HANDLE
	FirstChild UINT
}

type NONCLIENTMETRICSA struct {
	Size            UINT
	BorderWidth     int
	ScrollWidth     int
	ScrollHeight    int
	CaptionWidth    int
	CaptionHeight   int
	CaptionFont     LOGFONTA
	SmCaptionWidth  int
	SmCaptionHeight int
	SmCaptionFont   LOGFONTA
	MenuWidth       int
	MenuHeight      int
	MenuFont        LOGFONTA
	StatusFont      LOGFONTA
	MessageFont     LOGFONTA
}
type NONCLIENTMETRICSW struct {
	Size            UINT
	BorderWidth     int
	ScrollWidth     int
	ScrollHeight    int
	CaptionWidth    int
	CaptionHeight   int
	CaptionFont     LOGFONTW
	SmCaptionWidth  int
	SmCaptionHeight int
	SmCaptionFont   LOGFONTW
	MenuWidth       int
	MenuHeight      int
	MenuFont        LOGFONTW
	StatusFont      LOGFONTW
	MessageFont     LOGFONTW
}

type MINIMIZEDMETRICS struct {
	Size    UINT
	Width   int
	HorzGap int
	VertGap int
	Arrange int
}

type ICONMETRICSA struct {
	Size        UINT
	HorzSpacing int
	VertSpacing int
	TitleWrap   int
	fFont       LOGFONTA
}
type ICONMETRICSW struct {
	Size        UINT
	HorzSpacing int
	VertSpacing int
	TitleWrap   int
	fFont       LOGFONTW
}

type ANIMATIONINFO struct {
	Size       UINT
	MinAnimate int
}

type SERIALKEYS struct {
	Size       UINT
	Flags      DWORD
	ActivePort VString
	Port       VString
	BaudRate   UINT
	PortState  UINT
	Active     UINT
}

type HIGHCONTRASTA struct {
	Size          UINT
	Flags         DWORD
	DefaultScheme VString
}

type FILTERKEYS struct {
	Size       UINT
	Flags      DWORD
	WaitMSec   DWORD
	DelayMSec  DWORD
	RepeatMSec DWORD
	BounceMSec DWORD
}

type STICKYKEYS struct {
	Size  UINT
	Flags DWORD
}

type MOUSEKEYS struct {
	Size           UINT
	Flags          DWORD
	MaxSpeed       DWORD
	TimeToMaxSpeed DWORD
	CtrlSpeed      DWORD
	_              DWORD
	_              DWORD
}

type ACCESSTIMEOUT struct {
	Size        UINT
	Flags       DWORD
	TimeOutMSec DWORD
}

type SOUNDSENTRY struct {
	Size                  UINT
	Flags                 DWORD
	FSTextEffect          DWORD
	FSTextEffectMSec      DWORD
	FSTextEffectColorBits DWORD
	FSGrafEffect          DWORD
	FSGrafEffectMSec      DWORD
	FSGrafEffectColor     DWORD
	WindowsEffect         DWORD
	WindowsEffectMSec     DWORD
	WindowsEffectDLL      VString
	WindowsEffectOrdinal  DWORD
}

type TOGGLEKEYS struct {
	Size  UINT
	Flags DWORD
}

type MONITORINFOEXA struct {
	MI     MONITORINFO
	Device [CCHDEVICENAME]AChar
}

type MONITORINFOEXW struct {
	MI     MONITORINFO
	Device [CCHDEVICENAME]WChar
}

type RID_DEVICE_INFO_MOUSE struct {
	Id              DWORD
	NumberOfButtons DWORD
	SampleRate      DWORD
}

type RID_DEVICE_INFO_KEYBOARD struct {
	Type                 DWORD
	SubType              DWORD
	KeyboardMode         DWORD
	NumberOfFunctionKeys DWORD
	NumberOfIndicators   DWORD
	NumberOfKeysTotal    DWORD
}

type RID_DEVICE_INFO_HID struct {
	VendorId      DWORD
	ProductId     DWORD
	VersionNumber DWORD
	UsagePage     USHORT
	Usage         USHORT
}

type RID_DEVICE_INFO struct {
	Size DWORD
	Type DWORD
	//TODO(t):union
	// Mouse    RID_DEVICE_INFO_MOUSE
	Keyboard RID_DEVICE_INFO_KEYBOARD
	// HID      RID_DEVICE_INFO_HID
}

type MSGBOX_TYPE DWORD

//#if(WINVER >= 0x0400)
// * MessageBox() Flags
const (
	MB_OK MSGBOX_TYPE = iota
	MB_OKCANCEL
	MB_ABORTRETRYIGNORE
	MB_YESNOCANCEL
	MB_YESNO
	MB_RETRYCANCEL
	MB_CANCELTRYCONTINUE
)
const (
	MB_ICONHAND MSGBOX_TYPE = 0x10 << iota
	MB_ICONQUESTION
	MB_ICONEXCLAMATION
	MB_ICONASTERISK
	MB_USERICON
	MB_DEFBUTTON2
	MB_DEFBUTTON3
	_
	_
	MB_SYSTEMMODAL
	MB_TASKMODAL
	MB_HELP
	MB_NOFOCUS
	MB_SETFOREGROUND
	MB_DEFAULT_DESKTOP_ONLY
	MB_TOPMOST
	MB_RIGHT
	MB_RTLREADING
	MB_SERVICE_NOTIFICATION
)
const (
	MB_DEFBUTTON1 MSGBOX_TYPE = 0
	MB_DEFBUTTON4 MSGBOX_TYPE = 0x300
	MB_APPLMODAL  MSGBOX_TYPE = 0

	MB_TYPEMASK MSGBOX_TYPE = 0xF
	MB_ICONMASK MSGBOX_TYPE = 0xF0
	MB_DEFMASK  MSGBOX_TYPE = 0xF00
	MB_MODEMASK MSGBOX_TYPE = 0x3000
	MB_MISCMASK MSGBOX_TYPE = 0xC000

	MB_ICONWARNING     = MB_ICONEXCLAMATION
	MB_ICONERROR       = MB_ICONHAND
	MB_ICONINFORMATION = MB_ICONASTERISK
	MB_ICONSTOP        = MB_ICONHAND
)

var ( // Cannot be const as they are pseudo-*string
	IDI_APPLICATION = (PVString)(unsafe.Pointer(uintptr(32512)))
	IDI_HAND        = (PVString)(unsafe.Pointer(uintptr(32513)))
	IDI_QUESTION    = (PVString)(unsafe.Pointer(uintptr(32514)))
	IDI_EXCLAMATION = (PVString)(unsafe.Pointer(uintptr(32515)))
	IDI_ASTERISK    = (PVString)(unsafe.Pointer(uintptr(32516)))
	IDI_WINLOGO     = (PVString)(unsafe.Pointer(uintptr(32517)))
	IDI_WARNING     = IDI_EXCLAMATION
	IDI_ERROR       = IDI_HAND
	IDI_INFORMATION = IDI_ASTERISK
)

const (
	GWL_WNDPROC    = -4
	GWL_HINSTANCE  = -6
	GWL_HWNDPARENT = -8
	GWL_STYLE      = -16
	GWL_EXSTYLE    = -20
	GWL_USERDATA   = -21
	GWL_ID         = -12
)

//WIN64
// GWL_WNDPROC
// GWL_HINSTANCE
// GWL_HWNDPARENT
// GWL_USERDATA

///*
// * Window Messages
// */

type WINDOW_MESSAGE UINT

const (
	WM_NULL WINDOW_MESSAGE = iota
	WM_CREATE
	WM_DESTROY
	WM_MOVE
	_
	WM_SIZE
	WM_ACTIVATE
	///*
	// * WM_ACTIVATE state values
	// */
	//#define     WA_INACTIVE     0
	//#define     WA_ACTIVE       1
	//#define     WA_CLICKACTIVE  2

	WM_SETFOCUS
	WM_KILLFOCUS
	_
	WM_ENABLE
	WM_SETREDRAW
	WM_SETTEXT
	WM_GETTEXT
	WM_GETTEXTLENGTH
	WM_PAINT
	WM_CLOSE // 0x0010
	WM_QUERYENDSESSION
	WM_QUIT
	WM_QUERYOPEN
	WM_ERASEBKGND
	WM_SYSCOLORCHANGE
	WM_ENDSESSION
	_
	WM_SHOWWINDOW
	_
	WM_WININICHANGE

	//#if(WINVER >= 0x0400)
	//#define WM_SETTINGCHANGE                WM_WININICHANGE
	//#endif /* WINVER >= 0x0400 */

	WM_DEVMODECHANGE
	WM_ACTIVATEAPP
	WM_FONTCHANGE
	WM_TIMECHANGE
	WM_CANCELMODE
	WM_SETCURSOR // 0x0020
	WM_MOUSEACTIVATE
	WM_CHILDACTIVATE
	WM_QUEUESYNC
	WM_GETMINMAXINFO
	_
	WM_PAINTICON
	WM_ICONERASEBKGND
	WM_NEXTDLGCTL
	_
	WM_SPOOLERSTATUS
	WM_DRAWITEM
	WM_MEASUREITEM
	WM_DELETEITEM
	WM_VKEYTOITEM
	WM_CHARTOITEM
	WM_SETFONT // 0x0030
	WM_GETFONT
	WM_SETHOTKEY
	WM_GETHOTKEY
	_
	_
	_
	WM_QUERYDRAGICON
	_
	WM_COMPAREITEM
	_
	_
	_
	WM_GETOBJECT
	_
	_
	_ // 0x0040
	WM_COMPACTING
	_
	_
	WM_COMMNOTIFY // no longer suported
	_
	WM_WINDOWPOSCHANGING
	WM_WINDOWPOSCHANGED
	WM_POWER
	// wParam for WM_POWER window message and DRV_POWER driver notification
	//#define PWR_OK              1
	//#define PWR_FAIL            (-1)
	//#define PWR_SUSPENDREQUEST  1
	//#define PWR_SUSPENDRESUME   2
	//#define PWR_CRITICALRESUME  3
	_
	WM_COPYDATA
	WM_CANCELJOURNAL
	_
	_
	WM_NOTIFY
	_
	WM_INPUTLANGCHANGEREQUEST // 0x0050
	WM_INPUTLANGCHANGE
	WM_TCARD
	WM_HELP
	WM_USERCHANGED
	WM_NOTIFYFORMAT

//#define NFR_ANSI                             1
//#define NFR_UNICODE                          2
//#define NF_QUERY                             3
//#define NF_REQUERY                           4
)
const (
	WM_CONTEXTMENU WINDOW_MESSAGE = iota + 0x007B
	WM_STYLECHANGING
	WM_STYLECHANGED
	WM_DISPLAYCHANGE
	WM_GETICON
	WM_SETICON // 0x80
	WM_NCCREATE
	WM_NCDESTROY
	WM_NCCALCSIZE
	WM_NCHITTEST
	WM_NCPAINT
	WM_NCACTIVATE
	WM_GETDLGCODE
	WM_SYNCPAINT
)
const (
	WM_NCMOUSEMOVE WINDOW_MESSAGE = iota + 0x00A0
	WM_NCLBUTTONDOWN
	WM_NCLBUTTONUP
	WM_NCLBUTTONDBLCLK
	WM_NCRBUTTONDOWN
	WM_NCRBUTTONUP
	WM_NCRBUTTONDBLCLK
	WM_NCMBUTTONDOWN
	WM_NCMBUTTONUP
	WM_NCMBUTTONDBLCLK
	_
	WM_NCXBUTTONDOWN
	WM_NCXBUTTONUP
	WM_NCXBUTTONDBLCLK
)
const (
	WM_INPUT   WINDOW_MESSAGE = iota + 0x00FF
	WM_KEYDOWN                //0x100
	WM_KEYUP
	WM_CHAR
	WM_DEADCHAR
	WM_SYSKEYDOWN
	WM_SYSKEYUP
	WM_SYSCHAR
	WM_SYSDEADCHAR
	_
	WM_UNICHAR
	_
	_
	_
	WM_IME_STARTCOMPOSITION
	WM_IME_ENDCOMPOSITION
	WM_IME_COMPOSITION
	WM_INITDIALOG //0x110
	WM_COMMAND
	WM_SYSCOMMAND
	WM_TIMER
	WM_HSCROLL
	WM_VSCROLL
	WM_INITMENU
	WM_INITMENUPOPUP
	_
	_
	_
	_
	_
	_
	_
	WM_MENUSELECT
	WM_MENUCHAR // 0x120
	WM_ENTERIDLE
	WM_MENURBUTTONUP
	WM_MENUDRAG
	WM_MENUGETOBJECT
	WM_UNINITMENUPOPUP
	WM_MENUCOMMAND
	WM_CHANGEUISTATE
	WM_UPDATEUISTATE
	WM_QUERYUISTATE
	// LOWORD(wParam) values in WM_*UISTATE
	// UIS_SET                         1
	// UIS_CLEAR                       2
	// UIS_INITIALIZE                  3
	// HIWORD(wParam) values in WM_*UISTATE
	// UISF_HIDEFOCUS                  0x1
	// UISF_HIDEACCEL                  0x2
	// UISF_ACTIVE                     0x4
	_
	_
	_
	_
	_
	_
	_ //0x130
	_
	WM_CTLCOLORMSGBOX
	WM_CTLCOLOREDIT
	WM_CTLCOLORLISTBOX
	WM_CTLCOLORBTN
	WM_CTLCOLORDLG
	WM_CTLCOLORSCROLLBAR
	WM_CTLCOLORSTATIC // 0x0138

	WM_KEYFIRST = WM_KEYDOWN
	WM_KEYLAST  = WM_UNICHAR
	// UNICODE_NOCHAR                  0xFFFF
	WM_IME_KEYLAST = WM_IME_COMPOSITION

	// MN_GETHMENU                     0x01E1
)
const (
	WM_MOUSEFIRST WINDOW_MESSAGE = iota + 0x200
	WM_MOUSEMOVE
	WM_LBUTTONDOWN
	WM_LBUTTONUP
	WM_LBUTTONDBLCLK
	WM_RBUTTONDOWN
	WM_RBUTTONUP
	WM_RBUTTONDBLCLK
	WM_MBUTTONDOWN
	WM_MBUTTONUP
	WM_MBUTTONDBLCLK
	WM_MOUSEWHEEL
	WM_XBUTTONDOWN
	WM_XBUTTONUP
	WM_XBUTTONDBLCLK
	_
	// Value for rolling one detent
	// WHEEL_DELTA                     120
	// GET_WHEEL_DELTA_WPARAM(wParam)  ((short)HIWORD(wParam))
	// Setting to scroll one page for SPI_GET/SETWHEELSCROLLLINES
	// WHEEL_PAGESCROLL                (UINT_MAX)

	// GET_KEYSTATE_WPARAM(wParam)     (LOWORD(wParam))
	// GET_NCHITTEST_WPARAM(wParam)    ((short)LOWORD(wParam))
	// GET_XBUTTON_WPARAM(wParam)      (HIWORD(wParam))

	// XButton values are WORD flags
	// XBUTTON1      0x0001
	// XBUTTON2      0x0002
	// Were there to be an XBUTTON3, its value would be 0x0004

	WM_PARENTNOTIFY // 0x210
	WM_ENTERMENULOOP
	WM_EXITMENULOOP
	WM_NEXTMENU
	WM_SIZING
	WM_CAPTURECHANGED
	WM_MOVING
	_

	WM_POWERBROADCAST
	//PBT_APMQUERYSUSPEND             0x0000
	//PBT_APMQUERYSTANDBY             0x0001
	//PBT_APMQUERYSUSPENDFAILED       0x0002
	//PBT_APMQUERYSTANDBYFAILED       0x0003
	//PBT_APMSUSPEND                  0x0004
	//PBT_APMSTANDBY                  0x0005
	//PBT_APMRESUMECRITICAL           0x0006
	//PBT_APMRESUMESUSPEND            0x0007
	//PBT_APMRESUMESTANDBY            0x0008
	//PBTF_APMRESUMEFROMFAILURE       0x00000001
	//PBT_APMBATTERYLOW               0x0009
	//PBT_APMPOWERSTATUSCHANGE        0x000A
	//PBT_APMOEMEVENT                 0x000B
	//PBT_APMRESUMEAUTOMATIC          0x0012
	WM_DEVICECHANGE // 0x219
	_
	_
	_
	_
	_
	_
	WM_MDICREATE // 0x220
	WM_MDIDESTROY
	WM_MDIACTIVATE
	WM_MDIRESTORE
	WM_MDINEXT
	WM_MDIMAXIMIZE
	WM_MDITILE
	WM_MDICASCADE
	WM_MDIICONARRANGE
	WM_MDIGETACTIVE
	_
	_
	_
	_
	_
	_
	WM_MDISETMENU // 0x230
	WM_ENTERSIZEMOVE
	WM_EXITSIZEMOVE
	WM_DROPFILES
	WM_MDIREFRESHMENU //0x234
)
const (
	WM_IME_SETCONTEXT WINDOW_MESSAGE = 0x281 + iota
	WM_IME_NOTIFY
	WM_IME_CONTROL
	WM_IME_COMPOSITIONFULL
	WM_IME_SELECT
	WM_IME_CHAR
	_
	WM_IME_REQUEST
)
const (
	WM_IME_KEYDOWN WINDOW_MESSAGE = 0x290 + iota
	WM_IME_KEYUP
)
const (
	WM_NCMOUSEHOVER WINDOW_MESSAGE = 0x2A0 + iota
	WM_MOUSEHOVER
	WM_NCMOUSELEAVE
	WM_MOUSELEAVE
)
const (
	WM_WTSSESSION_CHANGE WINDOW_MESSAGE = 0x2B1
	WM_TABLET_FIRST      WINDOW_MESSAGE = 0x2c0
	WM_TABLET_LAST       WINDOW_MESSAGE = 0x2df
)
const (
	WM_CUT WINDOW_MESSAGE = 0x300 + iota
	WM_COPY
	WM_PASTE
	WM_CLEAR
	WM_UNDO
	WM_RENDERFORMAT
	WM_RENDERALLFORMATS
	WM_DESTROYCLIPBOARD
	WM_DRAWCLIPBOARD
	WM_PAINTCLIPBOARD
	WM_VSCROLLCLIPBOARD
	WM_SIZECLIPBOARD
	WM_ASKCBFORMATNAME
	WM_CHANGECBCHAIN
	WM_HSCROLLCLIPBOARD
	WM_QUERYNEWPALETTE
	WM_PALETTEISCHANGING // 0x310
	WM_PALETTECHANGED
	WM_HOTKEY
	_
	_
	_
	_
	WM_PRINT
	WM_PRINTCLIENT
	WM_APPCOMMAND
	WM_THEMECHANGED // 0x31A
)
const (
	WM_HANDHELDFIRST WINDOW_MESSAGE = 0x0358
	WM_HANDHELDLAST  WINDOW_MESSAGE = 0x035F
	WM_AFXFIRST      WINDOW_MESSAGE = 0x0360
	WM_AFXLAST       WINDOW_MESSAGE = 0x037F
	WM_PENWINFIRST   WINDOW_MESSAGE = 0x0380
	WM_PENWINLAST    WINDOW_MESSAGE = 0x038F
	WM_APP           WINDOW_MESSAGE = 0x8000
	WM_USER          WINDOW_MESSAGE = 0x0400
)

var (
	OCR_NORMAL      = (PVString)(unsafe.Pointer(uintptr(32512)))
	OCR_IBEAM       = (PVString)(unsafe.Pointer(uintptr(32513)))
	OCR_WAIT        = (PVString)(unsafe.Pointer(uintptr(32514)))
	OCR_CROSS       = (PVString)(unsafe.Pointer(uintptr(32515)))
	OCR_UP          = (PVString)(unsafe.Pointer(uintptr(32516)))
	OCR_SIZENWSE    = (PVString)(unsafe.Pointer(uintptr(32642)))
	OCR_SIZENESW    = (PVString)(unsafe.Pointer(uintptr(32643)))
	OCR_SIZEWE      = (PVString)(unsafe.Pointer(uintptr(32644)))
	OCR_SIZENS      = (PVString)(unsafe.Pointer(uintptr(32645)))
	OCR_SIZEALL     = (PVString)(unsafe.Pointer(uintptr(32646)))
	OCR_NO          = (PVString)(unsafe.Pointer(uintptr(32648)))
	OCR_HAND        = (PVString)(unsafe.Pointer(uintptr(32649)))
	OCR_APPSTARTING = (PVString)(unsafe.Pointer(uintptr(32650)))
	// OCR_SIZE   32640   OBSOLETE: use OCR_SIZEALL
	// OCR_ICON   32641   OBSOLETE: use OCR_NORMAL
	// OCR_ICOCUR 32647   OBSOLETE: use OIC_WINLOGO
)

var (
	IDC_ARROW       = (PVString)(unsafe.Pointer(uintptr(32512))) // MAKEINTRESOURCE(32512)
	IDC_IBEAM       = (PVString)(unsafe.Pointer(uintptr(32513)))
	IDC_WAIT        = (PVString)(unsafe.Pointer(uintptr(32514)))
	IDC_CROSS       = (PVString)(unsafe.Pointer(uintptr(32515)))
	IDC_UPARROW     = (PVString)(unsafe.Pointer(uintptr(32516)))
	IDC_SIZENWSE    = (PVString)(unsafe.Pointer(uintptr(32642)))
	IDC_SIZENESW    = (PVString)(unsafe.Pointer(uintptr(32643)))
	IDC_SIZEWE      = (PVString)(unsafe.Pointer(uintptr(32644)))
	IDC_SIZENS      = (PVString)(unsafe.Pointer(uintptr(32645)))
	IDC_SIZEALL     = (PVString)(unsafe.Pointer(uintptr(32646)))
	IDC_NO          = (PVString)(unsafe.Pointer(uintptr(32648)))
	IDC_HAND        = (PVString)(unsafe.Pointer(uintptr(32649)))
	IDC_APPSTARTING = (PVString)(unsafe.Pointer(uintptr(32650)))
	IDC_HELP        = (PVString)(unsafe.Pointer(uintptr(32651)))
	// IDC_SIZE 32640  OBSOLETE: use IDC_SIZEALL
	// IDC_ICON 32641  OBSOLETE: use IDC_ARROW
)

const (
	COLOR_SCROLLBAR = iota
	COLOR_BACKGROUND
	COLOR_ACTIVECAPTION
	COLOR_INACTIVECAPTION
	COLOR_MENU
	COLOR_WINDOW
	COLOR_WINDOWFRAME
	COLOR_MENUTEXT
	COLOR_WINDOWTEXT
	COLOR_CAPTIONTEXT
	COLOR_ACTIVEBORDER
	COLOR_INACTIVEBORDER
	COLOR_APPWORKSPACE
	COLOR_HIGHLIGHT
	COLOR_HIGHLIGHTTEXT
	COLOR_BTNFACE
	COLOR_BTNSHADOW
	COLOR_GRAYTEXT
	COLOR_BTNTEXT
	COLOR_INACTIVECAPTIONTEXT
	COLOR_BTNHIGHLIGHT
	COLOR_3DDKSHADOW
	COLOR_3DLIGHT
	COLOR_INFOTEXT
	COLOR_INFOBK
	_
	COLOR_HOTLIGHT
	COLOR_GRADIENTACTIVECAPTION
	COLOR_GRADIENTINACTIVECAPTION
	COLOR_MENUHILIGHT
	COLOR_MENUBAR
	COLOR_DESKTOP     = COLOR_BACKGROUND
	COLOR_3DFACE      = COLOR_BTNFACE
	COLOR_3DSHADOW    = COLOR_BTNSHADOW
	COLOR_3DHIGHLIGHT = COLOR_BTNHIGHLIGHT
	COLOR_3DHILIGHT   = COLOR_BTNHIGHLIGHT
	COLOR_BTNHILIGHT  = COLOR_BTNHIGHLIGHT
)

const (
	SW_HIDE = iota
	SW_SHOWNORMAL
	SW_SHOWMINIMIZED
	SW_SHOWMAXIMIZED
	SW_SHOWNOACTIVATE
	SW_SHOW
	SW_MINIMIZE
	SW_SHOWMINNOACTIVE
	SW_SHOWNA
	SW_RESTORE
	SW_SHOWDEFAULT
	SW_FORCEMINIMIZE
	SW_NORMAL   = SW_SHOWNORMAL
	SW_MAXIMIZE = SW_SHOWMAXIMIZED
	SW_MAX      = SW_FORCEMINIMIZE
)
const ( // Old
	HIDE_WINDOW = iota
	SHOW_OPENWINDOW
	SHOW_ICONWINDOW
	SHOW_FULLSCREEN
	SHOW_OPENNOACTIVATE
)
const (
	MF_INSERT = 0x00000000
	MF_CHANGE = 0x00000080
	MF_APPEND = 0x00000100
	MF_DELETE = 0x00000200
	MF_REMOVE = 0x00001000

	MF_BYCOMMAND  = 0x00000000
	MF_BYPOSITION = 0x00000400

	MF_SEPARATOR = 0x00000800

	MF_ENABLED  = 0x00000000
	MF_GRAYED   = 0x00000001
	MF_DISABLED = 0x00000002

	MF_UNCHECKED       = 0x00000000
	MF_CHECKED         = 0x00000008
	MF_USECHECKBITMAPS = 0x00000200

	MF_STRING    = 0x00000000
	MF_BITMAP    = 0x00000004
	MF_OWNERDRAW = 0x00000100

	MF_POPUP        = 0x00000010
	MF_MENUBARBREAK = 0x00000020
	MF_MENUBREAK    = 0x00000040

	MF_UNHILITE = 0x00000000
	MF_HILITE   = 0x00000080

	MF_DEFAULT      = 0x00001000
	MF_SYSMENU      = 0x00002000
	MF_HELP         = 0x00004000
	MF_RIGHTJUSTIFY = 0x00004000
	MF_MOUSESELECT  = 0x00008000

	MFT_STRING       = MF_STRING
	MFT_BITMAP       = MF_BITMAP
	MFT_MENUBARBREAK = MF_MENUBARBREAK
	MFT_MENUBREAK    = MF_MENUBREAK
	MFT_OWNERDRAW    = MF_OWNERDRAW
	MFT_RADIOCHECK   = 0x00000200
	MFT_SEPARATOR    = MF_SEPARATOR
	MFT_RIGHTORDER   = 0x00002000
	MFT_RIGHTJUSTIFY = MF_RIGHTJUSTIFY

	MFS_GRAYED    = 0x00000003
	MFS_DISABLED  = MFS_GRAYED
	MFS_CHECKED   = MF_CHECKED
	MFS_HILITE    = MF_HILITE
	MFS_ENABLED   = MF_ENABLED
	MFS_UNCHECKED = MF_UNCHECKED
	MFS_UNHILITE  = MF_UNHILITE
	MFS_DEFAULT   = MF_DEFAULT
)

const (
	FILE_FLAG_FIRST_PIPE_INSTANCE = 0x080000 << iota
	FILE_FLAG_OPEN_NO_RECALL
	FILE_FLAG_OPEN_REPARSE_POINT
	_
	_
	FILE_FLAG_POSIX_SEMANTICS
	FILE_FLAG_BACKUP_SEMANTICS
	FILE_FLAG_DELETE_ON_CLOSE
	FILE_FLAG_SEQUENTIAL_SCAN
	FILE_FLAG_RANDOM_ACCESS
	FILE_FLAG_NO_BUFFERING
	FILE_FLAG_OVERLAPPED
	FILE_FLAG_WRITE_THROUGH
)

const (
	CREATE_NEW = iota + 1
	CREATE_ALWAYS
	OPEN_EXISTING
	OPEN_ALWAYS
	TRUNCATE_EXISTING
)

const (
	FILE_SHARE_READ = 1 << iota
	FILE_SHARE_WRITE
	FILE_SHARE_DELETE
)
const (
	FILE_ATTRIBUTE_READONLY = 1 << iota
	FILE_ATTRIBUTE_HIDDEN
	FILE_ATTRIBUTE_SYSTEM
	_
	FILE_ATTRIBUTE_DIRECTORY
	FILE_ATTRIBUTE_ARCHIVE
	FILE_ATTRIBUTE_DEVICE
	FILE_ATTRIBUTE_NORMAL
	FILE_ATTRIBUTE_TEMPORARY
	FILE_ATTRIBUTE_SPARSE_FILE
	FILE_ATTRIBUTE_REPARSE_POINT
	FILE_ATTRIBUTE_COMPRESSED
	FILE_ATTRIBUTE_OFFLINE
	FILE_ATTRIBUTE_NOT_CONTENT_INDEXED
	FILE_ATTRIBUTE_ENCRYPTED
)
const (
	SECURITY_ANONYMOUS        = SecurityAnonymous << 16
	SECURITY_IDENTIFICATION   = SecurityIdentification << 16
	SECURITY_IMPERSONATION    = SecurityImpersonation << 16
	SECURITY_DELEGATION       = SecurityDelegation << 16
	SECURITY_CONTEXT_TRACKING = 0x40000
	SECURITY_EFFECTIVE_ONLY   = 0x80000
	SECURITY_SQOS_PRESENT     = 0x100000
	SECURITY_VALID_SQOS_FLAGS = 0x1F0000
)
