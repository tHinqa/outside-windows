// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package winuser provides API definitions for accessing
//user32.dll.
package winuser

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside/types"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/user32"
)

func CopyCursor(cur T.HCURSOR) (T.HCURSOR, error) {
	i, e := CopyIcon(T.HICON(cur))
	return T.HCURSOR(i),e
}

func CreateDialog(instance T.HINSTANCE, templateName VString, wndParent T.HWND, dialogFunc T.DLGPROC) (T.HWND, error) {
	return CreateDialogParam(instance, templateName, wndParent, dialogFunc, 0)
}

func CreateDialogIndirect(instance T.HINSTANCE, template T.DLGTEMPLATE, wndParent T.HWND,dialogFunc T.DLGPROC) (T.HWND, error) {
	return CreateDialogIndirectParam(instance, template, wndParent, dialogFunc, 0)
}

func CreateWindow(className, windowName VString, style T.WINDOW_STYLE, x, y, width, height int, wndParent T.HWND, menu T.HMENU, instance T.HINSTANCE, param *T.VOID) (T.HWND, error) {
	return CreateWindowEx(0, className, windowName, style, x, y, width, height, wndParent, menu, instance, param)
}

func DialogBox(instance T.HINSTANCE, template VString, wndParent T.HWND, dialogFunc T.DLGPROC) (T.INT_PTR, error) {
	return DialogBoxParam(instance, template, wndParent, dialogFunc, 0)
}

func DialogBoxIndirect(instance T.HINSTANCE, template T.DLGTEMPLATE, wndParent T.HWND, dialogFunc T.DLGPROC) (T.INT_PTR, error) {
	return DialogBoxIndirectParam(instance, template, wndParent, dialogFunc, 0)
}

func EnumTaskWindows(task T.DWORD, fn T.WNDENUMPROC, param T.LPARAM) (T.BOOL, error) {
	return EnumThreadWindows(task, fn, param)
}

func GetNextWindow(Wnd T.HWND, cmd T.UINT) (T.HWND, error) {
	return GetWindow(Wnd, cmd)
}

func GetSysModalWindow() {}

func GetWindowTask(Wnd T.HWND) (T.HANDLE, error) {
	i, e := GetWindowThreadProcessId(Wnd, nil)
	return T.HANDLE(i), e
}

func SetSysModalWindow(Wnd T.HWND) {}

type (
	va_list  T.Fake_type_Fix_me
)

