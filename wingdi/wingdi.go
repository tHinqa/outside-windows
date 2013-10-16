// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package wingdi provides API definitions for accessing
//gdi32.dll and opengl32.dll.
package wingdi

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	_ "github.com/tHinqa/outside/win32/gdi32"
	_ "github.com/tHinqa/outside/win32/opengl32"
)

var (
	//GetBitmapBits is obsolete; use GetDIBits
	GetBitmapBits func(bit T.HBITMAP, c T.LONG, bits *T.VOID) (T.LONG, error)

	//GetCharWidth is obsolete; use GetCharWidth32
	GetCharWidth func(dc T.HDC, first, last T.UINT, buffer *T.INT) (T.BOOL, error)

	//SetBitmapBits is obsolete; use SetDIBits
	SetBitmapBits func(bm T.HBITMAP, cb T.DWORD, bits *T.VOID) (T.LONG, error)

	AddFontResource func(VString) (int, error)

	AnimatePalette func(
		p T.HPALETTE,
		startIndex T.UINT,
		entries T.UINT,
		pe *T.PALETTEENTRY) (T.BOOL, error)

	Arc func(
		dc T.HDC, x1, y1, x2, y2, x3, y3, x4, y4 int) (T.BOOL, error)

	BitBlt func(
		dc T.HDC,
		x, y, cx, cy int,
		src T.HDC,
		x1, y1 int,
		rop T.DWORD) (T.BOOL, error)

	CancelDC func(
		dc T.HDC) (T.BOOL, error)

	Chord func(
		dc T.HDC, x1, y1, x2, y2, x3, y3, x4, y4 int) (T.BOOL, error)

	ChoosePixelFormat func(T.HDC, *T.PIXELFORMATDESCRIPTOR) (int, error)

	CloseMetaFile func(T.HDC) (T.HMETAFILE, error)

	CombineRgn func(dst, src1, src2 T.HRGN, mode int) (int, error)

	CopyMetaFile func(T.HMETAFILE, VString) (T.HMETAFILE, error)

	CreateBitmap func(
		width, height int,
		planes, bitCount T.UINT,
		bits *T.VOID) (T.HBITMAP, error)

	CreateBitmapIndirect func(*T.BITMAP) (T.HBITMAP, error)

	CreateBrushIndirect func(*T.LOGBRUSH) (T.HBRUSH, error)

	CreateCompatibleBitmap func(dc T.HDC, cx, cy int) (T.HBITMAP, error)

	CreateDiscardableBitmap func(dc T.HDC, cx, cy int) (T.HBITMAP, error)

	CreateCompatibleDC func(dc T.HDC) (T.HDC, error)

	CreateDCA func(
		driver, device, port VString, dm *T.DEVMODEA) (T.HDC, error)

	CreateDCW func(
		driver, device, port VString, dm *T.DEVMODEW) (T.HDC, error)

	CreateDIBitmap func(
		dc T.HDC,
		bmih *T.BITMAPINFOHEADER,
		init T.DWORD,
		bits *T.VOID,
		bmi *T.BITMAPINFO,
		usage T.UINT) (T.HBITMAP, error)

	CreateDIBPatternBrush func(h T.HGLOBAL, usage T.UINT) (T.HBRUSH, error)

	CreateDIBPatternBrushPt func(
		packedDIB *T.VOID, usage T.UINT) (T.HBRUSH, error)

	CreateEllipticRgn func(x1, y1, x2, y2 int) (T.HRGN, error)

	CreateEllipticRgnIndirect func(*T.RECT) (T.HRGN, error)

	CreateFontIndirectA func(*T.LOGFONTA) (T.HFONT, error)

	CreateFontIndirectW func(*T.LOGFONTW) (T.HFONT, error)

	CreateFont func(
		height, width, escapement, orientation, weight int,
		italic, underline, strikeOut, charSet, outPrecision,
		clipPrecision, quality, pitchAndFamily T.DWORD,
		faceName VString) (T.HFONT, error)

	CreateHatchBrush func(
		hatch int,
		color T.COLORREF) (T.HBRUSH, error)

	CreateICA func(
		driver, device, port VString, dm *T.DEVMODEA) (T.HDC, error)

	CreateICW func(
		driver, device, port VString, dm *T.DEVMODEW) (T.HDC, error)

	CreateMetaFile func(file VString) (T.HDC, error)

	CreatePalette func(*T.LOGPALETTE) (T.HPALETTE, error)

	CreatePen func(style, width int, color T.COLORREF) (T.HPEN, error)

	CreatePenIndirect func(*T.LOGPEN) (T.HPEN, error)

	CreatePolyPolygonRgn func(
		ptl *T.POINT, c *T.INT, poly, mode int) (T.HRGN, error)

	CreatePatternBrush func(T.HBITMAP) (T.HBRUSH, error)

	CreateRectRgn func(x1, y1, x2, y2 int) (T.HRGN, error)

	CreateRectRgnIndirect func(*T.RECT) (T.HRGN, error)

	CreateRoundRectRgn func(x1, y1, x2, y2, w, h int) (T.HRGN, error)

	CreateScalableFontResource func(
		hidden T.DWORD,
		font, file, path VString) (T.BOOL, error)

	CreateSolidBrush func(T.COLORREF) (T.HBRUSH, error)

	DeleteDC func(T.HDC) (T.BOOL, error)

	DeleteMetaFile func(T.HMETAFILE) (T.BOOL, error)

	DeleteObject func(T.HGDIOBJ) (T.BOOL, error)

	DescribePixelFormat func(
		dc T.HDC,
		pixelFormat int,
		bytes T.UINT,
		pfd *T.PIXELFORMATDESCRIPTOR) (int, error)

	DeviceCapabilitiesA func(
		device, port VString,
		capability T.WORD,
		output OVString,
		dm *T.DEVMODEA) (int, error)

	DeviceCapabilitiesW func(
		device, port VString,
		capability T.WORD,
		output OVString,
		dm *T.DEVMODEW) (int, error)

	DrawEscape func(dc T.HDC, escape, cIn int, in T.AString) (int, error)

	Ellipse func(dc T.HDC, left, top, right, bottom int) (T.BOOL, error)

	EnumFontFamiliesExA func(
		dc T.HDC,
		lf *T.LOGFONTA,
		p T.FONTENUMPROCA,
		l T.LPARAM,
		flags T.DWORD) (int, error)

	EnumFontFamiliesExW func(
		dc T.HDC,
		lf *T.LOGFONTW,
		p T.FONTENUMPROCW,
		l T.LPARAM,
		flags T.DWORD) (int, error)

	EnumFontFamiliesA func(
		dc T.HDC,
		lf T.AString,
		p T.FONTENUMPROCA,
		l T.LPARAM) (int, error)

	EnumFontFamiliesW func(
		dc T.HDC,
		lf T.WString,
		p T.FONTENUMPROCW,
		l T.LPARAM) (int, error)

	EnumFontsA func(
		dc T.HDC,
		lf T.AString,
		p T.FONTENUMPROCA,
		l T.LPARAM) (int, error)

	EnumFontsW func(
		dc T.HDC,
		lf T.WString,
		p T.FONTENUMPROCW,
		l T.LPARAM) (int, error)

	EnumObjects func(
		dc T.HDC,
		typ int,
		p T.GOBJENUMPROC,
		l T.LPARAM) (int, error)

	EqualRgn func(rgn1, rgn2 T.HRGN) (T.BOOL, error)

	Escape func(
		dc T.HDC, escape, cIn int, in T.AString, out *T.VOID) (int, error)

	///////////////////////////////////////////////////////
	ExtEscape func(
		dc T.HDC,
		escape, input int,
		inData VString,
		output int,
		outData OVString) (int, error)

	ExcludeClipRect func(
		dc T.HDC, left, top, right, bottom int) (int, error)

	ExtCreateRegion func(
		x *T.XFORM, count T.DWORD, data *T.RGNDATA) (T.HRGN, error)

	ExtFloodFill func(dc T.HDC, x, y int, color T.COLORREF, typ T.UINT) (T.BOOL, error)

	FillRgn func(dc T.HDC, rgn T.HRGN, br T.HBRUSH) (T.BOOL, error)

	FloodFill func(dc T.HDC, x, y int, color T.COLORREF) (T.BOOL, error)

	FrameRgn func(
		dc T.HDC, rgn T.HRGN, br T.HBRUSH, w, h int) (T.BOOL, error)

	GetROP2 func(dc T.HDC) (int, error)

	GetAspectRatioFilterEx func(dc T.HDC, size *T.SIZE) (T.BOOL, error)

	GetBkColor func(dc T.HDC) (T.COLORREF, error)

	GetDCBrushColor func(dc T.HDC) (T.COLORREF, error)

	GetDCPenColor func(dc T.HDC) (T.COLORREF, error)

	GetBkMode func(dc T.HDC) (int, error)

	GetBitmapDimensionEx func(bit T.HBITMAP, size *T.SIZE) (T.BOOL, error)

	GetBoundsRect func(dc T.HDC, rect *T.RECT, flags T.UINT) (T.UINT, error)

	GetBrushOrgEx func(dc T.HDC, pt *T.POINT) (T.BOOL, error)

	GetCharWidth32 func(
		dc T.HDC, first, last T.UINT, buffer *T.INT) (T.BOOL, error)

	GetCharWidthFloat func(
		dc T.HDC, first, last T.UINT, buffer *T.FLOAT) (T.BOOL, error)

	GetCharABCWidths func(
		dc T.HDC, first, last T.UINT, abc *T.ABC) (T.BOOL, error)

	GetCharABCWidthsFloat func(
		dc T.HDC, first T.UINT, last T.UINT, abc *T.ABCFLOAT) (T.BOOL, error)

	GetClipBox func(dc T.HDC, rect *T.RECT) (int, error)

	GetClipRgn func(dc T.HDC, rgn T.HRGN) (int, error)

	GetMetaRgn func(dc T.HDC, rgn T.HRGN) (int, error)

	GetCurrentObject func(dc T.HDC, typ T.UINT) (T.HGDIOBJ, error)

	GetCurrentPositionEx func(dc T.HDC, pt *T.POINT) (T.BOOL, error)

	GetDeviceCaps func(dc T.HDC, index int) (int, error)

	GetDIBits func(
		dc T.HDC,
		bm T.HBITMAP,
		start, lines T.UINT,
		vBits *T.VOID,
		bmi *T.BITMAPINFO,
		usage T.UINT) (int, error)

	GetFontData func(
		dc T.HDC,
		Table, Offset T.DWORD,
		Buffer *T.VOID,
		cBuffer T.DWORD) (T.DWORD, error)

	GetGlyphOutline func(
		dc T.HDC,
		Char T.UINT,
		Format T.UINT,
		gm *T.GLYPHMETRICS,
		cBuffer T.DWORD,
		Buffer *T.VOID,
		mat2 *T.MAT2) (T.DWORD, error)

	GetGraphicsMode func(dc T.HDC) (int, error)

	GetMapMode func(dc T.HDC) (int, error)

	GetMetaFileBitsEx func(
		MF T.HMETAFILE,
		cBuffer T.UINT,
		Data *T.VOID) (T.UINT, error)

	GetMetaFile func(Name VString) (T.HMETAFILE, error)

	GetNearestColor func(dc T.HDC, color T.COLORREF) (T.COLORREF, error)

	GetNearestPaletteIndex func(
		h T.HPALETTE, color T.COLORREF) (T.UINT, error)

	GetObjectType func(h T.HGDIOBJ) (T.DWORD, error)

	GetOutlineTextMetricsA func(
		dc T.HDC,
		cCopy T.UINT,
		otm *T.OUTLINETEXTMETRICA) (T.UINT, error)

	GetOutlineTextMetricsW func(
		dc T.HDC,
		cCopy T.UINT,
		otm *T.OUTLINETEXTMETRICW) (T.UINT, error)

	GetPaletteEntries func(
		pal T.HPALETTE,
		start, entries T.UINT,
		palEntries *T.PALETTEENTRY) (T.UINT, error)

	GetPixel func(
		dc T.HDC,
		x, y int) (T.COLORREF, error)

	GetPixelFormat func(dc T.HDC) (int, error)

	GetPolyFillMode func(dc T.HDC) (int, error)

	GetRasterizerCaps func(
		raststat *T.RASTERIZER_STATUS,
		cBytes T.UINT) (T.BOOL, error)

	GetRandomRgn func(dc T.HDC, rgn T.HRGN, i T.INT) (int, error)

	GetRegionData func(
		rgn T.HRGN, count T.DWORD, rgnData *T.RGNDATA) (T.DWORD, error)

	GetRgnBox func(rgn T.HRGN, rc *T.RECT) (int, error)

	GetStockObject func(i int) (T.HGDIOBJ, error)

	GetStretchBltMode func(dc T.HDC) (int, error)

	GetSystemPaletteEntries func(
		dc T.HDC,
		start, entries T.UINT,
		palEntries *T.PALETTEENTRY) (T.UINT, error)

	GetSystemPaletteUse func(hdc T.HDC) (T.UINT, error)

	GetTextCharacterExtra func(hdc T.HDC) (int, error)

	GetTextAlign func(hdc T.HDC) (T.UINT, error)

	GetTextColor func(hdc T.HDC) (T.COLORREF, error)

	GetTextExtentPoint func(
		dc T.HDC, s VString, c int, sz *T.SIZE) (T.BOOL, error)

	GetTextExtentPoint32 func(
		dc T.HDC, s VString, c int, sizl *T.SIZE) (T.BOOL, error)

	GetTextExtentExPoint func(
		dc T.HDC,
		s VString,
		cStr, maxExtent int,
		fit, dx *T.INT,
		size *T.SIZE) (T.BOOL, error)

	GetTextCharset func(dc T.HDC) (int, error)

	GetTextCharsetInfo func(
		dc T.HDC,
		sig *T.FONTSIGNATURE,
		flags T.DWORD) (int, error)

	TranslateCharsetInfo func(
		Src *T.DWORD,
		Cs *T.CHARSETINFO,
		Flags T.DWORD) (T.BOOL, error)

	GetFontLanguageInfo func(dc T.HDC) (T.DWORD, error)

	GetCharacterPlacementA func(
		dc T.HDC,
		String T.AString,
		Count,
		MexExtent int,
		Results *T.GCP_RESULTS,
		Flags T.DWORD) (T.DWORD, error)

	GetCharacterPlacementW func(
		dc T.HDC,
		String T.WString,
		Count,
		MexExtent int,
		Results *T.GCP_RESULTS,
		Flags T.DWORD) (T.DWORD, error)

	GetFontUnicodeRanges func(dc T.HDC, gs *T.GLYPHSET) (T.DWORD, error)

	GetGlyphIndices func(
		dc T.HDC,
		str VString,
		c int,
		gi *T.WORD,
		l T.DWORD) (T.DWORD, error)

	GetTextExtentPointI func(
		dc T.HDC, giIn *T.WORD, gi int, size *T.SIZE) (T.BOOL, error)

	GetTextExtentExPointI func(
		hdc T.HDC,
		lpwszString *T.WORD,
		cwchString int,
		nMaxExtent int,
		lpnFit *T.INT,
		lpnDx *T.INT,
		Size *T.SIZE) (T.BOOL, error)

	GetCharWidthI func(
		dc T.HDC,
		giFirst T.UINT,
		cgi T.UINT,
		gi *T.WORD,
		Widths *T.INT) (T.BOOL, error)

	GetCharABCWidthsI func(
		dc T.HDC,
		giFirst T.UINT,
		cgi T.UINT,
		gi *T.WORD,
		abc *T.ABC) (T.BOOL, error)

	AddFontResourceEx func(
		name VString, l T.DWORD, res *T.VOID) (int, error)

	RemoveFontResourceEx func(
		name VString, l T.DWORD, dv *T.VOID) (T.BOOL, error)

	AddFontMemResourceEx func(
		fileView *T.VOID,
		size T.DWORD,
		_ *T.VOID,
		numFonts *T.DWORD) (T.HANDLE, error)

	RemoveFontMemResourceEx func(h T.HANDLE) (T.BOOL, error)

	CreateFontIndirectExA func(*T.ENUMLOGFONTEXDVA) (T.HFONT, error)

	CreateFontIndirectExW func(*T.ENUMLOGFONTEXDVW) (T.HFONT, error)

	GetViewportExtEx func(dc T.HDC, size *T.SIZE) (T.BOOL, error)

	GetViewportOrgEx func(dc T.HDC, point *T.POINT) (T.BOOL, error)

	GetWindowExtEx func(dc T.HDC, size *T.SIZE) (T.BOOL, error)

	GetWindowOrgEx func(dc T.HDC, point *T.POINT) (T.BOOL, error)

	IntersectClipRect func(
		dc T.HDC, left, top, right, bottom int) (int, error)

	InvertRgn func(dc T.HDC, rgn T.HRGN) (T.BOOL, error)

	LineDDA func(
		xStart, yStart, xEnd, yEnd int,
		Proc T.LINEDDAPROC,
		data T.LPARAM) (T.BOOL, error)

	LineTo func(hdc T.HDC, x, y int) (T.BOOL, error)

	MaskBlt func(
		Dest T.HDC,
		xDest, yDest, width, height int,
		Src T.HDC,
		xSrc, ySrc int,
		Mask T.HBITMAP,
		xMask, yMask int,
		rop T.DWORD) (T.BOOL, error)

	PlgBlt func(
		Dest T.HDC,
		Point *T.POINT,
		Src T.HDC,
		xSrc, ySrc, width, height int,
		mask T.HBITMAP,
		xMask, yMask int) (T.BOOL, error)

	OffsetClipRgn func(dc T.HDC, x, y int) (int, error)

	OffsetRgn func(rgn T.HRGN, x, y int) (int, error)

	PatBlt func(dc T.HDC, x, y, w, h int, rop T.DWORD) (T.BOOL, error)

	Pie func(
		hdc T.HDC,
		left, top, right, bottom,
		xr1, yr1, xr2, yr2 int) (T.BOOL, error)

	PlayMetaFile func(dc T.HDC, mf T.HMETAFILE) (T.BOOL, error)

	PaintRgn func(dc T.HDC, rgn T.HRGN) (T.BOOL, error)

	PolyPolygon func(
		dc T.HDC, pt *T.POINT, psz *T.INT, sz int) (T.BOOL, error)

	PtInRegion func(rgn T.HRGN, x int, y int) (T.BOOL, error)

	PtVisible func(dc T.HDC, x int, y int) (T.BOOL, error)

	RectInRegion func(rgn T.HRGN, rect *T.RECT) (T.BOOL, error)

	RectVisible func(dc T.HDC, rect *T.RECT) (T.BOOL, error)

	Rectangle func(dc T.HDC, left, top, right, bottom int) (T.BOOL, error)

	RestoreDC func(dc T.HDC, savedDC int) (T.BOOL, error)

	ResetDCA func(dc T.HDC, dm *T.DEVMODEA) (T.HDC, error)

	ResetDCW func(dc T.HDC, dm *T.DEVMODEW) (T.HDC, error)

	RealizePalette func(dc T.HDC) (T.UINT, error)

	RemoveFontResource func(FileName VString) (T.BOOL, error)

	RoundRect func(
		dc T.HDC,
		left, top, right, bottom, width, height int) (T.BOOL, error)

	ResizePalette func(pal T.HPALETTE, n T.UINT) (T.BOOL, error)

	SaveDC func(dc T.HDC) (int, error)

	SelectClipRgn func(dc T.HDC, rgn T.HRGN) (int, error)

	ExtSelectClipRgn func(dc T.HDC, rgn T.HRGN, mode int) (int, error)

	SetMetaRgn func(dc T.HDC) (int, error)

	SelectObject func(dc T.HDC, h T.HGDIOBJ) (T.HGDIOBJ, error)

	SelectPalette func(
		dc T.HDC, pal T.HPALETTE, forceBkgd T.BOOL) (T.HPALETTE, error)

	SetBkColor func(dc T.HDC, color T.COLORREF) (T.COLORREF, error)

	SetDCBrushColor func(dc T.HDC, color T.COLORREF) (T.COLORREF, error)

	SetDCPenColor func(dc T.HDC, color T.COLORREF) (T.COLORREF, error)

	SetBkMode func(dc T.HDC, mode int) (int, error)

	SetBoundsRect func(
		dc T.HDC, rect *T.RECT, flags T.UINT) (T.UINT, error)

	SetDIBits func(
		dc T.HDC,
		bm T.HBITMAP,
		start, lines T.UINT,
		Bits *T.VOID,
		bmi *T.BITMAPINFO,
		ColorUse T.UINT) (int, error)

	SetDIBitsToDevice func(
		dc T.HDC,
		xDest, yDest int,
		w, h T.DWORD,
		xSrc, ySrc int,
		startScan, lines T.UINT,
		Bits *T.VOID,
		bmi *T.BITMAPINFO,
		ColorUse T.UINT) (int, error)

	SetMapperFlags func(dc T.HDC, flags T.DWORD) (T.DWORD, error)

	SetGraphicsMode func(dc T.HDC, mode int) (int, error)

	SetMapMode func(dc T.HDC, mode int) (int, error)

	SetLayout func(dc T.HDC, l T.DWORD) (T.DWORD, error)

	GetLayout func(dc T.HDC) (T.DWORD, error)

	SetMetaFileBitsEx func(
		buffer T.UINT, data *T.BYTE) (T.HMETAFILE, error)

	SetPaletteEntries func(
		pal T.HPALETTE,
		start, entries T.UINT,
		PalEntries *T.PALETTEENTRY) (T.UINT, error)

	SetPixel func(
		dc T.HDC, x, y int, color T.COLORREF) (T.COLORREF, error)

	SetPixelV func(
		hdc T.HDC, x, y int, color T.COLORREF) (T.BOOL, error)

	SetPixelFormat func(
		dc T.HDC,
		format int,
		pfd *T.PIXELFORMATDESCRIPTOR) (T.BOOL, error)

	SetPolyFillMode func(dc T.HDC, mode int) (int, error)

	StretchBlt func(
		Dest T.HDC,
		xDest, yDest, wDest, hDest int,
		Src T.HDC,
		xSrc, ySrc, wSrc, hSrc int,
		rop T.DWORD) (T.BOOL, error)

	SetRectRgn func(
		rgn T.HRGN,
		left, top, right, bottom int) (T.BOOL, error)

	StretchDIBits func(
		hdc T.HDC,
		xDest, yDest,
		destWidth, destHeight,
		xSrc, ySrc,
		srcWidth, srcHeight int,
		bits *T.VOID,
		bmi *T.BITMAPINFO,
		usage T.UINT,
		rop T.DWORD) (int, error)

	SetROP2 func(dc T.HDC, rop2 int) (int, error)

	SetStretchBltMode func(dc T.HDC, mode int) (int, error)

	SetSystemPaletteUse func(dc T.HDC, use T.UINT) (T.UINT, error)

	SetTextCharacterExtra func(dc T.HDC, extra int) (int, error)

	SetTextColor func(dc T.HDC, color T.COLORREF) (T.COLORREF, error)

	SetTextAlign func(dc T.HDC, align T.UINT) (T.UINT, error)

	SetTextJustification func(dc T.HDC, extra, count int) (T.BOOL, error)

	UpdateColors func(hdc T.HDC) (T.BOOL, error)

	AlphaBlend func(
		dest T.HDC,
		xoriginDest, yoriginDest,
		wDest, hDest int,
		src T.HDC,
		xoriginSrc, yoriginSrc,
		wSrc, hSrc int,
		ftn T.BLENDFUNCTION) (T.BOOL, error)

	TransparentBlt func(
		Dest T.HDC,
		xoriginDest, yoriginDest,
		wDest, hDest int,
		Src T.HDC,
		xoriginSrc, yoriginSrc,
		wSrc, hSrc int,
		crTransparent T.UINT) (T.BOOL, error)

	GradientFill func(
		dc T.HDC,
		Vertex *T.TRIVERTEX,
		nVertex T.ULONG,
		Mesh *T.VOID,
		nMesh T.ULONG,
		Mode T.ULONG) (T.BOOL, error)

	PlayMetaFileRecord func(
		dc T.HDC,
		HandleTable *T.HANDLETABLE,
		MR *T.METARECORD,
		Objs T.UINT) (T.BOOL, error)

	EnumMetaFile func(
		dc T.HDC,
		mf T.HMETAFILE,
		proc T.MFENUMPROC,
		param T.LPARAM) (T.BOOL, error)

	CloseEnhMetaFile func(
		dc T.HDC) (T.HENHMETAFILE, error)

	CopyEnhMetaFile func(
		Enh T.HENHMETAFILE,
		FileName VString) (T.HENHMETAFILE, error)

	CreateEnhMetaFile func(
		dc T.HDC,
		Filename VString,
		rc *T.RECT,
		Desc VString) (T.HDC, error)

	DeleteEnhMetaFile func(mf T.HENHMETAFILE) (T.BOOL, error)

	EnumEnhMetaFile func(
		dc T.HDC,
		mf T.HENHMETAFILE,
		proc T.ENHMFENUMPROC,
		param *T.VOID,
		Rect *T.RECT) (T.BOOL, error)

	GetEnhMetaFile func(Name VString) (T.HENHMETAFILE, error)

	GetEnhMetaFileBits func(
		emf T.HENHMETAFILE, size T.UINT, data *T.BYTE) (T.UINT, error)

	GetEnhMetaFileDescription func(
		emf T.HENHMETAFILE,
		cBuffer T.UINT,
		description OVString) (T.UINT, error)

	GetEnhMetaFileHeader func(
		emf T.HENHMETAFILE,
		size T.UINT,
		enhMetaHeader *T.ENHMETAHEADER) (T.UINT, error)

	GetEnhMetaFilePaletteEntries func(
		emf T.HENHMETAFILE,
		NumEntries T.UINT,
		PaletteEntries *T.PALETTEENTRY) (T.UINT, error)

	GetEnhMetaFilePixelFormat func(
		emf T.HENHMETAFILE,
		cBuffer T.UINT,
		pfd *T.PIXELFORMATDESCRIPTOR) (T.UINT, error)

	GetWinMetaFileBits func(
		emf T.HENHMETAFILE,
		cData16 T.UINT,
		Data16 *T.BYTE,
		MapMode T.INT,
		Ref T.HDC) (T.UINT, error)

	PlayEnhMetaFile func(
		dc T.HDC,
		mf T.HENHMETAFILE,
		rect *T.RECT) (T.BOOL, error)

	PlayEnhMetaFileRecord func(
		dc T.HDC,
		ht *T.HANDLETABLE,
		mr *T.ENHMETARECORD,
		cht T.UINT) (T.BOOL, error)

	SetEnhMetaFileBits func(
		nSize T.UINT,
		pb *T.BYTE) (T.HENHMETAFILE, error)

	SetWinMetaFileBits func(
		Size T.UINT,
		Meta16Data *T.BYTE,
		Ref T.HDC,
		MFP *T.METAFILEPICT) (T.HENHMETAFILE, error)

	GdiComment func(
		dc T.HDC,
		Size T.UINT,
		Data *T.BYTE) (T.BOOL, error)

	GetTextMetricsA func(dc T.HDC, tm *T.TEXTMETRICA) (T.BOOL, error)

	GetTextMetricsW func(dc T.HDC, tm *T.TEXTMETRICW) (T.BOOL, error)

	AngleArc func(
		dc T.HDC,
		x, y int,
		r T.DWORD,
		StartAngle T.FLOAT,
		SweepAngle T.FLOAT) (T.BOOL, error)

	PolyPolyline func(
		dc T.HDC, pt *T.POINT, psz *T.DWORD, sz T.DWORD) (T.BOOL, error)

	GetWorldTransform func(dc T.HDC, xf *T.XFORM) (T.BOOL, error)

	SetWorldTransform func(dc T.HDC, xf *T.XFORM) (T.BOOL, error)

	ModifyWorldTransform func(
		dc T.HDC, xf *T.XFORM, mode T.DWORD) (T.BOOL, error)

	CombineTransform func(
		out *T.XFORM, xf1, xf2 *T.XFORM) (T.BOOL, error)

	CreateDIBSection func(
		dc T.HDC,
		bmi *T.BITMAPINFO,
		usage T.UINT,
		Bits **T.VOID,
		Section T.HANDLE,
		offset T.DWORD) (T.HBITMAP, error)

	GetDIBColorTable func(
		dc T.HDC,
		Start T.UINT,
		Entries T.UINT,
		rgbq *T.RGBQUAD) (T.UINT, error)

	SetDIBColorTable func(
		dc T.HDC,
		Start T.UINT,
		Entries T.UINT,
		rgbq *T.RGBQUAD) (T.UINT, error)

	SetColorAdjustment func(
		dc T.HDC, ca *T.COLORADJUSTMENT) (T.BOOL, error)

	GetColorAdjustment func(
		dc T.HDC, ca *T.COLORADJUSTMENT) (T.BOOL, error)

	CreateHalftonePalette func(dc T.HDC) (T.HPALETTE, error)

	StartDoc func(dc T.HDC, di *T.DOCINFO) (int, error)

	EndDoc func(dc T.HDC) (int, error)

	StartPage func(dc T.HDC) (int, error)

	EndPage func(dc T.HDC) (int, error)

	AbortDoc func(dc T.HDC) (int, error)

	SetAbortProc func(dc T.HDC, proc T.ABORTPROC) (int, error)

	AbortPath func(hdc T.HDC) (T.BOOL, error)

	ArcTo func(
		dc T.HDC,
		left, top, right, bottom,
		xr1, yr1, xr2, yr2 int) (T.BOOL, error)

	BeginPath func(dc T.HDC) (T.BOOL, error)

	CloseFigure func(dc T.HDC) (T.BOOL, error)

	EndPath func(dc T.HDC) (T.BOOL, error)

	FillPath func(dc T.HDC) (T.BOOL, error)

	FlattenPath func(dc T.HDC) (T.BOOL, error)

	GetPath func(dc T.HDC, pt *T.POINT, t *T.BYTE, cpt int) (int, error)

	PathToRegion func(hdc T.HDC) (T.HRGN, error)

	PolyDraw func(dc T.HDC, pt *T.POINT, t *T.BYTE, cpt int) (T.BOOL, error)

	SelectClipPath func(dc T.HDC, mode int) (T.BOOL, error)

	SetArcDirection func(dc T.HDC, dir int) (int, error)

	SetMiterLimit func(dc T.HDC, limit T.FLOAT, old *T.FLOAT) (T.BOOL, error)

	StrokeAndFillPath func(dc T.HDC) (T.BOOL, error)

	StrokePath func(dc T.HDC) (T.BOOL, error)

	WidenPath func(dc T.HDC) (T.BOOL, error)

	ExtCreatePen func(
		PenStyle T.DWORD,
		Width T.DWORD,
		LBrush *T.LOGBRUSH,
		cStyle T.DWORD,
		Style *T.DWORD) (T.HPEN, error)

	GetMiterLimit func(dc T.HDC, limit *T.FLOAT) (T.BOOL, error)

	GetArcDirection func(dc T.HDC) (int, error)

	GetObject func(h T.HANDLE, c int, v *T.VOID) (int, error)

	MoveToEx func(dc T.HDC, x, y int, pt *T.POINT) (T.BOOL, error)

	TextOut func(
		dc T.HDC, x, y int, lpString VString, c int) (T.BOOL, error)

	ExtTextOut func(
		dc T.HDC,
		x, y int,
		options T.UINT,
		rect *T.RECT,
		String VString,
		c T.UINT,
		Dx *T.INT) (T.BOOL, error)

	PolyTextOut func(
		dc T.HDC, pt *T.POLYTEXT, nstrings int) (T.BOOL, error)

	CreatePolygonRgn func(
		ptl *T.POINT,
		cPoint int,
		Mode int) (T.HRGN, error)

	DPtoLP func(dc T.HDC, pt *T.POINT, c int) (T.BOOL, error)

	LPtoDP func(dc T.HDC, pt *T.POINT, c int) (T.BOOL, error)

	Polygon func(dc T.HDC, pt *T.POINT, cpt int) (T.BOOL, error)

	Polyline func(dc T.HDC, pt *T.POINT, cpt int) (T.BOOL, error)

	PolyBezier func(dc T.HDC, pt *T.POINT, cpt T.DWORD) (T.BOOL, error)

	PolyBezierTo func(dc T.HDC, pt *T.POINT, cpt T.DWORD) (T.BOOL, error)

	PolylineTo func(dc T.HDC, pt *T.POINT, cpt T.DWORD) (T.BOOL, error)

	SetViewportExtEx func(dc T.HDC, x, y int, sz *T.SIZE) (T.BOOL, error)

	SetViewportOrgEx func(dc T.HDC, x, y int, pt *T.POINT) (T.BOOL, error)

	SetWindowExtEx func(dc T.HDC, x, y int, sz *T.SIZE) (T.BOOL, error)

	SetWindowOrgEx func(dc T.HDC, x, y int, pt *T.POINT) (T.BOOL, error)

	OffsetViewportOrgEx func(dc T.HDC, x, y int, pt *T.POINT) (T.BOOL, error)

	OffsetWindowOrgEx func(dc T.HDC, x, y int, pt *T.POINT) (T.BOOL, error)

	ScaleViewportExtEx func(
		hdc T.HDC, xn, xd, yn, yd int, sz *T.SIZE) (T.BOOL, error)

	ScaleWindowExtEx func(
		dc T.HDC, xn, xd, yn, yd int, sz *T.SIZE) (T.BOOL, error)

	SetBitmapDimensionEx func(
		bm T.HBITMAP, w, h int, sz *T.SIZE) (T.BOOL, error)

	SetBrushOrgEx func(dc T.HDC, x, y int, pt *T.POINT) (T.BOOL, error)

	GetTextFace func(dc T.HDC, c int, name VString) (int, error)

	GetKerningPairs func(
		dc T.HDC, pairs T.DWORD, kernPair *T.KERNINGPAIR) (T.DWORD, error)

	GetDCOrgEx func(dc T.HDC, pt *T.POINT) (T.BOOL, error)

	FixBrushOrgEx func(
		dc T.HDC, x, y int, ptl *T.POINT) (T.BOOL, error)

	UnrealizeObject func(h T.HGDIOBJ) (T.BOOL, error)

	GdiFlush func() (T.BOOL, error)

	GdiSetBatchLimit func(dw T.DWORD) (T.DWORD, error)

	GdiGetBatchLimit func() (T.DWORD, error)

	SetICMMode func(dc T.HDC, mode int) (int, error)

	CheckColorsInGamut func(
		dc T.HDC,
		RGBTriple *T.VOID,
		Buffer *T.VOID,
		Count T.DWORD) (T.BOOL, error)

	GetColorSpace func(dc T.HDC) (T.HCOLORSPACE, error)

	GetLogColorSpaceA func(
		ColorSpace T.HCOLORSPACE,
		Buffer *T.LOGCOLORSPACEA,
		Size T.DWORD) (T.BOOL, error)

	GetLogColorSpaceW func(
		ColorSpace T.HCOLORSPACE,
		Buffer *T.LOGCOLORSPACEW,
		Size T.DWORD) (T.BOOL, error)

	CreateColorSpaceA func(
		lcs *T.LOGCOLORSPACEA) (T.HCOLORSPACE, error)

	CreateColorSpaceW func(
		lcs *T.LOGCOLORSPACEW) (T.HCOLORSPACE, error)

	SetColorSpace func(
		dc T.HDC, cs T.HCOLORSPACE) (T.HCOLORSPACE, error)

	DeleteColorSpace func(cs T.HCOLORSPACE) (T.BOOL, error)

	GetICMProfile func(
		dc T.HDC, bufSize *T.DWORD, filename VString) (T.BOOL, error)

	SetICMProfile func(dc T.HDC, fileName VString) (T.BOOL, error)

	GetDeviceGammaRamp func(dc T.HDC, ramp *T.VOID) (T.BOOL, error)

	SetDeviceGammaRamp func(dc T.HDC, ramp *T.VOID) (T.BOOL, error)

	ColorMatchToTarget func(
		dc T.HDC, target T.HDC, action T.DWORD) (T.BOOL, error)

	EnumICMProfilesA func(
		dc T.HDC, proc T.ICMENUMPROCA, param T.LPARAM) (int, error)

	EnumICMProfilesW func(
		dc T.HDC, proc T.ICMENUMPROCW, param T.LPARAM) (int, error)

	UpdateICMRegKey func(
		reserved T.DWORD,
		cmid, fileName VString,
		command T.UINT) (T.BOOL, error)

	ColorCorrectPalette func(
		dc T.HDC,
		pal T.HPALETTE,
		first T.DWORD,
		num T.DWORD) (T.BOOL, error)

	WglCopyContext func(T.HGLRC, T.HGLRC, T.UINT) (T.BOOL, error)

	WglCreateContext func(T.HDC) (T.HGLRC, error)

	WglCreateLayerContext func(T.HDC, int) (T.HGLRC, error)

	WglDeleteContext func(T.HGLRC) (T.BOOL, error)

	WglGetCurrentContext func(T.VOID) (T.HGLRC, error)

	WglGetCurrentDC func(T.VOID) (T.HDC, error)

	WglGetProcAddress func(T.AString) (T.PROC, error)

	WglMakeCurrent func(T.HDC, T.HGLRC) (T.BOOL, error)

	WglShareLists func(T.HGLRC, T.HGLRC) (T.BOOL, error)

	WglUseFontBitmaps func(T.HDC, T.DWORD, T.DWORD, T.DWORD) (T.BOOL, error)

	SwapBuffers func(T.HDC) (T.BOOL, error)

	WglUseFontOutlines func(
		T.HDC, T.DWORD, T.DWORD, T.DWORD,
		T.FLOAT, T.FLOAT, int, *T.GLYPHMETRICSFLOAT) (T.BOOL, error)

	WglDescribeLayerPlane func(
		T.HDC, int, int, T.UINT, *T.LAYERPLANEDESCRIPTOR) (T.BOOL, error)

	WglSetLayerPaletteEntries func(
		T.HDC, int, int, int, *T.COLORREF) (int, error)

	WglGetLayerPaletteEntries func(
		T.HDC, int, int, int, *T.COLORREF) (int, error)

	WglRealizeLayerPalette func(T.HDC, int, T.BOOL) (T.BOOL, error)

	WglSwapLayerBuffers func(T.HDC, T.WGL_SWAP_FLAG) (T.BOOL, error)

	WglSwapMultipleBuffers func(T.UINT, *T.WGLSWAP) (T.DWORD, error)
)

