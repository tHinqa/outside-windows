// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package winbase provides API definitions for accessing
//kernel32.dll nad advapi32.dll.
package winbase

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside/types"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/advapi32"
	_ "github.com/tHinqa/outside/win32/kernel32"
	"unsafe"
)

//GetFreeSpace is obsolete; instead use:
//	GlobalMemoryStatus
func GetFreeSpace() int {
	return 0x100000
}

func InterlockedAnd64(
	destination *T.LONGLONG, value T.LONGLONG) (old T.LONGLONG) {
	for {
		old = *destination
		if InterlockedCompareExchange64(destination,
			old&value, old) == old {
			return
		}
	}
}

func InterlockedCompareExchangePointer(
	destination *T.LONG, exChange uintptr,
	comperand uintptr) uintptr {
	return (uintptr)(unsafe.Pointer(uintptr(InterlockedCompareExchange(destination, T.LONG(exChange), T.LONG(comperand)))))
}

func InterlockedDecrement64(
	destination *T.LONGLONG, value T.LONGLONG) (old T.LONGLONG) {
	for {
		old = *destination
		if InterlockedCompareExchange64(destination,
			old-1, old) == old {
			return
		}
	}
}

func InterlockedExchange64(
	target *T.LONGLONG, value T.LONGLONG) (old T.LONGLONG) {
	for {
		old = *target
		if InterlockedCompareExchange64(target,
			value, old) == old {
			return
		}
	}
}

func InterlockedExchangeAdd64(
	addend *T.LONGLONG, value T.LONGLONG) (old T.LONGLONG) {
	for {
		old = *addend
		if InterlockedCompareExchange64(addend,
			old+value, old) == old {
			return
		}
	}
}

func InterlockedExchangePointer(
	target *T.LONG, value T.LONG) uintptr {
	return (uintptr)(unsafe.Pointer(uintptr(InterlockedExchange(target, value))))
}

func InterlockedIncrement64(
	destination *T.LONGLONG, value T.LONGLONG) (old T.LONGLONG) {
	for {
		old = *destination
		if InterlockedCompareExchange64(destination,
			old+1, old) == old {
			return
		}
	}
}

//BUG: IDs with 0-9 are not DOC'd
func InterlockedOr64(
	destination *T.LONGLONG, value T.LONGLONG) (old T.LONGLONG) {
	for {
		old = *destination
		if InterlockedCompareExchange64(destination,
			old|value, old) == old {
			return
		}
	}
}

func InterlockedXor64(
	destination *T.LONGLONG, value T.LONGLONG) (old T.LONGLONG) {
	for {
		old = *destination
		if InterlockedCompareExchange64(destination,
			old^value, old) == old {
			return
		}
	}
}

//Obsolete
func Yield() {}

type (
	char    T.Fake_type_Fix_me
	va_list T.Fake_type_Fix_me
	long    T.Fake_type_Fix_me
)

