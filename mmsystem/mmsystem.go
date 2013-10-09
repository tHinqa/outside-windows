// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package mmsystem provides API definitions for accessing
//winmm.dll.
package mmsystem

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/winmm"
)

var (
	CloseDriver func(Driver HDRVR, Param1, Param2 LPARAM) LRESULT

	OpenDriver func(
		DriverName, SectionName WString, Param2 LPARAM) HDRVR

	SendDriverMessage func(
		Driver HDRVR, Message UINT, Param1, Param2 LPARAM) LRESULT

	DrvGetModuleHandle func(Driver HDRVR) HMODULE

	GetDriverModuleHandle func(Driver HDRVR) HMODULE

	DefDriverProc func(
		DriverIdentifier DWORD_PTR,
		Drvr HDRVR,
		Msg UINT,
		Param1, Param2 LPARAM) LRESULT

	SndPlaySound func(Sound VString, fSound UINT) BOOL

	PlaySound func(Sound VString, mod HMODULE, fSound DWORD) BOOL

	WaveOutGetNumDevs func() UINT

	WaveOutGetDevCapsA func(
		deviceID UINT_PTR, woc *WAVEOUTCAPSA, size UINT) MMRESULT

	WaveOutGetDevCapsW func(
		deviceID UINT_PTR, woc *WAVEOUTCAPSW, size UINT) MMRESULT

	WaveOutGetVolume func(WO HWAVEOUT, Volume *DWORD) MMRESULT

	WaveOutSetVolume func(WO HWAVEOUT, Volume DWORD) MMRESULT

	WaveOutGetErrorText func(
		MmrError MMRESULT,
		Text OVString,
		sText UINT) MMRESULT

	WaveOutOpen func(
		WO *HWAVEOUT,
		DeviceID UINT,
		WFX *WAVEFORMATEX,
		Callback,
		Instance DWORD_PTR,
		fOpen DWORD) MMRESULT

	WaveOutClose func(WO HWAVEOUT) MMRESULT

	WaveOutPrepareHeader func(
		WO HWAVEOUT, WH *WAVEHDR, cWH UINT) MMRESULT

	WaveOutUnprepareHeader func(
		WO HWAVEOUT, WH *WAVEHDR, cWH UINT) MMRESULT

	WaveOutWrite func(
		WO HWAVEOUT,
		WH *WAVEHDR,
		sWH UINT) MMRESULT

	WaveOutPause func(WO HWAVEOUT) MMRESULT

	WaveOutRestart func(WO HWAVEOUT) MMRESULT

	WaveOutReset func(WO HWAVEOUT) MMRESULT

	WaveOutBreakLoop func(WO HWAVEOUT) MMRESULT

	WaveOutGetPosition func(
		WO HWAVEOUT, mmt *MMTIME, cbmmt UINT) MMRESULT

	WaveOutGetPitch func(WO HWAVEOUT, Pitch *DWORD) MMRESULT

	WaveOutSetPitch func(WO HWAVEOUT, Pitch DWORD) MMRESULT

	WaveOutGetPlaybackRate func(
		WO HWAVEOUT, Rate *DWORD) MMRESULT

	WaveOutSetPlaybackRate func(WO HWAVEOUT, Rate DWORD) MMRESULT

	WaveOutGetID func(WO HWAVEOUT, DeviceID *UINT) MMRESULT

	WaveOutMessage func(
		WO HWAVEOUT, Msg UINT, D1, D2 DWORD_PTR) MMRESULT

	WaveInGetNumDevs func() UINT

	WaveInGetDevCapsA func(
		DeviceID UINT_PTR, WIC *WAVEINCAPSA, cWIC UINT) MMRESULT

	WaveInGetDevCapsW func(
		DeviceID UINT_PTR, WIC *WAVEINCAPSW, cWIC UINT) MMRESULT

	WaveInGetErrorText func(
		MmrError MMRESULT, Text OVString, cText UINT) MMRESULT

	WaveInOpen func(
		WI *HWAVEIN,
		DeviceID UINT,
		Wfx *WAVEFORMATEX,
		Callback, Instance DWORD_PTR,
		fdwOpen DWORD) MMRESULT

	WaveInClose func(WI HWAVEIN) MMRESULT

	WaveInPrepareHeader func(
		WI HWAVEIN, Wh *WAVEHDR, cWH UINT) MMRESULT

	WaveInUnprepareHeader func(
		WI HWAVEIN, WH *WAVEHDR, cbwh UINT) MMRESULT

	WaveInAddBuffer func(
		WI HWAVEIN, WH *WAVEHDR, cWH UINT) MMRESULT

	WaveInStart func(WI HWAVEIN) MMRESULT

	WaveInStop func(WI HWAVEIN) MMRESULT

	WaveInReset func(WI HWAVEIN) MMRESULT

	WaveInGetPosition func(
		WI HWAVEIN, MMT *MMTIME, cMMT UINT) MMRESULT

	WaveInGetID func(WI HWAVEIN, DeviceID *UINT) MMRESULT

	WaveInMessage func(
		WI HWAVEIN, Msg UINT, D1, D2 DWORD_PTR) MMRESULT

	MidiOutGetNumDevs func() UINT

	MidiStreamOpen func(
		MS *HMIDISTRM,
		DeviceID *UINT,
		sMidi DWORD,
		Callback,
		Instance DWORD_PTR,
		fOpen DWORD) MMRESULT

	MidiStreamClose func(MS HMIDISTRM) MMRESULT

	MidiStreamProperty func(
		MS HMIDISTRM, RopData *BYTE, Property DWORD) MMRESULT

	MidiStreamPosition func(
		MS HMIDISTRM, MMT *MMTIME, cMMT UINT) MMRESULT

	MidiStreamOut func(
		MS HMIDISTRM, MH *MIDIHDR, cMH UINT) MMRESULT

	MidiStreamPause func(MS HMIDISTRM) MMRESULT

	MidiStreamRestart func(MS HMIDISTRM) MMRESULT

	MidiStreamStop func(MS HMIDISTRM) MMRESULT

	MidiConnect func(
		MI HMIDI, MO HMIDIOUT, Reserved *VOID) MMRESULT

	MidiDisconnect func(
		MI HMIDI, MO HMIDIOUT, Reserved *VOID) MMRESULT

	MidiOutGetDevCapsA func(
		DeviceID UINT_PTR, MOC *MIDIOUTCAPSA, cMOC UINT) MMRESULT

	MidiOutGetDevCapsW func(
		DeviceID UINT_PTR, MOC *MIDIOUTCAPSW, cMOC UINT) MMRESULT

	MidiOutGetVolume func(Mo HMIDIOUT, Volume *DWORD) MMRESULT

	MidiOutSetVolume func(Mo HMIDIOUT, Volume DWORD) MMRESULT

	MidiOutGetErrorText func(
		MmrError MMRESULT, Text OVString, cText UINT) MMRESULT

	MidiOutOpen func(
		MO *HMIDIOUT,
		DeviceID UINT,
		Callback,
		Instance DWORD_PTR,
		fOpen DWORD) MMRESULT

	MidiOutClose func(MO HMIDIOUT) MMRESULT

	MidiOutPrepareHeader func(
		MO HMIDIOUT, MH *MIDIHDR, sMH UINT) MMRESULT

	MidiOutUnprepareHeader func(
		MO HMIDIOUT, MH *MIDIHDR, sMH UINT) MMRESULT

	MidiOutShortMsg func(MO HMIDIOUT, Msg DWORD) MMRESULT

	MidiOutLongMsg func(
		MO HMIDIOUT, MH *MIDIHDR, sMH UINT) MMRESULT

	MidiOutReset func(MO HMIDIOUT) MMRESULT

	MidiOutCachePatches func(
		MO HMIDIOUT, Bank UINT, P *WORD, fCache UINT) MMRESULT

	MidiOutCacheDrumPatches func(
		MO HMIDIOUT,
		Patch UINT,
		sPatch *WORD,
		fCache UINT) MMRESULT

	MidiOutGetID func(
		MO HMIDIOUT,
		DeviceID *UINT) MMRESULT

	MidiOutMessage func(
		MO HMIDIOUT, Msg UINT, D1, D2 DWORD_PTR) MMRESULT

	MidiInGetNumDevs func() UINT

	MidiInGetDevCapsA func(
		DeviceID UINT_PTR, MIC *MIDIINCAPSA, cMIC UINT) MMRESULT

	MidiInGetDevCapsW func(
		DeviceID UINT_PTR, MIC *MIDIINCAPSW, cMIC UINT) MMRESULT

	MidiInGetErrorText func(
		MmrError MMRESULT, Text OVString, cText UINT) MMRESULT

	MidiInOpen func(
		MI *HMIDIIN,
		DeviceID UINT,
		Callback,
		Instance DWORD_PTR,
		fOpen DWORD) MMRESULT

	MidiInClose func(MI HMIDIIN) MMRESULT

	MidiInPrepareHeader func(
		MI HMIDIIN, MH *MIDIHDR, cMH UINT) MMRESULT

	MidiInUnprepareHeader func(
		MI HMIDIIN, MH *MIDIHDR, cMH UINT) MMRESULT

	MidiInAddBuffer func(
		MI HMIDIIN, MH *MIDIHDR, cMH UINT) MMRESULT

	MidiInStart func(MI HMIDIIN) MMRESULT

	MidiInStop func(MI HMIDIIN) MMRESULT

	MidiInReset func(MI HMIDIIN) MMRESULT

	MidiInGetID func(MI HMIDIIN, DeviceID *UINT) MMRESULT

	MidiInMessage func(
		MI HMIDIIN, Msg UINT, D1, D2 DWORD_PTR) MMRESULT

	AuxGetNumDevs func() UINT

	AuxGetDevCapsA func(
		DeviceID UINT_PTR, AC *AUXCAPSA, cAC UINT) MMRESULT

	AuxGetDevCapsW func(
		DeviceID UINT_PTR, AC *AUXCAPSW, cAC UINT) MMRESULT

	AuxSetVolume func(DeviceID UINT, Volume DWORD) MMRESULT

	AuxGetVolume func(DeviceID UINT, Volume *DWORD) MMRESULT

	AuxOutMessage func(
		DeviceID, Msg UINT, D1, D2 DWORD_PTR) MMRESULT

	MixerGetNumDevs func() UINT

	MixerGetDevCapsA func(
		MxId UINT_PTR, mc *MIXERCAPSA, sMxCaps UINT) MMRESULT

	MixerGetDevCapsW func(
		MxId UINT_PTR, mc *MIXERCAPSW, sMxCaps UINT) MMRESULT

	MixerOpen func(
		Mx *HMIXER, MxId UINT,
		Callback, Instance DWORD_PTR, fOpen DWORD) MMRESULT

	MixerClose func(Mx HMIXER) MMRESULT

	MixerMessage func(
		Mx HMIXER, Msg UINT, Param1, Param2 DWORD_PTR) DWORD

	MixerGetLineInfoA func(
		Mxobj HMIXEROBJ, ml **MIXERLINEA, fInfo DWORD) MMRESULT

	MixerGetLineInfoW func(
		Mxobj HMIXEROBJ, ml **MIXERLINEW, fInfo DWORD) MMRESULT

	MixerGetID func(
		Mxobj HMIXEROBJ,
		uMxId *UINT,
		fdwId DWORD) MMRESULT

	MixerGetControlDetails func(
		Mxobj HMIXEROBJ,
		Mxcd *MIXERCONTROLDETAILS,
		fDetails DWORD) MMRESULT

	MixerSetControlDetails func(
		Mxobj HMIXEROBJ,
		Mxcd *MIXERCONTROLDETAILS,
		fDetails DWORD) MMRESULT

	TimeGetSystemTime func(Mmt *MMTIME, sMMT UINT) MMRESULT

	TimeGetTime func() DWORD

	TimeSetEvent func(
		Delay UINT,
		Resolution UINT,
		TC *TIMECALLBACK,
		User DWORD_PTR,
		fEvent UINT) MMRESULT

	TimeKillEvent func(
		TimerID UINT) MMRESULT

	TimeGetDevCaps func(
		TC *TIMECAPS,
		sTC UINT) MMRESULT

	TimeBeginPeriod func(Period UINT) MMRESULT

	TimeEndPeriod func(Period UINT) MMRESULT

	JoyGetNumDevs func() UINT

	JoyGetDevCapsA func(
		JoyID UINT_PTR, JC *JOYCAPSA, sJC UINT) MMRESULT

	JoyGetDevCapsW func(
		JoyID UINT_PTR, JC *JOYCAPSW, sJC UINT) MMRESULT

	JoyGetPos func(JoyID UINT, JI *JOYINFO) MMRESULT

	JoyGetPosEx func(JoyID UINT, JI *JOYINFOEX) MMRESULT

	JoyGetThreshold func(JoyID UINT, Threshold *UINT) MMRESULT

	JoyReleaseCapture func(JoyID UINT) MMRESULT

	JoySetCapture func(
		Wnd HWND, JoyID, Period UINT, fChanged BOOL) MMRESULT

	JoySetThreshold func(JoyID UINT, Threshold UINT) MMRESULT

	MmioStringToFOURCC func(S VString, Flags UINT) FOURCC

	MmioInstallIOProc func(
		FccIOProc FOURCC, IOProc *MMIOPROC,
		Flags DWORD) *MMIOPROC

	MmioOpen func(
		FileName VString,
		Mmioinfo *MMIOINFO,
		fOpen DWORD) HMMIO

	MmioRename func(
		FileName, NewFileName VString,
		Mmioinfo *MMIOINFO,
		fRename DWORD) MMRESULT

	MmioClose func(Mmio HMMIO, fClose UINT) MMRESULT

	MmioRead func(Mmio HMMIO, S *Char, sS LONG) LONG

	MmioWrite func(Mmio HMMIO, S *Char, sS LONG) LONG

	MmioSeek func(Mmio HMMIO, Offset LONG, Origin int) LONG

	MmioGetInfo func(
		Mmio HMMIO, MmioInfo *MMIOINFO, fInfo UINT) MMRESULT

	MmioSetInfo func(
		Mmio HMMIO, Mmioinfo *MMIOINFO, fInfo UINT) MMRESULT

	MmioSetBuffer func(
		Mmio HMMIO,
		Buffer *Char,
		sBuffer LONG,
		fuBuffer UINT) MMRESULT

	MmioFlush func(Mmio HMMIO, fFlush UINT) MMRESULT

	MmioAdvance func(
		Mmio HMMIO, Mmioinfo *MMIOINFO, fAdvance UINT) MMRESULT

	MmioSendMessage func(
		Mmio HMMIO, Msg UINT, Param1, Param2 LPARAM) LRESULT

	MmioDescend func(
		Mmio HMMIO,
		Mmcki *MMCKINFO,
		MmckiParent *MMCKINFO,
		fDescend UINT) MMRESULT

	MmioAscend func(
		Mmio HMMIO, Mmcki *MMCKINFO, fAscend UINT) MMRESULT

	MmioCreateChunk func(
		Mmio HMMIO, Mmcki *MMCKINFO, fCreate UINT) MMRESULT

	MciSendCommand func(
		MciId MCIDEVICEID,
		Msg UINT,
		Param1, Param2 DWORD_PTR) MCIERROR

	MciSendString func(
		Command VString,
		ReturnString OVString,
		ReturnLength UINT,
		WndCallback HWND) MCIERROR

	MciGetDeviceID func(Device VString) MCIDEVICEID

	MciGetDeviceIDFromElementID func(
		ElementID DWORD, Type VString) MCIDEVICEID

	MciGetErrorString func(
		err MCIERROR, text OAString, textSize UINT) BOOL

	MciSetYieldProc func(
		MciId MCIDEVICEID,
		YieldProc YIELDPROC,
		YieldData DWORD) BOOL

	MciGetCreatorTask func(MciId MCIDEVICEID) HTASK

	MciGetYieldProc func(
		MciId MCIDEVICEID, YieldData *DWORD) YIELDPROC
)