/*
TODO(t): msimg32
{"AlphaBlend", &AlphaBlend},
{"GradientFill", &GradientFill},
{"TransparentBlt", &TransparentBlt},

TODO(t): winspool
{"DeviceCapabilities", &DeviceCapabilities},
*/

var WinGdiApis = outside.Apis{
	{"AbortDoc", &AbortDoc},
	{"AbortPath", &AbortPath},
	{"AddFontMemResourceEx", &AddFontMemResourceEx},
	{"AngleArc", &AngleArc},
	{"AnimatePalette", &AnimatePalette},
	{"Arc", &Arc},
	{"ArcTo", &ArcTo},
	{"BeginPath", &BeginPath},
	{"BitBlt", &BitBlt},
	{"CancelDC", &CancelDC},
	{"CheckColorsInGamut", &CheckColorsInGamut},
	{"ChoosePixelFormat", &ChoosePixelFormat},
	{"Chord", &Chord},
	{"CloseEnhMetaFile", &CloseEnhMetaFile},
	{"CloseFigure", &CloseFigure},
	{"CloseMetaFile", &CloseMetaFile},
	{"ColorCorrectPalette", &ColorCorrectPalette},
	{"ColorMatchToTarget", &ColorMatchToTarget},
	{"CombineRgn", &CombineRgn},
	{"CombineTransform", &CombineTransform},
	{"CreateBitmap", &CreateBitmap},
	{"CreateBitmapIndirect", &CreateBitmapIndirect},
	{"CreateBrushIndirect", &CreateBrushIndirect},
	{"CreateCompatibleBitmap", &CreateCompatibleBitmap},
	{"CreateCompatibleDC", &CreateCompatibleDC},
	{"CreateDIBitmap", &CreateDIBitmap},
	{"CreateDIBPatternBrush", &CreateDIBPatternBrush},
	{"CreateDIBPatternBrushPt", &CreateDIBPatternBrushPt},
	{"CreateDIBSection", &CreateDIBSection},
	{"CreateDiscardableBitmap", &CreateDiscardableBitmap},
	{"CreateEllipticRgn", &CreateEllipticRgn},
	{"CreateEllipticRgnIndirect", &CreateEllipticRgnIndirect},
	{"CreateHalftonePalette", &CreateHalftonePalette},
	{"CreateHatchBrush", &CreateHatchBrush},
	{"CreatePalette", &CreatePalette},
	{"CreatePatternBrush", &CreatePatternBrush},
	{"CreatePen", &CreatePen},
	{"CreatePenIndirect", &CreatePenIndirect},
	{"CreatePolygonRgn", &CreatePolygonRgn},
	{"CreatePolyPolygonRgn", &CreatePolyPolygonRgn},
	{"CreateRectRgn", &CreateRectRgn},
	{"CreateRectRgnIndirect", &CreateRectRgnIndirect},
	{"CreateRoundRectRgn", &CreateRoundRectRgn},
	{"CreateSolidBrush", &CreateSolidBrush},
	{"DeleteColorSpace", &DeleteColorSpace},
	{"DeleteDC", &DeleteDC},
	{"DeleteEnhMetaFile", &DeleteEnhMetaFile},
	{"DeleteMetaFile", &DeleteMetaFile},
	{"DeleteObject", &DeleteObject},
	{"DescribePixelFormat", &DescribePixelFormat},
	{"DPtoLP", &DPtoLP},
	{"DrawEscape", &DrawEscape},
	{"Ellipse", &Ellipse},
	{"EndDoc", &EndDoc},
	{"EndPage", &EndPage},
	{"EndPath", &EndPath},
	{"EnumEnhMetaFile", &EnumEnhMetaFile},
	{"EnumMetaFile", &EnumMetaFile},
	{"EnumObjects", &EnumObjects},
	{"EqualRgn", &EqualRgn},
	{"Escape", &Escape},
	{"ExcludeClipRect", &ExcludeClipRect},
	{"ExtCreatePen", &ExtCreatePen},
	{"ExtCreateRegion", &ExtCreateRegion},
	{"ExtEscape", &ExtEscape},
	{"ExtFloodFill", &ExtFloodFill},
	{"ExtSelectClipRgn", &ExtSelectClipRgn},
	{"FillPath", &FillPath},
	{"FillRgn", &FillRgn},
	{"FixBrushOrgEx", &FixBrushOrgEx},
	{"FlattenPath", &FlattenPath},
	{"FloodFill", &FloodFill},
	{"FrameRgn", &FrameRgn},
	{"GdiComment", &GdiComment},
	{"GdiFlush", &GdiFlush},
	{"GdiGetBatchLimit", &GdiGetBatchLimit},
	{"GdiSetBatchLimit", &GdiSetBatchLimit},
	{"GetArcDirection", &GetArcDirection},
	{"GetAspectRatioFilterEx", &GetAspectRatioFilterEx},
	{"GetBitmapBits", &GetBitmapBits},
	{"GetBitmapDimensionEx", &GetBitmapDimensionEx},
	{"GetBkColor", &GetBkColor},
	{"GetBkMode", &GetBkMode},
	{"GetBoundsRect", &GetBoundsRect},
	{"GetBrushOrgEx", &GetBrushOrgEx},
	{"GetCharABCWidthsI", &GetCharABCWidthsI},
	{"GetCharWidthI", &GetCharWidthI},
	{"GetClipBox", &GetClipBox},
	{"GetClipRgn", &GetClipRgn},
	{"GetColorAdjustment", &GetColorAdjustment},
	{"GetColorSpace", &GetColorSpace},
	{"GetCurrentObject", &GetCurrentObject},
	{"GetCurrentPositionEx", &GetCurrentPositionEx},
	{"GetDCBrushColor", &GetDCBrushColor},
	{"GetDCOrgEx", &GetDCOrgEx},
	{"GetDCPenColor", &GetDCPenColor},
	{"GetDeviceCaps", &GetDeviceCaps},
	{"GetDeviceGammaRamp", &GetDeviceGammaRamp},
	{"GetDIBColorTable", &GetDIBColorTable},
	{"GetDIBits", &GetDIBits},
	{"GetEnhMetaFileBits", &GetEnhMetaFileBits},
	{"GetEnhMetaFileHeader", &GetEnhMetaFileHeader},
	{"GetEnhMetaFilePaletteEntries", &GetEnhMetaFilePaletteEntries},
	{"GetEnhMetaFilePixelFormat", &GetEnhMetaFilePixelFormat},
	{"GetFontData", &GetFontData},
	{"GetFontLanguageInfo", &GetFontLanguageInfo},
	{"GetFontUnicodeRanges", &GetFontUnicodeRanges},
	{"GetGlyphOutline", &GetGlyphOutline},
	{"GetGraphicsMode", &GetGraphicsMode},
	{"GetKerningPairs", &GetKerningPairs},
	{"GetLayout", &GetLayout},
	{"GetMapMode", &GetMapMode},
	{"GetMetaFileBitsEx", &GetMetaFileBitsEx},
	{"GetMetaRgn", &GetMetaRgn},
	{"GetMiterLimit", &GetMiterLimit},
	{"GetNearestColor", &GetNearestColor},
	{"GetNearestPaletteIndex", &GetNearestPaletteIndex},
	{"GetObjectType", &GetObjectType},
	{"GetPaletteEntries", &GetPaletteEntries},
	{"GetPath", &GetPath},
	{"GetPixel", &GetPixel},
	{"GetPixelFormat", &GetPixelFormat},
	{"GetPolyFillMode", &GetPolyFillMode},
	{"GetRandomRgn", &GetRandomRgn},
	{"GetRasterizerCaps", &GetRasterizerCaps},
	{"GetRegionData", &GetRegionData},
	{"GetRgnBox", &GetRgnBox},
	{"GetROP2", &GetROP2},
	{"GetStockObject", &GetStockObject},
	{"GetStretchBltMode", &GetStretchBltMode},
	{"GetSystemPaletteEntries", &GetSystemPaletteEntries},
	{"GetSystemPaletteUse", &GetSystemPaletteUse},
	{"GetTextAlign", &GetTextAlign},
	{"GetTextCharacterExtra", &GetTextCharacterExtra},
	{"GetTextCharset", &GetTextCharset},
	{"GetTextCharsetInfo", &GetTextCharsetInfo},
	{"GetTextColor", &GetTextColor},
	{"GetTextExtentExPointI", &GetTextExtentExPointI},
	{"GetTextExtentPointI", &GetTextExtentPointI},
	{"GetViewportExtEx", &GetViewportExtEx},
	{"GetViewportOrgEx", &GetViewportOrgEx},
	{"GetWindowExtEx", &GetWindowExtEx},
	{"GetWindowOrgEx", &GetWindowOrgEx},
	{"GetWinMetaFileBits", &GetWinMetaFileBits},
	{"GetWorldTransform", &GetWorldTransform},
	{"IntersectClipRect", &IntersectClipRect},
	{"InvertRgn", &InvertRgn},
	{"LineDDA", &LineDDA},
	{"LineTo", &LineTo},
	{"LPtoDP", &LPtoDP},
	{"MaskBlt", &MaskBlt},
	{"ModifyWorldTransform", &ModifyWorldTransform},
	{"MoveToEx", &MoveToEx},
	{"OffsetClipRgn", &OffsetClipRgn},
	{"OffsetRgn", &OffsetRgn},
	{"OffsetViewportOrgEx", &OffsetViewportOrgEx},
	{"OffsetWindowOrgEx", &OffsetWindowOrgEx},
	{"PaintRgn", &PaintRgn},
	{"PatBlt", &PatBlt},
	{"PathToRegion", &PathToRegion},
	{"Pie", &Pie},
	{"PlayEnhMetaFile", &PlayEnhMetaFile},
	{"PlayEnhMetaFileRecord", &PlayEnhMetaFileRecord},
	{"PlayMetaFile", &PlayMetaFile},
	{"PlayMetaFileRecord", &PlayMetaFileRecord},
	{"PlgBlt", &PlgBlt},
	{"PolyBezier", &PolyBezier},
	{"PolyBezierTo", &PolyBezierTo},
	{"PolyDraw", &PolyDraw},
	{"Polygon", &Polygon},
	{"Polyline", &Polyline},
	{"PolylineTo", &PolylineTo},
	{"PolyPolygon", &PolyPolygon},
	{"PolyPolyline", &PolyPolyline},
	{"PtInRegion", &PtInRegion},
	{"PtVisible", &PtVisible},
	{"RealizePalette", &RealizePalette},
	{"Rectangle", &Rectangle},
	{"RectInRegion", &RectInRegion},
	{"RectVisible", &RectVisible},
	{"RemoveFontMemResourceEx", &RemoveFontMemResourceEx},
	{"ResizePalette", &ResizePalette},
	{"RestoreDC", &RestoreDC},
	{"RoundRect", &RoundRect},
	{"SaveDC", &SaveDC},
	{"ScaleViewportExtEx", &ScaleViewportExtEx},
	{"ScaleWindowExtEx", &ScaleWindowExtEx},
	{"SelectClipPath", &SelectClipPath},
	{"SelectClipRgn", &SelectClipRgn},
	{"SelectObject", &SelectObject},
	{"SelectPalette", &SelectPalette},
	{"SetAbortProc", &SetAbortProc},
	{"SetArcDirection", &SetArcDirection},
	{"SetBitmapBits", &SetBitmapBits},
	{"SetBitmapDimensionEx", &SetBitmapDimensionEx},
	{"SetBkColor", &SetBkColor},
	{"SetBkMode", &SetBkMode},
	{"SetBoundsRect", &SetBoundsRect},
	{"SetBrushOrgEx", &SetBrushOrgEx},
	{"SetColorAdjustment", &SetColorAdjustment},
	{"SetColorSpace", &SetColorSpace},
	{"SetDCBrushColor", &SetDCBrushColor},
	{"SetDCPenColor", &SetDCPenColor},
	{"SetDeviceGammaRamp", &SetDeviceGammaRamp},
	{"SetDIBColorTable", &SetDIBColorTable},
	{"SetDIBits", &SetDIBits},
	{"SetDIBitsToDevice", &SetDIBitsToDevice},
	{"SetEnhMetaFileBits", &SetEnhMetaFileBits},
	{"SetGraphicsMode", &SetGraphicsMode},
	{"SetICMMode", &SetICMMode},
	{"SetLayout", &SetLayout},
	{"SetMapMode", &SetMapMode},
	{"SetMapperFlags", &SetMapperFlags},
	{"SetMetaFileBitsEx", &SetMetaFileBitsEx},
	{"SetMetaRgn", &SetMetaRgn},
	{"SetMiterLimit", &SetMiterLimit},
	{"SetPaletteEntries", &SetPaletteEntries},
	{"SetPixel", &SetPixel},
	{"SetPixelFormat", &SetPixelFormat},
	{"SetPixelV", &SetPixelV},
	{"SetPolyFillMode", &SetPolyFillMode},
	{"SetRectRgn", &SetRectRgn},
	{"SetROP2", &SetROP2},
	{"SetStretchBltMode", &SetStretchBltMode},
	{"SetSystemPaletteUse", &SetSystemPaletteUse},
	{"SetTextAlign", &SetTextAlign},
	{"SetTextCharacterExtra", &SetTextCharacterExtra},
	{"SetTextColor", &SetTextColor},
	{"SetTextJustification", &SetTextJustification},
	{"SetViewportExtEx", &SetViewportExtEx},
	{"SetViewportOrgEx", &SetViewportOrgEx},
	{"SetWindowExtEx", &SetWindowExtEx},
	{"SetWindowOrgEx", &SetWindowOrgEx},
	{"SetWinMetaFileBits", &SetWinMetaFileBits},
	{"SetWorldTransform", &SetWorldTransform},
	{"StartPage", &StartPage},
	{"StretchBlt", &StretchBlt},
	{"StretchDIBits", &StretchDIBits},
	{"StrokeAndFillPath", &StrokeAndFillPath},
	{"StrokePath", &StrokePath},
	{"SwapBuffers", &SwapBuffers},
	{"TranslateCharsetInfo", &TranslateCharsetInfo},
	{"UnrealizeObject", &UnrealizeObject},
	{"UpdateColors", &UpdateColors},
	{"wglCopyContext", &WglCopyContext},
	{"wglCreateContext", &WglCreateContext},
	{"wglCreateLayerContext", &WglCreateLayerContext},
	{"wglDeleteContext", &WglDeleteContext},
	{"wglDescribeLayerPlane", &WglDescribeLayerPlane},
	{"wglGetCurrentContext", &WglGetCurrentContext},
	{"wglGetCurrentDC", &WglGetCurrentDC},
	{"wglGetLayerPaletteEntries", &WglGetLayerPaletteEntries},
	{"wglGetProcAddress", &WglGetProcAddress},
	{"wglMakeCurrent", &WglMakeCurrent},
	{"wglRealizeLayerPalette", &WglRealizeLayerPalette},
	{"wglSetLayerPaletteEntries", &WglSetLayerPaletteEntries},
	{"wglShareLists", &WglShareLists},
	{"wglSwapLayerBuffers", &WglSwapLayerBuffers},
	{"wglSwapMultipleBuffers", &WglSwapMultipleBuffers},
	{"WidenPath", &WidenPath},
}

