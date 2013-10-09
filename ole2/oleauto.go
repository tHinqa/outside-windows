package ole2

import (
	. "github.com/tHinqa/outside"
	//	_ "github.com/tHinqa/outside/win32/ole32"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/oleaut32"
)

type ()

var (
	SysAllocString func(*OLECHAR) BSTR

	SysReAllocString func(*BSTR, *OLECHAR) INT

	SysAllocStringLen func(*OLECHAR, UINT) BSTR

	SysReAllocStringLen func(*BSTR, *OLECHAR, UINT) INT

	SysFreeString func(BSTR)

	SysStringLen func(BSTR) UINT

	SysStringByteLen func(s BSTR) UINT

	SysAllocStringByteLen func(s AString, leng UINT) BSTR

	DosDateTimeToVariantTime func(
		DosDate, DosTime USHORT, Time *DOUBLE) INT

	VariantTimeToDosDateTime func(
		Time DOUBLE, DosDate, DosTime *USHORT) INT

	SystemTimeToVariantTime func(
		SystemTime *SYSTEMTIME, Time *DOUBLE) INT

	VariantTimeToSystemTime func(
		Time DOUBLE, SystemTime *SYSTEMTIME) INT

	SafeArrayAllocDescriptor func(Dims UINT, Out **SAFEARRAY)

	SafeArrayAllocDescriptorEx func(
		VT VARTYPE, Dims UINT, Out **SAFEARRAY)

	SafeArrayAllocData func(SA *SAFEARRAY)

	SafeArrayCreate func(
		VT VARTYPE, Dims UINT,
		SABound *SAFEARRAYBOUND) *SAFEARRAY

	SafeArrayCreateEx func(
		VT VARTYPE,
		Dims UINT,
		SABound *SAFEARRAYBOUND,
		pvExtra *VOID) *SAFEARRAY

	SafeArrayCopyData func(
		Source *SAFEARRAY,
		Target *SAFEARRAY)

	SafeArrayDestroyDescriptor func(
		SA *SAFEARRAY)

	SafeArrayDestroyData func(
		SA *SAFEARRAY)

	SafeArrayDestroy func(
		SA *SAFEARRAY)

	SafeArrayRedim func(
		SA *SAFEARRAY,
		psaboundNew *SAFEARRAYBOUND)

	SafeArrayGetDim func(
		SA *SAFEARRAY) UINT

	SafeArrayGetElemsize func(
		SA *SAFEARRAY) UINT

	SafeArrayGetUBound func(
		SA *SAFEARRAY,
		nDim UINT,
		plUbound *LONG)
	SafeArrayGetLBound func(
		SA *SAFEARRAY,
		nDim UINT,
		plLbound *LONG)
	SafeArrayLock func(
		SA *SAFEARRAY)
	SafeArrayUnlock func(
		SA *SAFEARRAY)
	SafeArrayAccessData func(
		SA *SAFEARRAY,
		ppvData **VOID)
	SafeArrayUnaccessData func(
		SA *SAFEARRAY)
	SafeArrayGetElement func(
		SA *SAFEARRAY,
		rgIndices *LONG,
		pv *VOID)
	SafeArrayPutElement func(
		SA *SAFEARRAY,
		rgIndices *LONG,
		pv *VOID)
	SafeArrayCopy func(
		SA *SAFEARRAY,
		Out **SAFEARRAY)
	SafeArrayPtrOfIndex func(
		SA *SAFEARRAY,
		rgIndices *LONG,
		ppvData **VOID)
	SafeArraySetRecordInfo func(
		SA *SAFEARRAY,
		prinfo *IRecordInfo)
	SafeArrayGetRecordInfo func(
		SA *SAFEARRAY,
		prinfo **IRecordInfo)
	SafeArraySetIID func(
		SA *SAFEARRAY,
		guid REFGUID)
	SafeArrayGetIID func(
		SA *SAFEARRAY,
		pguid *GUID)
	SafeArrayGetVartype func(
		SA *SAFEARRAY,
		pvt *VARTYPE)
	SafeArrayCreateVector func(
		VT VARTYPE,
		lLbound LONG,
		Elements ULONG) *SAFEARRAY
	SafeArrayCreateVectorEx func(
		VT VARTYPE,
		lLbound LONG,
		Elements ULONG,
		pvExtra *VOID) *SAFEARRAY

	VariantInit func(
		pvarg *VARIANTARG)
	VariantClear func(
		pvarg *VARIANTARG)
	VariantCopy func(
		Dest *VARIANTARG,
		Src *VARIANTARG)
	VariantCopyInd func(
		pvarDest *VARIANT,
		Src *VARIANTARG)
	VariantChangeType func(
		Dest *VARIANTARG,
		pvarSrc *VARIANTARG,
		wFlags USHORT,
		VT VARTYPE)
	VariantChangeTypeEx func(
		Dest *VARIANTARG,
		pvarSrc *VARIANTARG,
		lcid LCID,
		wFlags USHORT,
		VT VARTYPE)

	VectorFromBstr func(S BSTR, SA **SAFEARRAY)

	BstrFromVector func(SA *SAFEARRAY, S *BSTR)

	VarUI1FromI2 func(
		sIn SHORT,
		pbOut *BYTE)
	VarUI1FromI4 func(
		lIn LONG,
		pbOut *BYTE)
	VarUI1FromI8 func(
		i64In LONG64,
		pbOut *BYTE)
	VarUI1FromR4 func(
		fltIn FLOAT,
		pbOut *BYTE)
	VarUI1FromR8 func(
		dblIn DOUBLE,
		pbOut *BYTE)
	VarUI1FromCy func(
		In CY,
		pbOut *BYTE)
	VarUI1FromDate func(
		dateIn DATE,
		pbOut *BYTE)
	VarUI1FromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		pbOut *BYTE)
	VarUI1FromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		pbOut *BYTE)
	VarUI1FromBool func(
		boolIn VARIANT_BOOL,
		pbOut *BYTE)
	VarUI1FromI1 func(
		In Char,
		pbOut *BYTE)
	VarUI1FromUI2 func(
		uiIn USHORT,
		pbOut *BYTE)
	VarUI1FromUI4 func(
		ulIn ULONG,
		pbOut *BYTE)
	VarUI1FromUI8 func(
		ui64In ULONG64,
		pbOut *BYTE)
	VarUI1FromDec func(
		pdecIn *DECIMAL,
		pbOut *BYTE)

	VarI2FromUI1 func(
		bIn BYTE,
		Out *SHORT)
	VarI2FromI4 func(
		lIn LONG,
		Out *SHORT)
	VarI2FromI8 func(
		i64In LONG64,
		Out *SHORT)
	VarI2FromR4 func(
		fltIn FLOAT,
		Out *SHORT)
	VarI2FromR8 func(
		dblIn DOUBLE,
		Out *SHORT)
	VarI2FromCy func(
		In CY,
		Out *SHORT)
	VarI2FromDate func(
		dateIn DATE,
		Out *SHORT)
	VarI2FromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		Out *SHORT)
	VarI2FromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		Out *SHORT)
	VarI2FromBool func(
		boolIn VARIANT_BOOL,
		Out *SHORT)
	VarI2FromI1 func(
		In Char,
		Out *SHORT)
	VarI2FromUI2 func(
		uiIn USHORT,
		Out *SHORT)
	VarI2FromUI4 func(
		ulIn ULONG,
		Out *SHORT)
	VarI2FromUI8 func(
		ui64In ULONG64,
		Out *SHORT)
	VarI2FromDec func(
		pdecIn *DECIMAL,
		Out *SHORT)

	VarI4FromUI1 func(
		bIn BYTE,
		plOut *LONG)
	VarI4FromI2 func(
		sIn SHORT,
		plOut *LONG)
	VarI4FromI8 func(
		i64In LONG64,
		plOut *LONG)
	VarI4FromR4 func(
		fltIn FLOAT,
		plOut *LONG)
	VarI4FromR8 func(
		dblIn DOUBLE,
		plOut *LONG)
	VarI4FromCy func(
		In CY,
		plOut *LONG)
	VarI4FromDate func(
		dateIn DATE,
		plOut *LONG)
	VarI4FromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		plOut *LONG)
	VarI4FromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		plOut *LONG)
	VarI4FromBool func(
		boolIn VARIANT_BOOL,
		plOut *LONG)
	VarI4FromI1 func(
		In Char,
		plOut *LONG)
	VarI4FromUI2 func(
		uiIn USHORT,
		plOut *LONG)
	VarI4FromUI4 func(
		ulIn ULONG,
		plOut *LONG)
	VarI4FromUI8 func(
		ui64In ULONG64,
		plOut *LONG)
	VarI4FromDec func(
		pdecIn *DECIMAL,
		plOut *LONG)

	VarI4FromInt func(in INT, out *LONG)

	VarI8FromUI1 func(
		bIn BYTE,
		pi64Out *LONG64)
	VarI8FromI2 func(
		sIn SHORT,
		pi64Out *LONG64)

	VarI8FromI4 func(in LONG, out *LONG64)

	VarI8FromR4 func(
		fltIn FLOAT,
		pi64Out *LONG64)
	VarI8FromR8 func(
		dblIn DOUBLE,
		pi64Out *LONG64)
	VarI8FromCy func(
		In CY,
		pi64Out *LONG64)
	VarI8FromDate func(
		dateIn DATE,
		pi64Out *LONG64)
	VarI8FromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags Unsigned_long,
		pi64Out *LONG64)
	VarI8FromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		pi64Out *LONG64)
	VarI8FromBool func(
		boolIn VARIANT_BOOL,
		pi64Out *LONG64)
	VarI8FromI1 func(
		In Char,
		pi64Out *LONG64)
	VarI8FromUI2 func(
		uiIn USHORT,
		pi64Out *LONG64)
	VarI8FromUI4 func(
		ulIn ULONG,
		pi64Out *LONG64)
	VarI8FromUI8 func(
		ui64In ULONG64,
		pi64Out *LONG64)
	VarI8FromDec func(
		pdecIn *DECIMAL,
		pi64Out *LONG64)
	VarI8FromInt func(
		intIn INT,
		pi64Out *LONG64)

	VarR4FromUI1 func(
		bIn BYTE,
		pfltOut *FLOAT)
	VarR4FromI2 func(
		sIn SHORT,
		pfltOut *FLOAT)
	VarR4FromI4 func(
		lIn LONG,
		pfltOut *FLOAT)
	VarR4FromI8 func(
		i64In LONG64,
		pfltOut *FLOAT)
	VarR4FromR8 func(
		dblIn DOUBLE,
		pfltOut *FLOAT)
	VarR4FromCy func(
		In CY,
		pfltOut *FLOAT)
	VarR4FromDate func(
		dateIn DATE,
		pfltOut *FLOAT)
	VarR4FromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		pfltOut *FLOAT)
	VarR4FromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		pfltOut *FLOAT)
	VarR4FromBool func(
		boolIn VARIANT_BOOL,
		pfltOut *FLOAT)
	VarR4FromI1 func(
		In Char,
		pfltOut *FLOAT)
	VarR4FromUI2 func(
		uiIn USHORT,
		pfltOut *FLOAT)
	VarR4FromUI4 func(
		ulIn ULONG,
		pfltOut *FLOAT)
	VarR4FromUI8 func(
		ui64In ULONG64,
		pfltOut *FLOAT)
	VarR4FromDec func(
		pdecIn *DECIMAL,
		pfltOut *FLOAT)

	VarR8FromUI1 func(
		bIn BYTE,
		pdblOut *DOUBLE)
	VarR8FromI2 func(
		sIn SHORT,
		pdblOut *DOUBLE)
	VarR8FromI4 func(
		lIn LONG,
		pdblOut *DOUBLE)
	VarR8FromI8 func(
		i64In LONG64,
		pdblOut *DOUBLE)
	VarR8FromR4 func(
		fltIn FLOAT,
		pdblOut *DOUBLE)
	VarR8FromCy func(
		In CY,
		pdblOut *DOUBLE)
	VarR8FromDate func(
		dateIn DATE,
		pdblOut *DOUBLE)
	VarR8FromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		pdblOut *DOUBLE)
	VarR8FromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		pdblOut *DOUBLE)
	VarR8FromBool func(
		boolIn VARIANT_BOOL,
		pdblOut *DOUBLE)
	VarR8FromI1 func(
		In Char,
		pdblOut *DOUBLE)
	VarR8FromUI2 func(
		uiIn USHORT,
		pdblOut *DOUBLE)
	VarR8FromUI4 func(
		ulIn ULONG,
		pdblOut *DOUBLE)
	VarR8FromUI8 func(
		ui64In ULONG64,
		pdblOut *DOUBLE)
	VarR8FromDec func(
		pdecIn *DECIMAL,
		pdblOut *DOUBLE)

	VarDateFromUI1 func(
		bIn BYTE,
		pdateOut *DATE)
	VarDateFromI2 func(
		sIn SHORT,
		pdateOut *DATE)
	VarDateFromI4 func(
		lIn LONG,
		pdateOut *DATE)
	VarDateFromI8 func(
		i64In LONG64,
		pdateOut *DATE)
	VarDateFromR4 func(
		fltIn FLOAT,
		pdateOut *DATE)
	VarDateFromR8 func(
		dblIn DOUBLE,
		pdateOut *DATE)
	VarDateFromCy func(
		In CY,
		pdateOut *DATE)
	VarDateFromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		pdateOut *DATE)
	VarDateFromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		pdateOut *DATE)
	VarDateFromBool func(
		boolIn VARIANT_BOOL,
		pdateOut *DATE)
	VarDateFromI1 func(
		In Char,
		pdateOut *DATE)
	VarDateFromUI2 func(
		uiIn USHORT,
		pdateOut *DATE)
	VarDateFromUI4 func(
		ulIn ULONG,
		pdateOut *DATE)
	VarDateFromUI8 func(
		ui64In ULONG64,
		pdateOut *DATE)
	VarDateFromDec func(
		pdecIn *DECIMAL,
		pdateOut *DATE)

	VarCyFromUI1 func(
		bIn BYTE,
		pcyOut *CY)
	VarCyFromI2 func(
		sIn SHORT,
		pcyOut *CY)
	VarCyFromI4 func(
		lIn LONG,
		pcyOut *CY)
	VarCyFromI8 func(
		i64In LONG64,
		pcyOut *CY)
	VarCyFromR4 func(
		fltIn FLOAT,
		pcyOut *CY)
	VarCyFromR8 func(
		dblIn DOUBLE,
		pcyOut *CY)
	VarCyFromDate func(
		dateIn DATE,
		pcyOut *CY)
	VarCyFromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		pcyOut *CY)
	VarCyFromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		pcyOut *CY)
	VarCyFromBool func(
		boolIn VARIANT_BOOL,
		pcyOut *CY)
	VarCyFromI1 func(
		In Char,
		pcyOut *CY)
	VarCyFromUI2 func(
		uiIn USHORT,
		pcyOut *CY)
	VarCyFromUI4 func(
		ulIn ULONG,
		pcyOut *CY)
	VarCyFromUI8 func(
		ui64In ULONG64,
		pcyOut *CY)
	VarCyFromDec func(
		pdecIn *DECIMAL,
		pcyOut *CY)

	VarBstrFromUI1 func(
		bVal BYTE,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromI2 func(
		iVal SHORT,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromI4 func(
		lIn LONG,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromI8 func(
		i64In LONG64,
		lcid LCID,
		Flags Unsigned_long,
		pbstrOut *BSTR)
	VarBstrFromR4 func(
		fltIn FLOAT,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromR8 func(
		dblIn DOUBLE,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromCy func(
		In CY,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromDate func(
		dateIn DATE,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromBool func(
		boolIn VARIANT_BOOL,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromI1 func(
		In Char,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromUI2 func(
		uiIn USHORT,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromUI4 func(
		ulIn ULONG,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)
	VarBstrFromUI8 func(
		ui64In ULONG64,
		lcid LCID,
		Flags Unsigned_long,
		pbstrOut *BSTR)
	VarBstrFromDec func(
		pdecIn *DECIMAL,
		lcid LCID,
		Flags ULONG,
		pbstrOut *BSTR)

	VarBoolFromUI1 func(
		bIn BYTE,
		pboolOut *VARIANT_BOOL)
	VarBoolFromI2 func(
		sIn SHORT,
		pboolOut *VARIANT_BOOL)
	VarBoolFromI4 func(
		lIn LONG,
		pboolOut *VARIANT_BOOL)
	VarBoolFromI8 func(
		i64In LONG64,
		pboolOut *VARIANT_BOOL)
	VarBoolFromR4 func(
		fltIn FLOAT,
		pboolOut *VARIANT_BOOL)
	VarBoolFromR8 func(
		dblIn DOUBLE,
		pboolOut *VARIANT_BOOL)
	VarBoolFromDate func(
		dateIn DATE,
		pboolOut *VARIANT_BOOL)
	VarBoolFromCy func(
		In CY,
		pboolOut *VARIANT_BOOL)
	VarBoolFromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		pboolOut *VARIANT_BOOL)
	VarBoolFromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		pboolOut *VARIANT_BOOL)
	VarBoolFromI1 func(
		In Char,
		pboolOut *VARIANT_BOOL)
	VarBoolFromUI2 func(
		uiIn USHORT,
		pboolOut *VARIANT_BOOL)
	VarBoolFromUI4 func(
		ulIn ULONG,
		pboolOut *VARIANT_BOOL)
	VarBoolFromUI8 func(
		i64In ULONG64,
		pboolOut *VARIANT_BOOL)
	VarBoolFromDec func(
		pdecIn *DECIMAL,
		pboolOut *VARIANT_BOOL)

	VarI1FromUI1 func(
		bIn BYTE,
		pcOut *Char)

	VarI1FromI2 func(
		uiIn SHORT,
		pcOut *Char)

	VarI1FromI4 func(
		lIn LONG,
		pcOut *Char)

	VarI1FromI8 func(
		i64In LONG64,
		pcOut *Char)

	VarI1FromR4 func(
		fltIn FLOAT,
		pcOut *Char)

	VarI1FromR8 func(
		dblIn DOUBLE,
		pcOut *Char)

	VarI1FromDate func(
		dateIn DATE,
		pcOut *Char)

	VarI1FromCy func(
		In CY,
		pcOut *Char)

	VarI1FromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		pcOut *Char)

	VarI1FromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		pcOut *Char)

	VarI1FromBool func(
		boolIn VARIANT_BOOL,
		pcOut *Char)

	VarI1FromUI2 func(
		uiIn USHORT,
		pcOut *Char)

	VarI1FromUI4 func(
		ulIn ULONG,
		pcOut *Char)

	VarI1FromUI8 func(
		i64In ULONG64,
		pcOut *Char)

	VarI1FromDec func(
		pdecIn *DECIMAL,
		pcOut *Char)

	VarUI2FromUI1 func(
		bIn BYTE,
		puiOut *USHORT)
	VarUI2FromI2 func(
		uiIn SHORT,
		puiOut *USHORT)
	VarUI2FromI4 func(
		lIn LONG,
		puiOut *USHORT)
	VarUI2FromI8 func(
		i64In LONG64,
		puiOut *USHORT)
	VarUI2FromR4 func(
		fltIn FLOAT,
		puiOut *USHORT)
	VarUI2FromR8 func(
		dblIn DOUBLE,
		puiOut *USHORT)
	VarUI2FromDate func(
		dateIn DATE,
		puiOut *USHORT)
	VarUI2FromCy func(
		In CY,
		puiOut *USHORT)
	VarUI2FromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		puiOut *USHORT)
	VarUI2FromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		puiOut *USHORT)
	VarUI2FromBool func(
		boolIn VARIANT_BOOL,
		puiOut *USHORT)
	VarUI2FromI1 func(
		In Char,
		puiOut *USHORT)
	VarUI2FromUI4 func(
		ulIn ULONG,
		puiOut *USHORT)
	VarUI2FromUI8 func(
		i64In ULONG64,
		puiOut *USHORT)
	VarUI2FromDec func(
		pdecIn *DECIMAL,
		puiOut *USHORT)

	VarUI4FromUI1 func(
		bIn BYTE,
		pulOut *ULONG)
	VarUI4FromI2 func(
		uiIn SHORT,
		pulOut *ULONG)
	VarUI4FromI4 func(
		lIn LONG,
		pulOut *ULONG)
	VarUI4FromI8 func(
		i64In LONG64,
		plOut *ULONG)
	VarUI4FromR4 func(
		fltIn FLOAT,
		pulOut *ULONG)
	VarUI4FromR8 func(
		dblIn DOUBLE,
		pulOut *ULONG)
	VarUI4FromDate func(
		dateIn DATE,
		pulOut *ULONG)
	VarUI4FromCy func(
		In CY,
		pulOut *ULONG)
	VarUI4FromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		pulOut *ULONG)
	VarUI4FromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		pulOut *ULONG)
	VarUI4FromBool func(
		boolIn VARIANT_BOOL,
		pulOut *ULONG)
	VarUI4FromI1 func(
		In Char,
		pulOut *ULONG)
	VarUI4FromUI2 func(
		uiIn USHORT,
		pulOut *ULONG)
	VarUI4FromUI8 func(
		ui64In ULONG64,
		plOut *ULONG)
	VarUI4FromDec func(
		pdecIn *DECIMAL,
		pulOut *ULONG)

	VarUI8FromUI1 func(
		bIn BYTE,
		pi64Out *ULONG64)
	VarUI8FromI2 func(
		sIn SHORT,
		pi64Out *ULONG64)
	VarUI8FromI4 func(
		lIn LONG,
		pi64Out *ULONG64)
	VarUI8FromI8 func(
		ui64In LONG64,
		pi64Out *ULONG64)
	VarUI8FromR4 func(
		fltIn FLOAT,
		pi64Out *ULONG64)
	VarUI8FromR8 func(
		dblIn DOUBLE,
		pi64Out *ULONG64)
	VarUI8FromCy func(
		In CY,
		pi64Out *ULONG64)
	VarUI8FromDate func(
		dateIn DATE,
		pi64Out *ULONG64)
	VarUI8FromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags Unsigned_long,
		pi64Out *ULONG64)
	VarUI8FromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		pi64Out *ULONG64)
	VarUI8FromBool func(
		boolIn VARIANT_BOOL,
		pi64Out *ULONG64)
	VarUI8FromI1 func(
		In Char,
		pi64Out *ULONG64)
	VarUI8FromUI2 func(
		uiIn USHORT,
		pi64Out *ULONG64)
	VarUI8FromUI4 func(
		ulIn ULONG,
		pi64Out *ULONG64)
	VarUI8FromDec func(
		pdecIn *DECIMAL,
		pi64Out *ULONG64)
	VarUI8FromInt func(
		intIn INT,
		pi64Out *ULONG64)

	VarDecFromUI1 func(
		bIn BYTE,
		pdecOut *DECIMAL)
	VarDecFromI2 func(
		uiIn SHORT,
		pdecOut *DECIMAL)
	VarDecFromI4 func(
		lIn LONG,
		pdecOut *DECIMAL)
	VarDecFromI8 func(
		i64In LONG64,
		pdecOut *DECIMAL)
	VarDecFromR4 func(
		fltIn FLOAT,
		pdecOut *DECIMAL)
	VarDecFromR8 func(
		dblIn DOUBLE,
		pdecOut *DECIMAL)
	VarDecFromDate func(
		dateIn DATE,
		pdecOut *DECIMAL)
	VarDecFromCy func(
		In CY,
		pdecOut *DECIMAL)
	VarDecFromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		pdecOut *DECIMAL)
	VarDecFromDisp func(
		pdispIn *IDispatch,
		lcid LCID,
		pdecOut *DECIMAL)
	VarDecFromBool func(
		boolIn VARIANT_BOOL,
		pdecOut *DECIMAL)
	VarDecFromI1 func(
		In Char,
		pdecOut *DECIMAL)
	VarDecFromUI2 func(
		uiIn USHORT,
		pdecOut *DECIMAL)
	VarDecFromUI4 func(
		ulIn ULONG,
		pdecOut *DECIMAL)
	VarDecFromUI8 func(
		ui64In ULONG64, pdecOut *DECIMAL)

	VarParseNumFromStr func(
		strIn *OLECHAR,
		lcid LCID,
		Flags ULONG,
		pnumprs *NUMPARSE,
		rgbDig *BYTE)

	VarNumFromParseNum func(
		pnumprs *NUMPARSE,
		rgbDig *BYTE,
		VtBits ULONG,
		pvar *VARIANT)

	VarAdd func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarAnd func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarCat func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarDiv func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarEqv func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarIdiv func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarImp func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarMod func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarMul func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarOr func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarPow func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarSub func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)
	VarXor func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		pvarResult *VARIANT)

	VarAbs func(
		pvarIn *VARIANT,
		pvarResult *VARIANT)
	VarFix func(
		pvarIn *VARIANT,
		pvarResult *VARIANT)
	VarInt func(
		pvarIn *VARIANT,
		pvarResult *VARIANT)
	VarNeg func(
		pvarIn *VARIANT,
		pvarResult *VARIANT)
	VarNot func(
		pvarIn *VARIANT,
		pvarResult *VARIANT)

	VarRound func(
		pvarIn *VARIANT,
		Decimals int,
		pvarResult *VARIANT)

	VarCmp func(
		pvarLeft *VARIANT,
		pvarRight *VARIANT,
		lcid LCID,
		Flags ULONG)

	VarDecAdd func(
		pdecLeft *DECIMAL,
		pdecRight *DECIMAL,
		pdecResult *DECIMAL)
	VarDecDiv func(
		pdecLeft *DECIMAL,
		pdecRight *DECIMAL,
		pdecResult *DECIMAL)
	VarDecMul func(
		pdecLeft *DECIMAL,
		pdecRight *DECIMAL,
		pdecResult *DECIMAL)
	VarDecSub func(
		pdecLeft *DECIMAL,
		pdecRight *DECIMAL,
		pdecResult *DECIMAL)

	VarDecAbs func(
		pdecIn *DECIMAL,
		pdecResult *DECIMAL)
	VarDecFix func(
		pdecIn *DECIMAL,
		pdecResult *DECIMAL)
	VarDecInt func(
		pdecIn *DECIMAL,
		pdecResult *DECIMAL)
	VarDecNeg func(
		pdecIn *DECIMAL,
		pdecResult *DECIMAL)

	VarDecRound func(
		pdecIn *DECIMAL,
		Decimals int,
		pdecResult *DECIMAL)

	VarDecCmp func(
		pdecLeft *DECIMAL,
		pdecRight *DECIMAL)
	VarDecCmpR8 func(
		pdecLeft *DECIMAL,
		dblRight DOUBLE)

	VarCyAdd func(
		Left CY,
		Right CY,
		pcyResult *CY)
	VarCyMul func(
		Left CY,
		Right CY,
		pcyResult *CY)
	VarCyMulI4 func(
		Left CY,
		lRight LONG,
		pcyResult *CY)
	VarCyMulI8 func(
		Left CY,
		lRight LONG64,
		pcyResult *CY)
	VarCySub func(
		Left CY,
		Right CY,
		pcyResult *CY)

	VarCyAbs func(
		In CY,
		pcyResult *CY)
	VarCyFix func(
		In CY,
		pcyResult *CY)
	VarCyInt func(
		In CY,
		pcyResult *CY)
	VarCyNeg func(
		In CY,
		pcyResult *CY)

	VarCyRound func(
		In CY,
		Decimals int,
		pcyResult *CY)

	VarCyCmp func(
		Left CY,
		Right CY)
	VarCyCmpR8 func(
		Left CY,
		dblRight DOUBLE)

	VarBstrCat func(
		bstrLeft BSTR,
		bstrRight BSTR,
		pbstrResult *BSTR)
	VarBstrCmp func(
		bstrLeft BSTR,
		bstrRight BSTR,
		lcid LCID,
		Flags ULONG)
	VarR8Pow func(
		dblLeft DOUBLE,
		dblRight DOUBLE,
		pdblResult *DOUBLE)
	VarR4CmpR8 func(
		fltLeft FLOAT,
		dblRight DOUBLE)
	VarR8Round func(
		dblIn DOUBLE,
		Decimals int,
		pdblResult *DOUBLE)

	VarDateFromUdate func(
		pudateIn *UDATE,
		Flags ULONG,
		pdateOut *DATE)
	VarDateFromUdateEx func(
		pudateIn *UDATE,
		lcid LCID,
		Flags ULONG,
		pdateOut *DATE)
	VarUdateFromDate func(
		dateIn DATE,
		Flags ULONG,
		pudateOut *UDATE)

	GetAltMonthNames func(
		lcid LCID,
		prgp ***OLESTR)

	VarFormat func(
		pvarIn *VARIANT,
		pstrFormat *OLESTR,
		iFirstDay int,
		iFirstWeek int,
		Flags ULONG,
		pbstrOut *BSTR)
	VarFormatDateTime func(
		pvarIn *VARIANT,
		iNamedFormat int,
		Flags ULONG,
		pbstrOut *BSTR)
	VarFormatNumber func(
		pvarIn *VARIANT,
		iNumDig int,
		iIncLead int,
		iUseParens int,
		iGroup int,
		Flags ULONG,
		pbstrOut *BSTR)
	VarFormatPercent func(
		pvarIn *VARIANT,
		iNumDig int,
		iIncLead int,
		iUseParens int,
		iGroup int,
		Flags ULONG,
		pbstrOut *BSTR)
	VarFormatCurrency func(
		pvarIn *VARIANT,
		iNumDig int,
		iIncLead int,
		iUseParens int,
		iGroup int,
		Flags ULONG,
		pbstrOut *BSTR)

	VarWeekdayName func(
		iWeekday int,
		fAbbrev int,
		iFirstDay int,
		Flags ULONG,
		pbstrOut *BSTR)
	VarMonthName func(
		iMonth int,
		fAbbrev int,
		Flags ULONG,
		pbstrOut *BSTR)

	VarFormatFromTokens func(
		pvarIn *VARIANT,
		pstrFormat *OLESTR,
		pbTokCur *BYTE,
		Flags ULONG,
		pbstrOut *BSTR,
		lcid LCID)
	VarTokenizeFormatString func(
		pstrFormat *OLESTR,
		rgbTok *BYTE,
		cbTok int,
		iFirstDay int,
		iFirstWeek int,
		lcid LCID,
		pcbActual *int)

	LHashValOfNameSysA func(
		syskind SYSKIND,
		lcid LCID,
		szName AString) ULONG

	LHashValOfNameSys func(
		syskind SYSKIND,
		lcid LCID,
		szName *OLECHAR) ULONG

	LoadTypeLib func(
		szFile *OLECHAR,
		pptlib **ITypeLib)

	LoadTypeLibEx func(
		szFile *OLESTR,
		regkind REGKIND,
		pptlib **ITypeLib)

	LoadRegTypeLib func(
		rguid REFGUID,
		VerMajor WORD,
		VerMinor WORD,
		lcid LCID,
		pptlib **ITypeLib)

	QueryPathOfRegTypeLib func(
		guid REFGUID,
		Maj USHORT,
		Min USHORT,
		lcid LCID,
		PathName *BSTR)

	RegisterTypeLib func(
		ptlib *ITypeLib,
		szFullPath *OLECHAR,
		szHelpDir *OLECHAR)

	UnRegisterTypeLib func(
		libID REFGUID,
		wVerMajor WORD,
		wVerMinor WORD,
		lcid LCID,
		syskind SYSKIND)

	CreateTypeLib func(
		syskind SYSKIND,
		szFile *OLECHAR,
		ppctlib **ICreateTypeLib)

	CreateTypeLib2 func(
		syskind SYSKIND,
		szFile *OLESTR,
		ppctlib **ICreateTypeLib2)

	DispGetParam func(
		pdispparams *DISPPARAMS,
		position UINT,
		vtTarg VARTYPE,
		pvarResult *VARIANT,
		puArgErr *UINT)

	DispGetIDsOfNames func(
		ptinfo *ITypeInfo,
		rgszNames **OLECHAR,
		Names UINT,
		rgdispid *DISPID)

	DispInvoke func(
		_this *VOID,
		ptinfo *ITypeInfo,
		dispidMember DISPID,
		wFlags WORD,
		pparams *DISPPARAMS,
		pvarResult *VARIANT,
		pexcepinfo *EXCEPINFO,
		puArgErr *UINT)

	CreateDispTypeInfo func(
		pidata *INTERFACEDATA,
		lcid LCID,
		pptinfo **ITypeInfo)

	CreateStdDispatch func(
		Outer *IUnknown,
		pvThis *VOID,
		ptinfo *ITypeInfo,
		ppunkStdDisp **IUnknown)

	DispCallFunc func(
		pvInstance *VOID,
		oVft ULONG_PTR,
		cc CALLCONV,
		vtReturn VARTYPE,
		Actuals UINT,
		prgvt *VARTYPE,
		prgpvarg **VARIANTARG,
		Result *VARIANT)

	RegisterActiveObject func(
		Unk *IUnknown,
		rclsid REFCLSID,
		Flags DWORD,
		pdwRegister *DWORD)

	RevokeActiveObject func(
		Register DWORD,
		pvReserved *VOID)

	GetActiveObject func(
		rclsid REFCLSID,
		pvReserved *VOID,
		ppunk **IUnknown)

	SetErrorInfo func(
		Reserved ULONG,
		perrinfo *IErrorInfo)
	GetErrorInfo func(
		Reserved ULONG,
		pperrinfo **IErrorInfo)
	CreateErrorInfo func(
		pperrinfo **ICreateErrorInfo)

	GetRecordInfoFromTypeInfo func(
		TypeInfo *ITypeInfo,
		RecInfo **IRecordInfo)

	GetRecordInfoFromGuids func(
		rGuidTypeLib REFGUID,
		uVerMajor ULONG,
		uVerMinor ULONG,
		lcid LCID,
		rGuidTypeInfo REFGUID,
		RecInfo **IRecordInfo)

	OaBuildVersion func() ULONG

	ClearCustData func(custData *CUSTDATA)
)

