// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package winuser provides API definitions for accessing
//user32.dll.
package winuser

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/user32"
)

func CopyCursor(cur HCURSOR) HCURSOR {
	return HCURSOR(CopyIcon(HICON(cur)))
}

func CreateDialog(Instance HINSTANCE, TemplateName VString, WndParent HWND, DialogFunc DLGPROC) HWND {
	return CreateDialogParam(Instance, TemplateName, WndParent, DialogFunc, 0)
}

func CreateDialogIndirect(Instance HINSTANCE, Template DLGTEMPLATE, WndParent HWND, DialogFunc DLGPROC) HWND {
	return CreateDialogIndirectParam(Instance, Template, WndParent, DialogFunc, 0)
}

func CreateWindow(className, WindowName VString, Style WINDOW_STYLE, X, Y, Width, Height int, WndParent HWND, Menu HMENU, Instance HINSTANCE, Param *VOID) HWND {
	return CreateWindowEx(0, className, WindowName, Style, X, Y, Width, Height, WndParent, Menu, Instance, Param)
}

func DialogBox(Instance HINSTANCE, Template VString, WndParent HWND, DialogFunc DLGPROC) INT_PTR {
	return DialogBoxParam(Instance, Template, WndParent, DialogFunc, 0)
}

func DialogBoxIndirect(Instance HINSTANCE, Template DLGTEMPLATE, WndParent HWND, DialogFunc DLGPROC) INT_PTR {
	return DialogBoxIndirectParam(Instance, Template, WndParent, DialogFunc, 0)
}

func EnumTaskWindows(Task DWORD, fn WNDENUMPROC, Param LPARAM) BOOL {
	return EnumThreadWindows(Task, fn, Param)
}

func GetNextWindow(Wnd HWND, cmd UINT) HWND {
	return GetWindow(Wnd, cmd)
}

func GetSysModalWindow() {}

func GetWindowTask(Wnd HWND) HANDLE {
	return HANDLE(GetWindowThreadProcessId(Wnd, nil))
}

func SetSysModalWindow(Wnd HWND) {}

type (
	va_list  Fake_type_Fix_me
	variadic Fake_type_Fix_me
)

