package commctrl

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/comctl32"
)

var (
	InitMUILanguage func(uiLang LANGID)

	GetMUILanguage func() LANGID

	InitCommonControls func()

	InitCommonControlsEx func(*INITCOMMONCONTROLSEX) BOOL

	ImageList_Create func(
		x, y int, flags UINT, Initial, Grow int) HIMAGELIST

	ImageList_Destroy func(iml HIMAGELIST) BOOL

	ImageList_GetImageCount func(iml HIMAGELIST) int

	ImageList_SetImageCount func(
		iml HIMAGELIST, NewCount UINT) BOOL

	ImageList_Add func(
		iml HIMAGELIST, Image HBITMAP, Mask HBITMAP) int

	ImageList_ReplaceIcon func(
		iml HIMAGELIST, i int, icon HICON) int

	ImageList_SetBkColor func(
		iml HIMAGELIST, clrBk COLORREF) COLORREF

	ImageList_GetBkColor func(iml HIMAGELIST) COLORREF

	ImageList_SetOverlayImage func(
		iml HIMAGELIST, Image, Overlay int) BOOL

	ImageList_Draw func(
		iml HIMAGELIST, i int, Dst HDC, x, y int, Style UINT) BOOL

	ImageList_Replace func(
		iml HIMAGELIST, i int, Image HBITMAP, Mask HBITMAP) BOOL

	ImageList_AddMasked func(
		iml HIMAGELIST, Image HBITMAP, Mask COLORREF) int

	ImageList_DrawEx func(
		iml HIMAGELIST,
		i int,
		Dst HDC,
		x, y, dx, dy int,
		Bk, Fg COLORREF,
		Style UINT) BOOL

	ImageList_DrawIndirect func(
		imldp *IMAGELISTDRAWPARAMS) BOOL

	ImageList_Remove func(iml HIMAGELIST, i int) BOOL

	ImageList_GetIcon func(
		himl HIMAGELIST, i int, flags UINT) HICON

	ImageList_LoadImage func(
		i HINSTANCE,
		bmp VString,
		cx, Grow int,
		Mask COLORREF,
		Type, Flags UINT) HIMAGELIST

	ImageList_BeginDrag func(
		imlTrack HIMAGELIST, Track, xHotspot, yHotspot int) BOOL

	ImageList_EndDrag func()

	ImageList_DragEnter func(Lock HWND, x, y int) BOOL

	ImageList_DragLeave func(Lock HWND) BOOL

	ImageList_DragMove func(x, y int) BOOL

	ImageList_SetDragCursorImage func(
		imlDrag HIMAGELIST, Drag, xHotspot, yHotspot int) BOOL

	ImageList_DragShowNolock func(Show BOOL) BOOL

	ImageList_GetDragImage func(
		pt, Hotspot *POINT) HIMAGELIST

	ImageList_Read func(stm *STREAM) HIMAGELIST

	ImageList_Write func(iml HIMAGELIST, stm *STREAM) BOOL

	ImageList_ReadEx func(
		Flags DWORD, stm *STREAM, riid REFIID, v **VOID) HRESULT

	ImageList_WriteEx func(
		iml HIMAGELIST, Flags DWORD, stm *STREAM) HRESULT

	ImageList_GetIconSize func(iml HIMAGELIST, x, y *int) BOOL

	ImageList_SetIconSize func(iml HIMAGELIST, x, y int) BOOL

	ImageList_GetImageInfo func(
		iml HIMAGELIST, i int, ImageInfo *IMAGEINFO) BOOL

	ImageList_Merge func(iml1 HIMAGELIST,
		i1 int, iml2 HIMAGELIST, i2, x, y int) HIMAGELIST

	ImageList_Duplicate func(iml HIMAGELIST) HIMAGELIST

	ImageList_Copy func(Dst HIMAGELIST, iDst int,
		Src HIMAGELIST, iSrc int, Flags UINT) BOOL

	CreateToolbarEx func(
		wnd HWND,
		s DWORD,
		ID UINT,
		Bitmaps int,
		BMInst HINSTANCE,
		BMID UINT_PTR,
		Buttons *TBBUTTON,
		NumButtons,
		xButton,
		yButton,
		xBitmap,
		yBitmap int,
		StructSize UINT) HWND

	CreateMappedBitmap func(
		Instance HINSTANCE,
		Bitmap INT_PTR,
		Flags UINT,
		ColorMap *COLORMAP,
		NumMaps int) HBITMAP

	DrawStatusText func(
		DC HDC, rc *RECT, Text VString, Flags UINT)

	CreateStatusWindow func(
		style LONG, Text VString, Parent HWND, ID UINT) HWND

	MenuHelp func(
		Msg UINT,
		WParam WPARAM,
		LParam LPARAM,
		MainMenu HMENU,
		Inst HINSTANCE,
		Status HWND,
		IDs *UINT)

	ShowHideMenuCtl func(Wnd HWND, Flags UINT_PTR, Info *INT) BOOL

	GetEffectiveClientRect func(Wnd HWND, rc *RECT, Info *INT)

	MakeDragList func(LB HWND) BOOL

	DrawInsert func(Parent, LB HWND, Item int)

	LBItemFromPt func(LB HWND, pt POINT, AutoScroll BOOL) int

	CreateUpDownControl func(
		Style DWORD,
		x, y, cx, cy int,
		Parent HWND,
		ID int,
		Inst HINSTANCE,
		Buddy HWND,
		Upper, Lower, Pos int) HWND

	DSA_Create func(Item int, ItemGrow int) HDSA

	DSA_Destroy func(dsa HDSA) BOOL

	DSA_DestroyCallback func(
		dsa HDSA, CB *FNDSAENUMCALLBACK, Data *VOID)

	DSA_GetItemPtr func(dsa HDSA, i int) *VOID

	DSA_InsertItem func(dsa HDSA, i int, pitem *VOID) int

	DPA_Create func(ItemGrow int) HDPA

	DPA_Destroy func(dpa HDPA) BOOL

	DPA_DeletePtr func(dpa HDPA, i int) *VOID

	DPA_DeleteAllPtrs func(dpa HDPA) BOOL

	DPA_EnumCallback func(
		dpa HDPA, CB *FNDPAENUMCALLBACK, Data *VOID)

	DPA_DestroyCallback func(
		dpa HDPA, CB *FNDPAENUMCALLBACK, Data *VOID)

	DPA_SetPtr func(dpa HDPA, i int, p *VOID) BOOL

	DPA_InsertPtr func(dpa HDPA, i int, p *VOID) int

	DPA_GetPtr func(dpa HDPA, i INT_PTR) *VOID

	DPA_Sort func(
		dpa HDPA, Compare *FNDPACOMPARE, LParam LPARAM) BOOL

	DPA_Search func(
		dpa HDPA,
		Find *VOID,
		Start int,
		Compare *FNDPACOMPARE,
		LParam LPARAM,
		options UINT) int

	Str_SetPtrW func(
		ps *WString, //TODO(t):This gets free'd !!!
		s WString) BOOL

	//Was named _TrackMouseEvent and TrackMouseEvent
	//clashes with winuser.TrackMouseEvent
	TrackMouseEvent_ func(
		EventTrack *TRACKMOUSEEVENT) BOOL

	FlatSB_EnableScrollBar func(HWND, int, UINT) BOOL

	FlatSB_ShowScrollBar func(w HWND, code int, f BOOL) BOOL

	FlatSB_GetScrollRange func(w HWND, code int, s, e *INT) BOOL

	FlatSB_GetScrollInfo func(
		w HWND, code int, si *SCROLLINFO) BOOL

	FlatSB_GetScrollPos func(w HWND, code int) int

	FlatSB_GetScrollProp func(w HWND, propIndex int, p *INT) BOOL

	FlatSB_SetScrollPos func(
		w HWND, code, pos int, fRedraw BOOL) int

	FlatSB_SetScrollInfo func(
		w HWND, code int, si *SCROLLINFO, fRedraw BOOL) int

	FlatSB_SetScrollRange func(
		w HWND, code, min, max int, fRedraw BOOL) int

	FlatSB_SetScrollProp func(
		w HWND, index UINT, newValue INT_PTR, f BOOL) BOOL

	InitializeFlatSB func(HWND) BOOL

	UninitializeFlatSB func(HWND) HRESULT

	SetWindowSubclass func(
		Wnd HWND,
		Subclass SUBCLASSPROC,
		IdSubclass UINT_PTR,
		RefData DWORD_PTR) BOOL

	GetWindowSubclass func(
		Wnd HWND,
		Subclass SUBCLASSPROC,
		IdSubclass UINT_PTR,
		RefData *DWORD_PTR) BOOL

	RemoveWindowSubclass func(
		hWnd HWND,
		Subclass SUBCLASSPROC,
		IdSubclass UINT_PTR) BOOL

	DefSubclassProc func(
		Wnd HWND,
		Msg UINT,
		WParam WPARAM,
		LParam LPARAM) LRESULT

	DrawShadowText func(
		dc HDC,
		Text WString,
		c UINT,
		rc *RECT,
		Flags DWORD,
		clrText, clrShadow COLORREF,
		xOffset, yOffset int) int
)

