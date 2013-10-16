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
		accessStatus *T.BOOL) T.BOOL

	AccessCheckAndAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		objectTypeName, objectName VString,
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		desiredAccess T.DWORD,
		genericMapping *T.GENERIC_MAPPING,
		objectCreation T.BOOL,
		grantedAccess *T.DWORD,
		accessStatus, generateOnClose *T.BOOL) T.BOOL

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
		accessStatus *T.BOOL) T.BOOL

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
		accessStatus, generateOnClose *T.BOOL) T.BOOL

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
		accessStatusList *T.DWORD) T.BOOL

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
		generateOnClose *T.BOOL) T.BOOL

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
		generateOnClose *T.BOOL) T.BOOL

	ActivateActCtx func(
		actCtx T.HANDLE, cookie *T.ULONG_PTR) T.BOOL

	AddAccessAllowedAce func(
		acl *T.ACL,
		aceRevision, accessMask T.DWORD,
		sid *T.SID) T.BOOL

	AddAccessAllowedAceEx func(
		acl *T.ACL,
		aceRevision, aceFlags, accessMask T.DWORD,
		sid *T.SID) T.BOOL

	AddAccessAllowedObjectAce func(
		acl *T.ACL,
		aceRevision, aceFlags, accessMask T.DWORD,
		objectTypeGuid, inheritedObjectTypeGuid *T.GUID,
		sid *T.SID) T.BOOL

	AddAccessDeniedAce func(
		acl *T.ACL,
		aceRevision, accessMask T.DWORD,
		sid *T.SID) T.BOOL

	AddAccessDeniedAceEx func(
		acl *T.ACL,
		aceRevision, aceFlags, accessMask T.DWORD,
		sid *T.SID) T.BOOL

	AddAccessDeniedObjectAce func(
		acl *T.ACL,
		aceRevision, aceFlags, accessMask T.DWORD,
		objectTypeGuid, inheritedObjectTypeGuid *T.GUID,
		sid *T.SID) T.BOOL

	AddAce func(
		acl *T.ACL,
		aceRevision, startingAceIndex T.DWORD,
		aceList *T.VOID,
		aceListLength T.DWORD) T.BOOL

	AddAtom func(string_ VString) T.ATOM

	AddAuditAccessAce func(
		acl *T.ACL,
		aceRevision, accessMask T.DWORD,
		sid *T.SID,
		auditSuccess, auditFailure T.BOOL) T.BOOL

	AddAuditAccessAceEx func(
		acl *T.ACL,
		aceRevision, aceFlags,
		accessMask T.DWORD,
		sid *T.SID,
		auditSuccess, auditFailure T.BOOL) T.BOOL

	AddAuditAccessObjectAce func(
		acl *T.ACL,
		aceRevision, aceFlags, accessMask,
		objectTypeGuid, inheritedObjectTypeGuid *T.GUID,
		sid *T.SID,
		auditSuccess, auditFailure T.BOOL) T.BOOL

	AddRefActCtx func(actCtx T.HANDLE)

	AddVectoredContinueHandler func(
		first T.ULONG,
		handler T.VECTORED_EXCEPTION_HANDLER) *T.VOID

	AddVectoredExceptionHandler func(
		first T.ULONG,
		handler T.VECTORED_EXCEPTION_HANDLER) *T.VOID

	AdjustTokenGroups func(
		tokenHandle T.HANDLE,
		resetToDefault T.BOOL,
		newState *T.TOKEN_GROUPS,
		bufferLength T.DWORD,
		previousState *T.TOKEN_GROUPS,
		returnLength *T.DWORD) T.BOOL

	AdjustTokenPrivileges func(
		tokenHandle T.HANDLE,
		disableAllPrivileges T.BOOL,
		newState *T.TOKEN_PRIVILEGES,
		bufferLength T.DWORD,
		previousState *T.TOKEN_PRIVILEGES,
		returnLength *T.DWORD) T.BOOL

	AllocateAndInitializeSid func(
		identifierAuthority *T.SID_IDENTIFIER_AUTHORITY,
		subAuthorityCount T.BYTE,
		subAuthority0, subAuthority1,
		subAuthority2, subAuthority3,
		subAuthority4, subAuthority5,
		subAuthority6, subAuthority7 T.DWORD,
		sid **T.SID) T.BOOL

	AllocateLocallyUniqueId func(luid *T.LUID) T.BOOL

	AllocateUserPhysicalPages func(
		process T.HANDLE,
		numberOfPages, pageArray *T.ULONG_PTR) T.BOOL

	AreAllAccessesGranted func(
		grantedAccess, desiredAccess T.DWORD) T.BOOL

	AreAnyAccessesGranted func(
		grantedAccess, desiredAccess T.DWORD) T.BOOL

	AreFileApisANSI func()

	AssignProcessToJobObject func(job, process T.HANDLE) T.BOOL

	BackupEventLog func(
		eventLog T.HANDLE, backupFileName VString) T.BOOL

	BackupRead func(
		file T.HANDLE,
		buffer *T.BYTE,
		numberOfBytesToRead T.DWORD,
		numberOfBytesRead *T.DWORD,
		abort, processSecurity T.BOOL,
		context **T.VOID) T.BOOL

	BackupSeek func(
		file T.HANDLE,
		lowBytesToSeek, highBytesToSeek T.DWORD,
		lowByteSeeked, highByteSeeked *T.DWORD,
		context **T.VOID) T.BOOL

	BackupWrite func(
		file T.HANDLE,
		buffer *T.BYTE,
		numberOfBytesToWrite T.DWORD,
		numberOfBytesWritten *T.DWORD,
		abort, processSecurity T.BOOL,
		context **T.VOID) T.BOOL

	Beep func(freq, duration T.DWORD) T.BOOL

	BeginUpdateResource func(
		fileName VString,
		deleteExistingResources T.BOOL) T.HANDLE

	BindIoCompletionCallback func(
		fileHandle T.HANDLE,
		function T.OVERLAPPED_COMPLETION_ROUTINE,
		flags T.ULONG) T.BOOL

	BuildCommDCB func(def VString, dcb *T.DCB) T.BOOL

	BuildCommDCBAndTimeouts func(
		def VString,
		dcb *T.DCB,
		commTimeouts *T.COMMTIMEOUTS) T.BOOL

	CallNamedPipe func(
		namedPipeName VString,
		inBuffer *T.VOID,
		inBufferSize T.DWORD,
		outBuffer *T.VOID,
		outBufferSize T.DWORD,
		bytesRead *T.DWORD,
		timeOut T.DWORD) T.BOOL

	CancelDeviceWakeupRequest func(device T.HANDLE) T.BOOL

	CancelIo func(file T.HANDLE) T.BOOL

	CancelTimerQueueTimer func(
		timerQueue, timer T.HANDLE) T.BOOL

	CancelWaitableTimer func(timer T.HANDLE) T.BOOL

	ChangeTimerQueueTimer func(
		timerQueue, timer T.HANDLE,
		dueTime, period T.ULONG) T.BOOL

	CheckNameLegalDOS8Dot3 func(
		name VString,
		oemName T.OAString,
		oemNameSize T.DWORD,
		nameContainsSpaces, nameLegal *T.BOOL) T.BOOL

	CheckRemoteDebuggerPresent func(
		process T.HANDLE, debuggerPresent *T.BOOL) T.BOOL

	CheckTokenMembership func(
		tokenHandleHANDLE,
		sidToCheck *T.SID,
		isMember *T.BOOL) T.BOOL

	ClearCommBreak func(file T.HANDLE) T.BOOL

	ClearCommError func(
		file T.HANDLE, errors *T.DWORD, stat *T.COMSTAT) T.BOOL

	ClearEventLog func(
		eventLog T.HANDLE, backupFileName VString) T.BOOL

	CloseEncryptedFileRaw func(context *T.VOID)

	CloseEventLog func(eventLog T.HANDLE) T.BOOL

	CloseHandle func(object T.HANDLE) T.BOOL

	CommConfigDialog func(
		name VString, wnd T.HWND, cc *T.COMMCONFIG) T.BOOL

	CompareFileTime func(
		fileTime1, fileTime2 *T.FILETIME) T.LONG

	ConnectNamedPipe func(
		namedPipe T.HANDLE, overlapped *T.OVERLAPPED) T.BOOL

	ContinueDebugEvent func(
		processId, threadId, continueStatus T.DWORD) T.BOOL

	ConvertFiberToThread func() T.BOOL

	ConvertThreadToFiber func(parameter *T.VOID) *T.VOID

	ConvertThreadToFiberEx func(
		parameter *T.VOID, flags T.DWORD) *T.VOID

	ConvertToAutoInheritPrivateObjectSecurity func(
		parentDescriptor,
		currentSecurityDescriptor *T.SECURITY_DESCRIPTOR,
		newSecurityDescriptor **T.SECURITY_DESCRIPTOR,
		objectType *T.GUID,
		isDirectoryObject T.BOOLEAN,
		genericMapping *T.GENERIC_MAPPING) T.BOOL

	CopyFile func(
		existingFileName, newFileName VString,
		failIfExists T.BOOL) T.BOOL

	CopyFileEx func(
		existingFileName, newFileName VString,
		progressRoutine T.PROGRESS_ROUTINE,
		data *T.VOID,
		cancel *T.BOOL,
		copyFlags T.DWORD) T.BOOL

	CopySid func(
		destinationLength T.DWORD,
		destination, source *T.SID) T.BOOL

	CreateActCtx func(actCtx T.ACTCTX) T.HANDLE

	CreateDirectory func(
		pathName VString,
		securityAttributes *T.SECURITY_ATTRIBUTES) T.BOOL

	CreateDirectoryEx func(
		templateDirectory, newDirectory VString,
		securityAttributes *T.SECURITY_ATTRIBUTES) T.BOOL

	CreateEvent func(
		eventAttributes *T.SECURITY_ATTRIBUTES,
		manualReset, initialState T.BOOL,
		name VString) T.HANDLE

	CreateFiber func(
		stackSize T.SIZE_T,
		startAddress T.FIBER_START_ROUTINE,
		parameter *T.VOID) *T.VOID

	CreateFiberEx func(
		stackCommitSize, stackReserveSize T.SIZE_T,
		flags T.DWORD,
		startAddress T.FIBER_START_ROUTINE,
		parameter *T.VOID) *T.VOID

	CreateFile func(
		fileName VString,
		desiredAccess, shareMode T.DWORD,
		securityAttributes *T.SECURITY_ATTRIBUTES,
		creationDisposition, flagsAndAttributes T.DWORD,
		templateFile T.HANDLE) T.HANDLE

	CreateFileMapping func(
		file T.HANDLE,
		fileMappingAttributes *T.SECURITY_ATTRIBUTES,
		protect, maximumSizeHigh, maximumSizeLow T.DWORD,
		name VString) T.HANDLE

	CreateHardLink func(
		fileName, existingFileName VString,
		securityAttributes *T.SECURITY_ATTRIBUTES) T.BOOL

	CreateIoCompletionPort func(
		fileHandle, existingCompletionPort T.HANDLE,
		completionKey T.ULONG_PTR,
		numberOfConcurrentThreads T.DWORD) T.HANDLE

	CreateJobObject func(
		jobAttributes *T.SECURITY_ATTRIBUTES,
		name VString) T.HANDLE

	CreateJobSet func(
		numJob T.ULONG,
		userJobSet *T.JOB_SET_ARRAY,
		flags T.ULONG) T.BOOL

	CreateMailslot func(
		name VString,
		maxMessageSize, readTimeout T.DWORD,
		securityAttributes *T.SECURITY_ATTRIBUTES) T.HANDLE

	CreateMemoryResourceNotification func(
		T.MEMORY_RESOURCE_NOTIFICATION_TYPE) T.HANDLE

	CreateMutex func(
		mutexAttributes *T.SECURITY_ATTRIBUTES,
		initialOwner T.BOOL,
		name VString) T.HANDLE

	CreateNamedPipe func(
		name VString,
		openMode, pipeMode, maxInstances,
		outBufferSize, inBufferSize, defaultTimeOut T.DWORD,
		securityAttributes *T.SECURITY_ATTRIBUTES) T.HANDLE

	CreatePipe func(
		readPipe, writePipe *T.HANDLE,
		pipeAttributes *T.SECURITY_ATTRIBUTES,
		size T.DWORD) T.BOOL

	CreatePrivateObjectSecurity func(
		parentDescriptor,
		creatorDescriptor *T.SECURITY_DESCRIPTOR,
		newDescriptor **T.SECURITY_DESCRIPTOR,
		isDirectoryObject T.BOOL,
		token T.HANDLE,
		genericMapping *T.GENERIC_MAPPING) T.BOOL

	CreatePrivateObjectSecurityEx func(
		parentDescriptor,
		creatorDescriptor *T.SECURITY_DESCRIPTOR,
		newDescriptor **T.SECURITY_DESCRIPTOR,
		objectType *T.GUID,
		isContainerObject T.BOOL,
		autoInheritFlags T.ULONG,
		token T.HANDLE,
		genericMapping *T.GENERIC_MAPPING) T.BOOL

	CreatePrivateObjectSecurityWithMultipleInheritance func(
		parentDescriptor, creatorDescriptor,
		newDescriptor **T.SECURITY_DESCRIPTOR,
		objectTypes **T.GUID,
		guidCount T.ULONG,
		isContainerObject T.BOOL,
		autoInheritFlags T.ULONG,
		token T.HANDLE,
		genericMapping *T.GENERIC_MAPPING) T.BOOL

	CreateProcess func(
		applicationName, commandLine VString,
		processAttributes,
		threadAttributes *T.SECURITY_ATTRIBUTES,
		inheritHandles T.BOOL,
		creationFlags T.DWORD,
		environment *T.VOID,
		currentDirectory VString,
		startupInfo *T.STARTUPINFO,
		processInformation *T.PROCESS_INFORMATION) T.BOOL

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
		processInformation *T.PROCESS_INFORMATION) T.BOOL

	CreateProcessWithLogonW func(
		username, domain, password T.WString,
		logonFlags T.DWORD,
		applicationName T.WString,
		commandLine T.OWString,
		creationFlags T.DWORD,
		environment *T.VOID,
		currentDirectory T.WString,
		startupInfo *T.STARTUPINFO,
		processInformation *T.PROCESS_INFORMATION) T.BOOL

	CreateProcessWithTokenW func(
		token T.HANDLE,
		logonFlags T.DWORD,
		applicationName T.WString,
		commandLine T.OWString,
		creationFlags T.DWORD,
		environment *T.VOID,
		currentDirectory T.OWString,
		startupInfo *T.STARTUPINFO,
		processInformation *T.PROCESS_INFORMATION) T.BOOL

	CreateRemoteThread func(
		process T.HANDLE,
		threadAttributes *T.SECURITY_ATTRIBUTES,
		stackSize T.SIZE_T,
		startAddress T.THREAD_START_ROUTINE,
		parameter *T.VOID,
		creationFlags T.DWORD,
		threadId *T.DWORD) T.HANDLE

	CreateRestrictedToken func(
		existingTokenHandle T.HANDLE,
		flags, disableSidCount, deletePrivilegeCount,
		restrictedSidCount T.DWORD,
		newTokenHandle *T.HANDLE) T.BOOL

	CreateSemaphore func(
		semaphoreAttributes *T.SECURITY_ATTRIBUTES,
		initialCount, maximumCount T.LONG,
		name VString) T.HANDLE

	CreateTapePartition func(
		device T.HANDLE,
		partitionMethod, count, size T.DWORD) T.DWORD

	CreateThread func(
		threadAttributes *T.SECURITY_ATTRIBUTES,
		stackSize T.SIZE_T,
		startAddress T.THREAD_START_ROUTINE,
		parameter *T.VOID,
		creationFlags T.DWORD,
		threadId *T.DWORD) T.HANDLE

	CreateTimerQueue func() T.HANDLE

	CreateTimerQueueTimer func(
		newTimer *T.HANDLE,
		timerQueue T.HANDLE,
		callback T.WAITORTIMERCALLBACK,
		parameter *T.VOID,
		dueTime, period T.DWORD,
		flags T.ULONG) T.BOOL

	CreateWaitableTimer func(
		timerAttributes *T.SECURITY_ATTRIBUTES,
		manualReset T.BOOL,
		timerName VString) T.HANDLE

	CreateWellKnownSid func(
		wellKnownSidType T.WELL_KNOWN_SID_TYPE,
		domainSid, sid *T.SID,
		sidl *T.DWORD) T.BOOL

	DeactivateActCtx func(
		flags T.DWORD, cookie T.ULONG_PTR) T.BOOL

	DebugActiveProcess func(processId T.DWORD) T.BOOL

	DebugActiveProcessStop func(processId T.DWORD) T.BOOL

	DebugBreak func()

	DebugBreakProcess func(process T.HANDLE) T.BOOL

	DebugSetProcessKillOnExit func(killOnExit T.BOOL) T.BOOL

	DecodePointer func(ptr *T.VOID) *T.VOID

	DecodeSystemPointer func(ptr *T.VOID) *T.VOID

	DecryptFile func(fileName VString, reserved T.DWORD) T.BOOL

	DefineDosDevice func(
		flags T.DWORD,
		deviceName, targetPath VString) T.BOOL

	DeleteAce func(acl *T.ACL, aceIndex T.DWORD) T.BOOL

	DeleteAtom func(atom T.ATOM) T.ATOM

	DeleteCriticalSection func(*T.CRITICAL_SECTION)

	DeleteFiber func(fiber *T.VOID)

	DeleteFile func(fileName VString) T.BOOL

	DeleteTimerQueue func(timerQueue T.HANDLE) T.BOOL

	DeleteTimerQueueEx func(
		timerQueue, completionEvent T.HANDLE) T.BOOL

	DeleteTimerQueueTimer func(
		timerQueue, timer, completionEvent T.HANDLE) T.BOOL

	DeleteVolumeMountPoint func(
		volumeMountPoint VString) T.BOOL

	DeregisterEventSource func(eventLog T.HANDLE) T.BOOL

	DestroyPrivateObjectSecurity func(
		objectDescriptor **T.SECURITY_DESCRIPTOR) T.BOOL

	DeviceIoControl func(
		device T.HANDLE,
		ioControlCode T.DWORD,
		inBuffer *T.VOID,
		inBufferSize T.DWORD,
		outBuffer *T.VOID,
		outBufferSize, bytesReturned *T.DWORD,
		overlapped *T.OVERLAPPED) T.BOOL

	DisableThreadLibraryCalls func(libModule T.HMODULE) T.BOOL

	DisconnectNamedPipe func(namedPipe T.HANDLE) T.BOOL

	DnsHostnameToComputerName func(
		hostname VString,
		computerName OVString,
		size *T.DWORD) T.BOOL

	DosDateTimeToFileTime func(
		fatDate, fatTime T.WORD, fileTime *T.FILETIME) T.BOOL

	DuplicateHandle func(
		sourceProcessHandle, sourceHandle,
		targetProcessHandle T.HANDLE,
		targetHandle *T.HANDLE,
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		options T.DWORD) T.BOOL

	DuplicateToken func(
		existingTokenHandle T.HANDLE,
		impersonationLevel T.SECURITY_IMPERSONATION_LEVEL,
		duplicateTokenHandle *T.HANDLE) T.BOOL

	DuplicateTokenEx func(
		existingToken T.HANDLE,
		desiredAccess T.DWORD,
		tokenAttributes *T.SECURITY_ATTRIBUTES,
		impersonationLevel T.SECURITY_IMPERSONATION_LEVEL,
		tokenType T.TOKEN_TYPE,
		newToken *T.HANDLE) T.BOOL

	EncodePointer func(ptr *T.VOID) *T.VOID

	EncodeSystemPointer func(ptr *T.VOID) *T.VOID

	EncryptFile func(fileName VString) T.BOOL

	EndUpdateResource func(
		update T.HANDLE, discard T.BOOL) T.BOOL

	EnterCriticalSection func(*T.CRITICAL_SECTION)

	EnumResourceLanguages func(
		module T.HMODULE,
		type_, name VString,
		enumFunc T.ENUMRESLANGPROC,
		param T.LONG_PTR) T.BOOL

	EnumResourceNames func(
		module T.HMODULE,
		type_ VString,
		enumFunc T.ENUMRESNAMEPROC,
		param T.LONG_PTR) T.BOOL

	EnumResourceTypes func(
		module T.HMODULE,
		enumFunc T.ENUMRESTYPEPROC,
		param T.LONG_PTR) T.BOOL

	EnumSystemFirmwareTables func(
		firmwareTableProviderSignature T.DWORD,
		firmwareTableEnumBuffer *T.VOID,
		bufferSize T.DWORD) T.UINT

	EqualDomainSid func(
		sid1, sid2 *T.SID, equal *T.BOOL) T.BOOL

	EqualPrefixSid func(
		sid1, sid2 *T.SID) T.BOOL

	EqualSid func(sid1, sid2 *T.SID) T.BOOL

	EraseTape func(
		device T.HANDLE,
		eraseType T.DWORD,
		immediate T.BOOL) T.DWORD

	EscapeCommFunction func(file T.HANDLE, function T.DWORD) T.BOOL

	ExitProcess func(exitCode T.UINT)

	ExitThread func(exitCode T.DWORD)

	ExpandEnvironmentStrings func(
		src, dst VString, size T.DWORD) T.DWORD

	FatalAppExit func(action T.UINT, messageText VString)

	FatalExit func(exitCode int)

	FileEncryptionStatus func(
		fileName VString, status *T.DWORD) T.BOOL

	FileTimeToDosDateTime func(
		fileTime *T.FILETIME, fatDate, fatTime *T.WORD) T.BOOL

	FileTimeToLocalFileTime func(
		fileTime, localFileTime *T.FILETIME) T.BOOL

	FileTimeToSystemTime func(
		fileTime, systemTime *T.SYSTEMTIME) T.BOOL

	FindActCtxSectionGuid func(
		flags T.DWORD,
		extensionGuid *T.GUID,
		sectionId T.ULONG,
		guidToFind *T.GUID,
		returnedData *T.ACTCTX_SECTION_KEYED_DATA) T.BOOL

	FindActCtxSectionString func(
		flags T.DWORD,
		extensionGuid *T.GUID,
		sectionId T.ULONG,
		string_ToFind VString,
		returnedData *T.ACTCTX_SECTION_KEYED_DATA) T.BOOL

	FindAtom func(string_ VString) T.ATOM

	FindClose func(findFile T.HANDLE) T.BOOL

	FindCloseChangeNotification func(
		changeHandle T.HANDLE) T.BOOL

	FindFirstChangeNotification func(
		pathName VString,
		watchSubtree T.BOOL,
		notifyFilter T.DWORD) T.HANDLE

	FindFirstFileA func(
		fileName T.AString,
		findFileData *T.WIN32_FIND_DATAA) T.HANDLE

	FindFirstFileW func(
		fileName T.WString,
		findFileData *T.WIN32_FIND_DATAW) T.HANDLE

	FindFirstFileEx func(
		fileName VString,
		infoLevelId T.FINDEX_INFO_LEVELS,
		findFileData *T.VOID,
		searchOp T.FINDEX_SEARCH_OPS,
		searchFilter *T.VOID,
		additionalFlags T.DWORD) T.HANDLE

	FindFirstFreeAce func(acl *T.ACL, ace **T.VOID) T.BOOL

	FindFirstVolume func(
		volumeName OVString, bufferLength T.DWORD) T.HANDLE

	FindFirstVolumeMountPoint func(
		rootPathName VString,
		volumeMountPoint OVString,
		bufferLength T.DWORD) T.HANDLE

	FindFirstStreamW func(
		fileName T.WString,
		infoLevel T.STREAM_INFO_LEVELS,
		findStreamData *T.VOID,
		flags T.DWORD) T.HANDLE

	FindNextStreamW func(
		findStream T.HANDLE, data *T.VOID) T.BOOL

	FindNextChangeNotification func(
		changeHandle T.HANDLE) T.BOOL

	FindNextFileA func(
		findFile T.HANDLE, data *T.WIN32_FIND_DATAA) T.BOOL

	FindNextFileW func(
		findFile T.HANDLE, data *T.WIN32_FIND_DATAW) T.BOOL

	FindNextVolume func(
		findVolume T.HANDLE,
		volumeName OVString,
		bufferLength T.DWORD) T.BOOL

	FindNextVolumeMountPoint func(
		findVolumeMountPoint T.HANDLE,
		volumeMountPoint OVString,
		bufferLength T.DWORD) T.BOOL

	FindResource func(
		module T.HMODULE,
		name, typ VString) T.HRSRC

	FindResourceEx func(
		module T.HMODULE,
		typ, name VString,
		language T.WORD) T.HRSRC

	FindVolumeClose func(findVolume T.HANDLE) T.BOOL

	FindVolumeMountPointClose func(
		findVolumeMountPoint T.HANDLE) T.BOOL

	FlsAlloc func(callback T.FLS_CALLBACK_FUNCTION) T.DWORD

	FlsFree func(flsIndex T.DWORD) T.BOOL

	FlsGetValue func(flsIndex T.DWORD) *T.VOID

	FlsSetValue func(flsIndex T.DWORD, flsData *T.VOID) T.BOOL

	FlushFileBuffers func(file T.HANDLE) T.BOOL

	FlushInstructionCache func(
		process T.HANDLE, baseAddress *T.VOID, size T.SIZE_T) T.BOOL

	FlushViewOfFile func(
		baseAddress *T.VOID, numberOfBytesToFlush T.SIZE_T) T.BOOL

	FormatMessage func(
		flags T.DWORD,
		source *T.VOID,
		messageId, languageId T.DWORD,
		buffer VString,
		size T.DWORD,
		arguments *va_list) T.DWORD

	FreeEnvironmentStrings func(VString) T.BOOL

	FreeLibrary func(libModule T.HMODULE) (T.BOOL, error)

	FreeLibraryAndExitThread func(
		libModule T.HMODULE, exitCode T.DWORD)

	FreeResource func(resData T.HGLOBAL) T.BOOL

	FreeSid func(sid *T.SID) *T.VOID

	FreeUserPhysicalPages func(
		process T.HANDLE,
		numberOfPages, pageArray *T.ULONG_PTR) T.BOOL

	GetAce func(
		acl *T.ACL, aceIndex T.DWORD, ace **T.VOID) T.BOOL

	GetAclInformation func(
		acl *T.ACL,
		info *T.VOID,
		length T.DWORD,
		class T.ACL_INFORMATION_CLASS) T.BOOL

	GetAtomName func(
		atom T.ATOM, buffer VString, size int) T.UINT

	GetBinaryType func(
		applicationName VString, binaryType *T.DWORD) T.BOOL

	GetCommandLine func() VString

	GetCommConfig func(
		commDev T.HANDLE, cc *T.COMMCONFIG, size *T.DWORD) T.BOOL

	GetCommMask func(file T.HANDLE, evtMask *T.DWORD) T.BOOL

	GetCommModemStatus func(
		file T.HANDLE, modemStat *T.DWORD) T.BOOL

	GetCommProperties func(
		file T.HANDLE, cp *T.COMMPROP) T.BOOL

	GetCommState func(file T.HANDLE, dcn *T.DCB) T.BOOL

	GetCommTimeouts func(
		file T.HANDLE, commTimeouts *T.COMMTIMEOUTS) T.BOOL

	GetCompressedFileSize func(
		fileName VString, fileSizeHigh *T.DWORD) T.DWORD

	GetComputerName func(
		buffer OVString, size *T.DWORD) T.BOOL

	GetComputerNameEx func(
		nameType T.COMPUTER_NAME_FORMAT,
		buffer OVString,
		size *T.DWORD) T.BOOL

	GetCurrentActCtx func(actCtx *T.HANDLE) T.BOOL

	GetCurrentDirectory func(
		bufferLength T.DWORD, buffer VString) T.DWORD

	GetCurrentHwProfileA func(
		hwProfileInfo *T.HW_PROFILE_INFOA) T.BOOL

	GetCurrentHwProfileW func(
		hwProfileInfo *T.HW_PROFILE_INFOW) T.BOOL

	GetCurrentProcess func() T.HANDLE

	GetCurrentProcessId func() T.DWORD

	GetCurrentProcessorNumber func() T.DWORD

	GetCurrentThread func() T.HANDLE

	GetCurrentThreadId func() T.DWORD

	GetDefaultCommConfig func(
		name VString, cc *T.COMMCONFIG, size *T.DWORD) T.BOOL

	GetDevicePowerState func(
		device T.HANDLE, on *T.BOOL) T.BOOL

	GetDiskFreeSpace func(
		rootPathName VString,
		sectorsPerCluster,
		bytesPerSector,
		numberOfFreeClusters,
		totalNumberOfClusters *T.DWORD) T.BOOL

	GetDiskFreeSpaceEx func(
		directoryName VString,
		freeBytesAvailableToCaller,
		totalNumberOfBytes,
		totalNumberOfFreeBytes *T.ULARGE_INTEGER) T.BOOL

	GetDllDirectory func(
		bufferLength T.DWORD, buffer VString) T.DWORD

	GetDriveType func(rootPathName VString) T.UINT

	GetEnvironmentStrings func() T.OMVString

	GetEnvironmentVariable func(
		name, buffer VString, size T.DWORD) T.DWORD

	GetEventLogInformation func(
		eventLog T.HANDLE,
		infoLevel T.DWORD,
		buffer *T.VOID,
		bufSize T.DWORD,
		bytesNeeded *T.DWORD) T.BOOL

	GetExitCodeProcess func(process T.HANDLE, exitCode *T.DWORD) T.BOOL

	GetExitCodeThread func(thread T.HANDLE, exitCode *T.DWORD) T.BOOL

	GetFileAttributes func(fileName VString) T.DWORD

	GetFileAttributesEx func(
		fileName VString,
		infoLevelId T.GET_FILEEX_INFO_LEVELS,
		fileInformation *T.VOID) T.BOOL

	GetFileInformationByHandle func(
		file T.HANDLE,
		fileInformation *T.BY_HANDLE_FILE_INFORMATION) T.BOOL

	GetFileSecurity func(
		fileName VString,
		requestedInformation T.SECURITY_INFORMATION,
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		length T.DWORD,
		lengthNeeded *T.DWORD) T.BOOL

	GetFileSize func(
		file T.HANDLE, fileSizeHigh *T.DWORD) T.DWORD

	GetFileSizeEx func(
		file T.HANDLE, fileSize *T.LARGE_INTEGER) T.BOOL

	GetFileTime func(
		file T.HANDLE,
		creationTime, lastAccessTime,
		lastWriteTime *T.FILETIME) T.BOOL

	GetFileType func(file T.HANDLE) T.DWORD

	GetFirmwareEnvironmentVariable func(
		name, guid VString, buffer *T.VOID, size T.DWORD) T.DWORD

	GetFullPathName func(
		fileName VString,
		bufferLength T.DWORD,
		buffer VString,
		filePart OVString) T.DWORD

	GetHandleInformation func(
		object T.HANDLE, flags *T.DWORD) T.BOOL

	GetKernelObjectSecurity func(
		handle T.HANDLE,
		requestedInformation T.SECURITY_INFORMATION,
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		length T.DWORD,
		lengthNeeded *T.DWORD) T.BOOL

	GetLargePageMinimum func() T.SIZE_T

	GetLastError func() T.DWORD

	GetLengthSid func(sid *T.SID) T.DWORD

	GetLocalTime func(systemTime *T.SYSTEMTIME)

	GetLogicalDrives func() T.DWORD

	GetLogicalDriveStrings func(
		bufferLength T.DWORD,
		buffer T.OMVString) T.DWORD

	GetLogicalProcessorInformation func(
		buffer *T.SYSTEM_LOGICAL_PROCESSOR_INFORMATION,
		returnedLength *T.DWORD) T.BOOL

	GetLongPathName func(
		shortPath VString,
		longPath OVString,
		buffer T.DWORD) T.DWORD

	GetMailslotInfo func(
		mailslot T.HANDLE,
		maxMessageSize, nextSize, messageCount,
		readTimeout *T.DWORD) T.BOOL

	GetModuleFileName func(
		module T.HMODULE, filename VString, size T.DWORD) T.DWORD

	GetModuleHandle func(moduleName VString) T.HMODULE

	GetModuleHandleEx func(
		flags T.DWORD,
		moduleName VString,
		module *T.HMODULE) T.BOOL

	GetNamedPipeHandleState func(
		namedPipe T.HANDLE,
		state, curInstances, maxCollectionCount,
		collectDataTimeout *T.DWORD,
		userName VString,
		maxUserNameSize T.DWORD) T.BOOL

	GetNamedPipeInfo func(
		namedPipe T.HANDLE,
		flags, outBufferSize, inBufferSize,
		maxInstances *T.DWORD) T.BOOL

	GetNativeSystemInfo func(systemInfo *T.SYSTEM_INFO)

	GetNumaAvailableMemoryNode func(
		node T.WChar,
		availableBytes *T.ULONGLONG) T.BOOL

	GetNumaHighestNodeNumber func(
		highestNodeNumber *T.ULONG) T.BOOL

	GetNumaNodeProcessorMask func(
		node T.WChar,
		processorMask *T.ULONGLONG) T.BOOL

	GetNumaProcessorNode func(
		processor T.WChar,
		nodeNumber *T.WChar) T.BOOL

	GetNumberOfEventLogRecords func(
		eventLog T.HANDLE,
		numberOfRecords *T.DWORD) T.BOOL

	GetOldestEventLogRecord func(
		eventLog T.HANDLE,
		oldestRecord *T.DWORD) T.BOOL

	GetOverlappedResult func(
		file T.HANDLE,
		overlapped *T.OVERLAPPED,
		numberOfBytesTransferred *T.DWORD,
		wait T.BOOL) T.BOOL

	GetPriorityClass func(process T.HANDLE) T.DWORD

	GetPrivateObjectSecurity func(
		objectDescriptor *T.SECURITY_DESCRIPTOR,
		securityInformation T.SECURITY_INFORMATION,
		resultantDescriptor *T.SECURITY_DESCRIPTOR,
		descriptorLength T.DWORD,
		returnLength *T.DWORD) T.BOOL

	GetPrivateProfileInt func(
		appName, keyName VString,
		default_ T.INT,
		fileName VString) T.UINT

	GetPrivateProfileSection func(
		appName, returnedString VString,
		size T.DWORD,
		fileName VString) T.DWORD

	GetPrivateProfileSectionNames func(
		returnBuffer VString,
		size T.DWORD,
		fileName VString) T.DWORD

	GetPrivateProfileString func(
		appName, keyName, default_, returned VString,
		size T.DWORD,
		fileName VString) T.DWORD

	GetPrivateProfileStruct func(
		section, key VString,
		struct_ *T.VOID,
		sizeStruct T.UINT,
		file VString) T.BOOL

	GetProcAddress func(
		module T.HMODULE, procName T.AString) T.FARPROC

	GetProcessAffinityMask func(
		process T.HANDLE,
		processAffinityMask,
		systemAffinityMask *T.DWORD_PTR) T.BOOL

	GetProcessHandleCount func(
		process T.HANDLE, handleCount *T.DWORD) T.BOOL

	GetProcessHeap func() T.HANDLE

	GetProcessHeaps func(
		numberOfHeaps T.DWORD,
		processHeaps *T.HANDLE) T.DWORD

	GetProcessId func(process T.HANDLE) T.DWORD

	GetProcessIdOfThread func(thread T.HANDLE) T.DWORD

	GetProcessIoCounters func(
		process T.HANDLE, ioCounters *T.IO_COUNTERS) T.BOOL

	GetProcessPriorityBoost func(
		process T.HANDLE, disablePriorityBoost *T.BOOL) T.BOOL

	GetProcessShutdownParameters func(
		level, flags *T.DWORD) T.BOOL

	GetProcessTimes func(
		process T.HANDLE,
		creationTime, exitTime, kernelTime,
		userTime *T.FILETIME) T.BOOL

	GetProcessVersion func(
		processId T.DWORD) T.DWORD

	GetProcessWorkingSetSize func(
		process T.HANDLE,
		minWorkingSetSize, maxWorkingSetSize *T.SIZE_T) T.BOOL

	GetProcessWorkingSetSizeEx func(
		process T.HANDLE,
		minWorkingSetSize, maxWorkingSetSize *T.SIZE_T,
		flags *T.DWORD) T.BOOL

	GetProfileInt func(
		appName, keyName VString,
		default_ T.INT) T.UINT

	GetProfileSection func(
		appName, returnedString VString,
		size T.DWORD) T.DWORD

	GetProfileString func(
		appName, keyName, default_, returned VString,
		size T.DWORD) T.DWORD

	GetQueuedCompletionStatus func(
		completionPort T.HANDLE,
		numberOfBytesTransferred *T.DWORD,
		completionKey *T.ULONG_PTR,
		overlapped **T.OVERLAPPED,
		milliseconds T.DWORD) T.BOOL

	GetSecurityDescriptorControl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		control *T.SECURITY_DESCRIPTOR_CONTROL,
		revision *T.DWORD) T.BOOL

	GetSecurityDescriptorDacl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		daclPresent *T.BOOL,
		dacl **T.ACL,
		daclDefaulted *T.BOOL) T.BOOL

	GetSecurityDescriptorGroup func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		group **T.SID,
		groupDefaulted *T.BOOL) T.BOOL

	GetSecurityDescriptorLength func(
		securityDescriptor *T.SECURITY_DESCRIPTOR) T.DWORD

	GetSecurityDescriptorOwner func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		owner **T.SID,
		ownerDefaulted *T.BOOL) T.BOOL

	GetSecurityDescriptorRMControl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		rmControl *T.WChar) T.DWORD

	GetSecurityDescriptorSacl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		saclPresent *T.BOOL,
		sacl **T.ACL,
		saclDefaulted *T.BOOL) T.BOOL

	GetShortPathName func(
		longPath VString,
		shortPath OVString,
		buffer T.DWORD) T.DWORD

	GetSidIdentifierAuthority func(
		sid *T.SID) *T.SID_IDENTIFIER_AUTHORITY

	GetSidLengthRequired func(subAuthorityCount T.WChar) T.DWORD

	GetSidSubAuthority func(
		sid *T.SID, subAuthority T.DWORD) *T.DWORD

	GetSidSubAuthorityCount func(sid *T.SID) *T.WChar

	GetStartupInfo func(startupInfo *T.STARTUPINFO)

	GetStdHandle func(stdHandle T.DWORD) T.HANDLE

	GetSystemDirectory func(buffer VString, size T.UINT) T.UINT

	GetSystemFileCacheSize func(
		minFileCacheSize, maxFileCacheSize *T.SIZE_T,
		flags *T.DWORD) T.BOOL

	GetSystemFirmwareTable func(
		firmwareTableProviderSignature,
		firmwareTableID T.DWORD,
		firmwareTableBuffer *T.VOID,
		bufferSize T.DWORD) T.UINT

	GetSystemInfo func(systemInfo *T.SYSTEM_INFO)

	GetSystemPowerStatus func(
		systemPowerStatus *T.SYSTEM_POWER_STATUS) T.BOOL

	GetSystemRegistryQuota func(
		quotaAllowed, quotaUsed *T.DWORD) T.BOOL

	GetSystemTime func(
		systemTime *T.SYSTEMTIME)

	GetSystemTimeAdjustment func(
		timeAdjustment, timeIncrement *T.DWORD,
		timeAdjustmentDisabled *T.BOOL) T.BOOL

	GetSystemTimeAsFileTime func(
		systemTimeAsFileTime *T.FILETIME)

	GetSystemTimes func(
		idleTime, kernelTime, userTime *T.FILETIME) T.BOOL

	GetSystemWindowsDirectory func(
		buffer VString, size T.UINT) T.UINT

	GetSystemWow64Directory func(
		buffer VString, size T.UINT) T.UINT

	GetTapeParameters func(
		device T.HANDLE,
		operation T.DWORD,
		size *T.DWORD,
		tapeInformation *T.VOID) T.DWORD

	GetTapePosition func(
		device T.HANDLE,
		positionType T.DWORD,
		partition, offsetLow, offsetHigh *T.DWORD) T.DWORD

	GetTapeStatus func(device T.HANDLE) T.DWORD

	GetTempFileName func(
		pathName, prefixString VString,
		unique T.UINT,
		tempFileName VString) T.UINT

	GetTempPath func(bufferLength T.DWORD, buffer VString) T.DWORD

	GetThreadContext func(thread T.HANDLE, context *T.CONTEXT) T.BOOL

	GetThreadId func(thread T.HANDLE) T.DWORD

	GetThreadIOPendingFlag func(
		thread T.HANDLE, ioIsPending *T.BOOL) T.BOOL

	GetThreadPriority func(thread T.HANDLE) int

	GetThreadPriorityBoost func(
		thread T.HANDLE, disablePriorityBoost *T.BOOL) T.BOOL

	GetThreadSelectorEntry func(
		thread T.HANDLE,
		selector T.DWORD,
		selectorEntry *T.LDT_ENTRY) T.BOOL

	GetThreadTimes func(
		thread T.HANDLE,
		creationTime, exitTime, kernelTime,
		userTime *T.FILETIME) T.BOOL

	GetTickCount func() T.DWORD

	GetTimeZoneInformation func(
		timeZoneInformation *T.TIME_ZONE_INFORMATION) T.DWORD

	GetTokenInformation func(
		tokenHandle T.HANDLE,
		tokenInformationClass T.TOKEN_INFORMATION_CLASS,
		tokenInformation *T.VOID,
		tokenInformationLength T.DWORD,
		returnLength *T.DWORD) T.BOOL

	GetUserName func(buffer OVString, size *T.DWORD) T.BOOL

	GetVersion func() T.DWORD

	GetVersionExA func(versionInformation *T.OSVERSIONINFOA) T.BOOL

	GetVersionExW func(versionInformation *T.OSVERSIONINFOW) T.BOOL

	GetVolumeInformation func(
		rootPathName, volumeNameBuffer VString,
		volumeNameSize T.DWORD,
		volumeSerialNumber, maximumComponentLength,
		fileSystemFlags *T.DWORD,
		fileSystemNameBuffer VString,
		fileSystemNameSize T.DWORD) T.BOOL

	GetVolumeNameForVolumeMountPoint func(
		volumeMountPoint VString,
		volumeName OVString,
		bufferLength T.DWORD) T.BOOL

	GetVolumePathName func(
		fileName VString,
		volumePathName OVString,
		bufferLength T.DWORD) T.BOOL

	GetVolumePathNamesForVolumeName func(
		volumeName VString,
		volumePathNames T.OMVString,
		bufferLength T.DWORD,
		returnLength *T.DWORD) T.BOOL

	GetWindowsAccountDomainSid func(
		sid, domainSid *T.SID, domainSidl *T.DWORD) T.BOOL

	GetWindowsDirectory func(buffer OVString, size T.UINT) T.UINT

	GetWriteWatch func(
		flags T.DWORD,
		baseAddress *T.VOID,
		regionSize T.SIZE_T,
		addresses **T.VOID,
		count *T.ULONG_PTR,
		granularity *T.ULONG) T.UINT

	GlobalAddAtom func(s VString) T.ATOM

	GlobalAlloc func(flags T.UINT, bytes T.SIZE_T) T.HGLOBAL

	GlobalCompact func(minFree T.DWORD) T.SIZE_T

	GlobalDeleteAtom func(a T.ATOM) T.ATOM

	GlobalFindAtom func(s VString) T.ATOM

	GlobalFix func(mem T.HGLOBAL)

	GlobalFlags func(mem T.HGLOBAL) T.UINT

	GlobalFree func(mem T.HGLOBAL) T.HGLOBAL

	GlobalGetAtomName func(
		atom T.ATOM, buffer OVString, size int) T.UINT

	GlobalHandle func(mem *T.VOID) T.HGLOBAL

	GlobalLock func(mem T.HGLOBAL) *T.VOID

	GlobalMemoryStatus func(buffer *T.MEMORYSTATUS)

	GlobalMemoryStatusEx func(buffer *T.MEMORYSTATUSEX) T.BOOL

	GlobalReAlloc func(
		mem T.HGLOBAL, bytes T.SIZE_T, flags T.UINT) T.HGLOBAL

	GlobalSize func(mem T.HGLOBAL) T.SIZE_T

	GlobalUnfix func(mem T.HGLOBAL)

	GlobalUnlock func(mem T.HGLOBAL) T.BOOL

	GlobalUnWire func(mem T.HGLOBAL) T.BOOL

	GlobalWire func(mem T.HGLOBAL) *T.VOID

	HeapAlloc func(heap T.HANDLE, flags T.DWORD, bytes T.SIZE_T) *T.VOID

	HeapCompact func(Heap T.HANDLE, flags T.DWORD) T.SIZE_T

	HeapCreate func(
		options T.DWORD, initialSize, maximumSize T.SIZE_T) T.HANDLE

	HeapDestroy func(Heap T.HANDLE) T.BOOL

	HeapFree func(Heap T.HANDLE, flags T.DWORD, mem *T.VOID) T.BOOL

	HeapLock func(Heap T.HANDLE) T.BOOL

	HeapQueryInformation func(
		h T.HANDLE,
		ic T.HEAP_INFORMATION_CLASS,
		info *T.VOID,
		infoLength T.SIZE_T,
		returnLength *T.SIZE_T) T.BOOL

	HeapReAlloc func(
		heap T.HANDLE, flags T.DWORD, mem *T.VOID, bytes T.SIZE_T) *T.VOID

	HeapSetInformation func(
		h, T.HANDLE,
		ic T.HEAP_INFORMATION_CLASS,
		info *T.VOID,
		infoLength T.SIZE_T) T.BOOL

	HeapSize func(heap T.HANDLE, flags T.DWORD, mem *T.VOID) T.SIZE_T

	HeapUnlock func(Heap T.HANDLE) T.BOOL

	HeapValidate func(heap T.HANDLE, flags T.DWORD, mem *T.VOID) T.BOOL

	HeapWalk func(heap T.HANDLE, entry *T.PROCESS_HEAP_ENTRY) T.BOOL

	Hread func(file T.HFILE, buffer *T.VOID, bytes long) long

	Hwrite func(file T.HFILE, buffer *T.CH, bytes long) long

	ImpersonateAnonymousToken func(threadHandle T.HANDLE) T.BOOL

	ImpersonateLoggedOnUser func(token T.HANDLE) T.BOOL

	ImpersonateNamedPipeClient func(namedPipe T.HANDLE) T.BOOL

	ImpersonateSelf func(
		impersonationLevel T.SECURITY_IMPERSONATION_LEVEL) T.BOOL

	InitAtomTable func(size T.DWORD) T.BOOL

	InitializeAcl func(
		acl *T.ACL, aclLength T.DWORD, aclRevision T.DWORD) T.BOOL

	InitializeCriticalSection func(*T.CRITICAL_SECTION)

	InitializeCriticalSectionAndSpinCount func(
		cs *T.CRITICAL_SECTION, spinCount T.DWORD) T.BOOL

	InitializeSecurityDescriptor func(
		secDesc *T.SECURITY_DESCRIPTOR,
		revision T.DWORD) T.BOOL

	InitializeSid func(
		sid *T.SID,
		idAuth *T.SID_IDENTIFIER_AUTHORITY,
		subAuthorityCount T.BYTE) T.BOOL

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

	IsBadCodePtr func(T.FARPROC) T.BOOL

	IsBadHugeReadPtr func(*T.VOID, T.UINT_PTR) T.BOOL

	IsBadHugeWritePtr func(*T.VOID, T.UINT_PTR) T.BOOL

	IsBadReadPtr func(*T.VOID, T.UINT_PTR) T.BOOL

	IsBadStringPtr func(s VString, max T.UINT_PTR) T.BOOL

	IsBadWritePtr func(*T.VOID, T.UINT_PTR) T.BOOL

	IsDebuggerPresent func() T.BOOL

	IsProcessInJob func(
		processHandle T.HANDLE,
		jobHandle T.HANDLE,
		result *T.BOOL) T.BOOL

	IsProcessorFeaturePresent func(processorFeature T.DWORD) T.BOOL

	IsSystemResumeAutomatic func() T.BOOL

	IsTextUnicode func(v *T.VOID, size int, result *T.INT) T.BOOL

	IsTokenRestricted func(token T.HANDLE) T.BOOL

	IsTokenUntrusted func(token T.HANDLE) T.BOOL

	IsValidAcl func(*T.ACL) T.BOOL

	IsValidSecurityDescriptor func(*T.SECURITY_DESCRIPTOR) T.BOOL

	IsValidSid func(*T.SID) T.BOOL

	IsWellKnownSid func(*T.SID, T.WELL_KNOWN_SID_TYPE) T.BOOL

	IsWow64Process func(process T.HANDLE, wow64Process *T.BOOL) T.BOOL

	Lclose func(file T.HFILE) T.HFILE

	Lcreat func(pathName T.AString, attribute int) T.HFILE

	LeaveCriticalSection func(*T.CRITICAL_SECTION)

	Llseek func(file T.HFILE, offset T.LONG, origin int) T.LONG

	LoadLibrary func(libFileName VString) (T.HMODULE, error)

	LoadLibraryEx func(
		libFileName VString, file T.HANDLE, flags T.DWORD) (T.HMODULE, error)

	LoadModule func(
		moduleName T.AString, parameterBlock *T.VOID) (T.DWORD, error)

	LoadResource func(
		module T.HMODULE, resInfo T.HRSRC) (T.HGLOBAL, error)

	LocalAlloc func(
		flags T.UINT, bytes T.SIZE_T) T.HLOCAL

	LocalCompact func(minFree T.UINT) T.SIZE_T

	LocalFileTimeToFileTime func(
		localFileTime, fileTime *T.FILETIME) T.BOOL

	LocalFlags func(mem T.HLOCAL) T.UINT

	LocalFree func(mem T.HLOCAL) T.HLOCAL

	LocalHandle func(mem *T.VOID) T.HLOCAL

	LocalLock func(mem T.HLOCAL) *T.VOID

	LocalReAlloc func(
		mem T.HLOCAL, bytes T.SIZE_T, flags T.UINT) T.HLOCAL

	LocalShrink func(mem T.HLOCAL, newSize T.UINT) T.SIZE_T

	LocalSize func(mem T.HLOCAL) T.SIZE_T

	LocalUnlock func(mem T.HLOCAL) T.BOOL

	LockFile func(
		file T.HANDLE,
		fileOffsetLow, fileOffsetHigh,
		bytesToLockLow, bytesToLockHigh T.DWORD) T.BOOL

	LockFileEx func(
		file T.HANDLE,
		flags, reserved,
		bytesToLockLow, bytesToLockHigh T.DWORD,
		overlapped *T.OVERLAPPED) T.BOOL

	LockResource func(resData T.HGLOBAL) *T.VOID

	LogonUser func(
		username, domain, password VString,
		logonType, logonProvider T.DWORD,
		token *T.HANDLE) T.BOOL

	LogonUserEx func(
		username, domain, password VString,
		logonType, logonProvider T.DWORD,
		token *T.HANDLE,
		logonSid **T.SID,
		profileBuffer **T.VOID,
		profileLength *T.DWORD,
		quotaLimits *T.QUOTA_LIMITS) T.BOOL

	LookupAccountName func(
		systemName VString,
		accountName OVString,
		sid *T.SID,
		nSid *T.DWORD,
		referencedDomainName OVString,
		nRefDomName *T.DWORD,
		use *T.SID_NAME_USE) T.BOOL

	LookupAccountSid func(
		systemName VString,
		sid *T.SID,
		name OVString,
		nName *T.DWORD,
		referencedDomainName OVString,
		nRefDomName *T.DWORD,
		use *T.SID_NAME_USE) T.BOOL

	LookupPrivilegeDisplayName func(
		systemName, name VString,
		displayName OVString,
		nDispName, languageId *T.DWORD) T.BOOL

	LookupPrivilegeName func(
		systemName VString,
		luid *T.LUID,
		name OVString,
		nName *T.DWORD) T.BOOL

	LookupPrivilegeValue func(
		systemName, name VString, luid *T.LUID) T.BOOL

	Lopen func(pathName T.AString, readWrite int) T.HFILE

	Lread func(file T.HFILE, buffer *T.VOID, bytes T.UINT) T.UINT

	Lstrcat func(string_1, string_2 VString) VString

	Lstrcmp func(string_1, string_2 VString) int

	Lstrcmpi func(string_1, string_2 VString) int

	Lstrcpy func(string_1, string_2 VString) VString

	Lstrcpyn func(
		string_1, string_2 VString,
		maxLength int) VString

	Lstrlen func(string_ VString) int

	Lwrite func(file T.HFILE, buffer T.CH, bytes T.UINT) T.UINT

	MakeAbsoluteSD func(
		selfRelativeSecurityDescriptor,
		absoluteSecurityDescriptor *T.SECURITY_DESCRIPTOR,
		absoluteSecurityDescriptorSize *T.DWORD,
		dacl *T.ACL, daclSize *T.DWORD,
		sacl *T.ACL, saclSize *T.DWORD,
		owner *T.SID, ownerSize *T.DWORD,
		primaryGroup *T.SID, primaryGroupSize *T.DWORD) T.BOOL

	MakeAbsoluteSD2 func(
		selfRelativeSecurityDescriptor *T.SECURITY_DESCRIPTOR,
		bufferSize *T.DWORD) T.BOOL

	MakeSelfRelativeSD func(
		absoluteSecurityDescriptor,
		selfRelativeSecurityDescriptor *T.SECURITY_DESCRIPTOR,
		bufferLength *T.DWORD) T.BOOL

	MapGenericMask func(
		accessMask *T.DWORD, genericMapping *T.GENERIC_MAPPING)

	MapUserPhysicalPages func(
		virtualAddress *T.VOID,
		numberOfPages T.ULONG_PTR,
		pageArray *T.ULONG_PTR) T.BOOL

	MapUserPhysicalPagesScatter func(
		virtualAddresses **T.VOID,
		numberOfPages T.ULONG_PTR,
		pageArray *T.ULONG_PTR) T.BOOL

	MapViewOfFile func(
		fileMappingObject T.HANDLE,
		desiredAccess, fileOffsetHigh, fileOffsetLow T.DWORD,
		numberOfBytesToMap T.SIZE_T) *T.VOID

	MapViewOfFileEx func(
		fileMappingObject T.HANDLE,
		desiredAccess, fileOffsetHigh, fileOffsetLow T.DWORD,
		numberOfBytesToMap T.SIZE_T,
		baseAddress *T.VOID) *T.VOID

	MoveFile func(
		existingFileName, newFileName VString) T.BOOL

	MoveFileEx func(
		existingFileName, newFileName VString,
		flags T.DWORD) T.BOOL

	MoveFileWithProgress func(
		existingFileName,
		newFileName VString,
		progressRoutine T.PROGRESS_ROUTINE,
		data *T.VOID,
		flags T.DWORD) T.BOOL

	MulDiv func(number, numerator, denominator int) int

	NeedCurrentDirectoryForExePath func(
		exeName VString) T.BOOL

	NotifyChangeEventLog func(
		eventLog T.HANDLE, event T.HANDLE) T.BOOL

	ObjectCloseAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		generateOnClose T.BOOL) T.BOOL

	ObjectDeleteAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		generateOnClose T.BOOL) T.BOOL

	ObjectOpenAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		objectTypeName, objectName VString,
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		clientToken T.HANDLE,
		desiredAccess, grantedAccess T.DWORD,
		privileges *T.PRIVILEGE_SET,
		objectCreation, accessGranted T.BOOL,
		generateOnClose *T.BOOL) T.BOOL

	ObjectPrivilegeAuditAlarm func(
		subsystemName VString,
		handleId *T.VOID,
		clientToken T.HANDLE,
		desiredAccess T.DWORD,
		privileges *T.PRIVILEGE_SET,
		accessGranted T.BOOL) T.BOOL

	OpenBackupEventLog func(
		uncServerName, fileName VString) T.HANDLE

	OpenEncryptedFileRaw func(
		fileName VString, flags T.ULONG, context **T.VOID) T.DWORD

	OpenEvent func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		name VString) T.HANDLE

	OpenEventLog func(
		uncServerName, sourceName VString) T.HANDLE

	OpenFile func(
		fileName T.AString,
		reOpenBuff *T.OFSTRUCT,
		style T.UINT) T.HFILE

	OpenFileMapping func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		name VString) T.HANDLE

	OpenJobObject func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		name VString) T.HANDLE

	OpenMutex func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		name VString) T.HANDLE

	OpenProcess func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		processId T.DWORD) T.HANDLE

	OpenProcessToken func(
		processHandle T.HANDLE,
		desiredAccess T.DWORD,
		tokenHandle *T.HANDLE) T.BOOL

	OpenSemaphore func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		name VString) T.HANDLE

	OpenThread func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		threadId T.DWORD) T.HANDLE

	OpenThreadToken func(
		threadHandle T.HANDLE,
		desiredAccess T.DWORD,
		openAsSelf T.BOOL,
		tokenHandle *T.HANDLE) T.BOOL

	OpenWaitableTimer func(
		desiredAccess T.DWORD,
		inheritHandle T.BOOL,
		timerName VString) T.HANDLE

	OutputDebugString func(
		outputString VString)

	PeekNamedPipe func(
		namedPipe T.HANDLE,
		buffer *T.VOID,
		bufferSize T.DWORD,
		bytesRead, totalBytesAvail *T.DWORD,
		bytesLeftThisMessage *T.DWORD) T.BOOL

	PostQueuedCompletionStatus func(
		completionPort T.HANDLE,
		numberOfBytesTransferred T.DWORD,
		completionKey T.ULONG_PTR,
		overlapped *T.OVERLAPPED) T.BOOL

	PrepareTape func(
		device T.HANDLE,
		operation T.DWORD,
		immediate T.BOOL) T.DWORD

	PrivilegeCheck func(
		clientToken T.HANDLE,
		requiredPrivileges *T.PRIVILEGE_SET,
		result *T.BOOL) T.BOOL

	PrivilegedServiceAuditAlarm func(
		subsystemName, serviceName VString,
		clientToken T.HANDLE,
		privileges *T.PRIVILEGE_SET,
		accessGranted T.BOOL) T.BOOL

	ProcessIdToSessionId func(
		processId T.DWORD,
		sessionId *T.DWORD) T.BOOL

	PulseEvent func(event T.HANDLE) T.BOOL

	PurgeComm func(file T.HANDLE, flags T.DWORD) T.BOOL

	QueryActCtxW func(
		flags T.DWORD,
		actCtx T.HANDLE,
		subInstance *T.VOID,
		infoClass T.ULONG,
		buffer *T.VOID,
		size T.SIZE_T,
		writtenOrRequired *T.SIZE_T) T.BOOL

	QueryDepthSList func(listHead *T.SLIST_HEADER) T.USHORT

	QueryDosDevice func(
		deviceName VString,
		targetPath OVString,
		max T.DWORD) T.DWORD

	QueryInformationJobObject func(
		job T.HANDLE,
		class T.JOBOBJECTINFOCLASS,
		info *T.VOID,
		length T.DWORD,
		returnLength *T.DWORD) T.BOOL

	QueryMemoryResourceNotification func(
		notificationHandle T.HANDLE, state *T.BOOL) T.BOOL

	QueryPerformanceCounter func(
		performanceCount *T.LARGE_INTEGER) T.BOOL

	QueryPerformanceFrequency func(
		frequency *T.LARGE_INTEGER) T.BOOL

	QueueUserAPC func(
		apc T.APCFUNC, thread T.HANDLE, data T.ULONG_PTR) T.DWORD

	QueueUserWorkItem func(
		function T.THREAD_START_ROUTINE,
		context *T.VOID,
		flags T.ULONG) T.BOOL

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
		completionRoutine T.OVERLAPPED_COMPLETION_ROUTINE) T.BOOL

	ReadEncryptedFileRaw func(
		exportCallback T.FE_EXPORT_FUNC,
		callbackContext,
		context *T.VOID) T.DWORD

	ReadEventLog func(
		eventLog T.HANDLE,
		readFlags, recordOffset T.DWORD,
		buffer *T.VOID,
		numberOfBytesToRead T.DWORD,
		bytesRead *T.DWORD,
		minNumberOfBytesNeeded *T.DWORD) T.BOOL

	ReadFile func(
		file T.HANDLE,
		buffer *T.VOID,
		numberOfBytesToRead T.DWORD,
		numberOfBytesRead *T.DWORD,
		overlapped *T.OVERLAPPED) T.BOOL

	ReadFileEx func(
		file T.HANDLE,
		buffer *T.VOID,
		numberOfBytesToRead T.DWORD,
		overlapped *T.OVERLAPPED,
		completionRoutine T.OVERLAPPED_COMPLETION_ROUTINE) T.BOOL

	ReadFileScatter func(
		file T.HANDLE,
		segmentArray *T.FILE_SEGMENT_ELEMENT, // []
		numberOfBytesToRead T.DWORD,
		reserved *T.DWORD,
		overlapped *T.OVERLAPPED) T.BOOL

	ReadProcessMemory func(
		process T.HANDLE,
		baseAddress, buffer *T.VOID,
		size T.SIZE_T,
		numberOfBytesRead *T.SIZE_T) T.BOOL

	RegisterEventSource func(
		uncServerName, sourceName VString) T.HANDLE

	RegisterWaitForSingleObject func(
		newWaitObject *T.HANDLE,
		object T.HANDLE,
		callback T.WAITORTIMERCALLBACK,
		context *T.VOID,
		milliseconds, flags T.ULONG) T.BOOL

	RegisterWaitForSingleObjectEx func(
		object T.HANDLE,
		callback T.WAITORTIMERCALLBACK,
		context *T.VOID,
		milliseconds, flags T.ULONG) T.HANDLE

	ReleaseActCtx func(actCtx T.HANDLE)

	ReleaseMutex func(mutex T.HANDLE) T.BOOL

	ReleaseSemaphore func(
		semaphore T.HANDLE,
		releaseCount T.LONG,
		previousCount *T.LONG) T.BOOL

	RemoveDirectory func(pathName VString) T.BOOL

	RemoveVectoredContinueHandler func(Handle *T.VOID) T.ULONG

	RemoveVectoredExceptionHandler func(Handle *T.VOID) T.ULONG

	ReOpenFile func(
		originalFile T.HANDLE,
		desiredAccess, shareMode,
		flagsAndAttributes T.DWORD) T.HANDLE

	ReplaceFile func(
		replacedFileName,
		replacementFileName,
		backupFileName VString,
		replaceFlags T.DWORD,
		exclude, reserved *T.VOID) T.BOOL

	ReportEvent func(
		eventLog T.HANDLE,
		type_, category T.WORD,
		eventID T.DWORD,
		userSid *T.SID,
		numStrings T.WORD,
		dataSize T.DWORD,
		string_s T.OMVString,
		rawData *T.VOID) T.BOOL

	RequestDeviceWakeup func(device T.HANDLE) T.BOOL

	RequestWakeupLatency func(latency T.LATENCY_TIME) T.BOOL

	ResetEvent func(event T.HANDLE) T.BOOL

	ResetWriteWatch func(
		baseAddress *T.VOID, regionSize T.SIZE_T) T.UINT

	RestoreLastError func(errCode T.DWORD)

	ResumeThread func(thread T.HANDLE) T.DWORD

	RevertToSelf func() T.BOOL

	SearchPath func(
		path, fileName, extension VString,
		bufferLength T.DWORD,
		buffer OVString,
		filePart *VString) T.DWORD

	SetAclInformation func(
		acl *T.ACL,
		aclInformation *T.VOID,
		aclInformationLength T.DWORD,
		aclInformationClass T.ACL_INFORMATION_CLASS) T.BOOL

	SetCommBreak func(file T.HANDLE) T.BOOL

	SetCommConfig func(
		commDev T.HANDLE, cc *T.COMMCONFIG, size T.DWORD) T.BOOL

	SetCommMask func(file T.HANDLE, evtMask T.DWORD) T.BOOL

	SetCommState func(file T.HANDLE, dcb *T.DCB) T.BOOL

	SetCommTimeouts func(
		file T.HANDLE, commTimeouts *T.COMMTIMEOUTS) T.BOOL

	SetComputerName func(computerName VString) T.BOOL

	SetComputerNameEx func(
		nameType T.COMPUTER_NAME_FORMAT,
		buffer VString) T.BOOL

	SetCriticalSectionSpinCount func(
		criticalSection *T.CRITICAL_SECTION,
		spinCount T.DWORD) T.DWORD

	SetCurrentDirectory func(pathName VString) T.BOOL

	SetDefaultCommConfig func(
		name VString, cc *T.COMMCONFIG, size T.DWORD) T.BOOL

	SetDllDirectory func(pathName VString) T.BOOL

	SetEndOfFile func(file T.HANDLE) T.BOOL

	SetEnvironmentStrings func(newEnvironment T.MVString) T.BOOL

	SetEnvironmentVariable func(name, value VString) T.BOOL

	SetErrorMode func(mode T.UINT) T.UINT

	SetEvent func(event T.HANDLE) T.BOOL

	SetFileApisToANSI func()

	SetFileApisToOEM func()

	SetFileAttributes func(name VString, attributes T.DWORD) T.BOOL

	SetFilePointer func(
		file T.HANDLE,
		distanceToMove T.LONG,
		distanceToMoveHigh *T.LONG,
		moveMethod T.DWORD) T.DWORD

	SetFilePointerEx func(
		file T.HANDLE,
		distanceToMove T.LARGE_INTEGER,
		newFilePointer *T.LARGE_INTEGER,
		moveMethod T.DWORD) T.BOOL

	SetFileSecurity func(
		fileName VString,
		info T.SECURITY_INFORMATION,
		descriptor *T.SECURITY_DESCRIPTOR) T.BOOL

	SetFileShortName func(
		file T.HANDLE, shortName VString) T.BOOL

	SetFileTime func(
		file T.HANDLE,
		creationTime, lastAccessTime,
		lastWriteTime *T.FILETIME) T.BOOL

	SetFileValidData func(
		file T.HANDLE, validDataLength T.LONGLONG) T.BOOL

	SetFirmwareEnvironmentVariable func(
		name, guid VString, value *T.VOID, size T.DWORD) T.BOOL

	SetHandleCount func(number T.UINT) T.UINT

	SetHandleInformation func(
		object T.HANDLE, mask T.DWORD, flags T.DWORD) T.BOOL

	SetInformationJobObject func(
		job T.HANDLE,
		jobObjectInformationClass T.JOBOBJECTINFOCLASS,
		jobObjectInformation *T.VOID,
		jobObjectInformationLength T.DWORD) T.BOOL

	SetKernelObjectSecurity func(
		handle T.HANDLE,
		securityInformation T.SECURITY_INFORMATION,
		securityDescriptor *T.SECURITY_DESCRIPTOR) T.BOOL

	SetLastError func(errCode T.DWORD)

	SetLocalTime func(systemTime *T.SYSTEMTIME) T.BOOL

	SetMailslotInfo func(
		mailslot T.HANDLE,
		readTimeout T.DWORD) T.BOOL

	SetMessageWaitingIndicator func(
		msgIndicator T.HANDLE,
		msgCount T.ULONG) T.BOOL

	SetNamedPipeHandleState func(
		namedPipe T.HANDLE,
		mode, maxCollectionCount,
		collectDataTimeout *T.DWORD) T.BOOL

	SetPriorityClass func(
		process T.HANDLE, priorityClass T.DWORD) T.BOOL

	SetPrivateObjectSecurity func(
		securityInformation T.SECURITY_INFORMATION,
		modificationDescriptor *T.SECURITY_DESCRIPTOR,
		objectsSecurityDescriptor **T.SECURITY_DESCRIPTOR,
		genericMapping *T.GENERIC_MAPPING,
		token T.HANDLE) T.BOOL

	SetPrivateObjectSecurityEx func(
		securityInformation T.SECURITY_INFORMATION,
		modificationDescriptor *T.SECURITY_DESCRIPTOR,
		objectsSecurityDescriptor **T.SECURITY_DESCRIPTOR,
		autoInheritFlags T.ULONG,
		genericMapping *T.GENERIC_MAPPING,
		token T.HANDLE) T.BOOL

	SetProcessAffinityMask func(
		process T.HANDLE, processAffinityMask T.DWORD_PTR) T.BOOL

	SetProcessPriorityBoost func(
		process T.HANDLE, disablePriorityBoost T.BOOL) T.BOOL

	SetProcessShutdownParameters func(
		level, flags T.DWORD) T.BOOL

	SetProcessWorkingSetSize func(
		process T.HANDLE,
		minWorkingSetSize, maxWorkingSetSize T.SIZE_T) T.BOOL

	SetProcessWorkingSetSizeEx func(
		process T.HANDLE,
		minWorkingSetSize, maxWorkingSetSize T.SIZE_T,
		flags T.DWORD) T.BOOL

	SetSecurityDescriptorControl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		controlBitsOfInterest,
		controlBitsToSet T.SECURITY_DESCRIPTOR_CONTROL) T.BOOL

	SetSecurityDescriptorDacl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		daclPresent T.BOOL,
		dacl *T.ACL,
		daclDefaulted T.BOOL) T.BOOL

	SetSecurityDescriptorGroup func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		group *T.SID,
		groupDefaulted T.BOOL) T.BOOL

	SetSecurityDescriptorOwner func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		owner *T.SID,
		ownerDefaulted T.BOOL) T.BOOL

	SetSecurityDescriptorRMControl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		rmControl *T.WChar) T.DWORD

	SetSecurityDescriptorSacl func(
		securityDescriptor *T.SECURITY_DESCRIPTOR,
		saclPresent T.BOOL,
		sacl *T.ACL,
		saclDefaulted T.BOOL) T.BOOL

	SetStdHandle func(stdHandle T.DWORD, handle T.HANDLE) T.BOOL

	SetSystemFileCacheSize func(
		minFileCacheSize, maxFileCacheSize T.SIZE_T,
		flags T.DWORD) T.BOOL

	SetSystemPowerState func(suspend, force T.BOOL) T.BOOL

	SetSystemTime func(systemTime *T.SYSTEMTIME) T.BOOL

	SetSystemTimeAdjustment func(
		timeAdjustment T.DWORD,
		timeAdjustmentDisabled T.BOOL) T.BOOL

	SetTapeParameters func(
		device T.HANDLE,
		operation T.DWORD,
		tapeInformation *T.VOID) T.DWORD

	SetTapePosition func(
		device T.HANDLE,
		positionMethod,
		partition,
		offsetLow,
		offsetHigh T.DWORD,
		immediate T.BOOL) T.DWORD

	SetThreadAffinityMask func(
		thread T.HANDLE,
		affinityMask T.DWORD_PTR) T.DWORD_PTR

	SetThreadContext func(
		thread T.HANDLE, context *T.CONTEXT) T.BOOL

	SetThreadExecutionState func(
		flags T.EXECUTION_STATE) T.EXECUTION_STATE

	SetThreadIdealProcessor func(
		thread T.HANDLE, idealProcessor T.DWORD) T.DWORD

	SetThreadPriority func(
		thread T.HANDLE, priority int) T.BOOL

	SetThreadPriorityBoost func(
		thread T.HANDLE, disablePriorityBoost T.BOOL) T.BOOL

	SetThreadStackGuarantee func(
		stackSizeInBytes *T.ULONG) T.BOOL

	SetThreadToken func(
		thread *T.HANDLE, token T.HANDLE) T.BOOL

	SetTimerQueueTimer func(
		timerQueue T.HANDLE,
		callback T.WAITORTIMERCALLBACK,
		parameter *T.VOID,
		dueTime, period T.DWORD,
		preferIo T.BOOL) T.HANDLE

	SetTimeZoneInformation func(
		timeZoneInformation *T.TIME_ZONE_INFORMATION) T.BOOL

	SetTokenInformation func(
		tokenHandle T.HANDLE,
		tokenInformationClass T.TOKEN_INFORMATION_CLASS,
		tokenInformation *T.VOID,
		tokenInformationLength T.DWORD) T.BOOL

	SetUnhandledExceptionFilter func(
		tef *T.TOP_LEVEL_EXCEPTION_FILTER) *T.TOP_LEVEL_EXCEPTION_FILTER

	SetupComm func(
		file T.HANDLE, inQueue, outQueue T.DWORD) T.BOOL

	SetVolumeLabel func(
		rootPathName, volumeName VString) T.BOOL

	SetVolumeMountPoint func(
		volumeMountPoint, volumeName VString) T.BOOL

	SetWaitableTimer func(
		timer T.HANDLE,
		dueTime *T.LARGE_INTEGER,
		period T.LONG,
		completionRoutine T.TIMERAPCROUTINE,
		argToCompletionRoutine *T.VOID,
		resume T.BOOL) T.BOOL

	SignalObjectAndWait func(
		objectToSignal, objectToWaitOn T.HANDLE,
		milliseconds T.DWORD,
		alertable T.BOOL) T.DWORD

	SizeofResource func(module T.HMODULE, resInfo T.HRSRC) T.DWORD

	Sleep func(milliseconds T.DWORD)

	SleepEx func(milliseconds T.DWORD, alertable T.BOOL) T.DWORD

	SuspendThread func(thread T.HANDLE) T.DWORD

	SwitchToFiber func(fiber *T.VOID)

	SwitchToThread func() T.BOOL

	SystemTimeToFileTime func(
		systemTime *T.SYSTEMTIME, fileTime *T.FILETIME) T.BOOL

	SystemTimeToTzSpecificLocalTime func(
		timeZoneInformation *T.TIME_ZONE_INFORMATION,
		universalTime, localTime *T.SYSTEMTIME) T.BOOL

	TerminateJobObject func(job T.HANDLE, exitCode T.UINT) T.BOOL

	TerminateProcess func(process T.HANDLE, exitCode T.UINT) T.BOOL

	TerminateThread func(thread T.HANDLE, exitCode T.DWORD) T.BOOL

	TlsAlloc func() T.DWORD

	TlsFree func(tlsIndex T.DWORD) T.BOOL

	TlsGetValue func(tlsIndex T.DWORD) *T.VOID

	TlsSetValue func(tlsIndex T.DWORD, tlsValue *T.VOID) T.BOOL

	TransactNamedPipe func(
		namedPipe T.HANDLE,
		inBuffer *T.VOID, inBufferSize T.DWORD,
		outBuffer *T.VOID, outBufferSize T.DWORD,
		bytesRead *T.DWORD,
		overlapped *T.OVERLAPPED) T.BOOL

	TransmitCommChar func(file T.HANDLE, c char) T.BOOL

	TryEnterCriticalSection func(
		criticalSection *T.CRITICAL_SECTION) T.BOOL

	TzSpecificLocalTimeToSystemTime func(
		timeZoneInformation *T.TIME_ZONE_INFORMATION,
		local, universal *T.SYSTEMTIME) T.BOOL

	UnlockFile func(
		file T.HANDLE,
		offsetLow, offsetHigh, bytesLow, bytesHigh T.DWORD) T.BOOL

	UnlockFileEx func(
		file T.HANDLE,
		reserved, bytesLow, bytesHigh T.DWORD,
		o *T.OVERLAPPED) T.BOOL

	UnmapViewOfFile func(baseAddress *T.VOID) T.BOOL

	UnregisterWait func(waitHandle T.HANDLE) T.BOOL

	UnregisterWaitEx func(
		waitHandle, completionEvent T.HANDLE) T.BOOL

	UpdateResource func(
		update T.HANDLE,
		typ, name VString,
		language T.WORD,
		data *T.VOID,
		count T.DWORD) T.BOOL

	VerifyVersionInfoA func(
		versionInformation *T.OSVERSIONINFOEXA,
		typeMask T.DWORD,
		conditionMask T.DWORDLONG) T.BOOL

	VerifyVersionInfoW func(
		versionInformation *T.OSVERSIONINFOEXW,
		typeMask T.DWORD,
		conditionMask T.DWORDLONG) T.BOOL

	VirtualAlloc func(
		address *T.VOID,
		size T.SIZE_T,
		allocationType, protect T.DWORD) *T.VOID

	VirtualAllocEx func(
		process T.HANDLE,
		address *T.VOID,
		size T.SIZE_T,
		allocationType, protect T.DWORD) *T.VOID

	VirtualFree func(
		address *T.VOID, size T.SIZE_T, freeType T.DWORD) T.BOOL

	VirtualFreeEx func(
		process T.HANDLE,
		address *T.VOID,
		size T.SIZE_T,
		freeType T.DWORD) T.BOOL

	VirtualLock func(address *T.VOID, size T.SIZE_T) T.BOOL

	VirtualProtect func(
		address *T.VOID,
		size T.SIZE_T,
		newProtect T.DWORD,
		oldProtect *T.DWORD) T.BOOL

	VirtualProtectEx func(
		process T.HANDLE,
		address *T.VOID,
		size T.SIZE_T,
		newProtect T.DWORD,
		oldProtect *T.DWORD) T.BOOL

	VirtualQuery func(
		address *T.VOID,
		buffer *T.MEMORY_BASIC_INFORMATION,
		length T.SIZE_T) T.SIZE_T

	VirtualQueryEx func(
		process T.HANDLE,
		address *T.VOID,
		buffer *T.MEMORY_BASIC_INFORMATION,
		length T.SIZE_T) T.SIZE_T

	VirtualUnlock func(address *T.VOID, size T.SIZE_T) T.BOOL

	WaitCommEvent func(
		file T.HANDLE,
		evtMask *T.DWORD,
		overlapped *T.OVERLAPPED) T.BOOL

	WaitForDebugEvent func(
		debugEvent *T.DEBUG_EVENT,
		milliseconds T.DWORD) T.BOOL

	WaitForMultipleObjects func(
		count T.DWORD,
		handles *T.HANDLE,
		waitAll T.BOOL,
		milliseconds T.DWORD) T.DWORD

	WaitForMultipleObjectsEx func(
		count T.DWORD,
		handles *T.HANDLE,
		waitAll T.BOOL,
		milliseconds T.DWORD,
		alertable T.BOOL) T.DWORD

	WaitForSingleObject func(
		handle T.HANDLE, milliseconds T.DWORD) T.DWORD

	WaitForSingleObjectEx func(
		handle T.HANDLE,
		milliseconds T.DWORD,
		alertable T.BOOL) T.DWORD

	WaitNamedPipe func(
		namedPipeName VString,
		timeOut T.DWORD) T.BOOL

	//WinExec is obsolete; instead use:
	//	CreateProcess
	WinExec func(cmdLine string, cmdShow T.UINT) T.UINT

	Wow64DisableWow64FsRedirection func(
		oldValue **T.VOID) T.BOOL

	Wow64EnableWow64FsRedirection func(
		wow64FsEnableRedirection T.BOOLEAN) T.BOOLEAN

	Wow64RevertWow64FsRedirection func(olValue *T.VOID) T.BOOL

	WriteEncryptedFileRaw func(
		importCallback T.FE_IMPORT_FUNC,
		callbackContext, context *T.VOID) T.DWORD

	WriteFile func(
		file T.HANDLE,
		buffer *T.VOID,
		numberOfBytesToWrite T.DWORD,
		numberOfBytesWritten *T.DWORD,
		overlapped *T.OVERLAPPED) T.BOOL

	WriteFileEx func(
		file T.HANDLE,
		buffer *T.VOID,
		numberOfBytesToWrite T.DWORD,
		overlapped *T.OVERLAPPED,
		completionRoutine T.OVERLAPPED_COMPLETION_ROUTINE) T.BOOL

	WriteFileGather func(
		file T.HANDLE,
		segmentArray *T.FILE_SEGMENT_ELEMENT, // []
		numberOfBytesToWrite T.DWORD,
		reserved *T.DWORD,
		overlapped *T.OVERLAPPED) T.BOOL

	//WritePrivateProfileSection is obsolete; instead use:
	//	registry
	WritePrivateProfileSection func(
		appName, string_, fileName VString) T.BOOL

	//WritePrivateProfileString is obsolete; instead use:
	//	registry
	WritePrivateProfileString func(
		appName, keyName, string_, fileName VString) T.BOOL

	//WritePrivateProfileStruct is obsolete; instead use:
	//	registry
	WritePrivateProfileStruct func(
		section, key VString,
		struct_ *T.VOID,
		sizeStruct T.UINT,
		file VString) T.BOOL

	WriteProcessMemory func(
		process T.HANDLE,
		baseAddress, buffer *T.VOID,
		size T.SIZE_T,
		numberOfBytesWritten *T.SIZE_T) T.BOOL

	//WriteProfileSection is obsolete; instead use:
	//	registry
	WriteProfileSection func(
		appName, string_ VString) T.BOOL

	//WriteProfileString is obsolete; instead use:
	//	registry
	WriteProfileString func(
		appName, keyName, string_ VString) T.BOOL

	WriteTapemark func(
		device T.HANDLE,
		tapemarkType, tapemarkCount T.DWORD,
		immediate T.BOOL) T.DWORD

	WTSGetActiveConsoleSessionId func() T.DWORD

	ZombifyActCtx func(actCtx T.HANDLE) T.BOOL

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
