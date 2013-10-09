package wingdi

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/gdi32"
	_ "github.com/tHinqa/outside/win32/opengl32"
)

//GetBitmapBits is obsolete; use GetDIBits
var GetBitmapBits func(bit HBITMAP, c LONG, Bits *VOID) LONG

//GetCharWidth is obsolete; use GetCharWidth32
var GetCharWidth func(dc HDC, First, Last UINT, Buffer *INT) BOOL

//SetBitmapBits is obsolete; use SetDIBits
var SetBitmapBits func(bm HBITMAP, cb DWORD, Bits *VOID) LONG

var AddFontResource func(VString) int

var AnimatePalette func(
	p HPALETTE,
	startIndex UINT,
	entries UINT,
	pe *PALETTEENTRY) BOOL

var Arc func(
	dc HDC, x1, y1, x2, y2, x3, y3, x4, y4 int) BOOL

var BitBlt func(
	dc HDC,
	x, y, cx, cy int,
	src HDC,
	x1, y1 int,
	rop DWORD) BOOL

var CancelDC func(
	dc HDC) BOOL

var Chord func(
	dc HDC, x1, y1, x2, y2, x3, y3, x4, y4 int) BOOL

var ChoosePixelFormat func(HDC, *PIXELFORMATDESCRIPTOR) int

var CloseMetaFile func(HDC) HMETAFILE

var CombineRgn func(dst, src1, src2 HRGN, mode int) int

var CopyMetaFile func(HMETAFILE, VString) HMETAFILE

var CreateBitmap func(
	width, height int,
	planes, bitCount UINT,
	bits *VOID) HBITMAP

var CreateBitmapIndirect func(*BITMAP) HBITMAP

var CreateBrushIndirect func(*LOGBRUSH) HBRUSH

var CreateCompatibleBitmap func(dc HDC, cx, cy int) HBITMAP

var CreateDiscardableBitmap func(dc HDC, cx, cy int) HBITMAP

var CreateCompatibleDC func(dc HDC) HDC

var CreateDCA func(
	driver, device, port VString, dm *DEVMODEA) HDC

var CreateDCW func(
	driver, device, port VString, dm *DEVMODEW) HDC

var CreateDIBitmap func(
	dc HDC,
	bmih *BITMAPINFOHEADER,
	init DWORD,
	bits *VOID,
	bmi *BITMAPINFO,
	usage UINT) HBITMAP

var CreateDIBPatternBrush func(h HGLOBAL, usage UINT) HBRUSH

var CreateDIBPatternBrushPt func(
	packedDIB *VOID, usage UINT) HBRUSH

var CreateEllipticRgn func(x1, y1, x2, y2 int) HRGN

var CreateEllipticRgnIndirect func(*RECT) HRGN

var CreateFontIndirectA func(*LOGFONTA) HFONT

var CreateFontIndirectW func(*LOGFONTW) HFONT

var CreateFont func(
	height, width, escapement, orientation, weight int,
	italic, underline, strikeOut, charSet, outPrecision,
	clipPrecision, quality, pitchAndFamily DWORD,
	faceName VString) HFONT

var CreateHatchBrush func(
	hatch int,
	color COLORREF) HBRUSH

var CreateICA func(
	driver, device, port VString, dm *DEVMODEA) HDC

var CreateICW func(
	driver, device, port VString, dm *DEVMODEW) HDC

var CreateMetaFile func(file VString) HDC

var CreatePalette func(*LOGPALETTE) HPALETTE

var CreatePen func(style, width int, color COLORREF) HPEN

var CreatePenIndirect func(*LOGPEN) HPEN

var CreatePolyPolygonRgn func(
	ptl *POINT, c *INT, poly, mode int) HRGN

var CreatePatternBrush func(HBITMAP) HBRUSH

var CreateRectRgn func(x1, y1, x2, y2 int) HRGN

var CreateRectRgnIndirect func(*RECT) HRGN

var CreateRoundRectRgn func(x1, y1, x2, y2, w, h int) HRGN

