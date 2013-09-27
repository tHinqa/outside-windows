// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package winbase.
package winbase

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
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
	destination *LONGLONG, value LONGLONG) (old LONGLONG) {
	for {
		old = *destination
		if InterlockedCompareExchange64(destination,
			old&value, old) == old {
			return
		}
	}
}

func InterlockedCompareExchangePointer(
	destination *LONG, exChange uintptr,
	comperand uintptr) uintptr {
	return (uintptr)(unsafe.Pointer(uintptr(InterlockedCompareExchange(destination, LONG(exChange), LONG(comperand)))))
}

func InterlockedDecrement64(
	destination *LONGLONG, value LONGLONG) (old LONGLONG) {
	for {
		old = *destination
		if InterlockedCompareExchange64(destination,
			old-1, old) == old {
			return
		}
	}
}

func InterlockedExchange64(
	target *LONGLONG, value LONGLONG) (old LONGLONG) {
	for {
		old = *target
		if InterlockedCompareExchange64(target,
			value, old) == old {
			return
		}
	}
}

func InterlockedExchangeAdd64(
	addend *LONGLONG, value LONGLONG) (old LONGLONG) {
	for {
		old = *addend
		if InterlockedCompareExchange64(addend,
			old+value, old) == old {
			return
		}
	}
}

func InterlockedExchangePointer(
	target *LONG, value LONG) uintptr {
	return (uintptr)(unsafe.Pointer(uintptr(InterlockedExchange(target, value))))
}

func InterlockedIncrement64(
	destination *LONGLONG, value LONGLONG) (old LONGLONG) {
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
	destination *LONGLONG, value LONGLONG) (old LONGLONG) {
	for {
		old = *destination
		if InterlockedCompareExchange64(destination,
			old|value, old) == old {
			return
		}
	}
}

func InterlockedXor64(
	destination *LONGLONG, value LONGLONG) (old LONGLONG) {
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
	char    Fake_type_Fix_me
	va_list Fake_type_Fix_me
	long    Fake_type_Fix_me
)

