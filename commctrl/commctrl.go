// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package commctrl provides API definitions for accessing
//comctl32.dll.
package commctrl

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside/types"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/comctl32"
)

var (
	InitMUILanguage func(uiLang T.LANGID)

	GetMUILanguage func() T.LANGID

	InitCommonControls func()

	InitCommonControlsEx func(*T.INITCOMMONCONTROLSEX) T.BOOL

	ImageList_Create func(
		x, y int, flags T.UINT, initial, grow int) T.HIMAGELIST

	ImageList_Destroy func(iml T.HIMAGELIST) T.BOOL

	ImageList_GetImageCount func(iml T.HIMAGELIST) int

	ImageList_SetImageCount func(
		iml T.HIMAGELIST, newCount T.UINT) T.BOOL

	ImageList_Add func(
		iml T.HIMAGELIST, image T.HBITMAP, mask T.HBITMAP) int

	ImageList_ReplaceIcon func(
		iml T.HIMAGELIST, i int, icon T.HICON) int

	ImageList_SetBkColor func(
		iml T.HIMAGELIST, clrBk T.COLORREF) T.COLORREF

	ImageList_GetBkColor func(iml T.HIMAGELIST) T.COLORREF

	ImageList_SetOverlayImage func(
		iml T.HIMAGELIST, T.Image, Overlay int) T.BOOL

	ImageList_Draw func(
		iml T.HIMAGELIST, i int, dst T.HDC, x, y int, style T.UINT) T.BOOL

	ImageList_Replace func(
		iml T.HIMAGELIST, i int, image T.HBITMAP, mask T.HBITMAP) T.BOOL

	ImageList_AddMasked func(
		iml T.HIMAGELIST, image T.HBITMAP, mask T.COLORREF) int

	ImageList_DrawEx func(
		iml T.HIMAGELIST,
		i int,
		dst T.HDC,
		x, y, dx, dy int,
		bk, fg T.COLORREF,
		style T.UINT) T.BOOL

	ImageList_DrawIndirect func(
		imldp *T.IMAGELISTDRAWPARAMS) T.BOOL

	ImageList_Remove func(iml T.HIMAGELIST, i int) T.BOOL

	ImageList_GetIcon func(
		himl T.HIMAGELIST, i int, flags T.UINT) T.HICON

	ImageList_LoadImage func(
		i T.HINSTANCE,
		bmp VString,
		cx, grow int,
		mask T.COLORREF,
		typ, flags T.UINT) T.HIMAGELIST

	ImageList_BeginDrag func(
		imlTrack T.HIMAGELIST, Track, xHotspot, yHotspot int) T.BOOL

	ImageList_EndDrag func()

	ImageList_DragEnter func(lock T.HWND, x, y int) T.BOOL

	ImageList_DragLeave func(lock T.HWND) T.BOOL

	ImageList_DragMove func(x, y int) T.BOOL

	ImageList_SetDragCursorImage func(
		imlDrag T.HIMAGELIST, drag, xHotspot, yHotspot int) T.BOOL

	ImageList_DragShowNolock func(show T.BOOL) T.BOOL

	ImageList_GetDragImage func(
		pt, hotspot *T.POINT) T.HIMAGELIST

	ImageList_Read func(stm *T.STREAM) T.HIMAGELIST

	ImageList_Write func(iml T.HIMAGELIST, stm *T.STREAM) T.BOOL

	ImageList_ReadEx func(
		flags T.DWORD, stm *T.STREAM, riid T.REFIID, v **T.VOID) T.HRESULT

	ImageList_WriteEx func(
		iml T.HIMAGELIST, flags T.DWORD, stm *T.STREAM) T.HRESULT

	ImageList_GetIconSize func(iml T.HIMAGELIST, x, y *int) T.BOOL

	ImageList_SetIconSize func(iml T.HIMAGELIST, x, y int) T.BOOL

	ImageList_GetImageInfo func(
		iml T.HIMAGELIST, i int, imageInfo *T.IMAGEINFO) T.BOOL

	ImageList_Merge func(iml1 T.HIMAGELIST,
		i1 int, iml2 T.HIMAGELIST, i2, x, y int) T.HIMAGELIST

	ImageList_Duplicate func(iml T.HIMAGELIST) T.HIMAGELIST

	ImageList_Copy func(dst T.HIMAGELIST, iDst int,
		src T.HIMAGELIST, iSrc int, flags T.UINT) T.BOOL

	CreateToolbarEx func(
		wnd T.HWND,
		s T.DWORD,
		id T.UINT,
		bitmaps int,
		bmInst T.HINSTANCE,
		bmID T.UINT_PTR,
		buttons *T.TBBUTTON,
		numButtons,
		xButton,
		yButton,
		xBitmap,
		yBitmap int,
		structSize T.UINT) T.HWND

	CreateMappedBitmap func(
		instance T.HINSTANCE,
		bitmap T.INT_PTR,
		flags T.UINT,
		colorMap *T.COLORMAP,
		NumMaps int) T.HBITMAP

	DrawStatusText func(
		dc T.HDC, rc *T.RECT, text VString, flags T.UINT)

	CreateStatusWindow func(
		style T.LONG, text VString, parent T.HWND, id T.UINT) T.HWND

	MenuHelp func(
		msg T.UINT,
		wParam T.WPARAM,
		lParam T.LPARAM,
		mainMenu T.HMENU,
		inst T.HINSTANCE,
		status T.HWND,
		ids *T.UINT)

	ShowHideMenuCtl func(wnd T.HWND, flags T.UINT_PTR, info *T.INT) T.BOOL

	GetEffectiveClientRect func(wnd T.HWND, rc *T.RECT, info *T.INT)

	MakeDragList func(lb T.HWND) T.BOOL

	DrawInsert func(parent, lb T.HWND, item int)

	LBItemFromPt func(lb T.HWND, pt T.POINT, autoScroll T.BOOL) int

	CreateUpDownControl func(
		style T.DWORD,
		x, y, cx, cy int,
		parent T.HWND,
		id int,
		inst T.HINSTANCE,
		buddy T.HWND,
		upper, lower, pos int) T.HWND

	DSA_Create func(item , itemGrow int) T.HDSA

	DSA_Destroy func(dsa T.HDSA) T.BOOL

	DSA_DestroyCallback func(
		dsa T.HDSA, cb *T.FNDSAENUMCALLBACK, data *T.VOID)

	DSA_GetItemPtr func(dsa T.HDSA, i int) *T.VOID

	DSA_InsertItem func(dsa T.HDSA, i int, pitem *T.VOID) int

	DPA_Create func(itemGrow int) T.HDPA

	DPA_Destroy func(dpa T.HDPA) T.BOOL

	DPA_DeletePtr func(dpa T.HDPA, i int) *T.VOID

	DPA_DeleteAllPtrs func(dpa T.HDPA) T.BOOL

	DPA_EnumCallback func(
		dpa T.HDPA, cb *T.FNDPAENUMCALLBACK, data *T.VOID)

	DPA_DestroyCallback func(
		dpa T.HDPA, cb *T.FNDPAENUMCALLBACK, data *T.VOID)

	DPA_SetPtr func(dpa T.HDPA, i int, p *T.VOID) T.BOOL

	DPA_InsertPtr func(dpa T.HDPA, i int, p *T.VOID) int

	DPA_GetPtr func(dpa T.HDPA, i T.INT_PTR) *T.VOID

	DPA_Sort func(
		dpa T.HDPA, compare *T.FNDPACOMPARE, lParam T.LPARAM) T.BOOL

	DPA_Search func(
		dpa T.HDPA,
		find *T.VOID,
		start int,
		compare *T.FNDPACOMPARE,
		lParam T.LPARAM,
		options T.UINT) int

	Str_SetPtrW func(
		ps *T.WString, //TODO(t):This gets free'd !!!
		s T.WString) T.BOOL

	//Was named _TrackMouseEvent and TrackMouseEvent
	//clashes with winuser.TrackMouseEvent
	TrackMouseEvent_ func(
		EventTrack *T.TRACKMOUSEEVENT) T.BOOL

	FlatSB_EnableScrollBar func(T.HWND, int, T.UINT) T.BOOL

	FlatSB_ShowScrollBar func(w T.HWND, code int, f T.BOOL) T.BOOL

	FlatSB_GetScrollRange func(w T.HWND, code int, s, e *T.INT) T.BOOL

	FlatSB_GetScrollInfo func(
		w T.HWND, code int, si *T.SCROLLINFO) T.BOOL

	FlatSB_GetScrollPos func(w T.HWND, code int) int

	FlatSB_GetScrollProp func(w T.HWND, propIndex int, p *T.INT) T.BOOL

	FlatSB_SetScrollPos func(
		w T.HWND, code, pos int, fRedraw T.BOOL) int

	FlatSB_SetScrollInfo func(
		w T.HWND, code int, si *T.SCROLLINFO, fRedraw T.BOOL) int

	FlatSB_SetScrollRange func(
		w T.HWND, code, min, max int, fRedraw T.BOOL) int

	FlatSB_SetScrollProp func(
		w T.HWND, index T.UINT, newValue T.INT_PTR, f T.BOOL) T.BOOL

	InitializeFlatSB func(T.HWND) T.BOOL

	UninitializeFlatSB func(T.HWND) T.HRESULT

	SetWindowSubclass func(
		Wnd T.HWND,
		Subclass T.SUBCLASSPROC,
		IdSubclass T.UINT_PTR,
		RefData T.DWORD_PTR) T.BOOL

	GetWindowSubclass func(
		Wnd T.HWND,
		Subclass T.SUBCLASSPROC,
		IdSubclass T.UINT_PTR,
		RefData *T.DWORD_PTR) T.BOOL

	RemoveWindowSubclass func(
		hWnd T.HWND,
		Subclass T.SUBCLASSPROC,
		IdSubclass T.UINT_PTR) T.BOOL

	DefSubclassProc func(
		Wnd T.HWND,
		Msg T.UINT,
		WParam T.WPARAM,
		LParam T.LPARAM) T.LRESULT

	DrawShadowText func(
		dc T.HDC,
		Text T.WString,
		c T.UINT,
		rc *T.RECT,
		Flags T.DWORD,
		clrText, clrShadow T.COLORREF,
		xOffset, yOffset int) int
)