var ( //TODO(t): Verify all
	AccessCheck func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		clientToken T.HANDLE,
		desiredAccess T.DWORD,
		genericMapping *T.GENERIC_MAPPING,
		privilegeSet *T.PRIVILEGE_SET,
		privilegeSetLength, grantedAccess *T.DWORD,
		accessStatus *T.BOOL) (T.BOOL, error)

	AccessCheckAndAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		objectTypeName, objectName VString,
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		desiredAccess T.DWORD,
		genericMapping *T.GENERIC_MAPPING,
		objectCreation T.BOOL,
		grantedAccess *T.DWORD,
		accessStatus, generateOnClose *T.BOOL) (T.BOOL, error)

	AccessCheckByType func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		principalSelfSid *T.SID,
		clientToken T.HANDLE,
		desiredAccess T.DWORD,
		objectTypeList *T.OBJECT_TYPE_LIST,
		objectTypeListLength T.DWORD,
		genericMapping *T.GENERIC_MAPPING,
		privilegeSet *T.PRIVILEGE_SET,
		privilegeSetLength, grantedAccess *T.DWORD,
		accessStatus *T.BOOL) (T.BOOL, error)

	AccessCheckByTypeAndAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		objectTypeName,
		objectName VString,
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		principalSelfSid *T.SID,
		desiredAccess T.DWORD,
		auditType T.AUDIT_EVENT_TYPE,
		flags T.DWORD,
		objectTypeList *T.OBJECT_TYPE_LIST,
		objectTypeListLength T.DWORD,
		genericMapping *T.GENERIC_MAPPING,
		objectCreation T.BOOL,
		grantedAccess *T.DWORD,
		accessStatus, generateOnClose *T.BOOL) (T.BOOL, error)

	AccessCheckByTypeResultList func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		principalSelfSid *T.SID,
		clientToken T.HANDLE,
		desiredAccess T.DWORD,
		objectTypeList *T.OBJECT_TYPE_LIST,
		objectTypeListLength T.DWORD,
		genericMapping *T.GENERIC_MAPPING,
		privilegeSet *T.PRIVILEGE_SET,
		privilegeSetLength, grantedAccessList,
		accessStatusList *T.DWORD) (T.BOOL, error)

	AccessCheckByTypeResultListAndAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		objectTypeName, objectName VString,
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		principalSelfSid *T.SID,
		desiredAccess T.DWORD,
		auditType T.AUDIT_EVENT_TYPE,
		flags T.DWORD,
		objectTypeList *T.OBJECT_TYPE_LIST,
		objectTypeListLength T.DWORD,
		genericMapping *T.GENERIC_MAPPING,
		objectCreation T.BOOL,
		grantedAccess, accessStatusList *T.DWORD,
		generateOnClose *T.BOOL) (T.BOOL, error)

	AccessCheckByTypeResultListAndAuditAlarmByHandle func(
		subsystemName VString,
		handleId *T.VOID,
		clientToken T.HANDLE,
		objectTypeName, objectName VString,
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		principalSelfSid *T.SID,
		desiredAccess T.DWORD,
		auditType T.AUDIT_EVENT_TYPE,
		flags T.DWORD,
		objectTypeList *T.OBJECT_TYPE_LIST,
		objectTypeListLength T.DWORD,
		genericMapping *T.GENERIC_MAPPING,
		objectCreation T.BOOL,
		grantedAccess, accessStatusList *T.DWORD,
		generateOnClose *T.BOOL) (T.BOOL, error)

	ActivateActCtx func(
		actCtx T.HANDLE, cookie *T.ULONG_PTR) (T.BOOL, error)

	AddAccessAllowedAce func(
		acl *T.ACL,
		aceRevision, accessMask T.DWORD,
		sid *T.SID) (T.BOOL, error)

	AddAccessAllowedAceEx func(
		acl *T.ACL,
		aceRevision, aceFlags, accessMask T.DWORD,
		sid *T.SID) (T.BOOL, error)

	AddAccessAllowedObjectAce func(
		acl *T.ACL,
		aceRevision, aceFlags, accessMask T.DWORD,
		objectTypeGuid, inheritedObjectTypeGuid *T.GUID,
		sid *T.SID) (T.BOOL, error)

	AddAccessDeniedAce func(
		acl *T.ACL,
		aceRevision, accessMask T.DWORD,
		sid *T.SID) (T.BOOL, error)

	AddAccessDeniedAceEx func(
		acl *T.ACL,
		aceRevision, aceFlags, accessMask T.DWORD,
		sid *T.SID) (T.BOOL, error)

	AddAccessDeniedObjectAce func(
		acl *T.ACL,
		aceRevision, aceFlags, accessMask T.DWORD,
		objectTypeGuid, inheritedObjectTypeGuid *T.GUID,
		sid *T.SID) (T.BOOL, error)

	AddAce func(
		acl *T.ACL,
		aceRevision, startingAceIndex T.DWORD,
		aceList *T.VOID,
		aceListLength T.DWORD) (T.BOOL, error)

	AddAtom func(string_ VString) (T.ATOM, error)

	AddAuditAccessAce func(
		acl *T.ACL,
		aceRevision, accessMask T.DWORD,
		sid *T.SID,
		auditSuccess, auditFailure T.BOOL) (T.BOOL, error)

	AddAuditAccessAceEx func(
		acl *T.ACL,
		aceRevision, aceFlags,
		accessMask T.DWORD,
		sid *T.SID,
		auditSuccess, auditFailure T.BOOL) (T.BOOL, error)

	AddAuditAccessObjectAce func(
		acl *T.ACL,
		aceRevision, aceFlags, accessMask,
		objectTypeGuid, inheritedObjectTypeGuid *T.GUID,
		sid *T.SID,
		auditSuccess, auditFailure T.BOOL) (T.BOOL, error)

	AddRefActCtx func(actCtx T.HANDLE)

	AddVectoredContinueHandler func(
		first T.ULONG,
		handler T.VECTORED_EXCEPTION_HANDLER) (*T.VOID, error)

	AddVectoredExceptionHandler func(
		first T.ULONG,
		handler T.VECTORED_EXCEPTION_HANDLER) (*T.VOID, error)

	AdjustTokenGroups func(
		tokenHandle T.HANDLE,
		resetToDefault T.BOOL,
		newState *T.TOKEN_GROUPS,
		bufferLength T.DWORD,
		previousState *T.TOKEN_GROUPS,
		returnLength *T.DWORD) (T.BOOL, error)

	AdjustTokenPrivileges func(
		tokenHandle T.HANDLE,
		disableAllPrivileges T.BOOL,
		newState *T.TOKEN_PRIVILEGES,
		bufferLength T.DWORD,
		previousState *T.TOKEN_PRIVILEGES,
		returnLength *T.DWORD) (T.BOOL, error)

	AllocateAndInitializeSid func(
		identifierAuthority *T.SID_IDENTIFIER_AUTHORITY,
		subAuthorityCount T.BYTE,
		subAuthority0, subAuthority1,
		subAuthority2, subAuthority3,
		subAuthority4, subAuthority5,
		subAuthority6, subAuthority7 T.DWORD,
		sid **T.SID) (T.BOOL, error)

	AllocateLocallyUniqueId func(luid *T.LUID) (T.BOOL, error)

	AllocateUserPhysicalPages func(
		process T.HANDLE,
		numberOfPages, pageArray *T.ULONG_PTR) (T.BOOL, error)

	AreAllAccessesGranted func(
		grantedAccess, desiredAccess T.DWORD) (T.BOOL, error)

	AreAnyAccessesGranted func(
		grantedAccess, desiredAccess T.DWORD) (T.BOOL, error)

	AreFileApisANSI func()

	AssignProcessToJobObject func(job, process T.HANDLE) (T.BOOL, error)

	BackupEventLog func(
		eventLog T.HANDLE, backupFileName VString) (T.BOOL, error)

	BackupRead func(
		file T.HANDLE,
		buffer *T.BYTE,
		numberOfBytesToRead T.DWORD,
		numberOfBytesRead *T.DWORD,
		abort, processSecurity T.BOOL,
		context **T.VOID) (T.BOOL, error)

	BackupSeek func(
		file T.HANDLE,
		lowBytesToSeek, highBytesToSeek T.DWORD,
		lowByteSeeked, highByteSeeked *T.DWORD,
		context **T.VOID) (T.BOOL, error)

	BackupWrite func(
		file T.HANDLE,
		buffer *T.BYTE,
		numberOfBytesToWrite T.DWORD,
		numberOfBytesWritten *T.DWORD,
		abort, processSecurity T.BOOL,
		context **T.VOID) (T.BOOL, error)

	Beep func(freq, duration T.DWORD) (T.BOOL, error)

	BeginUpdateResource func(
		fileName VString,
		deleteExistingResources T.BOOL) (T.HANDLE, error)

	BindIoCompletionCallback func(
		fileHandle T.HANDLE,
		function T.OVERLAPPED_COMPLETION_ROUTINE,
		flags T.ULONG) (T.BOOL, error)

	BuildCommDCB func(def VString, dcb *T.DCB) (T.BOOL, error)

	BuildCommDCBAndTimeouts func(
		def VString,
		dcb *T.DCB,
		commTimeouts *T.COMMTIMEOUTS) (T.BOOL, error)

	CallNamedPipe func(
		namedPipeName VString,
		inBuffer *T.VOID,
		inBufferSize T.DWORD,
		outBuffer *T.VOID,
		outBufferSize T.DWORD,
		bytesRead *T.DWORD,
		timeOut T.DWORD) (T.BOOL, error)

	CancelDeviceWakeupRequest func(device T.HANDLE) (T.BOOL, error)

	CancelIo func(file T.HANDLE) (T.BOOL, error)

	CancelTimerQueueTimer func(
		timerQueue, timer T.HANDLE) (T.BOOL, error)

	CancelWaitableTimer func(timer T.HANDLE) (T.BOOL, error)

	ChangeTimerQueueTimer func(
		timerQueue, timer T.HANDLE,
		dueTime, period T.ULONG) (T.BOOL, error)

	CheckNameLegalDOS8Dot3 func(
		name VString,
		oemName T.OAString,
		oemNameSize T.DWORD,
		nameContainsSpaces, nameLegal *T.BOOL) (T.BOOL, error)

	CheckRemoteDebuggerPresent func(
		process T.HANDLE, debuggerPresent *T.BOOL) (T.BOOL, error)

	CheckTokenMembership func(
		tokenHandleHANDLE,
		sidToCheck *T.SID,
		isMember *T.BOOL) (T.BOOL, error)

	ClearCommBreak func(file T.HANDLE) (T.BOOL, error)

	ClearCommError func(
		file T.HANDLE, errors *T.DWORD, stat *T.COMSTAT) (T.BOOL, error)

	ClearEventLog func(
		eventLog T.HANDLE, backupFileName VString) (T.BOOL, error)

	CloseEncryptedFileRaw func(context *T.VOID)

	CloseEventLog func(eventLog T.HANDLE) (T.BOOL, error)

	CloseHandle func(object T.HANDLE) (T.BOOL, error)

	CommConfigDialog func(
		name VString, wnd T.HWND, cc *T.COMMCONFIG) (T.BOOL, error)

	CompareFileTime func(
		fileTime1, fileTime2 *T.FILETIME) (T.LONG, error)

	ConnectNamedPipe func(
		namedPipe T.HANDLE, overlapped *T.OVERLAPPED) (T.BOOL, error)

	ContinueDebugEvent func(
		processId, threadId, continueStatus T.DWORD) (T.BOOL, error)

	ConvertFiberToThread func() (T.BOOL, error)

	ConvertThreadToFiber func(parameter *T.VOID) (*T.VOID, error)

	ConvertThreadToFiberEx func(
		parameter *T.VOID, flags T.DWORD) (*T.VOID, error)

	ConvertToAutoInheritPrivateObjectSecurity func(
		parentDescriptor,
		currentSecurityDescriptor *T.SECURITY_DESCRIPTOR,
		newSecurityDescriptor **T.SECURITY_DESCRIPTOR,
		objectType *T.GUID,
		isDirectoryObject T.BOOLEAN,
		genericMapping *T.GENERIC_MAPPING) (T.BOOL, error)

	CopyFile func(
		existingFileName, newFileName VString,
		failIfExists T.BOOL) (T.BOOL, error)

	CopyFileEx func(
		existingFileName, newFileName VString,
		progressRoutine T.PROGRESS_ROUTINE,
		data *T.VOID,
		cancel *T.BOOL,
		copyFlags T.DWORD) (T.BOOL, error)

	CopySid func(
		destinationLength T.DWORD,
		destination, source *T.SID) (T.BOOL, error)

	CreateActCtx func(actCtx T.ACTCTX) (T.HANDLE, error)

	CreateDirectory func(
		pathName VString,
		securityAttributes *T.SECURITY_ATTRIBUTES) (T.BOOL, error)

	CreateDirectoryEx func(
		templateDirectory, newDirectory VString,
		securityAttributes *T.SECURITY_ATTRIBUTES) (T.BOOL, error)

	CreateEvent func(
		eventAttributes *T.SECURITY_ATTRIBUTES,
		manualReset, initialState T.BOOL,
		name VString) (T.HANDLE, error)

	CreateFiber func(
		stackSize T.SIZE_T,
		startAddress T.FIBER_START_ROUTINE,
		parameter *T.VOID) (*T.VOID, error)

	CreateFiberEx func(
		stackCommitSize, stackReserveSize T.SIZE_T,
		flags T.DWORD,
		startAddress T.FIBER_START_ROUTINE,
		parameter *T.VOID) (*T.VOID, error)

	CreateFile func(
		fileName VString,
		desiredAccess, shareMode T.DWORD,
		securityAttributes *T.SECURITY_ATTRIBUTES,
		creationDisposition, flagsAndAttributes T.DWORD,
		templateFile T.HANDLE) (T.HANDLE, error)

	CreateFileMapping func(
		file T.HANDLE,
		fileMappingAttributes *T.SECURITY_ATTRIBUTES,
		protect, maximumSizeHigh, maximumSizeLow T.DWORD,
		name VString) (T.HANDLE, error)

	CreateHardLink func(
		fileName, existingFileName VString,
		securityAttributes *T.SECURITY_ATTRIBUTES) (T.BOOL, error)

	CreateIoCompletionPort func(
		fileHandle, existingCompletionPort T.HANDLE,
		completionKey T.ULONG_PTR,
		numberOfConcurrentThreads T.DWORD) (T.HANDLE, error)

	CreateJobObject func(
		jobAttributes *T.SECURITY_ATTRIBUTES,
		name VString) (T.HANDLE, error)

	CreateJobSet func(
		numJob T.ULONG,
		userJobSet *T.JOB_SET_ARRAY,
		flags T.ULONG) (T.BOOL, error)

	CreateMailslot func(
		name VString,
		maxMessageSize, readTimeout T.DWORD,
		securityAttributes *T.SECURITY_ATTRIBUTES) (T.HANDLE, error)

	CreateMemoryResourceNotification func(
		T.MEMORY_RESOURCE_NOTIFICATION_TYPE) (T.HANDLE, error)

	CreateMutex func(
		mutexAttributes *T.SECURITY_ATTRIBUTES,
		initialOwner T.BOOL,
		name VString) (T.HANDLE, error)

	CreateNamedPipe func(
		name VString,
		openMode, pipeMode, maxInstances,
		outBufferSize, inBufferSize, defaultTimeOut T.DWORD,
		securityAttributes *T.SECURITY_ATTRIBUTES) (T.HANDLE, error)

	CreatePipe func(
		readPipe, writePipe *T.HANDLE,
		pipeAttributes *T.SECURITY_ATTRIBUTES,
		size T.DWORD) (T.BOOL, error)

	CreatePrivateObjectSecurity func(
		parentDescriptor,
		creatorDescriptor *T.SECURITY_DESCRIPTOR,
		newDescriptor **T.SECURITY_DESCRIPTOR,
		isDirectoryObject T.BOOL,
		token T.HANDLE,
		genericMapping *T.GENERIC_MAPPING) (T.BOOL, error)

	CreatePrivateObjectSecurityEx func(
		parentDescriptor,
		creatorDescriptor *T.SECURITY_DESCRIPTOR,
		newDescriptor **T.SECURITY_DESCRIPTOR,
		objectType *T.GUID,
		isContainerObject T.BOOL,
		autoInheritFlags T.ULONG,
		token T.HANDLE,
		genericMapping *T.GENERIC_MAPPING) (T.BOOL, error)

	CreatePrivateObjectSecurityWithMultipleInheritance func(
		parentDescriptor, creatorDescriptor,
		newDescriptor **T.SECURITY_DESCRIPTOR,
		objectTypes **T.GUID,
		guidCount T.ULONG,
		isContainerObject T.BOOL,
		autoInheritFlags T.ULONG,
		token T.HANDLE,
		genericMapping *T.GENERIC_MAPPING) (T.BOOL, error)

	CreateProcess func(
		applicationName, commandLine VString,
		processAttributes,
		threadAttributes *T.SECURITY_ATTRIBUTES,
		inheritHandles T.BOOL,
		creationFlags T.DWORD,
		environment *T.VOID,
		currentDirectory VString,
		startupInfo *T.STARTUPINFO,
		processInformation *T.PROCESS_INFORMATION) (T.BOOL, error)

	CreateProcessAsUser func(
		token T.HANDLE,
		applicationName, commandLine VString,
		processAttributes,
		threadAttributes *T.SECURITY_ATTRIBUTES,
		inheritHandles T.BOOL,
		creationFlags T.DWORD,
		environment *T.VOID,
		currentDirectory VString,
		startupInfo *T.STARTUPINFO,
		processInformation *T.PROCESS_INFORMATION) (T.BOOL, error)

	CreateProcessWithLogonW func(
		username, domain, password T.WString,
		logonFlags T.DWORD,
		applicationName T.WString,
		commandLine T.OWString,
		creationFlags T.DWORD,
		environment *T.VOID,
		currentDirectory T.WString,
		startupInfo *T.STARTUPINFO,
		processInformation *T.PROCESS_INFORMATION) (T.BOOL, error)

	CreateProcessWithTokenW func(
		token T.HANDLE,
		logonFlags T.DWORD,
		applicationName T.WString,
		commandLine T.OWString,
		creationFlags T.DWORD,
		environment *T.VOID,
		currentDirectory T.OWString,
		startupInfo *T.STARTUPINFO,
		processInformation *T.PROCESS_INFORMATION) (T.BOOL, error)

	CreateRemoteThread func(
		process T.HANDLE,
		threadAttributes *T.SECURITY_ATTRIBUTES,
		stackSize T.SIZE_T,
		startAddress T.THREAD_START_ROUTINE,
		parameter *T.VOID,
		creationFlags T.DWORD,
		threadId *T.DWORD) (T.HANDLE, error)

	CreateRestrictedToken func(
		existingTokenHandle T.HANDLE,
		flags, disableSidCount, deletePrivilegeCount,
		restrictedSidCount T.DWORD,
		newTokenHandle *T.HANDLE) (T.BOOL, error)

	CreateSemaphore func(
		semaphoreAttributes *T.SECURITY_ATTRIBUTES,
		initialCount, maximumCount T.LONG,
		name VString) (T.HANDLE, error)

	CreateTapePartition func(
		device T.HANDLE,
		partitionMethod, count, size T.DWORD) (T.DWORD, error)

	CreateThread func(
		threadAttributes *T.SECURITY_ATTRIBUTES,
		stackSize T.SIZE_T,
		startAddress T.THREAD_START_ROUTINE,
		parameter *T.VOID,
		creationFlags T.DWORD,
		threadId *T.DWORD) (T.HANDLE, error)

	CreateTimerQueue func() (T.HANDLE, error)

	CreateTimerQueueTimer func(
		newTimer *T.HANDLE,
		timerQueue T.HANDLE,
		callback T.WAITORTIMERCALLBACK,
		parameter *T.VOID,
		dueTime, period T.DWORD,
		flags T.ULONG) (T.BOOL, error)

	CreateWaitableTimer func(
		timerAttributes *T.SECURITY_ATTRIBUTES,
		manualReset T.BOOL,
		timerName VString) (T.HANDLE, error)

	CreateWellKnownSid func(
		wellKnownSidType T.WELL_KNOWN_SID_TYPE,
		domainSid, sid *T.SID,
		sidl *T.DWORD) (T.BOOL, error)

	DeactivateActCtx func(
		flags T.DWORD, cookie T.ULONG_PTR) (T.BOOL, error)

	DebugActiveProcess func(processId T.DWORD) (T.BOOL, error)

	DebugActiveProcessStop func(processId T.DWORD) (T.BOOL, error)

	DebugBreak func()

	DebugBreakProcess func(process T.HANDLE) (T.BOOL, error)

	DebugSetProcessKillOnExit func(killOnExit T.BOOL) (T.BOOL, error)

	DecodePointer func(ptr *T.VOID) (*T.VOID, error)

	DecodeSystemPointer func(ptr *T.VOID) (*T.VOID, error)

	DecryptFile func(fileName VString, reserved T.DWORD) (T.BOOL, error)

	DefineDosDevice func(
		flags T.DWORD,
		deviceName, targetPath VString) (T.BOOL, error)

	DeleteAce func(acl *T.ACL, aceIndex T.DWORD) (T.BOOL, error)

	DeleteAtom func(atom T.ATOM) (T.ATOM, error)

	DeleteCriticalSection func(*T.CRITICAL_SECTION)

	DeleteFiber func(fiber *T.VOID)

	DeleteFile func(fileName VString) (T.BOOL, error)

	DeleteTimerQueue func(timerQueue T.HANDLE) (T.BOOL, error)

	DeleteTimerQueueEx func(
		timerQueue, completionEvent T.HANDLE) (T.BOOL, error)

	DeleteTimerQueueTimer func(
		timerQueue, timer, completionEvent T.HANDLE) (T.BOOL, error)

	DeleteVolumeMountPoint func(
		volumeMountPoint VString) (T.BOOL, error)

	DeregisterEventSource func(eventLog T.HANDLE) (T.BOOL, error)

	DestroyPrivateObjectSecurity func(
		objectDescriptor **T.SECURITY_DESCRIPTOR) (T.BOOL, error)

	DeviceIoControl func(
		device T.HANDLE,
		ioControlCode T.DWORD,
		inBuffer *T.VOID,
		inBufferSize T.DWORD,
		outBuffer *T.VOID,
		outBufferSize, bytesReturned *T.DWORD,
		overlapped *T.OVERLAPPED) (T.BOOL, error)

	DisableThreadLibraryCalls func(libModule T.HMODULE) (T.BOOL, error)

	DisconnectNamedPipe func(namedPipe T.HANDLE) (T.BOOL, error)

	DnsHostnameToComputerName func(
		hostname VString,
		computerName OVString,
		size *T.DWORD) (T.BOOL, error)

	DosDateTimeToFileTime func(
		fatDate, fatTime T.WORD, fileTime *T.FILETIME) (T.BOOL, error)

	DuplicateHandle func(
		sourceProcessHandle, sourceHandle,
		targetProcessHandle T.HANDLE,
		targetHandle *T.HANDLE,
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		options T.DWORD) (T.BOOL, error)

	DuplicateToken func(
		existingTokenHandle T.HANDLE,
		impersonationLevel T.SECURITY_IMPERSONATION_LEVEL,
		duplicateTokenHandle *T.HANDLE) (T.BOOL, error)

	DuplicateTokenEx func(
		existingToken T.HANDLE,
		desiredAccess T.DWORD,
		tokenAttributes *T.SECURITY_ATTRIBUTES,
		impersonationLevel T.SECURITY_IMPERSONATION_LEVEL,
		tokenType T.TOKEN_TYPE,
		newToken *T.HANDLE) (T.BOOL, error)

	EncodePointer func(ptr *T.VOID) (*T.VOID, error)

	EncodeSystemPointer func(ptr *T.VOID) (*T.VOID, error)

	EncryptFile func(fileName VString) (T.BOOL, error)

	EndUpdateResource func(
		update T.HANDLE, discard T.BOOL) (T.BOOL, error)

	EnterCriticalSection func(*T.CRITICAL_SECTION)

	EnumResourceLanguages func(
		module T.HMODULE,
		type_, name VString,
		enumFunc T.ENUMRESLANGPROC,
		param T.LONG_PTR) (T.BOOL, error)

	EnumResourceNames func(
		module T.HMODULE,
		type_ VString,
		enumFunc T.ENUMRESNAMEPROC,
		param T.LONG_PTR) (T.BOOL, error)

	EnumResourceTypes func(
		module T.HMODULE,
		enumFunc T.ENUMRESTYPEPROC,
		param T.LONG_PTR) (T.BOOL, error)

	EnumSystemFirmwareTables func(
		firmwareTableProviderSignature T.DWORD,
		firmwareTableEnumBuffer *T.VOID,
		bufferSize T.DWORD) (T.UINT, error)

	EqualDomainSid func(
		sid1, sid2 *T.SID, equal *T.BOOL) (T.BOOL, error)

	EqualPrefixSid func(
		sid1, sid2 *T.SID) (T.BOOL, error)

	EqualSid func(sid1, sid2 *T.SID) (T.BOOL, error)

	EraseTape func(
		device T.HANDLE,
		eraseType T.DWORD,
		immediate T.BOOL) (T.DWORD, error)

	EscapeCommFunction func(file T.HANDLE, function T.DWORD) (T.BOOL, error)

	ExitProcess func(exitCode T.UINT)

	ExitThread func(exitCode T.DWORD)

	ExpandEnvironmentStrings func(
		src, dst VString, size T.DWORD) (T.DWORD, error)

	FatalAppExit func(action T.UINT, messageText VString)

	FatalExit func(exitCode int)

	FileEncryptionStatus func(
		fileName VString, status *T.DWORD) (T.BOOL, error)

	FileTimeToDosDateTime func(
		fileTime *T.FILETIME, fatDate, fatTime *T.WORD) (T.BOOL, error)

	FileTimeToLocalFileTime func(
		fileTime, localFileTime *T.FILETIME) (T.BOOL, error)

	FileTimeToSystemTime func(
		fileTime, systemTime *T.SYSTEMTIME) (T.BOOL, error)

	FindActCtxSectionGuid func(
		flags T.DWORD,
		extensionGuid *T.GUID,
		sectionId T.ULONG,
		guidToFind *T.GUID,
		returnedData *T.ACTCTX_SECTION_KEYED_DATA) (T.BOOL, error)

	FindActCtxSectionString func(
		flags T.DWORD,
		extensionGuid *T.GUID,
		sectionId T.ULONG,
		string_ToFind VString,
		returnedData *T.ACTCTX_SECTION_KEYED_DATA) (T.BOOL, error)

	FindAtom func(string_ VString) (T.ATOM, error)

	FindClose func(findFile T.HANDLE) (T.BOOL, error)

	FindCloseChangeNotification func(
		changeHandle T.HANDLE) (T.BOOL, error)

	FindFirstChangeNotification func(
		pathName VString,
		watchSubtree T.BOOL,
		notifyFilter T.DWORD) (T.HANDLE, error)

	FindFirstFileA func(
		fileName T.AString,
		findFileData *T.WIN32_FIND_DATAA) (T.HANDLE, error)

	FindFirstFileW func(
		fileName T.WString,
		findFileData *T.WIN32_FIND_DATAW) (T.HANDLE, error)

	FindFirstFileEx func(
		fileName VString,
		infoLevelId T.FINDEX_INFO_LEVELS,
		findFileData *T.VOID,
		searchOp T.FINDEX_SEARCH_OPS,
		searchFilter *T.VOID,
		additionalFlags T.DWORD) (T.HANDLE, error)

	FindFirstFreeAce func(acl *T.ACL, ace **T.VOID) (T.BOOL, error)

	FindFirstVolume func(
		volumeName OVString, bufferLength T.DWORD) (T.HANDLE, error)

	FindFirstVolumeMountPoint func(
		rootPathName VString,
		volumeMountPoint OVString,
		bufferLength T.DWORD) (T.HANDLE, error)

	FindFirstStreamW func(
		fileName T.WString,
		infoLevel T.STREAM_INFO_LEVELS,
		findStreamData *T.VOID,
		flags T.DWORD) (T.HANDLE, error)

	FindNextStreamW func(
		findStream T.HANDLE, data *T.VOID) (T.BOOL, error)

	FindNextChangeNotification func(
		changeHandle T.HANDLE) (T.BOOL, error)

	FindNextFileA func(
		findFile T.HANDLE, data *T.WIN32_FIND_DATAA) (T.BOOL, error)

	FindNextFileW func(
		findFile T.HANDLE, data *T.WIN32_FIND_DATAW) (T.BOOL, error)

	FindNextVolume func(
		findVolume T.HANDLE,
		volumeName OVString,
		bufferLength T.DWORD) (T.BOOL, error)

	FindNextVolumeMountPoint func(
		findVolumeMountPoint T.HANDLE,
		volumeMountPoint OVString,
		bufferLength T.DWORD) (T.BOOL, error)

	FindResource func(
		module T.HMODULE,
		name, typ VString) (T.HRSRC, error)

	FindResourceEx func(
		module T.HMODULE,
		typ, name VString,
		language T.WORD) (T.HRSRC, error)

	FindVolumeClose func(findVolume T.HANDLE) (T.BOOL, error)

	FindVolumeMountPointClose func(
		findVolumeMountPoint T.HANDLE) (T.BOOL, error)

	FlsAlloc func(callback T.FLS_CALLBACK_FUNCTION) (T.DWORD, error)

	FlsFree func(flsIndex T.DWORD) (T.BOOL, error)

	FlsGetValue func(flsIndex T.DWORD) (*T.VOID, error)

	FlsSetValue func(flsIndex T.DWORD, flsData *T.VOID) (T.BOOL, error)

	FlushFileBuffers func(file T.HANDLE) (T.BOOL, error)

	FlushInstructionCache func(
		process T.HANDLE, baseAddress *T.VOID, size T.SIZE_T) (T.BOOL, error)

	FlushViewOfFile func(
		baseAddress *T.VOID, numberOfBytesToFlush T.SIZE_T) (T.BOOL, error)

	FormatMessage func(
		flags T.DWORD,
		source *T.VOID,
		messageId, languageId T.DWORD,
		buffer VString,
		size T.DWORD,
		arguments *va_list) (T.DWORD, error)

	FreeEnvironmentStrings func(VString) (T.BOOL, error)

	FreeLibrary func(libModule T.HMODULE) (T.BOOL, error)

	FreeLibraryAndExitThread func(
		libModule T.HMODULE, exitCode T.DWORD)

	FreeResource func(resData T.HGLOBAL) (T.BOOL, error)

	FreeSid func(sid *T.SID) (*T.VOID, error)

	FreeUserPhysicalPages func(
		process T.HANDLE,
		numberOfPages, pageArray *T.ULONG_PTR) (T.BOOL, error)

	GetAce func(
		acl *T.ACL, aceIndex T.DWORD, ace **T.VOID) (T.BOOL, error)

	GetAclInformation func(
		acl *T.ACL,
		info *T.VOID,
		length T.DWORD,
		class T.ACL_INFORMATION_CLASS) (T.BOOL, error)

	GetAtomName func(
		atom T.ATOM, buffer VString, size int) (T.UINT, error)

	GetBinaryType func(
		applicationName VString, binaryType *T.DWORD) (T.BOOL, error)

	GetCommandLine func() (VString, error)

	GetCommConfig func(
		commDev T.HANDLE, cc *T.COMMCONFIG, size *T.DWORD) (T.BOOL, error)

	GetCommMask func(file T.HANDLE, evtMask *T.DWORD) (T.BOOL, error)

	GetCommModemStatus func(
		file T.HANDLE, modemStat *T.DWORD) (T.BOOL, error)

	GetCommProperties func(
		file T.HANDLE, cp *T.COMMPROP) (T.BOOL, error)

	GetCommState func(file T.HANDLE, dcn *T.DCB) (T.BOOL, error)

	GetCommTimeouts func(
		file T.HANDLE, commTimeouts *T.COMMTIMEOUTS) (T.BOOL, error)

	GetCompressedFileSize func(
		fileName VString, fileSizeHigh *T.DWORD) (T.DWORD, error)

	GetComputerName func(
		buffer OVString, size *T.DWORD) (T.BOOL, error)

	GetComputerNameEx func(
		nameType T.COMPUTER_NAME_FORMAT,
		buffer OVString,
		size *T.DWORD) (T.BOOL, error)

	GetCurrentActCtx func(actCtx *T.HANDLE) (T.BOOL, error)

	GetCurrentDirectory func(
		bufferLength T.DWORD, buffer VString) (T.DWORD, error)

	GetCurrentHwProfileA func(
		hwProfileInfo *T.HW_PROFILE_INFOA) (T.BOOL, error)

	GetCurrentHwProfileW func(
		hwProfileInfo *T.HW_PROFILE_INFOW) (T.BOOL, error)

	GetCurrentProcess func() (T.HANDLE, error)

	GetCurrentProcessId func() (T.DWORD, error)

	GetCurrentProcessorNumber func() (T.DWORD, error)

	GetCurrentThread func() (T.HANDLE, error)

	GetCurrentThreadId func() (T.DWORD, error)

	GetDefaultCommConfig func(
		name VString, cc *T.COMMCONFIG, size *T.DWORD) (T.BOOL, error)

	GetDevicePowerState func(
		device T.HANDLE, on *T.BOOL) (T.BOOL, error)

	GetDiskFreeSpace func(
		rootPathName VString,
		sectorsPerCluster,
		bytesPerSector,
		numberOfFreeClusters,
		totalNumberOfClusters *T.DWORD) (T.BOOL, error)

	GetDiskFreeSpaceEx func(
		directoryName VString,
		freeBytesAvailableToCaller,
		totalNumberOfBytes,
		totalNumberOfFreeBytes *T.ULARGE_INTEGER) (T.BOOL, error)

	GetDllDirectory func(
		bufferLength T.DWORD, buffer VString) (T.DWORD, error)

	GetDriveType func(rootPathName VString) (T.UINT, error)

	GetEnvironmentStrings func() (T.OMVString, error)

	GetEnvironmentVariable func(
		name, buffer VString, size T.DWORD) (T.DWORD, error)

	GetEventLogInformation func(
		eventLog T.HANDLE,
		infoLevel T.DWORD,
		buffer *T.VOID,
		bufSize T.DWORD,
		bytesNeeded *T.DWORD) (T.BOOL, error)

	GetExitCodeProcess func(process T.HANDLE, exitCode *T.DWORD) (T.BOOL, error)

	GetExitCodeThread func(thread T.HANDLE, exitCode *T.DWORD) (T.BOOL, error)

	GetFileAttributes func(fileName VString) (T.DWORD, error)

	GetFileAttributesEx func(
		fileName VString,
		infoLevelId T.GET_FILEEX_INFO_LEVELS,
		fileInformation *T.VOID) (T.BOOL, error)

	GetFileInformationByHandle func(
		file T.HANDLE,
		fileInformation *T.BY_HANDLE_FILE_INFORMATION) (T.BOOL, error)

	GetFileSecurity func(
		fileName VString,
		requestedInformation T.SECURITY_INFORMATION,
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		length T.DWORD,
		lengthNeeded *T.DWORD) (T.BOOL, error)

	GetFileSize func(
		file T.HANDLE, fileSizeHigh *T.DWORD) (T.DWORD, error)

	GetFileSizeEx func(
		file T.HANDLE, fileSize *T.LARGE_INTEGER) (T.BOOL, error)

	GetFileTime func(
		file T.HANDLE,
		creationTime, lastAccessTime,
		lastWriteTime *T.FILETIME) (T.BOOL, error)

	GetFileType func(file T.HANDLE) (T.DWORD, error)

	GetFirmwareEnvironmentVariable func(
		name, guid VString, buffer *T.VOID, size T.DWORD) (T.DWORD, error)

	GetFullPathName func(
		fileName VString,
		bufferLength T.DWORD,
		buffer VString,
		filePart OVString) (T.DWORD, error)

	GetHandleInformation func(
		object T.HANDLE, flags *T.DWORD) (T.BOOL, error)

	GetKernelObjectSecurity func(
		handle T.HANDLE,
		requestedInformation T.SECURITY_INFORMATION,
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		length T.DWORD,
		lengthNeeded *T.DWORD) (T.BOOL, error)

	GetLargePageMinimum func() (T.SIZE_T, error)

	GetLastError func() (T.DWORD, error)

	GetLengthSid func(sid *T.SID) (T.DWORD, error)

	GetLocalTime func(systemTime *T.SYSTEMTIME)

	GetLogicalDrives func() (T.DWORD, error)

	GetLogicalDriveStrings func(
		bufferLength T.DWORD,
		buffer T.OMVString) (T.DWORD, error)

	GetLogicalProcessorInformation func(
		buffer *T.SYSTEM_LOGICAL_PROCESSOR_INFORMATION,
		returnedLength *T.DWORD) (T.BOOL, error)

	GetLongPathName func(
		shortPath VString,
		longPath OVString,
		buffer T.DWORD) (T.DWORD, error)

	GetMailslotInfo func(
		mailslot T.HANDLE,
		maxMessageSize, nextSize, messageCount,
		readTimeout *T.DWORD) (T.BOOL, error)

	GetModuleFileName func(
		module T.HMODULE, filename VString, size T.DWORD) (T.DWORD, error)

	GetModuleHandle func(moduleName VString) (T.HMODULE, error)

	GetModuleHandleEx func(
		flags T.DWORD,
		moduleName VString,
		module *T.HMODULE) (T.BOOL, error)

	GetNamedPipeHandleState func(
		namedPipe T.HANDLE,
		state, curInstances, maxCollectionCount,
		collectDataTimeout *T.DWORD,
		userName VString,
		maxUserNameSize T.DWORD) (T.BOOL, error)

	GetNamedPipeInfo func(
		namedPipe T.HANDLE,
		flags, outBufferSize, inBufferSize,
		maxInstances *T.DWORD) (T.BOOL, error)

	GetNativeSystemInfo func(systemInfo *T.SYSTEM_INFO)

	GetNumaAvailableMemoryNode func(
		node T.WChar,
		availableBytes *T.ULONGLONG) (T.BOOL, error)

	GetNumaHighestNodeNumber func(
		highestNodeNumber *T.ULONG) (T.BOOL, error)

	GetNumaNodeProcessorMask func(
		node T.WChar,
		processorMask *T.ULONGLONG) (T.BOOL, error)

	GetNumaProcessorNode func(
		processor T.WChar,
		nodeNumber *T.WChar) (T.BOOL, error)

	GetNumberOfEventLogRecords func(
		eventLog T.HANDLE,
		numberOfRecords *T.DWORD) (T.BOOL, error)

	GetOldestEventLogRecord func(
		eventLog T.HANDLE,
		oldestRecord *T.DWORD) (T.BOOL, error)

	GetOverlappedResult func(
		file T.HANDLE,
		overlapped *T.OVERLAPPED,
		numberOfBytesTransferred *T.DWORD,
		wait T.BOOL) (T.BOOL, error)

	GetPriorityClass func(process T.HANDLE) (T.DWORD, error)

	GetPrivateObjectSecurity func(
		objectDescriptor *T.SECURITY_DESCRIPTOR,
		securityInformation T.SECURITY_INFORMATION,
		resultantDescriptor *T.SECURITY_DESCRIPTOR,
		descriptorLength T.DWORD,
		returnLength *T.DWORD) (T.BOOL, error)

	GetPrivateProfileInt func(
		appName, keyName VString,
		default_ T.INT,
		fileName VString) (T.UINT, error)

	GetPrivateProfileSection func(
		appName, returnedString VString,
		size T.DWORD,
		fileName VString) (T.DWORD, error)

	GetPrivateProfileSectionNames func(
		returnBuffer VString,
		size T.DWORD,
		fileName VString) (T.DWORD, error)

	GetPrivateProfileString func(
		appName, keyName, default_, returned VString,
		size T.DWORD,
		fileName VString) (T.DWORD, error)

	GetPrivateProfileStruct func(
		section, key VString,
		struct_ *T.VOID,
		sizeStruct T.UINT,
		file VString) (T.BOOL, error)

	GetProcAddress func(
		module T.HMODULE, procName T.AString) (T.FARPROC, error)

	GetProcessAffinityMask func(
		process T.HANDLE,
		processAffinityMask,
		systemAffinityMask *T.DWORD_PTR) (T.BOOL, error)

	GetProcessHandleCount func(
		process T.HANDLE, handleCount *T.DWORD) (T.BOOL, error)

	GetProcessHeap func() (T.HANDLE, error)

	GetProcessHeaps func(
		numberOfHeaps T.DWORD,
		processHeaps *T.HANDLE) (T.DWORD, error)

	GetProcessId func(process T.HANDLE) (T.DWORD, error)

	GetProcessIdOfThread func(thread T.HANDLE) (T.DWORD, error)

	GetProcessIoCounters func(
		process T.HANDLE, ioCounters *T.IO_COUNTERS) (T.BOOL, error)

	GetProcessPriorityBoost func(
		process T.HANDLE, disablePriorityBoost *T.BOOL) (T.BOOL, error)

	GetProcessShutdownParameters func(
		level, flags *T.DWORD) (T.BOOL, error)

	GetProcessTimes func(
		process T.HANDLE,
		creationTime, exitTime, kernelTime,
		userTime *T.FILETIME) (T.BOOL, error)

	GetProcessVersion func(
		processId T.DWORD) (T.DWORD, error)

	GetProcessWorkingSetSize func(
		process T.HANDLE,
		minWorkingSetSize, maxWorkingSetSize *T.SIZE_T) (T.BOOL, error)

	GetProcessWorkingSetSizeEx func(
		process T.HANDLE,
		minWorkingSetSize, maxWorkingSetSize *T.SIZE_T,
		flags *T.DWORD) (T.BOOL, error)

	GetProfileInt func(
		appName, keyName VString,
		default_ T.INT) (T.UINT, error)

	GetProfileSection func(
		appName, returnedString VString,
		size T.DWORD) (T.DWORD, error)

	GetProfileString func(
		appName, keyName, default_, returned VString,
		size T.DWORD) (T.DWORD, error)

	GetQueuedCompletionStatus func(
		completionPort T.HANDLE,
		numberOfBytesTransferred *T.DWORD,
		completionKey *T.ULONG_PTR,
		overlapped **T.OVERLAPPED,
		milliseconds T.DWORD) (T.BOOL, error)

	GetSecurityDescriptorControl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		control *T.SECURITY_DESCRIPTOR_CONTROL,
		revision *T.DWORD) (T.BOOL, error)

	GetSecurityDescriptorDacl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		daclPresent *T.BOOL,
		dacl **T.ACL,
		daclDefaulted *T.BOOL) (T.BOOL, error)

	GetSecurityDescriptorGroup func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		group **T.SID,
		groupDefaulted *T.BOOL) (T.BOOL, error)

	GetSecurityDescriptorLength func(
		securityDescriptor *T.SECURITY_DESCRIPTOR) (T.DWORD, error)

	GetSecurityDescriptorOwner func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		owner **T.SID,
		ownerDefaulted *T.BOOL) (T.BOOL, error)

	GetSecurityDescriptorRMControl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		rmControl *T.WChar) (T.DWORD, error)

	GetSecurityDescriptorSacl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		saclPresent *T.BOOL,
		sacl **T.ACL,
		saclDefaulted *T.BOOL) (T.BOOL, error)

	GetShortPathName func(
		longPath VString,
		shortPath OVString,
		buffer T.DWORD) (T.DWORD, error)

	GetSidIdentifierAuthority func(
		sid *T.SID) (*T.SID_IDENTIFIER_AUTHORITY, error)

	GetSidLengthRequired func(subAuthorityCount T.WChar) (T.DWORD, error)

	GetSidSubAuthority func(
		sid *T.SID, subAuthority T.DWORD) (*T.DWORD, error)

	GetSidSubAuthorityCount func(sid *T.SID) (*T.WChar, error)

	GetStartupInfo func(startupInfo *T.STARTUPINFO)

	GetStdHandle func(stdHandle T.DWORD) (T.HANDLE, error)

	GetSystemDirectory func(buffer VString, size T.UINT) (T.UINT, error)

	GetSystemFileCacheSize func(
		minFileCacheSize, maxFileCacheSize *T.SIZE_T,
		flags *T.DWORD) (T.BOOL, error)

	GetSystemFirmwareTable func(
		firmwareTableProviderSignature,
		firmwareTableID T.DWORD,
		firmwareTableBuffer *T.VOID,
		bufferSize T.DWORD) (T.UINT, error)

	GetSystemInfo func(systemInfo *T.SYSTEM_INFO)

	GetSystemPowerStatus func(
		systemPowerStatus *T.SYSTEM_POWER_STATUS) (T.BOOL, error)

	GetSystemRegistryQuota func(
		quotaAllowed, quotaUsed *T.DWORD) (T.BOOL, error)

	GetSystemTime func(
		systemTime *T.SYSTEMTIME)

	GetSystemTimeAdjustment func(
		timeAdjustment, timeIncrement *T.DWORD,
		timeAdjustmentDisabled *T.BOOL) (T.BOOL, error)

	GetSystemTimeAsFileTime func(
		systemTimeAsFileTime *T.FILETIME)

	GetSystemTimes func(
		idleTime, kernelTime, userTime *T.FILETIME) (T.BOOL, error)

	GetSystemWindowsDirectory func(
		buffer VString, size T.UINT) (T.UINT, error)

	GetSystemWow64Directory func(
		buffer VString, size T.UINT) (T.UINT, error)

	GetTapeParameters func(
		device T.HANDLE,
		operation T.DWORD,
		size *T.DWORD,
		tapeInformation *T.VOID) (T.DWORD, error)

	GetTapePosition func(
		device T.HANDLE,
		positionType T.DWORD,
		partition, offsetLow, offsetHigh *T.DWORD) (T.DWORD, error)

	GetTapeStatus func(device T.HANDLE) (T.DWORD, error)

	GetTempFileName func(
		pathName, prefixString VString,
		unique T.UINT,
		tempFileName VString) (T.UINT, error)

	GetTempPath func(bufferLength T.DWORD, buffer VString) (T.DWORD, error)

	GetThreadContext func(thread T.HANDLE, context *T.CONTEXT) (T.BOOL, error)

	GetThreadId func(thread T.HANDLE) (T.DWORD, error)

	GetThreadIOPendingFlag func(
		thread T.HANDLE, ioIsPending *T.BOOL) (T.BOOL, error)

	GetThreadPriority func(thread T.HANDLE) (int, error)

	GetThreadPriorityBoost func(
		thread T.HANDLE, disablePriorityBoost *T.BOOL) (T.BOOL, error)

	GetThreadSelectorEntry func(
		thread T.HANDLE,
		selector T.DWORD,
		selectorEntry *T.LDT_ENTRY) (T.BOOL, error)

	GetThreadTimes func(
		thread T.HANDLE,
		creationTime, exitTime, kernelTime,
		userTime *T.FILETIME) (T.BOOL, error)

	GetTickCount func() (T.DWORD, error)

	GetTimeZoneInformation func(
		timeZoneInformation *T.TIME_ZONE_INFORMATION) (T.DWORD, error)

	GetTokenInformation func(
		tokenHandle T.HANDLE,
		tokenInformationClass T.TOKEN_INFORMATION_CLASS,
		tokenInformation *T.VOID,
		tokenInformationLength T.DWORD,
		returnLength *T.DWORD) (T.BOOL, error)

	GetUserName func(buffer OVString, size *T.DWORD) (T.BOOL, error)

	GetVersion func() (T.DWORD, error)

	GetVersionExA func(versionInformation *T.OSVERSIONINFOA) (T.BOOL, error)

	GetVersionExW func(versionInformation *T.OSVERSIONINFOW) (T.BOOL, error)

	GetVolumeInformation func(
		rootPathName, volumeNameBuffer VString,
		volumeNameSize T.DWORD,
		volumeSerialNumber, maximumComponentLength,
		fileSystemFlags *T.DWORD,
		fileSystemNameBuffer VString,
		fileSystemNameSize T.DWORD) (T.BOOL, error)

	GetVolumeNameForVolumeMountPoint func(
		volumeMountPoint VString,
		volumeName OVString,
		bufferLength T.DWORD) (T.BOOL, error)

	GetVolumePathName func(
		fileName VString,
		volumePathName OVString,
		bufferLength T.DWORD) (T.BOOL, error)

	GetVolumePathNamesForVolumeName func(
		volumeName VString,
		volumePathNames T.OMVString,
		bufferLength T.DWORD,
		returnLength *T.DWORD) (T.BOOL, error)

	GetWindowsAccountDomainSid func(
		sid, domainSid *T.SID, domainSidl *T.DWORD) (T.BOOL, error)

	GetWindowsDirectory func(buffer OVString, size T.UINT) (T.UINT, error)

	GetWriteWatch func(
		flags T.DWORD,
		baseAddress *T.VOID,
		regionSize T.SIZE_T,
		addresses **T.VOID,
		count *T.ULONG_PTR,
		granularity *T.ULONG) (T.UINT, error)

	GlobalAddAtom func(s VString) (T.ATOM, error)

	GlobalAlloc func(flags T.UINT, bytes T.SIZE_T) (T.HGLOBAL, error)

	GlobalCompact func(minFree T.DWORD) (T.SIZE_T, error)

	GlobalDeleteAtom func(a T.ATOM) (T.ATOM, error)

	GlobalFindAtom func(s VString) (T.ATOM, error)

	GlobalFix func(mem T.HGLOBAL)

	GlobalFlags func(mem T.HGLOBAL) (T.UINT, error)

	GlobalFree func(mem T.HGLOBAL) (T.HGLOBAL, error)

	GlobalGetAtomName func(
		atom T.ATOM, buffer OVString, size int) (T.UINT, error)

	GlobalHandle func(mem *T.VOID) (T.HGLOBAL, error)

	GlobalLock func(mem T.HGLOBAL) (*T.VOID, error)

	GlobalMemoryStatus func(buffer *T.MEMORYSTATUS)

	GlobalMemoryStatusEx func(buffer *T.MEMORYSTATUSEX) (T.BOOL, error)

	GlobalReAlloc func(
		mem T.HGLOBAL, bytes T.SIZE_T, flags T.UINT) (T.HGLOBAL, error)

	GlobalSize func(mem T.HGLOBAL) (T.SIZE_T, error)

	GlobalUnfix func(mem T.HGLOBAL)

	GlobalUnlock func(mem T.HGLOBAL) (T.BOOL, error)

	GlobalUnWire func(mem T.HGLOBAL) (T.BOOL, error)

	GlobalWire func(mem T.HGLOBAL) (*T.VOID, error)

	HeapAlloc func(heap T.HANDLE, flags T.DWORD, bytes T.SIZE_T) (*T.VOID, error)

	HeapCompact func(Heap T.HANDLE, flags T.DWORD) (T.SIZE_T, error)

	HeapCreate func(
		options T.DWORD, initialSize, maximumSize T.SIZE_T) (T.HANDLE, error)

	HeapDestroy func(Heap T.HANDLE) (T.BOOL, error)

	HeapFree func(Heap T.HANDLE, flags T.DWORD, mem *T.VOID) (T.BOOL, error)

	HeapLock func(Heap T.HANDLE) (T.BOOL, error)

	HeapQueryInformation func(
		h T.HANDLE,
		ic T.HEAP_INFORMATION_CLASS,
		info *T.VOID,
		infoLength T.SIZE_T,
		returnLength *T.SIZE_T) (T.BOOL, error)

	HeapReAlloc func(
		heap T.HANDLE, flags T.DWORD, mem *T.VOID, bytes T.SIZE_T) (*T.VOID, error)

	HeapSetInformation func(
		h, T.HANDLE,
		ic T.HEAP_INFORMATION_CLASS,
		info *T.VOID,
		infoLength T.SIZE_T) (T.BOOL, error)

	HeapSize func(heap T.HANDLE, flags T.DWORD, mem *T.VOID) (T.SIZE_T, error)

	HeapUnlock func(Heap T.HANDLE) (T.BOOL, error)

	HeapValidate func(heap T.HANDLE, flags T.DWORD, mem *T.VOID) (T.BOOL, error)

	HeapWalk func(heap T.HANDLE, entry *T.PROCESS_HEAP_ENTRY) (T.BOOL, error)

	Hread func(file T.HFILE, buffer *T.VOID, bytes long) (long, error)

	Hwrite func(file T.HFILE, buffer *T.CH, bytes long) (long, error)

	ImpersonateAnonymousToken func(threadHandle T.HANDLE) (T.BOOL, error)

	ImpersonateLoggedOnUser func(token T.HANDLE) (T.BOOL, error)

	ImpersonateNamedPipeClient func(namedPipe T.HANDLE) (T.BOOL, error)

	ImpersonateSelf func(
		impersonationLevel T.SECURITY_IMPERSONATION_LEVEL) (T.BOOL, error)

	InitAtomTable func(size T.DWORD) (T.BOOL, error)

	InitializeAcl func(
		acl *T.ACL, aclLength T.DWORD, aclRevision T.DWORD) (T.BOOL, error)

	InitializeCriticalSection func(*T.CRITICAL_SECTION)

	InitializeCriticalSectionAndSpinCount func(
		cs *T.CRITICAL_SECTION, spinCount T.DWORD) (T.BOOL, error)

	InitializeSecurityDescriptor func(
		secDesc *T.SECURITY_DESCRIPTOR,
		revision T.DWORD) (T.BOOL, error)

	InitializeSid func(
		sid *T.SID,
		idAuth *T.SID_IDENTIFIER_AUTHORITY,
		subAuthorityCount T.BYTE) (T.BOOL, error)

	InitializeSListHead func(listHead T.SLIST_HEADER)

	InterlockedCompareExchange func(
		destination *T.LONG, exchange, comperand T.LONG) T.LONG

	InterlockedCompareExchange64 func(
		destination *T.LONGLONG,
		exchange, comperand T.LONGLONG) T.LONGLONG

	InterlockedDecrement func(addend *T.LONG) T.LONG

	InterlockedExchange func(target *T.LONG, value T.LONG) T.LONG

	InterlockedExchangeAdd func(addend *T.LONG, value T.LONG) T.LONG

	InterlockedFlushSList func(*T.SLIST_HEADER) *T.SLIST_ENTRY

	InterlockedIncrement func(addend *T.LONG) T.LONG

	InterlockedPopEntrySList func(*T.SLIST_HEADER) *T.SLIST_ENTRY

	InterlockedPushEntrySList func(
		*T.SLIST_HEADER, *T.SLIST_ENTRY) *T.SLIST_ENTRY

	IsBadCodePtr func(T.FARPROC) (T.BOOL, error)

	IsBadHugeReadPtr func(*T.VOID, T.UINT_PTR) (T.BOOL, error)

	IsBadHugeWritePtr func(*T.VOID, T.UINT_PTR) (T.BOOL, error)

	IsBadReadPtr func(*T.VOID, T.UINT_PTR) (T.BOOL, error)

	IsBadStringPtr func(s VString, max T.UINT_PTR) (T.BOOL, error)

	IsBadWritePtr func(*T.VOID, T.UINT_PTR) (T.BOOL, error)

	IsDebuggerPresent func() (T.BOOL, error)

	IsProcessInJob func(
		processHandle T.HANDLE,
		jobHandle T.HANDLE,
		result *T.BOOL) (T.BOOL, error)

	IsProcessorFeaturePresent func(processorFeature T.DWORD) (T.BOOL, error)

	IsSystemResumeAutomatic func() (T.BOOL, error)

	IsTextUnicode func(v *T.VOID, size int, result *T.INT) (T.BOOL, error)

	IsTokenRestricted func(token T.HANDLE) (T.BOOL, error)

	IsTokenUntrusted func(token T.HANDLE) (T.BOOL, error)

	IsValidAcl func(*T.ACL) (T.BOOL, error)

	IsValidSecurityDescriptor func(*T.SECURITY_DESCRIPTOR) (T.BOOL, error)

	IsValidSid func(*T.SID) (T.BOOL, error)

	IsWellKnownSid func(*T.SID, T.WELL_KNOWN_SID_TYPE) (T.BOOL, error)

	IsWow64Process func(process T.HANDLE, wow64Process *T.BOOL) (T.BOOL, error)

	Lclose func(file T.HFILE) (T.HFILE, error)

	Lcreat func(pathName T.AString, attribute int) (T.HFILE, error)

	LeaveCriticalSection func(*T.CRITICAL_SECTION)

	Llseek func(file T.HFILE, offset T.LONG, origin int) (T.LONG, error)

	LoadLibrary func(libFileName VString) (T.HMODULE, error)

	LoadLibraryEx func(
		libFileName VString, file T.HANDLE, flags T.DWORD) (T.HMODULE, error)

	LoadModule func(
		moduleName T.AString, parameterBlock *T.VOID) (T.DWORD, error)

	LoadResource func(
		module T.HMODULE, resInfo T.HRSRC) (T.HGLOBAL, error)

	LocalAlloc func(
		flags T.UINT, bytes T.SIZE_T) (T.HLOCAL, error)

	LocalCompact func(minFree T.UINT) (T.SIZE_T, error)

	LocalFileTimeToFileTime func(
		localFileTime, fileTime *T.FILETIME) (T.BOOL, error)

	LocalFlags func(mem T.HLOCAL) (T.UINT, error)

	LocalFree func(mem T.HLOCAL) (T.HLOCAL, error)

	LocalHandle func(mem *T.VOID) (T.HLOCAL, error)

	LocalLock func(mem T.HLOCAL) (*T.VOID, error)

	LocalReAlloc func(
		mem T.HLOCAL, bytes T.SIZE_T, flags T.UINT) (T.HLOCAL, error)

	LocalShrink func(mem T.HLOCAL, newSize T.UINT) (T.SIZE_T, error)

	LocalSize func(mem T.HLOCAL) (T.SIZE_T, error)

	LocalUnlock func(mem T.HLOCAL) (T.BOOL, error)

	LockFile func(
		file T.HANDLE,
		fileOffsetLow, fileOffsetHigh,
		bytesToLockLow, bytesToLockHigh T.DWORD) (T.BOOL, error)

	LockFileEx func(
		file T.HANDLE,
		flags, reserved,
		bytesToLockLow, bytesToLockHigh T.DWORD,
		overlapped *T.OVERLAPPED) (T.BOOL, error)

	LockResource func(resData T.HGLOBAL) (*T.VOID, error)

	LogonUser func(
		username, domain, password VString,
		logonType, logonProvider T.DWORD,
		token *T.HANDLE) (T.BOOL, error)

	LogonUserEx func(
		username, domain, password VString,
		logonType, logonProvider T.DWORD,
		token *T.HANDLE,
		logonSid **T.SID,
		profileBuffer **T.VOID,
		profileLength *T.DWORD,
		quotaLimits *T.QUOTA_LIMITS) (T.BOOL, error)

	LookupAccountName func(
		systemName VString,
		accountName OVString,
		sid *T.SID,
		nSid *T.DWORD,
		referencedDomainName OVString,
		nRefDomName *T.DWORD,
		use *T.SID_NAME_USE) (T.BOOL, error)

	LookupAccountSid func(
		systemName VString,
		sid *T.SID,
		name OVString,
		nName *T.DWORD,
		referencedDomainName OVString,
		nRefDomName *T.DWORD,
		use *T.SID_NAME_USE) (T.BOOL, error)

	LookupPrivilegeDisplayName func(
		systemName, name VString,
		displayName OVString,
		nDispName, languageId *T.DWORD) (T.BOOL, error)

	LookupPrivilegeName func(
		systemName VString,
		luid *T.LUID,
		name OVString,
		nName *T.DWORD) (T.BOOL, error)

	LookupPrivilegeValue func(
		systemName, name VString, luid *T.LUID) (T.BOOL, error)

	Lopen func(pathName T.AString, readWrite int) (T.HFILE, error)

	Lread func(file T.HFILE, buffer *T.VOID, bytes T.UINT) (T.UINT, error)

	Lstrcat func(string_1, string_2 VString) (VString, error)

	Lstrcmp func(string_1, string_2 VString) (int, error)

	Lstrcmpi func(string_1, string_2 VString) (int, error)

	Lstrcpy func(string_1, string_2 VString) (VString, error)

	Lstrcpyn func(
		string_1, string_2 VString,
		maxLength int) (VString, error)

	Lstrlen func(string_ VString) (int, error)

	Lwrite func(file T.HFILE, buffer T.CH, bytes T.UINT) (T.UINT, error)

	MakeAbsoluteSD func(
		selfRelativeSecurityDescriptor,
		absoluteSecurityDescriptor *T.SECURITY_DESCRIPTOR,
		absoluteSecurityDescriptorSize *T.DWORD,
		dacl *T.ACL, daclSize *T.DWORD,
		sacl *T.ACL, saclSize *T.DWORD,
		owner *T.SID, ownerSize *T.DWORD,
		primaryGroup *T.SID, primaryGroupSize *T.DWORD) (T.BOOL, error)

	MakeAbsoluteSD2 func(
		selfRelativeSecurityDescriptor *T.SECURITY_DESCRIPTOR,
		bufferSize *T.DWORD) (T.BOOL, error)

	MakeSelfRelativeSD func(
		absoluteSecurityDescriptor,
		selfRelativeSecurityDescriptor *T.SECURITY_DESCRIPTOR,
		bufferLength *T.DWORD) (T.BOOL, error)

	MapGenericMask func(
		accessMask *T.DWORD, genericMapping *T.GENERIC_MAPPING)

	MapUserPhysicalPages func(
		virtualAddress *T.VOID,
		numberOfPages T.ULONG_PTR,
		pageArray *T.ULONG_PTR) (T.BOOL, error)

	MapUserPhysicalPagesScatter func(
		virtualAddresses **T.VOID,
		numberOfPages T.ULONG_PTR,
		pageArray *T.ULONG_PTR) (T.BOOL, error)

	MapViewOfFile func(
		fileMappingObject T.HANDLE,
		desiredAccess, fileOffsetHigh, fileOffsetLow T.DWORD,
		numberOfBytesToMap T.SIZE_T) (*T.VOID, error)

	MapViewOfFileEx func(
		fileMappingObject T.HANDLE,
		desiredAccess, fileOffsetHigh, fileOffsetLow T.DWORD,
		numberOfBytesToMap T.SIZE_T,
		baseAddress *T.VOID) (*T.VOID, error)

	MoveFile func(
		existingFileName, newFileName VString) (T.BOOL, error)

	MoveFileEx func(
		existingFileName, newFileName VString,
		flags T.DWORD) (T.BOOL, error)

	MoveFileWithProgress func(
		existingFileName,
		newFileName VString,
		progressRoutine T.PROGRESS_ROUTINE,
		data *T.VOID,
		flags T.DWORD) (T.BOOL, error)

	MulDiv func(number, numerator, denominator int) (int, error)

	NeedCurrentDirectoryForExePath func(
		exeName VString) (T.BOOL, error)

	NotifyChangeEventLog func(
		eventLog T.HANDLE, event T.HANDLE) (T.BOOL, error)

	ObjectCloseAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		generateOnClose T.BOOL) (T.BOOL, error)

	ObjectDeleteAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		generateOnClose T.BOOL) (T.BOOL, error)

	ObjectOpenAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		objectTypeName, objectName VString,
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		clientToken T.HANDLE,
		desiredAccess, grantedAccess T.DWORD,
		privileges *T.PRIVILEGE_SET,
		objectCreation, accessGranted T.BOOL,
		generateOnClose *T.BOOL) (T.BOOL, error)

	ObjectPrivilegeAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		clientToken T.HANDLE,
		desiredAccess T.DWORD,
		privileges *T.PRIVILEGE_SET,
		accessGranted T.BOOL) (T.BOOL, error)

	OpenBackupEventLog func(
		uncServerName, fileName VString) (T.HANDLE, error)

	OpenEncryptedFileRaw func(
		fileName VString, flags T.ULONG, context **T.VOID) (T.DWORD, error)

	OpenEvent func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		name VString) (T.HANDLE, error)

	OpenEventLog func(
		uncServerName, sourceName VString) (T.HANDLE, error)

	OpenFile func(
		fileName T.AString,
		reOpenBuff *T.OFSTRUCT,
		style T.UINT) (T.HFILE, error)

	OpenFileMapping func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		name VString) (T.HANDLE, error)

	OpenJobObject func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		name VString) (T.HANDLE, error)

	OpenMutex func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		name VString) (T.HANDLE, error)

	OpenProcess func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		processId T.DWORD) (T.HANDLE, error)

	OpenProcessToken func(
		processHandle T.HANDLE,
		desiredAccess T.DWORD,
		tokenHandle *T.HANDLE) (T.BOOL, error)

	OpenSemaphore func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		name VString) (T.HANDLE, error)

	OpenThread func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		threadId T.DWORD) (T.HANDLE, error)

	OpenThreadToken func(
		threadHandle T.HANDLE,
		desiredAccess T.DWORD,
		openAsSelf T.BOOL,
		tokenHandle *T.HANDLE) (T.BOOL, error)

	OpenWaitableTimer func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		timerName VString) (T.HANDLE, error)

	OutputDebugString func(
		outputString VString)

	PeekNamedPipe func(
		namedPipe T.HANDLE,
		buffer *T.VOID,
		bufferSize T.DWORD,
		bytesRead, totalBytesAvail *T.DWORD,
		bytesLeftThisMessage *T.DWORD) (T.BOOL, error)

	PostQueuedCompletionStatus func(
		completionPort T.HANDLE,
		numberOfBytesTransferred T.DWORD,
		completionKey T.ULONG_PTR,
		overlapped *T.OVERLAPPED) (T.BOOL, error)

	PrepareTape func(
		device T.HANDLE,
		operation T.DWORD,
		immediate T.BOOL) (T.DWORD, error)

	PrivilegeCheck func(
		clientToken T.HANDLE,
		requiredPrivileges *T.PRIVILEGE_SET,
		result *T.BOOL) (T.BOOL, error)

	PrivilegedServiceAuditAlarm func(
		subsystemName, serviceName VString,
		clientToken T.HANDLE,
		privileges *T.PRIVILEGE_SET,
		accessGranted T.BOOL) (T.BOOL, error)

	ProcessIdToSessionId func(
		processId T.DWORD,
		sessionId *T.DWORD) (T.BOOL, error)

	PulseEvent func(event T.HANDLE) (T.BOOL, error)

	PurgeComm func(file T.HANDLE, flags T.DWORD) (T.BOOL, error)

	QueryActCtxW func(
		flags T.DWORD,
		actCtx T.HANDLE,
		subInstance *T.VOID,
		infoClass T.ULONG,
		buffer *T.VOID,
		size T.SIZE_T,
		writtenOrRequired *T.SIZE_T) (T.BOOL, error)

	QueryDepthSList func(listHead *T.SLIST_HEADER) (T.USHORT, error)

	QueryDosDevice func(
		deviceName VString,
		targetPath OVString,
		max T.DWORD) (T.DWORD, error)

	QueryInformationJobObject func(
		job T.HANDLE,
		class T.JOBOBJECTINFOCLASS,
		info *T.VOID,
		length T.DWORD,
		returnLength *T.DWORD) (T.BOOL, error)

	QueryMemoryResourceNotification func(
		notificationHandle T.HANDLE, state *T.BOOL) (T.BOOL, error)

	QueryPerformanceCounter func(
		performanceCount *T.LARGE_INTEGER) (T.BOOL, error)

	QueryPerformanceFrequency func(
		frequency *T.LARGE_INTEGER) (T.BOOL, error)

	QueueUserAPC func(
		apc T.APCFUNC, thread T.HANDLE, data T.ULONG_PTR) (T.DWORD, error)

	QueueUserWorkItem func(
		function T.THREAD_START_ROUTINE,
		context *T.VOID,
		flags T.ULONG) (T.BOOL, error)

	RaiseException func(
		exceptionCode,
		exceptionFlags,
		numberOfArguments T.DWORD,
		arguments *T.ULONG_PTR)

	ReadDirectoryChangesW func(
		directory T.HANDLE,
		buffer *T.VOID,
		bufferLength T.DWORD,
		watchSubtree T.BOOL,
		notifyFilter T.DWORD,
		bytesReturned *T.DWORD,
		overlapped *T.OVERLAPPED,
		completionRoutine T.OVERLAPPED_COMPLETION_ROUTINE) (T.BOOL, error)

	ReadEncryptedFileRaw func(
		exportCallback T.FE_EXPORT_FUNC,
		callbackContext,
		context *T.VOID) (T.DWORD, error)

	ReadEventLog func(
		eventLog T.HANDLE,
		readFlags, recordOffset T.DWORD,
		buffer *T.VOID,
		numberOfBytesToRead T.DWORD,
		bytesRead *T.DWORD,
		minNumberOfBytesNeeded *T.DWORD) (T.BOOL, error)

	ReadFile func(
		file T.HANDLE,
		buffer *T.VOID,
		numberOfBytesToRead T.DWORD,
		numberOfBytesRead *T.DWORD,
		overlapped *T.OVERLAPPED) (T.BOOL, error)

	ReadFileEx func(
		file T.HANDLE,
		buffer *T.VOID,
		numberOfBytesToRead T.DWORD,
		overlapped *T.OVERLAPPED,
		completionRoutine T.OVERLAPPED_COMPLETION_ROUTINE) (T.BOOL, error)

	ReadFileScatter func(
		file T.HANDLE,
		segmentArray *T.FILE_SEGMENT_ELEMENT, // []
		numberOfBytesToRead T.DWORD,
		reserved *T.DWORD,
		overlapped *T.OVERLAPPED) (T.BOOL, error)

	ReadProcessMemory func(
		process T.HANDLE,
		baseAddress, buffer *T.VOID,
		size T.SIZE_T,
		numberOfBytesRead *T.SIZE_T) (T.BOOL, error)

	RegisterEventSource func(
		uncServerName, sourceName VString) (T.HANDLE, error)

	RegisterWaitForSingleObject func(
		newWaitObject *T.HANDLE,
		object T.HANDLE,
		callback T.WAITORTIMERCALLBACK,
		context *T.VOID,
		milliseconds, flags T.ULONG) (T.BOOL, error)

	RegisterWaitForSingleObjectEx func(
		object T.HANDLE,
		callback T.WAITORTIMERCALLBACK,
		context *T.VOID,
		milliseconds, flags T.ULONG) (T.HANDLE, error)

	ReleaseActCtx func(actCtx T.HANDLE)

	ReleaseMutex func(mutex T.HANDLE) (T.BOOL, error)

	ReleaseSemaphore func(
		semaphore T.HANDLE,
		releaseCount T.LONG,
		previousCount *T.LONG) (T.BOOL, error)

	RemoveDirectory func(pathName VString) (T.BOOL, error)

	RemoveVectoredContinueHandler func(Handle *T.VOID) (T.ULONG, error)

	RemoveVectoredExceptionHandler func(Handle *T.VOID) (T.ULONG, error)

	ReOpenFile func(
		originalFile T.HANDLE,
		desiredAccess, shareMode,
		flagsAndAttributes T.DWORD) (T.HANDLE, error)

	ReplaceFile func(
		replacedFileName,
		replacementFileName,
		backupFileName VString,
		replaceFlags T.DWORD,
		exclude, reserved *T.VOID) (T.BOOL, error)

	ReportEvent func(
		eventLog T.HANDLE,
		type_, category T.WORD,
		eventID T.DWORD,
		userSid *T.SID,
		numStrings T.WORD,
		dataSize T.DWORD,
		string_s T.OMVString,
		rawData *T.VOID) (T.BOOL, error)

	RequestDeviceWakeup func(device T.HANDLE) (T.BOOL, error)

	RequestWakeupLatency func(latency T.LATENCY_TIME) (T.BOOL, error)

	ResetEvent func(event T.HANDLE) (T.BOOL, error)

	ResetWriteWatch func(
		baseAddress *T.VOID, regionSize T.SIZE_T) (T.UINT, error)

	RestoreLastError func(errCode T.DWORD)

	ResumeThread func(thread T.HANDLE) (T.DWORD, error)

	RevertToSelf func() (T.BOOL, error)

	SearchPath func(
		path, fileName, extension VString,
		bufferLength T.DWORD,
		buffer OVString,
		filePart *VString) (T.DWORD, error)

	SetAclInformation func(
		acl *T.ACL,
		aclInformation *T.VOID,
		aclInformationLength T.DWORD,
		aclInformationClass T.ACL_INFORMATION_CLASS) (T.BOOL, error)

	SetCommBreak func(file T.HANDLE) (T.BOOL, error)

	SetCommConfig func(
		commDev T.HANDLE, cc *T.COMMCONFIG, size T.DWORD) (T.BOOL, error)

	SetCommMask func(file T.HANDLE, evtMask T.DWORD) (T.BOOL, error)

	SetCommState func(file T.HANDLE, dcb *T.DCB) (T.BOOL, error)

	SetCommTimeouts func(
		file T.HANDLE, commTimeouts *T.COMMTIMEOUTS) (T.BOOL, error)

	SetComputerName func(computerName VString) (T.BOOL, error)

	SetComputerNameEx func(
		nameType T.COMPUTER_NAME_FORMAT,
		buffer VString) (T.BOOL, error)

	SetCriticalSectionSpinCount func(
		criticalSection *T.CRITICAL_SECTION,
		spinCount T.DWORD) (T.DWORD, error)

	SetCurrentDirectory func(pathName VString) (T.BOOL, error)

	SetDefaultCommConfig func(
		name VString, cc *T.COMMCONFIG, size T.DWORD) (T.BOOL, error)

	SetDllDirectory func(pathName VString) (T.BOOL, error)

	SetEndOfFile func(file T.HANDLE) (T.BOOL, error)

	SetEnvironmentStrings func(newEnvironment T.MVString) (T.BOOL, error)

	SetEnvironmentVariable func(name, value VString) (T.BOOL, error)

	SetErrorMode func(mode T.UINT) (T.UINT, error)

	SetEvent func(event T.HANDLE) (T.BOOL, error)

	SetFileApisToANSI func()

	SetFileApisToOEM func()

	SetFileAttributes func(name VString, attributes T.DWORD) (T.BOOL, error)

	SetFilePointer func(
		file T.HANDLE,
		distanceToMove T.LONG,
		distanceToMoveHigh *T.LONG,
		moveMethod T.DWORD) (T.DWORD, error)

	SetFilePointerEx func(
		file T.HANDLE,
		distanceToMove T.LARGE_INTEGER,
		newFilePointer *T.LARGE_INTEGER,
		moveMethod T.DWORD) (T.BOOL, error)

	SetFileSecurity func(
		fileName VString,
		info T.SECURITY_INFORMATION,
		descriptor *T.SECURITY_DESCRIPTOR) (T.BOOL, error)

	SetFileShortName func(
		file T.HANDLE, shortName VString) (T.BOOL, error)

	SetFileTime func(
		file T.HANDLE,
		creationTime, lastAccessTime,
		lastWriteTime *T.FILETIME) (T.BOOL, error)

	SetFileValidData func(
		file T.HANDLE, validDataLength T.LONGLONG) (T.BOOL, error)

	SetFirmwareEnvironmentVariable func(
		name, guid VString, value *T.VOID, size T.DWORD) (T.BOOL, error)

	SetHandleCount func(number T.UINT) (T.UINT, error)

	SetHandleInformation func(
		object T.HANDLE, mask T.DWORD, flags T.DWORD) (T.BOOL, error)

	SetInformationJobObject func(
		job T.HANDLE,
		jobObjectInformationClass T.JOBOBJECTINFOCLASS,
		jobObjectInformation *T.VOID,
		jobObjectInformationLength T.DWORD) (T.BOOL, error)

	SetKernelObjectSecurity func(
		handle T.HANDLE,
		securityInformation T.SECURITY_INFORMATION,
		securityDescriptor *T.SECURITY_DESCRIPTOR) (T.BOOL, error)

	SetLastError func(errCode T.DWORD)

	SetLocalTime func(systemTime *T.SYSTEMTIME) (T.BOOL, error)

	SetMailslotInfo func(
		mailslot T.HANDLE,
		readTimeout T.DWORD) (T.BOOL, error)

	SetMessageWaitingIndicator func(
		msgIndicator T.HANDLE,
		msgCount T.ULONG) (T.BOOL, error)

	SetNamedPipeHandleState func(
		namedPipe T.HANDLE,
		mode, maxCollectionCount,
		collectDataTimeout *T.DWORD) (T.BOOL, error)

	SetPriorityClass func(
		process T.HANDLE, priorityClass T.DWORD) (T.BOOL, error)

	SetPrivateObjectSecurity func(
		securityInformation T.SECURITY_INFORMATION,
		modificationDescriptor *T.SECURITY_DESCRIPTOR,
		objectsSecurityDescriptor **T.SECURITY_DESCRIPTOR,
		genericMapping *T.GENERIC_MAPPING,
		token T.HANDLE) (T.BOOL, error)

	SetPrivateObjectSecurityEx func(
		securityInformation T.SECURITY_INFORMATION,
		modificationDescriptor *T.SECURITY_DESCRIPTOR,
		objectsSecurityDescriptor **T.SECURITY_DESCRIPTOR,
		autoInheritFlags T.ULONG,
		genericMapping *T.GENERIC_MAPPING,
		token T.HANDLE) (T.BOOL, error)

	SetProcessAffinityMask func(
		process T.HANDLE, processAffinityMask T.DWORD_PTR) (T.BOOL, error)

	SetProcessPriorityBoost func(
		process T.HANDLE, disablePriorityBoost T.BOOL) (T.BOOL, error)

	SetProcessShutdownParameters func(
		level, flags T.DWORD) (T.BOOL, error)

	SetProcessWorkingSetSize func(
		process T.HANDLE,
		minWorkingSetSize, maxWorkingSetSize T.SIZE_T) (T.BOOL, error)

	SetProcessWorkingSetSizeEx func(
		process T.HANDLE,
		minWorkingSetSize, maxWorkingSetSize T.SIZE_T,
		flags T.DWORD) (T.BOOL, error)

	SetSecurityDescriptorControl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		controlBitsOfInterest,
		controlBitsToSet T.SECURITY_DESCRIPTOR_CONTROL) (T.BOOL, error)

	SetSecurityDescriptorDacl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		daclPresent T.BOOL,
		dacl *T.ACL,
		daclDefaulted T.BOOL) (T.BOOL, error)

	SetSecurityDescriptorGroup func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		group *T.SID,
		groupDefaulted T.BOOL) (T.BOOL, error)

	SetSecurityDescriptorOwner func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		owner *T.SID,
		ownerDefaulted T.BOOL) (T.BOOL, error)

	SetSecurityDescriptorRMControl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		rmControl *T.WChar) (T.DWORD, error)

	SetSecurityDescriptorSacl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		saclPresent T.BOOL,
		sacl *T.ACL,
		saclDefaulted T.BOOL) (T.BOOL, error)

	SetStdHandle func(stdHandle T.DWORD, handle T.HANDLE) (T.BOOL, error)

	SetSystemFileCacheSize func(
		minFileCacheSize, maxFileCacheSize T.SIZE_T,
		flags T.DWORD) (T.BOOL, error)

	SetSystemPowerState func(suspend, force T.BOOL) (T.BOOL, error)

	SetSystemTime func(systemTime *T.SYSTEMTIME) (T.BOOL, error)

	SetSystemTimeAdjustment func(
		timeAdjustment T.DWORD,
		timeAdjustmentDisabled T.BOOL) (T.BOOL, error)

	SetTapeParameters func(
		device T.HANDLE,
		operation T.DWORD,
		tapeInformation *T.VOID) (T.DWORD, error)

	SetTapePosition func(
		device T.HANDLE,
		positionMethod,
		partition,
		offsetLow,
		offsetHigh T.DWORD,
		immediate T.BOOL) (T.DWORD, error)

	SetThreadAffinityMask func(
		thread T.HANDLE,
		affinityMask T.DWORD_PTR) (T.DWORD_PTR, error)

	SetThreadContext func(
		thread T.HANDLE, context *T.CONTEXT) (T.BOOL, error)

	SetThreadExecutionState func(
		flags T.EXECUTION_STATE) (T.EXECUTION_STATE, error)

	SetThreadIdealProcessor func(
		thread T.HANDLE, idealProcessor T.DWORD) (T.DWORD, error)

	SetThreadPriority func(
		thread T.HANDLE, priority int) (T.BOOL, error)

	SetThreadPriorityBoost func(
		thread T.HANDLE, disablePriorityBoost T.BOOL) (T.BOOL, error)

	SetThreadStackGuarantee func(
		stackSizeInBytes *T.ULONG) (T.BOOL, error)

	SetThreadToken func(
		thread *T.HANDLE, token T.HANDLE) (T.BOOL, error)

	SetTimerQueueTimer func(
		timerQueue T.HANDLE,
		callback T.WAITORTIMERCALLBACK,
		parameter *T.VOID,
		dueTime, period T.DWORD,
		preferIo T.BOOL) (T.HANDLE, error)

	SetTimeZoneInformation func(
		timeZoneInformation *T.TIME_ZONE_INFORMATION) (T.BOOL, error)

	SetTokenInformation func(
		tokenHandle T.HANDLE,
		tokenInformationClass T.TOKEN_INFORMATION_CLASS,
		tokenInformation *T.VOID,
		tokenInformationLength T.DWORD) (T.BOOL, error)

	SetUnhandledExceptionFilter func(
		tef *T.TOP_LEVEL_EXCEPTION_FILTER) (*T.TOP_LEVEL_EXCEPTION_FILTER, error)

	SetupComm func(
		file T.HANDLE, inQueue, outQueue T.DWORD) (T.BOOL, error)

	SetVolumeLabel func(
		rootPathName, volumeName VString) (T.BOOL, error)

	SetVolumeMountPoint func(
		volumeMountPoint, volumeName VString) (T.BOOL, error)

	SetWaitableTimer func(
		timer T.HANDLE,
		dueTime *T.LARGE_INTEGER,
		period T.LONG,
		completionRoutine T.TIMERAPCROUTINE,
		argToCompletionRoutine *T.VOID,
		resume T.BOOL) (T.BOOL, error)

	SignalObjectAndWait func(
		objectToSignal, objectToWaitOn T.HANDLE,
		milliseconds T.DWORD,
		alertable T.BOOL) (T.DWORD, error)

	SizeofResource func(module T.HMODULE, resInfo T.HRSRC) (T.DWORD, error)

	Sleep func(milliseconds T.DWORD)

	SleepEx func(milliseconds T.DWORD, alertable T.BOOL) (T.DWORD, error)

	SuspendThread func(thread T.HANDLE) (T.DWORD, error)

	SwitchToFiber func(fiber *T.VOID)

	SwitchToThread func() (T.BOOL, error)

	SystemTimeToFileTime func(
		systemTime *T.SYSTEMTIME, fileTime *T.FILETIME) (T.BOOL, error)

	SystemTimeToTzSpecificLocalTime func(
		timeZoneInformation *T.TIME_ZONE_INFORMATION,
		universalTime, localTime *T.SYSTEMTIME) (T.BOOL, error)

	TerminateJobObject func(job T.HANDLE, exitCode T.UINT) (T.BOOL, error)

	TerminateProcess func(process T.HANDLE, exitCode T.UINT) (T.BOOL, error)

	TerminateThread func(thread T.HANDLE, exitCode T.DWORD) (T.BOOL, error)

	TlsAlloc func() (T.DWORD, error)

	TlsFree func(tlsIndex T.DWORD) (T.BOOL, error)

	TlsGetValue func(tlsIndex T.DWORD) (*T.VOID, error)

	TlsSetValue func(tlsIndex T.DWORD, tlsValue *T.VOID) (T.BOOL, error)

	TransactNamedPipe func(
		namedPipe T.HANDLE,
		inBuffer *T.VOID, inBufferSize T.DWORD,
		outBuffer *T.VOID, outBufferSize T.DWORD,
		bytesRead *T.DWORD,
		overlapped *T.OVERLAPPED) (T.BOOL, error)

	TransmitCommChar func(file T.HANDLE, c char) (T.BOOL, error)

	TryEnterCriticalSection func(
		criticalSection *T.CRITICAL_SECTION) (T.BOOL, error)

	TzSpecificLocalTimeToSystemTime func(
		timeZoneInformation *T.TIME_ZONE_INFORMATION,
		local, universal *T.SYSTEMTIME) (T.BOOL, error)

	UnlockFile func(
		file T.HANDLE,
		offsetLow, offsetHigh, bytesLow, bytesHigh T.DWORD) (T.BOOL, error)

	UnlockFileEx func(
		file T.HANDLE,
		reserved, bytesLow, bytesHigh T.DWORD,
		o *T.OVERLAPPED) (T.BOOL, error)

	UnmapViewOfFile func(baseAddress *T.VOID) (T.BOOL, error)

	UnregisterWait func(waitHandle T.HANDLE) (T.BOOL, error)

	UnregisterWaitEx func(
		waitHandle, completionEvent T.HANDLE) (T.BOOL, error)

	UpdateResource func(
		update T.HANDLE,
		typ, name VString,
		language T.WORD,
		data *T.VOID,
		count T.DWORD) (T.BOOL, error)

	VerifyVersionInfoA func(
		versionInformation *T.OSVERSIONINFOEXA,
		typeMask T.DWORD,
		conditionMask T.DWORDLONG) (T.BOOL, error)

	VerifyVersionInfoW func(
		versionInformation *T.OSVERSIONINFOEXW,
		typeMask T.DWORD,
		conditionMask T.DWORDLONG) (T.BOOL, error)

	VirtualAlloc func(
		address *T.VOID,
		size T.SIZE_T,
		allocationType, protect T.DWORD) (*T.VOID, error)

	VirtualAllocEx func(
		process T.HANDLE,
		address *T.VOID,
		size T.SIZE_T,
		allocationType, protect T.DWORD) (*T.VOID, error)

	VirtualFree func(
		address *T.VOID, size T.SIZE_T, freeType T.DWORD) (T.BOOL, error)

	VirtualFreeEx func(
		process T.HANDLE,
		address *T.VOID,
		size T.SIZE_T,
		freeType T.DWORD) (T.BOOL, error)

	VirtualLock func(address *T.VOID, size T.SIZE_T) (T.BOOL, error)

	VirtualProtect func(
		address *T.VOID,
		size T.SIZE_T,
		newProtect T.DWORD,
		oldProtect *T.DWORD) (T.BOOL, error)

	VirtualProtectEx func(
		process T.HANDLE,
		address *T.VOID,
		size T.SIZE_T,
		newProtect T.DWORD,
		oldProtect *T.DWORD) (T.BOOL, error)

	VirtualQuery func(
		address *T.VOID,
		buffer *T.MEMORY_BASIC_INFORMATION,
		length T.SIZE_T) (T.SIZE_T, error)

	VirtualQueryEx func(
		process T.HANDLE,
		address *T.VOID,
		buffer *T.MEMORY_BASIC_INFORMATION,
		length T.SIZE_T) (T.SIZE_T, error)

	VirtualUnlock func(address *T.VOID, size T.SIZE_T) (T.BOOL, error)

	WaitCommEvent func(
		file T.HANDLE,
		evtMask *T.DWORD,
		overlapped *T.OVERLAPPED) (T.BOOL, error)

	WaitForDebugEvent func(
		debugEvent *T.DEBUG_EVENT,
		milliseconds T.DWORD) (T.BOOL, error)

	WaitForMultipleObjects func(
		count T.DWORD,
		handles *T.HANDLE,
		waitAll T.BOOL,
		milliseconds T.DWORD) (T.DWORD, error)

	WaitForMultipleObjectsEx func(
		count T.DWORD,
		handles *T.HANDLE,
		waitAll T.BOOL,
		milliseconds T.DWORD,
		alertable T.BOOL) (T.DWORD, error)

	WaitForSingleObject func(
		handle T.HANDLE, milliseconds T.DWORD) (T.DWORD, error)

	WaitForSingleObjectEx func(
		handle T.HANDLE,
		milliseconds T.DWORD,
		alertable T.BOOL) (T.DWORD, error)

	WaitNamedPipe func(
		namedPipeName VString,
		timeOut T.DWORD) (T.BOOL, error)

	//WinExec is obsolete; instead use:
	//	CreateProcess
	WinExec func(cmdLine string, cmdShow T.UINT) (T.UINT, error)

	Wow64DisableWow64FsRedirection func(
		oldValue **T.VOID) (T.BOOL, error)

	Wow64EnableWow64FsRedirection func(
		wow64FsEnableRedirection T.BOOLEAN) (T.BOOLEAN, error)

	Wow64RevertWow64FsRedirection func(olValue *T.VOID) (T.BOOL, error)

	WriteEncryptedFileRaw func(
		importCallback T.FE_IMPORT_FUNC,
		callbackContext, context *T.VOID) (T.DWORD, error)

	WriteFile func(
		file T.HANDLE,
		buffer *T.VOID,
		numberOfBytesToWrite T.DWORD,
		numberOfBytesWritten *T.DWORD,
		overlapped *T.OVERLAPPED) (T.BOOL, error)

	WriteFileEx func(
		file T.HANDLE,
		buffer *T.VOID,
		numberOfBytesToWrite T.DWORD,
		overlapped *T.OVERLAPPED,
		completionRoutine T.OVERLAPPED_COMPLETION_ROUTINE) (T.BOOL, error)

	WriteFileGather func(
		file T.HANDLE,
		segmentArray *T.FILE_SEGMENT_ELEMENT, // []
		numberOfBytesToWrite T.DWORD,
		reserved *T.DWORD,
		overlapped *T.OVERLAPPED) (T.BOOL, error)

	//WritePrivateProfileSection is obsolete; instead use:
	//	registry
	WritePrivateProfileSection func(
		appName, string_, fileName VString) (T.BOOL, error)

	//WritePrivateProfileString is obsolete; instead use:
	//	registry
	WritePrivateProfileString func(
		appName, keyName, string_, fileName VString) (T.BOOL, error)

	//WritePrivateProfileStruct is obsolete; instead use:
	//	registry
	WritePrivateProfileStruct func(
		section, key VString,
		struct_ *T.VOID,
		sizeStruct T.UINT,
		file VString) (T.BOOL, error)

	WriteProcessMemory func(
		process T.HANDLE,
		baseAddress, buffer *T.VOID,
		size T.SIZE_T,
		numberOfBytesWritten *T.SIZE_T) (T.BOOL, error)

	//WriteProfileSection is obsolete; instead use:
	//	registry
	WriteProfileSection func(
		appName, string_ VString) (T.BOOL, error)

	//WriteProfileString is obsolete; instead use:
	//	registry
	WriteProfileString func(
		appName, keyName, string_ VString) (T.BOOL, error)

	WriteTapemark func(
		device T.HANDLE,
		tapemarkType, tapemarkCount T.DWORD,
		immediate T.BOOL) (T.DWORD, error)

	WTSGetActiveConsoleSessionId func() (T.DWORD, error)

	ZombifyActCtx func(actCtx T.HANDLE) (T.BOOL, error)