var (
	ActivateKeyboardLayout func(hkl T.HKL, flags T.UINT) (T.HKL, error)

	AdjustWindowRect func(
		rect *T.RECT, style T.DWORD, menu T.BOOL) (T.BOOL, error)

	AdjustWindowRectEx func(
		rect *T.RECT, style T.DWORD, menu T.BOOL, exStyle T.DWORD) (T.BOOL, error)

	AllowSetForegroundWindow func(processId T.DWORD) (T.BOOL, error)

	AnimateWindow func(w T.HWND, time T.DWORD, flags T.DWORD) (T.BOOL, error)

	AnyPopup func() (T.BOOL, error)

	AppendMenu func(
		m T.HMENU,
		flags T.UINT,
		idNewItem T.UINT_PTR,
		newItem VString) (T.BOOL, error)

	ArrangeIconicWindows func(w T.HWND) (T.UINT, error)

	AttachThreadInput func(attach, to T.DWORD, flag T.BOOL) (T.BOOL, error)

	BeginDeferWindowPos func(numWindows int) (T.HDWP, error)

	BeginPaint func(w T.HWND, paint *T.PAINTSTRUCT) (T.HDC, error)

	BringWindowToTop func(w T.HWND) (T.BOOL, error)

	BroadcastSystemMessage func(
		flags T.DWORD,
		info *T.DWORD,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM) (T.LONG, error)

	BroadcastSystemMessageEx func(
		flags T.BSM_FLAGS,
		recipients *T.DWORD,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM,
		info *T.BSMINFO) (T.LONG, error)

	CallMsgFilter func(msg *T.MSG, code int) (T.BOOL, error)

	CallNextHookEx func(
		hk T.HHOOK, code int, w T.WPARAM, l T.LPARAM) (T.LRESULT, error)

	CallWindowProc func(
		prevWndFunc T.WNDPROC,
		win T.HWND,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM) (T.LRESULT, error)

	CascadeWindows func(
		wndParent T.HWND,
		how T.UINT,
		rect *T.RECT,
		nKids T.UINT,
		kids *T.HWND) (T.WORD, error)

	ChangeClipboardChain func(wndRemove, wndNewNext T.HWND) (T.BOOL, error)

	ChangeDisplaySettingsA func(dm *T.DEVMODEA, flags T.DWORD) (T.LONG, error)

	ChangeDisplaySettingsW func(dm *T.DEVMODEW, flags T.DWORD) (T.LONG, error)

	ChangeDisplaySettingsExA func(
		deviceName VString,
		dm *T.DEVMODEA,
		w T.HWND,
		flags T.DWORD,
		lParam *T.VOID) (T.LONG, error)

	ChangeDisplaySettingsExW func(
		deviceName VString,
		dm *T.DEVMODEW,
		w T.HWND,
		flags T.DWORD,
		lParam *T.VOID) (T.LONG, error)

	ChangeMenu func(
		m T.HMENU,
		cmd T.UINT,
		NewItem VString,
		insert T.UINT,
		flags T.UINT) (T.BOOL, error)

	CharLower func(VString) (OVString, error)

	CharLowerBuff func(buff T.IOAString, length T.DWORD) (T.DWORD, error)

	CharNext func(VString) (OVString, error)

	CharNextExA func(
		codePage T.WORD, currentChar T.AString, flags T.DWORD) (T.AString, error)

	CharPrev func(start, current VString) (VString, error)

	CharPrevExA func(
		codePage T.WORD,
		start, currentChar T.AString,
		flags T.DWORD) (T.AString, error)

	CharToOem func(src VString, dst OVString) (T.BOOL, error)

	CharToOemBuff func(src, dst OVString, dstLength T.DWORD) (T.BOOL, error)

	CharUpper func(VString) (OVString, error)

	CharUpperBuff func(buff T.IOAString, length T.DWORD) (T.DWORD, error)

	CheckDlgButton func(dlg T.HWND, idButton int, check T.UINT) (T.BOOL, error)

	CheckMenuItem func(m T.HMENU, idCheckItem, check T.UINT) (T.DWORD, error)

	CheckMenuRadioItem func(
		m T.HMENU, first, last, check, flags T.UINT) (T.BOOL, error)

	CheckRadioButton func(
		dlg T.HWND, idFirst, idLast, idCheck int) (T.BOOL, error)

	ChildWindowFromPoint func(wndParent T.HWND, point T.POINT) (T.HWND, error)

	ChildWindowFromPointEx func(
		w T.HWND, p T.POINT, flags T.UINT) (T.HWND, error)

	ClientToScreen func(w T.HWND, p *T.POINT) (T.BOOL, error)

	ClipCursor func(rect *T.RECT) (T.BOOL, error)

	CloseClipboard func() (T.BOOL, error)

	CloseDesktop func(desktop T.HDESK) (T.BOOL, error)

	CloseWindow func(w T.HWND) (T.BOOL, error)

	CloseWindowStation func(winSta T.HWINSTA) (T.BOOL, error)

	CopyAcceleratorTable func(
		src T.HACCEL, dst *T.ACCEL, entries int) (int, error)

	CopyIcon func(icon T.HICON) (T.HICON, error)

	CopyImage func(
		h T.HANDLE, typ T.UINT, T.X, y int, flags T.UINT) (T.HANDLE, error)

	CopyRect func(dst, src *T.RECT) (T.BOOL, error)

	CountClipboardFormats func() (int, error)

	CreateAcceleratorTable func(accel *T.ACCEL, nAccel int) (T.HACCEL, error)

	CreateCaret func(
		w T.HWND, bitmap T.HBITMAP, width int, height int) (T.BOOL, error)

	CreateCursor func(
		inst T.HINSTANCE,
		xHotSpot, yHotSpot, width, height int,
		andPlane, xorPlane *T.VOID) (T.HCURSOR, error)

	CreateDesktopA func(
		desktop, device VString,
		devmode *T.DEVMODEA,
		flags T.DWORD,
		desiredAccess T.ACCESS_MASK,
		sa *T.SECURITY_ATTRIBUTES) (T.HDESK, error)

	CreateDesktopW func(
		desktop, device VString,
		devmode *T.DEVMODEW,
		flags T.DWORD,
		desiredAccess T.ACCESS_MASK,
		sa *T.SECURITY_ATTRIBUTES) (T.HDESK, error)

	CreateDialogIndirectParam func(
		instance T.HINSTANCE,
		template T.DLGTEMPLATE,
		wndParent T.HWND,
		dialogFunc T.DLGPROC,
		initParam T.LPARAM) (T.HWND, error)

	CreateDialogParam func(
		instance T.HINSTANCE,
		templateName VString,
		wndParent T.HWND,
		dialogFunc T.DLGPROC,
		initParam T.LPARAM) (T.HWND, error)

	CreateIcon func(
		instance T.HINSTANCE,
		width, height int,
		planes, bitsPixel T.BYTE,
		andbits, xorbits *T.BYTE) (T.HICON, error)

	CreateIconFromResource func(
		resbits *T.BYTE,
		resSize T.DWORD,
		icon T.BOOL,
		ver T.DWORD) (T.HICON, error)

	CreateIconFromResourceEx func(
		resbits *T.BYTE,
		resSize T.DWORD,
		icon T.BOOL,
		ver T.DWORD,
		xDesired, yDesired int,
		flags T.UINT) (T.HICON, error)

	CreateIconIndirect func(iconinfo *T.ICONINFO) (T.HICON, error)

	CreateMDIWindow func(
		className, windowName VString,
		style T.DWORD,
		x, y, width,
		height int,
		wndParent T.HWND,
		instance T.HINSTANCE,
		l T.LPARAM) (T.HWND, error)

	CreateMenu func() (T.HMENU, error)

	CreatePopupMenu func() (T.HMENU, error)

	CreateWindowEx func(
		exStyle T.WINDOW_STYLE_EX,
		className, windowName VString,
		style T.WINDOW_STYLE,
		x, y, width, height int,
		wndParent T.HWND,
		m T.HMENU,
		instance T.HINSTANCE,
		param *T.VOID) (T.HWND, error)

	CreateWindowStation func(
		winsta VString,
		flags T.DWORD,
		desiredAccess T.ACCESS_MASK,
		sa *T.SECURITY_ATTRIBUTES) (T.HWINSTA, error)

	DefDlgProc func(
		dlg T.HWND, msg T.UINT, w T.WPARAM, l T.LPARAM) (T.LRESULT, error)

	DeferWindowPos func(
		winPosInfo T.HDWP,
		w T.HWND,
		wndInsertAfter T.HWND,
		x, y, cx, cy int,
		flags T.UINT) (T.HDWP, error)

	DefFrameProc func(
		win T.HWND,
		mdiClient T.HWND,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM) (T.LRESULT, error)

	DefMDIChildProc func(
		win T.HWND, msg T.UINT, w T.WPARAM, l T.LPARAM) (T.LRESULT, error)

	DefRawInputProc func(rawInput *T.RAWINPUT, input T.INT, sizeHeader T.INT) (T.LRESULT, error)

	DefWindowProc func(h T.HWND, m T.WINDOW_MESSAGE, w T.WPARAM, l T.LPARAM) (T.LRESULT, error)

	DeleteMenu func(m T.HMENU, position T.UINT, flags T.UINT) (T.BOOL, error)

	DeregisterShellHookWindow func(w T.HWND) (T.BOOL, error)

	DestroyAcceleratorTable func(accel T.HACCEL) (T.BOOL, error)

	DestroyCaret func() (T.BOOL, error)

	DestroyCursor func(cursor T.HCURSOR) (T.BOOL, error)

	DestroyIcon func(icon T.HICON) (T.BOOL, error)

	DestroyMenu func(m T.HMENU) (T.BOOL, error)

	DestroyWindow func(w T.HWND) (T.BOOL, error)

	DialogBoxIndirectParam func(
		instance T.HINSTANCE,
		dialogTemplate T.DLGTEMPLATE,
		wndParent T.HWND,
		dialogFunc T.DLGPROC,
		initParam T.LPARAM) (T.INT_PTR, error)

	DialogBoxParam func(
		instance T.HINSTANCE,
		templateName VString,
		wndParent T.HWND,
		dialogFunc T.DLGPROC,
		initParam T.LPARAM) (T.INT_PTR, error)

	DisableProcessWindowsGhosting func()

	DispatchMessage func(msg *T.MSG) (T.LRESULT, error)

	DlgDirList func(
		dlg T.HWND,
		pathSpec VString,
		idListBox, idStaticPath int,
		fileType T.UINT) (int, error)

	DlgDirListComboBox func(
		dlg T.HWND,
		pathSpec VString,
		idComboBox, idStaticPath int,
		filetype T.UINT) (int, error)

	DlgDirSelectComboBoxEx func(
		dlg T.HWND,
		string_ VString,
		out int,
		comboBox int) (T.BOOL, error)

	DlgDirSelectEx func(
		dlg T.HWND,
		string_ VString,
		count int,
		listBox int) (T.BOOL, error)

	DragDetect func(w T.HWND, pt T.POINT) (T.BOOL, error)

	DragObject func(
		parent T.HWND,
		from T.HWND,
		fmt T.UINT,
		data T.ULONG_PTR,
		cur T.HCURSOR) (T.DWORD, error)

	DrawAnimatedRects func(
		hw T.HWND, ani int, from *T.RECT, to *T.RECT) (T.BOOL, error)

	DrawCaption func(
		w T.HWND, dc T.HDC, rect *T.RECT, flags T.UINT) (T.BOOL, error)

	DrawEdge func(
		hdc T.HDC, qrc *T.RECT, edge T.UINT, grfFlags T.UINT) (T.BOOL, error)

	DrawFocusRect func(dc T.HDC, rc *T.RECT) (T.BOOL, error)

	DrawFrameControl func(T.HDC, *T.RECT, T.UINT, T.UINT) (T.BOOL, error)

	DrawIcon func(dc T.HDC, x, y int, icon T.HICON) (T.BOOL, error)

	DrawIconEx func(
		dc T.HDC,
		xLeft, yTop int,
		icon T.HICON,
		xWidth, yWidth int,
		stepIfAniCur T.UINT,
		brFlickerFreeDraw T.HBRUSH,
		flags T.UINT) (T.BOOL, error)

	DrawMenuBar func(w T.HWND) (T.BOOL, error)

	DrawState func(
		dc T.HDC,
		brFore T.HBRUSH,
		fnCallBack T.DRAWSTATEPROC,
		lData T.LPARAM,
		wData T.WPARAM,
		x, y, cx, cy int,
		uFlags T.UINT) (T.BOOL, error)

	DrawText func(
		dc T.HDC,
		text VString,
		nText int,
		rc *T.RECT,
		format T.UINT) (int, error)

	DrawTextEx func(
		dc T.HDC,
		text VString,
		nText int,
		rc *T.RECT,
		format T.UINT,
		dtp *T.DRAWTEXTPARAMS) (int, error)

	EmptyClipboard func() (T.BOOL, error)

	EnableMenuItem func(
		m T.HMENU, idDEnableItem T.UINT, enable T.UINT) (T.BOOL, error)

	EnableScrollBar func(w T.HWND, sbflags T.UINT, arrows T.UINT) (T.BOOL, error)

	EnableWindow func(w T.HWND, enable T.BOOL) (T.BOOL, error)

	EndDeferWindowPos func(winPosInfo T.HDWP) (T.BOOL, error)

	EndDialog func(dlg T.HWND, result T.INT_PTR) (T.BOOL, error)

	EndMenu func() (T.BOOL, error)

	EndPaint func(w T.HWND, paint *T.PAINTSTRUCT) (T.BOOL, error)

	EndTask func(w T.HWND, shutDown T.BOOL, force T.BOOL) (T.BOOL, error)

	EnumChildWindows func(
		parent T.HWND, fn T.WNDENUMPROC, l T.LPARAM) (T.BOOL, error)

	EnumClipboardFormats func(format T.UINT) (T.UINT, error)

	EnumDesktops func(
		wsta T.HWINSTA, fn T.DESKTOPENUMPROC, l T.LPARAM) (T.BOOL, error)

	EnumDesktopWindows func(
		desktop T.HDESK, fn T.WNDENUMPROC, param T.LPARAM) (T.BOOL, error)

	EnumDisplayDevicesA func(
		device VString,
		devNum T.DWORD,
		displayDevice *T.DISPLAY_DEVICEA,
		flags T.DWORD) (T.BOOL, error)

	EnumDisplayDevicesW func(
		device VString,
		devNum T.DWORD,
		displayDevice *T.DISPLAY_DEVICEW,
		flags T.DWORD) (T.BOOL, error)

	EnumDisplayMonitors func(
		dc T.HDC,
		clip *T.RECT,
		fn T.MONITORENUMPROC,
		data T.LPARAM) (T.BOOL, error)

	EnumDisplaySettingsA func(
		deviceName VString,
		modeNum T.DWORD,
		dm *T.DEVMODEA) (T.BOOL, error)

	EnumDisplaySettingsW func(
		deviceName VString,
		modeNum T.DWORD,
		dm *T.DEVMODEW) (T.BOOL, error)

	EnumDisplaySettingsExA func(
		deviceName VString,
		modeNum T.DWORD,
		dm *T.DEVMODEA,
		flags T.DWORD) (T.BOOL, error)

	EnumDisplaySettingsExW func(
		deviceName VString,
		modeNum T.DWORD,
		dm *T.DEVMODEW,
		flags T.DWORD) (T.BOOL, error)

	EnumPropsA func(w T.HWND, enumFunc T.PROPENUMPROCA) (int, error)

	EnumPropsW func(w T.HWND, enumFunc T.PROPENUMPROCW) (int, error)

	EnumPropsExA func(
		w T.HWND, enumFunc T.PROPENUMPROCEXA, l T.LPARAM) (int, error)

	EnumPropsExW func(
		w T.HWND, enumFunc T.PROPENUMPROCEXW, l T.LPARAM) (int, error)

	EnumThreadWindows func(
		threadId T.DWORD, fn T.WNDENUMPROC, l T.LPARAM) (T.BOOL, error)

	EnumWindows func(enumFunc T.WNDENUMPROC, l T.LPARAM) (T.BOOL, error)

	EnumWindowStations func(
		enumFunc T.WINSTAENUMPROC, l T.LPARAM) (T.BOOL, error)

	EqualRect func(rc1 *T.RECT, rc2 *T.RECT) (T.BOOL, error)

	ExcludeUpdateRgn func(dc T.HDC, w T.HWND) (int, error)

	ExitWindowsEx func(flags T.UINT, reason T.DWORD) (T.BOOL, error)

	FillRect func(dc T.HDC, rc *T.RECT, br T.HBRUSH) (int, error)

	FindWindow func(className VString, windowName VString) (T.HWND, error)

	FindWindowEx func(
		wndParent, wndChildAfter T.HWND,
		class, window VString) (T.HWND, error)

	FlashWindow func(w T.HWND, invert T.BOOL) (T.BOOL, error)

	FlashWindowEx func(fwi *T.FLASHWINFO) (T.BOOL, error)

	FrameRect func(dc T.HDC, rc *T.RECT, br T.HBRUSH) (int, error)

	GetActiveWindow func() (T.HWND, error)

	GetAltTabInfo func(
		w T.HWND,
		item int,
		ati *T.ALTTABINFO,
		itemText OVString,
		size T.UINT) (T.BOOL, error)

	GetAncestor func(w T.HWND, flags T.UINT) (T.HWND, error)

	GetAsyncKeyState func(key int) (T.SHORT, error)

	GetCapture func() (T.HWND, error)

	GetCaretBlinkTime func() (T.UINT, error)

	GetCaretPos func(point *T.POINT) (T.BOOL, error)

	GetClassInfo func(
		instance T.HINSTANCE, className VString, wc *T.WNDCLASS) (T.BOOL, error)

	GetClassInfoEx func(
		instance T.HINSTANCE, class VString, wcx *T.WNDCLASSEX) (T.BOOL, error)

	GetClassLong func(w T.HWND, index int) (T.DWORD, error)

	GetClassLongPtr func(w T.HWND, index int) (T.ULONG_PTR, error)

	//GetClassName func(
	//	w T.HWND, className OVString, maxCount int) (int, error)
	GetClassName func(
		w T.HWND, className *T.WChar, maxCount int) (int, error)

	//GetClassWord is obsolete; instead use:
	//	GetClassLong
	GetClassWord func(w T.HWND, index int) (T.WORD, error)

	GetClientRect func(w T.HWND, rect *T.RECT) (T.BOOL, error)

	GetClipboardData func(format T.UINT) (T.HANDLE, error)

	GetClipboardFormatName func(
		format T.UINT, formatName VString, maxCount int) (int, error)

	GetClipboardOwner func() (T.HWND, error)

	GetClipboardSequenceNumber func() (T.DWORD, error)

	GetClipboardViewer func() (T.HWND, error)

	GetClipCursor func(rect *T.RECT) (T.BOOL, error)

	GetComboBoxInfo func(w T.HWND, cbi *T.COMBOBOXINFO) (T.BOOL, error)

	GetCursor func() (T.HCURSOR, error)

	GetCursorInfo func(ci *T.CURSORINFO) (T.BOOL, error)

	GetCursorPos func(point *T.POINT) (T.BOOL, error)

	GetDC func(w T.HWND) (T.HDC, error)

	GetDCEx func(w T.HWND, rgnClip T.HRGN, flags T.DWORD) (T.HDC, error)

	GetDesktopWindow func() (T.HWND, error)

	GetDialogBaseUnits func() (T.LONG, error)

	GetDlgCtrlID func(w T.HWND) (int, error)

	GetDlgItem func(dlg T.HWND, idDlgItem int) (T.HWND, error)

	GetDlgItemInt func(
		dlg T.HWND,
		idDlgItem int,
		translated *T.BOOL,
		signed T.BOOL) (T.UINT, error)

	GetDlgItemText func(
		dlg T.HWND,
		idDlgItem int,
		string_ OVString,
		max int) (T.UINT, error)

	GetDoubleClickTime func() (T.UINT, error)

	GetFocus func() (T.HWND, error)

	GetForegroundWindow func() (T.HWND, error)

	GetGuiResources func(process T.HANDLE, flags T.DWORD) (T.DWORD, error)

	GetGUIThreadInfo func(thread T.DWORD, ui *T.GUITHREADINFO) (T.BOOL, error)

	GetIconInfo func(icon T.HICON, iconinfo *T.ICONINFO) (T.BOOL, error)

	GetInputState func() (T.BOOL, error)

	GetKBCodePage func() (T.UINT, error)

	GetKeyboardLayout func(thread T.DWORD) (T.HKL, error)

	GetKeyboardLayoutList func(buff int, list *T.HKL) (int, error)

	GetKeyboardLayoutName func(klid VString) (T.BOOL, error)

	GetKeyboardState func(keyState *T.BYTE) (T.BOOL, error)

	GetKeyboardType func(typeFlag int) (int, error)

	GetKeyNameText func(
		param T.LONG, string_ OVString, size int) (int, error)

	GetKeyState func(virtKey int) (T.SHORT, error)

	GetLastActivePopup func(w T.HWND) (T.HWND, error)

	GetLastInputInfo func(lii *T.LASTINPUTINFO) (T.BOOL, error)

	GetLayeredWindowAttributes func(
		w T.HWND, crKey *T.COLORREF, alpha *T.BYTE, flags *T.DWORD) (T.BOOL, error)

	GetListBoxInfo func(w T.HWND) (T.DWORD, error)

	GetMenu func(w T.HWND) (T.HMENU, error)

	GetMenuBarInfo func(
		w T.HWND, object, item T.LONG, mbi *T.MENUBARINFO) (T.BOOL, error)

	GetMenuCheckMarkDimensions func() (T.LONG, error)

	GetMenuContextHelpId func(T.HMENU) (T.DWORD, error)

	GetMenuDefaultItem func(
		m T.HMENU, byPos T.UINT, gmdiFlags T.UINT) (T.UINT, error)

	GetMenuInfo func(T.HMENU, *T.MENUINFO) (T.BOOL, error)

	GetMenuItemCount func(m T.HMENU) (int, error)

	GetMenuItemID func(m T.HMENU, pos int) (T.UINT, error)

	GetMenuItemInfo func(
		m T.HMENU,
		item T.UINT,
		byPosition T.BOOL,
		mii *T.MENUITEMINFO) (T.BOOL, error)

	GetMenuItemRect func(
		w T.HWND, m T.HMENU, item T.UINT, rcItem *T.RECT) (T.BOOL, error)

	GetMenuState func(m T.HMENU, id T.UINT, flags T.UINT) (T.UINT, error)

	GetMenuString func(
		m T.HMENU,
		idItem T.UINT,
		string_ OVString,
		max int,
		flags T.UINT) (int, error)

	GetMessage func(
		m *T.MSG, w T.HWND, filterMin, filterMax T.UINT) (T.BOOL, error)

	GetMessageExtraInfo func() (T.LPARAM, error)

	GetMessagePos func() (T.DWORD, error)

	GetMessageTime func() (T.LONG, error)

	GetMonitorInfo func(monitor T.HMONITOR, mi *T.MONITORINFO) (T.BOOL, error)

	GetMouseMovePointsEx func(
		size T.UINT,
		pt, ptBuf *T.MOUSEMOVEPOINT,
		bufPoints int,
		resolution T.DWORD) (int, error)

	GetNextDlgGroupItem func(
		dlg, ctl T.HWND, previous T.BOOL) (T.HWND, error)

	GetNextDlgTabItem func(
		dlg, ctl T.HWND, previous T.BOOL) (T.HWND, error)

	GetOpenClipboardWindow func() (T.HWND, error)

	GetParent func(w T.HWND) (T.HWND, error)

	GetPriorityClipboardFormat func(
		formatPriorityList *T.UINT, formats int) (int, error)

	GetProcessDefaultLayout func(defaultLayout *T.DWORD) (T.BOOL, error)

	GetProcessWindowStation func() (T.HWINSTA, error)

	GetProp func(w T.HWND, string_ VString) (T.HANDLE, error)

	GetQueueStatus func(flags T.UINT) (T.DWORD, error)

	GetRawInputBuffer func(
		data *T.RAWINPUT, size *T.UINT, sizeHeader T.UINT) (T.UINT, error)

	GetRawInputData func(
		rawInput T.HRAWINPUT,
		command T.UINT,
		data *T.VOID,
		size *T.UINT,
		sizeHeader T.UINT) (T.UINT, error)

	GetRawInputDeviceInfo func(
		device T.HANDLE,
		command T.UINT,
		data *T.VOID,
		size *T.UINT) (T.UINT, error)

	GetRawInputDeviceList func(
		ridl *T.RAWINPUTDEVICELIST,
		numDevices *T.UINT,
		size T.UINT) (T.UINT, error)

	GetRegisteredRawInputDevices func(
		rid *T.RAWINPUTDEVICE, numDevices *T.UINT, size T.UINT) (T.UINT, error)

	GetScrollBarInfo func(
		w T.HWND, object T.LONG, sbi *T.SCROLLBARINFO) (T.BOOL, error)

	GetScrollInfo func(w T.HWND, bar int, si *T.SCROLLINFO) (T.BOOL, error)

	GetScrollPos func(w T.HWND, bar int) (int, error)

	GetScrollRange func(
		w T.HWND, bar int, minPos *T.INT, maxPos *T.INT) (T.BOOL, error)

	GetShellWindow func() (T.HWND, error)

	GetSubMenu func(m T.HMENU, pos int) (T.HMENU, error)

	GetSysColor func(index int) (T.DWORD, error)

	GetSysColorBrush func(index int) (T.HBRUSH, error)

	GetSystemMenu func(w T.HWND, revert T.BOOL) (T.HMENU, error)

	GetSystemMetrics func(index int) (int, error)

	GetTabbedTextExtent func(
		dc T.HDC,
		s VString,
		count, tabPositions int,
		tabStopPositions *T.INT) (T.DWORD, error)

	GetThreadDesktop func(threadId T.DWORD) (T.HDESK, error)

	GetTitleBarInfo func(w T.HWND, ti *T.TITLEBARINFO) (T.BOOL, error)

	GetTopWindow func(w T.HWND) (T.HWND, error)

	GetUpdateRect func(w T.HWND, rect *T.RECT, erase T.BOOL) (T.BOOL, error)

	GetUpdateRgn func(w T.HWND, rgn T.HRGN, erase T.BOOL) (int, error)

	GetUserObjectInformation func(
		obj T.HANDLE,
		index int,
		info *T.VOID,
		length T.DWORD,
		lengthNeeded *T.DWORD) (T.BOOL, error)

	GetUserObjectSecurity func(
		obj T.HANDLE,
		siRequested *T.SECURITY_INFORMATION,
		sd *T.SECURITY_DESCRIPTOR,
		length T.DWORD,
		lengthNeeded *T.DWORD) (T.BOOL, error)

	GetWindow func(w T.HWND, cmd T.UINT) (T.HWND, error)

	GetWindowContextHelpId func(T.HWND) (T.DWORD, error)

	GetWindowDC func(w T.HWND) (T.HDC, error)

	GetWindowInfo func(w T.HWND, wi *T.WINDOWINFO) (T.BOOL, error)

	GetWindowLong func(w T.HWND, index int) (T.LONG, error)

	GetWindowLongPtr func(w T.HWND, index int) (T.LONG_PTR, error)

	GetWindowModuleFileName func(
		w T.HWND, fileName OVString, fileNameMax T.UINT) (T.UINT, error)

	GetWindowPlacement func(w T.HWND, wndpl *T.WINDOWPLACEMENT) (T.BOOL, error)

	GetWindowRect func(w T.HWND, rect *T.RECT) (T.BOOL, error)

	GetWindowRgn func(w T.HWND, rgn T.HRGN) (int, error)

	GetWindowRgnBox func(w T.HWND, rc *T.RECT) (int, error)

	// GetWindowText func(w T.HWND, s OVString, maxCount int) (int, error)
	GetWindowText func(w T.HWND, s *T.WChar, maxCount int) (int, error)

	GetWindowTextLength func(w T.HWND) (int, error)

	GetWindowThreadProcessId func(
		w T.HWND, processId *T.DWORD) (T.DWORD, error)

	GetWindowWord func(w T.HWND, index int) (T.WORD, error)

	GrayString func(
		dc T.HDC,
		brush T.HBRUSH,
		outputFunc T.GRAYSTRINGPROC,
		data T.LPARAM,
		count, x, y, width, height int) (T.BOOL, error)

	HideCaret func(w T.HWND) (T.BOOL, error)

	HiliteMenuItem func(
		w T.HWND, m T.HMENU, idHiliteItem T.UINT, hilite T.UINT) (T.BOOL, error)

	InflateRect func(rc *T.RECT, x, y int) (T.BOOL, error)

	InSendMessage func() (T.BOOL, error)

	InSendMessageEx func(reserved *T.VOID) (T.DWORD, error)

	InsertMenu func(
		m T.HMENU,
		position T.UINT,
		flags T.UINT,
		idNewItem T.UINT_PTR,
		NewItem T.AString) (T.BOOL, error)

	InsertMenuItem func(
		m T.HMENU,
		item T.UINT,
		byPosition T.BOOL,
		mii T.MENUITEMINFO) (T.BOOL, error)

	InternalGetWindowText func(w T.HWND, s T.OWString, size int) (int, error)

	IntersectRect func(dst, src1, src2 *T.RECT) (T.BOOL, error)

	InvalidateRect func(w T.HWND, rect *T.RECT, erase T.BOOL) (T.BOOL, error)

	InvalidateRgn func(w T.HWND, rgn T.HRGN, erase T.BOOL) (T.BOOL, error)

	InvertRect func(dc T.HDC, rc *T.RECT) (T.BOOL, error)

	IsCharAlpha func(ch T.VChar) (T.BOOL, error)

	IsCharAlphaNumeric func(ch T.VChar) (T.BOOL, error)

	IsCharLower func(ch T.VChar) (T.BOOL, error)

	IsCharUpper func(ch T.VChar) (T.BOOL, error)

	IsChild func(wndParent T.HWND, w T.HWND) (T.BOOL, error)

	IsClipboardFormatAvailable func(format T.UINT) (T.BOOL, error)

	IsDialogMessage func(dlg T.HWND, msg *T.MSG) (T.BOOL, error)

	IsDlgButtonChecked func(dlg T.HWND, idButton int) (T.UINT, error)

	IsGUIThread func(convert T.BOOL) (T.BOOL, error)

	IsHungAppWindow func(w T.HWND) (T.BOOL, error)

	IsIconic func(w T.HWND) (T.BOOL, error)

	IsMenu func(m T.HMENU) (T.BOOL, error)

	IsRectEmpty func(rc *T.RECT) (T.BOOL, error)

	IsWindow func(w T.HWND) (T.BOOL, error)

	IsWindowEnabled func(w T.HWND) (T.BOOL, error)

	IsWindowUnicode func(w T.HWND) (T.BOOL, error)

	IsWindowVisible func(w T.HWND) (T.BOOL, error)

	IsWinEventHookInstalled func(event T.DWORD) (T.BOOL, error)

	IsWow64Message func() (T.BOOL, error)

	IsZoomed func(w T.HWND) (T.BOOL, error)

	Keybd_event func(
		vk T.BYTE, scan T.BYTE, flags T.DWORD, extraInfo T.ULONG_PTR)

	KillTimer func(w T.HWND, idEvent T.UINT_PTR) (T.BOOL, error)

	LoadAccelerators func(i T.HINSTANCE, tableName VString) (T.HACCEL, error)

	LoadBitmap func(i T.HINSTANCE, bitmapName VString) (T.HBITMAP, error)

	LoadCursor func(i T.HINSTANCE, cursorName PVString) (T.HCURSOR, error)

	LoadCursorFromFile func(fileName VString) (T.HCURSOR, error)

	LoadIcon func(i T.HINSTANCE, iconName VString) (T.HICON, error)

	LoadImage func(
		i T.HINSTANCE,
		name VString,
		type_ T.UINT,
		x, y int,
		load T.UINT) (T.HANDLE, error)

	LoadKeyboardLayout func(klid VString, flags T.UINT) (T.HKL, error)

	LoadMenu func(instance T.HINSTANCE, menuName VString) (T.HMENU, error)

	//TODO(t):How does this work (*T.VOID underneath)
	LoadMenuIndirect func(mt *T.MENUTEMPLATE) (T.HMENU, error)

	LoadString func(
		i T.HINSTANCE, id T.UINT, s VString, size int) (int, error)

	LockSetForegroundWindow func(lockCode T.UINT) (T.BOOL, error)

	LockWindowUpdate func(wndLock T.HWND) (T.BOOL, error)

	LockWorkStation func() (T.BOOL, error)

	LookupIconIdFromDirectory func(
		resbits *T.BYTE, icon T.BOOL) (int, error)

	LookupIconIdFromDirectoryEx func(
		resbits *T.BYTE,
		icon T.BOOL,
		xDesired, yDesired int,
		flags T.UINT) (int, error)

	MapDialogRect func(dlg T.HWND, rect *T.RECT) (T.BOOL, error)

	MapVirtualKey func(code T.UINT, mapType T.UINT) (T.UINT, error)

	MapVirtualKeyEx func(code T.UINT, mapType T.UINT, hkl T.HKL) (T.UINT, error)

	MapWindowPoints func(
		from T.HWND, to T.HWND, p *T.POINT, nPoints T.UINT) (int, error)

	MenuItemFromPoint func(w T.HWND, m T.HMENU, ptScreen T.POINT) (int, error)

	MessageBeep func(type_ T.UINT) (T.BOOL, error)

	MessageBox func(
		w T.HWND, text, caption VString, t T.MSGBOX_TYPE) (int, error)

	MessageBoxEx func(
		w T.HWND,
		text, caption VString,
		t T.MSGBOX_TYPE,
		languageId T.WORD) (int, error)

	MessageBoxIndirect func(mbp *T.MSGBOXPARAMS) (int, error)

	ModifyMenu func(
		mnu T.HMENU,
		position, flags T.UINT,
		idNewItem T.UINT_PTR,
		NewItem VString) (T.BOOL, error)

	MonitorFromPoint func(pt T.POINT, flags T.DWORD) (T.HMONITOR, error)

	MonitorFromRect func(rc *T.RECT, flags T.DWORD) (T.HMONITOR, error)

	MonitorFromWindow func(w T.HWND, flags T.DWORD) (T.HMONITOR, error)

	Mouse_event func(
		flags, x, y, data T.DWORD, extraInfo T.ULONG_PTR)

	MoveWindow func(
		w T.HWND, x, y, width, height int, repaint T.BOOL) (T.BOOL, error)

	MsgWaitForMultipleObjects func(
		count T.DWORD,
		h *T.HANDLE,
		waitAll T.BOOL,
		milliseconds T.DWORD,
		wakeMask T.DWORD) (T.DWORD, error)

	MsgWaitForMultipleObjectsEx func(
		count T.DWORD,
		h *T.HANDLE,
		milliseconds, wakeMask, flags T.DWORD) (T.DWORD, error)

	NotifyWinEvent func(
		event T.DWORD, w T.HWND, object, child T.LONG)

	OemKeyScan func(oemChar T.WORD) (T.DWORD, error)

	OemToChar func(src VString, dst OVString) (T.BOOL, error)

	OemToCharBuff func(
		src VString, dst OVString, dstLength T.DWORD) (T.BOOL, error)

	OffsetRect func(rc *T.RECT, x, y int) (T.BOOL, error)

	OpenClipboard func(wndNewOwner T.HWND) (T.BOOL, error)

	OpenDesktop func(
		desktop VString,
		flags T.DWORD,
		inherit T.BOOL,
		desiredAccess T.ACCESS_MASK) (T.HDESK, error)

	OpenIcon func(w T.HWND) (T.BOOL, error)

	OpenInputDesktop func(
		flags T.DWORD,
		inherit T.BOOL,
		desiredAccess T.ACCESS_MASK) (T.HDESK, error)

	OpenWindowStation func(
		winSta VString,
		inherit T.BOOL,
		desiredAccess T.ACCESS_MASK) (T.HWINSTA, error)

	PaintDesktop func(dc T.HDC) (T.BOOL, error)

	PeekMessage func(
		msg *T.MSG,
		w T.HWND,
		msgFilterMin T.UINT,
		msgFilterMax T.UINT,
		removeMsg T.UINT) (T.BOOL, error)

	PostMessage func(
		win T.HWND, msg T.UINT, w T.WPARAM, l T.LPARAM) (T.BOOL, error)

	PostQuitMessage func(exitCode int)

	PostThreadMessage func(
		thread T.DWORD, msg T.UINT, w T.WPARAM, l T.LPARAM) (T.BOOL, error)

	PrintWindow func(w T.HWND, blt T.HDC, flags T.UINT) (T.BOOL, error)

	PrivateExtractIcons func(
		fileName VString,
		iconIndex, xIcon, yIcon int,
		icon *T.HICON,
		iconid *T.UINT,
		icons, flags T.UINT) (T.UINT, error)

	PtInRect func(rc *T.RECT, pt T.POINT) (T.BOOL, error)

	RealChildWindowFromPoint func(
		parent T.HWND, parentClientCoords T.POINT) (T.HWND, error)

	RealGetWindowClass func(
		w T.HWND, className OVString, classNameMax T.UINT) (T.UINT, error)

	RedrawWindow func(
		w T.HWND, rcUpdate *T.RECT, rgnUpdate T.HRGN, flags T.UINT) (T.BOOL, error)

	RegisterClass func(*T.WNDCLASS) (T.ATOM, error)

	RegisterClassEx func(*T.WNDCLASSEX) (T.ATOM, error)

	RegisterClipboardFormat func(format VString) (T.UINT, error)

	RegisterDeviceNotification func(
		recipient T.HANDLE,
		notificationFilter *T.VOID, flags T.DWORD) (T.HDEVNOTIFY, error)

	RegisterHotKey func(
		w T.HWND, id int, modifiers T.UINT, vk T.UINT) (T.BOOL, error)

	RegisterRawInputDevices func(
		rids *T.RAWINPUTDEVICE, numDevices T.UINT, size T.UINT) (T.BOOL, error)

	RegisterShellHookWindow func(T.HWND) (T.BOOL, error)

	RegisterWindowMessage func(VString) (T.UINT, error)

	ReleaseCapture func() (T.BOOL, error)

	ReleaseDC func(T.HWND, T.HDC) (int, error)

	RemoveMenu func(m T.HMENU, position T.UINT, flags T.UINT) (T.BOOL, error)

	RemoveProp func(T.HWND, VString) (T.HANDLE, error)

	ReplyMessage func(T.LRESULT) (T.BOOL, error)

	ScreenToClient func(T.HWND, *T.POINT) (T.BOOL, error)

	ScrollDC func(
		dc T.HDC,
		x, y int,
		rcScroll, rcClip *T.RECT,
		rgnUpdate T.HRGN,
		rcUpdate *T.RECT) (T.BOOL, error)

	ScrollWindow func(
		w T.HWND, xAmount, yAmount int,
		rect, clipRect *T.RECT) (T.BOOL, error)

	ScrollWindowEx func(
		w T.HWND,
		x, y int,
		scroll, clip *T.RECT,
		rgnUpdate T.HRGN,
		rcUpdate *T.RECT,
		flags T.UINT) (int, error)

	SendDlgItemMessage func(
		dlg T.HWND,
		idDlgItem int,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM) (T.LRESULT, error)

	SendInput func(nInputs T.UINT, i *T.INPUT, size int) (T.UINT, error)

	SendMessage func(
		win T.HWND, msg T.UINT, w T.WPARAM, l T.LPARAM) (T.LRESULT, error)

	SendMessageCallback func(
		win T.HWND,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM,
		resultCallBack T.SENDASYNCPROC,
		data T.ULONG_PTR) (T.BOOL, error)

	SendMessageTimeout func(
		win T.HWND,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM,
		flags, timeout T.UINT,
		result *T.DWORD_PTR) (T.LRESULT, error)

	SendNotifyMessage func(
		win T.HWND, msg T.UINT, w T.WPARAM, l T.LPARAM) (T.BOOL, error)

	SetActiveWindow func(w T.HWND) (T.HWND, error)

	SetCapture func(w T.HWND) (T.HWND, error)

	SetCaretBlinkTime func(mSeconds T.UINT) (T.BOOL, error)

	SetCaretPos func(x, y int) (T.BOOL, error)

	//SetClassWord is obsolete; instead use:
	//	SetClassLong
	SetClassLong func(w T.HWND, index int, newLong T.LONG) (T.DWORD, error)

	SetClassLongPtr func(
		w T.HWND, index int, newLong T.LONG_PTR) (T.ULONG_PTR, error)

	SetClassWord func(w T.HWND, index int, newWord T.WORD) (T.WORD, error)

	SetClipboardData func(format T.UINT, mem T.HANDLE) (T.HANDLE, error)

	SetClipboardViewer func(wndNewViewer T.HWND) (T.HWND, error)

	SetCursor func(cursor T.HCURSOR) (T.HCURSOR, error)

	SetCursorPos func(x, y int) (T.BOOL, error)

	SetDebugErrorLevel func(level T.DWORD)

	SetDlgItemInt func(
		dlg T.HWND, idDlgItem int, value T.UINT, signed T.BOOL) (T.BOOL, error)

	SetDlgItemText func(
		dlg T.HWND, idDlgItem int, s VString) (T.BOOL, error)

	SetDoubleClickTime func(T.UINT) (T.BOOL, error)

	SetFocus func(T.HWND) (T.HWND, error)

	SetForegroundWindow func(T.HWND) (T.BOOL, error)

	SetKeyboardState func(keyState *T.BYTE) (T.BOOL, error)

	SetLastErrorEx func(errCode T.DWORD, type_ T.DWORD)

	SetLayeredWindowAttributes func(
		w T.HWND, crKey T.COLORREF, alpha T.BYTE, flags T.DWORD) (T.BOOL, error)

	SetMenu func(T.HWND, T.HMENU) (T.BOOL, error)

	SetMenuContextHelpId func(T.HMENU, T.DWORD) (T.BOOL, error)

	SetMenuDefaultItem func(m T.HMENU, item T.UINT, byPos T.UINT) (T.BOOL, error)

	SetMenuInfo func(T.HMENU, T.MENUINFO) (T.BOOL, error)

	SetMenuItemBitmaps func(
		m T.HMENU,
		position, flags T.UINT,
		bitmapUnchecked, bitmapChecked T.HBITMAP) (T.BOOL, error)

	SetMenuItemInfo func(
		m T.HMENU,
		item T.UINT,
		byPositon T.BOOL,
		mii T.MENUITEMINFO) (T.BOOL, error)

	SetMessageExtraInfo func(T.LPARAM) (T.LPARAM, error)

	SetMessageQueue func(messagesMax int) (T.BOOL, error)

	SetParent func(child T.HWND, newParent T.HWND) (T.HWND, error)

	SetProcessDefaultLayout func(defaultLayout T.DWORD) (T.BOOL, error)

	SetProcessWindowStation func(T.HWINSTA) (T.BOOL, error)

	SetProp func(w T.HWND, s VString, data T.HANDLE) (T.BOOL, error)

	SetRect func(rc *T.RECT, xLeft, yTop, xRight, yBottom int) (T.BOOL, error)

	SetRectEmpty func(*T.RECT) (T.BOOL, error)

	SetScrollInfo func(
		w T.HWND, bar int, si *T.SCROLLINFO, redraw T.BOOL) (int, error)

	SetScrollPos func(w T.HWND, bar, pos, redraw T.BOOL) (int, error)

	SetScrollRange func(
		w T.HWND, bar, minPos, maxPos int, redraw T.BOOL) (T.BOOL, error)

	SetSysColors func(
		nElements int, elements *T.INT, rgbValues *T.COLORREF) (T.BOOL, error)

	SetSystemCursor func(cur T.HCURSOR, id T.DWORD) (T.BOOL, error)

	SetThreadDesktop func(desktop T.HDESK) (T.BOOL, error)

	SetTimer func(
		w T.HWND,
		idEvent T.UINT_PTR,
		elapse T.UINT,
		fn T.TIMERPROC) (T.UINT_PTR, error)

	SetUserObjectInformation func(
		obj T.HANDLE, index int, info *T.VOID, length T.DWORD) (T.BOOL, error)

	SetUserObjectSecurity func(
		obj T.HANDLE,
		requested *T.SECURITY_INFORMATION,
		sid *T.SECURITY_DESCRIPTOR) (T.BOOL, error)

	SetWindowContextHelpId func(T.HWND, T.DWORD) (T.BOOL, error)

	SetWindowLong func(w T.HWND, index int, newLong T.LONG) (T.LONG, error)

	SetWindowLongPtr func(
		w T.HWND, index int, newLong T.LONG_PTR) (T.LONG_PTR, error)

	SetWindowPlacement func(w T.HWND, wndpl *T.WINDOWPLACEMENT) (T.BOOL, error)

	SetWindowPos func(
		w T.HWND,
		insertAfter T.HWND,
		x, y, cx, cy int,
		flags T.UINT) (T.BOOL, error)

	SetWindowRgn func(w T.HWND, rgn T.HRGN, redraw T.BOOL) (int, error)

	SetWindowsHook func(filterType int, fn T.HOOKPROC) (T.HHOOK, error)

	SetWindowsHookEx func(
		hook int,
		fn T.HOOKPROC,
		mod T.HINSTANCE,
		threadId T.DWORD) (T.HHOOK, error)

	SetWindowText func(w T.HWND, s VString) (T.BOOL, error)

	SetWindowWord func(w T.HWND, index int, newWord T.WORD) (T.WORD, error)

	SetWinEventHook func(
		min, max T.DWORD,
		modWinEventProc T.HMODULE,
		fn T.WINEVENTPROC,
		process, thread T.DWORD,
		f T.WINEVENT_FLAGS) (T.HWINEVENTHOOK, error)

	ShowCaret func(w T.HWND) (T.BOOL, error)

	ShowCursor func(show T.BOOL) (int, error)

	ShowOwnedPopups func(w T.HWND, show T.BOOL) (T.BOOL, error)

	ShowScrollBar func(w T.HWND, bar int, show T.BOOL) (T.BOOL, error)

	ShowWindow func(w T.HWND, cmdShow int) (T.BOOL, error)

	ShowWindowAsync func(w T.HWND, cmdShow int) (T.BOOL, error)

	SubtractRect func(dst, src1, src2 *T.RECT) (T.BOOL, error)

	SwapMouseButton func(swap T.BOOL) (T.BOOL, error)

	SwitchDesktop func(desktop T.HDESK) (T.BOOL, error)

	SwitchToThisWindow func(w T.HWND, unknown T.BOOL)

	SystemParametersInfo func(
		action, nParam T.UINT, param *T.VOID, winIni T.UINT) (T.BOOL, error)

	TabbedTextOut func(
		h T.HDC,
		x, y int,
		s VString,
		count, tabPositions int,
		tabStopPositions *T.INT,
		tabOrigin int) (T.LONG, error)

	TileWindows func(
		parent T.HWND,
		how T.UINT,
		rect *T.RECT,
		nKids T.UINT,
		kids *T.HWND) (T.WORD, error)

	ToAscii func(
		virtKey, scanCode T.UINT,
		keyState *T.BYTE,
		char *T.WORD,
		flags T.UINT) (int, error)

	ToAsciiEx func(
		virtKey, scanCode T.UINT,
		keyState *T.BYTE,
		char *T.WORD,
		flags T.UINT,
		hkl T.HKL) (int, error)

	ToUnicode func(
		virtKey, scanCode T.UINT,
		keyState *T.BYTE,
		buff T.OWString,
		size int,
		flags T.UINT) (int, error)

	ToUnicodeEx func(
		virtKey, scanCode T.UINT,
		keyState *T.BYTE,
		buff T.WString,
		size int,
		flags T.UINT,
		h T.HKL) (int, error)

	TrackMouseEvent func(eventTrack *T.TRACKMOUSEEVENT) (T.BOOL, error)

	TrackPopupMenu func(
		m T.HMENU,
		flags T.UINT,
		x, y, reserved int,
		w T.HWND,
		rect *T.RECT) (T.BOOL, error)

	TrackPopupMenuEx func(
		T.HMENU, T.UINT, int, int, T.HWND, *T.TPMPARAMS) (T.BOOL, error)

	TranslateAccelerator func(
		w T.HWND, accTable T.HACCEL, msg *T.MSG) (int, error)

	TranslateMDISysAccel func(wndClient T.HWND, msg *T.MSG) (T.BOOL, error)

	TranslateMessage func(msg *T.MSG) (T.BOOL, error)

	UnhookWindowsHook func(code int, filterProc T.HOOKPROC) (T.BOOL, error)

	UnhookWindowsHookEx func(hk T.HHOOK) (T.BOOL, error)

	UnhookWinEvent func(winEventHook T.HWINEVENTHOOK) (T.BOOL, error)

	UnionRect func(dst, src1, src2 *T.RECT) (T.BOOL, error)

	UnloadKeyboardLayout func(hkl T.HKL) (T.BOOL, error)

	UnregisterClass func(
		className VString, instance T.HINSTANCE) (T.BOOL, error)

	UnregisterDeviceNotification func(handle T.HDEVNOTIFY) (T.BOOL, error)

	UnregisterHotKey func(w T.HWND, id int) (T.BOOL, error)

	UpdateLayeredWindow func(
		w T.HWND,
		dst T.HDC,
		ptDst *T.POINT,
		s *T.SIZE,
		src T.HDC,
		ptSrc *T.POINT,
		crKey T.COLORREF,
		blend *T.BLENDFUNCTION,
		flags T.DWORD) (T.BOOL, error)

	UpdateLayeredWindowIndirect func(
		w T.HWND, ulwi *T.UPDATELAYEREDWINDOWINFO) (T.BOOL, error)

	UpdateWindow func(w T.HWND) (T.BOOL, error)

	UserHandleGrantAccess func(
		user T.HANDLE, job T.HANDLE, grant T.BOOL) (T.BOOL, error)

	ValidateRect func(w T.HWND, rect *T.RECT) (T.BOOL, error)

	ValidateRgn func(w T.HWND, rgn T.HRGN) (T.BOOL, error)

	VkKeyScan func(ch T.Char) (T.SHORT, error)

	VkKeyScanEx func(ch T.Char, dwhkl T.HKL) (T.SHORT, error)

	WaitForInputIdle func(
		process T.HANDLE, milliseconds T.DWORD) (T.DWORD, error)

	WaitMessage func() (T.BOOL, error)

	WindowFromDC func(dc T.HDC) (T.HWND, error)

	WindowFromPoint func(point T.POINT) (T.HWND, error)

	WinHelp func(
		main T.HWND,
		help VString,
		command T.UINT,
		data T.ULONG_PTR) (T.BOOL, error)

	//Wsprintf func(POVString, VString, ...VArg) (int, error)
	Wsprintf func(*T.WChar, VString, ...VArg) (int, error)

	Wvsprintf func(VString, VString, va_list) (int, error)
)