var (
	ActivateKeyboardLayout func(hkl HKL, flags UINT) HKL

	AdjustWindowRect func(
		rect *RECT, style DWORD, menu BOOL) BOOL

	AdjustWindowRectEx func(
		rect *RECT, style DWORD, menu BOOL, exStyle DWORD) BOOL

	AllowSetForegroundWindow func(processId DWORD) BOOL

	AnimateWindow func(w HWND, time DWORD, flags DWORD) BOOL

	AnyPopup func() BOOL

	AppendMenu func(
		m HMENU,
		flags UINT,
		idNewItem UINT_PTR,
		newItem VString) BOOL

	ArrangeIconicWindows func(w HWND) UINT

	AttachThreadInput func(attach, to DWORD, flag BOOL) BOOL

	BeginDeferWindowPos func(numWindows int) HDWP

	BeginPaint func(w HWND, paint *PAINTSTRUCT) HDC

	BringWindowToTop func(w HWND) BOOL

	BroadcastSystemMessage func(
		flags DWORD,
		info *DWORD,
		msg UINT,
		w WPARAM,
		l LPARAM) LONG

	BroadcastSystemMessageEx func(
		flags BSM_FLAGS,
		recipients *DWORD,
		msg UINT,
		w WPARAM,
		l LPARAM,
		info *BSMINFO) LONG

	CallMsgFilter func(msg *MSG, code int) BOOL

	CallNextHookEx func(
		hk HHOOK, code int, w WPARAM, l LPARAM) LRESULT

	CallWindowProc func(
		prevWndFunc WNDPROC,
		win HWND,
		msg UINT,
		w WPARAM,
		l LPARAM) LRESULT

	CascadeWindows func(
		wndParent HWND,
		how UINT,
		rect *RECT,
		nKids UINT,
		kids *HWND) WORD

	ChangeClipboardChain func(wndRemove, wndNewNext HWND) BOOL

	ChangeDisplaySettingsA func(dm *DEVMODEA, flags DWORD) LONG

	ChangeDisplaySettingsW func(dm *DEVMODEW, flags DWORD) LONG

	ChangeDisplaySettingsExA func(
		deviceName VString,
		dm *DEVMODEA,
		w HWND,
		flags DWORD,
		lParam *VOID) LONG

	ChangeDisplaySettingsExW func(
		deviceName VString,
		dm *DEVMODEW,
		w HWND,
		flags DWORD,
		lParam *VOID) LONG

	ChangeMenu func(
		m HMENU,
		cmd UINT,
		NewItem VString,
		insert UINT,
		flags UINT) BOOL

	CharLower func(VString) OVString

	CharLowerBuff func(buff IOAString, length DWORD) DWORD

	CharNext func(VString) OVString

	CharNextExA func(
		codePage WORD, currentChar AString, flags DWORD) AString

	CharPrev func(start, current VString) VString

	CharPrevExA func(
		codePage WORD,
		start, currentChar AString,
		flags DWORD) AString

	CharToOem func(src VString, dst OVString) BOOL

	CharToOemBuff func(src, dst OVString, dstLength DWORD) BOOL

	CharUpper func(VString) OVString

	CharUpperBuff func(buff IOAString, length DWORD) DWORD

	CheckDlgButton func(dlg HWND, idButton int, check UINT) BOOL

	CheckMenuItem func(m HMENU, idCheckItem, check UINT) DWORD

	CheckMenuRadioItem func(
		m HMENU, first, last, check, flags UINT) BOOL

	CheckRadioButton func(
		dlg HWND, idFirst, idLast, idCheck int) BOOL

	ChildWindowFromPoint func(wndParent HWND, point POINT) HWND

	ChildWindowFromPointEx func(
		w HWND, p POINT, flags UINT) HWND

	ClientToScreen func(w HWND, p *POINT) BOOL

	ClipCursor func(rect *RECT) BOOL

	CloseClipboard func() BOOL

	CloseDesktop func(desktop HDESK) BOOL

	CloseWindow func(w HWND) BOOL

	CloseWindowStation func(winSta HWINSTA) BOOL

	CopyAcceleratorTable func(
		src HACCEL, dst *ACCEL, entries int) int

	CopyIcon func(icon HICON) HICON

	CopyImage func(
		h HANDLE, typ UINT, X, y int, flags UINT) HANDLE

	CopyRect func(dst, src *RECT) BOOL

	CountClipboardFormats func() int

	CreateAcceleratorTable func(accel *ACCEL, nAccel int) HACCEL

	CreateCaret func(
		w HWND, bitmap HBITMAP, width int, height int) BOOL

	CreateCursor func(
		inst HINSTANCE,
		xHotSpot, yHotSpot, width, height int,
		andPlane, xorPlane *VOID) HCURSOR

	CreateDesktopA func(
		desktop, device VString,
		devmode *DEVMODEA,
		flags DWORD,
		desiredAccess ACCESS_MASK,
		sa *SECURITY_ATTRIBUTES) HDESK

	CreateDesktopW func(
		desktop, device VString,
		devmode *DEVMODEW,
		flags DWORD,
		desiredAccess ACCESS_MASK,
		sa *SECURITY_ATTRIBUTES) HDESK

	CreateDialogIndirectParam func(
		instance HINSTANCE,
		template DLGTEMPLATE,
		wndParent HWND,
		dialogFunc DLGPROC,
		initParam LPARAM) HWND

	CreateDialogParam func(
		instance HINSTANCE,
		templateName VString,
		wndParent HWND,
		dialogFunc DLGPROC,
		initParam LPARAM) HWND

	CreateIcon func(
		instance HINSTANCE,
		width, height int,
		planes, bitsPixel BYTE,
		andbits, xorbits *BYTE) HICON

	CreateIconFromResource func(
		resbits *BYTE,
		resSize DWORD,
		icon BOOL,
		ver DWORD) HICON

	CreateIconFromResourceEx func(
		resbits *BYTE,
		resSize DWORD,
		icon BOOL,
		ver DWORD,
		xDesired, yDesired int,
		flags UINT) HICON

	CreateIconIndirect func(iconinfo *ICONINFO) HICON

	CreateMDIWindow func(
		className, windowName VString,
		style DWORD,
		x, y, width,
		height int,
		wndParent HWND,
		instance HINSTANCE,
		l LPARAM) HWND

	CreateMenu func() HMENU

	CreatePopupMenu func() HMENU

	CreateWindowEx func(
		exStyle WINDOW_STYLE_EX,
		className, windowName VString,
		style WINDOW_STYLE,
		x, y, width, height int,
		wndParent HWND,
		m HMENU,
		instance HINSTANCE,
		param *VOID) HWND

	CreateWindowStation func(
		winsta VString,
		flags DWORD,
		desiredAccess ACCESS_MASK,
		sa *SECURITY_ATTRIBUTES) HWINSTA

	DefDlgProc func(
		dlg HWND, msg UINT, w WPARAM, l LPARAM) LRESULT

	DeferWindowPos func(
		winPosInfo HDWP,
		w HWND,
		wndInsertAfter HWND,
		x, y, cx, cy int,
		flags UINT) HDWP

	DefFrameProc func(
		win HWND,
		mdiClient HWND,
		msg UINT,
		w WPARAM,
		l LPARAM) LRESULT

	DefMDIChildProc func(
		win HWND, msg UINT, w WPARAM, l LPARAM) LRESULT

	DefRawInputProc func(rawInput *RAWINPUT, input INT, sizeHeader INT) LRESULT

	DefWindowProc func(h HWND, m WINDOW_MESSAGE, w WPARAM, l LPARAM) LRESULT

	DeleteMenu func(m HMENU, position UINT, flags UINT) BOOL

	DeregisterShellHookWindow func(w HWND) BOOL

	DestroyAcceleratorTable func(accel HACCEL) BOOL

	DestroyCaret func() BOOL

	DestroyCursor func(cursor HCURSOR) BOOL

	DestroyIcon func(icon HICON) BOOL

	DestroyMenu func(m HMENU) BOOL

	DestroyWindow func(w HWND) BOOL

	DialogBoxIndirectParam func(
		instance HINSTANCE,
		dialogTemplate DLGTEMPLATE,
		wndParent HWND,
		dialogFunc DLGPROC,
		initParam LPARAM) INT_PTR

	DialogBoxParam func(
		instance HINSTANCE,
		templateName VString,
		wndParent HWND,
		dialogFunc DLGPROC,
		initParam LPARAM) INT_PTR

	DisableProcessWindowsGhosting func()

	DispatchMessage func(msg *MSG) LRESULT

	DlgDirList func(
		dlg HWND,
		pathSpec VString,
		idListBox, idStaticPath int,
		fileType UINT) int

	DlgDirListComboBox func(
		dlg HWND,
		pathSpec VString,
		idComboBox, idStaticPath int,
		filetype UINT) int

	DlgDirSelectComboBoxEx func(
		dlg HWND,
		string_ VString,
		out int,
		comboBox int) BOOL

	DlgDirSelectEx func(
		dlg HWND,
		string_ VString,
		count int,
		listBox int) BOOL

	DragDetect func(w HWND, pt POINT) BOOL

	DragObject func(
		parent HWND,
		from HWND,
		fmt UINT,
		data ULONG_PTR,
		cur HCURSOR) DWORD

	DrawAnimatedRects func(
		hw HWND, ani int, from *RECT, to *RECT) BOOL

	DrawCaption func(
		w HWND, dc HDC, rect *RECT, flags UINT) BOOL

	DrawEdge func(
		hdc HDC, qrc *RECT, edge UINT, grfFlags UINT) BOOL

	DrawFocusRect func(dc HDC, rc *RECT) BOOL

	DrawFrameControl func(HDC, *RECT, UINT, UINT) BOOL

	DrawIcon func(dc HDC, x, y int, icon HICON) BOOL

	DrawIconEx func(
		dc HDC,
		xLeft, yTop int,
		icon HICON,
		xWidth, yWidth int,
		stepIfAniCur UINT,
		brFlickerFreeDraw HBRUSH,
		flags UINT) BOOL

	DrawMenuBar func(w HWND) BOOL

	DrawState func(
		dc HDC,
		brFore HBRUSH,
		fnCallBack DRAWSTATEPROC,
		lData LPARAM,
		wData WPARAM,
		x, y, cx, cy int,
		uFlags UINT) BOOL

	DrawText func(
		dc HDC,
		text VString,
		nText int,
		rc *RECT,
		format UINT) int

	DrawTextEx func(
		dc HDC,
		text VString,
		nText int,
		rc *RECT,
		format UINT,
		dtp *DRAWTEXTPARAMS) int

	EmptyClipboard func() BOOL

	EnableMenuItem func(
		m HMENU, idDEnableItem UINT, enable UINT) BOOL

	EnableScrollBar func(w HWND, sbflags UINT, arrows UINT) BOOL

	EnableWindow func(w HWND, enable BOOL) BOOL

	EndDeferWindowPos func(winPosInfo HDWP) BOOL

	EndDialog func(dlg HWND, result INT_PTR) BOOL

	EndMenu func() BOOL

	EndPaint func(w HWND, paint *PAINTSTRUCT) BOOL

	EndTask func(w HWND, shutDown BOOL, force BOOL) BOOL

	EnumChildWindows func(
		parent HWND, fn WNDENUMPROC, l LPARAM) BOOL

	EnumClipboardFormats func(format UINT) UINT

	EnumDesktops func(
		wsta HWINSTA, fn DESKTOPENUMPROC, l LPARAM) BOOL

	EnumDesktopWindows func(
		desktop HDESK, fn WNDENUMPROC, param LPARAM) BOOL

	EnumDisplayDevicesA func(
		device VString,
		devNum DWORD,
		displayDevice *DISPLAY_DEVICEA,
		flags DWORD) BOOL

	EnumDisplayDevicesW func(
		device VString,
		devNum DWORD,
		displayDevice *DISPLAY_DEVICEW,
		flags DWORD) BOOL

	EnumDisplayMonitors func(
		dc HDC,
		clip *RECT,
		fn MONITORENUMPROC,
		data LPARAM) BOOL

	EnumDisplaySettingsA func(
		deviceName VString,
		modeNum DWORD,
		dm *DEVMODEA) BOOL

	EnumDisplaySettingsW func(
		deviceName VString,
		modeNum DWORD,
		dm *DEVMODEW) BOOL

	EnumDisplaySettingsExA func(
		deviceName VString,
		modeNum DWORD,
		dm *DEVMODEA,
		flags DWORD) BOOL

	EnumDisplaySettingsExW func(
		deviceName VString,
		modeNum DWORD,
		dm *DEVMODEW,
		flags DWORD) BOOL

	EnumPropsA func(w HWND, enumFunc PROPENUMPROCA) int

	EnumPropsW func(w HWND, enumFunc PROPENUMPROCW) int

	EnumPropsExA func(
		w HWND, enumFunc PROPENUMPROCEXA, l LPARAM) int

	EnumPropsExW func(
		w HWND, enumFunc PROPENUMPROCEXW, l LPARAM) int

	EnumThreadWindows func(
		threadId DWORD, fn WNDENUMPROC, l LPARAM) BOOL

	EnumWindows func(enumFunc WNDENUMPROC, l LPARAM) BOOL

	EnumWindowStations func(
		enumFunc WINSTAENUMPROC, l LPARAM) BOOL

	EqualRect func(rc1 *RECT, rc2 *RECT) BOOL

	ExcludeUpdateRgn func(dc HDC, w HWND) int

	ExitWindowsEx func(flags UINT, reason DWORD) BOOL

	FillRect func(dc HDC, rc *RECT, br HBRUSH) int

	FindWindow func(className VString, windowName VString) HWND

	FindWindowEx func(
		wndParent, wndChildAfter HWND,
		class, window VString) HWND

	FlashWindow func(w HWND, invert BOOL) BOOL

	FlashWindowEx func(fwi *FLASHWINFO) BOOL

	FrameRect func(dc HDC, rc *RECT, br HBRUSH) int

	GetActiveWindow func() HWND

	GetAltTabInfo func(
		w HWND,
		item int,
		ati *ALTTABINFO,
		itemText OVString,
		size UINT) BOOL

	GetAncestor func(w HWND, flags UINT) HWND

	GetAsyncKeyState func(key int) SHORT

	GetCapture func() HWND

	GetCaretBlinkTime func() UINT

	GetCaretPos func(point *POINT) BOOL

	GetClassInfo func(
		instance HINSTANCE, className VString, wc *WNDCLASS) BOOL

	GetClassInfoEx func(
		instance HINSTANCE, class VString, wcx *WNDCLASSEX) BOOL

	GetClassLong func(w HWND, index int) DWORD

	GetClassLongPtr func(w HWND, index int) ULONG_PTR

	//GetClassName func(
	//	w HWND, className OVString, maxCount int) int
	GetClassName func(w HWND, className *WChar, maxCount int) int

	//GetClassWord is obsolete; instead use:
	//	GetClassLong
	GetClassWord func(w HWND, index int) WORD

	GetClientRect func(w HWND, rect *RECT) BOOL

	GetClipboardData func(format UINT) HANDLE

	GetClipboardFormatName func(
		format UINT, formatName VString, maxCount int) int

	GetClipboardOwner func() HWND

	GetClipboardSequenceNumber func() DWORD

	GetClipboardViewer func() HWND

	GetClipCursor func(rect *RECT) BOOL

	GetComboBoxInfo func(w HWND, cbi *COMBOBOXINFO) BOOL

	GetCursor func() HCURSOR

	GetCursorInfo func(ci *CURSORINFO) BOOL

	GetCursorPos func(point *POINT) BOOL

	GetDC func(w HWND) HDC

	GetDCEx func(w HWND, rgnClip HRGN, flags DWORD) HDC

	GetDesktopWindow func() HWND

	GetDialogBaseUnits func() LONG

	GetDlgCtrlID func(w HWND) int

	GetDlgItem func(dlg HWND, idDlgItem int) HWND

	GetDlgItemInt func(
		dlg HWND,
		idDlgItem int,
		translated *BOOL,
		signed BOOL) UINT

	GetDlgItemText func(
		dlg HWND,
		idDlgItem int,
		string_ OVString,
		max int) UINT

	GetDoubleClickTime func() UINT

	GetFocus func() HWND

	GetForegroundWindow func() HWND

	GetGuiResources func(process HANDLE, flags DWORD) DWORD

	GetGUIThreadInfo func(thread DWORD, ui *GUITHREADINFO) BOOL

	GetIconInfo func(icon HICON, iconinfo *ICONINFO) BOOL

	GetInputState func() BOOL

	GetKBCodePage func() UINT

	GetKeyboardLayout func(thread DWORD) HKL

	GetKeyboardLayoutList func(buff int, list *HKL) int

	GetKeyboardLayoutName func(klid VString) BOOL

	GetKeyboardState func(keyState *BYTE) BOOL

	GetKeyboardType func(typeFlag int) int

	GetKeyNameText func(
		param LONG, string_ OVString, size int) int

	GetKeyState func(virtKey int) SHORT

	GetLastActivePopup func(w HWND) HWND

	GetLastInputInfo func(lii *LASTINPUTINFO) BOOL

	GetLayeredWindowAttributes func(
		w HWND, crKey *COLORREF, alpha *BYTE, flags *DWORD) BOOL

	GetListBoxInfo func(w HWND) DWORD

	GetMenu func(w HWND) HMENU

	GetMenuBarInfo func(
		w HWND, object, item LONG, mbi *MENUBARINFO) BOOL

	GetMenuCheckMarkDimensions func() LONG

	GetMenuContextHelpId func(HMENU) DWORD

	GetMenuDefaultItem func(
		m HMENU, byPos UINT, gmdiFlags UINT) UINT

	GetMenuInfo func(HMENU, *MENUINFO) BOOL

	GetMenuItemCount func(m HMENU) int

	GetMenuItemID func(m HMENU, pos int) UINT

	GetMenuItemInfo func(
		m HMENU,
		item UINT,
		byPosition BOOL,
		mii *MENUITEMINFO) BOOL

	GetMenuItemRect func(
		w HWND, m HMENU, item UINT, rcItem *RECT) BOOL

	GetMenuState func(m HMENU, id UINT, flags UINT) UINT

	GetMenuString func(
		m HMENU,
		idItem UINT,
		string_ OVString,
		max int,
		flags UINT) int

	GetMessage func(
		m *MSG, w HWND, filterMin, filterMax UINT) BOOL

	GetMessageExtraInfo func() LPARAM

	GetMessagePos func() DWORD

	GetMessageTime func() LONG

	GetMonitorInfo func(monitor HMONITOR, mi *MONITORINFO) BOOL

	GetMouseMovePointsEx func(
		size UINT,
		pt, ptBuf *MOUSEMOVEPOINT,
		bufPoints int,
		resolution DWORD) int

	GetNextDlgGroupItem func(
		dlg, ctl HWND, previous BOOL) HWND

	GetNextDlgTabItem func(
		dlg, ctl HWND, previous BOOL) HWND

	GetOpenClipboardWindow func() HWND

	GetParent func(w HWND) HWND

	GetPriorityClipboardFormat func(
		formatPriorityList *UINT, formats int) int

	GetProcessDefaultLayout func(defaultLayout *DWORD) BOOL

	GetProcessWindowStation func() HWINSTA

	GetProp func(w HWND, string_ VString) HANDLE

	GetQueueStatus func(flags UINT) DWORD

	GetRawInputBuffer func(
		data *RAWINPUT, size *UINT, sizeHeader UINT) UINT

	GetRawInputData func(
		rawInput HRAWINPUT,
		command UINT,
		data *VOID,
		size *UINT,
		sizeHeader UINT) UINT

	GetRawInputDeviceInfo func(
		device HANDLE,
		command UINT,
		data *VOID,
		size *UINT) UINT

	GetRawInputDeviceList func(
		ridl *RAWINPUTDEVICELIST,
		numDevices *UINT,
		size UINT) UINT

	GetRegisteredRawInputDevices func(
		rid *RAWINPUTDEVICE, numDevices *UINT, size UINT) UINT

	GetScrollBarInfo func(
		w HWND, object LONG, sbi *SCROLLBARINFO) BOOL

	GetScrollInfo func(w HWND, bar int, si *SCROLLINFO) BOOL

	GetScrollPos func(w HWND, bar int) int

	GetScrollRange func(
		w HWND, bar int, minPos *INT, maxPos *INT) BOOL

	GetShellWindow func() HWND

	GetSubMenu func(m HMENU, pos int) HMENU

	GetSysColor func(index int) DWORD

	GetSysColorBrush func(index int) HBRUSH

	GetSystemMenu func(w HWND, revert BOOL) HMENU

	GetSystemMetrics func(index int) int

	GetTabbedTextExtent func(
		dc HDC,
		s VString,
		count, tabPositions int,
		tabStopPositions *INT) DWORD

	GetThreadDesktop func(threadId DWORD) HDESK

	GetTitleBarInfo func(w HWND, ti *TITLEBARINFO) BOOL

	GetTopWindow func(w HWND) HWND

	GetUpdateRect func(w HWND, rect *RECT, erase BOOL) BOOL

	GetUpdateRgn func(w HWND, rgn HRGN, erase BOOL) int

	GetUserObjectInformation func(
		obj HANDLE,
		index int,
		info *VOID,
		length DWORD,
		lengthNeeded *DWORD) BOOL

	GetUserObjectSecurity func(
		obj HANDLE,
		siRequested *SECURITY_INFORMATION,
		sd *SECURITY_DESCRIPTOR,
		length DWORD,
		lengthNeeded *DWORD) BOOL

	GetWindow func(w HWND, cmd UINT) HWND

	GetWindowContextHelpId func(HWND) DWORD

	GetWindowDC func(w HWND) HDC

	GetWindowInfo func(w HWND, wi *WINDOWINFO) BOOL

	GetWindowLong func(w HWND, index int) LONG

	GetWindowLongPtr func(w HWND, index int) LONG_PTR

	GetWindowModuleFileName func(
		w HWND, fileName OVString, fileNameMax UINT) UINT

	GetWindowPlacement func(w HWND, wndpl *WINDOWPLACEMENT) BOOL

	GetWindowRect func(w HWND, rect *RECT) BOOL

	GetWindowRgn func(w HWND, rgn HRGN) int

	GetWindowRgnBox func(w HWND, rc *RECT) int

	// GetWindowText func(w HWND, s OVString, maxCount int) int
	GetWindowText func(w HWND, s *WChar, maxCount int) int

	GetWindowTextLength func(w HWND) int

	GetWindowThreadProcessId func(
		w HWND, processId *DWORD) DWORD

	GetWindowWord func(w HWND, index int) WORD

	GrayString func(
		dc HDC,
		brush HBRUSH,
		outputFunc GRAYSTRINGPROC,
		data LPARAM,
		count, x, y, width, height int) BOOL

	HideCaret func(w HWND) BOOL

	HiliteMenuItem func(
		w HWND, m HMENU, idHiliteItem UINT, hilite UINT) BOOL

	InflateRect func(rc *RECT, x, y int) BOOL

	InSendMessage func() BOOL

	InSendMessageEx func(reserved *VOID) DWORD

	InsertMenu func(
		m HMENU,
		position UINT,
		flags UINT,
		idNewItem UINT_PTR,
		NewItem AString) BOOL

	InsertMenuItem func(
		m HMENU,
		item UINT,
		byPosition BOOL,
		mii MENUITEMINFO) BOOL

	InternalGetWindowText func(w HWND, s OWString, size int) int

	IntersectRect func(dst, src1, src2 *RECT) BOOL

	InvalidateRect func(w HWND, rect *RECT, erase BOOL) BOOL

	InvalidateRgn func(w HWND, rgn HRGN, erase BOOL) BOOL

	InvertRect func(dc HDC, rc *RECT) BOOL

	IsCharAlpha func(ch VChar) BOOL

	IsCharAlphaNumeric func(ch VChar) BOOL

	IsCharLower func(ch VChar) BOOL

	IsCharUpper func(ch VChar) BOOL

	IsChild func(wndParent HWND, w HWND) BOOL

	IsClipboardFormatAvailable func(format UINT) BOOL

	IsDialogMessage func(dlg HWND, msg *MSG) BOOL

	IsDlgButtonChecked func(dlg HWND, idButton int) UINT

	IsGUIThread func(convert BOOL) BOOL

	IsHungAppWindow func(w HWND) BOOL

	IsIconic func(w HWND) BOOL

	IsMenu func(m HMENU) BOOL

	IsRectEmpty func(rc *RECT) BOOL

	IsWindow func(w HWND) BOOL

	IsWindowEnabled func(w HWND) BOOL

	IsWindowUnicode func(w HWND) BOOL

	IsWindowVisible func(w HWND) BOOL

	IsWinEventHookInstalled func(event DWORD) BOOL

	IsWow64Message func() BOOL

	IsZoomed func(w HWND) BOOL

	Keybd_event func(
		vk BYTE, scan BYTE, flags DWORD, extraInfo ULONG_PTR)

	KillTimer func(w HWND, idEvent UINT_PTR) BOOL

	LoadAccelerators func(i HINSTANCE, tableName VString) HACCEL

	LoadBitmap func(i HINSTANCE, bitmapName VString) HBITMAP

	LoadCursor func(i HINSTANCE, cursorName PVString) HCURSOR

	LoadCursorFromFile func(fileName VString) HCURSOR

	LoadIcon func(i HINSTANCE, iconName VString) HICON

	LoadImage func(
		i HINSTANCE,
		name VString,
		type_ UINT,
		x, y int,
		load UINT) HANDLE

	LoadKeyboardLayout func(klid VString, flags UINT) HKL

	LoadMenu func(instance HINSTANCE, menuName VString) HMENU

	//TODO(t):How does yhis work (*VOID underneath)
	LoadMenuIndirect func(mt *MENUTEMPLATE) HMENU

	LoadString func(
		i HINSTANCE, id UINT, s VString, size int) int

	LockSetForegroundWindow func(lockCode UINT) BOOL

	LockWindowUpdate func(wndLock HWND) BOOL

	LockWorkStation func() BOOL

	LookupIconIdFromDirectory func(
		resbits *BYTE, icon BOOL) int

	LookupIconIdFromDirectoryEx func(
		resbits *BYTE,
		icon BOOL,
		xDesired, yDesired int,
		flags UINT) int

	MapDialogRect func(dlg HWND, rect *RECT) BOOL

	MapVirtualKey func(code UINT, mapType UINT) UINT

	MapVirtualKeyEx func(code UINT, mapType UINT, hkl HKL) UINT

	MapWindowPoints func(
		from HWND, to HWND, p *POINT, nPoints UINT) int

	MenuItemFromPoint func(w HWND, m HMENU, ptScreen POINT) int

	MessageBeep func(type_ UINT) BOOL
)