/* TODO(t):AFTER INIT
InterlockedIncrementAcquire = InterlockedIncrement
InterlockedIncrementRelease = InterlockedIncrement
InterlockedDecrementAcquire = InterlockedDecrement
InterlockedDecrementRelease = InterlockedDecrement
InterlockedCompareExchangeAcquire = InterlockedCompareExchange
InterlockedCompareExchangeRelease = InterlockedCompareExchange
InterlockedCompareExchangeAcquire64 = InterlockedCompareExchange64
InterlockedCompareExchangeRelease64 = InterlockedCompareExchange64
InterlockedCompareExchangePointerAcquire = InterlockedCompareExchangePointer
InterlockedCompareExchangePointerRelease = InterlockedCompareExchangePointer
*/
)

// NOTE(t): Watch out for void arguments and returns

//NOTE(t): exists but is not user function (builtin default)
//{"UnhandledExceptionFilter", &UnhandledExceptionFilter},
//TODO(t): Not xp32
//{"AddVectoredContinueHandler", &AddVectoredContinueHandler},
//{"ConvertThreadToFiberEx", &ConvertThreadToFiberEx},
//{"EndUpdateResource", &EndUpdateResource},
//{"EnumSystemFirmwareTables", &EnumSystemFirmwareTables},
//{"FindFirstStreamW", &FindFirstStreamW},
//{"FindNextStreamW", &FindNextStreamW},
//{"FlsAlloc", &FlsAlloc},
//{"FlsFree", &FlsFree},
//{"FlsGetValue", &FlsGetValue},
//{"FlsSetValue", &FlsSetValue},
//{"GetCurrentProcessorNumber", &GetCurrentProcessorNumber},
//{"GetLargePageMinimum", &GetLargePageMinimum},
//{"GetProcessIdOfThread", &GetProcessIdOfThread},
//{"GetProcessWorkingSetSizeEx", &GetProcessWorkingSetSizeEx},
//{"GetSystemFileCacheSize", &GetSystemFileCacheSize},
//{"GetSystemFirmwareTable", &GetSystemFirmwareTable},
//{"GetThreadId", &GetThreadId},
//{"InterlockedCompareExchange64", &InterlockedCompareExchange64},
//{"RemoveVectoredContinueHandler", &RemoveVectoredContinueHandler},
//{"SetProcessWorkingSetSizeEx", &SetProcessWorkingSetSizeEx},
//{"SetSystemFileCacheSize", &SetSystemFileCacheSize},
//{"SetThreadStackGuarantee", &SetThreadStackGuarantee},
//{"Wow64DisableWow64FsRedirection", &Wow64DisableWow64FsRedirection},
//{"Wow64EnableWow64FsRedirection", &Wow64EnableWow64FsRedirection},
//{"Wow64RevertWow64FsRedirection", &Wow64RevertWow64FsRedirection},