/*
func ImageList_AddIcon(iml HIMAGELIST, icon HICON) int {
	return ImageList_ReplaceIcon(iml, -1, icon)
}

func ImageList_RemoveAll(iml HIMAGELIST) BOOL {
	return ImageList_Remove(iml, -1)
}

func ImageList_ExtractIcon(hi int, iml HIMAGELIST, i int) HICON {
	return ImageList_GetIcon(iml, i, 0)
}

func ImageList_LoadBitmap(i HINSTANCE, bmp VString, cx, Grow int, Mask COLORREF) {
	var IMAGE_BITMAP UINT = 0 //TODO(t):FIX
	ImageList_LoadImage(i, bmp, cx, Grow, Mask, IMAGE_BITMAP, 0)
}

 #ifdef _WIN64
    BOOL   FlatSB_GetScrollPropPtr func(HWND, int propIndex, PINT_PTR);
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

var CommCtrlANSIApis = Apis{
	{"CreateStatusWindowA", &CreateStatusWindow},
	{"DrawStatusTextA", &DrawStatusText},
}

var CommCtrlUnicodeApis = Apis{
	{"CreateStatusWindowW", &CreateStatusWindow},
	{"DrawStatusTextW", &DrawStatusText},
}

var CommCtrlApis = Apis{
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