/*
func ImageList_AddIcon(iml T.HIMAGELIST, icon T.HICON) int {
	return ImageList_ReplaceIcon(iml, -1, icon)
}

func ImageList_RemoveAll(iml T.HIMAGELIST) T.BOOL {
	return ImageList_Remove(iml, -1)
}

func ImageList_ExtractIcon(hi int, iml T.HIMAGELIST, i int) T.HICON {
	return ImageList_GetIcon(iml, i, 0)
}

func ImageList_LoadBitmap(i T.HINSTANCE, bmp VString, cx, T.Grow int, T.Mask T.COLORREF) {
	var IMAGE_BITMAP T.UINT = 0 //TODO(t):FIX
	ImageList_LoadImage(i, bmp, cx, Grow, Mask, IMAGE_BITMAP, 0)
}

 #ifdef _WIN64
    T.BOOL   FlatSB_GetScrollPropPtr func(T.HWND, int propIndex, T.PINT_PTR);
 #else
 #define FlatSB_GetScrollPropPtr  FlatSB_GetScrollProp
 #endif
 #define FlatSB_SetScrollPropPtr FlatSB_SetScrollProp
*/

//TODO(t):Not on XP
//{"ImageList_ReadEx", &ImageList_ReadEx},
//{"GetWindowSubclass", &GetWindowSubclass},