var CreateScalableFontResource func(
	hidden DWORD,
	font, file, path VString) BOOL

var CreateSolidBrush func(COLORREF) HBRUSH

var DeleteDC func(HDC) BOOL

var DeleteMetaFile func(HMETAFILE) BOOL

var DeleteObject func(HGDIOBJ) BOOL

var DescribePixelFormat func(
	dc HDC,
	pixelFormat int,
	bytes UINT,
	pfd *PIXELFORMATDESCRIPTOR) int

var DeviceCapabilitiesA func(
	device, port VString,
	capability WORD,
	output OVString,
	dm *DEVMODEA) int

var DeviceCapabilitiesW func(
	device, port VString,
	capability WORD,
	output OVString,
	dm *DEVMODEW) int

var DrawEscape func(dc HDC, escape, cIn int, in AString) int

var Ellipse func(dc HDC, left, top, right, bottom int) BOOL

var EnumFontFamiliesExA func(
	dc HDC,
	lf *LOGFONTA,
	p FONTENUMPROCA,
	l LPARAM,
	flags DWORD) int

var EnumFontFamiliesExW func(
	dc HDC,
	lf *LOGFONTW,
	p FONTENUMPROCW,
	l LPARAM,
	flags DWORD) int

var EnumFontFamiliesA func(
	dc HDC,
	lf AString,
	p FONTENUMPROCA,
	l LPARAM) int

var EnumFontFamiliesW func(
	dc HDC,
	lf WString,
	p FONTENUMPROCW,
	l LPARAM) int

var EnumFontsA func(
	dc HDC,
	lf AString,
	p FONTENUMPROCA,
	l LPARAM) int

var EnumFontsW func(
	dc HDC,
	lf WString,
	p FONTENUMPROCW,
	l LPARAM) int

var EnumObjects func(
	dc HDC,
	typ int,
	p GOBJENUMPROC,
	l LPARAM) int

var EqualRgn func(rgn1, rgn2 HRGN) BOOL

var Escape func(
	dc HDC, escape, cIn int, in AString, out *VOID) int

///////////////////////////////////////////////////////
var ExtEscape func(
	dc HDC,
	escape, input int,
	inData VString,
	output int,
	outData OVString) int

var ExcludeClipRect func(
	dc HDC, left, top, right, bottom int) int

var ExtCreateRegion func(
	x *XFORM, count DWORD, data *RGNDATA) HRGN

var ExtFloodFill func(dc HDC, x, y int, color COLORREF, typ UINT) BOOL

var FillRgn func(dc HDC, rgn HRGN, br HBRUSH) BOOL

var FloodFill func(dc HDC, x, y int, color COLORREF) BOOL

var FrameRgn func(
	dc HDC, rgn HRGN, br HBRUSH, w, h int) BOOL

var GetROP2 func(dc HDC) int

var GetAspectRatioFilterEx func(dc HDC, size *SIZE) BOOL

var GetBkColor func(dc HDC) COLORREF

var GetDCBrushColor func(dc HDC) COLORREF

var GetDCPenColor func(dc HDC) COLORREF

var GetBkMode func(dc HDC) int

var GetBitmapDimensionEx func(bit HBITMAP, size *SIZE) BOOL

var GetBoundsRect func(dc HDC, rect *RECT, flags UINT) UINT

var GetBrushOrgEx func(dc HDC, pt *POINT) BOOL

var GetCharWidth32 func(
	dc HDC, First, Last UINT, Buffer *INT) BOOL

var GetCharWidthFloat func(
	dc HDC, First, Last UINT, Buffer *FLOAT) BOOL

var GetCharABCWidths func(
	dc HDC, First, Last UINT, ABC *ABC) BOOL

var GetCharABCWidthsFloat func(
	dc HDC, First UINT, Last UINT, ABC *ABCFLOAT) BOOL

var GetClipBox func(dc HDC, rect *RECT) int

var GetClipRgn func(dc HDC, rgn HRGN) int

var GetMetaRgn func(dc HDC, rgn HRGN) int