var WinGdiANSIApis = outside.Apis{
	{"AddFontResourceA", &AddFontResource},
	{"AddFontResourceExA", &AddFontResourceEx},
	{"CopyEnhMetaFileA", &CopyEnhMetaFile},
	{"CopyMetaFileA", &CopyMetaFile},
	{"CreateColorSpaceA", &CreateColorSpaceA},
	{"CreateDCA", &CreateDCA},
	{"CreateEnhMetaFileA", &CreateEnhMetaFile},
	{"CreateFontA", &CreateFont},
	{"CreateFontIndirectA", &CreateFontIndirectA},
	{"CreateFontIndirectExA", &CreateFontIndirectExA},
	{"CreateICA", &CreateICA},
	{"CreateMetaFileA", &CreateMetaFile},
	{"CreateScalableFontResourceA", &CreateScalableFontResource},
	{"EnumFontFamiliesA", &EnumFontFamiliesA},
	{"EnumFontFamiliesExA", &EnumFontFamiliesExA},
	{"EnumFontsA", &EnumFontsA},
	{"EnumICMProfilesA", &EnumICMProfilesA},
	{"ExtTextOutA", &ExtTextOut},
	{"GetCharABCWidthsA", &GetCharABCWidths},
	{"GetCharABCWidthsFloatA", &GetCharABCWidthsFloat},
	{"GetCharacterPlacementA", &GetCharacterPlacementA},
	{"GetCharWidthA", &GetCharWidth},
	{"GetCharWidth32A", &GetCharWidth32},
	{"GetCharWidthFloatA", &GetCharWidthFloat},
	{"GetEnhMetaFileA", &GetEnhMetaFile},
	{"GetEnhMetaFileDescriptionA", &GetEnhMetaFileDescription},
	{"GetGlyphIndicesA", &GetGlyphIndices},
	{"GetICMProfileA", &GetICMProfile},
	{"GetLogColorSpaceA", &GetLogColorSpaceA},
	{"GetMetaFileA", &GetMetaFile},
	{"GetObjectA", &GetObject},
	{"GetOutlineTextMetricsA", &GetOutlineTextMetricsA},
	{"GetTextExtentExPointA", &GetTextExtentExPoint},
	{"GetTextExtentPointA", &GetTextExtentPoint},
	{"GetTextExtentPoint32A", &GetTextExtentPoint32},
	{"GetTextFaceA", &GetTextFace},
	{"GetTextMetricsA", &GetTextMetricsA},
	{"PolyTextOutA", &PolyTextOut},
	{"RemoveFontResourceA", &RemoveFontResource},
	{"RemoveFontResourceExA", &RemoveFontResourceEx},
	{"ResetDCA", &ResetDCA},
	{"SetICMProfileA", &SetICMProfile},
	{"StartDocA", &StartDoc},
	{"TextOutA", &TextOut},
	{"UpdateICMRegKeyA", &UpdateICMRegKey},
	{"wglUseFontBitmapsA", &WglUseFontBitmaps},
	{"wglUseFontOutlinesA", &WglUseFontOutlines},
}

