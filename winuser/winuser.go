// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package winuser provides API definitions for accessing
//user32.dll.
package winuser

import (
	. "github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/user32"
)

func CopyCursor(cur T.HCURSOR) T.HCURSOR {
	return T.HCURSOR(CopyIcon(T.HICON(cur)))
}

func CreateDialog(instance T.HINSTANCE, templateName VString, wndParent T.HWND, dialogFunc T.DLGPROC) T.HWND {
	return CreateDialogParam(instance, templateName, wndParent, dialogFunc, 0)
}

func CreateDialogIndirect(instance T.HINSTANCE, template T.DLGTEMPLATE, wndParent T.HWND,dialogFunc T.DLGPROC) T.HWND {
	return CreateDialogIndirectParam(instance, template, wndParent, dialogFunc, 0)
}

func CreateWindow(className, windowName VString, style T.WINDOW_STYLE, x, y, width, height int, wndParent T.HWND, menu T.HMENU, instance T.HINSTANCE, param *T.VOID) T.HWND {
	return CreateWindowEx(0, className, windowName, style, x, y, width, height, wndParent, menu, instance, param)
}

func DialogBox(instance T.HINSTANCE, template VString, wndParent T.HWND, dialogFunc T.DLGPROC) T.INT_PTR {
	return DialogBoxParam(instance, template, wndParent, dialogFunc, 0)
}

func DialogBoxIndirect(instance T.HINSTANCE, template T.DLGTEMPLATE, wndParent T.HWND, dialogFunc T.DLGPROC) T.INT_PTR {
	return DialogBoxIndirectParam(instance, template, wndParent, dialogFunc, 0)
}

func EnumTaskWindows(task T.DWORD, fn T.WNDENUMPROC, param T.LPARAM) T.BOOL {
	return EnumThreadWindows(task, fn, param)
}

func GetNextWindow(Wnd T.HWND, cmd T.UINT) T.HWND {
	return GetWindow(Wnd, cmd)
}

func GetSysModalWindow() {}

func GetWindowTask(Wnd T.HWND) T.HANDLE {
	return T.HANDLE(GetWindowThreadProcessId(Wnd, nil))
}

func SetSysModalWindow(Wnd T.HWND) {}

type (
	va_list  T.Fake_type_Fix_me
	variadic T.Fake_type_Fix_me
)