//TODO(t): Not xp32
//{"IsWow64Message",&IsWow64Message},
//{"UpdateLayeredWindowIndirect",&UpdateLayeredWindowIndirect},

var WinUserApis = outside.Apis{
	{"ActivateKeyboardLayout", &ActivateKeyboardLayout},
	{"AdjustWindowRect", &AdjustWindowRect},
	{"AdjustWindowRectEx", &AdjustWindowRectEx},
	{"AllowSetForegroundWindow", &AllowSetForegroundWindow},
	{"AnimateWindow", &AnimateWindow},
	{"AnyPopup", &AnyPopup},
	{"ArrangeIconicWindows", &ArrangeIconicWindows},
	{"AttachThreadInput", &AttachThreadInput},
	{"BeginDeferWindowPos", &BeginDeferWindowPos},
	{"BeginPaint", &BeginPaint},
	{"BringWindowToTop", &BringWindowToTop},
	{"CallMsgFilter", &CallMsgFilter},
	{"CallNextHookEx", &CallNextHookEx},
	{"CascadeWindows", &CascadeWindows},
	{"ChangeClipboardChain", &ChangeClipboardChain},
	{"CheckDlgButton", &CheckDlgButton},
	{"CheckMenuItem", &CheckMenuItem},
	{"CheckMenuRadioItem", &CheckMenuRadioItem},
	{"CheckRadioButton", &CheckRadioButton},
	{"ChildWindowFromPoint", &ChildWindowFromPoint},
	{"ChildWindowFromPointEx", &ChildWindowFromPointEx},
	{"ClientToScreen", &ClientToScreen},
	{"ClipCursor", &ClipCursor},
	{"CloseClipboard", &CloseClipboard},
	{"CloseDesktop", &CloseDesktop},
	{"CloseWindow", &CloseWindow},
	{"CloseWindowStation", &CloseWindowStation},
	{"CopyIcon", &CopyIcon},
	{"CopyImage", &CopyImage},
	{"CopyRect", &CopyRect},
	{"CountClipboardFormats", &CountClipboardFormats},
	{"CreateCaret", &CreateCaret},
	{"CreateCursor", &CreateCursor},
	{"CreateIcon", &CreateIcon},
	{"CreateIconFromResource", &CreateIconFromResource},
	{"CreateIconFromResourceEx", &CreateIconFromResourceEx},
	{"CreateIconIndirect", &CreateIconIndirect},
	{"CreateMenu", &CreateMenu},
	{"CreatePopupMenu", &CreatePopupMenu},
	{"DeferWindowPos", &DeferWindowPos},
	{"DefRawInputProc", &DefRawInputProc},
	{"DeleteMenu", &DeleteMenu},
	{"DeregisterShellHookWindow", &DeregisterShellHookWindow},
	{"DestroyAcceleratorTable", &DestroyAcceleratorTable},
	{"DestroyCaret", &DestroyCaret},
	{"DestroyCursor", &DestroyCursor},
	{"DestroyIcon", &DestroyIcon},
	{"DestroyMenu", &DestroyMenu},
	{"DestroyWindow", &DestroyWindow},
	{"DisableProcessWindowsGhosting", &DisableProcessWindowsGhosting},
	{"DragDetect", &DragDetect},
	{"DragObject", &DragObject},
	{"DrawAnimatedRects", &DrawAnimatedRects},
	{"DrawCaption", &DrawCaption},
	{"DrawEdge", &DrawEdge},
	{"DrawFocusRect", &DrawFocusRect},
	{"DrawFrameControl", &DrawFrameControl},
	{"DrawIcon", &DrawIcon},
	{"DrawIconEx", &DrawIconEx},
	{"DrawMenuBar", &DrawMenuBar},
	{"EmptyClipboard", &EmptyClipboard},
	{"EnableMenuItem", &EnableMenuItem},
	{"EnableScrollBar", &EnableScrollBar},
	{"EnableWindow", &EnableWindow},
	{"EndDeferWindowPos", &EndDeferWindowPos},
	{"EndDialog", &EndDialog},
	{"EndMenu", &EndMenu},
	{"EndPaint", &EndPaint},
	{"EndTask", &EndTask},
	{"EnumChildWindows", &EnumChildWindows},
	{"EnumClipboardFormats", &EnumClipboardFormats},
	{"EnumDesktopWindows", &EnumDesktopWindows},
	{"EnumDisplayMonitors", &EnumDisplayMonitors},
	{"EnumThreadWindows", &EnumThreadWindows},
	{"EnumWindows", &EnumWindows},
	{"EqualRect", &EqualRect},
	{"ExcludeUpdateRgn", &ExcludeUpdateRgn},
	{"ExitWindowsEx", &ExitWindowsEx},
	{"FillRect", &FillRect},
	{"FlashWindow", &FlashWindow},
	{"FlashWindowEx", &FlashWindowEx},
	{"FrameRect", &FrameRect},
	{"GetActiveWindow", &GetActiveWindow},
	{"GetAltTabInfo", &GetAltTabInfo},
	{"GetAncestor", &GetAncestor},
	{"GetAsyncKeyState", &GetAsyncKeyState},
	{"GetCapture", &GetCapture},
	{"GetCaretBlinkTime", &GetCaretBlinkTime},
	{"GetCaretPos", &GetCaretPos},
	{"GetClassWord", &GetClassWord},
	{"GetClientRect", &GetClientRect},
	{"GetClipboardData", &GetClipboardData},
	{"GetClipboardOwner", &GetClipboardOwner},
	{"GetClipboardSequenceNumber", &GetClipboardSequenceNumber},
	{"GetClipboardViewer", &GetClipboardViewer},
	{"GetClipCursor", &GetClipCursor},
	{"GetComboBoxInfo", &GetComboBoxInfo},
	{"GetCursor", &GetCursor},
	{"GetCursorInfo", &GetCursorInfo},
	{"GetCursorPos", &GetCursorPos},
	{"GetDC", &GetDC},
	{"GetDCEx", &GetDCEx},
	{"GetDesktopWindow", &GetDesktopWindow},
	{"GetDialogBaseUnits", &GetDialogBaseUnits},
	{"GetDlgCtrlID", &GetDlgCtrlID},
	{"GetDlgItem", &GetDlgItem},
	{"GetDlgItemInt", &GetDlgItemInt},
	{"GetDoubleClickTime", &GetDoubleClickTime},
	{"GetFocus", &GetFocus},
	{"GetForegroundWindow", &GetForegroundWindow},
	{"GetGuiResources", &GetGuiResources},
	{"GetGUIThreadInfo", &GetGUIThreadInfo},
	{"GetIconInfo", &GetIconInfo},
	{"GetInputState", &GetInputState},
	{"GetKBCodePage", &GetKBCodePage},
	{"GetKeyboardLayout", &GetKeyboardLayout},
	{"GetKeyboardLayoutList", &GetKeyboardLayoutList},
	{"GetKeyboardState", &GetKeyboardState},
	{"GetKeyboardType", &GetKeyboardType},
	{"GetKeyState", &GetKeyState},
	{"GetLastActivePopup", &GetLastActivePopup},
	{"GetLastInputInfo", &GetLastInputInfo},
	{"GetLayeredWindowAttributes", &GetLayeredWindowAttributes},
	{"GetListBoxInfo", &GetListBoxInfo},
	{"GetMenu", &GetMenu},
	{"GetMenuBarInfo", &GetMenuBarInfo},
	{"GetMenuCheckMarkDimensions", &GetMenuCheckMarkDimensions},
	{"GetMenuContextHelpId", &GetMenuContextHelpId},
	{"GetMenuDefaultItem", &GetMenuDefaultItem},
	{"GetMenuInfo", &GetMenuInfo},
	{"GetMenuItemCount", &GetMenuItemCount},
	{"GetMenuItemID", &GetMenuItemID},
	{"GetMenuItemRect", &GetMenuItemRect},
	{"GetMenuState", &GetMenuState},
	{"GetMessageExtraInfo", &GetMessageExtraInfo},
	{"GetMessagePos", &GetMessagePos},
	{"GetMessageTime", &GetMessageTime},
	{"GetMouseMovePointsEx", &GetMouseMovePointsEx},
	{"GetNextDlgGroupItem", &GetNextDlgGroupItem},
	{"GetNextDlgTabItem", &GetNextDlgTabItem},
	{"GetOpenClipboardWindow", &GetOpenClipboardWindow},
	{"GetParent", &GetParent},
	{"GetPriorityClipboardFormat", &GetPriorityClipboardFormat},
	{"GetProcessDefaultLayout", &GetProcessDefaultLayout},
	{"GetProcessWindowStation", &GetProcessWindowStation},
	{"GetQueueStatus", &GetQueueStatus},
	{"GetRawInputBuffer", &GetRawInputBuffer},
	{"GetRawInputData", &GetRawInputData},
	{"GetRawInputDeviceList", &GetRawInputDeviceList},
	{"GetRegisteredRawInputDevices", &GetRegisteredRawInputDevices},
	{"GetScrollBarInfo", &GetScrollBarInfo},
	{"GetScrollInfo", &GetScrollInfo},
	{"GetScrollPos", &GetScrollPos},
	{"GetScrollRange", &GetScrollRange},
	{"GetShellWindow", &GetShellWindow},
	{"GetSubMenu", &GetSubMenu},
	{"GetSysColor", &GetSysColor},
	{"GetSysColorBrush", &GetSysColorBrush},
	{"GetSystemMenu", &GetSystemMenu},
	{"GetSystemMetrics", &GetSystemMetrics},
	{"GetThreadDesktop", &GetThreadDesktop},
	{"GetTitleBarInfo", &GetTitleBarInfo},
	{"GetTopWindow", &GetTopWindow},
	{"GetUpdateRect", &GetUpdateRect},
	{"GetUpdateRgn", &GetUpdateRgn},
	{"GetUserObjectSecurity", &GetUserObjectSecurity},
	{"GetWindow", &GetWindow},
	{"GetWindowContextHelpId", &GetWindowContextHelpId},
	{"GetWindowDC", &GetWindowDC},
	{"GetWindowInfo", &GetWindowInfo},
	{"GetWindowModuleFileName", &GetWindowModuleFileName},
	{"GetWindowPlacement", &GetWindowPlacement},
	{"GetWindowRect", &GetWindowRect},
	{"GetWindowRgn", &GetWindowRgn},
	{"GetWindowRgnBox", &GetWindowRgnBox},
	{"GetWindowThreadProcessId", &GetWindowThreadProcessId},
	{"GetWindowWord", &GetWindowWord},
	{"HideCaret", &HideCaret},
	{"HiliteMenuItem", &HiliteMenuItem},
	{"InflateRect", &InflateRect},
	{"InSendMessage", &InSendMessage},
	{"InSendMessageEx", &InSendMessageEx},
	{"InternalGetWindowText", &InternalGetWindowText},
	{"IntersectRect", &IntersectRect},
	{"InvalidateRect", &InvalidateRect},
	{"InvalidateRgn", &InvalidateRgn},
	{"InvertRect", &InvertRect},
	{"IsChild", &IsChild},
	{"IsClipboardFormatAvailable", &IsClipboardFormatAvailable},
	{"IsDlgButtonChecked", &IsDlgButtonChecked},
	{"IsGUIThread", &IsGUIThread},
	{"IsHungAppWindow", &IsHungAppWindow},
	{"IsIconic", &IsIconic},
	{"IsMenu", &IsMenu},
	{"IsRectEmpty", &IsRectEmpty},
	{"IsWindow", &IsWindow},
	{"IsWindowEnabled", &IsWindowEnabled},
	{"IsWindowUnicode", &IsWindowUnicode},
	{"IsWindowVisible", &IsWindowVisible},
	{"IsWinEventHookInstalled", &IsWinEventHookInstalled},
	{"IsZoomed", &IsZoomed},
	{"keybd_event", &Keybd_event},
	{"KillTimer", &KillTimer},
	{"LockSetForegroundWindow", &LockSetForegroundWindow},
	{"LockWindowUpdate", &LockWindowUpdate},
	{"LockWorkStation", &LockWorkStation},
	{"LookupIconIdFromDirectory", &LookupIconIdFromDirectory},
	{"LookupIconIdFromDirectoryEx", &LookupIconIdFromDirectoryEx},
	{"MapDialogRect", &MapDialogRect},
	{"MapWindowPoints", &MapWindowPoints},
	{"MenuItemFromPoint", &MenuItemFromPoint},
	{"MessageBeep", &MessageBeep},
	{"MonitorFromPoint", &MonitorFromPoint},
	{"MonitorFromRect", &MonitorFromRect},
	{"MonitorFromWindow", &MonitorFromWindow},
	{"mouse_event", &Mouse_event},
	{"MoveWindow", &MoveWindow},
	{"MsgWaitForMultipleObjects", &MsgWaitForMultipleObjects},
	{"MsgWaitForMultipleObjectsEx", &MsgWaitForMultipleObjectsEx},
	{"NotifyWinEvent", &NotifyWinEvent},
	{"OemKeyScan", &OemKeyScan},
	{"OffsetRect", &OffsetRect},
	{"OpenClipboard", &OpenClipboard},
	{"OpenIcon", &OpenIcon},
	{"OpenInputDesktop", &OpenInputDesktop},
	{"PaintDesktop", &PaintDesktop},
	{"PostQuitMessage", &PostQuitMessage},
	{"PrintWindow", &PrintWindow},
	{"PtInRect", &PtInRect},
	{"RealChildWindowFromPoint", &RealChildWindowFromPoint},
	{"RealGetWindowClass", &RealGetWindowClass},
	{"RedrawWindow", &RedrawWindow},
	{"RegisterHotKey", &RegisterHotKey},
	{"RegisterRawInputDevices", &RegisterRawInputDevices},
	{"RegisterShellHookWindow", &RegisterShellHookWindow},
	{"ReleaseCapture", &ReleaseCapture},
	{"ReleaseDC", &ReleaseDC},
	{"RemoveMenu", &RemoveMenu},
	{"ReplyMessage", &ReplyMessage},
	{"ScreenToClient", &ScreenToClient},
	{"ScrollDC", &ScrollDC},
	{"ScrollWindow", &ScrollWindow},
	{"ScrollWindowEx", &ScrollWindowEx},
	{"SendInput", &SendInput},
	{"SetActiveWindow", &SetActiveWindow},
	{"SetCapture", &SetCapture},
	{"SetCaretBlinkTime", &SetCaretBlinkTime},
	{"SetCaretPos", &SetCaretPos},
	{"SetClassWord", &SetClassWord},
	{"SetClipboardData", &SetClipboardData},
	{"SetClipboardViewer", &SetClipboardViewer},
	{"SetCursor", &SetCursor},
	{"SetCursorPos", &SetCursorPos},
	{"SetDebugErrorLevel", &SetDebugErrorLevel},
	{"SetDlgItemInt", &SetDlgItemInt},
	{"SetDoubleClickTime", &SetDoubleClickTime},
	{"SetFocus", &SetFocus},
	{"SetForegroundWindow", &SetForegroundWindow},
	{"SetKeyboardState", &SetKeyboardState},
	{"SetLastErrorEx", &SetLastErrorEx},
	{"SetLayeredWindowAttributes", &SetLayeredWindowAttributes},
	{"SetMenu", &SetMenu},
	{"SetMenuContextHelpId", &SetMenuContextHelpId},
	{"SetMenuDefaultItem", &SetMenuDefaultItem},
	{"SetMenuInfo", &SetMenuInfo},
	{"SetMenuItemBitmaps", &SetMenuItemBitmaps},
	{"SetMessageExtraInfo", &SetMessageExtraInfo},
	{"SetMessageQueue", &SetMessageQueue},
	{"SetParent", &SetParent},
	{"SetProcessDefaultLayout", &SetProcessDefaultLayout},
	{"SetProcessWindowStation", &SetProcessWindowStation},
	{"SetRect", &SetRect},
	{"SetRectEmpty", &SetRectEmpty},
	{"SetScrollInfo", &SetScrollInfo},
	{"SetScrollPos", &SetScrollPos},
	{"SetScrollRange", &SetScrollRange},
	{"SetSysColors", &SetSysColors},
	{"SetSystemCursor", &SetSystemCursor},
	{"SetThreadDesktop", &SetThreadDesktop},
	{"SetTimer", &SetTimer},
	{"SetUserObjectSecurity", &SetUserObjectSecurity},
	{"SetWindowContextHelpId", &SetWindowContextHelpId},
	{"SetWindowPlacement", &SetWindowPlacement},
	{"SetWindowPos", &SetWindowPos},
	{"SetWindowRgn", &SetWindowRgn},
	{"SetWindowWord", &SetWindowWord},
	{"SetWinEventHook", &SetWinEventHook},
	{"ShowCaret", &ShowCaret},
	{"ShowCursor", &ShowCursor},
	{"ShowOwnedPopups", &ShowOwnedPopups},
	{"ShowScrollBar", &ShowScrollBar},
	{"ShowWindow", &ShowWindow},
	{"ShowWindowAsync", &ShowWindowAsync},
	{"SubtractRect", &SubtractRect},
	{"SwapMouseButton", &SwapMouseButton},
	{"SwitchDesktop", &SwitchDesktop},
	{"SwitchToThisWindow", &SwitchToThisWindow},
	{"TileWindows", &TileWindows},
	{"ToAscii", &ToAscii},
	{"ToAsciiEx", &ToAsciiEx},
	{"ToUnicode", &ToUnicode},
	{"ToUnicodeEx", &ToUnicodeEx},
	{"TrackMouseEvent", &TrackMouseEvent},
	{"TrackPopupMenu", &TrackPopupMenu},
	{"TrackPopupMenuEx", &TrackPopupMenuEx},
	{"TranslateAccelerator", &TranslateAccelerator},
	{"TranslateMDISysAccel", &TranslateMDISysAccel},
	{"TranslateMessage", &TranslateMessage},
	{"UnhookWindowsHook", &UnhookWindowsHook},
	{"UnhookWindowsHookEx", &UnhookWindowsHookEx},
	{"UnhookWinEvent", &UnhookWinEvent},
	{"UnionRect", &UnionRect},
	{"UnloadKeyboardLayout", &UnloadKeyboardLayout},
	{"UnregisterDeviceNotification", &UnregisterDeviceNotification},
	{"UnregisterHotKey", &UnregisterHotKey},
	{"UpdateLayeredWindow", &UpdateLayeredWindow},
	{"UpdateWindow", &UpdateWindow},
	{"UserHandleGrantAccess", &UserHandleGrantAccess},
	{"ValidateRect", &ValidateRect},
	{"ValidateRgn", &ValidateRgn},
	{"WaitForInputIdle", &WaitForInputIdle},
	{"WaitMessage", &WaitMessage},
	{"WindowFromDC", &WindowFromDC},
	{"WindowFromPoint", &WindowFromPoint},
}