var MessageBox func(
	w HWND, text, caption VString, t MSGBOX_TYPE) int

var (
	MessageBoxEx func(
		w HWND,
		text, caption VString,
		t MSGBOX_TYPE,
		languageId WORD) int

	MessageBoxIndirect func(mbp *MSGBOXPARAMS) int

	ModifyMenu func(
		mnu HMENU,
		position, flags UINT,
		idNewItem UINT_PTR,
		NewItem VString) BOOL

	MonitorFromPoint func(pt POINT, flags DWORD) HMONITOR

	MonitorFromRect func(rc *RECT, flags DWORD) HMONITOR

	MonitorFromWindow func(w HWND, flags DWORD) HMONITOR

	Mouse_event func(
		flags, x, y, data DWORD, extraInfo ULONG_PTR)

	MoveWindow func(
		w HWND, x, y, width, height int, repaint BOOL) BOOL

	MsgWaitForMultipleObjects func(
		count DWORD,
		h *HANDLE,
		waitAll BOOL,
		milliseconds DWORD,
		wakeMask DWORD) DWORD

	MsgWaitForMultipleObjectsEx func(
		count DWORD,
		h *HANDLE,
		milliseconds, wakeMask, flags DWORD) DWORD

	NotifyWinEvent func(
		event DWORD, w HWND, object, child LONG)

	OemKeyScan func(oemChar WORD) DWORD

	OemToChar func(src VString, dst OVString) BOOL

	OemToCharBuff func(
		src VString, dst OVString, dstLength DWORD) BOOL

	OffsetRect func(rc *RECT, x, y int) BOOL

	OpenClipboard func(wndNewOwner HWND) BOOL

	OpenDesktop func(
		desktop VString,
		flags DWORD,
		inherit BOOL,
		desiredAccess ACCESS_MASK) HDESK

	OpenIcon func(w HWND) BOOL

	OpenInputDesktop func(
		flags DWORD,
		inherit BOOL,
		desiredAccess ACCESS_MASK) HDESK

	OpenWindowStation func(
		winSta VString,
		inherit BOOL,
		desiredAccess ACCESS_MASK) HWINSTA

	PaintDesktop func(dc HDC) BOOL

	PeekMessage func(
		msg *MSG,
		w HWND,
		msgFilterMin UINT,
		msgFilterMax UINT,
		removeMsg UINT) BOOL

	PostMessage func(
		win HWND, msg UINT, w WPARAM, l LPARAM) BOOL

	PostQuitMessage func(exitCode int)

	PostThreadMessage func(
		thread DWORD, msg UINT, w WPARAM, l LPARAM) BOOL

	PrintWindow func(w HWND, blt HDC, flags UINT) BOOL

	PrivateExtractIcons func(
		fileName VString,
		iconIndex, xIcon, yIcon int,
		icon *HICON,
		iconid *UINT,
		icons, flags UINT) UINT

	PtInRect func(rc *RECT, pt POINT) BOOL

	RealChildWindowFromPoint func(
		parent HWND, parentClientCoords POINT) HWND

	RealGetWindowClass func(
		w HWND, className OVString, classNameMax UINT) UINT

	RedrawWindow func(
		w HWND, rcUpdate *RECT, rgnUpdate HRGN, flags UINT) BOOL

	RegisterClass func(*WNDCLASS) ATOM

	RegisterClassEx func(*WNDCLASSEX) ATOM

	RegisterClipboardFormat func(format VString) UINT

	RegisterDeviceNotification func(
		recipient HANDLE,
		notificationFilter *VOID, flags DWORD) HDEVNOTIFY

	RegisterHotKey func(
		w HWND, id int, modifiers UINT, vk UINT) BOOL

	RegisterRawInputDevices func(
		rids *RAWINPUTDEVICE, numDevices UINT, size UINT) BOOL

	RegisterShellHookWindow func(HWND) BOOL

	RegisterWindowMessage func(VString) UINT

	ReleaseCapture func() BOOL

	ReleaseDC func(HWND, HDC) int

	RemoveMenu func(m HMENU, position UINT, flags UINT) BOOL

	RemoveProp func(HWND, VString) HANDLE

	ReplyMessage func(LRESULT) BOOL

	ScreenToClient func(HWND, *POINT) BOOL

	ScrollDC func(
		dc HDC,
		x, y int,
		rcScroll, rcClip *RECT,
		rgnUpdate HRGN,
		rcUpdate *RECT) BOOL

	ScrollWindow func(
		w HWND, xAmount, yAmount int,
		rect, clipRect *RECT) BOOL

	ScrollWindowEx func(
		w HWND,
		x, y int,
		scroll, clip *RECT,
		rgnUpdate HRGN,
		rcUpdate *RECT,
		flags UINT) int

	SendDlgItemMessage func(
		dlg HWND,
		idDlgItem int,
		msg UINT,
		w WPARAM,
		l LPARAM) LRESULT

	SendInput func(nInputs UINT, i *INPUT, size int) UINT

	SendMessage func(
		win HWND, msg UINT, w WPARAM, l LPARAM) LRESULT

	SendMessageCallback func(
		win HWND,
		msg UINT,
		w WPARAM,
		l LPARAM,
		resultCallBack SENDASYNCPROC,
		data ULONG_PTR) BOOL

	SendMessageTimeout func(
		win HWND,
		msg UINT,
		w WPARAM,
		l LPARAM,
		flags, timeout UINT,
		result *DWORD_PTR) LRESULT

	SendNotifyMessage func(
		win HWND, msg UINT, w WPARAM, l LPARAM) BOOL

	SetActiveWindow func(w HWND) HWND

	SetCapture func(w HWND) HWND

	SetCaretBlinkTime func(mSeconds UINT) BOOL

	SetCaretPos func(x, y int) BOOL

	//SetClassWord is obsolete; instead use:
	//	SetClassLong
	SetClassLong func(w HWND, index int, newLong LONG) DWORD

	SetClassLongPtr func(
		w HWND, index int, newLong LONG_PTR) ULONG_PTR

	SetClassWord func(w HWND, index int, newWord WORD) WORD

	SetClipboardData func(format UINT, mem HANDLE) HANDLE

	SetClipboardViewer func(wndNewViewer HWND) HWND

	SetCursor func(cursor HCURSOR) HCURSOR

	SetCursorPos func(x, y int) BOOL

	SetDebugErrorLevel func(level DWORD)

	SetDlgItemInt func(
		dlg HWND, idDlgItem int, value UINT, signed BOOL) BOOL

	SetDlgItemText func(
		dlg HWND, idDlgItem int, s VString) BOOL

	SetDoubleClickTime func(UINT) BOOL

	SetFocus func(HWND) HWND

	SetForegroundWindow func(HWND) BOOL

	SetKeyboardState func(keyState *BYTE) BOOL

	SetLastErrorEx func(errCode DWORD, type_ DWORD)

	SetLayeredWindowAttributes func(
		w HWND, crKey COLORREF, alpha BYTE, flags DWORD) BOOL

	SetMenu func(HWND, HMENU) BOOL

	SetMenuContextHelpId func(HMENU, DWORD) BOOL

	SetMenuDefaultItem func(m HMENU, item UINT, byPos UINT) BOOL

	SetMenuInfo func(HMENU, MENUINFO) BOOL

	SetMenuItemBitmaps func(
		m HMENU,
		position, flags UINT,
		bitmapUnchecked, bitmapChecked HBITMAP) BOOL

	SetMenuItemInfo func(
		m HMENU,
		item UINT,
		byPositon BOOL,
		mii MENUITEMINFO) BOOL

	SetMessageExtraInfo func(LPARAM) LPARAM

	SetMessageQueue func(messagesMax int) BOOL

	SetParent func(child HWND, newParent HWND) HWND

	SetProcessDefaultLayout func(defaultLayout DWORD) BOOL

	SetProcessWindowStation func(HWINSTA) BOOL

	SetProp func(w HWND, s VString, data HANDLE) BOOL

	SetRect func(rc *RECT, xLeft, yTop, xRight, yBottom int) BOOL

	SetRectEmpty func(*RECT) BOOL

	SetScrollInfo func(
		w HWND, bar int, si *SCROLLINFO, redraw BOOL) int

	SetScrollPos func(w HWND, bar, pos, redraw BOOL) int

	SetScrollRange func(
		w HWND, bar, minPos, maxPos int, redraw BOOL) BOOL

	SetSysColors func(
		nElements int, elements *INT, rgbValues *COLORREF) BOOL

	SetSystemCursor func(cur HCURSOR, id DWORD) BOOL

	SetThreadDesktop func(desktop HDESK) BOOL

	SetTimer func(
		w HWND,
		idEvent UINT_PTR,
		elapse UINT,
		fn TIMERPROC) UINT_PTR

	SetUserObjectInformation func(
		obj HANDLE, index int, info *VOID, length DWORD) BOOL

	SetUserObjectSecurity func(
		obj HANDLE,
		requested *SECURITY_INFORMATION,
		sid *SECURITY_DESCRIPTOR) BOOL

	SetWindowContextHelpId func(HWND, DWORD) BOOL

	SetWindowLong func(w HWND, index int, newLong LONG) LONG

	SetWindowLongPtr func(
		w HWND, index int, newLong LONG_PTR) LONG_PTR

	SetWindowPlacement func(w HWND, wndpl *WINDOWPLACEMENT) BOOL

	SetWindowPos func(
		w HWND,
		insertAfter HWND,
		x, y, cx, cy int,
		flags UINT) BOOL

	SetWindowRgn func(w HWND, rgn HRGN, redraw BOOL) int

	SetWindowsHook func(filterType int, fn HOOKPROC) HHOOK

	SetWindowsHookEx func(
		hook int,
		fn HOOKPROC,
		mod HINSTANCE,
		threadId DWORD) HHOOK

	SetWindowText func(w HWND, s VString) BOOL

	SetWindowWord func(w HWND, index int, newWord WORD) WORD

	SetWinEventHook func(
		min, max DWORD,
		modWinEventProc HMODULE,
		fn WINEVENTPROC,
		process, thread DWORD,
		f WINEVENT_FLAGS) HWINEVENTHOOK

	ShowCaret func(w HWND) BOOL

	ShowCursor func(show BOOL) int

	ShowOwnedPopups func(w HWND, show BOOL) BOOL

	ShowScrollBar func(w HWND, bar int, show BOOL) BOOL

	ShowWindow func(w HWND, cmdShow int) BOOL

	ShowWindowAsync func(w HWND, cmdShow int) BOOL

	SubtractRect func(dst, src1, src2 *RECT) BOOL

	SwapMouseButton func(swap BOOL) BOOL

	SwitchDesktop func(desktop HDESK) BOOL

	SwitchToThisWindow func(w HWND, unknown BOOL)

	SystemParametersInfo func(
		action, nParam UINT, param *VOID, winIni UINT) BOOL

	TabbedTextOut func(
		h HDC,
		x, y int,
		s VString,
		count, tabPositions int,
		tabStopPositions *INT,
		tabOrigin int) LONG

	TileWindows func(
		parent HWND,
		how UINT,
		rect *RECT,
		nKids UINT,
		kids *HWND) WORD

	ToAscii func(
		virtKey, scanCode UINT,
		keyState *BYTE,
		char *WORD,
		flags UINT) int

	ToAsciiEx func(
		virtKey, scanCode UINT,
		keyState *BYTE,
		char *WORD,
		flags UINT,
		hkl HKL) int

	ToUnicode func(
		virtKey, scanCode UINT,
		keyState *BYTE,
		buff OWString,
		size int,
		flags UINT) int

	ToUnicodeEx func(
		virtKey, scanCode UINT,
		keyState *BYTE,
		buff WString,
		size int,
		flags UINT,
		h HKL) int

	TrackMouseEvent func(eventTrack *TRACKMOUSEEVENT) BOOL

	TrackPopupMenu func(
		m HMENU,
		flags UINT,
		x, y, reserved int,
		w HWND,
		rect *RECT) BOOL

	TrackPopupMenuEx func(
		HMENU, UINT, int, int, HWND, *TPMPARAMS) BOOL

	TranslateAccelerator func(
		w HWND, accTable HACCEL, msg *MSG) int

	TranslateMDISysAccel func(wndClient HWND, msg *MSG) BOOL

	TranslateMessage func(msg *MSG) BOOL

	UnhookWindowsHook func(code int, filterProc HOOKPROC) BOOL

	UnhookWindowsHookEx func(hk HHOOK) BOOL

	UnhookWinEvent func(winEventHook HWINEVENTHOOK) BOOL

	UnionRect func(dst, src1, src2 *RECT) BOOL

	UnloadKeyboardLayout func(hkl HKL) BOOL

	UnregisterClass func(
		className VString, instance HINSTANCE) BOOL

	UnregisterDeviceNotification func(handle HDEVNOTIFY) BOOL

	UnregisterHotKey func(w HWND, id int) BOOL

	UpdateLayeredWindow func(
		w HWND,
		dst HDC,
		ptDst *POINT,
		s *SIZE,
		src HDC,
		ptSrc *POINT,
		crKey COLORREF,
		blend *BLENDFUNCTION,
		flags DWORD) BOOL

	UpdateLayeredWindowIndirect func(
		w HWND, ulwi *UPDATELAYEREDWINDOWINFO) BOOL

	UpdateWindow func(w HWND) BOOL

	UserHandleGrantAccess func(
		user HANDLE, job HANDLE, grant BOOL) BOOL

	ValidateRect func(w HWND, rect *RECT) BOOL

	ValidateRgn func(w HWND, rgn HRGN) BOOL

	VkKeyScan func(ch Char) SHORT

	VkKeyScanEx func(ch Char, dwhkl HKL) SHORT

	WaitForInputIdle func(
		process HANDLE, milliseconds DWORD) DWORD

	WaitMessage func() BOOL

	WindowFromDC func(dc HDC) HWND

	WindowFromPoint func(point POINT) HWND

	WinHelp func(
		main HWND,
		help VString,
		command UINT,
		data ULONG_PTR) BOOL

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
var WinUserUnicodeApis = Apis{
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
	{"GetWindowLongW", &GetWindowLongPtr}, // NOTE(t): Not on xp32
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
	{"SetWindowLongW", &SetWindowLongPtr}, // NOTE(t): Not on xp32
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
	{"GetClassLongA", &GetClassLongPtr}, // NOTE(t): Not on xp32
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
	{"GetWindowLongA", &GetWindowLongPtr}, // NOTE(t): Not on xp32
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
	{"SetClassLongA", &SetClassLongPtr}, // NOTE(t): Not on xp32
	{"SetDlgItemTextA", &SetDlgItemText},
	{"SetMenuItemInfoA", &SetMenuItemInfo},
	{"SetPropA", &SetProp},
	{"SetUserObjectInformationA", &SetUserObjectInformation},
	{"SetWindowLongA", &SetWindowLong},
	{"SetWindowLongA", &SetWindowLongPtr}, // NOTE(t): Not on xp32
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
