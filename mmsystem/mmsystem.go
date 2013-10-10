// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package mmsystem provides API definitions for accessing
//winmm.dll.
package mmsystem

import (
	. "github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/winmm"
)

var (
	CloseDriver func(driver T.HDRVR, param1, param2 T.LPARAM) T.LRESULT

	OpenDriver func(
		driverName, sectionName T.WString, param2 T.LPARAM) T.HDRVR

	SendDriverMessage func(
		driver T.HDRVR, message T.UINT, param1, param2 T.LPARAM) T.LRESULT

	DrvGetModuleHandle func(driver T.HDRVR) T.HMODULE

	GetDriverModuleHandle func(driver T.HDRVR) T.HMODULE

	DefDriverProc func(
		driverIdentifier T.DWORD_PTR,
		drvr T.HDRVR,
		msg T.UINT,
		param1, param2 T.LPARAM) T.LRESULT

	SndPlaySound func(sound VString, fSound T.UINT) T.BOOL

	PlaySound func(sound VString, mod T.HMODULE, fSound T.DWORD) T.BOOL

	WaveOutGetNumDevs func() T.UINT

	WaveOutGetDevCapsA func(
		deviceID T.UINT_PTR, woc *T.WAVEOUTCAPSA, size T.UINT) T.MMRESULT

	WaveOutGetDevCapsW func(
		deviceID T.UINT_PTR, woc *T.WAVEOUTCAPSW, size T.UINT) T.MMRESULT

	WaveOutGetVolume func(wo T.HWAVEOUT, volume *T.DWORD) T.MMRESULT

	WaveOutSetVolume func(wo T.HWAVEOUT, volume T.DWORD) T.MMRESULT

	WaveOutGetErrorText func(
		mmrError T.MMRESULT,
		text OVString,
		sText T.UINT) T.MMRESULT

	WaveOutOpen func(
		wo *T.HWAVEOUT,
		deviceID T.UINT,
		wfx *T.WAVEFORMATEX,
		callback,
		instance T.DWORD_PTR,
		fOpen T.DWORD) T.MMRESULT

	WaveOutClose func(wo T.HWAVEOUT) T.MMRESULT

	WaveOutPrepareHeader func(
		wo T.HWAVEOUT, wh *T.WAVEHDR, cWH T.UINT) T.MMRESULT

	WaveOutUnprepareHeader func(
		wo T.HWAVEOUT, wh *T.WAVEHDR, cWH T.UINT) T.MMRESULT

	WaveOutWrite func(
		wo T.HWAVEOUT,
		wh *T.WAVEHDR,
		sWH T.UINT) T.MMRESULT

	WaveOutPause func(wo T.HWAVEOUT) T.MMRESULT

	WaveOutRestart func(wo T.HWAVEOUT) T.MMRESULT

	WaveOutReset func(wo T.HWAVEOUT) T.MMRESULT

	WaveOutBreakLoop func(wo T.HWAVEOUT) T.MMRESULT

	WaveOutGetPosition func(
		wo T.HWAVEOUT, mmt *T.MMTIME, cbmmt T.UINT) T.MMRESULT

	WaveOutGetPitch func(wo T.HWAVEOUT, pitch *T.DWORD) T.MMRESULT

	WaveOutSetPitch func(wo T.HWAVEOUT, pitch T.DWORD) T.MMRESULT

	WaveOutGetPlaybackRate func(
		wo T.HWAVEOUT, rate  *T.DWORD) T.MMRESULT

	WaveOutSetPlaybackRate func(wo T.HWAVEOUT, rate  T.DWORD) T.MMRESULT

	WaveOutGetID func(wo T.HWAVEOUT, devideID *T.UINT) T.MMRESULT

	WaveOutMessage func(
		wo T.HWAVEOUT, msg T.UINT, d1, d2 T.DWORD_PTR) T.MMRESULT

	WaveInGetNumDevs func() T.UINT

	WaveInGetDevCapsA func(
		deviceID T.UINT_PTR, wic *T.WAVEINCAPSA, cWIC T.UINT) T.MMRESULT

	WaveInGetDevCapsW func(
		deviceID T.UINT_PTR, wic *T.WAVEINCAPSW, cWIC T.UINT) T.MMRESULT

	WaveInGetErrorText func(
		mmrError T.MMRESULT, text OVString, cText T.UINT) T.MMRESULT

	WaveInOpen func(
		wi *T.HWAVEIN,
		deviceID T.UINT,
		wfx *T.WAVEFORMATEX,
		callback, instance T.DWORD_PTR,
		fdwOpen T.DWORD) T.MMRESULT

	WaveInClose func(wi T.HWAVEIN) T.MMRESULT

	WaveInPrepareHeader func(
		wi T.HWAVEIN, wh *T.WAVEHDR, cWH T.UINT) T.MMRESULT

	WaveInUnprepareHeader func(
		wi T.HWAVEIN, wh *T.WAVEHDR, cbwh T.UINT) T.MMRESULT

	WaveInAddBuffer func(
		wi T.HWAVEIN, wh *T.WAVEHDR, cWH T.UINT) T.MMRESULT

	WaveInStart func(wi T.HWAVEIN) T.MMRESULT

	WaveInStop func(wi T.HWAVEIN) T.MMRESULT

	WaveInReset func(wi T.HWAVEIN) T.MMRESULT

	WaveInGetPosition func(
		wi T.HWAVEIN, mmt *T.MMTIME, cMMT T.UINT) T.MMRESULT

	WaveInGetID func(wi T.HWAVEIN, devideID *T.UINT) T.MMRESULT

	WaveInMessage func(
		wi T.HWAVEIN,  msg T.UINT, d1, d2 T.DWORD_PTR) T.MMRESULT

	MidiOutGetNumDevs func() T.UINT

	MidiStreamOpen func(
		ms *T.HMIDISTRM,
		deviceID *T.UINT,
		sMidi T.DWORD,
		callback,
		instance T.DWORD_PTR,
		fOpen T.DWORD) T.MMRESULT

	MidiStreamClose func(ms T.HMIDISTRM) T.MMRESULT

	MidiStreamProperty func(
		ms T.HMIDISTRM, ropData *T.BYTE, property T.DWORD) T.MMRESULT

	MidiStreamPosition func(
		ms T.HMIDISTRM, mmt  *T.MMTIME, cMMT T.UINT) T.MMRESULT

	MidiStreamOut func(
		ms T.HMIDISTRM, mh *T.MIDIHDR, cMH T.UINT) T.MMRESULT

	MidiStreamPause func(ms T.HMIDISTRM) T.MMRESULT

	MidiStreamRestart func(ms T.HMIDISTRM) T.MMRESULT

	MidiStreamStop func(ms T.HMIDISTRM) T.MMRESULT

	MidiConnect func(
		mi T.HMIDI, mo T.HMIDIOUT, _ *T.VOID) T.MMRESULT

	MidiDisconnect func(
		mi T.HMIDI, mo T.HMIDIOUT, _ *T.VOID) T.MMRESULT

	MidiOutGetDevCapsA func(
		deviceID T.UINT_PTR, moc *T.MIDIOUTCAPSA, cMOC T.UINT) T.MMRESULT

	MidiOutGetDevCapsW func(
		deviceID T.UINT_PTR, moc *T.MIDIOUTCAPSW, cMOC T.UINT) T.MMRESULT

	MidiOutGetVolume func(mo T.HMIDIOUT, volume *T.DWORD) T.MMRESULT

	MidiOutSetVolume func(mo T.HMIDIOUT, volume T.DWORD) T.MMRESULT

	MidiOutGetErrorText func(
		mmrError T.MMRESULT, text OVString, cText T.UINT) T.MMRESULT

	MidiOutOpen func(
		mo *T.HMIDIOUT,
		deviceID T.UINT,
		callback ,
		instance T.DWORD_PTR,
		fOpen T.DWORD) T.MMRESULT

	MidiOutClose func(mo T.HMIDIOUT) T.MMRESULT

	MidiOutPrepareHeader func(
		mo T.HMIDIOUT, mh *T.MIDIHDR, sMH T.UINT) T.MMRESULT

	MidiOutUnprepareHeader func(
		mo T.HMIDIOUT, mh *T.MIDIHDR, sMH T.UINT) T.MMRESULT

	MidiOutShortMsg func(mo T.HMIDIOUT,  msg T.DWORD) T.MMRESULT

	MidiOutLongMsg func(
		mo T.HMIDIOUT, mh *T.MIDIHDR, sMH T.UINT) T.MMRESULT

	MidiOutReset func(mo T.HMIDIOUT) T.MMRESULT

	MidiOutCachePatches func(
		mo T.HMIDIOUT, bank T.UINT, p *T.WORD, fCache T.UINT) T.MMRESULT

	MidiOutCacheDrumPatches func(
		mo T.HMIDIOUT,
		patch T.UINT,
		sPatch *T.WORD,
		fCache T.UINT) T.MMRESULT

	MidiOutGetID func(
		mo T.HMIDIOUT,
		deviceID *T.UINT) T.MMRESULT

	MidiOutMessage func(
		mo T.HMIDIOUT,  msg T.UINT, d1, d2 T.DWORD_PTR) T.MMRESULT

	MidiInGetNumDevs func() T.UINT

	MidiInGetDevCapsA func(
		deviceID T.UINT_PTR, mic *T.MIDIINCAPSA, cMIC T.UINT) T.MMRESULT

	MidiInGetDevCapsW func(
		deviceID T.UINT_PTR, mic *T.MIDIINCAPSW, cMIC T.UINT) T.MMRESULT

	MidiInGetErrorText func(
		mmrError T.MMRESULT, Text OVString, cText T.UINT) T.MMRESULT

	MidiInOpen func(
		mi *T.HMIDIIN,
		deviceID T.UINT,
		callback ,
		instance T.DWORD_PTR,
		fOpen T.DWORD) T.MMRESULT

	MidiInClose func(mi T.HMIDIIN) T.MMRESULT

	MidiInPrepareHeader func(
		mi T.HMIDIIN, mh *T.MIDIHDR, cMH T.UINT) T.MMRESULT

	MidiInUnprepareHeader func(
		mi T.HMIDIIN, mh *T.MIDIHDR, cMH T.UINT) T.MMRESULT

	MidiInAddBuffer func(
		mi T.HMIDIIN, mh *T.MIDIHDR, cMH T.UINT) T.MMRESULT

	MidiInStart func(mi T.HMIDIIN) T.MMRESULT

	MidiInStop func(mi T.HMIDIIN) T.MMRESULT

	MidiInReset func(mi T.HMIDIIN) T.MMRESULT

	MidiInGetID func(mi T.HMIDIIN, devideID *T.UINT) T.MMRESULT

	MidiInMessage func(
		mi T.HMIDIIN,  msg T.UINT, d1, d2 T.DWORD_PTR) T.MMRESULT

	AuxGetNumDevs func() T.UINT

	AuxGetDevCapsA func(
		deviceID T.UINT_PTR, ac *T.AUXCAPSA, cAC T.UINT) T.MMRESULT

	AuxGetDevCapsW func(
		deviceID T.UINT_PTR, ac *T.AUXCAPSW, cAC T.UINT) T.MMRESULT

	AuxSetVolume func(deviceID T.UINT, volume T.DWORD) T.MMRESULT

	AuxGetVolume func(deviceID T.UINT, volume *T.DWORD) T.MMRESULT

	AuxOutMessage func(
		deviceID,  msg T.UINT, d1, d2 T.DWORD_PTR) T.MMRESULT

	MixerGetNumDevs func() T.UINT

	MixerGetDevCapsA func(
		mxId T.UINT_PTR, mc *T.MIXERCAPSA, sMxCaps T.UINT) T.MMRESULT

	MixerGetDevCapsW func(
		mxId T.UINT_PTR, mc *T.MIXERCAPSW, sMxCaps T.UINT) T.MMRESULT

	MixerOpen func(
		mx *T.HMIXER, mxId T.UINT,
		callback , instance T.DWORD_PTR, fOpen T.DWORD) T.MMRESULT

	MixerClose func(mx T.HMIXER) T.MMRESULT

	MixerMessage func(
		mx T.HMIXER,  msg T.UINT, param1, param2 T.DWORD_PTR) T.DWORD

	MixerGetLineInfoA func(
		mxobj T.HMIXEROBJ, ml **T.MIXERLINEA, fInfo T.DWORD) T.MMRESULT

	MixerGetLineInfoW func(
		mxobj T.HMIXEROBJ, ml **T.MIXERLINEW, fInfo T.DWORD) T.MMRESULT

	MixerGetID func(
		mxobj T.HMIXEROBJ,
		uMxId *T.UINT,
		fdwId T.DWORD) T.MMRESULT

	MixerGetControlDetails func(
		mxobj T.HMIXEROBJ,
		mxcd *T.MIXERCONTROLDETAILS,
		fDetails T.DWORD) T.MMRESULT

	MixerSetControlDetails func(
		mxobj T.HMIXEROBJ,
		mxcd *T.MIXERCONTROLDETAILS,
		fDetails T.DWORD) T.MMRESULT

	TimeGetSystemTime func(mmt *T.MMTIME, sMMT T.UINT) T.MMRESULT

	TimeGetTime func() T.DWORD

	TimeSetEvent func(
		delay T.UINT,
		resolution T.UINT,
		tc *T.TIMECALLBACK,
		user T.DWORD_PTR,
		fEvent T.UINT) T.MMRESULT

	TimeKillEvent func(
		timerID T.UINT) T.MMRESULT

	TimeGetDevCaps func(
		tc *T.TIMECAPS,
		sTC T.UINT) T.MMRESULT

	TimeBeginPeriod func(period T.UINT) T.MMRESULT

	TimeEndPeriod func(period T.UINT) T.MMRESULT

	JoyGetNumDevs func() T.UINT

	JoyGetDevCapsA func(
		joyID T.UINT_PTR, jc *T.JOYCAPSA, sJC T.UINT) T.MMRESULT

	JoyGetDevCapsW func(
		joyID T.UINT_PTR, jc *T.JOYCAPSW, sJC T.UINT) T.MMRESULT

	JoyGetPos func(joyID T.UINT, ji *T.JOYINFO) T.MMRESULT

	JoyGetPosEx func(joyID T.UINT, ji *T.JOYINFOEX) T.MMRESULT

	JoyGetThreshold func(joyID T.UINT, threshold *T.UINT) T.MMRESULT

	JoyReleaseCapture func(joyID T.UINT) T.MMRESULT

	JoySetCapture func(
		wnd T.HWND, T.JoyID, period T.UINT, fChanged T.BOOL) T.MMRESULT

	JoySetThreshold func(joyID T.UINT, threshold T.UINT) T.MMRESULT

	MmioStringToFOURCC func(s VString, flags T.UINT) T.FOURCC

	MmioInstallIOProc func(
		fccIOProc T.FOURCC, ioProc *T.MMIOPROC,
		flags T.DWORD) *T.MMIOPROC

	MmioOpen func(
		fileName VString,
		mmioinfo *T.MMIOINFO,
		fOpen T.DWORD) T.HMMIO

	MmioRename func(
		fileName, newFileName VString,
		mmioinfo *T.MMIOINFO,
		fRename T.DWORD) T.MMRESULT

	MmioClose func(mmio T.HMMIO, fClose T.UINT) T.MMRESULT

	MmioRead func(mmio T.HMMIO, s *T.Char, sS T.LONG) T.LONG

	MmioWrite func(mmio T.HMMIO, s *T.Char, sS T.LONG) T.LONG

	MmioSeek func(mmio T.HMMIO, Offset T.LONG, Origin int) T.LONG

	MmioGetInfo func(
		mmio T.HMMIO, mmioInfo *T.MMIOINFO, fInfo T.UINT) T.MMRESULT

	MmioSetInfo func(
		mmio T.HMMIO, mmioinfo *T.MMIOINFO, fInfo T.UINT) T.MMRESULT

	MmioSetBuffer func(
		mmio T.HMMIO,
		buffer *T.Char,
		sBuffer T.LONG,
		fuBuffer T.UINT) T.MMRESULT

	MmioFlush func(mmio T.HMMIO, fFlush T.UINT) T.MMRESULT

	MmioAdvance func(
		mmio T.HMMIO, mmioinfo *T.MMIOINFO, fAdvance T.UINT) T.MMRESULT

	MmioSendMessage func(
		mmio T.HMMIO,  msg T.UINT, param1, param2 T.LPARAM) T.LRESULT

	MmioDescend func(
		mmio T.HMMIO,
		mmcki *T.MMCKINFO,
		mmckiParent *T.MMCKINFO,
		fDescend T.UINT) T.MMRESULT

	MmioAscend func(
		mmio T.HMMIO, mmcki *T.MMCKINFO, fAscend T.UINT) T.MMRESULT

	MmioCreateChunk func(
		mmio T.HMMIO, mmcki *T.MMCKINFO, fCreate T.UINT) T.MMRESULT

	MciSendCommand func(
		mciId T.MCIDEVICEID,
		msg T.UINT,
		param1, param2 T.DWORD_PTR) T.MCIERROR

	MciSendString func(
		command VString,
		returnString OVString,
		returnLength T.UINT,
		wndCallback T.HWND) T.MCIERROR

	MciGetDeviceID func(device VString) T.MCIDEVICEID

	MciGetDeviceIDFromElementID func(
		elementID T.DWORD, Type VString) T.MCIDEVICEID

	MciGetErrorString func(
		err T.MCIERROR, text T.OAString, textSize T.UINT) T.BOOL

	MciSetYieldProc func(
		mciId T.MCIDEVICEID,
		yieldProc T.YIELDPROC,
		yieldData T.DWORD) T.BOOL

	MciGetCreatorTask func(mciId T.MCIDEVICEID) T.HTASK

	MciGetYieldProc func(
		mciId T.MCIDEVICEID, yieldData *T.DWORD) T.YIELDPROC
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