var WinBaseApis = outside.Apis{
	{"_hread", &Hread},
	{"_hwrite", &Hwrite},
	{"_lclose", &Lclose},
	{"_lcreat", &Lcreat},
	{"_llseek", &Llseek},
	{"_lopen", &Lopen},
	{"_lread", &Lread},
	{"_lwrite", &Lwrite},
	{"AccessCheck", &AccessCheck},
	{"AccessCheckByType", &AccessCheckByType},
	{"AccessCheckByTypeResultList", &AccessCheckByTypeResultList},
	{"ActivateActCtx", &ActivateActCtx},
	{"AddAccessAllowedAce", &AddAccessAllowedAce},
	{"AddAccessAllowedAceEx", &AddAccessAllowedAceEx},
	{"AddAccessAllowedObjectAce", &AddAccessAllowedObjectAce},
	{"AddAccessDeniedAce", &AddAccessDeniedAce},
	{"AddAccessDeniedAceEx", &AddAccessDeniedAceEx},
	{"AddAccessDeniedObjectAce", &AddAccessDeniedObjectAce},
	{"AddAce", &AddAce},
	{"AddAuditAccessAce", &AddAuditAccessAce},
	{"AddAuditAccessAceEx", &AddAuditAccessAceEx},
	{"AddAuditAccessObjectAce", &AddAuditAccessObjectAce},
	{"AddRefActCtx", &AddRefActCtx},
	{"AddVectoredExceptionHandler", &AddVectoredExceptionHandler},
	{"AdjustTokenGroups", &AdjustTokenGroups},
	{"AdjustTokenPrivileges", &AdjustTokenPrivileges},
	{"AllocateAndInitializeSid", &AllocateAndInitializeSid},
	{"AllocateLocallyUniqueId", &AllocateLocallyUniqueId},
	{"AllocateUserPhysicalPages", &AllocateUserPhysicalPages},
	{"AreAllAccessesGranted", &AreAllAccessesGranted},
	{"AreAnyAccessesGranted", &AreAnyAccessesGranted},
	{"AreFileApisANSI", &AreFileApisANSI},
	{"AssignProcessToJobObject", &AssignProcessToJobObject},
	{"BackupRead", &BackupRead},
	{"BackupSeek", &BackupSeek},
	{"BackupWrite", &BackupWrite},
	{"Beep", &Beep},
	{"BindIoCompletionCallback", &BindIoCompletionCallback},
	{"CancelDeviceWakeupRequest", &CancelDeviceWakeupRequest},
	{"CancelIo", &CancelIo},
	{"CancelTimerQueueTimer", &CancelTimerQueueTimer},
	{"CancelWaitableTimer", &CancelWaitableTimer},
	{"ChangeTimerQueueTimer", &ChangeTimerQueueTimer},
	{"CheckRemoteDebuggerPresent", &CheckRemoteDebuggerPresent},
	{"ClearCommBreak", &ClearCommBreak},
	{"ClearCommError", &ClearCommError},
	{"CloseEncryptedFileRaw", &CloseEncryptedFileRaw},
	{"CloseEventLog", &CloseEventLog},
	{"CloseHandle", &CloseHandle},
	{"CompareFileTime", &CompareFileTime},
	{"ConnectNamedPipe", &ConnectNamedPipe},
	{"ContinueDebugEvent", &ContinueDebugEvent},
	{"ConvertFiberToThread", &ConvertFiberToThread},
	{"ConvertThreadToFiber", &ConvertThreadToFiber},
	{"ConvertToAutoInheritPrivateObjectSecurity", &ConvertToAutoInheritPrivateObjectSecurity},
	{"CopySid", &CopySid},
	{"CreateFiber", &CreateFiber},
	{"CreateFiberEx", &CreateFiberEx},
	{"CreateIoCompletionPort", &CreateIoCompletionPort},
	{"CreateJobSet", &CreateJobSet},
	{"CreateMemoryResourceNotification", &CreateMemoryResourceNotification},
	{"CreatePipe", &CreatePipe},
	{"CreatePrivateObjectSecurity", &CreatePrivateObjectSecurity},
	{"CreatePrivateObjectSecurityEx", &CreatePrivateObjectSecurityEx},
	{"CreatePrivateObjectSecurityWithMultipleInheritance", &CreatePrivateObjectSecurityWithMultipleInheritance},
	{"CreateRemoteThread", &CreateRemoteThread},
	{"CreateRestrictedToken", &CreateRestrictedToken},
	{"CreateTapePartition", &CreateTapePartition},
	{"CreateThread", &CreateThread},
	{"CreateTimerQueue", &CreateTimerQueue},
	{"CreateTimerQueueTimer", &CreateTimerQueueTimer},
	{"CreateWellKnownSid", &CreateWellKnownSid},
	{"DeactivateActCtx", &DeactivateActCtx},
	{"DebugActiveProcess", &DebugActiveProcess},
	{"DebugActiveProcessStop", &DebugActiveProcessStop},
	{"DebugBreak", &DebugBreak},
	{"DebugBreakProcess", &DebugBreakProcess},
	{"DebugSetProcessKillOnExit", &DebugSetProcessKillOnExit},
	{"DecodePointer", &DecodePointer},
	{"DecodeSystemPointer", &DecodeSystemPointer},
	{"DeleteAce", &DeleteAce},
	{"DeleteAtom", &DeleteAtom},
	{"DeleteCriticalSection", &DeleteCriticalSection},
	{"DeleteFiber", &DeleteFiber},
	{"DeleteTimerQueue", &DeleteTimerQueue},
	{"DeleteTimerQueueEx", &DeleteTimerQueueEx},
	{"DeleteTimerQueueTimer", &DeleteTimerQueueTimer},
	{"DeregisterEventSource", &DeregisterEventSource},
	{"DestroyPrivateObjectSecurity", &DestroyPrivateObjectSecurity},
	{"DeviceIoControl", &DeviceIoControl},
	{"DisableThreadLibraryCalls", &DisableThreadLibraryCalls},
	{"DisconnectNamedPipe", &DisconnectNamedPipe},
	{"DosDateTimeToFileTime", &DosDateTimeToFileTime},
	{"DuplicateHandle", &DuplicateHandle},
	{"DuplicateToken", &DuplicateToken},
	{"DuplicateTokenEx", &DuplicateTokenEx},
	{"EncodePointer", &EncodePointer},
	{"EncodeSystemPointer", &EncodeSystemPointer},
	{"EnterCriticalSection", &EnterCriticalSection},
	{"EqualDomainSid", &EqualDomainSid},
	{"EqualPrefixSid", &EqualPrefixSid},
	{"EqualSid", &EqualSid},
	{"EraseTape", &EraseTape},
	{"EscapeCommFunction", &EscapeCommFunction},
	{"ExitProcess", &ExitProcess},
	{"ExitThread", &ExitThread},
	{"FatalExit", &FatalExit},
	{"FileTimeToDosDateTime", &FileTimeToDosDateTime},
	{"FileTimeToLocalFileTime", &FileTimeToLocalFileTime},
	{"FileTimeToSystemTime", &FileTimeToSystemTime},
	{"FindActCtxSectionGuid", &FindActCtxSectionGuid},
	{"FindClose", &FindClose},
	{"FindCloseChangeNotification", &FindCloseChangeNotification},
	{"FindFirstFreeAce", &FindFirstFreeAce},
	{"FindNextChangeNotification", &FindNextChangeNotification},
	{"FindVolumeClose", &FindVolumeClose},
	{"FindVolumeMountPointClose", &FindVolumeMountPointClose},
	{"FlushFileBuffers", &FlushFileBuffers},
	{"FlushInstructionCache", &FlushInstructionCache},
	{"FlushViewOfFile", &FlushViewOfFile},
	{"FreeLibrary", &FreeLibrary},
	{"FreeLibraryAndExitThread", &FreeLibraryAndExitThread},
	{"FreeResource", &FreeResource},
	{"FreeSid", &FreeSid},
	{"FreeUserPhysicalPages", &FreeUserPhysicalPages},
	{"GetAce", &GetAce},
	{"GetAclInformation", &GetAclInformation},
	{"GetCommConfig", &GetCommConfig},
	{"GetCommMask", &GetCommMask},
	{"GetCommModemStatus", &GetCommModemStatus},
	{"GetCommProperties", &GetCommProperties},
	{"GetCommState", &GetCommState},
	{"GetCommTimeouts", &GetCommTimeouts},
	{"GetCurrentActCtx", &GetCurrentActCtx},
	{"GetCurrentProcess", &GetCurrentProcess},
	{"GetCurrentProcessId", &GetCurrentProcessId},
	{"GetCurrentThread", &GetCurrentThread},
	{"GetCurrentThreadId", &GetCurrentThreadId},
	{"GetDevicePowerState", &GetDevicePowerState},
	{"GetEventLogInformation", &GetEventLogInformation},
	{"GetExitCodeProcess", &GetExitCodeProcess},
	{"GetExitCodeThread", &GetExitCodeThread},
	{"GetFileInformationByHandle", &GetFileInformationByHandle},
	{"GetFileSize", &GetFileSize},
	{"GetFileSizeEx", &GetFileSizeEx},
	{"GetFileTime", &GetFileTime},
	{"GetFileType", &GetFileType},
	{"GetHandleInformation", &GetHandleInformation},
	{"GetKernelObjectSecurity", &GetKernelObjectSecurity},
	{"GetLastError", &GetLastError},
	{"GetLengthSid", &GetLengthSid},
	{"GetLocalTime", &GetLocalTime},
	{"GetLogicalDrives", &GetLogicalDrives},
	{"GetLogicalProcessorInformation", &GetLogicalProcessorInformation},
	{"GetMailslotInfo", &GetMailslotInfo},
	{"GetNamedPipeInfo", &GetNamedPipeInfo},
	{"GetNativeSystemInfo", &GetNativeSystemInfo},
	{"GetNumaAvailableMemoryNode", &GetNumaAvailableMemoryNode},
	{"GetNumaHighestNodeNumber", &GetNumaHighestNodeNumber},
	{"GetNumaNodeProcessorMask", &GetNumaNodeProcessorMask},
	{"GetNumaProcessorNode", &GetNumaProcessorNode},
	{"GetNumberOfEventLogRecords", &GetNumberOfEventLogRecords},
	{"GetOldestEventLogRecord", &GetOldestEventLogRecord},
	{"GetOverlappedResult", &GetOverlappedResult},
	{"GetPriorityClass", &GetPriorityClass},
	{"GetPrivateObjectSecurity", &GetPrivateObjectSecurity},
	{"GetProcAddress", &GetProcAddress},
	{"GetProcessAffinityMask", &GetProcessAffinityMask},
	{"GetProcessHandleCount", &GetProcessHandleCount},
	{"GetProcessHeap", &GetProcessHeap},
	{"GetProcessHeaps", &GetProcessHeaps},
	{"GetProcessId", &GetProcessId},
	{"GetProcessIoCounters", &GetProcessIoCounters},
	{"GetProcessPriorityBoost", &GetProcessPriorityBoost},
	{"GetProcessShutdownParameters", &GetProcessShutdownParameters},
	{"GetProcessTimes", &GetProcessTimes},
	{"GetProcessVersion", &GetProcessVersion},
	{"GetProcessWorkingSetSize", &GetProcessWorkingSetSize},
	{"GetQueuedCompletionStatus", &GetQueuedCompletionStatus},
	{"GetSecurityDescriptorControl", &GetSecurityDescriptorControl},
	{"GetSecurityDescriptorDacl", &GetSecurityDescriptorDacl},
	{"GetSecurityDescriptorGroup", &GetSecurityDescriptorGroup},
	{"GetSecurityDescriptorLength", &GetSecurityDescriptorLength},
	{"GetSecurityDescriptorOwner", &GetSecurityDescriptorOwner},
	{"GetSecurityDescriptorRMControl", &GetSecurityDescriptorRMControl},
	{"GetSecurityDescriptorSacl", &GetSecurityDescriptorSacl},
	{"GetSidIdentifierAuthority", &GetSidIdentifierAuthority},
	{"GetSidLengthRequired", &GetSidLengthRequired},
	{"GetSidSubAuthority", &GetSidSubAuthority},
	{"GetSidSubAuthorityCount", &GetSidSubAuthorityCount},
	{"GetStdHandle", &GetStdHandle},
	{"GetSystemInfo", &GetSystemInfo},
	{"GetSystemPowerStatus", &GetSystemPowerStatus},
	{"GetSystemRegistryQuota", &GetSystemRegistryQuota},
	{"GetSystemTime", &GetSystemTime},
	{"GetSystemTimeAdjustment", &GetSystemTimeAdjustment},
	{"GetSystemTimeAsFileTime", &GetSystemTimeAsFileTime},
	{"GetSystemTimes", &GetSystemTimes},
	{"GetTapeParameters", &GetTapeParameters},
	{"GetTapePosition", &GetTapePosition},
	{"GetTapeStatus", &GetTapeStatus},
	{"GetThreadContext", &GetThreadContext},
	{"GetThreadIOPendingFlag", &GetThreadIOPendingFlag},
	{"GetThreadPriority", &GetThreadPriority},
	{"GetThreadPriorityBoost", &GetThreadPriorityBoost},
	{"GetThreadSelectorEntry", &GetThreadSelectorEntry},
	{"GetThreadTimes", &GetThreadTimes},
	{"GetTickCount", &GetTickCount},
	{"GetTimeZoneInformation", &GetTimeZoneInformation},
	{"GetTokenInformation", &GetTokenInformation},
	{"GetVersion", &GetVersion},
	{"GetWindowsAccountDomainSid", &GetWindowsAccountDomainSid},
	{"GetWriteWatch", &GetWriteWatch},
	{"GlobalAlloc", &GlobalAlloc},
	{"GlobalCompact", &GlobalCompact},
	{"GlobalDeleteAtom", &GlobalDeleteAtom},
	{"GlobalFix", &GlobalFix},
	{"GlobalFlags", &GlobalFlags},
	{"GlobalFree", &GlobalFree},
	{"GlobalHandle", &GlobalHandle},
	{"GlobalLock", &GlobalLock},
	{"GlobalMemoryStatus", &GlobalMemoryStatus},
	{"GlobalMemoryStatusEx", &GlobalMemoryStatusEx},
	{"GlobalReAlloc", &GlobalReAlloc},
	{"GlobalSize", &GlobalSize},
	{"GlobalUnfix", &GlobalUnfix},
	{"GlobalUnlock", &GlobalUnlock},
	{"GlobalUnWire", &GlobalUnWire},
	{"GlobalWire", &GlobalWire},
	{"HeapAlloc", &HeapAlloc},
	{"HeapCompact", &HeapCompact},
	{"HeapCreate", &HeapCreate},
	{"HeapDestroy", &HeapDestroy},
	{"HeapFree", &HeapFree},
	{"HeapLock", &HeapLock},
	{"HeapQueryInformation", &HeapQueryInformation},
	{"HeapReAlloc", &HeapReAlloc},
	{"HeapSetInformation", &HeapSetInformation},
	{"HeapSize", &HeapSize},
	{"HeapUnlock", &HeapUnlock},
	{"HeapValidate", &HeapValidate},
	{"HeapWalk", &HeapWalk},
	{"ImpersonateAnonymousToken", &ImpersonateAnonymousToken},
	{"ImpersonateLoggedOnUser", &ImpersonateLoggedOnUser},
	{"ImpersonateNamedPipeClient", &ImpersonateNamedPipeClient},
	{"ImpersonateSelf", &ImpersonateSelf},
	{"InitAtomTable", &InitAtomTable},
	{"InitializeAcl", &InitializeAcl},
	{"InitializeCriticalSection", &InitializeCriticalSection},
	{"InitializeCriticalSectionAndSpinCount", &InitializeCriticalSectionAndSpinCount},
	{"InitializeSecurityDescriptor", &InitializeSecurityDescriptor},
	{"InitializeSid", &InitializeSid},
	{"InitializeSListHead", &InitializeSListHead},
	{"InterlockedCompareExchange", &InterlockedCompareExchange},
	{"InterlockedDecrement", &InterlockedDecrement},
	{"InterlockedExchange", &InterlockedExchange},
	{"InterlockedExchangeAdd", &InterlockedExchangeAdd},
	{"InterlockedFlushSList", &InterlockedFlushSList},
	{"InterlockedIncrement", &InterlockedIncrement},
	{"InterlockedPopEntrySList", &InterlockedPopEntrySList},
	{"InterlockedPushEntrySList", &InterlockedPushEntrySList},
	{"IsBadCodePtr", &IsBadCodePtr},
	{"IsBadHugeReadPtr", &IsBadHugeReadPtr},
	{"IsBadHugeWritePtr", &IsBadHugeWritePtr},
	{"IsBadReadPtr", &IsBadReadPtr},
	{"IsBadWritePtr", &IsBadWritePtr},
	{"IsDebuggerPresent", &IsDebuggerPresent},
	{"IsProcessInJob", &IsProcessInJob},
	{"IsProcessorFeaturePresent", &IsProcessorFeaturePresent},
	{"IsSystemResumeAutomatic", &IsSystemResumeAutomatic},
	{"IsTextUnicode", &IsTextUnicode},
	{"IsTokenRestricted", &IsTokenRestricted},
	{"IsTokenUntrusted", &IsTokenUntrusted},
	{"IsValidAcl", &IsValidAcl},
	{"IsValidSecurityDescriptor", &IsValidSecurityDescriptor},
	{"IsValidSid", &IsValidSid},
	{"IsWellKnownSid", &IsWellKnownSid},
	{"IsWow64Process", &IsWow64Process},
	{"LeaveCriticalSection", &LeaveCriticalSection},
	{"LoadModule", &LoadModule},
	{"LoadResource", &LoadResource},
	{"LocalAlloc", &LocalAlloc},
	{"LocalCompact", &LocalCompact},
	{"LocalFileTimeToFileTime", &LocalFileTimeToFileTime},
	{"LocalFlags", &LocalFlags},
	{"LocalFree", &LocalFree},
	{"LocalHandle", &LocalHandle},
	{"LocalLock", &LocalLock},
	{"LocalReAlloc", &LocalReAlloc},
	{"LocalShrink", &LocalShrink},
	{"LocalSize", &LocalSize},
	{"LocalUnlock", &LocalUnlock},
	{"LockFile", &LockFile},
	{"LockFileEx", &LockFileEx},
	{"LockResource", &LockResource},
	{"MakeAbsoluteSD", &MakeAbsoluteSD},
	{"MakeAbsoluteSD2", &MakeAbsoluteSD2},
	{"MakeSelfRelativeSD", &MakeSelfRelativeSD},
	{"MapGenericMask", &MapGenericMask},
	{"MapUserPhysicalPages", &MapUserPhysicalPages},
	{"MapUserPhysicalPagesScatter", &MapUserPhysicalPagesScatter},
	{"MapViewOfFile", &MapViewOfFile},
	{"MapViewOfFileEx", &MapViewOfFileEx},
	{"MulDiv", &MulDiv},
	{"NotifyChangeEventLog", &NotifyChangeEventLog},
	{"OpenFile", &OpenFile},
	{"OpenProcess", &OpenProcess},
	{"OpenProcessToken", &OpenProcessToken},
	{"OpenThread", &OpenThread},
	{"OpenThreadToken", &OpenThreadToken},
	{"PeekNamedPipe", &PeekNamedPipe},
	{"PostQueuedCompletionStatus", &PostQueuedCompletionStatus},
	{"PrepareTape", &PrepareTape},
	{"PrivilegeCheck", &PrivilegeCheck},
	{"ProcessIdToSessionId", &ProcessIdToSessionId},
	{"PulseEvent", &PulseEvent},
	{"PurgeComm", &PurgeComm},
	{"QueryActCtxW", &QueryActCtxW},
	{"QueryDepthSList", &QueryDepthSList},
	{"QueryInformationJobObject", &QueryInformationJobObject},
	{"QueryMemoryResourceNotification", &QueryMemoryResourceNotification},
	{"QueryPerformanceCounter", &QueryPerformanceCounter},
	{"QueryPerformanceFrequency", &QueryPerformanceFrequency},
	{"QueueUserAPC", &QueueUserAPC},
	{"QueueUserWorkItem", &QueueUserWorkItem},
	{"RaiseException", &RaiseException},
	{"ReadEncryptedFileRaw", &ReadEncryptedFileRaw},
	{"ReadFile", &ReadFile},
	{"ReadFileEx", &ReadFileEx},
	{"ReadFileScatter", &ReadFileScatter},
	{"ReadProcessMemory", &ReadProcessMemory},
	{"RegisterWaitForSingleObject", &RegisterWaitForSingleObject},
	{"RegisterWaitForSingleObjectEx", &RegisterWaitForSingleObjectEx},
	{"ReleaseActCtx", &ReleaseActCtx},
	{"ReleaseMutex", &ReleaseMutex},
	{"ReleaseSemaphore", &ReleaseSemaphore},
	{"RemoveVectoredExceptionHandler", &RemoveVectoredExceptionHandler},
	{"RequestDeviceWakeup", &RequestDeviceWakeup},
	{"RequestWakeupLatency", &RequestWakeupLatency},
	{"ResetEvent", &ResetEvent},
	{"ResetWriteWatch", &ResetWriteWatch},
	{"RestoreLastError", &RestoreLastError},
	{"ResumeThread", &ResumeThread},
	{"RevertToSelf", &RevertToSelf},
	{"SetAclInformation", &SetAclInformation},
	{"SetCommBreak", &SetCommBreak},
	{"SetCommConfig", &SetCommConfig},
	{"SetCommMask", &SetCommMask},
	{"SetCommState", &SetCommState},
	{"SetCommTimeouts", &SetCommTimeouts},
	{"SetCriticalSectionSpinCount", &SetCriticalSectionSpinCount},
	{"SetEndOfFile", &SetEndOfFile},
	{"SetErrorMode", &SetErrorMode},
	{"SetEvent", &SetEvent},
	{"SetFileApisToANSI", &SetFileApisToANSI},
	{"SetFileApisToOEM", &SetFileApisToOEM},
	{"SetFilePointer", &SetFilePointer},
	{"SetFilePointerEx", &SetFilePointerEx},
	{"SetFileTime", &SetFileTime},
	{"SetFileValidData", &SetFileValidData},
	{"SetHandleCount", &SetHandleCount},
	{"SetHandleInformation", &SetHandleInformation},
	{"SetInformationJobObject", &SetInformationJobObject},
	{"SetKernelObjectSecurity", &SetKernelObjectSecurity},
	{"SetLastError", &SetLastError},
	{"SetLocalTime", &SetLocalTime},
	{"SetMailslotInfo", &SetMailslotInfo},
	{"SetMessageWaitingIndicator", &SetMessageWaitingIndicator},
	{"SetNamedPipeHandleState", &SetNamedPipeHandleState},
	{"SetPriorityClass", &SetPriorityClass},
	{"SetPrivateObjectSecurity", &SetPrivateObjectSecurity},
	{"SetPrivateObjectSecurityEx", &SetPrivateObjectSecurityEx},
	{"SetProcessAffinityMask", &SetProcessAffinityMask},
	{"SetProcessPriorityBoost", &SetProcessPriorityBoost},
	{"SetProcessShutdownParameters", &SetProcessShutdownParameters},
	{"SetProcessWorkingSetSize", &SetProcessWorkingSetSize},
	{"SetSecurityDescriptorControl", &SetSecurityDescriptorControl},
	{"SetSecurityDescriptorDacl", &SetSecurityDescriptorDacl},
	{"SetSecurityDescriptorGroup", &SetSecurityDescriptorGroup},
	{"SetSecurityDescriptorOwner", &SetSecurityDescriptorOwner},
	{"SetSecurityDescriptorRMControl", &SetSecurityDescriptorRMControl},
	{"SetSecurityDescriptorSacl", &SetSecurityDescriptorSacl},
	{"SetStdHandle", &SetStdHandle},
	{"SetSystemPowerState", &SetSystemPowerState},
	{"SetSystemTime", &SetSystemTime},
	{"SetSystemTimeAdjustment", &SetSystemTimeAdjustment},
	{"SetTapeParameters", &SetTapeParameters},
	{"SetTapePosition", &SetTapePosition},
	{"SetThreadAffinityMask", &SetThreadAffinityMask},
	{"SetThreadContext", &SetThreadContext},
	{"SetThreadExecutionState", &SetThreadExecutionState},
	{"SetThreadIdealProcessor", &SetThreadIdealProcessor},
	{"SetThreadPriority", &SetThreadPriority},
	{"SetThreadPriorityBoost", &SetThreadPriorityBoost},
	{"SetThreadToken", &SetThreadToken},
	{"SetTimerQueueTimer", &SetTimerQueueTimer},
	{"SetTimeZoneInformation", &SetTimeZoneInformation},
	{"SetTokenInformation", &SetTokenInformation},
	{"SetUnhandledExceptionFilter", &SetUnhandledExceptionFilter},
	{"SetupComm", &SetupComm},
	{"SetWaitableTimer", &SetWaitableTimer},
	{"SignalObjectAndWait", &SignalObjectAndWait},
	{"SizeofResource", &SizeofResource},
	{"Sleep", &Sleep},
	{"SleepEx", &SleepEx},
	{"SuspendThread", &SuspendThread},
	{"SwitchToFiber", &SwitchToFiber},
	{"SwitchToThread", &SwitchToThread},
	{"SystemTimeToFileTime", &SystemTimeToFileTime},
	{"SystemTimeToTzSpecificLocalTime", &SystemTimeToTzSpecificLocalTime},
	{"TerminateJobObject", &TerminateJobObject},
	{"TerminateProcess", &TerminateProcess},
	{"TerminateThread", &TerminateThread},
	{"TlsAlloc", &TlsAlloc},
	{"TlsFree", &TlsFree},
	{"TlsGetValue", &TlsGetValue},
	{"TlsSetValue", &TlsSetValue},
	{"TransactNamedPipe", &TransactNamedPipe},
	{"TransmitCommChar", &TransmitCommChar},
	{"TryEnterCriticalSection", &TryEnterCriticalSection},
	{"TzSpecificLocalTimeToSystemTime", &TzSpecificLocalTimeToSystemTime},
	{"UnlockFile", &UnlockFile},
	{"UnlockFileEx", &UnlockFileEx},
	{"UnmapViewOfFile", &UnmapViewOfFile},
	{"UnregisterWait", &UnregisterWait},
	{"UnregisterWaitEx", &UnregisterWaitEx},
	{"VirtualAlloc", &VirtualAlloc},
	{"VirtualAllocEx", &VirtualAllocEx},
	{"VirtualFree", &VirtualFree},
	{"VirtualFreeEx", &VirtualFreeEx},
	{"VirtualLock", &VirtualLock},
	{"VirtualProtect", &VirtualProtect},
	{"VirtualProtectEx", &VirtualProtectEx},
	{"VirtualQuery", &VirtualQuery},
	{"VirtualQueryEx", &VirtualQueryEx},
	{"VirtualUnlock", &VirtualUnlock},
	{"WaitCommEvent", &WaitCommEvent},
	{"WaitForDebugEvent", &WaitForDebugEvent},
	{"WaitForMultipleObjects", &WaitForMultipleObjects},
	{"WaitForMultipleObjectsEx", &WaitForMultipleObjectsEx},
	{"WaitForSingleObject", &WaitForSingleObject},
	{"WaitForSingleObjectEx", &WaitForSingleObjectEx},
	{"WinExec", &WinExec},
	{"WriteEncryptedFileRaw", &WriteEncryptedFileRaw},
	{"WriteFile", &WriteFile},
	{"WriteFileEx", &WriteFileEx},
	{"WriteFileGather", &WriteFileGather},
	{"WriteProcessMemory", &WriteProcessMemory},
	{"WriteTapemark", &WriteTapemark},
	{"WTSGetActiveConsoleSessionId", &WTSGetActiveConsoleSessionId},
	{"ZombifyActCtx", &ZombifyActCtx},

	//TODO(t): ALWAYS Unicode so currently won't work!
	// No structure handling yet so 2 reasons

	{"CreateProcessWithLogonW", &CreateProcessWithLogonW},
	{"ReadDirectoryChangesW", &ReadDirectoryChangesW},

	//TODO(): Workaround?
	{"FindFirstFileA", &FindFirstFileA},
	{"FindFirstFileW", &FindFirstFileW},
	{"FindNextFileA", &FindNextFileA},
	{"FindNextFileW", &FindNextFileW},
	{"GetVersionExA", &GetVersionExA},
	{"GetVersionExW", &GetVersionExW},
	{"VerifyVersionInfoA", &VerifyVersionInfoA},
	{"VerifyVersionInfoW", &VerifyVersionInfoW},
}