//func OutputDebugStr(s VString) {  OutputDebugString(s) } // in winbase

var MMSystemANSIApis = Apis{
	{"auxGetDevCapsW", &AuxGetDevCapsW},
	{"joyGetDevCapsW", &JoyGetDevCapsW},
	{"mciGetDeviceIDW", &MciGetDeviceID},
	{"mciGetDeviceIDFromElementIDW", &MciGetDeviceIDFromElementID},
	{"mciGetErrorStringW", &MciGetErrorString},
	{"mciSendCommandW", &MciSendCommand},
	{"mciSendStringW", &MciSendString},
	{"midiInGetDevCapsW", &MidiInGetDevCapsW},
	{"midiInGetErrorTextW", &MidiInGetErrorText},
	{"midiOutGetDevCapsW", &MidiOutGetDevCapsW},
	{"midiOutGetErrorTextW", &MidiOutGetErrorText},
	{"mixerGetControlDetailsW", &MixerGetControlDetails},
	{"mixerGetDevCapsW", &MixerGetDevCapsW},
	{"mixerGetLineInfoW", &MixerGetLineInfoW},
	{"mmioInstallIOProcW", &MmioInstallIOProc},
	{"mmioOpenW", &MmioOpen},
	{"mmioRenameW", &MmioRename},
	{"PlaySoundW", &PlaySound},
	{"sndPlaySoundW", &SndPlaySound},
	{"waveInGetDevCapsW", &WaveInGetDevCapsW},
	{"waveInGetErrorTextW", &WaveInGetErrorText},
	{"waveOutGetDevCapsW", &WaveOutGetDevCapsW},
	{"waveOutGetErrorTextW", &WaveOutGetErrorText},
}

