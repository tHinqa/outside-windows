// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package winsock2 provides API definitions for accessing
//ws2_32.dll.
package winsock2

import (
	"github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	. "github.com/tHinqa/outside/types"
	_ "github.com/tHinqa/outside/win32/ws2_32"
)

var (
	Accept func(
		s T.SOCKET, addr *T.SOCKADDR, addrlen *int) T.SOCKET

	Bind func(s T.SOCKET, name *T.SOCKADDR, namelen int) int

	Closesocket func(s T.SOCKET) int

	Connect func(s T.SOCKET, name *T.SOCKADDR, namelen int) int

	Ioctlsocket func(s T.SOCKET, cmd T.LONG, argp *T.U_long) int

	Getpeername func(
		s T.SOCKET, name *T.SOCKADDR, namelen *int) int

	Getsockname func(
		s T.SOCKET, name *T.SOCKADDR, namelen *int) int

	Getsockopt func(s T.SOCKET,
		level, optname int, optval *T.Char, optlen *int) int

	Htonl func(hostlong T.U_long) T.U_long

	Htons func(hostshort T.U_short) T.U_short

	Inet_addr func(cp *T.Char) T.Unsigned_long

	Inet_ntoa func(in T.IN_ADDR) *T.Char

	Listen func(s T.SOCKET, backlog int) int

	Ntohl func(netlong T.U_long) T.U_long

	Ntohs func(netshort T.U_short) T.U_short

	Recv func(s T.SOCKET, buf *T.Char, leng, flags int) int

	Recvfrom func(s T.SOCKET, buf *T.Char,
		leng, flags int, from *T.SOCKADDR, fromlen *int) int

	Select func(nfds int, readfds,
		writefds, exceptfds *T.FD_SET, timeout *T.TIMEVAL) int

	Send func(s T.SOCKET, buf *T.Char, len int, flags int) int

	Sendto func(s T.SOCKET, buf *T.Char,
		leng, flags int, to *T.SOCKADDR, tolen int) int

	Setsockopt func(s T.SOCKET,
		level, optname int, optval *T.Char, optlen int) int

	Shutdown func(s T.SOCKET, how int) int

	Socket func(af, typ, protocol int) T.SOCKET

	Gethostbyaddr func(addr *T.Char, leng, typ int) *T.HOSTENT

	Gethostbyname func(name *T.Char) *T.HOSTENT

	Gethostname func(name *T.Char, namelen int) int

	Getservbyport func(port int, proto *T.Char) *T.SERVENT

	Getservbyname func(name, proto *T.Char) *T.SERVENT

	Getprotobynumber func(number int) *T.PROTOENT

	Getprotobyname func(name *T.Char) *T.PROTOENT

	WSAStartup func(
		VersionRequested T.WORD, WSAData *T.WSADATA) int

	WSACleanup func() int

	WSASetLastError func(Error int)

	WSAGetLastError func() int

	WSAIsBlocking func() T.BOOL

	WSAUnhookBlockingHook func() int

	WSASetBlockingHook func(BlockFunc T.FARPROC) T.FARPROC

	WSACancelBlockingCall func() int

	WSAAsyncGetServByName func(Wnd T.HWND, Msg T.U_int,
		name, proto, buf *T.Char, buflen int) T.HANDLE

	WSAAsyncGetServByPort func(Wnd T.HWND, Msg T.U_int,
		port int, proto, buf *T.Char, buflen int) T.HANDLE

	WSAAsyncGetProtoByName func(Wnd T.HWND,
		Msg T.U_int, name, buf *T.Char, buflen int) T.HANDLE

	WSAAsyncGetProtoByNumber func(Wnd T.HWND, Msg T.U_int,
		number int, buf *T.Char, buflen int) T.HANDLE

	WSAAsyncGetHostByName func(Wnd T.HWND, Msg T.U_int,
		name *T.Char, buf *T.Char, buflen int) T.HANDLE

	WSAAsyncGetHostByAddr func(Wnd T.HWND, Msg T.U_int, addr *T.Char,
		leng, typ int, buf *T.Char, buflen int) T.HANDLE

	WSACancelAsyncRequest func(AsyncTaskHandle T.HANDLE) int

	WSAAsyncSelect func(
		s T.SOCKET, Wnd T.HWND, Msg T.U_int, lEvent T.LONG) int

	WSAAccept func(
		s T.SOCKET,
		addr *T.SOCKADDR,
		addrlen *T.INT,
		fnCondition *T.CONDITIONPROC,
		CallbackData T.DWORD_PTR) T.SOCKET

	WSACloseEvent func(Event T.WSAEVENT) T.BOOL

	WSAConnect func(s T.SOCKET, name *T.SOCKADDR,
		namelen int,
		CallerData, CalleeData *T.WSABUF,
		SQOS, GQOS *T.QOS) int

	WSACreateEvent func() T.WSAEVENT

	WSADuplicateSocketA func(
		s T.SOCKET,
		ProcessId T.DWORD,
		ProtocolInfo *T.WSAPROTOCOL_INFOA) int

	WSADuplicateSocketW func(
		s T.SOCKET,
		ProcessId T.DWORD,
		ProtocolInfo *T.WSAPROTOCOL_INFOW) int

	WSAEnumNetworkEvents func(
		s T.SOCKET,
		EventObject T.WSAEVENT,
		NetworkEvents *T.WSANETWORKEVENTS) int

	WSAEnumProtocolsA func(
		Protocols *T.INT,
		ProtocolBuffer *T.WSAPROTOCOL_INFOA,
		BufferLength *T.DWORD) int

	WSAEnumProtocolsW func(
		Protocols *T.INT,
		ProtocolBuffer *T.WSAPROTOCOL_INFOW,
		BufferLength *T.DWORD) int

	WSAEventSelect func(
		s T.SOCKET,
		EventObject T.WSAEVENT,
		NetworkEvents T.LONG) int

	WSAGetOverlappedResult func(
		s T.SOCKET,
		Overlapped *T.WSAOVERLAPPED,
		Transfer *T.DWORD,
		fWait T.BOOL,
		Flags *T.DWORD) T.BOOL

	WSAGetQOSByName func(
		s T.SOCKET, QOSName *T.WSABUF, QOS *T.QOS) T.BOOL

	WSAHtonl func(
		s T.SOCKET, hostlong T.U_long, netlong *T.U_long) int

	WSAHtons func(
		s T.SOCKET, hostshort T.U_short, netshort *T.U_short) int

	//TODO(t):Callback indirection

	WSAIoctl func(
		s T.SOCKET,
		IoControlCode T.DWORD,
		InBuffer *T.VOID,
		sInBuffer T.DWORD,
		OutBuffer *T.VOID,
		sOutBuffer T.DWORD,
		BytesReturned *T.DWORD,
		Overlapped *T.WSAOVERLAPPED,
		CompletionRoutine *T.WSAOVERLAPPED_COMPLETION_ROUTINE) int

	WSAJoinLeaf func(s T.SOCKET, name *T.SOCKADDR, namelen int,
		CallerData, CalleeData *T.WSABUF,
		SQOS, GQOS *T.QOS, Flags T.DWORD) T.SOCKET

	WSANtohl func(
		s T.SOCKET, netlong T.U_long, hostlong *T.U_long) int

	WSANtohs func(
		s T.SOCKET, netshort T.U_short, hostshort *T.U_short) int

	//TODO(t):Callback indirection

	WSARecv func(s T.SOCKET, Buffers *T.WSABUF,
		BufferCount T.DWORD, NumberOfBytesRecvd, Flags *T.DWORD,
		Overlapped *T.WSAOVERLAPPED,
		CompletionRoutine *T.WSAOVERLAPPED_COMPLETION_ROUTINE) int

	WSARecvDisconnect func(s T.SOCKET,
		InboundDisconnectData *T.WSABUF) int

	//TODO(t):Callback indirection

	WSARecvFrom func(
		s T.SOCKET,
		Buffers *T.WSABUF,
		BufferCount T.DWORD,
		NumberOfBytesRecvd *T.DWORD,
		Flags *T.DWORD,
		From *T.SOCKADDR,
		Fromlen *T.INT,
		Overlapped *T.WSAOVERLAPPED,
		CompletionRoutine *T.WSAOVERLAPPED_COMPLETION_ROUTINE) int

	WSAResetEvent func(Event T.WSAEVENT) T.BOOL

	//TODO(t):Callback indirection

	WSASend func(
		s T.SOCKET,
		Buffers *T.WSABUF,
		BufferCount T.DWORD,
		NumberOfBytesSent *T.DWORD,
		Flags T.DWORD,
		Overlapped *T.WSAOVERLAPPED,
		CompletionRoutine *T.WSAOVERLAPPED_COMPLETION_ROUTINE) int

	WSASendDisconnect func(
		s T.SOCKET, OutboundDisconnectData *T.WSABUF) int

	//TODO(t):Callback indirection

	WSASendTo func(
		s T.SOCKET,
		Buffers *T.WSABUF,
		BufferCount T.DWORD,
		NumberOfBytesSent *T.DWORD,
		Flags T.DWORD,
		To *T.SOCKADDR,
		iTolen int,
		Overlapped *T.WSAOVERLAPPED,
		CompletionRoutine *T.WSAOVERLAPPED_COMPLETION_ROUTINE) int

	WSASetEvent func(Event T.WSAEVENT) T.BOOL

	WSASocketA func(
		af, typ, protocol int,
		ProtocolInfo *T.WSAPROTOCOL_INFOA,
		g T.GROUP,
		Flags T.DWORD) T.SOCKET
	WSASocketW func(
		af, typ, protocol int,
		ProtocolInfo *T.WSAPROTOCOL_INFOW,
		g T.GROUP,
		Flags T.DWORD) T.SOCKET

	WSAWaitForMultipleEvents func(
		nEvents T.DWORD,
		Events *T.WSAEVENT,
		WaitAll T.BOOL,
		Timeout T.DWORD,
		Alertable T.BOOL) T.DWORD

	WSAAddressToStringA func(
		Address *T.SOCKADDR,
		AddressLength T.DWORD,
		ProtocolInfo *T.WSAPROTOCOL_INFOA,
		AddressString OVString,
		AddressStringLength *T.DWORD) T.INT

	WSAAddressToStringW func(
		Address *T.SOCKADDR,
		AddressLength T.DWORD,
		ProtocolInfo *T.WSAPROTOCOL_INFOW,
		AddressString OVString,
		AddressStringLength *T.DWORD) T.INT

	WSAStringToAddressA func(
		AddressString VString,
		AddressFamily T.INT,
		ProtocolInfo *T.WSAPROTOCOL_INFOA,
		Address *T.SOCKADDR,
		AddressLength *T.INT) T.INT

	WSAStringToAddressW func(
		AddressString VString,
		AddressFamily T.INT,
		ProtocolInfo *T.WSAPROTOCOL_INFOW,
		Address *T.SOCKADDR,
		AddressLength *T.INT) T.INT

	WSALookupServiceBegin func(
		Restrictions *T.WSAQUERYSET,
		ControlFlags T.DWORD,
		Lookup *T.HANDLE) T.INT

	WSALookupServiceNext func(
		Lookup T.HANDLE,
		ControlFlags T.DWORD,
		BufferLength *T.DWORD,
		Results *T.WSAQUERYSET) T.INT

	WSANSPIoctl func(
		Lookup T.HANDLE,
		ControlCode T.DWORD,
		InBuffer *T.VOID,
		sInBuffer T.DWORD,
		OutBuffer *T.VOID,
		sOutBuffer T.DWORD,
		BytesReturned *T.DWORD,
		Completion *T.WSACOMPLETION) T.INT

	WSALookupServiceEnd func(Lookup T.HANDLE) T.INT

	WSAInstallServiceClass func(
		ServiceClassInfo *T.WSASERVICECLASSINFO) T.INT

	WSARemoveServiceClass func(ServiceClassId *T.GUID) T.INT

	WSAGetServiceClassInfo func(
		ProviderId *T.GUID,
		ServiceClassId *T.GUID,
		BufSize *T.DWORD,
		ServiceClassInfo *T.WSASERVICECLASSINFO) T.INT

	WSAEnumNameSpaceProviders func(
		BufferLength *T.DWORD,
		Buffer *T.WSANAMESPACE_INFO) T.INT

	WSAGetServiceClassNameByClassId func(
		ServiceClassId *T.GUID,
		ServiceClassName OVString,
		BufferLength *T.DWORD) T.INT

	WSASetService func(
		RegInfo *T.WSAQUERYSET,
		essoperation T.WSAESETSERVICEOP,
		ControlFlags T.DWORD) T.INT

	//TODO(t):Callback indirection

	WSAProviderConfigChange func(
		NotificationHandle *T.HANDLE,
		Overlapped *T.WSAOVERLAPPED,
		CompletionRoutine *T.WSAOVERLAPPED_COMPLETION_ROUTINE) T.INT
)