var (
	ActivateKeyboardLayout func(hkl T.HKL, flags T.UINT) T.HKL

	AdjustWindowRect func(
		rect *T.RECT, style T.DWORD, menu T.BOOL) T.BOOL

	AdjustWindowRectEx func(
		rect *T.RECT, style T.DWORD, menu T.BOOL, exStyle T.DWORD) T.BOOL

	AllowSetForegroundWindow func(processId T.DWORD) T.BOOL

	AnimateWindow func(w T.HWND, time T.DWORD, flags T.DWORD) T.BOOL

	AnyPopup func() T.BOOL

	AppendMenu func(
		m T.HMENU,
		flags T.UINT,
		idNewItem T.UINT_PTR,
		newItem VString) T.BOOL

	ArrangeIconicWindows func(w T.HWND) T.UINT

	AttachThreadInput func(attach, to T.DWORD, flag T.BOOL) T.BOOL

	BeginDeferWindowPos func(numWindows int) T.HDWP

	BeginPaint func(w T.HWND, paint *T.PAINTSTRUCT) T.HDC

	BringWindowToTop func(w T.HWND) T.BOOL

	BroadcastSystemMessage func(
		flags T.DWORD,
		info *T.DWORD,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM) T.LONG

	BroadcastSystemMessageEx func(
		flags T.BSM_FLAGS,
		recipients *T.DWORD,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM,
		info *T.BSMINFO) T.LONG

	CallMsgFilter func(msg *T.MSG, code int) T.BOOL

	CallNextHookEx func(
		hk T.HHOOK, code int, w T.WPARAM, l T.LPARAM) T.LRESULT

	CallWindowProc func(
		prevWndFunc T.WNDPROC,
		win T.HWND,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM) T.LRESULT

	CascadeWindows func(
		wndParent T.HWND,
		how T.UINT,
		rect *T.RECT,
		nKids T.UINT,
		kids *T.HWND) T.WORD

	ChangeClipboardChain func(wndRemove, wndNewNext T.HWND) T.BOOL

	ChangeDisplaySettingsA func(dm *T.DEVMODEA, flags T.DWORD) T.LONG

	ChangeDisplaySettingsW func(dm *T.DEVMODEW, flags T.DWORD) T.LONG

	ChangeDisplaySettingsExA func(
		deviceName VString,
		dm *T.DEVMODEA,
		w T.HWND,
		flags T.DWORD,
		lParam *T.VOID) T.LONG

	ChangeDisplaySettingsExW func(
		deviceName VString,
		dm *T.DEVMODEW,
		w T.HWND,
		flags T.DWORD,
		lParam *T.VOID) T.LONG

	ChangeMenu func(
		m T.HMENU,
		cmd T.UINT,
		NewItem VString,
		insert T.UINT,
		flags T.UINT) T.BOOL

	CharLower func(VString) OVString

	CharLowerBuff func(buff T.IOAString, length T.DWORD) T.DWORD

	CharNext func(VString) OVString

	CharNextExA func(
		codePage T.WORD, currentChar T.AString, flags T.DWORD) T.AString

	CharPrev func(start, current VString) VString

	CharPrevExA func(
		codePage T.WORD,
		start, currentChar T.AString,
		flags T.DWORD) T.AString

	CharToOem func(src VString, dst OVString) T.BOOL

	CharToOemBuff func(src, dst OVString, dstLength T.DWORD) T.BOOL

	CharUpper func(VString) OVString

	CharUpperBuff func(buff T.IOAString, length T.DWORD) T.DWORD

	CheckDlgButton func(dlg T.HWND, idButton int, check T.UINT) T.BOOL

	CheckMenuItem func(m T.HMENU, idCheckItem, check T.UINT) T.DWORD

	CheckMenuRadioItem func(
		m T.HMENU, first, last, check, flags T.UINT) T.BOOL

	CheckRadioButton func(
		dlg T.HWND, idFirst, idLast, idCheck int) T.BOOL

	ChildWindowFromPoint func(wndParent T.HWND, point T.POINT) T.HWND

	ChildWindowFromPointEx func(
		w T.HWND, p T.POINT, flags T.UINT) T.HWND

	ClientToScreen func(w T.HWND, p *T.POINT) T.BOOL

	ClipCursor func(rect *T.RECT) T.BOOL

	CloseClipboard func() T.BOOL

	CloseDesktop func(desktop T.HDESK) T.BOOL

	CloseWindow func(w T.HWND) T.BOOL

	CloseWindowStation func(winSta T.HWINSTA) T.BOOL

	CopyAcceleratorTable func(
		src T.HACCEL, dst *T.ACCEL, entries int) int

	CopyIcon func(icon T.HICON) T.HICON

	CopyImage func(
		h T.HANDLE, typ T.UINT, T.X, y int, flags T.UINT) T.HANDLE

	CopyRect func(dst, src *T.RECT) T.BOOL

	CountClipboardFormats func() int

	CreateAcceleratorTable func(accel *T.ACCEL, nAccel int) T.HACCEL

	CreateCaret func(
		w T.HWND, bitmap T.HBITMAP, width int, height int) T.BOOL

	CreateCursor func(
		inst T.HINSTANCE,
		xHotSpot, yHotSpot, width, height int,
		andPlane, xorPlane *T.VOID) T.HCURSOR

	CreateDesktopA func(
		desktop, device VString,
		devmode *T.DEVMODEA,
		flags T.DWORD,
		desiredAccess T.ACCESS_MASK,
		sa *T.SECURITY_ATTRIBUTES) T.HDESK

	CreateDesktopW func(
		desktop, device VString,
		devmode *T.DEVMODEW,
		flags T.DWORD,
		desiredAccess T.ACCESS_MASK,
		sa *T.SECURITY_ATTRIBUTES) T.HDESK

	CreateDialogIndirectParam func(
		instance T.HINSTANCE,
		template T.DLGTEMPLATE,
		wndParent T.HWND,
		dialogFunc T.DLGPROC,
		initParam T.LPARAM) T.HWND

	CreateDialogParam func(
		instance T.HINSTANCE,
		templateName VString,
		wndParent T.HWND,
		dialogFunc T.DLGPROC,
		initParam T.LPARAM) T.HWND

	CreateIcon func(
		instance T.HINSTANCE,
		width, height int,
		planes, bitsPixel T.BYTE,
		andbits, xorbits *T.BYTE) T.HICON

	CreateIconFromResource func(
		resbits *T.BYTE,
		resSize T.DWORD,
		icon T.BOOL,
		ver T.DWORD) T.HICON

	CreateIconFromResourceEx func(
		resbits *T.BYTE,
		resSize T.DWORD,
		icon T.BOOL,
		ver T.DWORD,
		xDesired, yDesired int,
		flags T.UINT) T.HICON

	CreateIconIndirect func(iconinfo *T.ICONINFO) T.HICON

	CreateMDIWindow func(
		className, windowName VString,
		style T.DWORD,
		x, y, width,
		height int,
		wndParent T.HWND,
		instance T.HINSTANCE,
		l T.LPARAM) T.HWND

	CreateMenu func() T.HMENU

	CreatePopupMenu func() T.HMENU

	CreateWindowEx func(
		exStyle T.WINDOW_STYLE_EX,
		className, windowName VString,
		style T.WINDOW_STYLE,
		x, y, width, height int,
		wndParent T.HWND,
		m T.HMENU,
		instance T.HINSTANCE,
		param *T.VOID) T.HWND

	CreateWindowStation func(
		winsta VString,
		flags T.DWORD,
		desiredAccess T.ACCESS_MASK,
		sa *T.SECURITY_ATTRIBUTES) T.HWINSTA

	DefDlgProc func(
		dlg T.HWND, msg T.UINT, w T.WPARAM, l T.LPARAM) T.LRESULT

	DeferWindowPos func(
		winPosInfo T.HDWP,
		w T.HWND,
		wndInsertAfter T.HWND,
		x, y, cx, cy int,
		flags T.UINT) T.HDWP

	DefFrameProc func(
		win T.HWND,
		mdiClient T.HWND,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM) T.LRESULT

	DefMDIChildProc func(
		win T.HWND, msg T.UINT, w T.WPARAM, l T.LPARAM) T.LRESULT

	DefRawInputProc func(rawInput *T.RAWINPUT, input T.INT, sizeHeader T.INT) T.LRESULT

	DefWindowProc func(h T.HWND, m T.WINDOW_MESSAGE, w T.WPARAM, l T.LPARAM) T.LRESULT

	DeleteMenu func(m T.HMENU, position T.UINT, flags T.UINT) T.BOOL

	DeregisterShellHookWindow func(w T.HWND) T.BOOL

	DestroyAcceleratorTable func(accel T.HACCEL) T.BOOL

	DestroyCaret func() T.BOOL

	DestroyCursor func(cursor T.HCURSOR) T.BOOL

	DestroyIcon func(icon T.HICON) T.BOOL

	DestroyMenu func(m T.HMENU) T.BOOL

	DestroyWindow func(w T.HWND) T.BOOL

	DialogBoxIndirectParam func(
		instance T.HINSTANCE,
		dialogTemplate T.DLGTEMPLATE,
		wndParent T.HWND,
		dialogFunc T.DLGPROC,
		initParam T.LPARAM) T.INT_PTR

	DialogBoxParam func(
		instance T.HINSTANCE,
		templateName VString,
		wndParent T.HWND,
		dialogFunc T.DLGPROC,
		initParam T.LPARAM) T.INT_PTR

	DisableProcessWindowsGhosting func()

	DispatchMessage func(msg *T.MSG) T.LRESULT

	DlgDirList func(
		dlg T.HWND,
		pathSpec VString,
		idListBox, idStaticPath int,
		fileType T.UINT) int

	DlgDirListComboBox func(
		dlg T.HWND,
		pathSpec VString,
		idComboBox, idStaticPath int,
		filetype T.UINT) int

	DlgDirSelectComboBoxEx func(
		dlg T.HWND,
		string_ VString,
		out int,
		comboBox int) T.BOOL

	DlgDirSelectEx func(
		dlg T.HWND,
		string_ VString,
		count int,
		listBox int) T.BOOL

	DragDetect func(w T.HWND, pt T.POINT) T.BOOL

	DragObject func(
		parent T.HWND,
		from T.HWND,
		fmt T.UINT,
		data T.ULONG_PTR,
		cur T.HCURSOR) T.DWORD

	DrawAnimatedRects func(
		hw T.HWND, ani int, from *T.RECT, to *T.RECT) T.BOOL

	DrawCaption func(
		w T.HWND, dc T.HDC, rect *T.RECT, flags T.UINT) T.BOOL

	DrawEdge func(
		hdc T.HDC, qrc *T.RECT, edge T.UINT, grfFlags T.UINT) T.BOOL

	DrawFocusRect func(dc T.HDC, rc *T.RECT) T.BOOL

	DrawFrameControl func(T.HDC, *T.RECT, T.UINT, T.UINT) T.BOOL

	DrawIcon func(dc T.HDC, x, y int, icon T.HICON) T.BOOL

	DrawIconEx func(
		dc T.HDC,
		xLeft, yTop int,
		icon T.HICON,
		xWidth, yWidth int,
		stepIfAniCur T.UINT,
		brFlickerFreeDraw T.HBRUSH,
		flags T.UINT) T.BOOL

	DrawMenuBar func(w T.HWND) T.BOOL

	DrawState func(
		dc T.HDC,
		brFore T.HBRUSH,
		fnCallBack T.DRAWSTATEPROC,
		lData T.LPARAM,
		wData T.WPARAM,
		x, y, cx, cy int,
		uFlags T.UINT) T.BOOL

	DrawText func(
		dc T.HDC,
		text VString,
		nText int,
		rc *T.RECT,
		format T.UINT) int

	DrawTextEx func(
		dc T.HDC,
		text VString,
		nText int,
		rc *T.RECT,
		format T.UINT,
		dtp *T.DRAWTEXTPARAMS) int

	EmptyClipboard func() T.BOOL

	EnableMenuItem func(
		m T.HMENU, idDEnableItem T.UINT, enable T.UINT) T.BOOL

	EnableScrollBar func(w T.HWND, sbflags T.UINT, arrows T.UINT) T.BOOL

	EnableWindow func(w T.HWND, enable T.BOOL) T.BOOL

	EndDeferWindowPos func(winPosInfo T.HDWP) T.BOOL

	EndDialog func(dlg T.HWND, result T.INT_PTR) T.BOOL

	EndMenu func() T.BOOL

	EndPaint func(w T.HWND, paint *T.PAINTSTRUCT) T.BOOL

	EndTask func(w T.HWND, shutDown T.BOOL, force T.BOOL) T.BOOL

	EnumChildWindows func(
		parent T.HWND, fn T.WNDENUMPROC, l T.LPARAM) T.BOOL

	EnumClipboardFormats func(format T.UINT) T.UINT

	EnumDesktops func(
		wsta T.HWINSTA, fn T.DESKTOPENUMPROC, l T.LPARAM) T.BOOL

	EnumDesktopWindows func(
		desktop T.HDESK, fn T.WNDENUMPROC, param T.LPARAM) T.BOOL

	EnumDisplayDevicesA func(
		device VString,
		devNum T.DWORD,
		displayDevice *T.DISPLAY_DEVICEA,
		flags T.DWORD) T.BOOL

	EnumDisplayDevicesW func(
		device VString,
		devNum T.DWORD,
		displayDevice *T.DISPLAY_DEVICEW,
		flags T.DWORD) T.BOOL

	EnumDisplayMonitors func(
		dc T.HDC,
		clip *T.RECT,
		fn T.MONITORENUMPROC,
		data T.LPARAM) T.BOOL

	EnumDisplaySettingsA func(
		deviceName VString,
		modeNum T.DWORD,
		dm *T.DEVMODEA) T.BOOL

	EnumDisplaySettingsW func(
		deviceName VString,
		modeNum T.DWORD,
		dm *T.DEVMODEW) T.BOOL

	EnumDisplaySettingsExA func(
		deviceName VString,
		modeNum T.DWORD,
		dm *T.DEVMODEA,
		flags T.DWORD) T.BOOL

	EnumDisplaySettingsExW func(
		deviceName VString,
		modeNum T.DWORD,
		dm *T.DEVMODEW,
		flags T.DWORD) T.BOOL

	EnumPropsA func(w T.HWND, enumFunc T.PROPENUMPROCA) int

	EnumPropsW func(w T.HWND, enumFunc T.PROPENUMPROCW) int

	EnumPropsExA func(
		w T.HWND, enumFunc T.PROPENUMPROCEXA, l T.LPARAM) int

	EnumPropsExW func(
		w T.HWND, enumFunc T.PROPENUMPROCEXW, l T.LPARAM) int

	EnumThreadWindows func(
		threadId T.DWORD, fn T.WNDENUMPROC, l T.LPARAM) T.BOOL

	EnumWindows func(enumFunc T.WNDENUMPROC, l T.LPARAM) T.BOOL

	EnumWindowStations func(
		enumFunc T.WINSTAENUMPROC, l T.LPARAM) T.BOOL

	EqualRect func(rc1 *T.RECT, rc2 *T.RECT) T.BOOL

	ExcludeUpdateRgn func(dc T.HDC, w T.HWND) int

	ExitWindowsEx func(flags T.UINT, reason T.DWORD) T.BOOL

	FillRect func(dc T.HDC, rc *T.RECT, br T.HBRUSH) int

	FindWindow func(className VString, windowName VString) T.HWND

	FindWindowEx func(
		wndParent, wndChildAfter T.HWND,
		class, window VString) T.HWND

	FlashWindow func(w T.HWND, invert T.BOOL) T.BOOL

	FlashWindowEx func(fwi *T.FLASHWINFO) T.BOOL

	FrameRect func(dc T.HDC, rc *T.RECT, br T.HBRUSH) int

	GetActiveWindow func() T.HWND

	GetAltTabInfo func(
		w T.HWND,
		item int,
		ati *T.ALTTABINFO,
		itemText OVString,
		size T.UINT) T.BOOL

	GetAncestor func(w T.HWND, flags T.UINT) T.HWND

	GetAsyncKeyState func(key int) T.SHORT

	GetCapture func() T.HWND

	GetCaretBlinkTime func() T.UINT

	GetCaretPos func(point *T.POINT) T.BOOL

	GetClassInfo func(
		instance T.HINSTANCE, className VString, wc *T.WNDCLASS) T.BOOL

	GetClassInfoEx func(
		instance T.HINSTANCE, class VString, wcx *T.WNDCLASSEX) T.BOOL

	GetClassLong func(w T.HWND, index int) T.DWORD

	GetClassLongPtr func(w T.HWND, index int) T.ULONG_PTR

	//GetClassName func(
	//	w T.HWND, className OVString, maxCount int) int
	GetClassName func(
		w T.HWND, className *T.WChar, maxCount int) int

	//GetClassWord is obsolete; instead use:
	//	GetClassLong
	GetClassWord func(w T.HWND, index int) T.WORD

	GetClientRect func(w T.HWND, rect *T.RECT) T.BOOL

	GetClipboardData func(format T.UINT) T.HANDLE

	GetClipboardFormatName func(
		format T.UINT, formatName VString, maxCount int) int

	GetClipboardOwner func() T.HWND

	GetClipboardSequenceNumber func() T.DWORD

	GetClipboardViewer func() T.HWND

	GetClipCursor func(rect *T.RECT) T.BOOL

	GetComboBoxInfo func(w T.HWND, cbi *T.COMBOBOXINFO) T.BOOL

	GetCursor func() T.HCURSOR

	GetCursorInfo func(ci *T.CURSORINFO) T.BOOL

	GetCursorPos func(point *T.POINT) T.BOOL

	GetDC func(w T.HWND) T.HDC

	GetDCEx func(w T.HWND, rgnClip T.HRGN, flags T.DWORD) T.HDC

	GetDesktopWindow func() T.HWND

	GetDialogBaseUnits func() T.LONG

	GetDlgCtrlID func(w T.HWND) int

	GetDlgItem func(dlg T.HWND, idDlgItem int) T.HWND

	GetDlgItemInt func(
		dlg T.HWND,
		idDlgItem int,
		translated *T.BOOL,
		signed T.BOOL) T.UINT

	GetDlgItemText func(
		dlg T.HWND,
		idDlgItem int,
		string_ OVString,
		max int) T.UINT

	GetDoubleClickTime func() T.UINT

	GetFocus func() T.HWND

	GetForegroundWindow func() T.HWND

	GetGuiResources func(process T.HANDLE, flags T.DWORD) T.DWORD

	GetGUIThreadInfo func(thread T.DWORD, ui *T.GUITHREADINFO) T.BOOL

	GetIconInfo func(icon T.HICON, iconinfo *T.ICONINFO) T.BOOL

	GetInputState func() T.BOOL

	GetKBCodePage func() T.UINT

	GetKeyboardLayout func(thread T.DWORD) T.HKL

	GetKeyboardLayoutList func(buff int, list *T.HKL) int

	GetKeyboardLayoutName func(klid VString) T.BOOL

	GetKeyboardState func(keyState *T.BYTE) T.BOOL

	GetKeyboardType func(typeFlag int) int

	GetKeyNameText func(
		param T.LONG, string_ OVString, size int) int

	GetKeyState func(virtKey int) T.SHORT

	GetLastActivePopup func(w T.HWND) T.HWND

	GetLastInputInfo func(lii *T.LASTINPUTINFO) T.BOOL

	GetLayeredWindowAttributes func(
		w T.HWND, crKey *T.COLORREF, alpha *T.BYTE, flags *T.DWORD) T.BOOL

	GetListBoxInfo func(w T.HWND) T.DWORD

	GetMenu func(w T.HWND) T.HMENU

	GetMenuBarInfo func(
		w T.HWND, object, item T.LONG, mbi *T.MENUBARINFO) T.BOOL

	GetMenuCheckMarkDimensions func() T.LONG

	GetMenuContextHelpId func(T.HMENU) T.DWORD

	GetMenuDefaultItem func(
		m T.HMENU, byPos T.UINT, gmdiFlags T.UINT) T.UINT

	GetMenuInfo func(T.HMENU, *T.MENUINFO) T.BOOL

	GetMenuItemCount func(m T.HMENU) int

	GetMenuItemID func(m T.HMENU, pos int) T.UINT

	GetMenuItemInfo func(
		m T.HMENU,
		item T.UINT,
		byPosition T.BOOL,
		mii *T.MENUITEMINFO) T.BOOL

	GetMenuItemRect func(
		w T.HWND, m T.HMENU, item T.UINT, rcItem *T.RECT) T.BOOL

	GetMenuState func(m T.HMENU, id T.UINT, flags T.UINT) T.UINT

	GetMenuString func(
		m T.HMENU,
		idItem T.UINT,
		string_ OVString,
		max int,
		flags T.UINT) int

	GetMessage func(
		m *T.MSG, w T.HWND, filterMin, filterMax T.UINT) T.BOOL

	GetMessageExtraInfo func() T.LPARAM

	GetMessagePos func() T.DWORD

	GetMessageTime func() T.LONG

	GetMonitorInfo func(monitor T.HMONITOR, mi *T.MONITORINFO) T.BOOL

	GetMouseMovePointsEx func(
		size T.UINT,
		pt, ptBuf *T.MOUSEMOVEPOINT,
		bufPoints int,
		resolution T.DWORD) int

	GetNextDlgGroupItem func(
		dlg, ctl T.HWND, previous T.BOOL) T.HWND

	GetNextDlgTabItem func(
		dlg, ctl T.HWND, previous T.BOOL) T.HWND

	GetOpenClipboardWindow func() T.HWND

	GetParent func(w T.HWND) T.HWND

	GetPriorityClipboardFormat func(
		formatPriorityList *T.UINT, formats int) int

	GetProcessDefaultLayout func(defaultLayout *T.DWORD) T.BOOL

	GetProcessWindowStation func() T.HWINSTA

	GetProp func(w T.HWND, string_ VString) T.HANDLE

	GetQueueStatus func(flags T.UINT) T.DWORD

	GetRawInputBuffer func(
		data *T.RAWINPUT, size *T.UINT, sizeHeader T.UINT) T.UINT

	GetRawInputData func(
		rawInput T.HRAWINPUT,
		command T.UINT,
		data *T.VOID,
		size *T.UINT,
		sizeHeader T.UINT) T.UINT

	GetRawInputDeviceInfo func(
		device T.HANDLE,
		command T.UINT,
		data *T.VOID,
		size *T.UINT) T.UINT

	GetRawInputDeviceList func(
		ridl *T.RAWINPUTDEVICELIST,
		numDevices *T.UINT,
		size T.UINT) T.UINT

	GetRegisteredRawInputDevices func(
		rid *T.RAWINPUTDEVICE, numDevices *T.UINT, size T.UINT) T.UINT

	GetScrollBarInfo func(
		w T.HWND, object T.LONG, sbi *T.SCROLLBARINFO) T.BOOL

	GetScrollInfo func(w T.HWND, bar int, si *T.SCROLLINFO) T.BOOL

	GetScrollPos func(w T.HWND, bar int) int

	GetScrollRange func(
		w T.HWND, bar int, minPos *T.INT, maxPos *T.INT) T.BOOL

	GetShellWindow func() T.HWND

	GetSubMenu func(m T.HMENU, pos int) T.HMENU

	GetSysColor func(index int) T.DWORD

	GetSysColorBrush func(index int) T.HBRUSH

	GetSystemMenu func(w T.HWND, revert T.BOOL) T.HMENU

	GetSystemMetrics func(index int) int

	GetTabbedTextExtent func(
		dc T.HDC,
		s VString,
		count, tabPositions int,
		tabStopPositions *T.INT) T.DWORD

	GetThreadDesktop func(threadId T.DWORD) T.HDESK

	GetTitleBarInfo func(w T.HWND, ti *T.TITLEBARINFO) T.BOOL

	GetTopWindow func(w T.HWND) T.HWND

	GetUpdateRect func(w T.HWND, rect *T.RECT, erase T.BOOL) T.BOOL

	GetUpdateRgn func(w T.HWND, rgn T.HRGN, erase T.BOOL) int

	GetUserObjectInformation func(
		obj T.HANDLE,
		index int,
		info *T.VOID,
		length T.DWORD,
		lengthNeeded *T.DWORD) T.BOOL

	GetUserObjectSecurity func(
		obj T.HANDLE,
		siRequested *T.SECURITY_INFORMATION,
		sd *T.SECURITY_DESCRIPTOR,
		length T.DWORD,
		lengthNeeded *T.DWORD) T.BOOL

	GetWindow func(w T.HWND, cmd T.UINT) T.HWND

	GetWindowContextHelpId func(T.HWND) T.DWORD

	GetWindowDC func(w T.HWND) T.HDC

	GetWindowInfo func(w T.HWND, wi *T.WINDOWINFO) T.BOOL

	GetWindowLong func(w T.HWND, index int) T.LONG

	GetWindowLongPtr func(w T.HWND, index int) T.LONG_PTR

	GetWindowModuleFileName func(
		w T.HWND, fileName OVString, fileNameMax T.UINT) T.UINT

	GetWindowPlacement func(w T.HWND, wndpl *T.WINDOWPLACEMENT) T.BOOL

	GetWindowRect func(w T.HWND, rect *T.RECT) T.BOOL

	GetWindowRgn func(w T.HWND, rgn T.HRGN) int

	GetWindowRgnBox func(w T.HWND, rc *T.RECT) int

	// GetWindowText func(w T.HWND, s OVString, maxCount int) int
	GetWindowText func(w T.HWND, s *T.WChar, maxCount int) int

	GetWindowTextLength func(w T.HWND) int

	GetWindowThreadProcessId func(
		w T.HWND, processId *T.DWORD) T.DWORD

	GetWindowWord func(w T.HWND, index int) T.WORD

	GrayString func(
		dc T.HDC,
		brush T.HBRUSH,
		outputFunc T.GRAYSTRINGPROC,
		data T.LPARAM,
		count, x, y, width, height int) T.BOOL

	HideCaret func(w T.HWND) T.BOOL

	HiliteMenuItem func(
		w T.HWND, m T.HMENU, idHiliteItem T.UINT, hilite T.UINT) T.BOOL

	InflateRect func(rc *T.RECT, x, y int) T.BOOL

	InSendMessage func() T.BOOL

	InSendMessageEx func(reserved *T.VOID) T.DWORD

	InsertMenu func(
		m T.HMENU,
		position T.UINT,
		flags T.UINT,
		idNewItem T.UINT_PTR,
		NewItem T.AString) T.BOOL

	InsertMenuItem func(
		m T.HMENU,
		item T.UINT,
		byPosition T.BOOL,
		mii T.MENUITEMINFO) T.BOOL

	InternalGetWindowText func(w T.HWND, s T.OWString, size int) int

	IntersectRect func(dst, src1, src2 *T.RECT) T.BOOL

	InvalidateRect func(w T.HWND, rect *T.RECT, erase T.BOOL) T.BOOL

	InvalidateRgn func(w T.HWND, rgn T.HRGN, erase T.BOOL) T.BOOL

	InvertRect func(dc T.HDC, rc *T.RECT) T.BOOL

	IsCharAlpha func(ch T.VChar) T.BOOL

	IsCharAlphaNumeric func(ch T.VChar) T.BOOL

	IsCharLower func(ch T.VChar) T.BOOL

	IsCharUpper func(ch T.VChar) T.BOOL

	IsChild func(wndParent T.HWND, w T.HWND) T.BOOL

	IsClipboardFormatAvailable func(format T.UINT) T.BOOL

	IsDialogMessage func(dlg T.HWND, msg *T.MSG) T.BOOL

	IsDlgButtonChecked func(dlg T.HWND, idButton int) T.UINT

	IsGUIThread func(convert T.BOOL) T.BOOL

	IsHungAppWindow func(w T.HWND) T.BOOL

	IsIconic func(w T.HWND) T.BOOL

	IsMenu func(m T.HMENU) T.BOOL

	IsRectEmpty func(rc *T.RECT) T.BOOL

	IsWindow func(w T.HWND) T.BOOL

	IsWindowEnabled func(w T.HWND) T.BOOL

	IsWindowUnicode func(w T.HWND) T.BOOL

	IsWindowVisible func(w T.HWND) T.BOOL

	IsWinEventHookInstalled func(event T.DWORD) T.BOOL

	IsWow64Message func() T.BOOL

	IsZoomed func(w T.HWND) T.BOOL

	Keybd_event func(
		vk T.BYTE, scan T.BYTE, flags T.DWORD, extraInfo T.ULONG_PTR)

	KillTimer func(w T.HWND, idEvent T.UINT_PTR) T.BOOL

	LoadAccelerators func(i T.HINSTANCE, tableName VString) T.HACCEL

	LoadBitmap func(i T.HINSTANCE, bitmapName VString) T.HBITMAP

	LoadCursor func(i T.HINSTANCE, cursorName PVString) T.HCURSOR

	LoadCursorFromFile func(fileName VString) T.HCURSOR

	LoadIcon func(i T.HINSTANCE, iconName VString) T.HICON

	LoadImage func(
		i T.HINSTANCE,
		name VString,
		type_ T.UINT,
		x, y int,
		load T.UINT) T.HANDLE

	LoadKeyboardLayout func(klid VString, flags T.UINT) T.HKL

	LoadMenu func(instance T.HINSTANCE, menuName VString) T.HMENU

	//TODO(t):How does this work (*T.VOID underneath)
	LoadMenuIndirect func(mt *T.MENUTEMPLATE) T.HMENU

	LoadString func(
		i T.HINSTANCE, id T.UINT, s VString, size int) int

	LockSetForegroundWindow func(lockCode T.UINT) T.BOOL

	LockWindowUpdate func(wndLock T.HWND) T.BOOL

	LockWorkStation func() T.BOOL

	LookupIconIdFromDirectory func(
		resbits *T.BYTE, icon T.BOOL) int

	LookupIconIdFromDirectoryEx func(
		resbits *T.BYTE,
		icon T.BOOL,
		xDesired, yDesired int,
		flags T.UINT) int

	MapDialogRect func(dlg T.HWND, rect *T.RECT) T.BOOL

	MapVirtualKey func(code T.UINT, mapType T.UINT) T.UINT

	MapVirtualKeyEx func(code T.UINT, mapType T.UINT, hkl T.HKL) T.UINT

	MapWindowPoints func(
		from T.HWND, to T.HWND, p *T.POINT, nPoints T.UINT) int

	MenuItemFromPoint func(w T.HWND, m T.HMENU, ptScreen T.POINT) int

	MessageBeep func(type_ T.UINT) T.BOOL

	MessageBox func(
		w T.HWND, text, caption VString, t T.MSGBOX_TYPE) int

	MessageBoxEx func(
		w T.HWND,
		text, caption VString,
		t T.MSGBOX_TYPE,
		languageId T.WORD) int

	MessageBoxIndirect func(mbp *T.MSGBOXPARAMS) int

	ModifyMenu func(
		mnu T.HMENU,
		position, flags T.UINT,
		idNewItem T.UINT_PTR,
		NewItem VString) T.BOOL

	MonitorFromPoint func(pt T.POINT, flags T.DWORD) T.HMONITOR

	MonitorFromRect func(rc *T.RECT, flags T.DWORD) T.HMONITOR

	MonitorFromWindow func(w T.HWND, flags T.DWORD) T.HMONITOR

	Mouse_event func(
		flags, x, y, data T.DWORD, extraInfo T.ULONG_PTR)

	MoveWindow func(
		w T.HWND, x, y, width, height int, repaint T.BOOL) T.BOOL

	MsgWaitForMultipleObjects func(
		count T.DWORD,
		h *T.HANDLE,
		waitAll T.BOOL,
		milliseconds T.DWORD,
		wakeMask T.DWORD) T.DWORD

	MsgWaitForMultipleObjectsEx func(
		count T.DWORD,
		h *T.HANDLE,
		milliseconds, wakeMask, flags T.DWORD) T.DWORD

	NotifyWinEvent func(
		event T.DWORD, w T.HWND, object, child T.LONG)

	OemKeyScan func(oemChar T.WORD) T.DWORD

	OemToChar func(src VString, dst OVString) T.BOOL

	OemToCharBuff func(
		src VString, dst OVString, dstLength T.DWORD) T.BOOL

	OffsetRect func(rc *T.RECT, x, y int) T.BOOL

	OpenClipboard func(wndNewOwner T.HWND) T.BOOL

	OpenDesktop func(
		desktop VString,
		flags T.DWORD,
		inherit T.BOOL,
		desiredAccess T.ACCESS_MASK) T.HDESK

	OpenIcon func(w T.HWND) T.BOOL

	OpenInputDesktop func(
		flags T.DWORD,
		inherit T.BOOL,
		desiredAccess T.ACCESS_MASK) T.HDESK

	OpenWindowStation func(
		winSta VString,
		inherit T.BOOL,
		desiredAccess T.ACCESS_MASK) T.HWINSTA

	PaintDesktop func(dc T.HDC) T.BOOL

	PeekMessage func(
		msg *T.MSG,
		w T.HWND,
		msgFilterMin T.UINT,
		msgFilterMax T.UINT,
		removeMsg T.UINT) T.BOOL

	PostMessage func(
		win T.HWND, msg T.UINT, w T.WPARAM, l T.LPARAM) T.BOOL

	PostQuitMessage func(exitCode int)

	PostThreadMessage func(
		thread T.DWORD, msg T.UINT, w T.WPARAM, l T.LPARAM) T.BOOL

	PrintWindow func(w T.HWND, blt T.HDC, flags T.UINT) T.BOOL

	PrivateExtractIcons func(
		fileName VString,
		iconIndex, xIcon, yIcon int,
		icon *T.HICON,
		iconid *T.UINT,
		icons, flags T.UINT) T.UINT

	PtInRect func(rc *T.RECT, pt T.POINT) T.BOOL

	RealChildWindowFromPoint func(
		parent T.HWND, parentClientCoords T.POINT) T.HWND

	RealGetWindowClass func(
		w T.HWND, className OVString, classNameMax T.UINT) T.UINT

	RedrawWindow func(
		w T.HWND, rcUpdate *T.RECT, rgnUpdate T.HRGN, flags T.UINT) T.BOOL

	RegisterClass func(*T.WNDCLASS) T.ATOM

	RegisterClassEx func(*T.WNDCLASSEX) T.ATOM

	RegisterClipboardFormat func(format VString) T.UINT

	RegisterDeviceNotification func(
		recipient T.HANDLE,
		notificationFilter *T.VOID, flags T.DWORD) T.HDEVNOTIFY

	RegisterHotKey func(
		w T.HWND, id int, modifiers T.UINT, vk T.UINT) T.BOOL

	RegisterRawInputDevices func(
		rids *T.RAWINPUTDEVICE, numDevices T.UINT, size T.UINT) T.BOOL

	RegisterShellHookWindow func(T.HWND) T.BOOL

	RegisterWindowMessage func(VString) T.UINT

	ReleaseCapture func() T.BOOL

	ReleaseDC func(T.HWND, T.HDC) int

	RemoveMenu func(m T.HMENU, position T.UINT, flags T.UINT) T.BOOL

	RemoveProp func(T.HWND, VString) T.HANDLE

	ReplyMessage func(T.LRESULT) T.BOOL

	ScreenToClient func(T.HWND, *T.POINT) T.BOOL

	ScrollDC func(
		dc T.HDC,
		x, y int,
		rcScroll, rcClip *T.RECT,
		rgnUpdate T.HRGN,
		rcUpdate *T.RECT) T.BOOL

	ScrollWindow func(
		w T.HWND, xAmount, yAmount int,
		rect, clipRect *T.RECT) T.BOOL

	ScrollWindowEx func(
		w T.HWND,
		x, y int,
		scroll, clip *T.RECT,
		rgnUpdate T.HRGN,
		rcUpdate *T.RECT,
		flags T.UINT) int

	SendDlgItemMessage func(
		dlg T.HWND,
		idDlgItem int,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM) T.LRESULT

	SendInput func(nInputs T.UINT, i *T.INPUT, size int) T.UINT

	SendMessage func(
		win T.HWND, msg T.UINT, w T.WPARAM, l T.LPARAM) T.LRESULT

	SendMessageCallback func(
		win T.HWND,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM,
		resultCallBack T.SENDASYNCPROC,
		data T.ULONG_PTR) T.BOOL

	SendMessageTimeout func(
		win T.HWND,
		msg T.UINT,
		w T.WPARAM,
		l T.LPARAM,
		flags, timeout T.UINT,
		result *T.DWORD_PTR) T.LRESULT

	SendNotifyMessage func(
		win T.HWND, msg T.UINT, w T.WPARAM, l T.LPARAM) T.BOOL

	SetActiveWindow func(w T.HWND) T.HWND

	SetCapture func(w T.HWND) T.HWND

	SetCaretBlinkTime func(mSeconds T.UINT) T.BOOL

	SetCaretPos func(x, y int) T.BOOL

	//SetClassWord is obsolete; instead use:
	//	SetClassLong
	SetClassLong func(w T.HWND, index int, newLong T.LONG) T.DWORD

	SetClassLongPtr func(
		w T.HWND, index int, newLong T.LONG_PTR) T.ULONG_PTR

	SetClassWord func(w T.HWND, index int, newWord T.WORD) T.WORD

	SetClipboardData func(format T.UINT, mem T.HANDLE) T.HANDLE

	SetClipboardViewer func(wndNewViewer T.HWND) T.HWND

	SetCursor func(cursor T.HCURSOR) T.HCURSOR

	SetCursorPos func(x, y int) T.BOOL

	SetDebugErrorLevel func(level T.DWORD)

	SetDlgItemInt func(
		dlg T.HWND, idDlgItem int, value T.UINT, signed T.BOOL) T.BOOL

	SetDlgItemText func(
		dlg T.HWND, idDlgItem int, s VString) T.BOOL

	SetDoubleClickTime func(T.UINT) T.BOOL

	SetFocus func(T.HWND) T.HWND

	SetForegroundWindow func(T.HWND) T.BOOL

	SetKeyboardState func(keyState *T.BYTE) T.BOOL

	SetLastErrorEx func(errCode T.DWORD, type_ T.DWORD)

	SetLayeredWindowAttributes func(
		w T.HWND, crKey T.COLORREF, alpha T.BYTE, flags T.DWORD) T.BOOL

	SetMenu func(T.HWND, T.HMENU) T.BOOL

	SetMenuContextHelpId func(T.HMENU, T.DWORD) T.BOOL

	SetMenuDefaultItem func(m T.HMENU, item T.UINT, byPos T.UINT) T.BOOL

	SetMenuInfo func(T.HMENU, T.MENUINFO) T.BOOL

	SetMenuItemBitmaps func(
		m T.HMENU,
		position, flags T.UINT,
		bitmapUnchecked, bitmapChecked T.HBITMAP) T.BOOL

	SetMenuItemInfo func(
		m T.HMENU,
		item T.UINT,
		byPositon T.BOOL,
		mii T.MENUITEMINFO) T.BOOL

	SetMessageExtraInfo func(T.LPARAM) T.LPARAM

	SetMessageQueue func(messagesMax int) T.BOOL

	SetParent func(child T.HWND, newParent T.HWND) T.HWND

	SetProcessDefaultLayout func(defaultLayout T.DWORD) T.BOOL

	SetProcessWindowStation func(T.HWINSTA) T.BOOL

	SetProp func(w T.HWND, s VString, data T.HANDLE) T.BOOL

	SetRect func(rc *T.RECT, xLeft, yTop, xRight, yBottom int) T.BOOL

	SetRectEmpty func(*T.RECT) T.BOOL

	SetScrollInfo func(
		w T.HWND, bar int, si *T.SCROLLINFO, redraw T.BOOL) int

	SetScrollPos func(w T.HWND, bar, pos, redraw T.BOOL) int

	SetScrollRange func(
		w T.HWND, bar, minPos, maxPos int, redraw T.BOOL) T.BOOL

	SetSysColors func(
		nElements int, elements *T.INT, rgbValues *T.COLORREF) T.BOOL

	SetSystemCursor func(cur T.HCURSOR, id T.DWORD) T.BOOL

	SetThreadDesktop func(desktop T.HDESK) T.BOOL

	SetTimer func(
		w T.HWND,
		idEvent T.UINT_PTR,
		elapse T.UINT,
		fn T.TIMERPROC) T.UINT_PTR

	SetUserObjectInformation func(
		obj T.HANDLE, index int, info *T.VOID, length T.DWORD) T.BOOL

	SetUserObjectSecurity func(
		obj T.HANDLE,
		requested *T.SECURITY_INFORMATION,
		sid *T.SECURITY_DESCRIPTOR) T.BOOL

	SetWindowContextHelpId func(T.HWND, T.DWORD) T.BOOL

	SetWindowLong func(w T.HWND, index int, newLong T.LONG) T.LONG

	SetWindowLongPtr func(
		w T.HWND, index int, newLong T.LONG_PTR) T.LONG_PTR

	SetWindowPlacement func(w T.HWND, wndpl *T.WINDOWPLACEMENT) T.BOOL

	SetWindowPos func(
		w T.HWND,
		insertAfter T.HWND,
		x, y, cx, cy int,
		flags T.UINT) T.BOOL

	SetWindowRgn func(w T.HWND, rgn T.HRGN, redraw T.BOOL) int

	SetWindowsHook func(filterType int, fn T.HOOKPROC) T.HHOOK

	SetWindowsHookEx func(
		hook int,
		fn T.HOOKPROC,
		mod T.HINSTANCE,
		threadId T.DWORD) T.HHOOK

	SetWindowText func(w T.HWND, s VString) T.BOOL

	SetWindowWord func(w T.HWND, index int, newWord T.WORD) T.WORD

	SetWinEventHook func(
		min, max T.DWORD,
		modWinEventProc T.HMODULE,
		fn T.WINEVENTPROC,
		process, thread T.DWORD,
		f T.WINEVENT_FLAGS) T.HWINEVENTHOOK

	ShowCaret func(w T.HWND) T.BOOL

	ShowCursor func(show T.BOOL) int

	ShowOwnedPopups func(w T.HWND, show T.BOOL) T.BOOL

	ShowScrollBar func(w T.HWND, bar int, show T.BOOL) T.BOOL

	ShowWindow func(w T.HWND, cmdShow int) T.BOOL

	ShowWindowAsync func(w T.HWND, cmdShow int) T.BOOL

	SubtractRect func(dst, src1, src2 *T.RECT) T.BOOL

	SwapMouseButton func(swap T.BOOL) T.BOOL

	SwitchDesktop func(desktop T.HDESK) T.BOOL

	SwitchToThisWindow func(w T.HWND, unknown T.BOOL)

	SystemParametersInfo func(
		action, nParam T.UINT, param *T.VOID, winIni T.UINT) T.BOOL

	TabbedTextOut func(
		h T.HDC,
		x, y int,
		s VString,
		count, tabPositions int,
		tabStopPositions *T.INT,
		tabOrigin int) T.LONG

	TileWindows func(
		parent T.HWND,
		how T.UINT,
		rect *T.RECT,
		nKids T.UINT,
		kids *T.HWND) T.WORD

	ToAscii func(
		virtKey, scanCode T.UINT,
		keyState *T.BYTE,
		char *T.WORD,
		flags T.UINT) int

	ToAsciiEx func(
		virtKey, scanCode T.UINT,
		keyState *T.BYTE,
		char *T.WORD,
		flags T.UINT,
		hkl T.HKL) int

	ToUnicode func(
		virtKey, scanCode T.UINT,
		keyState *T.BYTE,
		buff T.OWString,
		size int,
		flags T.UINT) int

	ToUnicodeEx func(
		virtKey, scanCode T.UINT,
		keyState *T.BYTE,
		buff T.WString,
		size int,
		flags T.UINT,
		h T.HKL) int

	TrackMouseEvent func(eventTrack *T.TRACKMOUSEEVENT) T.BOOL

	TrackPopupMenu func(
		m T.HMENU,
		flags T.UINT,
		x, y, reserved int,
		w T.HWND,
		rect *T.RECT) T.BOOL

	TrackPopupMenuEx func(
		T.HMENU, T.UINT, int, int, T.HWND, *T.TPMPARAMS) T.BOOL

	TranslateAccelerator func(
		w T.HWND, accTable T.HACCEL, msg *T.MSG) int

	TranslateMDISysAccel func(wndClient T.HWND, msg *T.MSG) T.BOOL

	TranslateMessage func(msg *T.MSG) T.BOOL

	UnhookWindowsHook func(code int, filterProc T.HOOKPROC) T.BOOL

	UnhookWindowsHookEx func(hk T.HHOOK) T.BOOL

	UnhookWinEvent func(winEventHook T.HWINEVENTHOOK) T.BOOL

	UnionRect func(dst, src1, src2 *T.RECT) T.BOOL

	UnloadKeyboardLayout func(hkl T.HKL) T.BOOL

	UnregisterClass func(
		className VString, instance T.HINSTANCE) T.BOOL

	UnregisterDeviceNotification func(handle T.HDEVNOTIFY) T.BOOL

	UnregisterHotKey func(w T.HWND, id int) T.BOOL

	UpdateLayeredWindow func(
		w T.HWND,
		dst T.HDC,
		ptDst *T.POINT,
		s *T.SIZE,
		src T.HDC,
		ptSrc *T.POINT,
		crKey T.COLORREF,
		blend *T.BLENDFUNCTION,
		flags T.DWORD) T.BOOL

	UpdateLayeredWindowIndirect func(
		w T.HWND, ulwi *T.UPDATELAYEREDWINDOWINFO) T.BOOL

	UpdateWindow func(w T.HWND) T.BOOL

	UserHandleGrantAccess func(
		user T.HANDLE, job T.HANDLE, grant T.BOOL) T.BOOL

	ValidateRect func(w T.HWND, rect *T.RECT) T.BOOL

	ValidateRgn func(w T.HWND, rgn T.HRGN) T.BOOL

	VkKeyScan func(ch T.Char) T.SHORT

	VkKeyScanEx func(ch T.Char, dwhkl T.HKL) T.SHORT

	WaitForInputIdle func(
		process T.HANDLE, milliseconds T.DWORD) T.DWORD

	WaitMessage func() T.BOOL

	WindowFromDC func(dc T.HDC) T.HWND

	WindowFromPoint func(point T.POINT) T.HWND

	WinHelp func(
		main T.HWND,
		help VString,
		command T.UINT,
		data T.ULONG_PTR) T.BOOL

	Wsprintf func(VString, VString, ...variadic) int

	Wvsprintf func(VString, VString, va_list) int
)

//TODO(t): Not xp32
//{"IsWow64Message",&IsWow64Message},
//{"UpdateLayeredWindowIndirect",&UpdateLayeredWindowIndirect},

var WinUserApis = Apis{
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
var WinUserUnicodeApis =Apis{
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

var WinUserANSIApis = Apis{
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