//TODO(t): funcs ending with A and W
var WinUserUnicodeApis = outside.Apis{
	{"AppendMenuW", &AppendMenu},
	{"BroadcastSystemMessageExW", &BroadcastSystemMessageEx},
	{"BroadcastSystemMessageW", &BroadcastSystemMessage},
	{"CallWindowProcW", &CallWindowProc},
	{"ChangeDisplaySettingsExW", &ChangeDisplaySettingsExW},
	{"ChangeDisplaySettingsW", &ChangeDisplaySettingsW},
	{"ChangeMenuW", &ChangeMenu},
	{"CharLowerBuffW", &CharLowerBuff},
	{"CharLowerW", &CharLower},
	{"CharNextW", &CharNext},
	{"CharPrevW", &CharPrev},
	{"CharToOemBuffW", &CharToOemBuff},
	{"CharToOemW", &CharToOem},
	{"CharUpperBuffW", &CharUpperBuff},
	{"CharUpperW", &CharUpper},
	{"CopyAcceleratorTableW", &CopyAcceleratorTable},
	{"CreateAcceleratorTableW", &CreateAcceleratorTable},
	{"CreateDesktopW", &CreateDesktopW},
	{"CreateDialogIndirectParamW", &CreateDialogIndirectParam},
	{"CreateDialogParamW", &CreateDialogParam},
	{"CreateMDIWindowW", &CreateMDIWindow},
	{"CreateWindowExW", &CreateWindowEx},
	{"CreateWindowStationW", &CreateWindowStation},
	{"DefDlgProcW", &DefDlgProc},
	{"DefFrameProcW", &DefFrameProc},
	{"DefMDIChildProcW", &DefMDIChildProc},
	{"DefWindowProcW", &DefWindowProc},
	{"DispatchMessageW", &DispatchMessage},
	{"DlgDirListComboBoxW", &DlgDirListComboBox},
	{"DlgDirListW", &DlgDirList},
	{"DlgDirSelectComboBoxExW", &DlgDirSelectComboBoxEx},
	{"DlgDirSelectExW", &DlgDirSelectEx},
	{"DrawStateW", &DrawState},
	{"DrawTextExW", &DrawTextEx},
	{"DrawTextW", &DrawText},
	{"EnumDesktopsW", &EnumDesktops},
	{"EnumDisplayDevicesW", &EnumDisplayDevicesW},
	{"EnumDisplaySettingsExW", &EnumDisplaySettingsExW},
	{"EnumDisplaySettingsW", &EnumDisplaySettingsW},
	{"EnumPropsExW", &EnumPropsExW},
	{"EnumPropsW", &EnumPropsW},
	{"EnumWindowStationsW", &EnumWindowStations},
	{"FindWindowExW", &FindWindowEx},
	{"FindWindowW", &FindWindow},
	{"GetClassInfoExW", &GetClassInfoEx},
	{"GetClassInfoW", &GetClassInfo},
	{"GetClassLongW", &GetClassLong},
	{"GetClassLongW", &GetClassLongPtr}, //NOTE(t): Not on xp32
	{"GetClassNameW", &GetClassName},
	{"GetClipboardFormatNameW", &GetClipboardFormatName},
	{"GetDlgItemTextW", &GetDlgItemText},
	{"GetKeyboardLayoutNameW", &GetKeyboardLayoutName},
	{"GetKeyNameTextW", &GetKeyNameText},
	{"GetMenuItemInfoW", &GetMenuItemInfo},
	{"GetMenuStringW", &GetMenuString},
	{"GetMessageW", &GetMessage},
	{"GetMonitorInfoW", &GetMonitorInfo},
	{"GetPropW", &GetProp},
	{"GetRawInputDeviceInfoW", &GetRawInputDeviceInfo},
	{"GetTabbedTextExtentW", &GetTabbedTextExtent},
	{"GetUserObjectInformationW", &GetUserObjectInformation},
	{"GetWindowLongW", &GetWindowLong},
	{"GetWindowLongW", &GetWindowLongPtr}, //NOTE(t): Not on xp32
	{"GetWindowTextLengthW", &GetWindowTextLength},
	{"GetWindowTextW", &GetWindowText},
	{"GrayStringW", &GrayString},
	{"InsertMenuItemW", &InsertMenuItem},
	{"InsertMenuW", &InsertMenu},
	{"IsCharAlphaNumericW", &IsCharAlphaNumeric},
	{"IsCharAlphaW", &IsCharAlpha},
	{"IsCharLowerW", &IsCharLower},
	{"IsCharUpperW", &IsCharUpper},
	{"IsDialogMessageW", &IsDialogMessage},
	{"LoadAcceleratorsW", &LoadAccelerators},
	{"LoadBitmapW", &LoadBitmap},
	{"LoadCursorFromFileW", &LoadCursorFromFile},
	{"LoadCursorW", &LoadCursor},
	{"LoadIconW", &LoadIcon},
	{"LoadImageW", &LoadImage},
	{"LoadKeyboardLayoutW", &LoadKeyboardLayout},
	{"LoadMenuIndirectW", &LoadMenuIndirect},
	{"LoadMenuW", &LoadMenu},
	{"LoadStringW", &LoadString},
	{"MapVirtualKeyExW", &MapVirtualKeyEx},
	{"MapVirtualKeyW", &MapVirtualKey},
	{"MessageBoxExW", &MessageBoxEx},
	{"MessageBoxIndirectW", &MessageBoxIndirect},
	{"MessageBoxW", &MessageBox},
	{"ModifyMenuW", &ModifyMenu},
	{"OemToCharBuffW", &OemToCharBuff},
	{"OemToCharW", &OemToChar},
	{"OpenDesktopW", &OpenDesktop},
	{"OpenWindowStationW", &OpenWindowStation},
	{"PeekMessageW", &PeekMessage},
	{"PostMessageW", &PostMessage},
	{"PostThreadMessageW", &PostThreadMessage},
	{"PrivateExtractIconsW", &PrivateExtractIcons},
	{"RegisterClassExW", &RegisterClassEx},
	{"RegisterClassW", &RegisterClass},
	{"RegisterClipboardFormatW", &RegisterClipboardFormat},
	{"RegisterDeviceNotificationW", &RegisterDeviceNotification},
	{"RegisterWindowMessageW", &RegisterWindowMessage},
	{"RemovePropW", &RemoveProp},
	{"SendDlgItemMessageW", &SendDlgItemMessage},
	{"SendMessageCallbackW", &SendMessageCallback},
	{"SendMessageTimeoutW", &SendMessageTimeout},
	{"SendMessageW", &SendMessage},
	{"SendNotifyMessageW", &SendNotifyMessage},
	{"SetClassLongW", &SetClassLong},
	{"SetClassLongW", &SetClassLongPtr}, //NOTE(t): Not on xp32
	{"SetDlgItemTextW", &SetDlgItemText},
	{"SetMenuItemInfoW", &SetMenuItemInfo},
	{"SetPropW", &SetProp},
	{"SetUserObjectInformationW", &SetUserObjectInformation},
	{"SetWindowLongW", &SetWindowLong},
	{"SetWindowLongW", &SetWindowLongPtr}, //NOTE(t): Not on xp32
	{"SetWindowsHookExW", &SetWindowsHookEx},
	{"SetWindowsHookW", &SetWindowsHook},
	{"SetWindowTextW", &SetWindowText},
	{"SystemParametersInfoW", &SystemParametersInfo},
	{"TabbedTextOutW", &TabbedTextOut},
	{"UnregisterClassW", &UnregisterClass},
	{"VkKeyScanExW", &VkKeyScanEx},
	{"VkKeyScanW", &VkKeyScan},
	{"WinHelpW", &WinHelp},
	{"wsprintfW", &Wsprintf},
	{"wvsprintfW", &Wvsprintf},
}