var GetCurrentObject func(dc HDC, typ UINT) HGDIOBJ

var GetCurrentPositionEx func(dc HDC, pt *POINT) BOOL

var GetDeviceCaps func(dc HDC, index int) int

var GetDIBits func(
	dc HDC,
	bm HBITMAP,
	start, lines UINT,
	vBits *VOID,
	bmi *BITMAPINFO,
	usage UINT) int

var GetFontData func(
	dc HDC,
	Table, Offset DWORD,
	Buffer *VOID,
	cBuffer DWORD) DWORD

var GetGlyphOutline func(
	dc HDC,
	Char UINT,
	Format UINT,
	gm *GLYPHMETRICS,
	cBuffer DWORD,
	Buffer *VOID,
	mat2 *MAT2) DWORD

var GetGraphicsMode func(dc HDC) int

var GetMapMode func(dc HDC) int

var GetMetaFileBitsEx func(
	MF HMETAFILE,
	cBuffer UINT,
	Data *VOID) UINT

var GetMetaFile func(Name VString) HMETAFILE

var GetNearestColor func(dc HDC, color COLORREF) COLORREF

var GetNearestPaletteIndex func(
	h HPALETTE, color COLORREF) UINT

var GetObjectType func(h HGDIOBJ) DWORD

var GetOutlineTextMetricsA func(
	dc HDC,
	cCopy UINT,
	otm *OUTLINETEXTMETRICA) UINT

var GetOutlineTextMetricsW func(
	dc HDC,
	cCopy UINT,
	otm *OUTLINETEXTMETRICW) UINT

var GetPaletteEntries func(
	pal HPALETTE,
	Start, Entries UINT,
	PalEntries *PALETTEENTRY) UINT

var GetPixel func(
	dc HDC,
	x, y int) COLORREF

var GetPixelFormat func(dc HDC) int

var GetPolyFillMode func(dc HDC) int

var GetRasterizerCaps func(
	raststat *RASTERIZER_STATUS,
	cBytes UINT) BOOL

var GetRandomRgn func(dc HDC, rgn HRGN, i INT) int

var GetRegionData func(
	rgn HRGN, Count DWORD, RgnData *RGNDATA) DWORD

var GetRgnBox func(rgn HRGN, rc *RECT) int

var GetStockObject func(i int) HGDIOBJ

var GetStretchBltMode func(dc HDC) int

var GetSystemPaletteEntries func(
	dc HDC,
	Start, Entries UINT,
	PalEntries *PALETTEENTRY) UINT

var GetSystemPaletteUse func(hdc HDC) UINT

var GetTextCharacterExtra func(hdc HDC) int

var GetTextAlign func(hdc HDC) UINT

var GetTextColor func(hdc HDC) COLORREF

var GetTextExtentPoint func(
	dc HDC, String VString, c int, sz *SIZE) BOOL

var GetTextExtentPoint32 func(
	dc HDC, String VString, c int, sizl *SIZE) BOOL

var GetTextExtentExPoint func(
	dc HDC,
	String VString,
	cString, MaxExtent int,
	Fit, Dx *INT,
	Size *SIZE) BOOL

var GetTextCharset func(dc HDC) int

var GetTextCharsetInfo func(
	dc HDC,
	Sig *FONTSIGNATURE,
	Flags DWORD) int

var TranslateCharsetInfo func(
	Src *DWORD,
	Cs *CHARSETINFO,
	Flags DWORD) BOOL

var GetFontLanguageInfo func(dc HDC) DWORD

var GetCharacterPlacementA func(
	dc HDC,
	String AString,
	Count,
	MexExtent int,
	Results *GCP_RESULTS,
	Flags DWORD) DWORD

var GetCharacterPlacementW func(
	dc HDC,
	String WString,
	Count,
	MexExtent int,
	Results *GCP_RESULTS,
	Flags DWORD) DWORD

var GetFontUnicodeRanges func(dc HDC, gs *GLYPHSET) DWORD

var GetGlyphIndices func(
	dc HDC,
	str VString,
	c int,
	gi *WORD,
	l DWORD) DWORD