func VarI4FromI4(in int, out *LONG) { *out = LONG(in) }

//VarI4FromI8 func(LONG64 i64In, LONG *plOut) DUP
//VarI4FromUI8 func(ULONG64 ui64In, LONG *plOut) DUP

//#define VarUI4FromUI4(in, pOut) (*(pOut) = (in))

//#define VarUI8FromUI8(in, pOut) (*(pOut) = (in))
//#define VarI8FromI8(in, pOut)   (*(pOut) = (in))

/*
   #define VarUI1FromInt       VarUI1FromI4
   #define VarUI1FromUint      VarUI1FromUI4
   #define VarI2FromInt        VarI2FromI4
   #define VarI2FromUint       VarI2FromUI4
   #define VarI4FromUint       VarI4FromUI4
   #define VarI8FromInt        VarI8FromI4
   #define VarI8FromUint       VarI8FromUI4
   #define VarR4FromInt        VarR4FromI4
   #define VarR4FromUint       VarR4FromUI4
   #define VarR8FromInt        VarR8FromI4
   #define VarR8FromUint       VarR8FromUI4
   #define VarDateFromInt      VarDateFromI4
   #define VarDateFromUint     VarDateFromUI4
   #define VarCyFromInt        VarCyFromI4
   #define VarCyFromUint       VarCyFromUI4
   #define VarBstrFromInt      VarBstrFromI4
   #define VarBstrFromUint     VarBstrFromUI4
   #define VarBoolFromInt      VarBoolFromI4
   #define VarBoolFromUint     VarBoolFromUI4
   #define VarI1FromInt        VarI1FromI4
   #define VarI1FromUint       VarI1FromUI4
   #define VarUI2FromInt       VarUI2FromI4
   #define VarUI2FromUint      VarUI2FromUI4
   #define VarUI4FromInt       VarUI4FromI4
   #define VarUI4FromUint      VarUI4FromUI4
   #define VarDecFromInt       VarDecFromI4
   #define VarDecFromUint      VarDecFromUI4
   #define VarIntFromUI1       VarI4FromUI1
   #define VarIntFromI2        VarI4FromI2
   #define VarIntFromI4        VarI4FromI4
   #define VarIntFromI8        VarI4FromI8
   #define VarIntFromR4        VarI4FromR4
   #define VarIntFromR8        VarI4FromR8
   #define VarIntFromDate      VarI4FromDate
   #define VarIntFromCy        VarI4FromCy
   #define VarIntFromStr       VarI4FromStr
   #define VarIntFromDisp      VarI4FromDisp
   #define VarIntFromBool      VarI4FromBool
   #define VarIntFromI1        VarI4FromI1
   #define VarIntFromUI2       VarI4FromUI2
   #define VarIntFromUI4       VarI4FromUI4
   #define VarIntFromUI8       VarI4FromUI8
   #define VarIntFromDec       VarI4FromDec
   #define VarIntFromUint      VarI4FromUI4
   #define VarUintFromUI1      VarUI4FromUI1
   #define VarUintFromI2       VarUI4FromI2
   #define VarUintFromI4       VarUI4FromI4
   #define VarUintFromI8       VarUI4FromI8
   #define VarUintFromR4       VarUI4FromR4
   #define VarUintFromR8       VarUI4FromR8
   #define VarUintFromDate     VarUI4FromDate
   #define VarUintFromCy       VarUI4FromCy
   #define VarUintFromStr      VarUI4FromStr
   #define VarUintFromDisp     VarUI4FromDisp
   #define VarUintFromBool     VarUI4FromBool
   #define VarUintFromI1       VarUI4FromI1
   #define VarUintFromUI2      VarUI4FromUI2
   #define VarUintFromUI4      VarUI4FromUI4
   #define VarUintFromUI8      VarUI4FromUI8
   #define VarUintFromDec      VarUI4FromDec
   #define VarUintFromInt      VarUI4FromI4
*/

