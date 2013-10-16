// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package mmsystem provides API definitions for accessing
//winmm.dll.
package mmsystem

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside/types"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/winmm"
)

var (
	CloseDriver func(driver T.HDRVR, param1, param2 T.LPARAM) (T.LRESULT, error)

	OpenDriver func(
		driverName, sectionName T.WString, param2 T.LPARAM) (T.HDRVR, error)

	SendDriverMessage func(
		driver T.HDRVR, message T.UINT, param1, param2 T.LPARAM) (T.LRESULT, error)

	DrvGetModuleHandle func(driver T.HDRVR) (T.HMODULE, error)

	GetDriverModuleHandle func(driver T.HDRVR) (T.HMODULE, error)

	DefDriverProc func(
		driverIdentifier T.DWORD_PTR,
		drvr T.HDRVR,
		msg T.UINT,
		param1, param2 T.LPARAM) (T.LRESULT, error)

	SndPlaySound func(sound VString, fSound T.UINT) (T.BOOL, error)

	PlaySound func(sound VString, mod T.HMODULE, fSound T.DWORD) (T.BOOL, error)

	WaveOutGetNumDevs func() (T.UINT, error)

	WaveOutGetDevCapsA func(
		deviceID T.UINT_PTR, woc *T.WAVEOUTCAPSA, size T.UINT) (T.MMRESULT, error)

	WaveOutGetDevCapsW func(
		deviceID T.UINT_PTR, woc *T.WAVEOUTCAPSW, size T.UINT) (T.MMRESULT, error)

	WaveOutGetVolume func(wo T.HWAVEOUT, volume *T.DWORD) (T.MMRESULT, error)

	WaveOutSetVolume func(wo T.HWAVEOUT, volume T.DWORD) (T.MMRESULT, error)

	WaveOutGetErrorText func(
		mmrError T.MMRESULT,
		text OVString,
		sText T.UINT) (T.MMRESULT, error)

	WaveOutOpen func(
		wo *T.HWAVEOUT,
		deviceID T.UINT,
		wfx *T.WAVEFORMATEX,
		callback,
		instance T.DWORD_PTR,
		fOpen T.DWORD) (T.MMRESULT, error)

	WaveOutClose func(wo T.HWAVEOUT) (T.MMRESULT, error)

	WaveOutPrepareHeader func(
		wo T.HWAVEOUT, wh *T.WAVEHDR, cWH T.UINT) (T.MMRESULT, error)

	WaveOutUnprepareHeader func(
		wo T.HWAVEOUT, wh *T.WAVEHDR, cWH T.UINT) (T.MMRESULT, error)

	WaveOutWrite func(
		wo T.HWAVEOUT,
		wh *T.WAVEHDR,
		sWH T.UINT) (T.MMRESULT, error)

	WaveOutPause func(wo T.HWAVEOUT) (T.MMRESULT, error)

	WaveOutRestart func(wo T.HWAVEOUT) (T.MMRESULT, error)

	WaveOutReset func(wo T.HWAVEOUT) (T.MMRESULT, error)

	WaveOutBreakLoop func(wo T.HWAVEOUT) (T.MMRESULT, error)

	WaveOutGetPosition func(
		wo T.HWAVEOUT, mmt *T.MMTIME, cbmmt T.UINT) (T.MMRESULT, error)

	WaveOutGetPitch func(wo T.HWAVEOUT, pitch *T.DWORD) (T.MMRESULT, error)

	WaveOutSetPitch func(wo T.HWAVEOUT, pitch T.DWORD) (T.MMRESULT, error)

	WaveOutGetPlaybackRate func(
		wo T.HWAVEOUT, rate  *T.DWORD) (T.MMRESULT, error)

	WaveOutSetPlaybackRate func(wo T.HWAVEOUT, rate  T.DWORD) (T.MMRESULT, error)

	WaveOutGetID func(wo T.HWAVEOUT, devideID *T.UINT) (T.MMRESULT, error)

	WaveOutMessage func(
		wo T.HWAVEOUT, msg T.UINT, d1, d2 T.DWORD_PTR) (T.MMRESULT, error)

	WaveInGetNumDevs func() (T.UINT, error)

	WaveInGetDevCapsA func(
		deviceID T.UINT_PTR, wic *T.WAVEINCAPSA, cWIC T.UINT) (T.MMRESULT, error)

	WaveInGetDevCapsW func(
		deviceID T.UINT_PTR, wic *T.WAVEINCAPSW, cWIC T.UINT) (T.MMRESULT, error)

	WaveInGetErrorText func(
		mmrError T.MMRESULT, text OVString, cText T.UINT) (T.MMRESULT, error)

	WaveInOpen func(
		wi *T.HWAVEIN,
		deviceID T.UINT,
		wfx *T.WAVEFORMATEX,
		callback, instance T.DWORD_PTR,
		fdwOpen T.DWORD) (T.MMRESULT, error)

	WaveInClose func(wi T.HWAVEIN) (T.MMRESULT, error)

	WaveInPrepareHeader func(
		wi T.HWAVEIN, wh *T.WAVEHDR, cWH T.UINT) (T.MMRESULT, error)

	WaveInUnprepareHeader func(
		wi T.HWAVEIN, wh *T.WAVEHDR, cbwh T.UINT) (T.MMRESULT, error)

	WaveInAddBuffer func(
		wi T.HWAVEIN, wh *T.WAVEHDR, cWH T.UINT) (T.MMRESULT, error)

	WaveInStart func(wi T.HWAVEIN) (T.MMRESULT, error)

	WaveInStop func(wi T.HWAVEIN) (T.MMRESULT, error)

	WaveInReset func(wi T.HWAVEIN) (T.MMRESULT, error)

	WaveInGetPosition func(
		wi T.HWAVEIN, mmt *T.MMTIME, cMMT T.UINT) (T.MMRESULT, error)

	WaveInGetID func(wi T.HWAVEIN, devideID *T.UINT) (T.MMRESULT, error)

	WaveInMessage func(
		wi T.HWAVEIN,  msg T.UINT, d1, d2 T.DWORD_PTR) (T.MMRESULT, error)

	MidiOutGetNumDevs func() (T.UINT, error)

	MidiStreamOpen func(
		ms *T.HMIDISTRM,
		deviceID *T.UINT,
		sMidi T.DWORD,
		callback,
		instance T.DWORD_PTR,
		fOpen T.DWORD) (T.MMRESULT, error)

	MidiStreamClose func(ms T.HMIDISTRM) (T.MMRESULT, error)

	MidiStreamProperty func(
		ms T.HMIDISTRM, ropData *T.BYTE, property T.DWORD) (T.MMRESULT, error)

	MidiStreamPosition func(
		ms T.HMIDISTRM, mmt  *T.MMTIME, cMMT T.UINT) (T.MMRESULT, error)

	MidiStreamOut func(
		ms T.HMIDISTRM, mh *T.MIDIHDR, cMH T.UINT) (T.MMRESULT, error)

	MidiStreamPause func(ms T.HMIDISTRM) (T.MMRESULT, error)

	MidiStreamRestart func(ms T.HMIDISTRM) (T.MMRESULT, error)

	MidiStreamStop func(ms T.HMIDISTRM) (T.MMRESULT, error)

	MidiConnect func(
		mi T.HMIDI, mo T.HMIDIOUT, _ *T.VOID) (T.MMRESULT, error)

	MidiDisconnect func(
		mi T.HMIDI, mo T.HMIDIOUT, _ *T.VOID) (T.MMRESULT, error)

	MidiOutGetDevCapsA func(
		deviceID T.UINT_PTR, moc *T.MIDIOUTCAPSA, cMOC T.UINT) (T.MMRESULT, error)

	MidiOutGetDevCapsW func(
		deviceID T.UINT_PTR, moc *T.MIDIOUTCAPSW, cMOC T.UINT) (T.MMRESULT, error)

	MidiOutGetVolume func(mo T.HMIDIOUT, volume *T.DWORD) (T.MMRESULT, error)

	MidiOutSetVolume func(mo T.HMIDIOUT, volume T.DWORD) (T.MMRESULT, error)

	MidiOutGetErrorText func(
		mmrError T.MMRESULT, text OVString, cText T.UINT) (T.MMRESULT, error)

	MidiOutOpen func(
		mo *T.HMIDIOUT,
		deviceID T.UINT,
		callback ,
		instance T.DWORD_PTR,
		fOpen T.DWORD) (T.MMRESULT, error)

	MidiOutClose func(mo T.HMIDIOUT) (T.MMRESULT, error)

	MidiOutPrepareHeader func(
		mo T.HMIDIOUT, mh *T.MIDIHDR, sMH T.UINT) (T.MMRESULT, error)

	MidiOutUnprepareHeader func(
		mo T.HMIDIOUT, mh *T.MIDIHDR, sMH T.UINT) (T.MMRESULT, error)

	MidiOutShortMsg func(mo T.HMIDIOUT,  msg T.DWORD) (T.MMRESULT, error)

	MidiOutLongMsg func(
		mo T.HMIDIOUT, mh *T.MIDIHDR, sMH T.UINT) (T.MMRESULT, error)

	MidiOutReset func(mo T.HMIDIOUT) (T.MMRESULT, error)

	MidiOutCachePatches func(
		mo T.HMIDIOUT, bank T.UINT, p *T.WORD, fCache T.UINT) (T.MMRESULT, error)

	MidiOutCacheDrumPatches func(
		mo T.HMIDIOUT,
		patch T.UINT,
		sPatch *T.WORD,
		fCache T.UINT) (T.MMRESULT, error)

	MidiOutGetID func(
		mo T.HMIDIOUT,
		deviceID *T.UINT) (T.MMRESULT, error)

	MidiOutMessage func(
		mo T.HMIDIOUT,  msg T.UINT, d1, d2 T.DWORD_PTR) (T.MMRESULT, error)

	MidiInGetNumDevs func() (T.UINT, error)

	MidiInGetDevCapsA func(
		deviceID T.UINT_PTR, mic *T.MIDIINCAPSA, cMIC T.UINT) (T.MMRESULT, error)

	MidiInGetDevCapsW func(
		deviceID T.UINT_PTR, mic *T.MIDIINCAPSW, cMIC T.UINT) (T.MMRESULT, error)

	MidiInGetErrorText func(
		mmrError T.MMRESULT, Text OVString, cText T.UINT) (T.MMRESULT, error)

	MidiInOpen func(
		mi *T.HMIDIIN,
		deviceID T.UINT,
		callback ,
		instance T.DWORD_PTR,
		fOpen T.DWORD) (T.MMRESULT, error)

	MidiInClose func(mi T.HMIDIIN) (T.MMRESULT, error)

	MidiInPrepareHeader func(
		mi T.HMIDIIN, mh *T.MIDIHDR, cMH T.UINT) (T.MMRESULT, error)

	MidiInUnprepareHeader func(
		mi T.HMIDIIN, mh *T.MIDIHDR, cMH T.UINT) (T.MMRESULT, error)

	MidiInAddBuffer func(
		mi T.HMIDIIN, mh *T.MIDIHDR, cMH T.UINT) (T.MMRESULT, error)

	MidiInStart func(mi T.HMIDIIN) (T.MMRESULT, error)

	MidiInStop func(mi T.HMIDIIN) (T.MMRESULT, error)

	MidiInReset func(mi T.HMIDIIN) (T.MMRESULT, error)

	MidiInGetID func(mi T.HMIDIIN, devideID *T.UINT) (T.MMRESULT, error)

	MidiInMessage func(
		mi T.HMIDIIN,  msg T.UINT, d1, d2 T.DWORD_PTR) (T.MMRESULT, error)

	AuxGetNumDevs func() (T.UINT, error)

	AuxGetDevCapsA func(
		deviceID T.UINT_PTR, ac *T.AUXCAPSA, cAC T.UINT) (T.MMRESULT, error)

	AuxGetDevCapsW func(
		deviceID T.UINT_PTR, ac *T.AUXCAPSW, cAC T.UINT) (T.MMRESULT, error)

	AuxSetVolume func(deviceID T.UINT, volume T.DWORD) (T.MMRESULT, error)

	AuxGetVolume func(deviceID T.UINT, volume *T.DWORD) (T.MMRESULT, error)

	AuxOutMessage func(
		deviceID,  msg T.UINT, d1, d2 T.DWORD_PTR) (T.MMRESULT, error)

	MixerGetNumDevs func() (T.UINT, error)

	MixerGetDevCapsA func(
		mxId T.UINT_PTR, mc *T.MIXERCAPSA, sMxCaps T.UINT) (T.MMRESULT, error)

	MixerGetDevCapsW func(
		mxId T.UINT_PTR, mc *T.MIXERCAPSW, sMxCaps T.UINT) (T.MMRESULT, error)

	MixerOpen func(
		mx *T.HMIXER, mxId T.UINT,
		callback , instance T.DWORD_PTR, fOpen T.DWORD) (T.MMRESULT, error)

	MixerClose func(mx T.HMIXER) (T.MMRESULT, error)

	MixerMessage func(
		mx T.HMIXER,  msg T.UINT, param1, param2 T.DWORD_PTR) (T.DWORD, error)

	MixerGetLineInfoA func(
		mxobj T.HMIXEROBJ, ml **T.MIXERLINEA, fInfo T.DWORD) (T.MMRESULT, error)

	MixerGetLineInfoW func(
		mxobj T.HMIXEROBJ, ml **T.MIXERLINEW, fInfo T.DWORD) (T.MMRESULT, error)

	MixerGetID func(
		mxobj T.HMIXEROBJ,
		uMxId *T.UINT,
		fdwId T.DWORD) (T.MMRESULT, error)

	MixerGetControlDetails func(
		mxobj T.HMIXEROBJ,
		mxcd *T.MIXERCONTROLDETAILS,
		fDetails T.DWORD) (T.MMRESULT, error)

	MixerSetControlDetails func(
		mxobj T.HMIXEROBJ,
		mxcd *T.MIXERCONTROLDETAILS,
		fDetails T.DWORD) (T.MMRESULT, error)

	TimeGetSystemTime func(mmt *T.MMTIME, sMMT T.UINT) (T.MMRESULT, error)

	TimeGetTime func() (T.DWORD, error)

	TimeSetEvent func(
		delay T.UINT,
		resolution T.UINT,
		tc *T.TIMECALLBACK,
		user T.DWORD_PTR,
		fEvent T.UINT) (T.MMRESULT, error)

	TimeKillEvent func(
		timerID T.UINT) (T.MMRESULT, error)

	TimeGetDevCaps func(
		tc *T.TIMECAPS,
		sTC T.UINT) (T.MMRESULT, error)

	TimeBeginPeriod func(period T.UINT) (T.MMRESULT, error)

	TimeEndPeriod func(period T.UINT) (T.MMRESULT, error)

	JoyGetNumDevs func() (T.UINT, error)

	JoyGetDevCapsA func(
		joyID T.UINT_PTR, jc *T.JOYCAPSA, sJC T.UINT) (T.MMRESULT, error)

	JoyGetDevCapsW func(
		joyID T.UINT_PTR, jc *T.JOYCAPSW, sJC T.UINT) (T.MMRESULT, error)

	JoyGetPos func(joyID T.UINT, ji *T.JOYINFO) (T.MMRESULT, error)

	JoyGetPosEx func(joyID T.UINT, ji *T.JOYINFOEX) (T.MMRESULT, error)

	JoyGetThreshold func(joyID T.UINT, threshold *T.UINT) (T.MMRESULT, error)

	JoyReleaseCapture func(joyID T.UINT) (T.MMRESULT, error)

	JoySetCapture func(
		wnd T.HWND, T.JoyID, period T.UINT, fChanged T.BOOL) (T.MMRESULT, error)

	JoySetThreshold func(joyID T.UINT, threshold T.UINT) (T.MMRESULT, error)

	MmioStringToFOURCC func(s VString, flags T.UINT) (T.FOURCC, error)

	MmioInstallIOProc func(
		fccIOProc T.FOURCC, ioProc *T.MMIOPROC,
		flags T.DWORD) (*T.MMIOPROC, error)

	MmioOpen func(
		fileName VString,
		mmioinfo *T.MMIOINFO,
		fOpen T.DWORD) (T.HMMIO, error)

	MmioRename func(
		fileName, newFileName VString,
		mmioinfo *T.MMIOINFO,
		fRename T.DWORD) (T.MMRESULT, error)

	MmioClose func(mmio T.HMMIO, fClose T.UINT) (T.MMRESULT, error)

	MmioRead func(mmio T.HMMIO, s *T.Char, sS T.LONG) (T.LONG, error)

	MmioWrite func(mmio T.HMMIO, s *T.Char, sS T.LONG) (T.LONG, error)

	MmioSeek func(mmio T.HMMIO, Offset T.LONG, Origin int) (T.LONG, error)

	MmioGetInfo func(
		mmio T.HMMIO, mmioInfo *T.MMIOINFO, fInfo T.UINT) (T.MMRESULT, error)

	MmioSetInfo func(
		mmio T.HMMIO, mmioinfo *T.MMIOINFO, fInfo T.UINT) (T.MMRESULT, error)

	MmioSetBuffer func(
		mmio T.HMMIO,
		buffer *T.Char,
		sBuffer T.LONG,
		fuBuffer T.UINT) (T.MMRESULT, error)

	MmioFlush func(mmio T.HMMIO, fFlush T.UINT) (T.MMRESULT, error)

	MmioAdvance func(
		mmio T.HMMIO, mmioinfo *T.MMIOINFO, fAdvance T.UINT) (T.MMRESULT, error)

	MmioSendMessage func(
		mmio T.HMMIO,  msg T.UINT, param1, param2 T.LPARAM) (T.LRESULT, error)

	MmioDescend func(
		mmio T.HMMIO,
		mmcki *T.MMCKINFO,
		mmckiParent *T.MMCKINFO,
		fDescend T.UINT) (T.MMRESULT, error)

	MmioAscend func(
		mmio T.HMMIO, mmcki *T.MMCKINFO, fAscend T.UINT) (T.MMRESULT, error)

	MmioCreateChunk func(
		mmio T.HMMIO, mmcki *T.MMCKINFO, fCreate T.UINT) (T.MMRESULT, error)

	MciSendCommand func(
		mciId T.MCIDEVICEID,
		msg T.UINT,
		param1, param2 T.DWORD_PTR) (T.MCIERROR, error)

	MciSendString func(
		command VString,
		returnString OVString,
		returnLength T.UINT,
		wndCallback T.HWND) (T.MCIERROR, error)

	MciGetDeviceID func(device VString) (T.MCIDEVICEID, error)

	MciGetDeviceIDFromElementID func(
		elementID T.DWORD, Type VString) (T.MCIDEVICEID, error)

	MciGetErrorString func(
		err T.MCIERROR, text T.OAString, textSize T.UINT) (T.BOOL, error)

	MciSetYieldProc func(
		mciId T.MCIDEVICEID,
		yieldProc T.YIELDPROC,
		yieldData T.DWORD) (T.BOOL, error)

	MciGetCreatorTask func(mciId T.MCIDEVICEID) (T.HTASK, error)

	MciGetYieldProc func(
		mciId T.MCIDEVICEID, yieldData *T.DWORD) (T.YIELDPROC, error)
)

//func OutputDebugStr(s VString) {  OutputDebugString(s) } // in winbase

var MMSystemANSIApis = outside.Apis{
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

var MMSystemUnicodeApis = outside.Apis{
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

var MMSystemApis = outside.Apis{
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