var GetTextExtentPointI func(
	dc HDC, giIn *WORD, gi int, size *SIZE) BOOL

var GetTextExtentExPointI func(
	hdc HDC,
	lpwszString *WORD,
	cwchString int,
	nMaxExtent int,
	lpnFit *INT,
	lpnDx *INT,
	Size *SIZE) BOOL

var GetCharWidthI func(
	dc HDC,
	giFirst UINT,
	cgi UINT,
	gi *WORD,
	Widths *INT) BOOL

var GetCharABCWidthsI func(
	dc HDC,
	giFirst UINT,
	cgi UINT,
	gi *WORD,
	abc *ABC) BOOL

var AddFontResourceEx func(
	name VString, l DWORD, res *VOID) int

var RemoveFontResourceEx func(
	name VString, l DWORD, dv *VOID) BOOL

var AddFontMemResourceEx func(
	FileView *VOID,
	Size DWORD,
	Resrved *VOID,
	NumFonts *DWORD) HANDLE

var RemoveFontMemResourceEx func(h HANDLE) BOOL

var CreateFontIndirectExA func(*ENUMLOGFONTEXDVA) HFONT

var CreateFontIndirectExW func(*ENUMLOGFONTEXDVW) HFONT

var GetViewportExtEx func(dc HDC, size *SIZE) BOOL

var GetViewportOrgEx func(dc HDC, point *POINT) BOOL

var GetWindowExtEx func(dc HDC, size *SIZE) BOOL

var GetWindowOrgEx func(dc HDC, point *POINT) BOOL

var IntersectClipRect func(
	dc HDC, left, top, right, bottom int) int

var InvertRgn func(dc HDC, rgn HRGN) BOOL

var LineDDA func(
	xStart, yStart, xEnd, yEnd int,
	Proc LINEDDAPROC,
	data LPARAM) BOOL

var LineTo func(hdc HDC, x, y int) BOOL

var MaskBlt func(
	Dest HDC,
	xDest, yDest, width, height int,
	Src HDC,
	xSrc, ySrc int,
	Mask HBITMAP,
	xMask, yMask int,
	rop DWORD) BOOL

var PlgBlt func(
	Dest HDC,
	Point *POINT,
	Src HDC,
	xSrc, ySrc, width, height int,
	Mask HBITMAP,
	xMask, yMask int) BOOL

var OffsetClipRgn func(dc HDC, x, y int) int

var OffsetRgn func(rgn HRGN, x, y int) int

var PatBlt func(dc HDC, x, y, w, h int, rop DWORD) BOOL

var Pie func(
	hdc HDC,
	left, top, right, bottom,
	xr1, yr1, xr2, yr2 int) BOOL

var PlayMetaFile func(dc HDC, mf HMETAFILE) BOOL

var PaintRgn func(dc HDC, rgn HRGN) BOOL

var PolyPolygon func(
	dc HDC, pt *POINT, psz *INT, sz int) BOOL

var PtInRegion func(rgn HRGN, x int, y int) BOOL

var PtVisible func(dc HDC, x int, y int) BOOL

var RectInRegion func(rgn HRGN, rect *RECT) BOOL

var RectVisible func(dc HDC, rect *RECT) BOOL

var Rectangle func(dc HDC, left, top, right, bottom int) BOOL

var RestoreDC func(dc HDC, SavedDC int) BOOL

var ResetDCA func(dc HDC, dm *DEVMODEA) HDC

var ResetDCW func(dc HDC, dm *DEVMODEW) HDC

var RealizePalette func(dc HDC) UINT

var RemoveFontResource func(FileName VString) BOOL

var RoundRect func(
	dc HDC,
	left, top, right, bottom, width, height int) BOOL

var ResizePalette func(pal HPALETTE, n UINT) BOOL

var SaveDC func(dc HDC) int

var SelectClipRgn func(dc HDC, rgn HRGN) int

var ExtSelectClipRgn func(dc HDC, rgn HRGN, mode int) int