var WinSock2ANSIApis = outside.Apis{
	{"WSAAddressToStringA", &WSAAddressToStringA},
	{"WSADuplicateSocketA", &WSADuplicateSocketA},
	{"WSAEnumNameSpaceProvidersA", &WSAEnumNameSpaceProviders},
	{"WSAEnumProtocolsA", &WSAEnumProtocolsA},
	{"WSAGetServiceClassInfoA", &WSAGetServiceClassInfo},
	{"WSAGetServiceClassNameByClassIdA", &WSAGetServiceClassNameByClassId},
	{"WSAInstallServiceClassA", &WSAInstallServiceClass},
	{"WSALookupServiceBeginA", &WSALookupServiceBegin},
	{"WSALookupServiceNextA", &WSALookupServiceNext},
	{"WSASetServiceA", &WSASetService},
	{"WSASocketA", &WSASocketA},
	{"WSAStringToAddressA", &WSAStringToAddressA},
}

var WinSock2UnicodeApis = outside.Apis{
	{"WSAAddressToStringW", &WSAAddressToStringW},
	{"WSADuplicateSocketW", &WSADuplicateSocketW},
	{"WSAEnumNameSpaceProvidersW", &WSAEnumNameSpaceProviders},
	{"WSAEnumProtocolsW", &WSAEnumProtocolsW},
	{"WSAGetServiceClassInfoW", &WSAGetServiceClassInfo},
	{"WSAGetServiceClassNameByClassIdW", &WSAGetServiceClassNameByClassId},
	{"WSAInstallServiceClassW", &WSAInstallServiceClass},
	{"WSALookupServiceBeginW", &WSALookupServiceBegin},
	{"WSALookupServiceNextW", &WSALookupServiceNext},
	{"WSASetServiceW", &WSASetService},
	{"WSASocketW", &WSASocketW},
	{"WSAStringToAddressW", &WSAStringToAddressW},
}

