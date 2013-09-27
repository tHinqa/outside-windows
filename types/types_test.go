package types

import "testing"

func TestWM(t *testing.T) {
	if WM_CLOSE != 0x0010 ||
		WM_SETCURSOR != 0x0020 ||
		WM_SETFONT != 0x30 ||
		WM_COMPACTING != 0x41 ||
		WM_INPUTLANGCHANGEREQUEST != 0x50 ||
		WM_SETICON != 0x80 ||
		WM_NCMOUSEMOVE != 0xA0 ||
		WM_KEYDOWN != 0x100 ||
		WM_INITDIALOG != 0x110 ||
		WM_MENUCHAR != 0x120 ||
		WM_CTLCOLORMSGBOX != 0x132 ||
		WM_PARENTNOTIFY != 0x210 ||
		WM_DEVICECHANGE != 0x219 ||
		WM_MDICREATE != 0x220 ||
		WM_MDISETMENU != 0x0230 ||
		WM_PALETTEISCHANGING != 0x0310 ||
		WM_THEMECHANGED != 0x031A ||
		false {
		t.Error("WINDOW_MESSAGE integrity lost")
	}
}