var SetMetaRgn func(dc HDC) int

var SelectObject func(dc HDC, h HGDIOBJ) HGDIOBJ

var SelectPalette func(
	dc HDC, Pal HPALETTE, ForceBkgd BOOL) HPALETTE

var SetBkColor func(dc HDC, color COLORREF) COLORREF

var SetDCBrushColor func(dc HDC, color COLORREF) COLORREF

var SetDCPenColor func(dc HDC, color COLORREF) COLORREF

var SetBkMode func(dc HDC, mode int) int

var SetBoundsRect func(
	dc HDC, rect *RECT, flags UINT) UINT

var SetDIBits func(
	dc HDC,
	bm HBITMAP,
	start, Lines UINT,
	Bits *VOID,
	bmi *BITMAPINFO,
	ColorUse UINT) int

var SetDIBitsToDevice func(
	dc HDC,
	xDest, yDest int,
	w, h DWORD,
	xSrc, ySrc int,
	StartScan, Lines UINT,
	Bits *VOID,
	bmi *BITMAPINFO,
	ColorUse UINT) int

var SetMapperFlags func(dc HDC, flags DWORD) DWORD

var SetGraphicsMode func(dc HDC, Mode int) int

var SetMapMode func(dc HDC, Mode int) int

var SetLayout func(dc HDC, l DWORD) DWORD

var GetLayout func(dc HDC) DWORD

var SetMetaFileBitsEx func(
	Buffer UINT, Data *BYTE) HMETAFILE

var SetPaletteEntries func(
	pal HPALETTE,
	Start, Entries UINT,
	PalEntries *PALETTEENTRY) UINT

var SetPixel func(
	dc HDC, x, y int, color COLORREF) COLORREF

var SetPixelV func(
	hdc HDC, x, y int, color COLORREF) BOOL

var SetPixelFormat func(
	dc HDC,
	format int,
	pfd *PIXELFORMATDESCRIPTOR) BOOL

var SetPolyFillMode func(dc HDC, mode int) int

var StretchBlt func(
	Dest HDC,
	xDest, yDest, wDest, hDest int,
	Src HDC,
	xSrc, ySrc, wSrc, hSrc int,
	rop DWORD) BOOL

var SetRectRgn func(
	rgn HRGN,
	left, top, right, bottom int) BOOL

var StretchDIBits func(
	hdc HDC,
	xDest, yDest,
	DestWidth, DestHeight,
	xSrc, ySrc,
	SrcWidth, SrcHeight int,
	Bits *VOID,
	bmi *BITMAPINFO,
	Usage UINT,
	rop DWORD) int

var SetROP2 func(dc HDC, rop2 int) int

var SetStretchBltMode func(dc HDC, mode int) int

var SetSystemPaletteUse func(dc HDC, use UINT) UINT

var SetTextCharacterExtra func(dc HDC, extra int) int

var SetTextColor func(dc HDC, color COLORREF) COLORREF

var SetTextAlign func(dc HDC, align UINT) UINT

var SetTextJustification func(dc HDC, extra, count int) BOOL

var UpdateColors func(hdc HDC) BOOL

var AlphaBlend func(
	Dest HDC,
	xoriginDest, yoriginDest,
	wDest, hDest int,
	Src HDC,
	xoriginSrc, yoriginSrc,
	wSrc, hSrc int,
	ftn BLENDFUNCTION) BOOL

var TransparentBlt func(
	Dest HDC,
	xoriginDest, yoriginDest,
	wDest, hDest int,
	Src HDC,
	xoriginSrc, yoriginSrc,
	wSrc, hSrc int,
	crTransparent UINT) BOOL

var GradientFill func(
	dc HDC,
	Vertex *TRIVERTEX,
	nVertex ULONG,
	Mesh *VOID,
	nMesh ULONG,
	Mode ULONG) BOOL

var PlayMetaFileRecord func(
	dc HDC,
	HandleTable *HANDLETABLE,
	MR *METARECORD,
	Objs UINT) BOOL