//TODO(t): Not xp32
//{"CreateProcessWithTokenW", &CreateProcessWithTokenW},
//{"NeedCurrentDirectoryForExePathW", &NeedCurrentDirectoryForExePath},
//{"ReOpenFileW", &ReOpenFile},
//{"SetEnvironmentStringsW", &SetEnvironmentStrings},

var WinBaseUnicodeApis = outside.Apis{
	{"AccessCheckAndAuditAlarmW", &AccessCheckAndAuditAlarm},
	{"AccessCheckByTypeAndAuditAlarmW", &AccessCheckByTypeAndAuditAlarm},
	{"AccessCheckByTypeResultListAndAuditAlarmByHandleW", &AccessCheckByTypeResultListAndAuditAlarmByHandle},
	{"AccessCheckByTypeResultListAndAuditAlarmW", &AccessCheckByTypeResultListAndAuditAlarm},
	{"AddAtomW", &AddAtom},
	{"BackupEventLogW", &BackupEventLog},
	{"BeginUpdateResourceW", &BeginUpdateResource},
	{"BuildCommDCBAndTimeoutsW", &BuildCommDCBAndTimeouts},
	{"BuildCommDCBW", &BuildCommDCB},
	{"CallNamedPipeW", &CallNamedPipe},
	{"CheckNameLegalDOS8Dot3W", &CheckNameLegalDOS8Dot3},
	{"ClearEventLogW", &ClearEventLog},
	{"CommConfigDialogW", &CommConfigDialog},
	{"CopyFileExW", &CopyFileEx},
	{"CopyFileW", &CopyFile},
	{"CreateActCtxW", &CreateActCtx},
	{"CreateDirectoryExW", &CreateDirectoryEx},
	{"CreateDirectoryW", &CreateDirectory},
	{"CreateEventW", &CreateEvent},
	{"CreateFileMappingW", &CreateFileMapping},
	{"CreateFileW", &CreateFile},
	{"CreateHardLinkW", &CreateHardLink},
	{"CreateJobObjectW", &CreateJobObject},
	{"CreateMailslotW", &CreateMailslot},
	{"CreateMutexW", &CreateMutex},
	{"CreateNamedPipeW", &CreateNamedPipe},
	{"CreateProcessAsUserW", &CreateProcessAsUser},
	{"CreateProcessW", &CreateProcess},
	{"CreateSemaphoreW", &CreateSemaphore},
	{"CreateWaitableTimerW", &CreateWaitableTimer},
	{"DecryptFileW", &DecryptFile},
	{"DefineDosDeviceW", &DefineDosDevice},
	{"DeleteFileW", &DeleteFile},
	{"DeleteVolumeMountPointW", &DeleteVolumeMountPoint},
	{"DnsHostnameToComputerNameW", &DnsHostnameToComputerName},
	{"EncryptFileW", &EncryptFile},
	{"EnumResourceLanguagesW", &EnumResourceLanguages},
	{"EnumResourceNamesW", &EnumResourceNames},
	{"EnumResourceTypesW", &EnumResourceTypes},
	{"ExpandEnvironmentStringsW", &ExpandEnvironmentStrings},
	{"FatalAppExitW", &FatalAppExit},
	{"FileEncryptionStatusW", &FileEncryptionStatus},
	{"FindActCtxSectionStringW", &FindActCtxSectionString},
	{"FindAtomW", &FindAtom},
	{"FindFirstChangeNotificationW", &FindFirstChangeNotification},
	{"FindFirstFileExW", &FindFirstFileEx},
	{"FindFirstVolumeMountPointW", &FindFirstVolumeMountPoint},
	{"FindFirstVolumeW", &FindFirstVolume},
	{"FindNextVolumeMountPointW", &FindNextVolumeMountPoint},
	{"FindNextVolumeW", &FindNextVolume},
	{"FindResourceExW", &FindResourceEx},
	{"FindResourceW", &FindResource},
	{"FormatMessageW", &FormatMessage},
	{"FreeEnvironmentStringsW", &FreeEnvironmentStrings},
	{"GetAtomNameW", &GetAtomName},
	{"GetBinaryTypeW", &GetBinaryType},
	{"GetCommandLineW", &GetCommandLine},
	{"GetCompressedFileSizeW", &GetCompressedFileSize},
	{"GetComputerNameExW", &GetComputerNameEx},
	{"GetComputerNameW", &GetComputerName},
	{"GetCurrentDirectoryW", &GetCurrentDirectory},
	{"GetCurrentHwProfileW", &GetCurrentHwProfileW},
	{"GetDefaultCommConfigW", &GetDefaultCommConfig},
	{"GetDiskFreeSpaceExW", &GetDiskFreeSpaceEx},
	{"GetDiskFreeSpaceW", &GetDiskFreeSpace},
	{"GetDllDirectoryW", &GetDllDirectory},
	{"GetDriveTypeW", &GetDriveType},
	{"GetEnvironmentStringsW", &GetEnvironmentStrings},
	{"GetEnvironmentVariableW", &GetEnvironmentVariable},
	{"GetFileAttributesExW", &GetFileAttributesEx},
	{"GetFileAttributesW", &GetFileAttributes},
	{"GetFileSecurityW", &GetFileSecurity},
	{"GetFirmwareEnvironmentVariableW", &GetFirmwareEnvironmentVariable},
	{"GetFullPathNameW", &GetFullPathName},
	{"GetLogicalDriveStringsW", &GetLogicalDriveStrings},
	{"GetLongPathNameW", &GetLongPathName},
	{"GetModuleFileNameW", &GetModuleFileName},
	{"GetModuleHandleExW", &GetModuleHandleEx},
	{"GetModuleHandleW", &GetModuleHandle},
	{"GetNamedPipeHandleStateW", &GetNamedPipeHandleState},
	{"GetPrivateProfileIntW", &GetPrivateProfileInt},
	{"GetPrivateProfileSectionNamesW", &GetPrivateProfileSectionNames},
	{"GetPrivateProfileSectionW", &GetPrivateProfileSection},
	{"GetPrivateProfileStringW", &GetPrivateProfileString},
	{"GetPrivateProfileStructW", &GetPrivateProfileStruct},
	{"GetProfileIntW", &GetProfileInt},
	{"GetProfileSectionW", &GetProfileSection},
	{"GetProfileStringW", &GetProfileString},
	{"GetShortPathNameW", &GetShortPathName},
	{"GetStartupInfoW", &GetStartupInfo},
	{"GetSystemDirectoryW", &GetSystemDirectory},
	{"GetSystemWindowsDirectoryW", &GetSystemWindowsDirectory},
	{"GetSystemWow64DirectoryW", &GetSystemWow64Directory},
	{"GetTempFileNameW", &GetTempFileName},
	{"GetTempPathW", &GetTempPath},
	{"GetUserNameW", &GetUserName},
	{"GetVolumeInformationW", &GetVolumeInformation},
	{"GetVolumeNameForVolumeMountPointW", &GetVolumeNameForVolumeMountPoint},
	{"GetVolumePathNamesForVolumeNameW", &GetVolumePathNamesForVolumeName},
	{"GetVolumePathNameW", &GetVolumePathName},
	{"GetWindowsDirectoryW", &GetWindowsDirectory},
	{"GlobalAddAtomW", &GlobalAddAtom},
	{"GlobalFindAtomW", &GlobalFindAtom},
	{"GlobalGetAtomNameW", &GlobalGetAtomName},
	{"IsBadStringPtrW", &IsBadStringPtr},
	{"LoadLibraryExW", &LoadLibraryEx},
	{"LoadLibraryW", &LoadLibrary},
	{"LogonUserExW", &LogonUserEx},
	{"LogonUserW", &LogonUser},
	{"LookupAccountNameW", &LookupAccountName},
	{"LookupAccountSidW", &LookupAccountSid},
	{"LookupPrivilegeDisplayNameW", &LookupPrivilegeDisplayName},
	{"LookupPrivilegeNameW", &LookupPrivilegeName},
	{"LookupPrivilegeValueW", &LookupPrivilegeValue},
	{"lstrcatW", &Lstrcat},
	{"lstrcmpiW", &Lstrcmpi},
	{"lstrcmpW", &Lstrcmp},
	{"lstrcpynW", &Lstrcpyn},
	{"lstrcpyW", &Lstrcpy},
	{"lstrlenW", &Lstrlen},
	{"MoveFileExW", &MoveFileEx},
	{"MoveFileW", &MoveFile},
	{"MoveFileWithProgressW", &MoveFileWithProgress},
	{"ObjectCloseAuditAlarmW", &ObjectCloseAuditAlarm},
	{"ObjectDeleteAuditAlarmW", &ObjectDeleteAuditAlarm},
	{"ObjectOpenAuditAlarmW", &ObjectOpenAuditAlarm},
	{"ObjectPrivilegeAuditAlarmW", &ObjectPrivilegeAuditAlarm},
	{"OpenBackupEventLogW", &OpenBackupEventLog},
	{"OpenEncryptedFileRawW", &OpenEncryptedFileRaw},
	{"OpenEventLogW", &OpenEventLog},
	{"OpenEventW", &OpenEvent},
	{"OpenFileMappingW", &OpenFileMapping},
	{"OpenJobObjectW", &OpenJobObject},
	{"OpenMutexW", &OpenMutex},
	{"OpenSemaphoreW", &OpenSemaphore},
	{"OpenWaitableTimerW", &OpenWaitableTimer},
	{"OutputDebugStringW", &OutputDebugString},
	{"PrivilegedServiceAuditAlarmW", &PrivilegedServiceAuditAlarm},
	{"QueryDosDeviceW", &QueryDosDevice},
	{"ReadDirectoryChangesW", &ReadDirectoryChangesW},
	{"ReadEventLogW", &ReadEventLog},
	{"RegisterEventSourceW", &RegisterEventSource},
	{"RemoveDirectoryW", &RemoveDirectory},
	{"ReplaceFileW", &ReplaceFile},
	{"ReportEventW", &ReportEvent},
	{"SearchPathW", &SearchPath},
	{"SetComputerNameExW", &SetComputerNameEx},
	{"SetComputerNameW", &SetComputerName},
	{"SetCurrentDirectoryW", &SetCurrentDirectory},
	{"SetDefaultCommConfigW", &SetDefaultCommConfig},
	{"SetDllDirectoryW", &SetDllDirectory},
	{"SetEnvironmentVariableW", &SetEnvironmentVariable},
	{"SetFileAttributesW", &SetFileAttributes},
	{"SetFileSecurityW", &SetFileSecurity},
	{"SetFileShortNameW", &SetFileShortName},
	{"SetFirmwareEnvironmentVariableW", &SetFirmwareEnvironmentVariable},
	{"SetVolumeLabelW", &SetVolumeLabel},
	{"SetVolumeMountPointW", &SetVolumeMountPoint},
	{"UpdateResourceW", &UpdateResource},
	{"WritePrivateProfileSectionW", &WritePrivateProfileSection},
	{"WritePrivateProfileStringW", &WritePrivateProfileString},
	{"WritePrivateProfileStructW", &WritePrivateProfileStruct},
	{"WriteProfileSectionW", &WriteProfileSection},
	{"WriteProfileStringW", &WriteProfileString},
}

