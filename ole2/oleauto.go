// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package ole2 provides API definitions for accessing
//ole32aut32.dll and ole32.dll.
package ole2

import (
	. "github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/oleaut32"
)

type ()

var (
	SysAllocString func(*T.OLECHAR) T.BSTR

	SysReAllocString func(*T.BSTR, *T.OLECHAR) T.INT

	SysAllocStringLen func(*T.OLECHAR, T.UINT) T.BSTR

	SysReAllocStringLen func(*T.BSTR, *T.OLECHAR, T.UINT) T.INT

	SysFreeString func(T.BSTR)

	SysStringLen func(T.BSTR) T.UINT

	SysStringByteLen func(s T.BSTR) T.UINT

	SysAllocStringByteLen func(s T.AString, leng T.UINT) T.BSTR

	DosDateTimeToVariantTime func(
		dosDate, dosTime T.USHORT, time *T.DOUBLE) T.INT

	VariantTimeToDosDateTime func(
		time T.DOUBLE, dosDate, dosTime *T.USHORT) T.INT

	SystemTimeToVariantTime func(
		systemTime *T.SYSTEMTIME, time *T.DOUBLE) T.INT

	VariantTimeToSystemTime func(
		time T.DOUBLE, systemTime *T.SYSTEMTIME) T.INT

	SafeArrayAllocDescriptor func(dims T.UINT, out **T.SAFEARRAY)

	SafeArrayAllocDescriptorEx func(
		vt T.VARTYPE, dims T.UINT, out **T.SAFEARRAY)

	SafeArrayAllocData func(sa *T.SAFEARRAY)

	SafeArrayCreate func(
		vt T.VARTYPE, dims T.UINT,
		saBound *T.SAFEARRAYBOUND) *T.SAFEARRAY

	SafeArrayCreateEx func(
		vt T.VARTYPE,
		dims T.UINT,
		saBound *T.SAFEARRAYBOUND,
		pvExtra *T.VOID) *T.SAFEARRAY

	SafeArrayCopyData func(source, target *T.SAFEARRAY)

	SafeArrayDestroyDescriptor func(sa *T.SAFEARRAY)

	SafeArrayDestroyData func(sa *T.SAFEARRAY)

	SafeArrayDestroy func(sa *T.SAFEARRAY)

	SafeArrayRedim func(
		sa *T.SAFEARRAY, psaboundNew *T.SAFEARRAYBOUND)

	SafeArrayGetDim func(sa *T.SAFEARRAY) T.UINT

	SafeArrayGetElemsize func(sa *T.SAFEARRAY) T.UINT

	SafeArrayGetUBound func(
		sa *T.SAFEARRAY, nDim T.UINT, plUbound *T.LONG)

	SafeArrayGetLBound func(
		sa *T.SAFEARRAY, nDim T.UINT, plLbound *T.LONG)

	SafeArrayLock func(sa *T.SAFEARRAY)

	SafeArrayUnlock func(sa *T.SAFEARRAY)

	SafeArrayAccessData func(sa *T.SAFEARRAY, ppvData **T.VOID)

	SafeArrayUnaccessData func(sa *T.SAFEARRAY)

	SafeArrayGetElement func(
		sa *T.SAFEARRAY, rgIndices *T.LONG, pv *T.VOID)

	SafeArrayPutElement func(
		sa *T.SAFEARRAY, rgIndices *T.LONG, pv *T.VOID)

	SafeArrayCopy func(sa *T.SAFEARRAY, out **T.SAFEARRAY)

	SafeArrayPtrOfIndex func(
		sa *T.SAFEARRAY, rgIndices *T.LONG, ppvData **T.VOID)

	SafeArraySetRecordInfo func(
		sa *T.SAFEARRAY, prinfo *T.IRecordInfo)

	SafeArrayGetRecordInfo func(
		sa *T.SAFEARRAY, prinfo **T.IRecordInfo)

	SafeArraySetIID func(sa *T.SAFEARRAY, guid T.REFGUID)

	SafeArrayGetIID func(sa *T.SAFEARRAY, pguid *T.GUID)

	SafeArrayGetVartype func(sa *T.SAFEARRAY, pvt *T.VARTYPE)

	SafeArrayCreateVector func(vt T.VARTYPE,
		lLbound T.LONG, elements T.ULONG) *T.SAFEARRAY

	SafeArrayCreateVectorEx func(vt T.VARTYPE, lLbound T.LONG,
		elements T.ULONG, pvExtra *T.VOID) *T.SAFEARRAY

	VariantInit func(pvarg *T.VARIANTARG)

	VariantClear func(pvarg *T.VARIANTARG)

	VariantCopy func(dest, src *T.VARIANTARG)

	VariantCopyInd func(pvarDest *T.VARIANT, src *T.VARIANTARG)

	VariantChangeType func(dest, pvarSrc *T.VARIANTARG,
		wFlags T.USHORT, vt T.VARTYPE)

	VariantChangeTypeEx func(dest, pvarSrc *T.VARIANTARG,
		lcid T.LCID, wFlags T.USHORT, vt T.VARTYPE)

	VectorFromBstr func(s T.BSTR, sa **T.SAFEARRAY)

	BstrFromVector func(sa *T.SAFEARRAY, s *T.BSTR)

	VarUI1FromI2 func(sIn T.SHORT, pbOut *T.BYTE)

	VarUI1FromI4 func(lIn T.LONG, pbOut *T.BYTE)

	VarUI1FromI8 func(i64In T.LONG64, pbOut *T.BYTE)

	VarUI1FromR4 func(fltIn T.FLOAT, pbOut *T.BYTE)

	VarUI1FromR8 func(dblIn T.DOUBLE, pbOut *T.BYTE)

	VarUI1FromCy func(in T.CY, pbOut *T.BYTE)

	VarUI1FromDate func(dateIn T.DATE, pbOut *T.BYTE)

	VarUI1FromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, pbOut *T.BYTE)

	VarUI1FromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, pbOut *T.BYTE)

	VarUI1FromBool func(boolIn T.VARIANT_BOOL, pbOut *T.BYTE)

	VarUI1FromI1 func(in T.Char, pbOut *T.BYTE)

	VarUI1FromUI2 func(uiIn T.USHORT, pbOut *T.BYTE)

	VarUI1FromUI4 func(ulIn T.ULONG, pbOut *T.BYTE)

	VarUI1FromUI8 func(ui64In T.ULONG64, pbOut *T.BYTE)

	VarUI1FromDec func(pdecIn *T.DECIMAL, pbOut *T.BYTE)

	VarI2FromUI1 func(bIn T.BYTE, out *T.SHORT)

	VarI2FromI4 func(lIn T.LONG, out *T.SHORT)

	VarI2FromI8 func(i64In T.LONG64, out *T.SHORT)

	VarI2FromR4 func(fltIn T.FLOAT, out *T.SHORT)

	VarI2FromR8 func(dblIn T.DOUBLE, out *T.SHORT)

	VarI2FromCy func(in T.CY, out *T.SHORT)

	VarI2FromDate func(dateIn T.DATE, out *T.SHORT)

	VarI2FromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, out *T.SHORT)

	VarI2FromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, out *T.SHORT)

	VarI2FromBool func(boolIn T.VARIANT_BOOL, out *T.SHORT)

	VarI2FromI1 func(in T.Char, out *T.SHORT)

	VarI2FromUI2 func(uiIn T.USHORT, out *T.SHORT)

	VarI2FromUI4 func(ulIn T.ULONG, out *T.SHORT)

	VarI2FromUI8 func(ui64In T.ULONG64, out *T.SHORT)

	VarI2FromDec func(pdecIn *T.DECIMAL, out *T.SHORT)

	VarI4FromUI1 func(bIn T.BYTE, plOut *T.LONG)

	VarI4FromI2 func(sIn T.SHORT, plOut *T.LONG)

	VarI4FromI8 func(i64In T.LONG64, plOut *T.LONG)

	VarI4FromR4 func(fltIn T.FLOAT, plOut *T.LONG)

	VarI4FromR8 func(dblIn T.DOUBLE, plOut *T.LONG)

	VarI4FromCy func(in T.CY, plOut *T.LONG)

	VarI4FromDate func(dateIn T.DATE, plOut *T.LONG)

	VarI4FromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, plOut *T.LONG)

	VarI4FromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, plOut *T.LONG)

	VarI4FromBool func(boolIn T.VARIANT_BOOL, plOut *T.LONG)

	VarI4FromI1 func(in T.Char, plOut *T.LONG)

	VarI4FromUI2 func(uiIn T.USHORT, plOut *T.LONG)

	VarI4FromUI4 func(ulIn T.ULONG, plOut *T.LONG)

	VarI4FromUI8 func(ui64In T.ULONG64, plOut *T.LONG)

	VarI4FromDec func(pdecIn *T.DECIMAL, plOut *T.LONG)

	VarI4FromInt func(in T.INT, out *T.LONG)

	VarI8FromUI1 func(bIn T.BYTE, pi64Out *T.LONG64)

	VarI8FromI2 func(sIn T.SHORT, pi64Out *T.LONG64)

	VarI8FromI4 func(in T.LONG, out *T.LONG64)

	VarI8FromR4 func(fltIn T.FLOAT, pi64Out *T.LONG64)

	VarI8FromR8 func(dblIn T.DOUBLE, pi64Out *T.LONG64)

	VarI8FromCy func(in T.CY, pi64Out *T.LONG64)

	VarI8FromDate func(dateIn T.DATE, pi64Out *T.LONG64)

	VarI8FromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.Unsigned_long, pi64Out *T.LONG64)

	VarI8FromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, pi64Out *T.LONG64)

	VarI8FromBool func(boolIn T.VARIANT_BOOL, pi64Out *T.LONG64)

	VarI8FromI1 func(in T.Char, pi64Out *T.LONG64)

	VarI8FromUI2 func(uiIn T.USHORT, pi64Out *T.LONG64)

	VarI8FromUI4 func(ulIn T.ULONG, pi64Out *T.LONG64)

	VarI8FromUI8 func(ui64In T.ULONG64, pi64Out *T.LONG64)

	VarI8FromDec func(pdecIn *T.DECIMAL, pi64Out *T.LONG64)

	VarI8FromInt func(intIn T.INT, pi64Out *T.LONG64)

	VarR4FromUI1 func(bIn T.BYTE, pfltOut *T.FLOAT)

	VarR4FromI2 func(sIn T.SHORT, pfltOut *T.FLOAT)

	VarR4FromI4 func(lIn T.LONG, pfltOut *T.FLOAT)

	VarR4FromI8 func(i64In T.LONG64, pfltOut *T.FLOAT)

	VarR4FromR8 func(dblIn T.DOUBLE, pfltOut *T.FLOAT)

	VarR4FromCy func(in T.CY, pfltOut *T.FLOAT)

	VarR4FromDate func(dateIn T.DATE, pfltOut *T.FLOAT)

	VarR4FromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, pfltOut *T.FLOAT)

	VarR4FromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, pfltOut *T.FLOAT)

	VarR4FromBool func(boolIn T.VARIANT_BOOL, pfltOut *T.FLOAT)

	VarR4FromI1 func(in T.Char, pfltOut *T.FLOAT)

	VarR4FromUI2 func(uiIn T.USHORT, pfltOut *T.FLOAT)

	VarR4FromUI4 func(ulIn T.ULONG, pfltOut *T.FLOAT)

	VarR4FromUI8 func(ui64In T.ULONG64, pfltOut *T.FLOAT)

	VarR4FromDec func(pdecIn *T.DECIMAL, pfltOut *T.FLOAT)

	VarR8FromUI1 func(bIn T.BYTE, pdblOut *T.DOUBLE)

	VarR8FromI2 func(sIn T.SHORT, pdblOut *T.DOUBLE)

	VarR8FromI4 func(lIn T.LONG, pdblOut *T.DOUBLE)

	VarR8FromI8 func(i64In T.LONG64, pdblOut *T.DOUBLE)

	VarR8FromR4 func(fltIn T.FLOAT, pdblOut *T.DOUBLE)

	VarR8FromCy func(in T.CY, pdblOut *T.DOUBLE)

	VarR8FromDate func(dateIn T.DATE, pdblOut *T.DOUBLE)

	VarR8FromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, pdblOut *T.DOUBLE)

	VarR8FromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, pdblOut *T.DOUBLE)

	VarR8FromBool func(boolIn T.VARIANT_BOOL, pdblOut *T.DOUBLE)

	VarR8FromI1 func(in T.Char, pdblOut *T.DOUBLE)

	VarR8FromUI2 func(uiIn T.USHORT, pdblOut *T.DOUBLE)

	VarR8FromUI4 func(ulIn T.ULONG, pdblOut *T.DOUBLE)

	VarR8FromUI8 func(ui64In T.ULONG64, pdblOut *T.DOUBLE)

	VarR8FromDec func(pdecIn *T.DECIMAL, pdblOut *T.DOUBLE)

	VarDateFromUI1 func(bIn T.BYTE, pdateOut *T.DATE)

	VarDateFromI2 func(sIn T.SHORT, pdateOut *T.DATE)

	VarDateFromI4 func(lIn T.LONG, pdateOut *T.DATE)

	VarDateFromI8 func(i64In T.LONG64, pdateOut *T.DATE)

	VarDateFromR4 func(fltIn T.FLOAT, pdateOut *T.DATE)

	VarDateFromR8 func(dblIn T.DOUBLE, pdateOut *T.DATE)

	VarDateFromCy func(in T.CY, pdateOut *T.DATE)

	VarDateFromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, pdateOut *T.DATE)

	VarDateFromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, pdateOut *T.DATE)

	VarDateFromBool func(boolIn T.VARIANT_BOOL, pdateOut *T.DATE)

	VarDateFromI1 func(in T.Char, pdateOut *T.DATE)

	VarDateFromUI2 func(uiIn T.USHORT, pdateOut *T.DATE)

	VarDateFromUI4 func(ulIn T.ULONG, pdateOut *T.DATE)

	VarDateFromUI8 func(ui64In T.ULONG64, pdateOut *T.DATE)

	VarDateFromDec func(pdecIn *T.DECIMAL, pdateOut *T.DATE)

	VarCyFromUI1 func(bIn T.BYTE, pcyOut *T.CY)

	VarCyFromI2 func(sIn T.SHORT, pcyOut *T.CY)

	VarCyFromI4 func(lIn T.LONG, pcyOut *T.CY)

	VarCyFromI8 func(i64In T.LONG64, pcyOut *T.CY)

	VarCyFromR4 func(fltIn T.FLOAT, pcyOut *T.CY)

	VarCyFromR8 func(dblIn T.DOUBLE, pcyOut *T.CY)

	VarCyFromDate func(dateIn T.DATE, pcyOut *T.CY)

	VarCyFromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, pcyOut *T.CY)

	VarCyFromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, pcyOut *T.CY)

	VarCyFromBool func(boolIn T.VARIANT_BOOL, pcyOut *T.CY)

	VarCyFromI1 func(in T.Char, pcyOut *T.CY)

	VarCyFromUI2 func(uiIn T.USHORT, pcyOut *T.CY)

	VarCyFromUI4 func(ulIn T.ULONG, pcyOut *T.CY)

	VarCyFromUI8 func(ui64In T.ULONG64, pcyOut *T.CY)

	VarCyFromDec func(pdecIn *T.DECIMAL, pcyOut *T.CY)

	VarBstrFromUI1 func(bVal T.BYTE,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromI2 func(iVal T.SHORT,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromI4 func(lIn T.LONG,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromI8 func(i64In T.LONG64,
		lcid T.LCID, flags T.Unsigned_long, pbstrOut *T.BSTR)

	VarBstrFromR4 func(fltIn T.FLOAT,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromR8 func(dblIn T.DOUBLE,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromCy func(in T.CY,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromDate func(dateIn T.DATE,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromDisp func(pdispIn *T.IDispatch,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromBool func(boolIn T.VARIANT_BOOL,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromI1 func(in T.Char,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromUI2 func(uiIn T.USHORT,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromUI4 func(ulIn T.ULONG,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBstrFromUI8 func(ui64In T.ULONG64,
		lcid T.LCID, flags T.Unsigned_long, pbstrOut *T.BSTR)

	VarBstrFromDec func(pdecIn *T.DECIMAL,
		lcid T.LCID, flags T.ULONG, pbstrOut *T.BSTR)

	VarBoolFromUI1 func(bIn T.BYTE, pboolOut *T.VARIANT_BOOL)

	VarBoolFromI2 func(sIn T.SHORT, pboolOut *T.VARIANT_BOOL)

	VarBoolFromI4 func(lIn T.LONG, pboolOut *T.VARIANT_BOOL)

	VarBoolFromI8 func(i64In T.LONG64, pboolOut *T.VARIANT_BOOL)

	VarBoolFromR4 func(fltIn T.FLOAT, pboolOut *T.VARIANT_BOOL)

	VarBoolFromR8 func(dblIn T.DOUBLE, pboolOut *T.VARIANT_BOOL)

	VarBoolFromDate func(dateIn T.DATE, pboolOut *T.VARIANT_BOOL)

	VarBoolFromCy func(in T.CY, pboolOut *T.VARIANT_BOOL)

	VarBoolFromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, pboolOut *T.VARIANT_BOOL)

	VarBoolFromDisp func(pdispIn *T.IDispatch,
		lcid T.LCID, pboolOut *T.VARIANT_BOOL)

	VarBoolFromI1 func(in T.Char, pboolOut *T.VARIANT_BOOL)

	VarBoolFromUI2 func(uiIn T.USHORT, pboolOut *T.VARIANT_BOOL)

	VarBoolFromUI4 func(ulIn T.ULONG, pboolOut *T.VARIANT_BOOL)

	VarBoolFromUI8 func(i64In T.ULONG64, pboolOut *T.VARIANT_BOOL)

	VarBoolFromDec func(pdecIn *T.DECIMAL, pboolOut *T.VARIANT_BOOL)

	VarI1FromUI1 func(bIn T.BYTE, pcOut *T.Char)

	VarI1FromI2 func(uiIn T.SHORT, pcOut *T.Char)

	VarI1FromI4 func(lIn T.LONG, pcOut *T.Char)

	VarI1FromI8 func(i64In T.LONG64, pcOut *T.Char)

	VarI1FromR4 func(fltIn T.FLOAT, pcOut *T.Char)

	VarI1FromR8 func(dblIn T.DOUBLE, pcOut *T.Char)

	VarI1FromDate func(dateIn T.DATE, pcOut *T.Char)

	VarI1FromCy func(in T.CY, pcOut *T.Char)

	VarI1FromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, pcOut *T.Char)

	VarI1FromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, pcOut *T.Char)

	VarI1FromBool func(boolIn T.VARIANT_BOOL, pcOut *T.Char)

	VarI1FromUI2 func(uiIn T.USHORT, pcOut *T.Char)

	VarI1FromUI4 func(ulIn T.ULONG, pcOut *T.Char)

	VarI1FromUI8 func(i64In T.ULONG64, pcOut *T.Char)

	VarI1FromDec func(pdecIn *T.DECIMAL, pcOut *T.Char)

	VarUI2FromUI1 func(bIn T.BYTE, puiOut *T.USHORT)

	VarUI2FromI2 func(uiIn T.SHORT, puiOut *T.USHORT)

	VarUI2FromI4 func(lIn T.LONG, puiOut *T.USHORT)

	VarUI2FromI8 func(i64In T.LONG64, puiOut *T.USHORT)

	VarUI2FromR4 func(fltIn T.FLOAT, puiOut *T.USHORT)

	VarUI2FromR8 func(dblIn T.DOUBLE, puiOut *T.USHORT)

	VarUI2FromDate func(dateIn T.DATE, puiOut *T.USHORT)

	VarUI2FromCy func(in T.CY, puiOut *T.USHORT)

	VarUI2FromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, puiOut *T.USHORT)

	VarUI2FromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, puiOut *T.USHORT)

	VarUI2FromBool func(boolIn T.VARIANT_BOOL, puiOut *T.USHORT)

	VarUI2FromI1 func(in T.Char, puiOut *T.USHORT)

	VarUI2FromUI4 func(ulIn T.ULONG, puiOut *T.USHORT)

	VarUI2FromUI8 func(i64In T.ULONG64, puiOut *T.USHORT)

	VarUI2FromDec func(pdecIn *T.DECIMAL, puiOut *T.USHORT)

	VarUI4FromUI1 func(bIn T.BYTE, pulOut *T.ULONG)

	VarUI4FromI2 func(uiIn T.SHORT, pulOut *T.ULONG)

	VarUI4FromI4 func(lIn T.LONG, pulOut *T.ULONG)

	VarUI4FromI8 func(i64In T.LONG64, plOut *T.ULONG)

	VarUI4FromR4 func(fltIn T.FLOAT, pulOut *T.ULONG)

	VarUI4FromR8 func(dblIn T.DOUBLE, pulOut *T.ULONG)

	VarUI4FromDate func(dateIn T.DATE, pulOut *T.ULONG)

	VarUI4FromCy func(in T.CY, pulOut *T.ULONG)

	VarUI4FromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, pulOut *T.ULONG)

	VarUI4FromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, pulOut *T.ULONG)

	VarUI4FromBool func(boolIn T.VARIANT_BOOL, pulOut *T.ULONG)

	VarUI4FromI1 func(in T.Char, pulOut *T.ULONG)

	VarUI4FromUI2 func(uiIn T.USHORT, pulOut *T.ULONG)

	VarUI4FromUI8 func(ui64In T.ULONG64, plOut *T.ULONG)

	VarUI4FromDec func(pdecIn *T.DECIMAL, pulOut *T.ULONG)

	VarUI8FromUI1 func(bIn T.BYTE, pi64Out *T.ULONG64)

	VarUI8FromI2 func(sIn T.SHORT, pi64Out *T.ULONG64)

	VarUI8FromI4 func(lIn T.LONG, pi64Out *T.ULONG64)

	VarUI8FromI8 func(ui64In T.LONG64, pi64Out *T.ULONG64)

	VarUI8FromR4 func(fltIn T.FLOAT, pi64Out *T.ULONG64)

	VarUI8FromR8 func(dblIn T.DOUBLE, pi64Out *T.ULONG64)

	VarUI8FromCy func(in T.CY, pi64Out *T.ULONG64)

	VarUI8FromDate func(dateIn T.DATE, pi64Out *T.ULONG64)

	VarUI8FromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.Unsigned_long, pi64Out *T.ULONG64)

	VarUI8FromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, pi64Out *T.ULONG64)

	VarUI8FromBool func(
		boolIn T.VARIANT_BOOL, pi64Out *T.ULONG64)

	VarUI8FromI1 func(in T.Char, pi64Out *T.ULONG64)

	VarUI8FromUI2 func(uiIn T.USHORT, pi64Out *T.ULONG64)

	VarUI8FromUI4 func(ulIn T.ULONG, pi64Out *T.ULONG64)

	VarUI8FromDec func(pdecIn *T.DECIMAL, pi64Out *T.ULONG64)

	VarUI8FromInt func(intIn T.INT, pi64Out *T.ULONG64)

	VarDecFromUI1 func(bIn T.BYTE, pdecOut *T.DECIMAL)

	VarDecFromI2 func(uiIn T.SHORT, pdecOut *T.DECIMAL)

	VarDecFromI4 func(lIn T.LONG, pdecOut *T.DECIMAL)

	VarDecFromI8 func(i64In T.LONG64, pdecOut *T.DECIMAL)

	VarDecFromR4 func(fltIn T.FLOAT, pdecOut *T.DECIMAL)

	VarDecFromR8 func(dblIn T.DOUBLE, pdecOut *T.DECIMAL)

	VarDecFromDate func(dateIn T.DATE, pdecOut *T.DECIMAL)

	VarDecFromCy func(in T.CY, pdecOut *T.DECIMAL)

	VarDecFromStr func(strIn *T.OLECHAR,
		lcid T.LCID, flags T.ULONG, pdecOut *T.DECIMAL)

	VarDecFromDisp func(
		pdispIn *T.IDispatch, lcid T.LCID, pdecOut *T.DECIMAL)

	VarDecFromBool func(boolIn T.VARIANT_BOOL, pdecOut *T.DECIMAL)

	VarDecFromI1 func(in T.Char, pdecOut *T.DECIMAL)

	VarDecFromUI2 func(uiIn T.USHORT, pdecOut *T.DECIMAL)

	VarDecFromUI4 func(ulIn T.ULONG, pdecOut *T.DECIMAL)

	VarDecFromUI8 func(ui64In T.ULONG64, pdecOut *T.DECIMAL)

	VarParseNumFromStr func(strIn *T.OLECHAR, lcid T.LCID,
		flags T.ULONG, pnumprs *T.NUMPARSE, rgbDig *T.BYTE)

	VarNumFromParseNum func(pnumprs *T.NUMPARSE,
		rgbDig *T.BYTE, vtBits T.ULONG, pvar *T.VARIANT)

	VarAdd func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarAnd func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarCat func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarDiv func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarEqv func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarIdiv func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarImp func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarMod func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarMul func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarOr func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarPow func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarSub func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarXor func(pvarLeft, pvarRight, pvarResult *T.VARIANT)

	VarAbs func(pvarIn, pvarResult *T.VARIANT)

	VarFix func(pvarIn, pvarResult *T.VARIANT)

	VarInt func(pvarIn, pvarResult *T.VARIANT)

	VarNeg func(pvarIn, pvarResult *T.VARIANT)

	VarNot func(pvarIn, pvarResult *T.VARIANT)

	VarRound func(
		pvarIn *T.VARIANT, decimals int, pvarResult *T.VARIANT)

	VarCmp func(pvarLeft, pvarRight *T.VARIANT,
		lcid T.LCID, flags T.ULONG)

	VarDecAdd func(pdecLeft, pdecRight, pdecResult *T.DECIMAL)

	VarDecDiv func(pdecLeft, pdecRight, pdecResult *T.DECIMAL)

	VarDecMul func(pdecLeft, pdecRight, pdecResult *T.DECIMAL)

	VarDecSub func(pdecLeft, pdecRight, pdecResult *T.DECIMAL)

	VarDecAbs func(pdecIn, pdecResult *T.DECIMAL)

	VarDecFix func(pdecIn, pdecResult *T.DECIMAL)

	VarDecInt func(pdecIn, pdecResult *T.DECIMAL)

	VarDecNeg func(pdecIn, pdecResult *T.DECIMAL)

	VarDecRound func(
		pdecIn *T.DECIMAL, decimals int, pdecResult *T.DECIMAL)

	VarDecCmp func(pdecLeft, pdecRight *T.DECIMAL)

	VarDecCmpR8 func(pdecLeft *T.DECIMAL, dblRight T.DOUBLE)

	VarCyAdd func(left, right T.CY, pcyResult *T.CY)

	VarCyMul func(left, right T.CY, pcyResult *T.CY)

	VarCyMulI4 func(left T.CY, lRight T.LONG, pcyResult *T.CY)

	VarCyMulI8 func(left T.CY, lRight T.LONG64, pcyResult *T.CY)

	VarCySub func(left, right T.CY, pcyResult *T.CY)

	VarCyAbs func(in T.CY, pcyResult *T.CY)

	VarCyFix func(in T.CY, pcyResult *T.CY)

	VarCyInt func(in T.CY, pcyResult *T.CY)

	VarCyNeg func(in T.CY, pcyResult *T.CY)

	VarCyRound func(in T.CY, decimals int, pcyResult *T.CY)

	VarCyCmp func(left, right T.CY)

	VarCyCmpR8 func(left T.CY, dblRight T.DOUBLE)

	VarBstrCat func(bstrLeft, bstrRight T.BSTR,
		pbstrResult *T.BSTR)

	VarBstrCmp func(
		bstrLeft, bstrRight T.BSTR, lcid T.LCID, flags T.ULONG)

	VarR8Pow func(
		dblLeft, dblRight T.DOUBLE, pdblResult *T.DOUBLE)

	VarR4CmpR8 func(fltLeft T.FLOAT, dblRight T.DOUBLE)

	VarR8Round func(
		dblIn T.DOUBLE, decimals int, pdblResult *T.DOUBLE)

	VarDateFromUdate func(
		pudateIn *T.UDATE, flags T.ULONG, pdateOut *T.DATE)

	VarDateFromUdateEx func(pudateIn *T.UDATE,
		lcid T.LCID, flags T.ULONG, pdateOut *T.DATE)

	VarUdateFromDate func(
		dateIn T.DATE, flags T.ULONG, pudateOut *T.UDATE)

	GetAltMonthNames func(lcid T.LCID, prgp ***T.OLESTR)

	VarFormat func(pvarIn *T.VARIANT,
		pstrFormat *T.OLESTR, iFirstDay, iFirstWeek int,
		flags T.ULONG, pbstrOut *T.BSTR)

	VarFormatDateTime func(pvarIn *T.VARIANT,
		iNamedFormat int, flags T.ULONG, pbstrOut *T.BSTR)

	VarFormatNumber func(pvarIn *T.VARIANT,
		iNumDig, iIncLead, iUseParens, iGroup int,
		flags T.ULONG, pbstrOut *T.BSTR)

	VarFormatPercent func(pvarIn *T.VARIANT,
		iNumDig, iIncLead, iUseParens, iGroup int,
		flags T.ULONG, pbstrOut *T.BSTR)

	VarFormatCurrency func(pvarIn *T.VARIANT,
		iNumDig, iIncLead, iUseParens, iGroup int,
		flags T.ULONG, pbstrOut *T.BSTR)

	VarWeekdayName func(iWeekday, fAbbrev, iFirstDay int,
		flags T.ULONG, pbstrOut *T.BSTR)

	VarMonthName func(
		iMonth, fAbbrev int, flags T.ULONG, pbstrOut *T.BSTR)

	VarFormatFromTokens func(pvarIn *T.VARIANT,
		pstrFormat *T.OLESTR, pbTokCur *T.BYTE,
		flags T.ULONG, pbstrOut *T.BSTR, lcid T.LCID)

	VarTokenizeFormatString func(pstrFormat *T.OLESTR,
		rgbTok *T.BYTE, cbTok, iFirstDay, iFirstWeek int,
		lcid T.LCID, pcbActual *int)

	LHashValOfNameSysA func(
		syskind T.SYSKIND, lcid T.LCID, szName T.AString) T.ULONG

	LHashValOfNameSys func(syskind T.SYSKIND,
		lcid T.LCID, szName *T.OLECHAR) T.ULONG

	LoadTypeLib func(szFile *T.OLECHAR, pptlib **T.ITypeLib)

	LoadTypeLibEx func(
		szFile *T.OLESTR, regkind T.REGKIND, pptlib **T.ITypeLib)

	LoadRegTypeLib func(
		rguid T.REFGUID, verMajor, verMinor T.WORD,
		lcid T.LCID, pptlib **T.ITypeLib)

	QueryPathOfRegTypeLib func(guid T.REFGUID,
		maj, min T.USHORT, lcid T.LCID, pathName *T.BSTR)

	RegisterTypeLib func(
		ptlib *T.ITypeLib, szFullPath, szHelpDir *T.OLECHAR)

	UnRegisterTypeLib func(
		libID T.REFGUID, wVerMajor, wVerMinor T.WORD,
		lcid T.LCID, syskind T.SYSKIND)

	CreateTypeLib func(syskind T.SYSKIND,
		szFile *T.OLECHAR, ppctlib **T.ICreateTypeLib)

	CreateTypeLib2 func(syskind T.SYSKIND,
		szFile *T.OLESTR, ppctlib **T.ICreateTypeLib2)

	DispGetParam func(pdispparams *T.DISPPARAMS,
		position T.UINT, vtTarg T.VARTYPE,
		pvarResult *T.VARIANT, puArgErr *T.UINT)

	DispGetIDsOfNames func(ptinfo *T.ITypeInfo,
		rgszNames **T.OLECHAR, names T.UINT, rgdispid *T.DISPID)

	DispInvoke func(_this *T.VOID, ptinfo *T.ITypeInfo,
		dispidMember T.DISPID, wFlags T.WORD,
		pparams *T.DISPPARAMS, pvarResult *T.VARIANT,
		pexcepinfo *T.EXCEPINFO, puArgErr *T.UINT)

	CreateDispTypeInfo func(pidata *T.INTERFACEDATA,
		lcid T.LCID, pptinfo **T.ITypeInfo)

	CreateStdDispatch func(outer *T.IUnknown, pvThis *T.VOID,
		ptinfo *T.ITypeInfo, ppunkStdDisp **T.IUnknown)

	DispCallFunc func(pvInstance *T.VOID, oVft T.ULONG_PTR,
		cc T.CALLCONV, vtReturn T.VARTYPE, actuals T.UINT,
		prgvt *T.VARTYPE, prgpvarg **T.VARIANTARG,
		result *T.VARIANT)

	RegisterActiveObject func(unk *T.IUnknown,
		rclsid T.REFCLSID, flags T.DWORD, pdwRegister *T.DWORD)

	RevokeActiveObject func(register T.DWORD, pvReserved *T.VOID)

	GetActiveObject func(rclsid T.REFCLSID,
		pvReserved *T.VOID, ppunk **T.IUnknown)

	SetErrorInfo func(
		reserved T.ULONG, perrinfo *T.IErrorInfo)

	GetErrorInfo func(
		reserved T.ULONG, pperrinfo **T.IErrorInfo)

	CreateErrorInfo func(pperrinfo **T.ICreateErrorInfo)

	GetRecordInfoFromTypeInfo func(
		typeInfo *T.ITypeInfo, recInfo **T.IRecordInfo)

	GetRecordInfoFromGuids func(rGuidTypeLib T.REFGUID,
		uVerMajor, uVerMinor T.ULONG, lcid T.LCID,
		rGuidTypeInfo T.REFGUID, recInfo **T.IRecordInfo)

	OaBuildVersion func() T.ULONG

	ClearCustData func(custData *T.CUSTDATA)
)

func VarI4FromI4(in int, out *T.LONG) { *out = T.LONG(in) }

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

   #define T.IsHashValCompatible(lhashval1, lhashval2) \
               ((BOOL) ((0x00ff0000 & (lhashval1)) == (0x00ff0000 & (lhashval2))))
*/

//TODO(t): T.Normal funcs
//{"VarI4FromInt", VarI4FromI4},

//TODO(t): Only in header T.No implementation
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
