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
	GetBitmapBits func(bit T.HBITMAP, c T.LONG, bits *T.VOID) T.LONG

	//GetCharWidth is obsolete; use GetCharWidth32
	GetCharWidth func(dc T.HDC, first, last T.UINT, buffer *T.INT) T.BOOL

	//SetBitmapBits is obsolete; use SetDIBits
	SetBitmapBits func(bm T.HBITMAP, cb T.DWORD, bits *T.VOID) T.LONG

	AddFontResource func(VString) int

	AnimatePalette func(
		p T.HPALETTE,
		startIndex T.UINT,
		entries T.UINT,
		pe *T.PALETTEENTRY) T.BOOL

	Arc func(
		dc T.HDC, x1, y1, x2, y2, x3, y3, x4, y4 int) T.BOOL

	BitBlt func(
		dc T.HDC,
		x, y, cx, cy int,
		src T.HDC,
		x1, y1 int,
		rop T.DWORD) T.BOOL

	CancelDC func(
		dc T.HDC) T.BOOL

	Chord func(
		dc T.HDC, x1, y1, x2, y2, x3, y3, x4, y4 int) T.BOOL

	ChoosePixelFormat func(T.HDC, *T.PIXELFORMATDESCRIPTOR) int

	CloseMetaFile func(T.HDC) T.HMETAFILE

	CombineRgn func(dst, src1, src2 T.HRGN, mode int) int

	CopyMetaFile func(T.HMETAFILE, VString) T.HMETAFILE

	CreateBitmap func(
		width, height int,
		planes, bitCount T.UINT,
		bits *T.VOID) T.HBITMAP

	CreateBitmapIndirect func(*T.BITMAP) T.HBITMAP

	CreateBrushIndirect func(*T.LOGBRUSH) T.HBRUSH

	CreateCompatibleBitmap func(dc T.HDC, cx, cy int) T.HBITMAP

	CreateDiscardableBitmap func(dc T.HDC, cx, cy int) T.HBITMAP

	CreateCompatibleDC func(dc T.HDC) T.HDC

	CreateDCA func(
		driver, device, port VString, dm *T.DEVMODEA) T.HDC

	CreateDCW func(
		driver, device, port VString, dm *T.DEVMODEW) T.HDC

	CreateDIBitmap func(
		dc T.HDC,
		bmih *T.BITMAPINFOHEADER,
		init T.DWORD,
		bits *T.VOID,
		bmi *T.BITMAPINFO,
		usage T.UINT) T.HBITMAP

	CreateDIBPatternBrush func(h T.HGLOBAL, usage T.UINT) T.HBRUSH

	CreateDIBPatternBrushPt func(
		packedDIB *T.VOID, usage T.UINT) T.HBRUSH

	CreateEllipticRgn func(x1, y1, x2, y2 int) T.HRGN

	CreateEllipticRgnIndirect func(*T.RECT) T.HRGN

	CreateFontIndirectA func(*T.LOGFONTA) T.HFONT

	CreateFontIndirectW func(*T.LOGFONTW) T.HFONT

	CreateFont func(
		height, width, escapement, orientation, weight int,
		italic, underline, strikeOut, charSet, outPrecision,
		clipPrecision, quality, pitchAndFamily T.DWORD,
		faceName VString) T.HFONT

	CreateHatchBrush func(
		hatch int,
		color T.COLORREF) T.HBRUSH

	CreateICA func(
		driver, device, port VString, dm *T.DEVMODEA) T.HDC

	CreateICW func(
		driver, device, port VString, dm *T.DEVMODEW) T.HDC

	CreateMetaFile func(file VString) T.HDC

	CreatePalette func(*T.LOGPALETTE) T.HPALETTE

	CreatePen func(style, width int, color T.COLORREF) T.HPEN

	CreatePenIndirect func(*T.LOGPEN) T.HPEN

	CreatePolyPolygonRgn func(
		ptl *T.POINT, c *T.INT, poly, mode int) T.HRGN

	CreatePatternBrush func(T.HBITMAP) T.HBRUSH

	CreateRectRgn func(x1, y1, x2, y2 int) T.HRGN

	CreateRectRgnIndirect func(*T.RECT) T.HRGN

	CreateRoundRectRgn func(x1, y1, x2, y2, w, h int) T.HRGN

	CreateScalableFontResource func(
		hidden T.DWORD,
		font, file, path VString) T.BOOL

	CreateSolidBrush func(T.COLORREF) T.HBRUSH

	DeleteDC func(T.HDC) T.BOOL

	DeleteMetaFile func(T.HMETAFILE) T.BOOL

	DeleteObject func(T.HGDIOBJ) T.BOOL

	DescribePixelFormat func(
		dc T.HDC,
		pixelFormat int,
		bytes T.UINT,
		pfd *T.PIXELFORMATDESCRIPTOR) int

	DeviceCapabilitiesA func(
		device, port VString,
		capability T.WORD,
		output OVString,
		dm *T.DEVMODEA) int

	DeviceCapabilitiesW func(
		device, port VString,
		capability T.WORD,
		output OVString,
		dm *T.DEVMODEW) int

	DrawEscape func(dc T.HDC, escape, cIn int, in T.AString) int

	Ellipse func(dc T.HDC, left, top, right, bottom int) T.BOOL

	EnumFontFamiliesExA func(
		dc T.HDC,
		lf *T.LOGFONTA,
		p T.FONTENUMPROCA,
		l T.LPARAM,
		flags T.DWORD) int

	EnumFontFamiliesExW func(
		dc T.HDC,
		lf *T.LOGFONTW,
		p T.FONTENUMPROCW,
		l T.LPARAM,
		flags T.DWORD) int

	EnumFontFamiliesA func(
		dc T.HDC,
		lf T.AString,
		p T.FONTENUMPROCA,
		l T.LPARAM) int

	EnumFontFamiliesW func(
		dc T.HDC,
		lf T.WString,
		p T.FONTENUMPROCW,
		l T.LPARAM) int

	EnumFontsA func(
		dc T.HDC,
		lf T.AString,
		p T.FONTENUMPROCA,
		l T.LPARAM) int

	EnumFontsW func(
		dc T.HDC,
		lf T.WString,
		p T.FONTENUMPROCW,
		l T.LPARAM) int

	EnumObjects func(
		dc T.HDC,
		typ int,
		p T.GOBJENUMPROC,
		l T.LPARAM) int

	EqualRgn func(rgn1, rgn2 T.HRGN) T.BOOL

	Escape func(
		dc T.HDC, escape, cIn int, in T.AString, out *T.VOID) int

	///////////////////////////////////////////////////////
	ExtEscape func(
		dc T.HDC,
		escape, input int,
		inData VString,
		output int,
		outData OVString) int

	ExcludeClipRect func(
		dc T.HDC, left, top, right, bottom int) int

	ExtCreateRegion func(
		x *T.XFORM, count T.DWORD, data *T.RGNDATA) T.HRGN

	ExtFloodFill func(dc T.HDC, x, y int, color T.COLORREF, typ T.UINT) T.BOOL

	FillRgn func(dc T.HDC, rgn T.HRGN, br T.HBRUSH) T.BOOL

	FloodFill func(dc T.HDC, x, y int, color T.COLORREF) T.BOOL

	FrameRgn func(
		dc T.HDC, rgn T.HRGN, br T.HBRUSH, w, h int) T.BOOL

	GetROP2 func(dc T.HDC) int

	GetAspectRatioFilterEx func(dc T.HDC, size *T.SIZE) T.BOOL

	GetBkColor func(dc T.HDC) T.COLORREF

	GetDCBrushColor func(dc T.HDC) T.COLORREF

	GetDCPenColor func(dc T.HDC) T.COLORREF

	GetBkMode func(dc T.HDC) int

	GetBitmapDimensionEx func(bit T.HBITMAP, size *T.SIZE) T.BOOL

	GetBoundsRect func(dc T.HDC, rect *T.RECT, flags T.UINT) T.UINT

	GetBrushOrgEx func(dc T.HDC, pt *T.POINT) T.BOOL

	GetCharWidth32 func(
		dc T.HDC, first, last T.UINT, buffer *T.INT) T.BOOL

	GetCharWidthFloat func(
		dc T.HDC, first, last T.UINT, buffer *T.FLOAT) T.BOOL

	GetCharABCWidths func(
		dc T.HDC, first, last T.UINT, abc *T.ABC) T.BOOL

	GetCharABCWidthsFloat func(
		dc T.HDC, first T.UINT, last T.UINT, abc *T.ABCFLOAT) T.BOOL

	GetClipBox func(dc T.HDC, rect *T.RECT) int

	GetClipRgn func(dc T.HDC, rgn T.HRGN) int

	GetMetaRgn func(dc T.HDC, rgn T.HRGN) int

	GetCurrentObject func(dc T.HDC, typ T.UINT) T.HGDIOBJ

	GetCurrentPositionEx func(dc T.HDC, pt *T.POINT) T.BOOL

	GetDeviceCaps func(dc T.HDC, index int) int

	GetDIBits func(
		dc T.HDC,
		bm T.HBITMAP,
		start, lines T.UINT,
		vBits *T.VOID,
		bmi *T.BITMAPINFO,
		usage T.UINT) int

	GetFontData func(
		dc T.HDC,
		Table, Offset T.DWORD,
		Buffer *T.VOID,
		cBuffer T.DWORD) T.DWORD

	GetGlyphOutline func(
		dc T.HDC,
		Char T.UINT,
		Format T.UINT,
		gm *T.GLYPHMETRICS,
		cBuffer T.DWORD,
		Buffer *T.VOID,
		mat2 *T.MAT2) T.DWORD

	GetGraphicsMode func(dc T.HDC) int

	GetMapMode func(dc T.HDC) int

	GetMetaFileBitsEx func(
		MF T.HMETAFILE,
		cBuffer T.UINT,
		Data *T.VOID) T.UINT

	GetMetaFile func(Name VString) T.HMETAFILE

	GetNearestColor func(dc T.HDC, color T.COLORREF) T.COLORREF

	GetNearestPaletteIndex func(
		h T.HPALETTE, color T.COLORREF) T.UINT

	GetObjectType func(h T.HGDIOBJ) T.DWORD

	GetOutlineTextMetricsA func(
		dc T.HDC,
		cCopy T.UINT,
		otm *T.OUTLINETEXTMETRICA) T.UINT

	GetOutlineTextMetricsW func(
		dc T.HDC,
		cCopy T.UINT,
		otm *T.OUTLINETEXTMETRICW) T.UINT

	GetPaletteEntries func(
		pal T.HPALETTE,
		start, entries T.UINT,
		palEntries *T.PALETTEENTRY) T.UINT

	GetPixel func(
		dc T.HDC,
		x, y int) T.COLORREF

	GetPixelFormat func(dc T.HDC) int

	GetPolyFillMode func(dc T.HDC) int

	GetRasterizerCaps func(
		raststat *T.RASTERIZER_STATUS,
		cBytes T.UINT) T.BOOL

	GetRandomRgn func(dc T.HDC, rgn T.HRGN, i T.INT) int

	GetRegionData func(
		rgn T.HRGN, count T.DWORD, rgnData *T.RGNDATA) T.DWORD

	GetRgnBox func(rgn T.HRGN, rc *T.RECT) int

	GetStockObject func(i int) T.HGDIOBJ

	GetStretchBltMode func(dc T.HDC) int

	GetSystemPaletteEntries func(
		dc T.HDC,
		start, entries T.UINT,
		palEntries *T.PALETTEENTRY) T.UINT

	GetSystemPaletteUse func(hdc T.HDC) T.UINT

	GetTextCharacterExtra func(hdc T.HDC) int

	GetTextAlign func(hdc T.HDC) T.UINT

	GetTextColor func(hdc T.HDC) T.COLORREF

	GetTextExtentPoint func(
		dc T.HDC, s VString, c int, sz *T.SIZE) T.BOOL

	GetTextExtentPoint32 func(
		dc T.HDC, s VString, c int, sizl *T.SIZE) T.BOOL

	GetTextExtentExPoint func(
		dc T.HDC,
		s VString,
		cStr, maxExtent int,
		fit, dx *T.INT,
		size *T.SIZE) T.BOOL

	GetTextCharset func(dc T.HDC) int

	GetTextCharsetInfo func(
		dc T.HDC,
		sig *T.FONTSIGNATURE,
		flags T.DWORD) int

	TranslateCharsetInfo func(
		Src *T.DWORD,
		Cs *T.CHARSETINFO,
		Flags T.DWORD) T.BOOL

	GetFontLanguageInfo func(dc T.HDC) T.DWORD

	GetCharacterPlacementA func(
		dc T.HDC,
		String T.AString,
		Count,
		MexExtent int,
		Results *T.GCP_RESULTS,
		Flags T.DWORD) T.DWORD

	GetCharacterPlacementW func(
		dc T.HDC,
		String T.WString,
		Count,
		MexExtent int,
		Results *T.GCP_RESULTS,
		Flags T.DWORD) T.DWORD

	GetFontUnicodeRanges func(dc T.HDC, gs *T.GLYPHSET) T.DWORD

	GetGlyphIndices func(
		dc T.HDC,
		str VString,
		c int,
		gi *T.WORD,
		l T.DWORD) T.DWORD

	GetTextExtentPointI func(
		dc T.HDC, giIn *T.WORD, gi int, size *T.SIZE) T.BOOL

	GetTextExtentExPointI func(
		hdc T.HDC,
		lpwszString *T.WORD,
		cwchString int,
		nMaxExtent int,
		lpnFit *T.INT,
		lpnDx *T.INT,
		Size *T.SIZE) T.BOOL

	GetCharWidthI func(
		dc T.HDC,
		giFirst T.UINT,
		cgi T.UINT,
		gi *T.WORD,
		Widths *T.INT) T.BOOL

	GetCharABCWidthsI func(
		dc T.HDC,
		giFirst T.UINT,
		cgi T.UINT,
		gi *T.WORD,
		abc *T.ABC) T.BOOL

	AddFontResourceEx func(
		name VString, l T.DWORD, res *T.VOID) int

	RemoveFontResourceEx func(
		name VString, l T.DWORD, dv *T.VOID) T.BOOL

	AddFontMemResourceEx func(
		fileView *T.VOID,
		size T.DWORD,
		_ *T.VOID,
		numFonts *T.DWORD) T.HANDLE

	RemoveFontMemResourceEx func(h T.HANDLE) T.BOOL

	CreateFontIndirectExA func(*T.ENUMLOGFONTEXDVA) T.HFONT

	CreateFontIndirectExW func(*T.ENUMLOGFONTEXDVW) T.HFONT

	GetViewportExtEx func(dc T.HDC, size *T.SIZE) T.BOOL

	GetViewportOrgEx func(dc T.HDC, point *T.POINT) T.BOOL

	GetWindowExtEx func(dc T.HDC, size *T.SIZE) T.BOOL

	GetWindowOrgEx func(dc T.HDC, point *T.POINT) T.BOOL

	IntersectClipRect func(
		dc T.HDC, left, top, right, bottom int) int

	InvertRgn func(dc T.HDC, rgn T.HRGN) T.BOOL

	LineDDA func(
		xStart, yStart, xEnd, yEnd int,
		Proc T.LINEDDAPROC,
		data T.LPARAM) T.BOOL

	LineTo func(hdc T.HDC, x, y int) T.BOOL

	MaskBlt func(
		Dest T.HDC,
		xDest, yDest, width, height int,
		Src T.HDC,
		xSrc, ySrc int,
		Mask T.HBITMAP,
		xMask, yMask int,
		rop T.DWORD) T.BOOL

	PlgBlt func(
		Dest T.HDC,
		Point *T.POINT,
		Src T.HDC,
		xSrc, ySrc, width, height int,
		mask T.HBITMAP,
		xMask, yMask int) T.BOOL

	OffsetClipRgn func(dc T.HDC, x, y int) int

	OffsetRgn func(rgn T.HRGN, x, y int) int

	PatBlt func(dc T.HDC, x, y, w, h int, rop T.DWORD) T.BOOL

	Pie func(
		hdc T.HDC,
		left, top, right, bottom,
		xr1, yr1, xr2, yr2 int) T.BOOL

	PlayMetaFile func(dc T.HDC, mf T.HMETAFILE) T.BOOL

	PaintRgn func(dc T.HDC, rgn T.HRGN) T.BOOL

	PolyPolygon func(
		dc T.HDC, pt *T.POINT, psz *T.INT, sz int) T.BOOL

	PtInRegion func(rgn T.HRGN, x int, y int) T.BOOL

	PtVisible func(dc T.HDC, x int, y int) T.BOOL

	RectInRegion func(rgn T.HRGN, rect *T.RECT) T.BOOL

	RectVisible func(dc T.HDC, rect *T.RECT) T.BOOL

	Rectangle func(dc T.HDC, left, top, right, bottom int) T.BOOL

	RestoreDC func(dc T.HDC, savedDC int) T.BOOL

	ResetDCA func(dc T.HDC, dm *T.DEVMODEA) T.HDC

	ResetDCW func(dc T.HDC, dm *T.DEVMODEW) T.HDC

	RealizePalette func(dc T.HDC) T.UINT

	RemoveFontResource func(FileName VString) T.BOOL

	RoundRect func(
		dc T.HDC,
		left, top, right, bottom, width, height int) T.BOOL

	ResizePalette func(pal T.HPALETTE, n T.UINT) T.BOOL

	SaveDC func(dc T.HDC) int

	SelectClipRgn func(dc T.HDC, rgn T.HRGN) int

	ExtSelectClipRgn func(dc T.HDC, rgn T.HRGN, mode int) int

	SetMetaRgn func(dc T.HDC) int

	SelectObject func(dc T.HDC, h T.HGDIOBJ) T.HGDIOBJ

	SelectPalette func(
		dc T.HDC, pal T.HPALETTE, forceBkgd T.BOOL) T.HPALETTE

	SetBkColor func(dc T.HDC, color T.COLORREF) T.COLORREF

	SetDCBrushColor func(dc T.HDC, color T.COLORREF) T.COLORREF

	SetDCPenColor func(dc T.HDC, color T.COLORREF) T.COLORREF

	SetBkMode func(dc T.HDC, mode int) int

	SetBoundsRect func(
		dc T.HDC, rect *T.RECT, flags T.UINT) T.UINT

	SetDIBits func(
		dc T.HDC,
		bm T.HBITMAP,
		start, lines T.UINT,
		Bits *T.VOID,
		bmi *T.BITMAPINFO,
		ColorUse T.UINT) int

	SetDIBitsToDevice func(
		dc T.HDC,
		xDest, yDest int,
		w, h T.DWORD,
		xSrc, ySrc int,
		startScan, lines T.UINT,
		Bits *T.VOID,
		bmi *T.BITMAPINFO,
		ColorUse T.UINT) int

	SetMapperFlags func(dc T.HDC, flags T.DWORD) T.DWORD

	SetGraphicsMode func(dc T.HDC, mode int) int

	SetMapMode func(dc T.HDC, mode int) int

	SetLayout func(dc T.HDC, l T.DWORD) T.DWORD

	GetLayout func(dc T.HDC) T.DWORD

	SetMetaFileBitsEx func(
		buffer T.UINT, data *T.BYTE) T.HMETAFILE

	SetPaletteEntries func(
		pal T.HPALETTE,
		start, entries T.UINT,
		PalEntries *T.PALETTEENTRY) T.UINT

	SetPixel func(
		dc T.HDC, x, y int, color T.COLORREF) T.COLORREF

	SetPixelV func(
		hdc T.HDC, x, y int, color T.COLORREF) T.BOOL

	SetPixelFormat func(
		dc T.HDC,
		format int,
		pfd *T.PIXELFORMATDESCRIPTOR) T.BOOL

	SetPolyFillMode func(dc T.HDC, mode int) int

	StretchBlt func(
		Dest T.HDC,
		xDest, yDest, wDest, hDest int,
		Src T.HDC,
		xSrc, ySrc, wSrc, hSrc int,
		rop T.DWORD) T.BOOL

	SetRectRgn func(
		rgn T.HRGN,
		left, top, right, bottom int) T.BOOL

	StretchDIBits func(
		hdc T.HDC,
		xDest, yDest,
		destWidth, destHeight,
		xSrc, ySrc,
		srcWidth, srcHeight int,
		bits *T.VOID,
		bmi *T.BITMAPINFO,
		usage T.UINT,
		rop T.DWORD) int

	SetROP2 func(dc T.HDC, rop2 int) int

	SetStretchBltMode func(dc T.HDC, mode int) int

	SetSystemPaletteUse func(dc T.HDC, use T.UINT) T.UINT

	SetTextCharacterExtra func(dc T.HDC, extra int) int

	SetTextColor func(dc T.HDC, color T.COLORREF) T.COLORREF

	SetTextAlign func(dc T.HDC, align T.UINT) T.UINT

	SetTextJustification func(dc T.HDC, extra, count int) T.BOOL

	UpdateColors func(hdc T.HDC) T.BOOL

	AlphaBlend func(
		dest T.HDC,
		xoriginDest, yoriginDest,
		wDest, hDest int,
		src T.HDC,
		xoriginSrc, yoriginSrc,
		wSrc, hSrc int,
		ftn T.BLENDFUNCTION) T.BOOL

	TransparentBlt func(
		Dest T.HDC,
		xoriginDest, yoriginDest,
		wDest, hDest int,
		Src T.HDC,
		xoriginSrc, yoriginSrc,
		wSrc, hSrc int,
		crTransparent T.UINT) T.BOOL

	GradientFill func(
		dc T.HDC,
		Vertex *T.TRIVERTEX,
		nVertex T.ULONG,
		Mesh *T.VOID,
		nMesh T.ULONG,
		Mode T.ULONG) T.BOOL

	PlayMetaFileRecord func(
		dc T.HDC,
		HandleTable *T.HANDLETABLE,
		MR *T.METARECORD,
		Objs T.UINT) T.BOOL

	EnumMetaFile func(
		dc T.HDC,
		mf T.HMETAFILE,
		proc T.MFENUMPROC,
		param T.LPARAM) T.BOOL

	CloseEnhMetaFile func(
		dc T.HDC) T.HENHMETAFILE

	CopyEnhMetaFile func(
		Enh T.HENHMETAFILE,
		FileName VString) T.HENHMETAFILE

	CreateEnhMetaFile func(
		dc T.HDC,
		Filename VString,
		rc *T.RECT,
		Desc VString) T.HDC

	DeleteEnhMetaFile func(mf T.HENHMETAFILE) T.BOOL

	EnumEnhMetaFile func(
		dc T.HDC,
		mf T.HENHMETAFILE,
		proc T.ENHMFENUMPROC,
		param *T.VOID,
		Rect *T.RECT) T.BOOL

	GetEnhMetaFile func(Name VString) T.HENHMETAFILE

	GetEnhMetaFileBits func(
		emf T.HENHMETAFILE, size T.UINT, data *T.BYTE) T.UINT

	GetEnhMetaFileDescription func(
		emf T.HENHMETAFILE,
		cBuffer T.UINT,
		description OVString) T.UINT

	GetEnhMetaFileHeader func(
		emf T.HENHMETAFILE,
		size T.UINT,
		enhMetaHeader *T.ENHMETAHEADER) T.UINT

	GetEnhMetaFilePaletteEntries func(
		emf T.HENHMETAFILE,
		NumEntries T.UINT,
		PaletteEntries *T.PALETTEENTRY) T.UINT

	GetEnhMetaFilePixelFormat func(
		emf T.HENHMETAFILE,
		cBuffer T.UINT,
		pfd *T.PIXELFORMATDESCRIPTOR) T.UINT

	GetWinMetaFileBits func(
		emf T.HENHMETAFILE,
		cData16 T.UINT,
		Data16 *T.BYTE,
		MapMode T.INT,
		Ref T.HDC) T.UINT

	PlayEnhMetaFile func(
		dc T.HDC,
		mf T.HENHMETAFILE,
		rect *T.RECT) T.BOOL

	PlayEnhMetaFileRecord func(
		dc T.HDC,
		ht *T.HANDLETABLE,
		mr *T.ENHMETARECORD,
		cht T.UINT) T.BOOL

	SetEnhMetaFileBits func(
		nSize T.UINT,
		pb *T.BYTE) T.HENHMETAFILE

	SetWinMetaFileBits func(
		Size T.UINT,
		Meta16Data *T.BYTE,
		Ref T.HDC,
		MFP *T.METAFILEPICT) T.HENHMETAFILE

	GdiComment func(
		dc T.HDC,
		Size T.UINT,
		Data *T.BYTE) T.BOOL

	GetTextMetricsA func(dc T.HDC, tm *T.TEXTMETRICA) T.BOOL

	GetTextMetricsW func(dc T.HDC, tm *T.TEXTMETRICW) T.BOOL

	AngleArc func(
		dc T.HDC,
		x, y int,
		r T.DWORD,
		StartAngle T.FLOAT,
		SweepAngle T.FLOAT) T.BOOL

	PolyPolyline func(
		dc T.HDC, pt *T.POINT, psz *T.DWORD, sz T.DWORD) T.BOOL

	GetWorldTransform func(dc T.HDC, xf *T.XFORM) T.BOOL

	SetWorldTransform func(dc T.HDC, xf *T.XFORM) T.BOOL

	ModifyWorldTransform func(
		dc T.HDC, xf *T.XFORM, mode T.DWORD) T.BOOL

	CombineTransform func(
		out *T.XFORM, xf1, xf2 *T.XFORM) T.BOOL

	CreateDIBSection func(
		dc T.HDC,
		bmi *T.BITMAPINFO,
		usage T.UINT,
		Bits **T.VOID,
		Section T.HANDLE,
		offset T.DWORD) T.HBITMAP

	GetDIBColorTable func(
		dc T.HDC,
		Start T.UINT,
		Entries T.UINT,
		rgbq *T.RGBQUAD) T.UINT

	SetDIBColorTable func(
		dc T.HDC,
		Start T.UINT,
		Entries T.UINT,
		rgbq *T.RGBQUAD) T.UINT

	SetColorAdjustment func(
		dc T.HDC, ca *T.COLORADJUSTMENT) T.BOOL

	GetColorAdjustment func(
		dc T.HDC, ca *T.COLORADJUSTMENT) T.BOOL

	CreateHalftonePalette func(dc T.HDC) T.HPALETTE

	StartDoc func(dc T.HDC, di *T.DOCINFO) int

	EndDoc func(dc T.HDC) int

	StartPage func(dc T.HDC) int

	EndPage func(dc T.HDC) int

	AbortDoc func(dc T.HDC) int

	SetAbortProc func(dc T.HDC, proc T.ABORTPROC) int

	AbortPath func(hdc T.HDC) T.BOOL

	ArcTo func(
		dc T.HDC,
		left, top, right, bottom,
		xr1, yr1, xr2, yr2 int) T.BOOL

	BeginPath func(dc T.HDC) T.BOOL

	CloseFigure func(dc T.HDC) T.BOOL

	EndPath func(dc T.HDC) T.BOOL

	FillPath func(dc T.HDC) T.BOOL

	FlattenPath func(dc T.HDC) T.BOOL

	GetPath func(dc T.HDC, pt *T.POINT, t *T.BYTE, cpt int) int

	PathToRegion func(hdc T.HDC) T.HRGN

	PolyDraw func(dc T.HDC, pt *T.POINT, t *T.BYTE, cpt int) T.BOOL

	SelectClipPath func(dc T.HDC, mode int) T.BOOL

	SetArcDirection func(dc T.HDC, dir int) int

	SetMiterLimit func(dc T.HDC, limit T.FLOAT, old *T.FLOAT) T.BOOL

	StrokeAndFillPath func(dc T.HDC) T.BOOL

	StrokePath func(dc T.HDC) T.BOOL

	WidenPath func(dc T.HDC) T.BOOL

	ExtCreatePen func(
		PenStyle T.DWORD,
		Width T.DWORD,
		LBrush *T.LOGBRUSH,
		cStyle T.DWORD,
		Style *T.DWORD) T.HPEN

	GetMiterLimit func(dc T.HDC, limit *T.FLOAT) T.BOOL

	GetArcDirection func(dc T.HDC) int

	GetObject func(h T.HANDLE, c int, v *T.VOID) int

	MoveToEx func(dc T.HDC, x, y int, pt *T.POINT) T.BOOL

	TextOut func(
		dc T.HDC, x, y int, lpString VString, c int) T.BOOL

	ExtTextOut func(
		dc T.HDC,
		x, y int,
		options T.UINT,
		rect *T.RECT,
		String VString,
		c T.UINT,
		Dx *T.INT) T.BOOL

	PolyTextOut func(
		dc T.HDC, pt *T.POLYTEXT, nstrings int) T.BOOL

	CreatePolygonRgn func(
		ptl *T.POINT,
		cPoint int,
		Mode int) T.HRGN

	DPtoLP func(dc T.HDC, pt *T.POINT, c int) T.BOOL

	LPtoDP func(dc T.HDC, pt *T.POINT, c int) T.BOOL

	Polygon func(dc T.HDC, pt *T.POINT, cpt int) T.BOOL

	Polyline func(dc T.HDC, pt *T.POINT, cpt int) T.BOOL

	PolyBezier func(dc T.HDC, pt *T.POINT, cpt T.DWORD) T.BOOL

	PolyBezierTo func(dc T.HDC, pt *T.POINT, cpt T.DWORD) T.BOOL

	PolylineTo func(dc T.HDC, pt *T.POINT, cpt T.DWORD) T.BOOL

	SetViewportExtEx func(dc T.HDC, x, y int, sz *T.SIZE) T.BOOL

	SetViewportOrgEx func(dc T.HDC, x, y int, pt *T.POINT) T.BOOL

	SetWindowExtEx func(dc T.HDC, x, y int, sz *T.SIZE) T.BOOL

	SetWindowOrgEx func(dc T.HDC, x, y int, pt *T.POINT) T.BOOL

	OffsetViewportOrgEx func(dc T.HDC, x, y int, pt *T.POINT) T.BOOL

	OffsetWindowOrgEx func(dc T.HDC, x, y int, pt *T.POINT) T.BOOL

	ScaleViewportExtEx func(
		hdc T.HDC, xn, xd, yn, yd int, sz *T.SIZE) T.BOOL

	ScaleWindowExtEx func(
		dc T.HDC, xn, xd, yn, yd int, sz *T.SIZE) T.BOOL

	SetBitmapDimensionEx func(
		bm T.HBITMAP, w, h int, sz *T.SIZE) T.BOOL

	SetBrushOrgEx func(dc T.HDC, x, y int, pt *T.POINT) T.BOOL

	GetTextFace func(dc T.HDC, c int, name VString) int

	GetKerningPairs func(
		dc T.HDC, pairs T.DWORD, kernPair *T.KERNINGPAIR) T.DWORD

	GetDCOrgEx func(dc T.HDC, pt *T.POINT) T.BOOL

	FixBrushOrgEx func(
		dc T.HDC, x, y int, ptl *T.POINT) T.BOOL

	UnrealizeObject func(h T.HGDIOBJ) T.BOOL

	GdiFlush func() T.BOOL

	GdiSetBatchLimit func(dw T.DWORD) T.DWORD

	GdiGetBatchLimit func() T.DWORD

	SetICMMode func(dc T.HDC, mode int) int

	CheckColorsInGamut func(
		dc T.HDC,
		RGBTriple *T.VOID,
		Buffer *T.VOID,
		Count T.DWORD) T.BOOL

	GetColorSpace func(dc T.HDC) T.HCOLORSPACE

	GetLogColorSpaceA func(
		ColorSpace T.HCOLORSPACE,
		Buffer *T.LOGCOLORSPACEA,
		Size T.DWORD) T.BOOL

	GetLogColorSpaceW func(
		ColorSpace T.HCOLORSPACE,
		Buffer *T.LOGCOLORSPACEW,
		Size T.DWORD) T.BOOL

	CreateColorSpaceA func(
		lcs *T.LOGCOLORSPACEA) T.HCOLORSPACE

	CreateColorSpaceW func(
		lcs *T.LOGCOLORSPACEW) T.HCOLORSPACE

	SetColorSpace func(
		dc T.HDC, cs T.HCOLORSPACE) T.HCOLORSPACE

	DeleteColorSpace func(cs T.HCOLORSPACE) T.BOOL

	GetICMProfile func(
		dc T.HDC, bufSize *T.DWORD, filename VString) T.BOOL

	SetICMProfile func(dc T.HDC, fileName VString) T.BOOL

	GetDeviceGammaRamp func(dc T.HDC, ramp *T.VOID) T.BOOL

	SetDeviceGammaRamp func(dc T.HDC, ramp *T.VOID) T.BOOL

	ColorMatchToTarget func(
		dc T.HDC, target T.HDC, action T.DWORD) T.BOOL

	EnumICMProfilesA func(
		dc T.HDC, proc T.ICMENUMPROCA, param T.LPARAM) int

	EnumICMProfilesW func(
		dc T.HDC, proc T.ICMENUMPROCW, param T.LPARAM) int

	UpdateICMRegKey func(
		reserved T.DWORD,
		cmid, fileName VString,
		command T.UINT) T.BOOL

	ColorCorrectPalette func(
		dc T.HDC,
		pal T.HPALETTE,
		first T.DWORD,
		num T.DWORD) T.BOOL

	WglCopyContext func(T.HGLRC, T.HGLRC, T.UINT) T.BOOL

	WglCreateContext func(T.HDC) T.HGLRC

	WglCreateLayerContext func(T.HDC, int) T.HGLRC

	WglDeleteContext func(T.HGLRC) T.BOOL

	WglGetCurrentContext func(T.VOID) T.HGLRC

	WglGetCurrentDC func(T.VOID) T.HDC

	WglGetProcAddress func(T.AString) T.PROC

	WglMakeCurrent func(T.HDC, T.HGLRC) T.BOOL

	WglShareLists func(T.HGLRC, T.HGLRC) T.BOOL

	WglUseFontBitmaps func(T.HDC, T.DWORD, T.DWORD, T.DWORD) T.BOOL

	SwapBuffers func(T.HDC) T.BOOL

	WglUseFontOutlines func(
		T.HDC, T.DWORD, T.DWORD, T.DWORD,
		T.FLOAT, T.FLOAT, int, *T.GLYPHMETRICSFLOAT) T.BOOL

	WglDescribeLayerPlane func(
		T.HDC, int, int, T.UINT, *T.LAYERPLANEDESCRIPTOR) T.BOOL

	WglSetLayerPaletteEntries func(
		T.HDC, int, int, int, *T.COLORREF) int

	WglGetLayerPaletteEntries func(
		T.HDC, int, int, int, *T.COLORREF) int

	WglRealizeLayerPalette func(T.HDC, int, T.BOOL) T.BOOL

	WglSwapLayerBuffers func(T.HDC, T.WGL_SWAP_FLAG) T.BOOL

	WglSwapMultipleBuffers func(T.UINT, *T.WGLSWAP) T.DWORD
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