//TODO(t):Not on xp32
//{"NeedCurrentDirectoryForExePathA", &NeedCurrentDirectoryForExePath},
//{"ReOpenFileA", &ReOpenFile},
//{"SetEnvironmentStringsA", &SetEnvironmentStrings},

var WinBaseANSIApis = outside.Apis{
	{"AccessCheckAndAuditAlarmA", &AccessCheckAndAuditAlarm},
	{"AccessCheckByTypeAndAuditAlarmA", &AccessCheckByTypeAndAuditAlarm},
	{"AccessCheckByTypeResultListAndAuditAlarmA", &AccessCheckByTypeResultListAndAuditAlarm},
	{"AccessCheckByTypeResultListAndAuditAlarmByHandleA", &AccessCheckByTypeResultListAndAuditAlarmByHandle},
	{"AddAtomA", &AddAtom},
	{"BackupEventLogA", &BackupEventLog},
	{"BeginUpdateResourceA", &BeginUpdateResource},
	{"BuildCommDCBA", &BuildCommDCB},
	{"BuildCommDCBAndTimeoutsA", &BuildCommDCBAndTimeouts},
	{"CallNamedPipeA", &CallNamedPipe},
	{"CheckNameLegalDOS8Dot3A", &CheckNameLegalDOS8Dot3},
	{"ClearEventLogA", &ClearEventLog},
	{"CommConfigDialogA", &CommConfigDialog},
	{"CopyFileA", &CopyFile},
	{"CopyFileExA", &CopyFileEx},
	{"CreateActCtxA", &CreateActCtx},
	{"CreateDirectoryA", &CreateDirectory},
	{"CreateDirectoryExA", &CreateDirectoryEx},
	{"CreateEventA", &CreateEvent},
	{"CreateFileA", &CreateFile},
	{"CreateFileMappingA", &CreateFileMapping},
	{"CreateHardLinkA", &CreateHardLink},
	{"CreateJobObjectA", &CreateJobObject},
	{"CreateMailslotA", &CreateMailslot},
	{"CreateMutexA", &CreateMutex},
	{"CreateNamedPipeA", &CreateNamedPipe},
	{"CreateProcessA", &CreateProcess},
	{"CreateProcessAsUserA", &CreateProcessAsUser},
	{"CreateSemaphoreA", &CreateSemaphore},
	{"CreateWaitableTimerA", &CreateWaitableTimer},
	{"DecryptFileA", &DecryptFile},
	{"DefineDosDeviceA", &DefineDosDevice},
	{"DeleteFileA", &DeleteFile},
	{"DeleteVolumeMountPointA", &DeleteVolumeMountPoint},
	{"DnsHostnameToComputerNameA", &DnsHostnameToComputerName},
	{"EncryptFileA", &EncryptFile},
	{"EnumResourceLanguagesA", &EnumResourceLanguages},
	{"EnumResourceNamesA", &EnumResourceNames},
	{"EnumResourceTypesA", &EnumResourceTypes},
	{"ExpandEnvironmentStringsA", &ExpandEnvironmentStrings},
	{"FatalAppExitA", &FatalAppExit},
	{"FileEncryptionStatusA", &FileEncryptionStatus},
	{"FindActCtxSectionStringA", &FindActCtxSectionString},
	{"FindAtomA", &FindAtom},
	{"FindFirstChangeNotificationA", &FindFirstChangeNotification},
	{"FindFirstFileExA", &FindFirstFileEx},
	{"FindFirstVolumeA", &FindFirstVolume},
	{"FindFirstVolumeMountPointA", &FindFirstVolumeMountPoint},
	{"FindNextVolumeA", &FindNextVolume},
	{"FindNextVolumeMountPointA", &FindNextVolumeMountPoint},
	{"FindResourceA", &FindResource},
	{"FindResourceExA", &FindResourceEx},
	{"FormatMessageA", &FormatMessage},
	{"FreeEnvironmentStringsA", &FreeEnvironmentStrings},
	{"GetAtomNameA", &GetAtomName},
	{"GetBinaryTypeA", &GetBinaryType},
	{"GetCommandLineA", &GetCommandLine},
	{"GetCompressedFileSizeA", &GetCompressedFileSize},
	{"GetComputerNameA", &GetComputerName},
	{"GetComputerNameExA", &GetComputerNameEx},
	{"GetCurrentDirectoryA", &GetCurrentDirectory},
	{"GetCurrentHwProfileA", &GetCurrentHwProfileA},
	{"GetDefaultCommConfigA", &GetDefaultCommConfig},
	{"GetDiskFreeSpaceA", &GetDiskFreeSpace},
	{"GetDiskFreeSpaceExA", &GetDiskFreeSpaceEx},
	{"GetDllDirectoryA", &GetDllDirectory},
	{"GetDriveTypeA", &GetDriveType},
	{"GetEnvironmentStringsA", &GetEnvironmentStrings},
	{"GetEnvironmentVariableA", &GetEnvironmentVariable},
	{"GetFileAttributesA", &GetFileAttributes},
	{"GetFileAttributesExA", &GetFileAttributesEx},
	{"GetFileSecurityA", &GetFileSecurity},
	{"GetFirmwareEnvironmentVariableA", &GetFirmwareEnvironmentVariable},
	{"GetFullPathNameA", &GetFullPathName},
	{"GetLogicalDriveStringsA", &GetLogicalDriveStrings},
	{"GetLongPathNameA", &GetLongPathName},
	{"GetModuleFileNameA", &GetModuleFileName},
	{"GetModuleHandleA", &GetModuleHandle},
	{"GetModuleHandleExA", &GetModuleHandleEx},
	{"GetNamedPipeHandleStateA", &GetNamedPipeHandleState},
	{"GetPrivateProfileIntA", &GetPrivateProfileInt},
	{"GetPrivateProfileSectionA", &GetPrivateProfileSection},
	{"GetPrivateProfileSectionNamesA", &GetPrivateProfileSectionNames},
	{"GetPrivateProfileStringA", &GetPrivateProfileString},
	{"GetPrivateProfileStructA", &GetPrivateProfileStruct},
	{"GetProfileIntA", &GetProfileInt},
	{"GetProfileSectionA", &GetProfileSection},
	{"GetProfileStringA", &GetProfileString},
	{"GetShortPathNameA", &GetShortPathName},
	{"GetStartupInfoA", &GetStartupInfo},
	{"GetSystemDirectoryA", &GetSystemDirectory},
	{"GetSystemWindowsDirectoryA", &GetSystemWindowsDirectory},
	{"GetSystemWow64DirectoryA", &GetSystemWow64Directory},
	{"GetTempFileNameA", &GetTempFileName},
	{"GetTempPathA", &GetTempPath},
	{"GetUserNameA", &GetUserName},
	{"GetVolumeInformationA", &GetVolumeInformation},
	{"GetVolumeNameForVolumeMountPointA", &GetVolumeNameForVolumeMountPoint},
	{"GetVolumePathNameA", &GetVolumePathName},
	{"GetVolumePathNamesForVolumeNameA", &GetVolumePathNamesForVolumeName},
	{"GetWindowsDirectoryA", &GetWindowsDirectory},
	{"GlobalAddAtomA", &GlobalAddAtom},
	{"GlobalFindAtomA", &GlobalFindAtom},
	{"GlobalGetAtomNameA", &GlobalGetAtomName},
	{"IsBadStringPtrA", &IsBadStringPtr},
	{"LoadLibraryA", &LoadLibrary},
	{"LoadLibraryExA", &LoadLibraryEx},
	{"LogonUserA", &LogonUser},
	{"LogonUserExA", &LogonUserEx},
	{"LookupAccountNameA", &LookupAccountName},
	{"LookupAccountSidA", &LookupAccountSid},
	{"LookupPrivilegeDisplayNameA", &LookupPrivilegeDisplayName},
	{"LookupPrivilegeNameA", &LookupPrivilegeName},
	{"LookupPrivilegeValueA", &LookupPrivilegeValue},
	{"lstrcatA", &Lstrcat},
	{"lstrcmpA", &Lstrcmp},
	{"lstrcmpiA", &Lstrcmpi},
	{"lstrcpyA", &Lstrcpy},
	{"lstrcpynA", &Lstrcpyn},
	{"lstrlenA", &Lstrlen},
	{"MoveFileA", &MoveFile},
	{"MoveFileExA", &MoveFileEx},
	{"MoveFileWithProgressA", &MoveFileWithProgress},
	{"ObjectCloseAuditAlarmA", &ObjectCloseAuditAlarm},
	{"ObjectDeleteAuditAlarmA", &ObjectDeleteAuditAlarm},
	{"ObjectOpenAuditAlarmA", &ObjectOpenAuditAlarm},
	{"ObjectPrivilegeAuditAlarmA", &ObjectPrivilegeAuditAlarm},
	{"OpenBackupEventLogA", &OpenBackupEventLog},
	{"OpenEncryptedFileRawA", &OpenEncryptedFileRaw},
	{"OpenEventA", &OpenEvent},
	{"OpenEventLogA", &OpenEventLog},
	{"OpenFileMappingA", &OpenFileMapping},
	{"OpenJobObjectA", &OpenJobObject},
	{"OpenMutexA", &OpenMutex},
	{"OpenSemaphoreA", &OpenSemaphore},
	{"OpenWaitableTimerA", &OpenWaitableTimer},
	{"OutputDebugStringA", &OutputDebugString},
	{"PrivilegedServiceAuditAlarmA", &PrivilegedServiceAuditAlarm},
	{"QueryDosDeviceA", &QueryDosDevice},
	{"ReadEventLogA", &ReadEventLog},
	{"RegisterEventSourceA", &RegisterEventSource},
	{"RemoveDirectoryA", &RemoveDirectory},
	{"ReplaceFileA", &ReplaceFile},
	{"ReportEventA", &ReportEvent},
	{"SearchPathA", &SearchPath},
	{"SetComputerNameA", &SetComputerName},
	{"SetComputerNameExA", &SetComputerNameEx},
	{"SetCurrentDirectoryA", &SetCurrentDirectory},
	{"SetDefaultCommConfigA", &SetDefaultCommConfig},
	{"SetDllDirectoryA", &SetDllDirectory},
	{"SetEnvironmentVariableA", &SetEnvironmentVariable},
	{"SetFileAttributesA", &SetFileAttributes},
	{"SetFileSecurityA", &SetFileSecurity},
	{"SetFileShortNameA", &SetFileShortName},
	{"SetFirmwareEnvironmentVariableA", &SetFirmwareEnvironmentVariable},
	{"SetVolumeLabelA", &SetVolumeLabel},
	{"SetVolumeMountPointA", &SetVolumeMountPoint},
	{"UpdateResourceA", &UpdateResource},
	{"WaitNamedPipeA", &WaitNamedPipe},
	{"WritePrivateProfileSectionA", &WritePrivateProfileSection},
	{"WritePrivateProfileStringA", &WritePrivateProfileString},
	{"WritePrivateProfileStructA", &WritePrivateProfileStruct},
	{"WriteProfileSectionA", &WriteProfileSection},
	{"WriteProfileStringA", &WriteProfileString},
}
