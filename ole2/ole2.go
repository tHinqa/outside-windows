// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package ole2

import (
	. "github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/ole32"
)

var (
	CreateDataAdviseHolder func(daHolder **T.IDataAdviseHolder)

	OleBuildVersion func() T.DWORD

	ReadClassStg func(stg *T.IStorage, clsId *T.CLSID)

	WriteClassStg func(stg *T.IStorage, clsId T.REFCLSID)

	ReadClassStm func(stm *T.IStream, clsId *T.CLSID)

	WriteClassStm func(stm *T.IStream, clsId T.REFCLSID)

	WriteFmtUserTypeStg func(
		stg *T.IStorage,
		cf T.CLIPFORMAT,
		userType *T.OLESTR)

	ReadFmtUserTypeStg func(
		stg *T.IStorage,
		cf *T.CLIPFORMAT,
		userType **T.OLESTR)

	OleInitialize func(pvReserved *T.VOID)

	OleUninitialize func()

	OleQueryLinkFromData func(srcDataObject *T.IDataObject)

	OleQueryCreateFromData func(srcDataObject *T.IDataObject)

	OleCreate func(
		rClsId T.REFCLSID,
		riid T.REFIID,
		renderopt T.DWORD,
		formatEtc *T.FORMATETC,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateEx func(
		rClsId T.REFCLSID,
		riid T.REFIID,
		flags T.DWORD,
		renderopt T.DWORD,
		formats T.ULONG,
		advf *T.DWORD,
		formatEtc *T.FORMATETC,
		adviseSink *T.IAdviseSink,
		connection *T.DWORD,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateFromData func(
		srcDataObj *T.IDataObject,
		riid T.REFIID,
		renderopt T.DWORD,
		formatEtc *T.FORMATETC,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateFromDataEx func(
		srcDataObj *T.IDataObject,
		riid T.REFIID,
		flags T.DWORD,
		renderopt T.DWORD,
		formats T.ULONG,
		advf *T.DWORD,
		formatEtc *T.FORMATETC,
		adviseSink *T.IAdviseSink,
		connection *T.DWORD,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateLinkFromData func(
		srcDataObj *T.IDataObject,
		riid T.REFIID,
		renderopt T.DWORD,
		formatEtc *T.FORMATETC,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateLinkFromDataEx func(
		srcDataObj *T.IDataObject,
		riid T.REFIID,
		flags T.DWORD,
		renderopt T.DWORD,
		formats T.ULONG,
		advf *T.DWORD,
		formatEtc *T.FORMATETC,
		adviseSink *T.IAdviseSink,
		connection *T.DWORD,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateStaticFromData func(
		srcDataObj *T.IDataObject,
		riid T.REFIID,
		renderopt T.DWORD,
		formatEtc *T.FORMATETC,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateLink func(
		pmkLinkSrc *T.IMoniker,
		riid T.REFIID,
		renderopt T.DWORD,
		formatEtc *T.FORMATETC,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateLinkEx func(
		pmkLinkSrc *T.IMoniker,
		riid T.REFIID,
		flags T.DWORD,
		renderopt T.DWORD,
		formats T.ULONG,
		advf *T.DWORD,
		formatEtc *T.FORMATETC,
		adviseSink *T.IAdviseSink,
		connection *T.DWORD,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateLinkToFile func(
		fileName *T.OLESTR,
		riid T.REFIID,
		renderopt T.DWORD,
		formatEtc *T.FORMATETC,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateLinkToFileEx func(
		lpszFileName *T.OLESTR,
		riid T.REFIID,
		flags T.DWORD,
		renderopt T.DWORD,
		formats T.ULONG,
		advf *T.DWORD,
		formatEtc *T.FORMATETC,
		adviseSink *T.IAdviseSink,
		connection *T.DWORD,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateFromFile func(
		clsId T.REFCLSID,
		fileName *T.OLESTR,
		riid T.REFIID,
		renderopt T.DWORD,
		formatEtc *T.FORMATETC,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleCreateFromFileEx func(
		rClsId T.REFCLSID,
		fileName *T.OLESTR,
		riid T.REFIID,
		flags T.DWORD,
		renderopt T.DWORD,
		formats T.ULONG,
		advf *T.DWORD,
		formatEtc *T.FORMATETC,
		adviseSink *T.IAdviseSink,
		connection *T.DWORD,
		clientSite *T.IOleClientSite,
		stg *T.IStorage,
		obj **T.VOID)

	OleLoad func(
		stg *T.IStorage,
		riid T.REFIID,
		clientSite *T.IOleClientSite,
		obj **T.VOID)

	OleSave func(
		ps *T.IPersistStorage,
		stg *T.IStorage,
		sameAsLoad T.BOOL)

	OleLoadFromStream func(
		stm *T.STREAM,
		riidInterface T.REFIID,
		obj **T.VOID)

	OleSaveToStream func(
		ps *T.IPersistStream, stm *T.STREAM)

	OleSetContainedObject func(
		unknown *T.IUnknown, contained T.BOOL)

	OleNoteObjectVisible func(
		unknown *T.IUnknown, Visible T.BOOL)

	RegisterDragDrop func(
		wnd T.HWND, dropTarget *T.IDropTarget)

	RevokeDragDrop func(wnd T.HWND)

	DoDragDrop func(
		dataObj *T.IDataObject,
		dropSource *T.IDropSource,
		okEffects T.DWORD,
		effect *T.DWORD)

	OleSetClipboard func(dataObj *T.IDataObject)

	OleGetClipboard func(dataObj **T.IDataObject)

	OleFlushClipboard func()

	OleIsCurrentClipboard func(dataObj *T.IDataObject)

	OleCreateMenuDescriptor func(
		menuCombined T.HMENU,
		menuWidths *T.OLEMENUGROUPWIDTHS) T.HOLEMENU

	OleSetMenuDescriptor func(
		olemenu T.HOLEMENU,
		wndFrame T.HWND,
		wndActiveObject T.HWND,
		frame *T.IOleInPlaceFrame,
		activeObj *T.IOleInPlaceActiveObject)

	OleDestroyMenuDescriptor func(olemenu T.HOLEMENU)

	OleTranslateAccelerator func(
		frame *T.IOleInPlaceFrame,
		frameInfo *T.OLEINPLACEFRAMEINFO,
		msg *T.MSG)

	OleDuplicateData func(
		src T.HANDLE, format T.CLIPFORMAT, flags T.UINT) T.HANDLE

	OleDraw func(
		unknown *T.IUnknown,
		aspect T.DWORD,
		dcDraw T.HDC,
		rcBounds *T.RECT)

	OleRun func(unknown *T.IUnknown)

	OleIsRunning func(object *T.IOleObject) T.BOOL

	OleLockRunning func(
		unknown *T.IUnknown,
		lock T.BOOL,
		lastUnlockCloses T.BOOL)

	ReleaseStgMedium func(*T.STGMEDIUM)

	CreateOleAdviseHolder func(oaHolder **T.IOleAdviseHolder)

	OleCreateDefaultHandler func(
		clsId T.REFCLSID,
		unkOuter *T.IUnknown,
		riid T.REFIID,
		obj **T.VOID)

	OleCreateEmbeddingHelper func(
		clsId T.REFCLSID,
		unkOuter *T.IUnknown,
		flags T.DWORD,
		cf *T.IClassFactory,
		riid T.REFIID,
		obj **T.VOID)

	IsAccelerator func(
		accel T.HACCEL,
		accelEntries int,
		msg *T.MSG,
		cmd *T.WORD) T.BOOL

	OleGetIconOfFile func(
		path *T.OLESTR, useFileAsLabel T.BOOL) T.HGLOBAL

	OleGetIconOfClass func(
		rClsId T.REFCLSID,
		label *T.OLESTR,
		useTypeAsLabel T.BOOL) T.HGLOBAL

	OleMetafilePictFromIconAndLabel func(
		icon T.HICON,
		label *T.OLESTR,
		sourceFile *T.OLESTR,
		iconIndex T.UINT) T.HGLOBAL

	OleRegGetUserType func(
		clsId T.REFCLSID,
		formOfType T.DWORD,
		userType **T.OLESTR)

	OleRegGetMiscStatus func(
		clsId T.REFCLSID, aspect T.DWORD, status *T.DWORD)

	OleRegEnumFormatEtc func(
		clsId T.REFCLSID,
		direction T.DWORD,
		enum **T.IEnumFORMATETC)

	OleRegEnumVerbs func(clsId T.REFCLSID, enum **T.IEnumOLEVERB)

	OleConvertOLESTREAMToIStorage func(
		olestream *T.OLESTREAM,
		stg *T.IStorage,
		td *T.DVTARGETDEVICE)

	OleConvertIStorageToOLESTREAM func(
		stg *T.IStorage, Olestream *T.OLESTREAM)

	GetHGlobalFromILockBytes func(
		kByt *T.ILockBytes, global *T.HGLOBAL)

	CreateILockBytesOnHGlobal func(
		global T.HGLOBAL,
		deleteOnRelease T.BOOL,
		kByt **T.ILockBytes)

	GetHGlobalFromStream func(stm *T.STREAM, global *T.HGLOBAL)

	CreateStreamOnHGlobal func(
		global T.HGLOBAL,
		deleteOnRelease T.BOOL,
		stm **T.STREAM)

	OleDoAutoConvert func(stg *T.IStorage, clsidNew *T.CLSID)

	OleGetAutoConvert func(
		clsIdOld T.REFCLSID, clsidNew *T.CLSID)

	OleSetAutoConvert func(
		clsIdOld T.REFCLSID, clsIdNew T.REFCLSID)

	GetConvertStg func(stg *T.IStorage)

	SetConvertStg func(stg *T.IStorage, convert T.BOOL)

	OleConvertIStorageToOLESTREAMEx func(
		stg *T.IStorage,
		format T.CLIPFORMAT,
		width T.LONG,
		height T.LONG,
		size T.DWORD,
		medium *T.STGMEDIUM,
		olestm *T.OLESTREAM)

	OleConvertOLESTREAMToIStorageEx func(
		olestm *T.OLESTREAM,
		stg *T.IStorage,
		format *T.CLIPFORMAT,
		width *T.LONG,
		height *T.LONG,
		size *T.DWORD,
		medium *T.STGMEDIUM)
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