var WinGdiUnicodeApis = outside.Apis{
	{"AddFontResourceW", &AddFontResource},
	{"AddFontResourceExW", &AddFontResourceEx},
	{"CopyEnhMetaFileW", &CopyEnhMetaFile},
	{"CopyMetaFileW", &CopyMetaFile},
	{"CreateColorSpaceW", &CreateColorSpaceW},
	{"CreateDCW", &CreateDCW},
	{"CreateEnhMetaFileW", &CreateEnhMetaFile},
	{"CreateFontW", &CreateFont},
	{"CreateFontIndirectW", &CreateFontIndirectW},
	{"CreateFontIndirectExW", &CreateFontIndirectExW},
	{"CreateICW", &CreateICW},
	{"CreateMetaFileW", &CreateMetaFile},
	{"CreateScalableFontResourceW", &CreateScalableFontResource},
	{"EnumFontFamiliesW", &EnumFontFamiliesW},
	{"EnumFontFamiliesExW", &EnumFontFamiliesExW},
	{"EnumFontsW", &EnumFontsW},
	{"EnumICMProfilesW", &EnumICMProfilesW},
	{"ExtTextOutW", &ExtTextOut},
	{"GetCharABCWidthsW", &GetCharABCWidths},
	{"GetCharABCWidthsFloatW", &GetCharABCWidthsFloat},
	{"GetCharacterPlacementW", &GetCharacterPlacementW},
	{"GetCharWidthW", &GetCharWidth},
	{"GetCharWidth32W", &GetCharWidth32},
	{"GetCharWidthFloatW", &GetCharWidthFloat},
	{"GetEnhMetaFileW", &GetEnhMetaFile},
	{"GetEnhMetaFileDescriptionW", &GetEnhMetaFileDescription},
	{"GetGlyphIndicesW", &GetGlyphIndices},
	{"GetICMProfileW", &GetICMProfile},
	{"GetLogColorSpaceW", &GetLogColorSpaceW},
	{"GetMetaFileW", &GetMetaFile},
	{"GetObjectW", &GetObject},
	{"GetOutlineTextMetricsW", &GetOutlineTextMetricsW},
	{"GetTextExtentExPointW", &GetTextExtentExPoint},
	{"GetTextExtentPointW", &GetTextExtentPoint},
	{"GetTextExtentPoint32W", &GetTextExtentPoint32},
	{"GetTextFaceW", &GetTextFace},
	{"GetTextMetricsW", &GetTextMetricsW},
	{"PolyTextOutW", &PolyTextOut},
	{"RemoveFontResourceW", &RemoveFontResource},
	{"RemoveFontResourceExW", &RemoveFontResourceEx},
	{"ResetDCW", &ResetDCW},
	{"SetICMProfileW", &SetICMProfile},
	{"StartDocW", &StartDoc},
	{"TextOutW", &TextOut},
	{"UpdateICMRegKeyW", &UpdateICMRegKey},
	{"wglUseFontBitmapsW", &WglUseFontBitmaps},
	{"wglUseFontOutlinesW", &WglUseFontOutlines},
}