var EnumMetaFile func(
	dc HDC,
	mf HMETAFILE,
	proc MFENUMPROC,
	param LPARAM) BOOL

var CloseEnhMetaFile func(
	dc HDC) HENHMETAFILE

var CopyEnhMetaFile func(
	Enh HENHMETAFILE,
	FileName VString) HENHMETAFILE

var CreateEnhMetaFile func(
	dc HDC,
	Filename VString,
	rc *RECT,
	Desc VString) HDC

var DeleteEnhMetaFile func(mf HENHMETAFILE) BOOL

var EnumEnhMetaFile func(
	dc HDC,
	mf HENHMETAFILE,
	proc ENHMFENUMPROC,
	param *VOID,
	Rect *RECT) BOOL

var GetEnhMetaFile func(Name VString) HENHMETAFILE

var GetEnhMetaFileBits func(
	EMF HENHMETAFILE, Size UINT, Data *BYTE) UINT

var GetEnhMetaFileDescription func(
	emf HENHMETAFILE,
	cBuffer UINT,
	Description OVString) UINT

var GetEnhMetaFileHeader func(
	emf HENHMETAFILE,
	Size UINT,
	EnhMetaHeader *ENHMETAHEADER) UINT

var GetEnhMetaFilePaletteEntries func(
	emf HENHMETAFILE,
	NumEntries UINT,
	PaletteEntries *PALETTEENTRY) UINT

var GetEnhMetaFilePixelFormat func(
	emf HENHMETAFILE,
	cBuffer UINT,
	pfd *PIXELFORMATDESCRIPTOR) UINT

var GetWinMetaFileBits func(
	emf HENHMETAFILE,
	cData16 UINT,
	Data16 *BYTE,
	MapMode INT,
	Ref HDC) UINT

var PlayEnhMetaFile func(
	dc HDC,
	mf HENHMETAFILE,
	rect *RECT) BOOL

var PlayEnhMetaFileRecord func(
	dc HDC,
	ht *HANDLETABLE,
	mr *ENHMETARECORD,
	cht UINT) BOOL

var SetEnhMetaFileBits func(
	nSize UINT,
	pb *BYTE) HENHMETAFILE

var SetWinMetaFileBits func(
	Size UINT,
	Meta16Data *BYTE,
	Ref HDC,
	MFP *METAFILEPICT) HENHMETAFILE

var GdiComment func(
	dc HDC,
	Size UINT,
	Data *BYTE) BOOL

var GetTextMetricsA func(dc HDC, tm *TEXTMETRICA) BOOL

var GetTextMetricsW func(dc HDC, tm *TEXTMETRICW) BOOL

var AngleArc func(
	dc HDC,
	x, y int,
	r DWORD,
	StartAngle FLOAT,
	SweepAngle FLOAT) BOOL

var PolyPolyline func(
	dc HDC, pt *POINT, psz *DWORD, sz DWORD) BOOL

var GetWorldTransform func(dc HDC, xf *XFORM) BOOL

var SetWorldTransform func(dc HDC, xf *XFORM) BOOL

var ModifyWorldTransform func(
	dc HDC, xf *XFORM, mode DWORD) BOOL

var CombineTransform func(
	out *XFORM, xf1, xf2 *XFORM) BOOL

var CreateDIBSection func(
	dc HDC,
	bmi *BITMAPINFO,
	usage UINT,
	Bits **VOID,
	Section HANDLE,
	offset DWORD) HBITMAP

var GetDIBColorTable func(
	dc HDC,
	Start UINT,
	Entries UINT,
	rgbq *RGBQUAD) UINT

var SetDIBColorTable func(
	dc HDC,
	Start UINT,
	Entries UINT,
	rgbq *RGBQUAD) UINT

var SetColorAdjustment func(
	dc HDC, ca *COLORADJUSTMENT) BOOL

var GetColorAdjustment func(
	dc HDC, ca *COLORADJUSTMENT) BOOL

var CreateHalftonePalette func(dc HDC) HPALETTE

var StartDoc func(dc HDC, di *DOCINFO) int

var EndDoc func(dc HDC) int