var WinSock2Apis = outside.Apis{
	{"accept", &Accept},
	{"bind", &Bind},
	{"closesocket", &Closesocket},
	{"connect", &Connect},
	{"ioctlsocket", &Ioctlsocket},
	{"gethostbyaddr", &Gethostbyaddr},
	{"gethostbyname", &Gethostbyname},
	{"gethostname", &Gethostname},
	{"getpeername", &Getpeername},
	{"getservbyname", &Getservbyname},
	{"getservbyport", &Getservbyport},
	{"getprotobyname", &Getprotobyname},
	{"getprotobynumber", &Getprotobynumber},
	{"getsockname", &Getsockname},
	{"getsockopt", &Getsockopt},
	{"htonl", &Htonl},
	{"htons", &Htons},
	{"inet_addr", &Inet_addr},
	{"inet_ntoa", &Inet_ntoa},
	{"listen", &Listen},
	{"ntohl", &Ntohl},
	{"ntohs", &Ntohs},
	{"recv", &Recv},
	{"recvfrom", &Recvfrom},
	{"select", &Select},
	{"send", &Send},
	{"sendto", &Sendto},
	{"setsockopt", &Setsockopt},
	{"shutdown", &Shutdown},
	{"socket", &Socket},
	{"WSAAccept", &WSAAccept},
	{"WSAAsyncGetHostByAddr", &WSAAsyncGetHostByAddr},
	{"WSAAsyncGetHostByName", &WSAAsyncGetHostByName},
	{"WSAAsyncGetProtoByName", &WSAAsyncGetProtoByName},
	{"WSAAsyncGetProtoByNumber", &WSAAsyncGetProtoByNumber},
	{"WSAAsyncGetServByName", &WSAAsyncGetServByName},
	{"WSAAsyncGetServByPort", &WSAAsyncGetServByPort},
	{"WSAAsyncSelect", &WSAAsyncSelect},
	{"WSACancelAsyncRequest", &WSACancelAsyncRequest},
	{"WSACancelBlockingCall", &WSACancelBlockingCall},
	{"WSACleanup", &WSACleanup},
	{"WSACloseEvent", &WSACloseEvent},
	{"WSAConnect", &WSAConnect},
	{"WSACreateEvent", &WSACreateEvent},
	{"WSAEnumNetworkEvents", &WSAEnumNetworkEvents},
	{"WSAEventSelect", &WSAEventSelect},
	{"WSAGetLastError", &WSAGetLastError},
	{"WSAGetOverlappedResult", &WSAGetOverlappedResult},
	{"WSAGetQOSByName", &WSAGetQOSByName},
	{"WSAHtonl", &WSAHtonl},
	{"WSAHtons", &WSAHtons},
	{"WSAIoctl", &WSAIoctl},
	{"WSAIsBlocking", &WSAIsBlocking},
	{"WSAJoinLeaf", &WSAJoinLeaf},
	{"WSALookupServiceEnd", &WSALookupServiceEnd},
	{"WSANSPIoctl", &WSANSPIoctl},
	{"WSANtohl", &WSANtohl},
	{"WSANtohs", &WSANtohs},
	{"WSAProviderConfigChange", &WSAProviderConfigChange},
	{"WSARecv", &WSARecv},
	{"WSARecvDisconnect", &WSARecvDisconnect},
	{"WSARecvFrom", &WSARecvFrom},
	{"WSARemoveServiceClass", &WSARemoveServiceClass},
	{"WSAResetEvent", &WSAResetEvent},
	{"WSASend", &WSASend},
	{"WSASendDisconnect", &WSASendDisconnect},
	{"WSASendTo", &WSASendTo},
	{"WSASetBlockingHook", &WSASetBlockingHook},
	{"WSASetEvent", &WSASetEvent},
	{"WSASetLastError", &WSASetLastError},
	{"WSAStartup", &WSAStartup},
	{"WSAUnhookBlockingHook", &WSAUnhookBlockingHook},
	{"WSAWaitForMultipleEvents", &WSAWaitForMultipleEvents},
}
