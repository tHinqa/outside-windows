package ole2

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/ole32"
)

var (
	CreateDataAdviseHolder func(DAHolder **IDataAdviseHolder)

	OleBuildVersion func() DWORD

	ReadClassStg func(Stg *IStorage, ClsId *CLSID)

	WriteClassStg func(Stg *IStorage, ClsId REFCLSID)

	ReadClassStm func(Stm *IStream, ClsId *CLSID)

	WriteClassStm func(Stm *IStream, ClsId REFCLSID)

	WriteFmtUserTypeStg func(
		Stg *IStorage,
		CF CLIPFORMAT,
		UserType *OLESTR)

	ReadFmtUserTypeStg func(
		Stg *IStorage,
		CF *CLIPFORMAT,
		UserType **OLESTR)

	OleInitialize func(pvReserved *VOID)

	OleUninitialize func()

	OleQueryLinkFromData func(SrcDataObject *IDataObject)

	OleQueryCreateFromData func(SrcDataObject *IDataObject)

	OleCreate func(
		RClsId REFCLSID,
		Riid REFIID,
		Renderopt DWORD,
		FormatEtc *FORMATETC,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateEx func(
		RClsId REFCLSID,
		Riid REFIID,
		Flags DWORD,
		Renderopt DWORD,
		Formats ULONG,
		Advf *DWORD,
		FormatEtc *FORMATETC,
		AdviseSink *IAdviseSink,
		Connection *DWORD,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateFromData func(
		SrcDataObj *IDataObject,
		Riid REFIID,
		Renderopt DWORD,
		FormatEtc *FORMATETC,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateFromDataEx func(
		SrcDataObj *IDataObject,
		Riid REFIID,
		Flags DWORD,
		Renderopt DWORD,
		Formats ULONG,
		Advf *DWORD,
		FormatEtc *FORMATETC,
		AdviseSink *IAdviseSink,
		Connection *DWORD,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateLinkFromData func(
		SrcDataObj *IDataObject,
		Riid REFIID,
		Renderopt DWORD,
		FormatEtc *FORMATETC,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateLinkFromDataEx func(
		SrcDataObj *IDataObject,
		Riid REFIID,
		Flags DWORD,
		Renderopt DWORD,
		Formats ULONG,
		Advf *DWORD,
		FormatEtc *FORMATETC,
		AdviseSink *IAdviseSink,
		Connection *DWORD,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateStaticFromData func(
		SrcDataObj *IDataObject,
		Riid REFIID,
		Renderopt DWORD,
		FormatEtc *FORMATETC,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateLink func(
		pmkLinkSrc *IMoniker,
		Riid REFIID,
		Renderopt DWORD,
		FormatEtc *FORMATETC,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateLinkEx func(
		pmkLinkSrc *IMoniker,
		Riid REFIID,
		Flags DWORD,
		Renderopt DWORD,
		Formats ULONG,
		Advf *DWORD,
		FormatEtc *FORMATETC,
		AdviseSink *IAdviseSink,
		Connection *DWORD,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateLinkToFile func(
		FileName *OLESTR,
		Riid REFIID,
		Renderopt DWORD,
		FormatEtc *FORMATETC,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateLinkToFileEx func(
		lpszFileName *OLESTR,
		Riid REFIID,
		Flags DWORD,
		Renderopt DWORD,
		Formats ULONG,
		Advf *DWORD,
		FormatEtc *FORMATETC,
		AdviseSink *IAdviseSink,
		Connection *DWORD,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateFromFile func(
		ClsId REFCLSID,
		FileName *OLESTR,
		Riid REFIID,
		Renderopt DWORD,
		FormatEtc *FORMATETC,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleCreateFromFileEx func(
		RClsId REFCLSID,
		FileName *OLESTR,
		Riid REFIID,
		Flags DWORD,
		Renderopt DWORD,
		Formats ULONG,
		Advf *DWORD,
		FormatEtc *FORMATETC,
		AdviseSink *IAdviseSink,
		Connection *DWORD,
		ClientSite *IOleClientSite,
		Stg *IStorage,
		Obj **VOID)

	OleLoad func(
		Stg *IStorage,
		Riid REFIID,
		ClientSite *IOleClientSite,
		Obj **VOID)

	OleSave func(
		PS *IPersistStorage,
		Stg *IStorage,
		SameAsLoad BOOL)

	OleLoadFromStream func(
		Stm *STREAM,
		RiidInterface REFIID,
		Obj **VOID)

	OleSaveToStream func(
		PStm *IPersistStream, Stm *STREAM)

	OleSetContainedObject func(
		Unknown *IUnknown, Contained BOOL)

	OleNoteObjectVisible func(
		Unknown *IUnknown, Visible BOOL)

	RegisterDragDrop func(
		Wnd HWND, DropTarget *IDropTarget)

	RevokeDragDrop func(Wnd HWND)

	DoDragDrop func(
		DataObj *IDataObject,
		DropSource *IDropSource,
		OKEffects DWORD,
		Effect *DWORD)

	OleSetClipboard func(DataObj *IDataObject)

	OleGetClipboard func(DataObj **IDataObject)

	OleFlushClipboard func()

	OleIsCurrentClipboard func(DataObj *IDataObject)

	OleCreateMenuDescriptor func(
		MenuCombined HMENU,
		MenuWidths *OLEMENUGROUPWIDTHS) HOLEMENU

	OleSetMenuDescriptor func(
		Olemenu HOLEMENU,
		WndFrame HWND,
		WndActiveObject HWND,
		Frame *IOleInPlaceFrame,
		ActiveObj *IOleInPlaceActiveObject)

	OleDestroyMenuDescriptor func(Olemenu HOLEMENU)

	OleTranslateAccelerator func(
		Frame *IOleInPlaceFrame,
		FrameInfo *OLEINPLACEFRAMEINFO,
		Msg *MSG)

	OleDuplicateData func(
		Src HANDLE, Format CLIPFORMAT, Flags UINT) HANDLE

	OleDraw func(
		Unknown *IUnknown,
		Aspect DWORD,
		DCDraw HDC,
		RCBounds *RECT)

	OleRun func(Unknown *IUnknown)

	OleIsRunning func(Object *IOleObject) BOOL

	OleLockRunning func(
		Unknown *IUnknown,
		Lock BOOL,
		LastUnlockCloses BOOL)

	ReleaseStgMedium func(*STGMEDIUM)

	CreateOleAdviseHolder func(OAHolder **IOleAdviseHolder)

	OleCreateDefaultHandler func(
		ClsId REFCLSID,
		UnkOuter *IUnknown,
		Riid REFIID,
		Obj **VOID)

	OleCreateEmbeddingHelper func(
		ClsId REFCLSID,
		UnkOuter *IUnknown,
		Flags DWORD,
		CF *IClassFactory,
		Riid REFIID,
		Obj **VOID)

	IsAccelerator func(
		Accel HACCEL,
		AccelEntries int,
		Msg *MSG,
		Cmd *WORD) BOOL

	OleGetIconOfFile func(
		Path *OLESTR, UseFileAsLabel BOOL) HGLOBAL

	OleGetIconOfClass func(
		RClsId REFCLSID,
		Label *OLESTR,
		UseTypeAsLabel BOOL) HGLOBAL

	OleMetafilePictFromIconAndLabel func(
		Icon HICON,
		Label *OLESTR,
		SourceFile *OLESTR,
		IconIndex UINT) HGLOBAL

	OleRegGetUserType func(
		ClsId REFCLSID,
		FormOfType DWORD,
		UserType **OLESTR)

	OleRegGetMiscStatus func(
		ClsId REFCLSID, Aspect DWORD, Status *DWORD)

	OleRegEnumFormatEtc func(
		ClsId REFCLSID,
		Direction DWORD,
		Enum **IEnumFORMATETC)

	OleRegEnumVerbs func(ClsId REFCLSID, Enum **IEnumOLEVERB)

	OleConvertOLESTREAMToIStorage func(
		Olestream *OLESTREAM,
		Stg *IStorage,
		TD *DVTARGETDEVICE)

	OleConvertIStorageToOLESTREAM func(
		Stg *IStorage, Olestream *OLESTREAM)

	GetHGlobalFromILockBytes func(
		KByt *ILockBytes, global *HGLOBAL)

	CreateILockBytesOnHGlobal func(
		Global HGLOBAL,
		DeleteOnRelease BOOL,
		KByt **ILockBytes)

	GetHGlobalFromStream func(Stm *STREAM, Global *HGLOBAL)

	CreateStreamOnHGlobal func(
		Global HGLOBAL,
		DeleteOnRelease BOOL,
		Stm **STREAM)

	OleDoAutoConvert func(Stg *IStorage, ClsidNew *CLSID)

	OleGetAutoConvert func(
		ClsIdOld REFCLSID, ClsidNew *CLSID)

	OleSetAutoConvert func(
		ClsIdOld REFCLSID, ClsIdNew REFCLSID)

	GetConvertStg func(Stg *IStorage)

	SetConvertStg func(Stg *IStorage, Convert BOOL)

	OleConvertIStorageToOLESTREAMEx func(
		Stg *IStorage,
		Format CLIPFORMAT,
		Width LONG,
		Height LONG,
		Size DWORD,
		Medium *STGMEDIUM,
		Olestm *OLESTREAM)

	OleConvertOLESTREAMToIStorageEx func(
		Olestm *OLESTREAM,
		Stg *IStorage,
		Format *CLIPFORMAT,
		Width *LONG,
		Height *LONG,
		Size *DWORD,
		Medium *STGMEDIUM)
)

var Ole2Apis = Apis{
	{"CreateDataAdviseHolder", &CreateDataAdviseHolder},
	{"CreateILockBytesOnHGlobal", &CreateILockBytesOnHGlobal},
	{"CreateOleAdviseHolder", &CreateOleAdviseHolder},
	{"CreateStreamOnHGlobal", &CreateStreamOnHGlobal},
	{"DoDragDrop", &DoDragDrop},
	{"GetConvertStg", &GetConvertStg},
	{"GetHGlobalFromILockBytes", &GetHGlobalFromILockBytes},
	{"GetHGlobalFromStream", &GetHGlobalFromStream},
	{"IsAccelerator", &IsAccelerator},
	{"OleBuildVersion", &OleBuildVersion},
	{"OleConvertIStorageToOLESTREAM", &OleConvertIStorageToOLESTREAM},
	{"OleConvertIStorageToOLESTREAMEx", &OleConvertIStorageToOLESTREAMEx},
	{"OleConvertOLESTREAMToIStorage", &OleConvertOLESTREAMToIStorage},
	{"OleConvertOLESTREAMToIStorageEx", &OleConvertOLESTREAMToIStorageEx},
	{"OleCreate", &OleCreate},
	{"OleCreateDefaultHandler", &OleCreateDefaultHandler},
	{"OleCreateEmbeddingHelper", &OleCreateEmbeddingHelper},
	{"OleCreateEx", &OleCreateEx},
	{"OleCreateFromData", &OleCreateFromData},
	{"OleCreateFromDataEx", &OleCreateFromDataEx},
	{"OleCreateFromFile", &OleCreateFromFile},
	{"OleCreateFromFileEx", &OleCreateFromFileEx},
	{"OleCreateLink", &OleCreateLink},
	{"OleCreateLinkEx", &OleCreateLinkEx},
	{"OleCreateLinkFromData", &OleCreateLinkFromData},
	{"OleCreateLinkFromDataEx", &OleCreateLinkFromDataEx},
	{"OleCreateLinkToFile", &OleCreateLinkToFile},
	{"OleCreateLinkToFileEx", &OleCreateLinkToFileEx},
	{"OleCreateMenuDescriptor", &OleCreateMenuDescriptor},
	{"OleCreateStaticFromData", &OleCreateStaticFromData},
	{"OleDestroyMenuDescriptor", &OleDestroyMenuDescriptor},
	{"OleDoAutoConvert", &OleDoAutoConvert},
	{"OleDraw", &OleDraw},
	{"OleDuplicateData", &OleDuplicateData},
	{"OleFlushClipboard", &OleFlushClipboard},
	{"OleGetAutoConvert", &OleGetAutoConvert},
	{"OleGetClipboard", &OleGetClipboard},
	{"OleGetIconOfClass", &OleGetIconOfClass},
	{"OleGetIconOfFile", &OleGetIconOfFile},
	{"OleInitialize", &OleInitialize},
	{"OleIsCurrentClipboard", &OleIsCurrentClipboard},
	{"OleIsRunning", &OleIsRunning},
	{"OleLoad", &OleLoad},
	{"OleLoadFromStream", &OleLoadFromStream},
	{"OleLockRunning", &OleLockRunning},
	{"OleMetafilePictFromIconAndLabel", &OleMetafilePictFromIconAndLabel},
	{"OleNoteObjectVisible", &OleNoteObjectVisible},
	{"OleQueryCreateFromData", &OleQueryCreateFromData},
	{"OleQueryLinkFromData", &OleQueryLinkFromData},
	{"OleRegEnumFormatEtc", &OleRegEnumFormatEtc},
	{"OleRegEnumVerbs", &OleRegEnumVerbs},
	{"OleRegGetMiscStatus", &OleRegGetMiscStatus},
	{"OleRegGetUserType", &OleRegGetUserType},
	{"OleRun", &OleRun},
	{"OleSave", &OleSave},
	{"OleSaveToStream", &OleSaveToStream},
	{"OleSetAutoConvert", &OleSetAutoConvert},
	{"OleSetClipboard", &OleSetClipboard},
	{"OleSetContainedObject", &OleSetContainedObject},
	{"OleSetMenuDescriptor", &OleSetMenuDescriptor},
	{"OleTranslateAccelerator", &OleTranslateAccelerator},
	{"OleUninitialize", &OleUninitialize},
	{"ReadClassStg", &ReadClassStg},
	{"ReadClassStm", &ReadClassStm},
	{"ReadFmtUserTypeStg", &ReadFmtUserTypeStg},
	{"RegisterDragDrop", &RegisterDragDrop},
	{"ReleaseStgMedium", &ReleaseStgMedium},
	{"RevokeDragDrop", &RevokeDragDrop},
	{"SetConvertStg", &SetConvertStg},
	{"WriteClassStg", &WriteClassStg},
	{"WriteClassStm", &WriteClassStm},
	{"WriteFmtUserTypeStg", &WriteFmtUserTypeStg},
}