var ( //TODO(t): Verify all
	AccessCheck func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		clientToken HANDLE,
		desiredAccess DWORD,
		genericMapping *GENERIC_MAPPING,
		privilegeSet *PRIVILEGE_SET,
		privilegeSetLength, grantedAccess *DWORD,
		accessStatus *BOOL) BOOL

	AccessCheckAndAuditAlarm func(
		subsystemName VString,
		handleId *VOID,
		objectTypeName, objectName VString,
		securityDescriptor *SECURITY_DESCRIPTOR,
		desiredAccess DWORD,
		genericMapping *GENERIC_MAPPING,
		objectCreation BOOL,
		grantedAccess *DWORD,
		accessStatus, generateOnClose *BOOL) BOOL

	AccessCheckByType func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		principalSelfSid *SID,
		clientToken HANDLE,
		desiredAccess DWORD,
		objectTypeList *OBJECT_TYPE_LIST,
		objectTypeListLength DWORD,
		genericMapping *GENERIC_MAPPING,
		privilegeSet *PRIVILEGE_SET,
		privilegeSetLength, grantedAccess *DWORD,
		accessStatus *BOOL) BOOL

	AccessCheckByTypeAndAuditAlarm func(
		subsystemName VString,
		handleId *VOID,
		objectTypeName,
		objectName VString,
		securityDescriptor *SECURITY_DESCRIPTOR,
		principalSelfSid *SID,
		desiredAccess DWORD,
		auditType AUDIT_EVENT_TYPE,
		flags DWORD,
		objectTypeList *OBJECT_TYPE_LIST,
		objectTypeListLength DWORD,
		genericMapping *GENERIC_MAPPING,
		objectCreation BOOL,
		grantedAccess *DWORD,
		accessStatus, generateOnClose *BOOL) BOOL

	AccessCheckByTypeResultList func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		principalSelfSid *SID,
		clientToken HANDLE,
		desiredAccess DWORD,
		objectTypeList *OBJECT_TYPE_LIST,
		objectTypeListLength DWORD,
		genericMapping *GENERIC_MAPPING,
		privilegeSet *PRIVILEGE_SET,
		privilegeSetLength, grantedAccessList,
		accessStatusList *DWORD) BOOL

	AccessCheckByTypeResultListAndAuditAlarm func(
		subsystemName VString,
		handleId *VOID,
		objectTypeName, objectName VString,
		securityDescriptor *SECURITY_DESCRIPTOR,
		principalSelfSid *SID,
		desiredAccess DWORD,
		auditType AUDIT_EVENT_TYPE,
		flags DWORD,
		objectTypeList *OBJECT_TYPE_LIST,
		objectTypeListLength DWORD,
		genericMapping *GENERIC_MAPPING,
		objectCreation BOOL,
		grantedAccess, accessStatusList *DWORD,
		generateOnClose *BOOL) BOOL

	AccessCheckByTypeResultListAndAuditAlarmByHandle func(
		subsystemName VString,
		handleId *VOID,
		clientToken HANDLE,
		objectTypeName, objectName VString,
		securityDescriptor *SECURITY_DESCRIPTOR,
		principalSelfSid *SID,
		desiredAccess DWORD,
		auditType AUDIT_EVENT_TYPE,
		flags DWORD,
		objectTypeList *OBJECT_TYPE_LIST,
		objectTypeListLength DWORD,
		genericMapping *GENERIC_MAPPING,
		objectCreation BOOL,
		grantedAccess, accessStatusList *DWORD,
		generateOnClose *BOOL) BOOL

	ActivateActCtx func(
		actCtx HANDLE, cookie *ULONG_PTR) BOOL

	AddAccessAllowedAce func(
		acl *ACL,
		aceRevision, accessMask DWORD,
		sid *SID) BOOL

	AddAccessAllowedAceEx func(
		acl *ACL,
		aceRevision, aceFlags, accessMask DWORD,
		sid *SID) BOOL

	AddAccessAllowedObjectAce func(
		acl *ACL,
		aceRevision, aceFlags, accessMask DWORD,
		objectTypeGuid, inheritedObjectTypeGuid *GUID,
		sid *SID) BOOL

	AddAccessDeniedAce func(
		acl *ACL,
		aceRevision, accessMask DWORD,
		sid *SID) BOOL

	AddAccessDeniedAceEx func(
		acl *ACL,
		aceRevision, aceFlags, accessMask DWORD,
		sid *SID) BOOL

	AddAccessDeniedObjectAce func(
		acl *ACL,
		aceRevision, aceFlags, accessMask DWORD,
		objectTypeGuid, inheritedObjectTypeGuid *GUID,
		sid *SID) BOOL

	AddAce func(
		acl *ACL,
		aceRevision, startingAceIndex DWORD,
		aceList *VOID,
		aceListLength DWORD) BOOL

	AddAtom func(string_ VString) ATOM

	AddAuditAccessAce func(
		acl *ACL,
		aceRevision, accessMask DWORD,
		sid *SID,
		auditSuccess, auditFailure BOOL) BOOL

	AddAuditAccessAceEx func(
		acl *ACL,
		aceRevision, aceFlags,
		accessMask DWORD,
		sid *SID,
		auditSuccess, auditFailure BOOL) BOOL

	AddAuditAccessObjectAce func(
		acl *ACL,
		aceRevision, aceFlags, accessMask,
		objectTypeGuid, inheritedObjectTypeGuid *GUID,
		sid *SID,
		auditSuccess, auditFailure BOOL) BOOL

	AddRefActCtx func(actCtx HANDLE)

	AddVectoredContinueHandler func(
		first ULONG,
		handler VECTORED_EXCEPTION_HANDLER) *VOID

	AddVectoredExceptionHandler func(
		first ULONG,
		handler VECTORED_EXCEPTION_HANDLER) *VOID

	AdjustTokenGroups func(
		tokenHandle HANDLE,
		resetToDefault BOOL,
		newState *TOKEN_GROUPS,
		bufferLength DWORD,
		previousState *TOKEN_GROUPS,
		returnLength *DWORD) BOOL

	AdjustTokenPrivileges func(
		tokenHandle HANDLE,
		disableAllPrivileges BOOL,
		newState *TOKEN_PRIVILEGES,
		bufferLength DWORD,
		previousState *TOKEN_PRIVILEGES,
		returnLength *DWORD) BOOL

	AllocateAndInitializeSid func(
		identifierAuthority *SID_IDENTIFIER_AUTHORITY,
		subAuthorityCount BYTE,
		subAuthority0, subAuthority1,
		subAuthority2, subAuthority3,
		subAuthority4, subAuthority5,
		subAuthority6, subAuthority7 DWORD,
		sid **SID) BOOL

	AllocateLocallyUniqueId func(luid *LUID) BOOL

	AllocateUserPhysicalPages func(
		process HANDLE,
		numberOfPages, pageArray *ULONG_PTR) BOOL

	AreAllAccessesGranted func(
		grantedAccess, desiredAccess DWORD) BOOL

	AreAnyAccessesGranted func(
		grantedAccess, desiredAccess DWORD) BOOL

	AreFileApisANSI func()

	AssignProcessToJobObject func(job, process HANDLE) BOOL

	BackupEventLog func(
		eventLog HANDLE, backupFileName VString) BOOL

	BackupRead func(
		file HANDLE,
		buffer *BYTE,
		numberOfBytesToRead DWORD,
		numberOfBytesRead *DWORD,
		abort, processSecurity BOOL,
		context **VOID) BOOL

	BackupSeek func(
		file HANDLE,
		lowBytesToSeek, HighBytesToSeek DWORD,
		lowByteSeeked, HighByteSeeked *DWORD,
		context **VOID) BOOL

	BackupWrite func(
		file HANDLE,
		buffer *BYTE,
		numberOfBytesToWrite DWORD,
		numberOfBytesWritten *DWORD,
		abort, processSecurity BOOL,
		context **VOID) BOOL

	Beep func(freq, duration DWORD) BOOL

	BeginUpdateResource func(
		fileName VString,
		deleteExistingResources BOOL) HANDLE

	BindIoCompletionCallback func(
		fileHandle HANDLE,
		function OVERLAPPED_COMPLETION_ROUTINE,
		flags ULONG) BOOL

	BuildCommDCB func(def VString, dcb *DCB) BOOL

	BuildCommDCBAndTimeouts func(
		def VString,
		dcb *DCB,
		commTimeouts *COMMTIMEOUTS) BOOL

	CallNamedPipe func(
		namedPipeName VString,
		inBuffer *VOID,
		inBufferSize DWORD,
		outBuffer *VOID,
		outBufferSize DWORD,
		bytesRead *DWORD,
		timeOut DWORD) BOOL

	CancelDeviceWakeupRequest func(device HANDLE) BOOL

	CancelIo func(file HANDLE) BOOL

	CancelTimerQueueTimer func(
		timerQueue, timer HANDLE) BOOL

	CancelWaitableTimer func(timer HANDLE) BOOL

	ChangeTimerQueueTimer func(
		timerQueue, timer HANDLE,
		dueTime, period ULONG) BOOL

	CheckNameLegalDOS8Dot3 func(
		name VString,
		oemName OAString,
		oemNameSize DWORD,
		nameContainsSpaces, nameLegal *BOOL) BOOL

	CheckRemoteDebuggerPresent func(
		process HANDLE, debuggerPresent *BOOL) BOOL

	CheckTokenMembership func(
		tokenHandleHANDLE,
		sidToCheck *SID,
		isMember *BOOL) BOOL

	ClearCommBreak func(file HANDLE) BOOL

	ClearCommError func(
		file HANDLE, errors *DWORD, stat *COMSTAT) BOOL

	ClearEventLog func(
		eventLog HANDLE, backupFileName VString) BOOL

	CloseEncryptedFileRaw func(context *VOID)

	CloseEventLog func(eventLog HANDLE) BOOL

	CloseHandle func(object HANDLE) BOOL

	CommConfigDialog func(
		name VString, wnd HWND, cc *COMMCONFIG) BOOL

	CompareFileTime func(
		fileTime1, fileTime2 *FILETIME) LONG

	ConnectNamedPipe func(
		namedPipe HANDLE, overlapped *OVERLAPPED) BOOL

	ContinueDebugEvent func(
		processId, threadId, continueStatus DWORD) BOOL

	ConvertFiberToThread func() BOOL

	ConvertThreadToFiber func(parameter *VOID) *VOID

	ConvertThreadToFiberEx func(
		parameter *VOID, flags DWORD) *VOID

	ConvertToAutoInheritPrivateObjectSecurity func(
		parentDescriptor,
		currentSecurityDescriptor *SECURITY_DESCRIPTOR,
		newSecurityDescriptor **SECURITY_DESCRIPTOR,
		objectType *GUID,
		isDirectoryObject BOOLEAN,
		genericMapping *GENERIC_MAPPING) BOOL

	CopyFile func(
		existingFileName, newFileName VString,
		failIfExists BOOL) BOOL

	CopyFileEx func(
		existingFileName, newFileName VString,
		progressRoutine PROGRESS_ROUTINE,
		data *VOID,
		cancel *BOOL,
		copyFlags DWORD) BOOL

	CopySid func(
		destinationLength DWORD,
		destination, source *SID) BOOL

	CreateActCtx func(actCtx ACTCTX) HANDLE

	CreateDirectory func(
		pathName VString,
		securityAttributes *SECURITY_ATTRIBUTES) BOOL

	CreateDirectoryEx func(
		templateDirectory, newDirectory VString,
		securityAttributes *SECURITY_ATTRIBUTES) BOOL

	CreateEvent func(
		eventAttributes *SECURITY_ATTRIBUTES,
		manualReset, initialState BOOL,
		name VString) HANDLE

	CreateFiber func(
		stackSize SIZE_T,
		startAddress FIBER_START_ROUTINE,
		parameter *VOID) *VOID

	CreateFiberEx func(
		stackCommitSize, stackReserveSize SIZE_T,
		flags DWORD,
		startAddress FIBER_START_ROUTINE,
		parameter *VOID) *VOID

	CreateFile func(
		fileName VString,
		desiredAccess, shareMode DWORD,
		securityAttributes *SECURITY_ATTRIBUTES,
		creationDisposition, flagsAndAttributes DWORD,
		templateFile HANDLE) HANDLE

	CreateFileMapping func(
		file HANDLE,
		fileMappingAttributes *SECURITY_ATTRIBUTES,
		protect, maximumSizeHigh, maximumSizeLow DWORD,
		name VString) HANDLE

	CreateHardLink func(
		fileName, existingFileName VString,
		securityAttributes *SECURITY_ATTRIBUTES) BOOL

	CreateIoCompletionPort func(
		fileHandle, existingCompletionPort HANDLE,
		completionKey ULONG_PTR,
		numberOfConcurrentThreads DWORD) HANDLE

	CreateJobObject func(
		jobAttributes *SECURITY_ATTRIBUTES,
		name VString) HANDLE

	CreateJobSet func(
		numJob ULONG,
		userJobSet *JOB_SET_ARRAY,
		flags ULONG) BOOL

	CreateMailslot func(
		name VString,
		maxMessageSize, readTimeout DWORD,
		securityAttributes *SECURITY_ATTRIBUTES) HANDLE

	CreateMemoryResourceNotification func(
		MEMORY_RESOURCE_NOTIFICATION_TYPE) HANDLE

	CreateMutex func(
		mutexAttributes *SECURITY_ATTRIBUTES,
		initialOwner BOOL,
		name VString) HANDLE

	CreateNamedPipe func(
		name VString,
		openMode, pipeMode, maxInstances,
		outBufferSize, inBufferSize, defaultTimeOut DWORD,
		securityAttributes *SECURITY_ATTRIBUTES) HANDLE

	CreatePipe func(
		readPipe, writePipe *HANDLE,
		pipeAttributes *SECURITY_ATTRIBUTES,
		size DWORD) BOOL

	CreatePrivateObjectSecurity func(
		parentDescriptor,
		creatorDescriptor *SECURITY_DESCRIPTOR,
		newDescriptor **SECURITY_DESCRIPTOR,
		isDirectoryObject BOOL,
		token HANDLE,
		genericMapping *GENERIC_MAPPING) BOOL

	CreatePrivateObjectSecurityEx func(
		parentDescriptor,
		creatorDescriptor *SECURITY_DESCRIPTOR,
		newDescriptor **SECURITY_DESCRIPTOR,
		objectType *GUID,
		isContainerObject BOOL,
		autoInheritFlags ULONG,
		token HANDLE,
		genericMapping *GENERIC_MAPPING) BOOL

	CreatePrivateObjectSecurityWithMultipleInheritance func(
		parentDescriptor, creatorDescriptor,
		newDescriptor **SECURITY_DESCRIPTOR,
		objectTypes **GUID,
		guidCount ULONG,
		isContainerObject BOOL,
		autoInheritFlags ULONG,
		token HANDLE,
		genericMapping *GENERIC_MAPPING) BOOL

	CreateProcess func(
		applicationName, commandLine VString,
		processAttributes,
		threadAttributes *SECURITY_ATTRIBUTES,
		inheritHandles BOOL,
		creationFlags DWORD,
		environment *VOID,
		currentDirectory VString,
		startupInfo *STARTUPINFO,
		processInformation *PROCESS_INFORMATION) BOOL

	CreateProcessAsUser func(
		token HANDLE,
		applicationName, commandLine VString,
		processAttributes,
		threadAttributes *SECURITY_ATTRIBUTES,
		inheritHandles BOOL,
		creationFlags DWORD,
		environment *VOID,
		currentDirectory VString,
		startupInfo *STARTUPINFO,
		processInformation *PROCESS_INFORMATION) BOOL

	CreateProcessWithLogonW func(
		username, domain, password WString,
		logonFlags DWORD,
		applicationName WString,
		commandLine OWString,
		creationFlags DWORD,
		environment *VOID,
		currentDirectory WString,
		startupInfo *STARTUPINFO,
		processInformation *PROCESS_INFORMATION) BOOL

	CreateProcessWithTokenW func(
		token HANDLE,
		logonFlags DWORD,
		applicationName WString,
		commandLine OWString,
		creationFlags DWORD,
		environment *VOID,
		currentDirectory OWString,
		startupInfo *STARTUPINFO,
		processInformation *PROCESS_INFORMATION) BOOL

	CreateRemoteThread func(
		process HANDLE,
		threadAttributes *SECURITY_ATTRIBUTES,
		stackSize SIZE_T,
		startAddress THREAD_START_ROUTINE,
		parameter *VOID,
		creationFlags DWORD,
		threadId *DWORD) HANDLE

	CreateRestrictedToken func(
		existingTokenHandle HANDLE,
		flags, disableSidCount, deletePrivilegeCount,
		restrictedSidCount DWORD,
		newTokenHandle *HANDLE) BOOL

	CreateSemaphore func(
		semaphoreAttributes *SECURITY_ATTRIBUTES,
		initialCount, maximumCount LONG,
		name VString) HANDLE

	CreateTapePartition func(
		device HANDLE,
		partitionMethod, count, size DWORD) DWORD

	CreateThread func(
		threadAttributes *SECURITY_ATTRIBUTES,
		stackSize SIZE_T,
		startAddress THREAD_START_ROUTINE,
		parameter *VOID,
		creationFlags DWORD,
		threadId *DWORD) HANDLE

	CreateTimerQueue func() HANDLE

	CreateTimerQueueTimer func(
		newTimer *HANDLE,
		timerQueue HANDLE,
		callback WAITORTIMERCALLBACK,
		parameter *VOID,
		dueTime, period DWORD,
		flags ULONG) BOOL

	CreateWaitableTimer func(
		timerAttributes *SECURITY_ATTRIBUTES,
		manualReset BOOL,
		timerName VString) HANDLE

	CreateWellKnownSid func(
		wellKnownSidType WELL_KNOWN_SID_TYPE,
		domainSid, sid *SID,
		sidl *DWORD) BOOL

	DeactivateActCtx func(
		flags DWORD, cookie ULONG_PTR) BOOL

	DebugActiveProcess func(processId DWORD) BOOL

	DebugActiveProcessStop func(processId DWORD) BOOL

	DebugBreak func()

	DebugBreakProcess func(process HANDLE) BOOL

	DebugSetProcessKillOnExit func(killOnExit BOOL) BOOL

	DecodePointer func(ptr *VOID) *VOID

	DecodeSystemPointer func(ptr *VOID) *VOID

	DecryptFile func(fileName VString, reserved DWORD) BOOL

	DefineDosDevice func(
		flags DWORD,
		deviceName, targetPath VString) BOOL

	DeleteAce func(acl *ACL, aceIndex DWORD) BOOL

	DeleteAtom func(atom ATOM) ATOM

	DeleteCriticalSection func(*CRITICAL_SECTION)

	DeleteFiber func(fiber *VOID)

	DeleteFile func(fileName VString) BOOL

	DeleteTimerQueue func(timerQueue HANDLE) BOOL

	DeleteTimerQueueEx func(
		timerQueue, completionEvent HANDLE) BOOL

	DeleteTimerQueueTimer func(
		timerQueue, timer, completionEvent HANDLE) BOOL

	DeleteVolumeMountPoint func(
		volumeMountPoint VString) BOOL

	DeregisterEventSource func(eventLog HANDLE) BOOL

	DestroyPrivateObjectSecurity func(
		objectDescriptor **SECURITY_DESCRIPTOR) BOOL

	DeviceIoControl func(
		device HANDLE,
		ioControlCode DWORD,
		inBuffer *VOID,
		inBufferSize DWORD,
		outBuffer *VOID,
		outBufferSize, bytesReturned *DWORD,
		overlapped *OVERLAPPED) BOOL

	DisableThreadLibraryCalls func(libModule HMODULE) BOOL

	DisconnectNamedPipe func(namedPipe HANDLE) BOOL

	DnsHostnameToComputerName func(
		hostname VString,
		computerName OVString,
		size *DWORD) BOOL

	DosDateTimeToFileTime func(
		fatDate, fatTime WORD, fileTime *FILETIME) BOOL

	DuplicateHandle func(
		sourceProcessHandle, sourceHandle,
		targetProcessHandle HANDLE,
		targetHandle *HANDLE,
		desiredAccess DWORD,
		inheritHandle BOOL,
		options DWORD) BOOL

	DuplicateToken func(
		existingTokenHandle HANDLE,
		impersonationLevel SECURITY_IMPERSONATION_LEVEL,
		duplicateTokenHandle *HANDLE) BOOL

	DuplicateTokenEx func(
		existingToken HANDLE,
		desiredAccess DWORD,
		tokenAttributes *SECURITY_ATTRIBUTES,
		impersonationLevel SECURITY_IMPERSONATION_LEVEL,
		tokenType TOKEN_TYPE,
		newToken *HANDLE) BOOL

	EncodePointer func(ptr *VOID) *VOID

	EncodeSystemPointer func(ptr *VOID) *VOID

	EncryptFile func(fileName VString) BOOL

	EndUpdateResource func(
		update HANDLE, discard BOOL) BOOL

	EnterCriticalSection func(*CRITICAL_SECTION)

	EnumResourceLanguages func(
		module HMODULE,
		type_, name VString,
		enumFunc ENUMRESLANGPROC,
		param LONG_PTR) BOOL

	EnumResourceNames func(
		module HMODULE,
		type_ VString,
		enumFunc ENUMRESNAMEPROC,
		param LONG_PTR) BOOL

	EnumResourceTypes func(
		module HMODULE,
		enumFunc ENUMRESTYPEPROC,
		param LONG_PTR) BOOL

	EnumSystemFirmwareTables func(
		firmwareTableProviderSignature DWORD,
		firmwareTableEnumBuffer *VOID,
		bufferSize DWORD) UINT

	EqualDomainSid func(
		sid1, sid2 *SID, equal *BOOL) BOOL

	EqualPrefixSid func(
		sid1, sid2 *SID) BOOL

	EqualSid func(sid1, sid2 *SID) BOOL

	EraseTape func(
		device HANDLE,
		eraseType DWORD,
		immediate BOOL) DWORD

	EscapeCommFunction func(file HANDLE, function DWORD) BOOL

	ExitProcess func(exitCode UINT)

	ExitThread func(exitCode DWORD)

	ExpandEnvironmentStrings func(
		src, dst VString, size DWORD) DWORD

	FatalAppExit func(action UINT, messageText VString)

	FatalExit func(exitCode int)

	FileEncryptionStatus func(
		fileName VString, status *DWORD) BOOL

	FileTimeToDosDateTime func(
		fileTime *FILETIME, fatDate, fatTime *WORD) BOOL

	FileTimeToLocalFileTime func(
		fileTime, localFileTime *FILETIME) BOOL

	FileTimeToSystemTime func(
		fileTime, systemTime *SYSTEMTIME) BOOL

	FindActCtxSectionGuid func(
		flags DWORD,
		extensionGuid *GUID,
		sectionId ULONG,
		guidToFind *GUID,
		returnedData *ACTCTX_SECTION_KEYED_DATA) BOOL

	FindActCtxSectionString func(
		flags DWORD,
		extensionGuid *GUID,
		sectionId ULONG,
		string_ToFind VString,
		returnedData *ACTCTX_SECTION_KEYED_DATA) BOOL

	FindAtom func(string_ VString) ATOM

	FindClose func(findFile HANDLE) BOOL

	FindCloseChangeNotification func(
		changeHandle HANDLE) BOOL

	FindFirstChangeNotification func(
		pathName VString,
		watchSubtree BOOL,
		notifyFilter DWORD) HANDLE

	FindFirstFileA func(
		fileName AString,
		findFileData *WIN32_FIND_DATAA) HANDLE

	FindFirstFileW func(
		fileName WString,
		findFileData *WIN32_FIND_DATAW) HANDLE

	FindFirstFileEx func(
		fileName VString,
		infoLevelId FINDEX_INFO_LEVELS,
		findFileData *VOID,
		searchOp FINDEX_SEARCH_OPS,
		searchFilter *VOID,
		additionalFlags DWORD) HANDLE

	FindFirstFreeAce func(acl *ACL, ace **VOID) BOOL

	FindFirstVolume func(
		volumeName OVString, bufferLength DWORD) HANDLE

	FindFirstVolumeMountPoint func(
		rootPathName VString,
		volumeMountPoint OVString,
		bufferLength DWORD) HANDLE

	FindFirstStreamW func(
		fileName WString,
		infoLevel STREAM_INFO_LEVELS,
		findStreamData *VOID,
		flags DWORD) HANDLE

	FindNextStreamW func(
		findStream HANDLE, data *VOID) BOOL

	FindNextChangeNotification func(
		changeHandle HANDLE) BOOL

	FindNextFileA func(
		findFile HANDLE, data *WIN32_FIND_DATAA) BOOL

	FindNextFileW func(
		findFile HANDLE, data *WIN32_FIND_DATAW) BOOL

	FindNextVolume func(
		findVolume HANDLE,
		volumeName OVString,
		bufferLength DWORD) BOOL

	FindNextVolumeMountPoint func(
		findVolumeMountPoint HANDLE,
		volumeMountPoint OVString,
		bufferLength DWORD) BOOL

	FindResource func(
		module HMODULE,
		name, typ VString) HRSRC

	FindResourceEx func(
		module HMODULE,
		typ, name VString,
		language WORD) HRSRC

	FindVolumeClose func(findVolume HANDLE) BOOL

	FindVolumeMountPointClose func(
		findVolumeMountPoint HANDLE) BOOL

	FlsAlloc func(callback FLS_CALLBACK_FUNCTION) DWORD

	FlsFree func(flsIndex DWORD) BOOL

	FlsGetValue func(flsIndex DWORD) *VOID

	FlsSetValue func(flsIndex DWORD, flsData *VOID) BOOL

	FlushFileBuffers func(file HANDLE) BOOL

	FlushInstructionCache func(
		process HANDLE, baseAddress *VOID, size SIZE_T) BOOL

	FlushViewOfFile func(
		baseAddress *VOID, numberOfBytesToFlush SIZE_T) BOOL

	FormatMessage func(
		flags DWORD,
		source *VOID,
		messageId, languageId DWORD,
		buffer VString,
		size DWORD,
		arguments *va_list) DWORD

	FreeEnvironmentStrings func(VString) BOOL

	FreeLibrary func(libModule HMODULE) BOOL

	FreeLibraryAndExitThread func(
		libModule HMODULE, exitCode DWORD)

	FreeResource func(resData HGLOBAL) BOOL

	FreeSid func(sid *SID) *VOID

	FreeUserPhysicalPages func(
		process HANDLE,
		numberOfPages, pageArray *ULONG_PTR) BOOL

	GetAce func(
		acl *ACL, aceIndex DWORD, ace **VOID) BOOL

	GetAclInformation func(
		acl *ACL,
		info *VOID,
		length DWORD,
		class ACL_INFORMATION_CLASS) BOOL

	GetAtomName func(
		atom ATOM, buffer VString, size int) UINT

	GetBinaryType func(
		applicationName VString, binaryType *DWORD) BOOL

	GetCommandLine func() VString

	GetCommConfig func(
		commDev HANDLE, cc *COMMCONFIG, size *DWORD) BOOL

	GetCommMask func(file HANDLE, evtMask *DWORD) BOOL

	GetCommModemStatus func(
		file HANDLE, modemStat *DWORD) BOOL

	GetCommProperties func(
		file HANDLE, cp *COMMPROP) BOOL

	GetCommState func(file HANDLE, dcn *DCB) BOOL

	GetCommTimeouts func(
		file HANDLE, commTimeouts *COMMTIMEOUTS) BOOL

	GetCompressedFileSize func(
		fileName VString, fileSizeHigh *DWORD) DWORD

	GetComputerName func(
		buffer OVString, size *DWORD) BOOL

	GetComputerNameEx func(
		nameType COMPUTER_NAME_FORMAT,
		buffer OVString,
		size *DWORD) BOOL

	GetCurrentActCtx func(actCtx *HANDLE) BOOL

	GetCurrentDirectory func(
		bufferLength DWORD, buffer VString) DWORD

	GetCurrentHwProfileA func(
		hwProfileInfo *HW_PROFILE_INFOA) BOOL

	GetCurrentHwProfileW func(
		hwProfileInfo *HW_PROFILE_INFOW) BOOL

	GetCurrentProcess func() HANDLE

	GetCurrentProcessId func() DWORD

	GetCurrentProcessorNumber func() DWORD

	GetCurrentThread func() HANDLE

	GetCurrentThreadId func() DWORD

	GetDefaultCommConfig func(
		name VString, cc *COMMCONFIG, size *DWORD) BOOL

	GetDevicePowerState func(
		device HANDLE, on *BOOL) BOOL

	GetDiskFreeSpace func(
		rootPathName VString,
		sectorsPerCluster,
		bytesPerSector,
		numberOfFreeClusters,
		totalNumberOfClusters *DWORD) BOOL

	GetDiskFreeSpaceEx func(
		directoryName VString,
		freeBytesAvailableToCaller,
		totalNumberOfBytes,
		totalNumberOfFreeBytes *ULARGE_INTEGER) BOOL

	GetDllDirectory func(
		bufferLength DWORD, buffer VString) DWORD

	GetDriveType func(rootPathName VString) UINT

	GetEnvironmentStrings func() OMVString

	GetEnvironmentVariable func(
		name, buffer VString, size DWORD) DWORD

	GetEventLogInformation func(
		eventLog HANDLE,
		infoLevel DWORD,
		buffer *VOID,
		bufSize DWORD,
		bytesNeeded *DWORD) BOOL

	GetExitCodeProcess func(process HANDLE, exitCode *DWORD) BOOL

	GetExitCodeThread func(thread HANDLE, exitCode *DWORD) BOOL

	GetFileAttributes func(fileName VString) DWORD

	GetFileAttributesEx func(
		fileName VString,
		infoLevelId GET_FILEEX_INFO_LEVELS,
		fileInformation *VOID) BOOL

	GetFileInformationByHandle func(
		file HANDLE,
		fileInformation *BY_HANDLE_FILE_INFORMATION) BOOL

	GetFileSecurity func(
		fileName VString,
		requestedInformation SECURITY_INFORMATION,
		securityDescriptor *SECURITY_DESCRIPTOR,
		length DWORD,
		lengthNeeded *DWORD) BOOL

	GetFileSize func(
		file HANDLE, fileSizeHigh *DWORD) DWORD

	GetFileSizeEx func(
		file HANDLE, fileSize *LARGE_INTEGER) BOOL

	GetFileTime func(
		file HANDLE,
		creationTime, lastAccessTime,
		lastWriteTime *FILETIME) BOOL

	GetFileType func(file HANDLE) DWORD

	GetFirmwareEnvironmentVariable func(
		name, guid VString, buffer *VOID, size DWORD) DWORD

	GetFullPathName func(
		fileName VString,
		bufferLength DWORD,
		buffer VString,
		filePart OVString) DWORD

	GetHandleInformation func(
		object HANDLE, flags *DWORD) BOOL

	GetKernelObjectSecurity func(
		handle HANDLE,
		requestedInformation SECURITY_INFORMATION,
		securityDescriptor *SECURITY_DESCRIPTOR,
		length DWORD,
		lengthNeeded *DWORD) BOOL

	GetLargePageMinimum func() SIZE_T

	GetLastError func() DWORD

	GetLengthSid func(sid *SID) DWORD

	GetLocalTime func(systemTime *SYSTEMTIME)

	GetLogicalDrives func() DWORD

	GetLogicalDriveStrings func(
		bufferLength DWORD,
		buffer OMVString) DWORD

	GetLogicalProcessorInformation func(
		buffer *SYSTEM_LOGICAL_PROCESSOR_INFORMATION,
		returnedLength *DWORD) BOOL

	GetLongPathName func(
		shortPath VString,
		longPath OVString,
		buffer DWORD) DWORD

	GetMailslotInfo func(
		mailslot HANDLE,
		maxMessageSize, nextSize, messageCount,
		readTimeout *DWORD) BOOL

	GetModuleFileName func(
		module HMODULE, filename VString, size DWORD) DWORD

	GetModuleHandle func(moduleName VString) HMODULE

	GetModuleHandleEx func(
		flags DWORD,
		moduleName VString,
		module *HMODULE) BOOL

	GetNamedPipeHandleState func(
		namedPipe HANDLE,
		state, curInstances, maxCollectionCount,
		collectDataTimeout *DWORD,
		userName VString,
		maxUserNameSize DWORD) BOOL

	GetNamedPipeInfo func(
		namedPipe HANDLE,
		flags, outBufferSize, inBufferSize,
		maxInstances *DWORD) BOOL

	GetNativeSystemInfo func(systemInfo *SYSTEM_INFO)

	GetNumaAvailableMemoryNode func(
		node WChar,
		availableBytes *ULONGLONG) BOOL

	GetNumaHighestNodeNumber func(
		highestNodeNumber *ULONG) BOOL

	GetNumaNodeProcessorMask func(
		node WChar,
		processorMask *ULONGLONG) BOOL

	GetNumaProcessorNode func(
		processor WChar,
		nodeNumber *WChar) BOOL

	GetNumberOfEventLogRecords func(
		eventLog HANDLE,
		numberOfRecords *DWORD) BOOL

	GetOldestEventLogRecord func(
		eventLog HANDLE,
		oldestRecord *DWORD) BOOL

	GetOverlappedResult func(
		file HANDLE,
		overlapped *OVERLAPPED,
		numberOfBytesTransferred *DWORD,
		wait BOOL) BOOL

	GetPriorityClass func(process HANDLE) DWORD

	GetPrivateObjectSecurity func(
		objectDescriptor *SECURITY_DESCRIPTOR,
		securityInformation SECURITY_INFORMATION,
		resultantDescriptor *SECURITY_DESCRIPTOR,
		descriptorLength DWORD,
		returnLength *DWORD) BOOL

	GetPrivateProfileInt func(
		appName, keyName VString,
		default_ INT,
		fileName VString) UINT

	GetPrivateProfileSection func(
		appName, returnedString VString,
		size DWORD,
		fileName VString) DWORD

	GetPrivateProfileSectionNames func(
		returnBuffer VString,
		size DWORD,
		fileName VString) DWORD

	GetPrivateProfileString func(
		appName, keyName, default_, returned VString,
		size DWORD,
		fileName VString) DWORD

	GetPrivateProfileStruct func(
		section, key VString,
		struct_ *VOID,
		sizeStruct UINT,
		file VString) BOOL

	GetProcAddress func(
		module HMODULE, procName AString) FARPROC

	GetProcessAffinityMask func(
		process HANDLE,
		processAffinityMask,
		systemAffinityMask *DWORD_PTR) BOOL

	GetProcessHandleCount func(
		process HANDLE, HandleCount *DWORD) BOOL

	GetProcessHeap func() HANDLE

	GetProcessHeaps func(
		numberOfHeaps DWORD,
		processHeaps *HANDLE) DWORD

	GetProcessId func(process HANDLE) DWORD

	GetProcessIdOfThread func(thread HANDLE) DWORD

	GetProcessIoCounters func(
		process HANDLE, ioCounters *IO_COUNTERS) BOOL

	GetProcessPriorityBoost func(
		process HANDLE, disablePriorityBoost *BOOL) BOOL

	GetProcessShutdownParameters func(
		level, flags *DWORD) BOOL

	GetProcessTimes func(
		process HANDLE,
		creationTime, exitTime, kernelTime,
		userTime *FILETIME) BOOL

	GetProcessVersion func(
		processId DWORD) DWORD

	GetProcessWorkingSetSize func(
		process HANDLE,
		minWorkingSetSize, maxWorkingSetSize *SIZE_T) BOOL

	GetProcessWorkingSetSizeEx func(
		process HANDLE,
		minWorkingSetSize, maxWorkingSetSize *SIZE_T,
		flags *DWORD) BOOL

	GetProfileInt func(
		appName, keyName VString,
		default_ INT) UINT

	GetProfileSection func(
		appName, returnedString VString,
		size DWORD) DWORD

	GetProfileString func(
		appName, keyName, default_, returned VString,
		size DWORD) DWORD

	GetQueuedCompletionStatus func(
		completionPort HANDLE,
		numberOfBytesTransferred *DWORD,
		completionKey *ULONG_PTR,
		overlapped **OVERLAPPED,
		milliseconds DWORD) BOOL

	GetSecurityDescriptorControl func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		control *SECURITY_DESCRIPTOR_CONTROL,
		revision *DWORD) BOOL

	GetSecurityDescriptorDacl func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		daclPresent *BOOL,
		dacl **ACL,
		daclDefaulted *BOOL) BOOL

	GetSecurityDescriptorGroup func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		group **SID,
		groupDefaulted *BOOL) BOOL

	GetSecurityDescriptorLength func(
		securityDescriptor *SECURITY_DESCRIPTOR) DWORD

	GetSecurityDescriptorOwner func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		owner **SID,
		ownerDefaulted *BOOL) BOOL

	GetSecurityDescriptorRMControl func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		rmControl *WChar) DWORD

	GetSecurityDescriptorSacl func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		saclPresent *BOOL,
		sacl **ACL,
		saclDefaulted *BOOL) BOOL

	GetShortPathName func(
		longPath VString,
		shortPath OVString,
		buffer DWORD) DWORD

	GetSidIdentifierAuthority func(
		sid *SID) *SID_IDENTIFIER_AUTHORITY

	GetSidLengthRequired func(subAuthorityCount WChar) DWORD

	GetSidSubAuthority func(
		sid *SID, subAuthority DWORD) *DWORD

	GetSidSubAuthorityCount func(sid *SID) *WChar

	GetStartupInfo func(startupInfo *STARTUPINFO)

	GetStdHandle func(stdHandle DWORD) HANDLE

	GetSystemDirectory func(buffer VString, size UINT) UINT

	GetSystemFileCacheSize func(
		minFileCacheSize, maxFileCacheSize *SIZE_T,
		flags *DWORD) BOOL

	GetSystemFirmwareTable func(
		firmwareTableProviderSignature,
		firmwareTableID DWORD,
		firmwareTableBuffer *VOID,
		bufferSize DWORD) UINT

	GetSystemInfo func(systemInfo *SYSTEM_INFO)

	GetSystemPowerStatus func(
		systemPowerStatus *SYSTEM_POWER_STATUS) BOOL

	GetSystemRegistryQuota func(
		quotaAllowed, quotaUsed *DWORD) BOOL

	GetSystemTime func(
		systemTime *SYSTEMTIME)

	GetSystemTimeAdjustment func(
		timeAdjustment, timeIncrement *DWORD,
		timeAdjustmentDisabled *BOOL) BOOL

	GetSystemTimeAsFileTime func(
		systemTimeAsFileTime *FILETIME)

	GetSystemTimes func(
		idleTime, kernelTime, userTime *FILETIME) BOOL

	GetSystemWindowsDirectory func(
		buffer VString, size UINT) UINT

	GetSystemWow64Directory func(
		buffer VString, size UINT) UINT

	GetTapeParameters func(
		device HANDLE,
		operation DWORD,
		size *DWORD,
		tapeInformation *VOID) DWORD

	GetTapePosition func(
		device HANDLE,
		positionType DWORD,
		partition, offsetLow, offsetHigh *DWORD) DWORD

	GetTapeStatus func(device HANDLE) DWORD

	GetTempFileName func(
		pathName, prefixString VString,
		unique UINT,
		tempFileName VString) UINT

	GetTempPath func(bufferLength DWORD, buffer VString) DWORD

	GetThreadContext func(thread HANDLE, context *CONTEXT) BOOL

	GetThreadId func(thread HANDLE) DWORD

	GetThreadIOPendingFlag func(
		thread HANDLE, ioIsPending *BOOL) BOOL

	GetThreadPriority func(thread HANDLE) int

	GetThreadPriorityBoost func(
		thread HANDLE, disablePriorityBoost *BOOL) BOOL

	GetThreadSelectorEntry func(
		thread HANDLE,
		selector DWORD,
		selectorEntry *LDT_ENTRY) BOOL

	GetThreadTimes func(
		thread HANDLE,
		creationTime, exitTime, kernelTime,
		userTime *FILETIME) BOOL

	GetTickCount func() DWORD

	GetTimeZoneInformation func(
		timeZoneInformation *TIME_ZONE_INFORMATION) DWORD

	GetTokenInformation func(
		tokenHandle HANDLE,
		tokenInformationClass TOKEN_INFORMATION_CLASS,
		tokenInformation *VOID,
		tokenInformationLength DWORD,
		returnLength *DWORD) BOOL

	GetUserName func(buffer OVString, size *DWORD) BOOL

	GetVersion func() DWORD

	GetVersionExA func(versionInformation *OSVERSIONINFOA) BOOL

	GetVersionExW func(versionInformation *OSVERSIONINFOW) BOOL

	GetVolumeInformation func(
		rootPathName, volumeNameBuffer VString,
		volumeNameSize DWORD,
		volumeSerialNumber, maximumComponentLength,
		fileSystemFlags *DWORD,
		fileSystemNameBuffer VString,
		fileSystemNameSize DWORD) BOOL

	GetVolumeNameForVolumeMountPoint func(
		volumeMountPoint VString,
		volumeName OVString,
		bufferLength DWORD) BOOL

	GetVolumePathName func(
		fileName VString,
		volumePathName OVString,
		bufferLength DWORD) BOOL

	GetVolumePathNamesForVolumeName func(
		volumeName VString,
		volumePathNames OMVString,
		bufferLength DWORD,
		returnLength *DWORD) BOOL

	GetWindowsAccountDomainSid func(
		sid, domainSid *SID, domainSidl *DWORD) BOOL

	GetWindowsDirectory func(buffer OVString, size UINT) UINT

	GetWriteWatch func(
		flags DWORD,
		baseAddress *VOID,
		regionSize SIZE_T,
		addresses **VOID,
		count *ULONG_PTR,
		granularity *ULONG) UINT

	GlobalAddAtom func(s VString) ATOM

	GlobalAlloc func(flags UINT, bytes SIZE_T) HGLOBAL

	GlobalCompact func(minFree DWORD) SIZE_T

	GlobalDeleteAtom func(a ATOM) ATOM

	GlobalFindAtom func(s VString) ATOM

	GlobalFix func(mem HGLOBAL)

	GlobalFlags func(mem HGLOBAL) UINT

	GlobalFree func(mem HGLOBAL) HGLOBAL

	GlobalGetAtomName func(
		atom ATOM, buffer OVString, size int) UINT

	GlobalHandle func(mem *VOID) HGLOBAL

	GlobalLock func(mem HGLOBAL) *VOID

	GlobalMemoryStatus func(buffer *MEMORYSTATUS)

	GlobalMemoryStatusEx func(buffer *MEMORYSTATUSEX) BOOL

	GlobalReAlloc func(
		mem HGLOBAL, bytes SIZE_T, flags UINT) HGLOBAL

	GlobalSize func(mem HGLOBAL) SIZE_T

	GlobalUnfix func(mem HGLOBAL)

	GlobalUnlock func(mem HGLOBAL) BOOL

	GlobalUnWire func(mem HGLOBAL) BOOL

	GlobalWire func(mem HGLOBAL) *VOID

	HeapAlloc func(heap HANDLE, flags DWORD, bytes SIZE_T) *VOID

	HeapCompact func(Heap HANDLE, flags DWORD) SIZE_T

	HeapCreate func(
		options DWORD, initialSize, maximumSize SIZE_T) HANDLE

	HeapDestroy func(Heap HANDLE) BOOL

	HeapFree func(Heap HANDLE, flags DWORD, mem *VOID) BOOL

	HeapLock func(Heap HANDLE) BOOL

	HeapQueryInformation func(
		h HANDLE,
		ic HEAP_INFORMATION_CLASS,
		info *VOID,
		infoLength SIZE_T,
		returnLength *SIZE_T) BOOL

	HeapReAlloc func(
		heap HANDLE, flags DWORD, mem *VOID, bytes SIZE_T) *VOID

	HeapSetInformation func(
		h, HANDLE,
		ic HEAP_INFORMATION_CLASS,
		info *VOID,
		infoLength SIZE_T) BOOL

	HeapSize func(heap HANDLE, flags DWORD, mem *VOID) SIZE_T

	HeapUnlock func(Heap HANDLE) BOOL

	HeapValidate func(heap HANDLE, flags DWORD, mem *VOID) BOOL

	HeapWalk func(heap HANDLE, entry *PROCESS_HEAP_ENTRY) BOOL

	Hread func(file HFILE, buffer *VOID, bytes long) long

	Hwrite func(file HFILE, buffer *CH, bytes long) long

	ImpersonateAnonymousToken func(threadHandle HANDLE) BOOL

	ImpersonateLoggedOnUser func(token HANDLE) BOOL

	ImpersonateNamedPipeClient func(namedPipe HANDLE) BOOL

	ImpersonateSelf func(
		impersonationLevel SECURITY_IMPERSONATION_LEVEL) BOOL

	InitAtomTable func(size DWORD) BOOL

	InitializeAcl func(
		acl *ACL, aclLength DWORD, aclRevision DWORD) BOOL

	InitializeCriticalSection func(*CRITICAL_SECTION)

	InitializeCriticalSectionAndSpinCount func(
		cs *CRITICAL_SECTION, spinCount DWORD) BOOL

	InitializeSecurityDescriptor func(
		secDesc *SECURITY_DESCRIPTOR,
		revision DWORD) BOOL

	InitializeSid func(
		sid *SID,
		idAuth *SID_IDENTIFIER_AUTHORITY,
		subAuthorityCount BYTE) BOOL

	InitializeSListHead func(listHead SLIST_HEADER)

	InterlockedCompareExchange func(
		destination *LONG, exchange, comperand LONG) LONG

	InterlockedCompareExchange64 func(
		destination *LONGLONG,
		exchange, comperand LONGLONG) LONGLONG

	InterlockedDecrement func(addend *LONG) LONG

	InterlockedExchange func(target *LONG, value LONG) LONG

	InterlockedExchangeAdd func(addend *LONG, value LONG) LONG

	InterlockedFlushSList func(*SLIST_HEADER) *SLIST_ENTRY

	InterlockedIncrement func(addend *LONG) LONG

	InterlockedPopEntrySList func(*SLIST_HEADER) *SLIST_ENTRY

	InterlockedPushEntrySList func(
		*SLIST_HEADER, *SLIST_ENTRY) *SLIST_ENTRY

	IsBadCodePtr func(FARPROC) BOOL

	IsBadHugeReadPtr func(*VOID, UINT_PTR) BOOL

	IsBadHugeWritePtr func(*VOID, UINT_PTR) BOOL

	IsBadReadPtr func(*VOID, UINT_PTR) BOOL

	IsBadStringPtr func(s VString, max UINT_PTR) BOOL

	IsBadWritePtr func(*VOID, UINT_PTR) BOOL

	IsDebuggerPresent func() BOOL

	IsProcessInJob func(
		processHandle HANDLE,
		jobHandle HANDLE,
		result *BOOL) BOOL

	IsProcessorFeaturePresent func(processorFeature DWORD) BOOL

	IsSystemResumeAutomatic func() BOOL

	IsTextUnicode func(v *VOID, size int, result *INT) BOOL

	IsTokenRestricted func(token HANDLE) BOOL

	IsTokenUntrusted func(token HANDLE) BOOL

	IsValidAcl func(*ACL) BOOL

	IsValidSecurityDescriptor func(*SECURITY_DESCRIPTOR) BOOL

	IsValidSid func(*SID) BOOL

	IsWellKnownSid func(*SID, WELL_KNOWN_SID_TYPE) BOOL

	IsWow64Process func(process HANDLE, wow64Process *BOOL) BOOL

	Lclose func(file HFILE) HFILE

	Lcreat func(pathName AString, attribute int) HFILE

	LeaveCriticalSection func(*CRITICAL_SECTION)

	Llseek func(file HFILE, offset LONG, origin int) LONG

	LoadLibrary func(libFileName VString) HMODULE

	LoadLibraryEx func(
		libFileName VString, file HANDLE, flags DWORD) HMODULE

	LoadModule func(
		moduleName AString, parameterBlock *VOID) DWORD

	LoadResource func(
		module HMODULE, resInfo HRSRC) HGLOBAL

	LocalAlloc func(
		flags UINT, bytes SIZE_T) HLOCAL

	LocalCompact func(minFree UINT) SIZE_T

	LocalFileTimeToFileTime func(
		localFileTime, fileTime *FILETIME) BOOL

	LocalFlags func(mem HLOCAL) UINT

	LocalFree func(mem HLOCAL) HLOCAL

	LocalHandle func(mem *VOID) HLOCAL

	LocalLock func(mem HLOCAL) *VOID

	LocalReAlloc func(
		mem HLOCAL, bytes SIZE_T, flags UINT) HLOCAL

	LocalShrink func(mem HLOCAL, newSize UINT) SIZE_T

	LocalSize func(mem HLOCAL) SIZE_T

	LocalUnlock func(mem HLOCAL) BOOL

	LockFile func(
		file HANDLE,
		fileOffsetLow, fileOffsetHigh,
		bytesToLockLow, bytesToLockHigh DWORD) BOOL

	LockFileEx func(
		file HANDLE,
		flags, reserved,
		bytesToLockLow, bytesToLockHigh DWORD,
		overlapped *OVERLAPPED) BOOL

	LockResource func(resData HGLOBAL) *VOID

	LogonUser func(
		username, domain, password VString,
		logonType, logonProvider DWORD,
		token *HANDLE) BOOL

	LogonUserEx func(
		username, domain, password VString,
		logonType, logonProvider DWORD,
		token *HANDLE,
		logonSid **SID,
		profileBuffer **VOID,
		profileLength *DWORD,
		quotaLimits *QUOTA_LIMITS) BOOL

	LookupAccountName func(
		systemName VString,
		accountName OVString,
		sid *SID,
		nSid *DWORD,
		referencedDomainName OVString,
		nRefDomName *DWORD,
		use *SID_NAME_USE) BOOL

	LookupAccountSid func(
		systemName VString,
		sid *SID,
		name OVString,
		nName *DWORD,
		referencedDomainName OVString,
		nRefDomName *DWORD,
		use *SID_NAME_USE) BOOL

	LookupPrivilegeDisplayName func(
		systemName, name VString,
		displayName OVString,
		nDispName, languageId *DWORD) BOOL

	LookupPrivilegeName func(
		systemName VString,
		luid *LUID,
		name OVString,
		nName *DWORD) BOOL

	LookupPrivilegeValue func(
		systemName, name VString, luid *LUID) BOOL

	Lopen func(pathName AString, readWrite int) HFILE

	Lread func(file HFILE, buffer *VOID, bytes UINT) UINT

	Lstrcat func(string_1, string_2 VString) VString

	Lstrcmp func(string_1, string_2 VString) int

	Lstrcmpi func(string_1, string_2 VString) int

	Lstrcpy func(string_1, string_2 VString) VString

	Lstrcpyn func(
		string_1, string_2 VString,
		maxLength int) VString

	Lstrlen func(string_ VString) int

	Lwrite func(file HFILE, buffer CH, bytes UINT) UINT

	MakeAbsoluteSD func(
		selfRelativeSecurityDescriptor,
		absoluteSecurityDescriptor *SECURITY_DESCRIPTOR,
		absoluteSecurityDescriptorSize *DWORD,
		dacl *ACL, daclSize *DWORD,
		sacl *ACL, saclSize *DWORD,
		owner *SID, ownerSize *DWORD,
		primaryGroup *SID, primaryGroupSize *DWORD) BOOL

	MakeAbsoluteSD2 func(
		selfRelativeSecurityDescriptor *SECURITY_DESCRIPTOR,
		bufferSize *DWORD) BOOL

	MakeSelfRelativeSD func(
		absoluteSecurityDescriptor,
		selfRelativeSecurityDescriptor *SECURITY_DESCRIPTOR,
		bufferLength *DWORD) BOOL

	MapGenericMask func(
		accessMask *DWORD, genericMapping *GENERIC_MAPPING)

	MapUserPhysicalPages func(
		virtualAddress *VOID,
		numberOfPages ULONG_PTR,
		pageArray *ULONG_PTR) BOOL

	MapUserPhysicalPagesScatter func(
		virtualAddresses **VOID,
		numberOfPages ULONG_PTR,
		pageArray *ULONG_PTR) BOOL

	MapViewOfFile func(
		fileMappingObject HANDLE,
		desiredAccess, fileOffsetHigh, fileOffsetLow DWORD,
		numberOfBytesToMap SIZE_T) *VOID

	MapViewOfFileEx func(
		fileMappingObject HANDLE,
		desiredAccess, fileOffsetHigh, fileOffsetLow DWORD,
		numberOfBytesToMap SIZE_T,
		baseAddress *VOID) *VOID

	MoveFile func(
		existingFileName, newFileName VString) BOOL

	MoveFileEx func(
		existingFileName, newFileName VString,
		flags DWORD) BOOL

	MoveFileWithProgress func(
		existingFileName,
		newFileName VString,
		progressRoutine PROGRESS_ROUTINE,
		data *VOID,
		flags DWORD) BOOL

	MulDiv func(number, numerator, denominator int) int

	NeedCurrentDirectoryForExePath func(
		exeName VString) BOOL

	NotifyChangeEventLog func(
		eventLog HANDLE, event HANDLE) BOOL

	ObjectCloseAuditAlarm func(
		subsystemName VString,
		handleId *VOID,
		generateOnClose BOOL) BOOL

	ObjectDeleteAuditAlarm func(
		subsystemName VString,
		handleId *VOID,
		generateOnClose BOOL) BOOL

	ObjectOpenAuditAlarm func(
		subsystemName VString,
		handleId *VOID,
		objectTypeName, objectName VString,
		securityDescriptor *SECURITY_DESCRIPTOR,
		clientToken HANDLE,
		desiredAccess, grantedAccess DWORD,
		privileges *PRIVILEGE_SET,
		objectCreation, accessGranted BOOL,
		generateOnClose *BOOL) BOOL

	ObjectPrivilegeAuditAlarm func(
		subsystemName VString,
		handleId *VOID,
		clientToken HANDLE,
		desiredAccess DWORD,
		privileges *PRIVILEGE_SET,
		accessGranted BOOL) BOOL

	OpenBackupEventLog func(
		uncServerName, fileName VString) HANDLE

	OpenEncryptedFileRaw func(
		fileName VString, flags ULONG, context **VOID) DWORD

	OpenEvent func(
		desiredAccess DWORD,
		inheritHandle BOOL,
		name VString) HANDLE

	OpenEventLog func(
		uncServerName, sourceName VString) HANDLE

	OpenFile func(
		fileName AString,
		reOpenBuff *OFSTRUCT,
		style UINT) HFILE

	OpenFileMapping func(
		desiredAccess DWORD,
		inheritHandle BOOL,
		name VString) HANDLE

	OpenJobObject func(
		desiredAccess DWORD,
		inheritHandle BOOL,
		name VString) HANDLE

	OpenMutex func(
		desiredAccess DWORD,
		inheritHandle BOOL,
		name VString) HANDLE

	OpenProcess func(
		desiredAccess DWORD,
		inheritHandle BOOL,
		processId DWORD) HANDLE

	OpenProcessToken func(
		processHandle HANDLE,
		desiredAccess DWORD,
		tokenHandle *HANDLE) BOOL

	OpenSemaphore func(
		desiredAccess DWORD,
		inheritHandle BOOL,
		name VString) HANDLE

	OpenThread func(
		desiredAccess DWORD,
		inheritHandle BOOL,
		threadId DWORD) HANDLE

	OpenThreadToken func(
		threadHandle HANDLE,
		desiredAccess DWORD,
		openAsSelf BOOL,
		tokenHandle *HANDLE) BOOL

	OpenWaitableTimer func(
		desiredAccess DWORD,
		inheritHandle BOOL,
		timerName VString) HANDLE

	OutputDebugString func(
		outputString VString)

	PeekNamedPipe func(
		namedPipe HANDLE,
		buffer *VOID,
		bufferSize DWORD,
		bytesRead, totalBytesAvail *DWORD,
		bytesLeftThisMessage *DWORD) BOOL

	PostQueuedCompletionStatus func(
		completionPort HANDLE,
		numberOfBytesTransferred DWORD,
		completionKey ULONG_PTR,
		overlapped *OVERLAPPED) BOOL

	PrepareTape func(
		device HANDLE,
		operation DWORD,
		immediate BOOL) DWORD

	PrivilegeCheck func(
		clientToken HANDLE,
		requiredPrivileges *PRIVILEGE_SET,
		result *BOOL) BOOL

	PrivilegedServiceAuditAlarm func(
		subsystemName, serviceName VString,
		clientToken HANDLE,
		privileges *PRIVILEGE_SET,
		accessGranted BOOL) BOOL

	ProcessIdToSessionId func(
		processId DWORD,
		sessionId *DWORD) BOOL

	PulseEvent func(event HANDLE) BOOL

	PurgeComm func(file HANDLE, flags DWORD) BOOL

	QueryActCtxW func(
		flags DWORD,
		actCtx HANDLE,
		subInstance *VOID,
		infoClass ULONG,
		buffer *VOID,
		size SIZE_T,
		writtenOrRequired *SIZE_T) BOOL

	QueryDepthSList func(listHead *SLIST_HEADER) USHORT

	QueryDosDevice func(
		deviceName VString,
		targetPath OVString,
		max DWORD) DWORD

	QueryInformationJobObject func(
		job HANDLE,
		class JOBOBJECTINFOCLASS,
		info *VOID,
		length DWORD,
		returnLength *DWORD) BOOL

	QueryMemoryResourceNotification func(
		notificationHandle HANDLE, state *BOOL) BOOL

	QueryPerformanceCounter func(
		performanceCount *LARGE_INTEGER) BOOL

	QueryPerformanceFrequency func(
		frequency *LARGE_INTEGER) BOOL

	QueueUserAPC func(
		apc APCFUNC, thread HANDLE, data ULONG_PTR) DWORD

	QueueUserWorkItem func(
		function THREAD_START_ROUTINE,
		context *VOID,
		flags ULONG) BOOL

	RaiseException func(
		exceptionCode,
		exceptionFlags,
		numberOfArguments DWORD,
		arguments *ULONG_PTR)

	ReadDirectoryChangesW func(
		directory HANDLE,
		buffer *VOID,
		bufferLength DWORD,
		watchSubtree BOOL,
		notifyFilter DWORD,
		bytesReturned *DWORD,
		overlapped *OVERLAPPED,
		completionRoutine OVERLAPPED_COMPLETION_ROUTINE) BOOL

	ReadEncryptedFileRaw func(
		exportCallback FE_EXPORT_FUNC,
		callbackContext,
		context *VOID) DWORD

	ReadEventLog func(
		eventLog HANDLE,
		readFlags, recordOffset DWORD,
		buffer *VOID,
		numberOfBytesToRead DWORD,
		bytesRead *DWORD,
		minNumberOfBytesNeeded *DWORD) BOOL

	ReadFile func(
		file HANDLE,
		buffer *VOID,
		numberOfBytesToRead DWORD,
		numberOfBytesRead *DWORD,
		overlapped *OVERLAPPED) BOOL

	ReadFileEx func(
		file HANDLE,
		buffer *VOID,
		numberOfBytesToRead DWORD,
		overlapped *OVERLAPPED,
		completionRoutine OVERLAPPED_COMPLETION_ROUTINE) BOOL

	ReadFileScatter func(
		file HANDLE,
		segmentArray *FILE_SEGMENT_ELEMENT, // []
		numberOfBytesToRead DWORD,
		reserved *DWORD,
		overlapped *OVERLAPPED) BOOL

	ReadProcessMemory func(
		process HANDLE,
		baseAddress, buffer *VOID,
		size SIZE_T,
		numberOfBytesRead *SIZE_T) BOOL

	RegisterEventSource func(
		uncServerName, sourceName VString) HANDLE

	RegisterWaitForSingleObject func(
		newWaitObject *HANDLE,
		object HANDLE,
		callback WAITORTIMERCALLBACK,
		context *VOID,
		milliseconds, flags ULONG) BOOL

	RegisterWaitForSingleObjectEx func(
		object HANDLE,
		callback WAITORTIMERCALLBACK,
		context *VOID,
		milliseconds, flags ULONG) HANDLE

	ReleaseActCtx func(actCtx HANDLE)

	ReleaseMutex func(mutex HANDLE) BOOL

	ReleaseSemaphore func(
		semaphore HANDLE,
		releaseCount LONG,
		previousCount *LONG) BOOL

	RemoveDirectory func(pathName VString) BOOL

	RemoveVectoredContinueHandler func(Handle *VOID) ULONG

	RemoveVectoredExceptionHandler func(Handle *VOID) ULONG

	ReOpenFile func(
		originalFile HANDLE,
		desiredAccess, shareMode,
		flagsAndAttributes DWORD) HANDLE

	ReplaceFile func(
		replacedFileName,
		replacementFileName,
		backupFileName VString,
		replaceFlags DWORD,
		exclude, reserved *VOID) BOOL

	ReportEvent func(
		eventLog HANDLE,
		type_, category WORD,
		eventID DWORD,
		userSid *SID,
		numStrings WORD,
		dataSize DWORD,
		string_s OMVString,
		rawData *VOID) BOOL

	RequestDeviceWakeup func(device HANDLE) BOOL

	RequestWakeupLatency func(latency LATENCY_TIME) BOOL

	ResetEvent func(event HANDLE) BOOL

	ResetWriteWatch func(
		baseAddress *VOID, regionSize SIZE_T) UINT

	RestoreLastError func(errCode DWORD)

	ResumeThread func(thread HANDLE) DWORD

	RevertToSelf func() BOOL

	SearchPath func(
		path, fileName, extension VString,
		bufferLength DWORD,
		buffer OVString,
		filePart *VString) DWORD

	SetAclInformation func(
		acl *ACL,
		aclInformation *VOID,
		aclInformationLength DWORD,
		aclInformationClass ACL_INFORMATION_CLASS) BOOL

	SetCommBreak func(file HANDLE) BOOL

	SetCommConfig func(
		commDev HANDLE, cc *COMMCONFIG, size DWORD) BOOL

	SetCommMask func(file HANDLE, evtMask DWORD) BOOL

	SetCommState func(file HANDLE, dcb *DCB) BOOL

	SetCommTimeouts func(
		file HANDLE, commTimeouts *COMMTIMEOUTS) BOOL

	SetComputerName func(computerName VString) BOOL

	SetComputerNameEx func(
		nameType COMPUTER_NAME_FORMAT,
		buffer VString) BOOL

	SetCriticalSectionSpinCount func(
		criticalSection *CRITICAL_SECTION,
		spinCount DWORD) DWORD

	SetCurrentDirectory func(pathName VString) BOOL

	SetDefaultCommConfig func(
		name VString, cc *COMMCONFIG, size DWORD) BOOL

	SetDllDirectory func(pathName VString) BOOL

	SetEndOfFile func(file HANDLE) BOOL

	SetEnvironmentStrings func(newEnvironment MVString) BOOL

	SetEnvironmentVariable func(name, value VString) BOOL

	SetErrorMode func(mode UINT) UINT

	SetEvent func(event HANDLE) BOOL

	SetFileApisToANSI func()

	SetFileApisToOEM func()

	SetFileAttributes func(name VString, attributes DWORD) BOOL

	SetFilePointer func(
		file HANDLE,
		distanceToMove LONG,
		distanceToMoveHigh *LONG,
		moveMethod DWORD) DWORD

	SetFilePointerEx func(
		file HANDLE,
		distanceToMove LARGE_INTEGER,
		newFilePointer *LARGE_INTEGER,
		moveMethod DWORD) BOOL

	SetFileSecurity func(
		fileName VString,
		info SECURITY_INFORMATION,
		descriptor *SECURITY_DESCRIPTOR) BOOL

	SetFileShortName func(
		file HANDLE, shortName VString) BOOL

	SetFileTime func(
		file HANDLE,
		creationTime, lastAccessTime,
		lastWriteTime *FILETIME) BOOL

	SetFileValidData func(
		file HANDLE, validDataLength LONGLONG) BOOL

	SetFirmwareEnvironmentVariable func(
		name, guid VString, value *VOID, size DWORD) BOOL

	SetHandleCount func(number UINT) UINT

	SetHandleInformation func(
		object HANDLE, mask DWORD, flags DWORD) BOOL

	SetInformationJobObject func(
		job HANDLE,
		jobObjectInformationClass JOBOBJECTINFOCLASS,
		jobObjectInformation *VOID,
		jobObjectInformationLength DWORD) BOOL

	SetKernelObjectSecurity func(
		handle HANDLE,
		securityInformation SECURITY_INFORMATION,
		securityDescriptor *SECURITY_DESCRIPTOR) BOOL

	SetLastError func(errCode DWORD)

	SetLocalTime func(systemTime *SYSTEMTIME) BOOL

	SetMailslotInfo func(
		mailslot HANDLE,
		readTimeout DWORD) BOOL

	SetMessageWaitingIndicator func(
		msgIndicator HANDLE,
		msgCount ULONG) BOOL

	SetNamedPipeHandleState func(
		namedPipe HANDLE,
		mode, maxCollectionCount,
		collectDataTimeout *DWORD) BOOL

	SetPriorityClass func(
		process HANDLE, priorityClass DWORD) BOOL

	SetPrivateObjectSecurity func(
		securityInformation SECURITY_INFORMATION,
		modificationDescriptor *SECURITY_DESCRIPTOR,
		objectsSecurityDescriptor **SECURITY_DESCRIPTOR,
		genericMapping *GENERIC_MAPPING,
		token HANDLE) BOOL

	SetPrivateObjectSecurityEx func(
		securityInformation SECURITY_INFORMATION,
		modificationDescriptor *SECURITY_DESCRIPTOR,
		objectsSecurityDescriptor **SECURITY_DESCRIPTOR,
		autoInheritFlags ULONG,
		genericMapping *GENERIC_MAPPING,
		token HANDLE) BOOL

	SetProcessAffinityMask func(
		process HANDLE, processAffinityMask DWORD_PTR) BOOL

	SetProcessPriorityBoost func(
		process HANDLE, disablePriorityBoost BOOL) BOOL

	SetProcessShutdownParameters func(
		level, flags DWORD) BOOL

	SetProcessWorkingSetSize func(
		process HANDLE,
		minWorkingSetSize, maxWorkingSetSize SIZE_T) BOOL

	SetProcessWorkingSetSizeEx func(
		process HANDLE,
		minWorkingSetSize, maxWorkingSetSize SIZE_T,
		flags DWORD) BOOL

	SetSecurityDescriptorControl func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		controlBitsOfInterest,
		controlBitsToSet SECURITY_DESCRIPTOR_CONTROL) BOOL

	SetSecurityDescriptorDacl func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		daclPresent BOOL,
		dacl *ACL,
		daclDefaulted BOOL) BOOL

	SetSecurityDescriptorGroup func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		group *SID,
		groupDefaulted BOOL) BOOL

	SetSecurityDescriptorOwner func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		owner *SID,
		ownerDefaulted BOOL) BOOL

	SetSecurityDescriptorRMControl func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		rmControl *WChar) DWORD

	SetSecurityDescriptorSacl func(
		securityDescriptor *SECURITY_DESCRIPTOR,
		saclPresent BOOL,
		sacl *ACL,
		saclDefaulted BOOL) BOOL

	SetStdHandle func(stdHandle DWORD, Handle HANDLE) BOOL

	SetSystemFileCacheSize func(
		minFileCacheSize, maxFileCacheSize SIZE_T,
		flags DWORD) BOOL

	SetSystemPowerState func(suspend, force BOOL) BOOL

	SetSystemTime func(systemTime *SYSTEMTIME) BOOL

	SetSystemTimeAdjustment func(
		timeAdjustment DWORD,
		timeAdjustmentDisabled BOOL) BOOL

	SetTapeParameters func(
		device HANDLE,
		operation DWORD,
		tapeInformation *VOID) DWORD

	SetTapePosition func(
		device HANDLE,
		positionMethod,
		partition,
		offsetLow,
		offsetHigh DWORD,
		immediate BOOL) DWORD

	SetThreadAffinityMask func(
		thread HANDLE,
		affinityMask DWORD_PTR) DWORD_PTR

	SetThreadContext func(
		thread HANDLE, context *CONTEXT) BOOL

	SetThreadExecutionState func(
		flags EXECUTION_STATE) EXECUTION_STATE

	SetThreadIdealProcessor func(
		thread HANDLE, idealProcessor DWORD) DWORD

	SetThreadPriority func(
		thread HANDLE, priority int) BOOL

	SetThreadPriorityBoost func(
		thread HANDLE, disablePriorityBoost BOOL) BOOL

	SetThreadStackGuarantee func(
		stackSizeInBytes *ULONG) BOOL

	SetThreadToken func(
		thread *HANDLE, token HANDLE) BOOL

	SetTimerQueueTimer func(
		timerQueue HANDLE,
		callback WAITORTIMERCALLBACK,
		parameter *VOID,
		dueTime, period DWORD,
		preferIo BOOL) HANDLE

	SetTimeZoneInformation func(
		timeZoneInformation *TIME_ZONE_INFORMATION) BOOL

	SetTokenInformation func(
		tokenHandle HANDLE,
		tokenInformationClass TOKEN_INFORMATION_CLASS,
		tokenInformation *VOID,
		tokenInformationLength DWORD) BOOL

	SetUnhandledExceptionFilter func(
		tef *TOP_LEVEL_EXCEPTION_FILTER) *TOP_LEVEL_EXCEPTION_FILTER

	SetupComm func(
		file HANDLE, inQueue, outQueue DWORD) BOOL

	SetVolumeLabel func(
		rootPathName, volumeName VString) BOOL

	SetVolumeMountPoint func(
		volumeMountPoint, volumeName VString) BOOL

	SetWaitableTimer func(
		timer HANDLE,
		dueTime *LARGE_INTEGER,
		period LONG,
		completionRoutine TIMERAPCROUTINE,
		argToCompletionRoutine *VOID,
		resume BOOL) BOOL

	SignalObjectAndWait func(
		objectToSignal, objectToWaitOn HANDLE,
		milliseconds DWORD,
		alertable BOOL) DWORD

	SizeofResource func(module HMODULE, resInfo HRSRC) DWORD

	Sleep func(milliseconds DWORD)

	SleepEx func(milliseconds DWORD, alertable BOOL) DWORD

	SuspendThread func(thread HANDLE) DWORD

	SwitchToFiber func(fiber *VOID)

	SwitchToThread func() BOOL

	SystemTimeToFileTime func(
		systemTime *SYSTEMTIME, fileTime *FILETIME) BOOL

	SystemTimeToTzSpecificLocalTime func(
		timeZoneInformation *TIME_ZONE_INFORMATION,
		universalTime, localTime *SYSTEMTIME) BOOL

	TerminateJobObject func(job HANDLE, exitCode UINT) BOOL

	TerminateProcess func(process HANDLE, exitCode UINT) BOOL

	TerminateThread func(thread HANDLE, exitCode DWORD) BOOL

	TlsAlloc func() DWORD

	TlsFree func(tlsIndex DWORD) BOOL

	TlsGetValue func(tlsIndex DWORD) *VOID

	TlsSetValue func(tlsIndex DWORD, tlsValue *VOID) BOOL

	TransactNamedPipe func(
		namedPipe HANDLE,
		inBuffer *VOID, inBufferSize DWORD,
		outBuffer *VOID, outBufferSize DWORD,
		bytesRead *DWORD,
		overlapped *OVERLAPPED) BOOL

	TransmitCommChar func(file HANDLE, c char) BOOL

	TryEnterCriticalSection func(
		criticalSection *CRITICAL_SECTION) BOOL

	TzSpecificLocalTimeToSystemTime func(
		timeZoneInformation *TIME_ZONE_INFORMATION,
		local, universal *SYSTEMTIME) BOOL

	UnlockFile func(
		file HANDLE,
		offsetLow, offsetHigh, bytesLow, bytesHigh DWORD) BOOL

	UnlockFileEx func(
		file HANDLE,
		reserved, bytesLow, bytesHigh DWORD,
		o *OVERLAPPED) BOOL

	UnmapViewOfFile func(baseAddress *VOID) BOOL

	UnregisterWait func(waitHandle HANDLE) BOOL

	UnregisterWaitEx func(
		waitHandle, completionEvent HANDLE) BOOL

	UpdateResource func(
		update HANDLE,
		typ, name VString,
		language WORD,
		data *VOID,
		count DWORD) BOOL

	VerifyVersionInfoA func(
		versionInformation *OSVERSIONINFOEXA,
		typeMask DWORD,
		conditionMask DWORDLONG) BOOL

	VerifyVersionInfoW func(
		versionInformation *OSVERSIONINFOEXW,
		typeMask DWORD,
		conditionMask DWORDLONG) BOOL

	VirtualAlloc func(
		address *VOID,
		size SIZE_T,
		allocationType, protect DWORD) *VOID

	VirtualAllocEx func(
		process HANDLE,
		address *VOID,
		size SIZE_T,
		allocationType, protect DWORD) *VOID

	VirtualFree func(
		address *VOID, size SIZE_T, freeType DWORD) BOOL

	VirtualFreeEx func(
		process HANDLE,
		address *VOID,
		size SIZE_T,
		freeType DWORD) BOOL

	VirtualLock func(address *VOID, size SIZE_T) BOOL

	VirtualProtect func(
		address *VOID,
		size SIZE_T,
		newProtect DWORD,
		oldProtect *DWORD) BOOL

	VirtualProtectEx func(
		process HANDLE,
		address *VOID,
		size SIZE_T,
		newProtect DWORD,
		oldProtect *DWORD) BOOL

	VirtualQuery func(
		address *VOID,
		buffer *MEMORY_BASIC_INFORMATION,
		length SIZE_T) SIZE_T

	VirtualQueryEx func(
		process HANDLE,
		address *VOID,
		buffer *MEMORY_BASIC_INFORMATION,
		length SIZE_T) SIZE_T

	VirtualUnlock func(address *VOID, size SIZE_T) BOOL

	WaitCommEvent func(
		file HANDLE,
		evtMask *DWORD,
		overlapped *OVERLAPPED) BOOL

	WaitForDebugEvent func(
		debugEvent *DEBUG_EVENT,
		milliseconds DWORD) BOOL

	WaitForMultipleObjects func(
		count DWORD,
		handles *HANDLE,
		waitAll BOOL,
		milliseconds DWORD) DWORD

	WaitForMultipleObjectsEx func(
		count DWORD,
		handles *HANDLE,
		waitAll BOOL,
		milliseconds DWORD,
		alertable BOOL) DWORD

	WaitForSingleObject func(
		handle HANDLE, milliseconds DWORD) DWORD

	WaitForSingleObjectEx func(
		handle HANDLE,
		milliseconds DWORD,
		alertable BOOL) DWORD

	WaitNamedPipe func(
		namedPipeName VString,
		timeOut DWORD) BOOL

	//WinExec is obsolete; instead use:
	//	CreateProcess
	WinExec func(cmdLine string, cmdShow UINT) UINT

	Wow64DisableWow64FsRedirection func(
		oldValue **VOID) BOOL

	Wow64EnableWow64FsRedirection func(
		wow64FsEnableRedirection BOOLEAN) BOOLEAN

	Wow64RevertWow64FsRedirection func(olValue *VOID) BOOL

	WriteEncryptedFileRaw func(
		importCallback FE_IMPORT_FUNC,
		callbackContext, context *VOID) DWORD

	WriteFile func(
		file HANDLE,
		buffer *VOID,
		numberOfBytesToWrite DWORD,
		numberOfBytesWritten *DWORD,
		overlapped *OVERLAPPED) BOOL

	WriteFileEx func(
		file HANDLE,
		buffer *VOID,
		numberOfBytesToWrite DWORD,
		overlapped *OVERLAPPED,
		completionRoutine OVERLAPPED_COMPLETION_ROUTINE) BOOL

	WriteFileGather func(
		file HANDLE,
		segmentArray *FILE_SEGMENT_ELEMENT, // []
		numberOfBytesToWrite DWORD,
		reserved *DWORD,
		overlapped *OVERLAPPED) BOOL

	//WritePrivateProfileSection is obsolete; instead use:
	//	registry
	WritePrivateProfileSection func(
		appName, string_, fileName VString) BOOL

	//WritePrivateProfileString is obsolete; instead use:
	//	registry
	WritePrivateProfileString func(
		appName, keyName, string_, fileName VString) BOOL

	//WritePrivateProfileStruct is obsolete; instead use:
	//	registry
	WritePrivateProfileStruct func(
		section, key VString,
		struct_ *VOID,
		sizeStruct UINT,
		file VString) BOOL

	WriteProcessMemory func(
		process HANDLE,
		baseAddress, buffer *VOID,
		size SIZE_T,
		numberOfBytesWritten *SIZE_T) BOOL

	//WriteProfileSection is obsolete; instead use:
	//	registry
	WriteProfileSection func(
		appName, string_ VString) BOOL

	//WriteProfileString is obsolete; instead use:
	//	registry
	WriteProfileString func(
		appName, keyName, string_ VString) BOOL

	WriteTapemark func(
		device HANDLE,
		tapemarkType, tapemarkCount DWORD,
		immediate BOOL) DWORD

	WTSGetActiveConsoleSessionId func() DWORD

	ZombifyActCtx func(actCtx HANDLE) BOOL

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

var WinBaseApis = Apis{
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

var WinBaseUnicodeApis = Apis{
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

var WinBaseANSIApis = Apis{
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