var MMSystemUnicodeApis = Apis{
	{"auxGetDevCapsA", &AuxGetDevCapsA},
	{"joyGetDevCapsA", &JoyGetDevCapsA},
	{"mciGetDeviceIDA", &MciGetDeviceID},
	{"mciGetDeviceIDFromElementIDA", &MciGetDeviceIDFromElementID},
	{"mciGetErrorStringA", &MciGetErrorString},
	{"mciSendCommandA", &MciSendCommand},
	{"mciSendStringA", &MciSendString},
	{"midiInGetDevCapsA", &MidiInGetDevCapsA},
	{"midiInGetErrorTextA", &MidiInGetErrorText},
	{"midiOutGetDevCapsA", &MidiOutGetDevCapsA},
	{"midiOutGetErrorTextA", &MidiOutGetErrorText},
	{"mixerGetControlDetailsA", &MixerGetControlDetails},
	{"mixerGetDevCapsA", &MixerGetDevCapsA},
	{"mixerGetLineInfoA", &MixerGetLineInfoA},
	{"mmioInstallIOProcA", &MmioInstallIOProc},
	{"mmioOpenA", &MmioOpen},
	{"mmioRenameA", &MmioRename},
	{"PlaySound", &PlaySound}, // == PlaySoundA?
	{"PlaySoundA", &PlaySound},
	{"sndPlaySoundA", &SndPlaySound},
	{"waveInGetDevCapsA", &WaveInGetDevCapsA},
	{"waveInGetErrorTextA", &WaveInGetErrorText},
	{"waveOutGetDevCapsA", &WaveOutGetDevCapsA},
	{"waveOutGetErrorTextA", &WaveOutGetErrorText},
}