//TODO(t): Where
//{"DrawShadowText", &DrawShadowText},
//{"ImageList_WriteEx", &ImageList_WriteEx},

var CommCtrlANSIApis = outside.Apis{
	{"CreateStatusWindowA", &CreateStatusWindow},
	{"DrawStatusTextA", &DrawStatusText},
}

var CommCtrlUnicodeApis = outside.Apis{
	{"CreateStatusWindowW", &CreateStatusWindow},
	{"DrawStatusTextW", &DrawStatusText},
}

var CommCtrlApis = outside.Apis{
	{"CreateMappedBitmap", &CreateMappedBitmap},
	{"CreateToolbarEx", &CreateToolbarEx},
	{"CreateUpDownControl", &CreateUpDownControl},
	{"DefSubclassProc", &DefSubclassProc},
	{"DPA_Create", &DPA_Create},
	{"DPA_DeleteAllPtrs", &DPA_DeleteAllPtrs},
	{"DPA_DeletePtr", &DPA_DeletePtr},
	{"DPA_Destroy", &DPA_Destroy},
	{"DPA_DestroyCallback", &DPA_DestroyCallback},
	{"DPA_EnumCallback", &DPA_EnumCallback},
	{"DPA_GetPtr", &DPA_GetPtr},
	{"DPA_InsertPtr", &DPA_InsertPtr},
	{"DPA_Search", &DPA_Search},
	{"DPA_SetPtr", &DPA_SetPtr},
	{"DPA_Sort", &DPA_Sort},
	{"DrawInsert", &DrawInsert},
	{"DSA_Create", &DSA_Create},
	{"DSA_Destroy", &DSA_Destroy},
	{"DSA_DestroyCallback", &DSA_DestroyCallback},
	{"DSA_GetItemPtr", &DSA_GetItemPtr},
	{"DSA_InsertItem", &DSA_InsertItem},
	{"FlatSB_EnableScrollBar", &FlatSB_EnableScrollBar},
	{"FlatSB_GetScrollInfo", &FlatSB_GetScrollInfo},
	{"FlatSB_GetScrollPos", &FlatSB_GetScrollPos},
	{"FlatSB_GetScrollProp", &FlatSB_GetScrollProp},
	{"FlatSB_GetScrollRange", &FlatSB_GetScrollRange},
	{"FlatSB_SetScrollInfo", &FlatSB_SetScrollInfo},
	{"FlatSB_SetScrollPos", &FlatSB_SetScrollPos},
	{"FlatSB_SetScrollProp", &FlatSB_SetScrollProp},
	{"FlatSB_SetScrollRange", &FlatSB_SetScrollRange},
	{"FlatSB_ShowScrollBar", &FlatSB_ShowScrollBar},
	{"GetEffectiveClientRect", &GetEffectiveClientRect},
	{"GetMUILanguage", &GetMUILanguage},
	{"ImageList_Add", &ImageList_Add},
	{"ImageList_AddMasked", &ImageList_AddMasked},
	{"ImageList_BeginDrag", &ImageList_BeginDrag},
	{"ImageList_Copy", &ImageList_Copy},
	{"ImageList_Create", &ImageList_Create},
	{"ImageList_Destroy", &ImageList_Destroy},
	{"ImageList_DragEnter", &ImageList_DragEnter},
	{"ImageList_DragLeave", &ImageList_DragLeave},
	{"ImageList_DragMove", &ImageList_DragMove},
	{"ImageList_DragShowNolock", &ImageList_DragShowNolock},
	{"ImageList_Draw", &ImageList_Draw},
	{"ImageList_DrawEx", &ImageList_DrawEx},
	{"ImageList_DrawIndirect", &ImageList_DrawIndirect},
	{"ImageList_Duplicate", &ImageList_Duplicate},
	{"ImageList_EndDrag", &ImageList_EndDrag},
	{"ImageList_GetBkColor", &ImageList_GetBkColor},
	{"ImageList_GetDragImage", &ImageList_GetDragImage},
	{"ImageList_GetIcon", &ImageList_GetIcon},
	{"ImageList_GetIconSize", &ImageList_GetIconSize},
	{"ImageList_GetImageCount", &ImageList_GetImageCount},
	{"ImageList_GetImageInfo", &ImageList_GetImageInfo},
	{"ImageList_LoadImage", &ImageList_LoadImage},
	{"ImageList_Merge", &ImageList_Merge},
	{"ImageList_Read", &ImageList_Read},
	{"ImageList_Remove", &ImageList_Remove},
	{"ImageList_Replace", &ImageList_Replace},
	{"ImageList_ReplaceIcon", &ImageList_ReplaceIcon},
	{"ImageList_SetBkColor", &ImageList_SetBkColor},
	{"ImageList_SetDragCursorImage", &ImageList_SetDragCursorImage},
	{"ImageList_SetIconSize", &ImageList_SetIconSize},
	{"ImageList_SetImageCount", &ImageList_SetImageCount},
	{"ImageList_SetOverlayImage", &ImageList_SetOverlayImage},
	{"ImageList_Write", &ImageList_Write},
	{"InitCommonControls", &InitCommonControls},
	{"InitCommonControlsEx", &InitCommonControlsEx},
	{"InitializeFlatSB", &InitializeFlatSB},
	{"InitMUILanguage", &InitMUILanguage},
	{"LBItemFromPt", &LBItemFromPt},
	{"MakeDragList", &MakeDragList},
	{"MenuHelp", &MenuHelp},
	{"RemoveWindowSubclass", &RemoveWindowSubclass},
	{"SetWindowSubclass", &SetWindowSubclass},
	{"ShowHideMenuCtl", &ShowHideMenuCtl},
	{"Str_SetPtrW", &Str_SetPtrW},
	{"_TrackMouseEvent", &TrackMouseEvent_},
	{"UninitializeFlatSB", &UninitializeFlatSB},
}