/*
   #define LHashValOfName(lcid, szName) \
               LHashValOfNameSys(SYS_WIN32, lcid, szName)

   #define WHashValOfLHashVal(lhashval) \
               ((USHORT) (0x0000ffff & (lhashval)))

   #define IsHashValCompatible(lhashval1, lhashval2) \
               ((BOOL) ((0x00ff0000 & (lhashval1)) == (0x00ff0000 & (lhashval2))))
*/

//TODO(t): Normal funcs
//{"VarI4FromInt", VarI4FromI4},

//TODO(t): Only in header No implementation
//{"VarI8FromI4", &VarI8FromI4},
//{"VarI8FromInt", &VarI8FromI4},
//{"VarUI8FromI4", &VarUI8FromI4},
//{"VarUI8FromInt", &VarUI8FromInt},

var OleAutoApis = Apis{
	{"BstrFromVector", &BstrFromVector},
	{"ClearCustData", &ClearCustData},
	{"CreateDispTypeInfo", &CreateDispTypeInfo},
	{"CreateErrorInfo", &CreateErrorInfo},
	{"CreateStdDispatch", &CreateStdDispatch},
	{"CreateTypeLib", &CreateTypeLib},
	{"CreateTypeLib2", &CreateTypeLib2},
	{"DispCallFunc", &DispCallFunc},
	{"DispGetIDsOfNames", &DispGetIDsOfNames},
	{"DispGetParam", &DispGetParam},
	{"DispInvoke", &DispInvoke},
	{"DosDateTimeToVariantTime", &DosDateTimeToVariantTime},
	{"GetActiveObject", &GetActiveObject},
	{"GetAltMonthNames", &GetAltMonthNames},
	{"GetErrorInfo", &GetErrorInfo},
	{"GetRecordInfoFromGuids", &GetRecordInfoFromGuids},
	{"GetRecordInfoFromTypeInfo", &GetRecordInfoFromTypeInfo},
	{"LHashValOfNameSys", &LHashValOfNameSys},
	{"LHashValOfNameSysA", &LHashValOfNameSysA},
	{"LoadRegTypeLib", &LoadRegTypeLib},
	{"LoadTypeLib", &LoadTypeLib},
	{"LoadTypeLibEx", &LoadTypeLibEx},
	{"OaBuildVersion", &OaBuildVersion},
	{"QueryPathOfRegTypeLib", &QueryPathOfRegTypeLib},
	{"RegisterActiveObject", &RegisterActiveObject},
	{"RegisterTypeLib", &RegisterTypeLib},
	{"RevokeActiveObject", &RevokeActiveObject},
	{"SafeArrayAccessData", &SafeArrayAccessData},
	{"SafeArrayAllocData", &SafeArrayAllocData},
	{"SafeArrayAllocDescriptor", &SafeArrayAllocDescriptor},
	{"SafeArrayAllocDescriptorEx", &SafeArrayAllocDescriptorEx},
	{"SafeArrayCopy", &SafeArrayCopy},
	{"SafeArrayCopyData", &SafeArrayCopyData},
	{"SafeArrayCreate", &SafeArrayCreate},
	{"SafeArrayCreateEx", &SafeArrayCreateEx},
	{"SafeArrayCreateVector", &SafeArrayCreateVector},
	{"SafeArrayCreateVectorEx", &SafeArrayCreateVectorEx},
	{"SafeArrayDestroy", &SafeArrayDestroy},
	{"SafeArrayDestroyData", &SafeArrayDestroyData},
	{"SafeArrayDestroyDescriptor", &SafeArrayDestroyDescriptor},
	{"SafeArrayGetDim", &SafeArrayGetDim},
	{"SafeArrayGetElement", &SafeArrayGetElement},
	{"SafeArrayGetElemsize", &SafeArrayGetElemsize},
	{"SafeArrayGetIID", &SafeArrayGetIID},
	{"SafeArrayGetLBound", &SafeArrayGetLBound},
	{"SafeArrayGetRecordInfo", &SafeArrayGetRecordInfo},
	{"SafeArrayGetUBound", &SafeArrayGetUBound},
	{"SafeArrayGetVartype", &SafeArrayGetVartype},
	{"SafeArrayLock", &SafeArrayLock},
	{"SafeArrayPtrOfIndex", &SafeArrayPtrOfIndex},
	{"SafeArrayPutElement", &SafeArrayPutElement},
	{"SafeArrayRedim", &SafeArrayRedim},
	{"SafeArraySetIID", &SafeArraySetIID},
	{"SafeArraySetRecordInfo", &SafeArraySetRecordInfo},
	{"SafeArrayUnaccessData", &SafeArrayUnaccessData},
	{"SafeArrayUnlock", &SafeArrayUnlock},
	{"SetErrorInfo", &SetErrorInfo},
	{"SysAllocString", &SysAllocString},
	{"SysAllocStringByteLen", &SysAllocStringByteLen},
	{"SysAllocStringLen", &SysAllocStringLen},
	{"SysFreeString", &SysFreeString},
	{"SysReAllocString", &SysReAllocString},
	{"SysReAllocStringLen", &SysReAllocStringLen},
	{"SysStringByteLen", &SysStringByteLen},
	{"SysStringLen", &SysStringLen},
	{"SystemTimeToVariantTime", &SystemTimeToVariantTime},
	{"UnRegisterTypeLib", &UnRegisterTypeLib},
	{"VarAbs", &VarAbs},
	{"VarAdd", &VarAdd},
	{"VarAnd", &VarAnd},
	{"VarBoolFromCy", &VarBoolFromCy},
	{"VarBoolFromDate", &VarBoolFromDate},
	{"VarBoolFromDec", &VarBoolFromDec},
	{"VarBoolFromDisp", &VarBoolFromDisp},
	{"VarBoolFromI1", &VarBoolFromI1},
	{"VarBoolFromI2", &VarBoolFromI2},
	{"VarBoolFromI4", &VarBoolFromI4},
	{"VarBoolFromI8", &VarBoolFromI8},
	{"VarBoolFromR4", &VarBoolFromR4},
	{"VarBoolFromR8", &VarBoolFromR8},
	{"VarBoolFromStr", &VarBoolFromStr},
	{"VarBoolFromUI1", &VarBoolFromUI1},
	{"VarBoolFromUI2", &VarBoolFromUI2},
	{"VarBoolFromUI4", &VarBoolFromUI4},
	{"VarBoolFromUI8", &VarBoolFromUI8},
	{"VarBstrCat", &VarBstrCat},
	{"VarBstrCmp", &VarBstrCmp},
	{"VarBstrFromBool", &VarBstrFromBool},
	{"VarBstrFromCy", &VarBstrFromCy},
	{"VarBstrFromDate", &VarBstrFromDate},
	{"VarBstrFromDec", &VarBstrFromDec},
	{"VarBstrFromDisp", &VarBstrFromDisp},
	{"VarBstrFromI1", &VarBstrFromI1},
	{"VarBstrFromI2", &VarBstrFromI2},
	{"VarBstrFromI4", &VarBstrFromI4},
	{"VarBstrFromI8", &VarBstrFromI8},
	{"VarBstrFromR4", &VarBstrFromR4},
	{"VarBstrFromR8", &VarBstrFromR8},
	{"VarBstrFromUI1", &VarBstrFromUI1},
	{"VarBstrFromUI2", &VarBstrFromUI2},
	{"VarBstrFromUI4", &VarBstrFromUI4},
	{"VarBstrFromUI8", &VarBstrFromUI8},
	{"VarCat", &VarCat},
	{"VarCmp", &VarCmp},
	{"VarCyAbs", &VarCyAbs},
	{"VarCyAdd", &VarCyAdd},
	{"VarCyCmp", &VarCyCmp},
	{"VarCyCmpR8", &VarCyCmpR8},
	{"VarCyFix", &VarCyFix},
	{"VarCyFromBool", &VarCyFromBool},
	{"VarCyFromDate", &VarCyFromDate},
	{"VarCyFromDec", &VarCyFromDec},
	{"VarCyFromDisp", &VarCyFromDisp},
	{"VarCyFromI1", &VarCyFromI1},
	{"VarCyFromI2", &VarCyFromI2},
	{"VarCyFromI4", &VarCyFromI4},
	{"VarCyFromI8", &VarCyFromI8},
	{"VarCyFromR4", &VarCyFromR4},
	{"VarCyFromR8", &VarCyFromR8},
	{"VarCyFromStr", &VarCyFromStr},
	{"VarCyFromUI1", &VarCyFromUI1},
	{"VarCyFromUI2", &VarCyFromUI2},
	{"VarCyFromUI4", &VarCyFromUI4},
	{"VarCyFromUI8", &VarCyFromUI8},
	{"VarCyInt", &VarCyInt},
	{"VarCyMul", &VarCyMul},
	{"VarCyMulI4", &VarCyMulI4},
	{"VarCyMulI8", &VarCyMulI8},
	{"VarCyNeg", &VarCyNeg},
	{"VarCyRound", &VarCyRound},
	{"VarCySub", &VarCySub},
	{"VarDateFromBool", &VarDateFromBool},
	{"VarDateFromCy", &VarDateFromCy},
	{"VarDateFromDec", &VarDateFromDec},
	{"VarDateFromDisp", &VarDateFromDisp},
	{"VarDateFromI1", &VarDateFromI1},
	{"VarDateFromI2", &VarDateFromI2},
	{"VarDateFromI4", &VarDateFromI4},
	{"VarDateFromI8", &VarDateFromI8},
	{"VarDateFromR4", &VarDateFromR4},
	{"VarDateFromR8", &VarDateFromR8},
	{"VarDateFromStr", &VarDateFromStr},
	{"VarDateFromUdate", &VarDateFromUdate},
	{"VarDateFromUdateEx", &VarDateFromUdateEx},
	{"VarDateFromUI1", &VarDateFromUI1},
	{"VarDateFromUI2", &VarDateFromUI2},
	{"VarDateFromUI4", &VarDateFromUI4},
	{"VarDateFromUI8", &VarDateFromUI8},
	{"VarDecAbs", &VarDecAbs},
	{"VarDecAdd", &VarDecAdd},
	{"VarDecCmp", &VarDecCmp},
	{"VarDecCmpR8", &VarDecCmpR8},
	{"VarDecDiv", &VarDecDiv},
	{"VarDecFix", &VarDecFix},
	{"VarDecFromBool", &VarDecFromBool},
	{"VarDecFromCy", &VarDecFromCy},
	{"VarDecFromDate", &VarDecFromDate},
	{"VarDecFromDisp", &VarDecFromDisp},
	{"VarDecFromI1", &VarDecFromI1},
	{"VarDecFromI2", &VarDecFromI2},
	{"VarDecFromI4", &VarDecFromI4},
	{"VarDecFromI8", &VarDecFromI8},
	{"VarDecFromR4", &VarDecFromR4},
	{"VarDecFromR8", &VarDecFromR8},
	{"VarDecFromStr", &VarDecFromStr},
	{"VarDecFromUI1", &VarDecFromUI1},
	{"VarDecFromUI2", &VarDecFromUI2},
	{"VarDecFromUI4", &VarDecFromUI4},
	{"VarDecFromUI8", &VarDecFromUI8},
	{"VarDecInt", &VarDecInt},
	{"VarDecMul", &VarDecMul},
	{"VarDecNeg", &VarDecNeg},
	{"VarDecRound", &VarDecRound},
	{"VarDecSub", &VarDecSub},
	{"VarDiv", &VarDiv},
	{"VarEqv", &VarEqv},
	{"VarFix", &VarFix},
	{"VarFormat", &VarFormat},
	{"VarFormatCurrency", &VarFormatCurrency},
	{"VarFormatDateTime", &VarFormatDateTime},
	{"VarFormatFromTokens", &VarFormatFromTokens},
	{"VarFormatNumber", &VarFormatNumber},
	{"VarFormatPercent", &VarFormatPercent},
	{"VarI1FromBool", &VarI1FromBool},
	{"VarI1FromCy", &VarI1FromCy},
	{"VarI1FromDate", &VarI1FromDate},
	{"VarI1FromDec", &VarI1FromDec},
	{"VarI1FromDisp", &VarI1FromDisp},
	{"VarI1FromI2", &VarI1FromI2},
	{"VarI1FromI4", &VarI1FromI4},
	{"VarI1FromI8", &VarI1FromI8},
	{"VarI1FromR4", &VarI1FromR4},
	{"VarI1FromR8", &VarI1FromR8},
	{"VarI1FromStr", &VarI1FromStr},
	{"VarI1FromUI1", &VarI1FromUI1},
	{"VarI1FromUI2", &VarI1FromUI2},
	{"VarI1FromUI4", &VarI1FromUI4},
	{"VarI1FromUI8", &VarI1FromUI8},
	{"VarI2FromBool", &VarI2FromBool},
	{"VarI2FromCy", &VarI2FromCy},
	{"VarI2FromDate", &VarI2FromDate},
	{"VarI2FromDec", &VarI2FromDec},
	{"VarI2FromDisp", &VarI2FromDisp},
	{"VarI2FromI1", &VarI2FromI1},
	{"VarI2FromI4", &VarI2FromI4},
	{"VarI2FromI8", &VarI2FromI8},
	{"VarI2FromR4", &VarI2FromR4},
	{"VarI2FromR8", &VarI2FromR8},
	{"VarI2FromStr", &VarI2FromStr},
	{"VarI2FromUI1", &VarI2FromUI1},
	{"VarI2FromUI2", &VarI2FromUI2},
	{"VarI2FromUI4", &VarI2FromUI4},
	{"VarI2FromUI8", &VarI2FromUI8},
	{"VarI4FromBool", &VarI4FromBool},
	{"VarI4FromCy", &VarI4FromCy},
	{"VarI4FromDate", &VarI4FromDate},
	{"VarI4FromDec", &VarI4FromDec},
	{"VarI4FromDisp", &VarI4FromDisp},
	{"VarI4FromI1", &VarI4FromI1},
	{"VarI4FromI2", &VarI4FromI2},
	{"VarI4FromI8", &VarI4FromI8},
	{"VarI4FromR4", &VarI4FromR4},
	{"VarI4FromR8", &VarI4FromR8},
	{"VarI4FromStr", &VarI4FromStr},
	{"VarI4FromUI1", &VarI4FromUI1},
	{"VarI4FromUI2", &VarI4FromUI2},
	{"VarI4FromUI4", &VarI4FromUI4},
	{"VarI4FromUI8", &VarI4FromUI8},
	{"VarI8FromBool", &VarI8FromBool},
	{"VarI8FromCy", &VarI8FromCy},
	{"VarI8FromDate", &VarI8FromDate},
	{"VarI8FromDec", &VarI8FromDec},
	{"VarI8FromDisp", &VarI8FromDisp},
	{"VarI8FromI1", &VarI8FromI1},
	{"VarI8FromI2", &VarI8FromI2},
	{"VarI8FromR4", &VarI8FromR4},
	{"VarI8FromR8", &VarI8FromR8},
	{"VarI8FromStr", &VarI8FromStr},
	{"VarI8FromUI1", &VarI8FromUI1},
	{"VarI8FromUI2", &VarI8FromUI2},
	{"VarI8FromUI4", &VarI8FromUI4},
	{"VarI8FromUI8", &VarI8FromUI8},
	{"VariantChangeType", &VariantChangeType},
	{"VariantChangeTypeEx", &VariantChangeTypeEx},
	{"VariantClear", &VariantClear},
	{"VariantCopy", &VariantCopy},
	{"VariantCopyInd", &VariantCopyInd},
	{"VariantInit", &VariantInit},
	{"VariantTimeToDosDateTime", &VariantTimeToDosDateTime},
	{"VariantTimeToSystemTime", &VariantTimeToSystemTime},
	{"VarIdiv", &VarIdiv},
	{"VarImp", &VarImp},
	{"VarInt", &VarInt},
	{"VarMod", &VarMod},
	{"VarMonthName", &VarMonthName},
	{"VarMul", &VarMul},
	{"VarNeg", &VarNeg},
	{"VarNot", &VarNot},
	{"VarNumFromParseNum", &VarNumFromParseNum},
	{"VarOr", &VarOr},
	{"VarParseNumFromStr", &VarParseNumFromStr},
	{"VarPow", &VarPow},
	{"VarR4CmpR8", &VarR4CmpR8},
	{"VarR4FromBool", &VarR4FromBool},
	{"VarR4FromCy", &VarR4FromCy},
	{"VarR4FromDate", &VarR4FromDate},
	{"VarR4FromDec", &VarR4FromDec},
	{"VarR4FromDisp", &VarR4FromDisp},
	{"VarR4FromI1", &VarR4FromI1},
	{"VarR4FromI2", &VarR4FromI2},
	{"VarR4FromI4", &VarR4FromI4},
	{"VarR4FromI8", &VarR4FromI8},
	{"VarR4FromR8", &VarR4FromR8},
	{"VarR4FromStr", &VarR4FromStr},
	{"VarR4FromUI1", &VarR4FromUI1},
	{"VarR4FromUI2", &VarR4FromUI2},
	{"VarR4FromUI4", &VarR4FromUI4},
	{"VarR4FromUI8", &VarR4FromUI8},
	{"VarR8FromBool", &VarR8FromBool},
	{"VarR8FromCy", &VarR8FromCy},
	{"VarR8FromDate", &VarR8FromDate},
	{"VarR8FromDec", &VarR8FromDec},
	{"VarR8FromDisp", &VarR8FromDisp},
	{"VarR8FromI1", &VarR8FromI1},
	{"VarR8FromI2", &VarR8FromI2},
	{"VarR8FromI4", &VarR8FromI4},
	{"VarR8FromI8", &VarR8FromI8},
	{"VarR8FromR4", &VarR8FromR4},
	{"VarR8FromStr", &VarR8FromStr},
	{"VarR8FromUI1", &VarR8FromUI1},
	{"VarR8FromUI2", &VarR8FromUI2},
	{"VarR8FromUI4", &VarR8FromUI4},
	{"VarR8FromUI8", &VarR8FromUI8},
	{"VarR8Pow", &VarR8Pow},
	{"VarR8Round", &VarR8Round},
	{"VarRound", &VarRound},
	{"VarSub", &VarSub},
	{"VarTokenizeFormatString", &VarTokenizeFormatString},
	{"VarUdateFromDate", &VarUdateFromDate},
	{"VarUI1FromBool", &VarUI1FromBool},
	{"VarUI1FromCy", &VarUI1FromCy},
	{"VarUI1FromDate", &VarUI1FromDate},
	{"VarUI1FromDec", &VarUI1FromDec},
	{"VarUI1FromDisp", &VarUI1FromDisp},
	{"VarUI1FromI1", &VarUI1FromI1},
	{"VarUI1FromI2", &VarUI1FromI2},
	{"VarUI1FromI4", &VarUI1FromI4},
	{"VarUI1FromI8", &VarUI1FromI8},
	{"VarUI1FromR4", &VarUI1FromR4},
	{"VarUI1FromR8", &VarUI1FromR8},
	{"VarUI1FromStr", &VarUI1FromStr},
	{"VarUI1FromUI2", &VarUI1FromUI2},
	{"VarUI1FromUI4", &VarUI1FromUI4},
	{"VarUI1FromUI8", &VarUI1FromUI8},
	{"VarUI2FromBool", &VarUI2FromBool},
	{"VarUI2FromCy", &VarUI2FromCy},
	{"VarUI2FromDate", &VarUI2FromDate},
	{"VarUI2FromDec", &VarUI2FromDec},
	{"VarUI2FromDisp", &VarUI2FromDisp},
	{"VarUI2FromI1", &VarUI2FromI1},
	{"VarUI2FromI2", &VarUI2FromI2},
	{"VarUI2FromI4", &VarUI2FromI4},
	{"VarUI2FromI8", &VarUI2FromI8},
	{"VarUI2FromR4", &VarUI2FromR4},
	{"VarUI2FromR8", &VarUI2FromR8},
	{"VarUI2FromStr", &VarUI2FromStr},
	{"VarUI2FromUI1", &VarUI2FromUI1},
	{"VarUI2FromUI4", &VarUI2FromUI4},
	{"VarUI2FromUI8", &VarUI2FromUI8},
	{"VarUI4FromBool", &VarUI4FromBool},
	{"VarUI4FromCy", &VarUI4FromCy},
	{"VarUI4FromDate", &VarUI4FromDate},
	{"VarUI4FromDec", &VarUI4FromDec},
	{"VarUI4FromDisp", &VarUI4FromDisp},
	{"VarUI4FromI1", &VarUI4FromI1},
	{"VarUI4FromI2", &VarUI4FromI2},
	{"VarUI4FromI4", &VarUI4FromI4},
	{"VarUI4FromI8", &VarUI4FromI8},
	{"VarUI4FromR4", &VarUI4FromR4},
	{"VarUI4FromR8", &VarUI4FromR8},
	{"VarUI4FromStr", &VarUI4FromStr},
	{"VarUI4FromUI1", &VarUI4FromUI1},
	{"VarUI4FromUI2", &VarUI4FromUI2},
	{"VarUI4FromUI8", &VarUI4FromUI8},
	{"VarUI8FromBool", &VarUI8FromBool},
	{"VarUI8FromCy", &VarUI8FromCy},
	{"VarUI8FromDate", &VarUI8FromDate},
	{"VarUI8FromDec", &VarUI8FromDec},
	{"VarUI8FromDisp", &VarUI8FromDisp},
	{"VarUI8FromI1", &VarUI8FromI1},
	{"VarUI8FromI2", &VarUI8FromI2},
	{"VarUI8FromI8", &VarUI8FromI8},
	{"VarUI8FromR4", &VarUI8FromR4},
	{"VarUI8FromR8", &VarUI8FromR8},
	{"VarUI8FromStr", &VarUI8FromStr},
	{"VarUI8FromUI1", &VarUI8FromUI1},
	{"VarUI8FromUI2", &VarUI8FromUI2},
	{"VarUI8FromUI4", &VarUI8FromUI4},
	{"VarWeekdayName", &VarWeekdayName},
	{"VarXor", &VarXor},
	{"VectorFromBstr", &VectorFromBstr},
}