var MMSystemApis = Apis{
	{"auxGetNumDevs", &AuxGetNumDevs},
	{"auxGetVolume", &AuxGetVolume},
	{"auxOutMessage", &AuxOutMessage},
	{"auxSetVolume", &AuxSetVolume},
	{"CloseDriver", &CloseDriver},
	{"DefDriverProc", &DefDriverProc},
	{"DrvGetModuleHandle", &DrvGetModuleHandle},
	{"GetDriverModuleHandle", &GetDriverModuleHandle},
	{"joyGetNumDevs", &JoyGetNumDevs},
	{"joyGetPos", &JoyGetPos},
	{"joyGetPosEx", &JoyGetPosEx},
	{"joyGetThreshold", &JoyGetThreshold},
	{"joyReleaseCapture", &JoyReleaseCapture},
	{"joySetCapture", &JoySetCapture},
	{"joySetThreshold", &JoySetThreshold},
	{"mciGetCreatorTask", &MciGetCreatorTask},
	{"mciGetYieldProc", &MciGetYieldProc},
	{"mciSetYieldProc", &MciSetYieldProc},
	{"midiConnect", &MidiConnect},
	{"midiDisconnect", &MidiDisconnect},
	{"midiInAddBuffer", &MidiInAddBuffer},
	{"midiInClose", &MidiInClose},
	{"midiInGetID", &MidiInGetID},
	{"midiInGetNumDevs", &MidiInGetNumDevs},
	{"midiInMessage", &MidiInMessage},
	{"midiInOpen", &MidiInOpen},
	{"midiInPrepareHeader", &MidiInPrepareHeader},
	{"midiInReset", &MidiInReset},
	{"midiInStart", &MidiInStart},
	{"midiInStop", &MidiInStop},
	{"midiInUnprepareHeader", &MidiInUnprepareHeader},
	{"midiOutCacheDrumPatches", &MidiOutCacheDrumPatches},
	{"midiOutCachePatches", &MidiOutCachePatches},
	{"midiOutClose", &MidiOutClose},
	{"midiOutGetID", &MidiOutGetID},
	{"midiOutGetNumDevs", &MidiOutGetNumDevs},
	{"midiOutGetVolume", &MidiOutGetVolume},
	{"midiOutLongMsg", &MidiOutLongMsg},
	{"midiOutMessage", &MidiOutMessage},
	{"midiOutOpen", &MidiOutOpen},
	{"midiOutPrepareHeader", &MidiOutPrepareHeader},
	{"midiOutReset", &MidiOutReset},
	{"midiOutSetVolume", &MidiOutSetVolume},
	{"midiOutShortMsg", &MidiOutShortMsg},
	{"midiOutUnprepareHeader", &MidiOutUnprepareHeader},
	{"midiStreamClose", &MidiStreamClose},
	{"midiStreamOpen", &MidiStreamOpen},
	{"midiStreamOut", &MidiStreamOut},
	{"midiStreamPause", &MidiStreamPause},
	{"midiStreamPosition", &MidiStreamPosition},
	{"midiStreamProperty", &MidiStreamProperty},
	{"midiStreamRestart", &MidiStreamRestart},
	{"midiStreamStop", &MidiStreamStop},
	{"mixerClose", &MixerClose},
	{"mixerGetID", &MixerGetID},
	{"mixerGetNumDevs", &MixerGetNumDevs},
	{"mixerMessage", &MixerMessage},
	{"mixerOpen", &MixerOpen},
	{"mixerSetControlDetails", &MixerSetControlDetails},
	{"mmioAdvance", &MmioAdvance},
	{"mmioAscend", &MmioAscend},
	{"mmioClose", &MmioClose},
	{"mmioCreateChunk", &MmioCreateChunk},
	{"mmioDescend", &MmioDescend},
	{"mmioFlush", &MmioFlush},
	{"mmioGetInfo", &MmioGetInfo},
	{"mmioRead", &MmioRead},
	{"mmioSeek", &MmioSeek},
	{"mmioSendMessage", &MmioSendMessage},
	{"mmioSetBuffer", &MmioSetBuffer},
	{"mmioSetInfo", &MmioSetInfo},
	{"mmioStringToFOURCCA", &MmioStringToFOURCC},
	{"mmioWrite", &MmioWrite},
	{"OpenDriver", &OpenDriver},
	{"SendDriverMessage", &SendDriverMessage},
	{"timeBeginPeriod", &TimeBeginPeriod},
	{"timeEndPeriod", &TimeEndPeriod},
	{"timeGetDevCaps", &TimeGetDevCaps},
	{"timeGetSystemTime", &TimeGetSystemTime},
	{"timeGetTime", &TimeGetTime},
	{"timeKillEvent", &TimeKillEvent},
	{"timeSetEvent", &TimeSetEvent},
	{"waveInAddBuffer", &WaveInAddBuffer},
	{"waveInClose", &WaveInClose},
	{"waveInGetID", &WaveInGetID},
	{"waveInGetNumDevs", &WaveInGetNumDevs},
	{"waveInGetPosition", &WaveInGetPosition},
	{"waveInMessage", &WaveInMessage},
	{"waveInOpen", &WaveInOpen},
	{"waveInPrepareHeader", &WaveInPrepareHeader},
	{"waveInReset", &WaveInReset},
	{"waveInStart", &WaveInStart},
	{"waveInStop", &WaveInStop},
	{"waveInUnprepareHeader", &WaveInUnprepareHeader},
	{"waveOutBreakLoop", &WaveOutBreakLoop},
	{"waveOutClose", &WaveOutClose},
	{"waveOutGetID", &WaveOutGetID},
	{"waveOutGetNumDevs", &WaveOutGetNumDevs},
	{"waveOutGetPitch", &WaveOutGetPitch},
	{"waveOutGetPlaybackRate", &WaveOutGetPlaybackRate},
	{"waveOutGetPosition", &WaveOutGetPosition},
	{"waveOutGetVolume", &WaveOutGetVolume},
	{"waveOutMessage", &WaveOutMessage},
	{"waveOutOpen", &WaveOutOpen},
	{"waveOutPause", &WaveOutPause},
	{"waveOutPrepareHeader", &WaveOutPrepareHeader},
	{"waveOutReset", &WaveOutReset},
	{"waveOutRestart", &WaveOutRestart},
	{"waveOutSetPitch", &WaveOutSetPitch},
	{"waveOutSetPlaybackRate", &WaveOutSetPlaybackRate},
	{"waveOutSetVolume", &WaveOutSetVolume},
	{"waveOutUnprepareHeader", &WaveOutUnprepareHeader},
	{"waveOutWrite", &WaveOutWrite},
}