var WinUserANSIApis = outside.Apis{
	{"AppendMenuA", &AppendMenu},
	{"BroadcastSystemMessageA", &BroadcastSystemMessage},
	{"BroadcastSystemMessageExA", &BroadcastSystemMessageEx},
	{"CallWindowProcA", &CallWindowProc},
	{"ChangeDisplaySettingsA", &ChangeDisplaySettingsA},
	{"ChangeDisplaySettingsExA", &ChangeDisplaySettingsExA},
	{"ChangeMenuA", &ChangeMenu},
	{"CharLowerA", &CharLower},
	{"CharLowerBuffA", &CharLowerBuff},
	{"CharNextA", &CharNext},
	{"CharNextExA", &CharNextExA},
	{"CharPrevA", &CharPrev},
	{"CharPrevExA", &CharPrevExA},
	{"CharToOemA", &CharToOem},
	{"CharToOemBuffA", &CharToOemBuff},
	{"CharUpperA", &CharUpper},
	{"CharUpperBuffA", &CharUpperBuff},
	{"CopyAcceleratorTableA", &CopyAcceleratorTable},
	{"CreateAcceleratorTableA", &CreateAcceleratorTable},
	{"CreateDesktopA", &CreateDesktopA},
	{"CreateDialogIndirectParamA", &CreateDialogIndirectParam},
	{"CreateDialogParamA", &CreateDialogParam},
	{"CreateMDIWindowA", &CreateMDIWindow},
	{"CreateWindowExA", &CreateWindowEx},
	{"CreateWindowStationA", &CreateWindowStation},
	{"DefDlgProcA", &DefDlgProc},
	{"DefFrameProcA", &DefFrameProc},
	{"DefMDIChildProcA", &DefMDIChildProc},
	{"DefWindowProcA", &DefWindowProc},
	{"DispatchMessageA", &DispatchMessage},
	{"DlgDirListA", &DlgDirList},
	{"DlgDirListComboBoxA", &DlgDirListComboBox},
	{"DlgDirSelectComboBoxExA", &DlgDirSelectComboBoxEx},
	{"DlgDirSelectExA", &DlgDirSelectEx},
	{"DrawStateA", &DrawState},
	{"DrawTextA", &DrawText},
	{"DrawTextExA", &DrawTextEx},
	{"EnumDesktopsA", &EnumDesktops},
	{"EnumDisplayDevicesA", &EnumDisplayDevicesA},
	{"EnumDisplaySettingsA", &EnumDisplaySettingsA},
	{"EnumDisplaySettingsExA", &EnumDisplaySettingsExA},
	{"EnumPropsA", &EnumPropsA},
	{"EnumPropsExA", &EnumPropsExA},
	{"EnumWindowStationsA", &EnumWindowStations},
	{"FindWindowA", &FindWindow},
	{"FindWindowExA", &FindWindowEx},
	{"GetClassInfoA", &GetClassInfo},
	{"GetClassInfoExA", &GetClassInfoEx},
	{"GetClassLongA", &GetClassLong},
	{"GetClassLongA", &GetClassLongPtr}, //NOTE(t): Not on xp32
	{"GetClassNameA", &GetClassName},
	{"GetClipboardFormatNameA", &GetClipboardFormatName},
	{"GetDlgItemTextA", &GetDlgItemText},
	{"GetKeyboardLayoutNameA", &GetKeyboardLayoutName},
	{"GetKeyNameTextA", &GetKeyNameText},
	{"GetMenuItemInfoA", &GetMenuItemInfo},
	{"GetMenuStringA", &GetMenuString},
	{"GetMessageA", &GetMessage},
	{"GetMonitorInfoA", &GetMonitorInfo},
	{"GetPropA", &GetProp},
	{"GetRawInputDeviceInfoA", &GetRawInputDeviceInfo},
	{"GetTabbedTextExtentA", &GetTabbedTextExtent},
	{"GetUserObjectInformationA", &GetUserObjectInformation},
	{"GetWindowLongA", &GetWindowLong},
	{"GetWindowLongA", &GetWindowLongPtr}, //NOTE(t): Not on xp32
	{"GetWindowTextA", &GetWindowText},
	{"GetWindowTextLengthA", &GetWindowTextLength},
	{"GrayStringA", &GrayString},
	{"InsertMenuA", &InsertMenu},
	{"InsertMenuItemA", &InsertMenuItem},
	{"IsCharAlphaA", &IsCharAlpha},
	{"IsCharAlphaNumericA", &IsCharAlphaNumeric},
	{"IsCharLowerA", &IsCharLower},
	{"IsCharUpperA", &IsCharUpper},
	{"IsDialogMessageA", &IsDialogMessage},
	{"LoadAcceleratorsA", &LoadAccelerators},
	{"LoadBitmapA", &LoadBitmap},
	{"LoadCursorA", &LoadCursor},
	{"LoadCursorFromFileA", &LoadCursorFromFile},
	{"LoadIconA", &LoadIcon},
	{"LoadImageA", &LoadImage},
	{"LoadKeyboardLayoutA", &LoadKeyboardLayout},
	{"LoadMenuA", &LoadMenu},
	{"LoadMenuIndirectA", &LoadMenuIndirect},
	{"LoadStringA", &LoadString},
	{"MapVirtualKeyA", &MapVirtualKey},
	{"MapVirtualKeyExA", &MapVirtualKeyEx},
	{"MessageBoxA", &MessageBox},
	{"MessageBoxExA", &MessageBoxEx},
	{"MessageBoxIndirectA", &MessageBoxIndirect},
	{"ModifyMenuA", &ModifyMenu},
	{"OemToCharA", &OemToChar},
	{"OemToCharBuffA", &OemToCharBuff},
	{"OpenDesktopA", &OpenDesktop},
	{"OpenWindowStationA", &OpenWindowStation},
	{"PeekMessageA", &PeekMessage},
	{"PostMessageA", &PostMessage},
	{"PostThreadMessageA", &PostThreadMessage},
	{"PrivateExtractIconsA", &PrivateExtractIcons},
	{"RegisterClassA", &RegisterClass},
	{"RegisterClassExA", &RegisterClassEx},
	{"RegisterClipboardFormatA", &RegisterClipboardFormat},
	{"RegisterDeviceNotificationA", &RegisterDeviceNotification},
	{"RegisterWindowMessageA", &RegisterWindowMessage},
	{"RemovePropA", &RemoveProp},
	{"SendDlgItemMessageA", &SendDlgItemMessage},
	{"SendMessageA", &SendMessage},
	{"SendMessageCallbackA", &SendMessageCallback},
	{"SendMessageTimeoutA", &SendMessageTimeout},
	{"SendNotifyMessageA", &SendNotifyMessage},
	{"SetClassLongA", &SetClassLong},
	{"SetClassLongA", &SetClassLongPtr}, //NOTE(t): Not on xp32
	{"SetDlgItemTextA", &SetDlgItemText},
	{"SetMenuItemInfoA", &SetMenuItemInfo},
	{"SetPropA", &SetProp},
	{"SetUserObjectInformationA", &SetUserObjectInformation},
	{"SetWindowLongA", &SetWindowLong},
	{"SetWindowLongA", &SetWindowLongPtr}, //NOTE(t): Not on xp32
	{"SetWindowsHookA", &SetWindowsHook},
	{"SetWindowsHookExA", &SetWindowsHookEx},
	{"SetWindowTextA", &SetWindowText},
	{"SystemParametersInfoA", &SystemParametersInfo},
	{"TabbedTextOutA", &TabbedTextOut},
	{"UnregisterClassA", &UnregisterClass},
	{"VkKeyScanA", &VkKeyScan},
	{"VkKeyScanExA", &VkKeyScanEx},
	{"WinHelpA", &WinHelp},
	{"wsprintfA", &Wsprintf},
	{"wvsprintfA", &Wvsprintf},
}
