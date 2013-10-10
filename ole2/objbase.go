package ole2

import (
	. "github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/ole32"
)

var (
	CoBuildVersion func() T.DWORD

	CoInitialize func(reserved *T.VOID)

	CoUninitialize func()

	CoGetMalloc func(memCtx T.DWORD, mAlloc **T.IMalloc)

	CoGetCurrentProcess func() T.DWORD

	CoRegisterMallocSpy func(spy *T.IMallocSpy)

	CoRevokeMallocSpy func()

	CoCreateStandardMalloc func(memCtx T.DWORD, mAlloc **T.IMalloc)

	CoInitializeEx func(reserved *T.VOID, coInit T.DWORD)

	CoGetCallerTID func(tid *T.DWORD)

	CoRegisterInitializeSpy func(
		spy *T.IInitializeSpy, cookie *T.ULARGE_INTEGER)

	CoRevokeInitializeSpy func(cookie T.ULARGE_INTEGER)

	CoGetContextToken func(token *T.ULONG_PTR)

	CoGetSystemSecurityPermissions func(
		sdType T.COMSD, sd **T.SECURITY_DESCRIPTOR)

	//NOTE(t):Not implemented in non-debug dll
	//DebugCoGetRpcFault func(void) T.ULONG

	//NOTE(t):Not implemented in non-debug dll
	//DebugCoSetRpcFault func(T.ULONG) void

	CoGetObjectContext func(riid T.REFIID, v **T.VOID)

	CoGetClassObject func(
		id T.REFCLSID,
		ctx T.DWORD,
		reserved *T.VOID,
		riid T.REFIID,
		v **T.VOID)

	CoRegisterClassObject func(
		id T.REFCLSID,
		unk *T.IUnknown,
		ctx T.DWORD,
		flags T.DWORD,
		register *T.DWORD)

	CoRevokeClassObject func(register T.DWORD)

	CoResumeClassObjects func()

	CoSuspendClassObjects func()

	CoAddRefServerProcess func() T.ULONG

	CoReleaseServerProcess func() T.ULONG

	CoGetPSClsid func(riid T.REFIID, clsid *T.CLSID)

	CoRegisterPSClsid func(riid T.REFIID, rclsid T.REFCLSID)

	CoRegisterSurrogate func(surrogate *T.ISurrogate)

	CoGetMarshalSizeMax func(
		size *T.ULONG,
		riid T.REFIID,
		unk *T.IUnknown,
		destContextSize T.DWORD,
		destContext *T.VOID,
		mshlflags T.DWORD)

	CoMarshalInterface func(
		stm *T.STREAM,
		riid T.REFIID,
		unk *T.IUnknown,
		destContextSize T.DWORD,
		destContext *T.VOID,
		mshlflags T.DWORD)

	CoUnmarshalInterface func(
		stm *T.STREAM,
		riid T.REFIID,
		ppv **T.VOID)

	CoMarshalHresult func(
		pstm *T.STREAM,
		hresult T.HRESULT)

	CoUnmarshalHresult func(
		pstm *T.STREAM,
		phresult *T.HRESULT)

	CoReleaseMarshalData func(
		stm *T.STREAM)

	CoDisconnectObject func(
		unk *T.IUnknown,
		reserved T.DWORD)

	CoLockObjectExternal func(
		unk *T.IUnknown,
		fLock T.BOOL,
		fLastUnlockReleases T.BOOL)

	CoGetStandardMarshal func(
		riid T.REFIID,
		unk *T.IUnknown,
		destContextSize T.DWORD,
		destContext *T.VOID,
		mshlflags T.DWORD,
		marshal **T.IMarshal)

	CoGetStdMarshalEx func(
		unkOuter *T.IUnknown,
		smexflags T.DWORD,
		unkInner **T.IUnknown)

	CoIsHandlerConnected func(unk *T.IUnknown) T.BOOL

	CoMarshalInterThreadInterfaceInStream func(
		riid T.REFIID, unk *T.IUnknown, stm **T.STREAM)

	CoGetInterfaceAndReleaseStream func(
		stm *T.STREAM, iid T.REFIID, ppv **T.VOID)

	CoCreateFreeThreadedMarshaler func(
		outer *T.IUnknown, marshal **T.IUnknown)

	CoLoadLibrary func(libName *T.OLESTR, autoFree T.BOOL) T.HINSTANCE

	CoFreeLibrary func(inst T.HINSTANCE)

	CoFreeAllLibraries func()

	CoFreeUnusedLibraries func()

	CoFreeUnusedLibrariesEx func(unloadDelay, reserved T.DWORD)

	CoInitializeSecurity func(
		secDesc *T.SECURITY_DESCRIPTOR,
		authSvc T.LONG,
		asAuthSvc *T.SOLE_AUTHENTICATION_SERVICE,
		reserved1 *T.VOID,
		authnLevel T.DWORD,
		impLevel T.DWORD,
		authList *T.VOID,
		capabilities T.DWORD,
		reserved3 *T.VOID)

	CoGetCallContext func(
		riid T.REFIID,
		intrface **T.VOID)

	CoQueryProxyBlanket func(
		proxy *T.IUnknown,
		pwAuthnSvc *T.DWORD,
		authzSvc *T.DWORD,
		serverPrincName **T.OLECHAR,
		authnLevel *T.DWORD,
		impLevel *T.DWORD,
		authInfo *T.RPC_AUTH_IDENTITY_HANDLE,
		capabilites *T.DWORD)

	CoSetProxyBlanket func(
		proxy *T.IUnknown,
		authnSvc T.DWORD,
		authzSvc T.DWORD,
		serverPrincName *T.OLECHAR,
		authnLevel T.DWORD,
		impLevel T.DWORD,
		authInfo T.RPC_AUTH_IDENTITY_HANDLE,
		capabilities T.DWORD)

	CoCopyProxy func(proxy *T.IUnknown, cpy **T.IUnknown)

	CoQueryClientBlanket func(
		authnSvc, authzSvc *T.DWORD,
		serverPrincName **T.OLECHAR,
		authnLevel, impLevel *T.DWORD,
		privs *T.RPC_AUTHZ_HANDLE,
		capabilities *T.DWORD)

	CoImpersonateClient func()

	CoRevertToSelf func()

	CoQueryAuthenticationServices func(
		pcAuthSvc *T.DWORD,
		asAuthSvc **T.SOLE_AUTHENTICATION_SERVICE)

	CoSwitchCallContext func(
		newObject *T.IUnknown,
		oldObject **T.IUnknown)

	CoCreateInstance func(
		rclsid T.REFCLSID,
		unkOuter *T.IUnknown,
		clsContext T.DWORD,
		riid T.REFIID,
		ppv **T.VOID)

	CoGetInstanceFromFile func(
		serverInfo *T.COSERVERINFO,
		clsid *T.CLSID,
		outer *T.IUnknown,
		clsCtx T.DWORD,
		grfMode T.DWORD,
		pwszName *T.OLECHAR,
		count T.DWORD,
		results *T.MULTI_QI)

	CoGetInstanceFromIStorage func(
		serverInfo *T.COSERVERINFO,
		clsid *T.CLSID,
		outer *T.IUnknown,
		clsCtx T.DWORD,
		pstg *T.IStorage,
		count T.DWORD,
		results *T.MULTI_QI)

	CoCreateInstanceEx func(
		clsid T.REFCLSID,
		outer *T.IUnknown,
		// only relevant locally
		clsCtx T.DWORD,
		serverInfo *T.COSERVERINFO,
		count T.DWORD,
		results *T.MULTI_QI)

	CoGetCancelObject func(
		ThreadId T.DWORD,
		iid T.REFIID,
		unk **T.VOID)

	CoSetCancelObject func(unk *T.IUnknown)

	CoCancelCall func(ThreadId T.DWORD, ulTimeout T.ULONG)

	CoTestCancel func()

	CoEnableCallCancellation func(Reserved *T.VOID)

	CoDisableCallCancellation func(Reserved *T.VOID)

	CoAllowSetForegroundWindow func(
		unk *T.IUnknown, lpvReserved *T.VOID)

	DcomChannelSetHResult func(
		Reserved1 *T.VOID,
		Reserved2 *T.ULONG,
		appsHR T.HRESULT)

	StringFromCLSID func(
		rclsid T.REFCLSID, lplpsz **T.OLESTR)

	CLSIDFromString func(lpsz *T.OLESTR, pclsid *T.CLSID)

	StringFromIID func(rclsid T.REFIID, lplpsz **T.OLESTR)

	IIDFromString func(sz *T.OLESTR, iid *T.IID)

	CoIsOle1Class func(rclsid T.REFCLSID) T.BOOL

	ProgIDFromCLSID func(clsid T.REFCLSID, lplpszProgID **T.OLESTR)

	CLSIDFromProgID func(szProgID *T.OLESTR, clsid *T.CLSID)

	CLSIDFromProgIDEx func(szProgID *T.OLESTR, clsid *T.CLSID)

	StringFromGUID2 func(
		rguid T.REFGUID, lpsz *T.OLESTR, cchMax int) int

	CoCreateGuid func(pguid *T.GUID)

	CoFileTimeToDosDateTime func(
		FileTime *T.FILETIME, DosDate, DosTime *T.WORD) T.BOOL

	CoDosDateTimeToFileTime func(
		DosDate, DosTime T.WORD, FileTime *T.FILETIME) T.BOOL

	CoFileTimeNow func(FileTime *T.FILETIME)

	CoRegisterMessageFilter func(MessageFilter *T.IMessageFilter,
		oMessageFilter **T.IMessageFilter)

	CoRegisterChannelHook func(ExtensionUuid T.REFGUID,
		ChannelHook *T.IChannelHook)

	CoWaitForMultipleHandles func(Flags, Timeout T.DWORD,
		nHandles T.ULONG, Handles *T.HANDLE, lpdwindex *T.DWORD)

	CoInvalidateRemoteMachineBindings func(
		MachineName *T.OLESTR)

	CoGetTreatAsClass func(old T.REFCLSID, new *T.CLSID)

	CoTreatAsClass func(Old, New T.REFCLSID)

	DllGetClassObject func(
		rclsid T.REFCLSID, riid T.REFIID, ppv **T.VOID)

	DllCanUnloadNow func()

	CoTaskMemAlloc func(cb T.SIZE_T) *T.VOID

	CoTaskMemRealloc func(pv *T.VOID, cb T.SIZE_T) *T.VOID

	CoTaskMemFree func(pv *T.VOID)

	//CreateDataAdviseHolder func(T.LPDATAADVISEHOLDER *ppDAHolder) T.DUP T.IN T.OLE2

	CreateDataCache func(
		unkOuter *T.IUnknown,
		rclsid T.REFCLSID,
		iid T.REFIID,
		ppv **T.VOID)

	StgCreateDocfile func(pwcsName *T.OLECHAR,
		grfMode, reserved T.DWORD, ppstgOpen **T.IStorage)

	StgCreateDocfileOnILockBytes func(
		plkbyt *T.ILockBytes,
		grfMode, reserved T.DWORD,
		ppstgOpen **T.IStorage)

	StgOpenStorage func(
		pwcsName *T.OLECHAR,
		pstgPriority *T.IStorage,
		grfMode T.DWORD,
		snbExclude T.SNB,
		reserved T.DWORD,
		ppstgOpen **T.IStorage)

	StgOpenStorageOnILockBytes func(
		plkbyt *T.ILockBytes,
		pstgPriority *T.IStorage,
		grfMode T.DWORD,
		snbExclude T.SNB,
		reserved T.DWORD,
		ppstgOpen **T.IStorage)

	StgIsStorageFile func(name *T.OLECHAR)

	StgIsStorageILockBytes func(lkByt *T.ILockBytes)

	StgSetTimes func(
		name *T.OLECHAR, cTime, aTime, mTime *T.FILETIME)

	StgOpenAsyncDocfileOnIFillLockBytes func(
		pflb *T.IFillLockBytes,
		grfMode T.DWORD,
		asyncFlags T.DWORD,
		ppstgOpen **T.IStorage)

	StgGetIFillLockBytesOnILockBytes func(
		pilb *T.ILockBytes,
		ppflb **T.IFillLockBytes)

	StgGetIFillLockBytesOnFile func(
		pwcsName *T.OLECHAR,
		ppflb **T.IFillLockBytes)

	StgOpenLayoutDocfile func(
		pwcsDfName *T.OLECHAR,
		grfMode T.DWORD,
		reserved T.DWORD,
		ppstgOpen **T.IStorage)

	StgCreateStorageEx func(
		pwcsName *T.WChar,
		// TODO(t):WString

		grfMode T.DWORD,
		stgfmt T.DWORD,
		grfAttrs T.DWORD,
		StgOptions *T.STGOPTIONS,
		reserved *T.VOID,
		riid T.REFIID,
		ObjectOpen **T.VOID)

	StgOpenStorageEx func(
		pwcsName *T.WChar,
		// TODO(t):WString

		grfMode T.DWORD,
		stgfmt T.DWORD,
		grfAttrs T.DWORD,
		StgOptions *T.STGOPTIONS,
		reserved *T.VOID,
		riid T.REFIID,
		ObjectOpen **T.VOID)

	BindMoniker func(
		pmk *T.IMoniker,
		grfOpt T.DWORD,
		iidResult T.REFIID,
		ppvResult **T.VOID)

	CoInstall func(
		pbc *T.IBindCtx,
		Flags T.DWORD,
		ClassSpec *T.UCLSSPEC,
		Query *T.QUERYCONTEXT,
		CodeBase T.WString)

	CoGetObject func(
		Name T.WString,
		BindOptions *T.BIND_OPTS,
		riid T.REFIID,
		ppv **T.VOID)

	MkParseDisplayName func(
		pbc *T.IBindCtx,
		szUserName *T.OLESTR,
		pchEaten *T.ULONG,
		ppmk **T.IMoniker)

	MonikerRelativePathTo func(pmkSrc, pmkDest *T.IMoniker,
		ppmkRelPath **T.IMoniker, Reserved T.BOOL)

	MonikerCommonPrefixWith func(
		pmkThis, pmkOther *T.IMoniker, ppmkCommon **T.IMoniker)

	CreateBindCtx func(reserved T.DWORD, ppbc **T.IBindCtx)

	CreateGenericComposite func(pmkFirst, pmkRest *T.IMoniker,
		ppmkComposite **T.IMoniker)

	GetClassFile func(szFilename *T.OLESTR, pclsid *T.CLSID)

	CreateClassMoniker func(rclsid T.REFCLSID, ppmk **T.IMoniker)

	CreateFileMoniker func(
		lpszPathName *T.OLESTR, ppmk **T.IMoniker)

	CreateItemMoniker func(
		lpszDelim, lpszItem *T.OLESTR, ppmk **T.IMoniker)

	CreateAntiMoniker func(ppmk **T.IMoniker)

	CreatePointerMoniker func(
		unk *T.IUnknown, ppmk **T.IMoniker)

	CreateObjrefMoniker func(
		unk *T.IUnknown, ppmk **T.IMoniker)

	GetRunningObjectTable func(
		reserved T.DWORD,
		pprot **T.IRunningObjectTable)

	CreateStdProgressIndicator func(
		hwndParent T.HWND,
		Title *T.OLESTR,
		IbscCaller *T.IBindStatusCallback,
		Ibsc **T.IBindStatusCallback)
)