var StartPage func(dc HDC) int

var EndPage func(dc HDC) int

var AbortDoc func(dc HDC) int

var SetAbortProc func(dc HDC, proc ABORTPROC) int

var AbortPath func(hdc HDC) BOOL

var ArcTo func(
	dc HDC,
	left, top, right, bottom,
	xr1, yr1, xr2, yr2 int) BOOL

var BeginPath func(dc HDC) BOOL

var CloseFigure func(dc HDC) BOOL

var EndPath func(dc HDC) BOOL

var FillPath func(dc HDC) BOOL

var FlattenPath func(dc HDC) BOOL

var GetPath func(dc HDC, pt *POINT, t *BYTE, cpt int) int

var PathToRegion func(hdc HDC) HRGN

var PolyDraw func(dc HDC, pt *POINT, t *BYTE, cpt int) BOOL

var SelectClipPath func(dc HDC, mode int) BOOL

var SetArcDirection func(dc HDC, dir int) int

var SetMiterLimit func(dc HDC, limit FLOAT, old *FLOAT) BOOL

var StrokeAndFillPath func(dc HDC) BOOL

var StrokePath func(dc HDC) BOOL

var WidenPath func(dc HDC) BOOL

var ExtCreatePen func(
	PenStyle DWORD,
	Width DWORD,
	LBrush *LOGBRUSH,
	cStyle DWORD,
	Style *DWORD) HPEN

var GetMiterLimit func(dc HDC, limit *FLOAT) BOOL

var GetArcDirection func(dc HDC) int

var GetObject func(h HANDLE, c int, v *VOID) int

var MoveToEx func(dc HDC, x, y int, pt *POINT) BOOL

var TextOut func(
	dc HDC, x, y int, lpString VString, c int) BOOL

var ExtTextOut func(
	dc HDC,
	x, y int,
	options UINT,
	rect *RECT,
	String VString,
	c UINT,
	Dx *INT) BOOL

var PolyTextOut func(
	dc HDC, pt *POLYTEXT, nstrings int) BOOL

var CreatePolygonRgn func(
	ptl *POINT,
	cPoint int,
	Mode int) HRGN

var DPtoLP func(dc HDC, pt *POINT, c int) BOOL

var LPtoDP func(dc HDC, pt *POINT, c int) BOOL

var Polygon func(dc HDC, pt *POINT, cpt int) BOOL

var Polyline func(dc HDC, pt *POINT, cpt int) BOOL

var PolyBezier func(dc HDC, pt *POINT, cpt DWORD) BOOL

var PolyBezierTo func(dc HDC, pt *POINT, cpt DWORD) BOOL

var PolylineTo func(dc HDC, pt *POINT, cpt DWORD) BOOL

var SetViewportExtEx func(dc HDC, x, y int, sz *SIZE) BOOL

var SetViewportOrgEx func(dc HDC, x, y int, pt *POINT) BOOL

var SetWindowExtEx func(dc HDC, x, y int, sz *SIZE) BOOL

var SetWindowOrgEx func(dc HDC, x, y int, pt *POINT) BOOL

var OffsetViewportOrgEx func(dc HDC, x, y int, pt *POINT) BOOL

var OffsetWindowOrgEx func(dc HDC, x, y int, pt *POINT) BOOL

var ScaleViewportExtEx func(
	hdc HDC, xn, xd, yn, yd int, sz *SIZE) BOOL

var ScaleWindowExtEx func(
	dc HDC, xn, xd, yn, yd int, sz *SIZE) BOOL

var SetBitmapDimensionEx func(
	bm HBITMAP, w, h int, sz *SIZE) BOOL

var SetBrushOrgEx func(dc HDC, x, y int, pt *POINT) BOOL

var GetTextFace func(dc HDC, c int, Name VString) int

var GetKerningPairs func(
	dc HDC, Pairs DWORD, KernPair *KERNINGPAIR) DWORD

var GetDCOrgEx func(dc HDC, pt *POINT) BOOL

