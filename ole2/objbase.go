package ole2

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/ole32"
)

var (
	CoBuildVersion func() DWORD

	CoInitialize func(reserved *VOID)

	CoUninitialize func()

	CoGetMalloc func(memCtx DWORD, mAlloc **IMalloc)

	CoGetCurrentProcess func() DWORD

	CoRegisterMallocSpy func(spy *IMallocSpy)

	CoRevokeMallocSpy func()

	CoCreateStandardMalloc func(memCtx DWORD, mAlloc **IMalloc)

	CoInitializeEx func(reserved *VOID, coInit DWORD)

	CoGetCallerTID func(tid *DWORD)

	CoRegisterInitializeSpy func(
		spy *IInitializeSpy, cookie *ULARGE_INTEGER)

	CoRevokeInitializeSpy func(cookie ULARGE_INTEGER)

	CoGetContextToken func(token *ULONG_PTR)

	CoGetSystemSecurityPermissions func(
		sdType COMSD, sd **SECURITY_DESCRIPTOR)

	//NOTE(t):Not implemented in non-debug dll
	//DebugCoGetRpcFault func(void) ULONG

	//NOTE(t):Not implemented in non-debug dll
	//DebugCoSetRpcFault func(ULONG) void

	CoGetObjectContext func(riid REFIID, v **VOID)

	CoGetClassObject func(
		id REFCLSID,
		ctx DWORD,
		reserved *VOID,
		riid REFIID,
		v **VOID)

	CoRegisterClassObject func(
		id REFCLSID,
		unk *IUnknown,
		ctx DWORD,
		flags DWORD,
		register *DWORD)

	CoRevokeClassObject func(register DWORD)

	CoResumeClassObjects func()

	CoSuspendClassObjects func()

	CoAddRefServerProcess func() ULONG

	CoReleaseServerProcess func() ULONG

	CoGetPSClsid func(
		riid REFIID,
		clsid *CLSID)
	CoRegisterPSClsid func(
		riid REFIID,
		rclsid REFCLSID)

	CoRegisterSurrogate func(
		Surrogate *ISurrogate)

	CoGetMarshalSizeMax func(
		Size *ULONG,
		riid REFIID,
		Unk *IUnknown,
		DestContextSize DWORD,
		DestContext *VOID,
		mshlflags DWORD)
	CoMarshalInterface func(
		Stm *STREAM,
		riid REFIID,
		Unk *IUnknown,
		DestContextSize DWORD,
		DestContext *VOID,
		mshlflags DWORD)
	CoUnmarshalInterface func(
		Stm *STREAM,
		riid REFIID,
		ppv **VOID)
	CoMarshalHresult func(
		pstm *STREAM,
		hresult HRESULT)
	CoUnmarshalHresult func(
		pstm *STREAM,
		phresult *HRESULT)
	CoReleaseMarshalData func(
		Stm *STREAM)
	CoDisconnectObject func(
		Unk *IUnknown,
		Reserved DWORD)
	CoLockObjectExternal func(
		Unk *IUnknown,
		fLock BOOL,
		fLastUnlockReleases BOOL)
	CoGetStandardMarshal func(
		riid REFIID,
		Unk *IUnknown,
		DestContextSize DWORD,
		DestContext *VOID,
		mshlflags DWORD,
		Marshal **IMarshal)

	CoGetStdMarshalEx func(
		UnkOuter *IUnknown,
		smexflags DWORD,
		UnkInner **IUnknown)

	CoIsHandlerConnected func(
		Unk *IUnknown) BOOL

	CoMarshalInterThreadInterfaceInStream func(
		riid REFIID,
		Unk *IUnknown,
		Stm **STREAM)

	CoGetInterfaceAndReleaseStream func(
		Stm *STREAM,
		iid REFIID,
		ppv **VOID)

	CoCreateFreeThreadedMarshaler func(
		outer *IUnknown, marshal **IUnknown)

	CoLoadLibrary func(libName *OLESTR, autoFree BOOL) HINSTANCE

	CoFreeLibrary func(inst HINSTANCE)

	CoFreeAllLibraries func()

	CoFreeUnusedLibraries func()

	CoFreeUnusedLibrariesEx func(unloadDelay, reserved DWORD)

	CoInitializeSecurity func(
		SecDesc *SECURITY_DESCRIPTOR,
		AuthSvc LONG,
		asAuthSvc *SOLE_AUTHENTICATION_SERVICE,
		Reserved1 *VOID,
		AuthnLevel DWORD,
		ImpLevel DWORD,
		AuthList *VOID,
		Capabilities DWORD,
		Reserved3 *VOID)

	CoGetCallContext func(
		riid REFIID,
		Interface **VOID)

	CoQueryProxyBlanket func(
		Proxy *IUnknown,
		pwAuthnSvc *DWORD,
		AuthzSvc *DWORD,
		ServerPrincName **OLECHAR,
		AuthnLevel *DWORD,
		ImpLevel *DWORD,
		AuthInfo *RPC_AUTH_IDENTITY_HANDLE,
		Capabilites *DWORD)

	CoSetProxyBlanket func(
		Proxy *IUnknown,
		AuthnSvc DWORD,
		AuthzSvc DWORD,
		ServerPrincName *OLECHAR,
		AuthnLevel DWORD,
		ImpLevel DWORD,
		AuthInfo RPC_AUTH_IDENTITY_HANDLE,
		Capabilities DWORD)

	CoCopyProxy func(Proxy *IUnknown, Copy **IUnknown)

	CoQueryClientBlanket func(
		AuthnSvc, AuthzSvc *DWORD,
		ServerPrincName **OLECHAR,
		AuthnLevel, ImpLevel *DWORD,
		Privs *RPC_AUTHZ_HANDLE,
		Capabilities *DWORD)

	CoImpersonateClient func()

	CoRevertToSelf func()

	CoQueryAuthenticationServices func(
		pcAuthSvc *DWORD,
		asAuthSvc **SOLE_AUTHENTICATION_SERVICE)

	CoSwitchCallContext func(
		NewObject *IUnknown,
		OldObject **IUnknown)

	CoCreateInstance func(
		rclsid REFCLSID,
		UnkOuter *IUnknown,
		ClsContext DWORD,
		riid REFIID,
		ppv **VOID)

	CoGetInstanceFromFile func(
		ServerInfo *COSERVERINFO,
		Clsid *CLSID,
		Outer *IUnknown,
		ClsCtx DWORD,
		grfMode DWORD,
		pwszName *OLECHAR,
		Count DWORD,
		Results *MULTI_QI)

	CoGetInstanceFromIStorage func(
		ServerInfo *COSERVERINFO,
		Clsid *CLSID,
		Outer *IUnknown,
		ClsCtx DWORD,
		pstg *IStorage,
		Count DWORD,
		Results *MULTI_QI)

	CoCreateInstanceEx func(
		Clsid REFCLSID,
		Outer *IUnknown,
		// only relevant locally
		ClsCtx DWORD,
		ServerInfo *COSERVERINFO,
		Count DWORD,
		Results *MULTI_QI)

	CoGetCancelObject func(
		ThreadId DWORD,
		iid REFIID,
		Unk **VOID)

	CoSetCancelObject func(
		Unk *IUnknown)

	CoCancelCall func(
		ThreadId DWORD,
		ulTimeout ULONG)

	CoTestCancel func()

	CoEnableCallCancellation func(
		Reserved *VOID)

	CoDisableCallCancellation func(
		Reserved *VOID)

	CoAllowSetForegroundWindow func(
		Unk *IUnknown,
		lpvReserved *VOID)

	DcomChannelSetHResult func(
		Reserved1 *VOID,
		Reserved2 *ULONG,
		appsHR HRESULT)

	StringFromCLSID func(
		rclsid REFCLSID,
		lplpsz **OLESTR)
	CLSIDFromString func(
		lpsz *OLESTR,
		pclsid *CLSID)
	StringFromIID func(
		rclsid REFIID,
		lplpsz **OLESTR)
	IIDFromString func(
		sz *OLESTR,
		iid *IID)
	CoIsOle1Class func(
		rclsid REFCLSID) BOOL
	ProgIDFromCLSID func(
		clsid REFCLSID,
		lplpszProgID **OLESTR)
	CLSIDFromProgID func(
		szProgID *OLESTR,
		clsid *CLSID)
	CLSIDFromProgIDEx func(
		szProgID *OLESTR,
		clsid *CLSID)
	StringFromGUID2 func(
		rguid REFGUID,
		lpsz *OLESTR,
		cchMax int) int

	CoCreateGuid func(
		pguid *GUID)

	CoFileTimeToDosDateTime func(
		FileTime *FILETIME,
		DosDate *WORD,
		DosTime *WORD) BOOL

	CoDosDateTimeToFileTime func(
		DosDate WORD,
		DosTime WORD,
		FileTime *FILETIME) BOOL

	CoFileTimeNow func(
		FileTime *FILETIME)

	CoRegisterMessageFilter func(
		MessageFilter *IMessageFilter,
		oMessageFilter **IMessageFilter)

	CoRegisterChannelHook func(
		ExtensionUuid REFGUID,
		ChannelHook *IChannelHook)

	CoWaitForMultipleHandles func(
		Flags DWORD,
		Timeout DWORD,
		nHandles ULONG,
		Handles *HANDLE,
		lpdwindex *DWORD)

	CoInvalidateRemoteMachineBindings func(
		MachineName *OLESTR)

	CoGetTreatAsClass func(Old REFCLSID, New *CLSID)

	CoTreatAsClass func(
		Old REFCLSID,
		New REFCLSID)

	DllGetClassObject func(
		rclsid REFCLSID,
		riid REFIID,
		ppv **VOID)

	DllCanUnloadNow func()

	CoTaskMemAlloc func(
		cb SIZE_T) *VOID
	CoTaskMemRealloc func(
		pv *VOID,
		cb SIZE_T) *VOID
	CoTaskMemFree func(
		pv *VOID)

	//CreateDataAdviseHolder func(LPDATAADVISEHOLDER *ppDAHolder) DUP IN OLE2

	CreateDataCache func(
		UnkOuter *IUnknown,
		rclsid REFCLSID,
		iid REFIID,
		ppv **VOID)

	StgCreateDocfile func(
		pwcsName *OLECHAR,
		grfMode DWORD,
		reserved DWORD,
		ppstgOpen **IStorage)

	StgCreateDocfileOnILockBytes func(
		plkbyt *ILockBytes,
		grfMode DWORD,
		reserved DWORD,
		ppstgOpen **IStorage)

	StgOpenStorage func(
		pwcsName *OLECHAR,
		pstgPriority *IStorage,
		grfMode DWORD,
		snbExclude SNB,
		reserved DWORD,
		ppstgOpen **IStorage)
	StgOpenStorageOnILockBytes func(
		plkbyt *ILockBytes,
		pstgPriority *IStorage,
		grfMode DWORD,
		snbExclude SNB,
		reserved DWORD,
		ppstgOpen **IStorage)

	StgIsStorageFile func(name *OLECHAR)

	StgIsStorageILockBytes func(lkByt *ILockBytes)

	StgSetTimes func(
		name *OLECHAR, cTime, aTime, mTime *FILETIME)

	StgOpenAsyncDocfileOnIFillLockBytes func(
		pflb *IFillLockBytes,
		grfMode DWORD,
		asyncFlags DWORD,
		ppstgOpen **IStorage)

	StgGetIFillLockBytesOnILockBytes func(
		pilb *ILockBytes,
		ppflb **IFillLockBytes)

	StgGetIFillLockBytesOnFile func(
		pwcsName *OLECHAR,
		ppflb **IFillLockBytes)

	StgOpenLayoutDocfile func(
		pwcsDfName *OLECHAR,
		grfMode DWORD,
		reserved DWORD,
		ppstgOpen **IStorage)

	StgCreateStorageEx func(
		pwcsName *WChar,
		// TODO(t):WString

		grfMode DWORD,
		stgfmt DWORD,
		grfAttrs DWORD,
		StgOptions *STGOPTIONS,
		reserved *VOID,
		riid REFIID,
		ObjectOpen **VOID)

	StgOpenStorageEx func(
		pwcsName *WChar,
		// TODO(t):WString

		grfMode DWORD,
		stgfmt DWORD,
		grfAttrs DWORD,
		StgOptions *STGOPTIONS,
		reserved *VOID,
		riid REFIID,
		ObjectOpen **VOID)

	BindMoniker func(
		pmk *IMoniker,
		grfOpt DWORD,
		iidResult REFIID,
		ppvResult **VOID)

	CoInstall func(
		pbc *IBindCtx,
		Flags DWORD,
		ClassSpec *UCLSSPEC,
		Query *QUERYCONTEXT,
		CodeBase WString)

	CoGetObject func(
		Name WString,
		BindOptions *BIND_OPTS,
		riid REFIID,
		ppv **VOID)
	MkParseDisplayName func(
		pbc *IBindCtx,
		szUserName *OLESTR,
		pchEaten *ULONG,
		ppmk **IMoniker)
	MonikerRelativePathTo func(
		pmkSrc *IMoniker,
		pmkDest *IMoniker,
		ppmkRelPath **IMoniker,
		Reserved BOOL)
	MonikerCommonPrefixWith func(
		pmkThis *IMoniker,
		pmkOther *IMoniker,
		ppmkCommon **IMoniker)
	CreateBindCtx func(
		reserved DWORD,
		ppbc **IBindCtx)
	CreateGenericComposite func(
		pmkFirst *IMoniker,
		pmkRest *IMoniker,
		ppmkComposite **IMoniker)
	GetClassFile func(
		szFilename *OLESTR,
		pclsid *CLSID)

	CreateClassMoniker func(
		rclsid REFCLSID,
		ppmk **IMoniker)

	CreateFileMoniker func(
		lpszPathName *OLESTR,
		ppmk **IMoniker)

	CreateItemMoniker func(
		lpszDelim *OLESTR,
		lpszItem *OLESTR,
		ppmk **IMoniker)
	CreateAntiMoniker func(
		ppmk **IMoniker)
	CreatePointerMoniker func(
		Unk *IUnknown,
		ppmk **IMoniker)
	CreateObjrefMoniker func(
		Unk *IUnknown,
		ppmk **IMoniker)

	GetRunningObjectTable func(
		reserved DWORD,
		pprot **IRunningObjectTable)

	CreateStdProgressIndicator func(
		hwndParent HWND,
		Title *OLESTR,
		IbscCaller *IBindStatusCallback,
		Ibsc **IBindStatusCallback)
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