//TODO(t):Unimplemented?
//{"CoCreateStandardMalloc", &CoCreateStandardMalloc},
//{"DebugCoGetRpcFault", &DebugCoGetRpcFault},
//{"DebugCoSetRpcFault", &DebugCoSetRpcFault},
//{"DllCanUnloadNow", &DllCanUnloadNow},
//{"StgOpenLayoutDocfile", &StgOpenLayoutDocfile},

var ObjBaseApis = Apis{
	{"BindMoniker", &BindMoniker},
	{"CLSIDFromProgID", &CLSIDFromProgID},
	{"CLSIDFromProgIDEx", &CLSIDFromProgIDEx},
	{"CLSIDFromString", &CLSIDFromString},
	{"CoAddRefServerProcess", &CoAddRefServerProcess},
	{"CoAllowSetForegroundWindow", &CoAllowSetForegroundWindow},
	{"CoBuildVersion", &CoBuildVersion},
	{"CoCancelCall", &CoCancelCall},
	{"CoCopyProxy", &CoCopyProxy},
	{"CoCreateFreeThreadedMarshaler", &CoCreateFreeThreadedMarshaler},
	{"CoCreateGuid", &CoCreateGuid},
	{"CoCreateInstance", &CoCreateInstance},
	{"CoCreateInstanceEx", &CoCreateInstanceEx},
	{"CoDisableCallCancellation", &CoDisableCallCancellation},
	{"CoDisconnectObject", &CoDisconnectObject},
	{"CoDosDateTimeToFileTime", &CoDosDateTimeToFileTime},
	{"CoEnableCallCancellation", &CoEnableCallCancellation},
	{"CoFileTimeNow", &CoFileTimeNow},
	{"CoFileTimeToDosDateTime", &CoFileTimeToDosDateTime},
	{"CoFreeAllLibraries", &CoFreeAllLibraries},
	{"CoFreeLibrary", &CoFreeLibrary},
	{"CoFreeUnusedLibraries", &CoFreeUnusedLibraries},
	{"CoFreeUnusedLibrariesEx", &CoFreeUnusedLibrariesEx},
	{"CoGetCallContext", &CoGetCallContext},
	{"CoGetCallerTID", &CoGetCallerTID},
	{"CoGetCancelObject", &CoGetCancelObject},
	{"CoGetClassObject", &CoGetClassObject},
	{"CoGetContextToken", &CoGetContextToken},
	{"CoGetCurrentProcess", &CoGetCurrentProcess},
	{"CoGetInstanceFromFile", &CoGetInstanceFromFile},
	{"CoGetInstanceFromIStorage", &CoGetInstanceFromIStorage},
	{"CoGetInterfaceAndReleaseStream", &CoGetInterfaceAndReleaseStream},
	{"CoGetMalloc", &CoGetMalloc},
	{"CoGetMarshalSizeMax", &CoGetMarshalSizeMax},
	{"CoGetObject", &CoGetObject},
	{"CoGetObjectContext", &CoGetObjectContext},
	{"CoGetPSClsid", &CoGetPSClsid},
	{"CoGetStandardMarshal", &CoGetStandardMarshal},
	{"CoGetStdMarshalEx", &CoGetStdMarshalEx},
	{"CoGetSystemSecurityPermissions", &CoGetSystemSecurityPermissions},
	{"CoGetTreatAsClass", &CoGetTreatAsClass},
	{"CoImpersonateClient", &CoImpersonateClient},
	{"CoInitialize", &CoInitialize},
	{"CoInitializeEx", &CoInitializeEx},
	{"CoInitializeSecurity", &CoInitializeSecurity},
	{"CoInstall", &CoInstall},
	{"CoInvalidateRemoteMachineBindings", &CoInvalidateRemoteMachineBindings},
	{"CoIsHandlerConnected", &CoIsHandlerConnected},
	{"CoIsOle1Class", &CoIsOle1Class},
	{"CoLoadLibrary", &CoLoadLibrary},
	{"CoLockObjectExternal", &CoLockObjectExternal},
	{"CoMarshalHresult", &CoMarshalHresult},
	{"CoMarshalInterface", &CoMarshalInterface},
	{"CoMarshalInterThreadInterfaceInStream", &CoMarshalInterThreadInterfaceInStream},
	{"CoQueryAuthenticationServices", &CoQueryAuthenticationServices},
	{"CoQueryClientBlanket", &CoQueryClientBlanket},
	{"CoQueryProxyBlanket", &CoQueryProxyBlanket},
	{"CoRegisterChannelHook", &CoRegisterChannelHook},
	{"CoRegisterClassObject", &CoRegisterClassObject},
	{"CoRegisterInitializeSpy", &CoRegisterInitializeSpy},
	{"CoRegisterMallocSpy", &CoRegisterMallocSpy},
	{"CoRegisterMessageFilter", &CoRegisterMessageFilter},
	{"CoRegisterPSClsid", &CoRegisterPSClsid},
	{"CoRegisterSurrogate", &CoRegisterSurrogate},
	{"CoReleaseMarshalData", &CoReleaseMarshalData},
	{"CoReleaseServerProcess", &CoReleaseServerProcess},
	{"CoResumeClassObjects", &CoResumeClassObjects},
	{"CoRevertToSelf", &CoRevertToSelf},
	{"CoRevokeClassObject", &CoRevokeClassObject},
	{"CoRevokeInitializeSpy", &CoRevokeInitializeSpy},
	{"CoRevokeMallocSpy", &CoRevokeMallocSpy},
	{"CoSetCancelObject", &CoSetCancelObject},
	{"CoSetProxyBlanket", &CoSetProxyBlanket},
	{"CoSuspendClassObjects", &CoSuspendClassObjects},
	{"CoSwitchCallContext", &CoSwitchCallContext},
	{"CoTaskMemAlloc", &CoTaskMemAlloc},
	{"CoTaskMemFree", &CoTaskMemFree},
	{"CoTaskMemRealloc", &CoTaskMemRealloc},
	{"CoTestCancel", &CoTestCancel},
	{"CoTreatAsClass", &CoTreatAsClass},
	{"CoUninitialize", &CoUninitialize},
	{"CoUnmarshalHresult", &CoUnmarshalHresult},
	{"CoUnmarshalInterface", &CoUnmarshalInterface},
	{"CoWaitForMultipleHandles", &CoWaitForMultipleHandles},
	{"CreateAntiMoniker", &CreateAntiMoniker},
	{"CreateBindCtx", &CreateBindCtx},
	{"CreateClassMoniker", &CreateClassMoniker},
	{"CreateDataCache", &CreateDataCache},
	{"CreateFileMoniker", &CreateFileMoniker},
	{"CreateGenericComposite", &CreateGenericComposite},
	{"CreateItemMoniker", &CreateItemMoniker},
	{"CreateObjrefMoniker", &CreateObjrefMoniker},
	{"CreatePointerMoniker", &CreatePointerMoniker},
	{"CreateStdProgressIndicator", &CreateStdProgressIndicator},
	{"DcomChannelSetHResult", &DcomChannelSetHResult},
	{"DllGetClassObject", &DllGetClassObject},
	{"GetClassFile", &GetClassFile},
	{"GetRunningObjectTable", &GetRunningObjectTable},
	{"IIDFromString", &IIDFromString},
	{"MkParseDisplayName", &MkParseDisplayName},
	{"MonikerCommonPrefixWith", &MonikerCommonPrefixWith},
	{"MonikerRelativePathTo", &MonikerRelativePathTo},
	{"ProgIDFromCLSID", &ProgIDFromCLSID},
	{"StgCreateDocfile", &StgCreateDocfile},
	{"StgCreateDocfileOnILockBytes", &StgCreateDocfileOnILockBytes},
	{"StgCreateStorageEx", &StgCreateStorageEx},
	{"StgGetIFillLockBytesOnFile", &StgGetIFillLockBytesOnFile},
	{"StgGetIFillLockBytesOnILockBytes", &StgGetIFillLockBytesOnILockBytes},
	{"StgIsStorageFile", &StgIsStorageFile},
	{"StgIsStorageILockBytes", &StgIsStorageILockBytes},
	{"StgOpenAsyncDocfileOnIFillLockBytes", &StgOpenAsyncDocfileOnIFillLockBytes},
	{"StgOpenStorage", &StgOpenStorage},
	{"StgOpenStorageEx", &StgOpenStorageEx},
	{"StgOpenStorageOnILockBytes", &StgOpenStorageOnILockBytes},
	{"StgSetTimes", &StgSetTimes},
	{"StringFromCLSID", &StringFromCLSID},
	{"StringFromGUID2", &StringFromGUID2},
	{"StringFromIID", &StringFromIID},
}