var FixBrushOrgEx func(
	dc HDC, x, y int, ptl *POINT) BOOL

var UnrealizeObject func(h HGDIOBJ) BOOL

var GdiFlush func() BOOL

var GdiSetBatchLimit func(dw DWORD) DWORD

var GdiGetBatchLimit func() DWORD

var SetICMMode func(dc HDC, mode int) int

var CheckColorsInGamut func(
	dc HDC,
	RGBTriple *VOID,
	Buffer *VOID,
	Count DWORD) BOOL

var GetColorSpace func(dc HDC) HCOLORSPACE

var GetLogColorSpaceA func(
	ColorSpace HCOLORSPACE,
	Buffer *LOGCOLORSPACEA,
	Size DWORD) BOOL

var GetLogColorSpaceW func(
	ColorSpace HCOLORSPACE,
	Buffer *LOGCOLORSPACEW,
	Size DWORD) BOOL

var CreateColorSpaceA func(
	lcs *LOGCOLORSPACEA) HCOLORSPACE

var CreateColorSpaceW func(
	lcs *LOGCOLORSPACEW) HCOLORSPACE

var SetColorSpace func(
	dc HDC, cs HCOLORSPACE) HCOLORSPACE

var DeleteColorSpace func(cs HCOLORSPACE) BOOL

var GetICMProfile func(
	dc HDC, BufSize *DWORD, Filename VString) BOOL

var SetICMProfile func(dc HDC, FileName VString) BOOL

var GetDeviceGammaRamp func(dc HDC, Ramp *VOID) BOOL

var SetDeviceGammaRamp func(dc HDC, Ramp *VOID) BOOL

var ColorMatchToTarget func(
	dc HDC, Target HDC, action DWORD) BOOL

var EnumICMProfilesA func(
	dc HDC, proc ICMENUMPROCA, param LPARAM) int

var EnumICMProfilesW func(
	dc HDC, proc ICMENUMPROCW, param LPARAM) int

var UpdateICMRegKey func(
	reserved DWORD,
	CMID, FileName VString,
	command UINT) BOOL

var ColorCorrectPalette func(
	dc HDC,
	Pal HPALETTE,
	First DWORD,
	num DWORD) BOOL

var WglCopyContext func(HGLRC, HGLRC, UINT) BOOL

var WglCreateContext func(HDC) HGLRC

var WglCreateLayerContext func(HDC, int) HGLRC

var WglDeleteContext func(HGLRC) BOOL

var WglGetCurrentContext func(VOID) HGLRC

var WglGetCurrentDC func(VOID) HDC

var WglGetProcAddress func(AString) PROC

var WglMakeCurrent func(HDC, HGLRC) BOOL

var WglShareLists func(HGLRC, HGLRC) BOOL

var WglUseFontBitmaps func(HDC, DWORD, DWORD, DWORD) BOOL

var SwapBuffers func(HDC) BOOL

var WglUseFontOutlines func(
	HDC, DWORD, DWORD, DWORD,
	FLOAT, FLOAT, int, *GLYPHMETRICSFLOAT) BOOL

var WglDescribeLayerPlane func(
	HDC, int, int, UINT, *LAYERPLANEDESCRIPTOR) BOOL

var WglSetLayerPaletteEntries func(
	HDC, int, int, int, *COLORREF) int

var WglGetLayerPaletteEntries func(
	HDC, int, int, int, *COLORREF) int

var WglRealizeLayerPalette func(HDC, int, BOOL) BOOL

var WglSwapLayerBuffers func(HDC, WGL_SWAP_FLAG) BOOL

var WglSwapMultipleBuffers func(UINT, *WGLSWAP) DWORD

/*
TODO(t): msimg32
{"AlphaBlend", &AlphaBlend},
{"GradientFill", &GradientFill},
{"TransparentBlt", &TransparentBlt},

TODO(t): winspool
{"DeviceCapabilities", &DeviceCapabilities},
*/

var WinGdiApis = Apis{
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

var WinGdiANSIApis = Apis{
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

var WinGdiUnicodeApis = Apis{
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
