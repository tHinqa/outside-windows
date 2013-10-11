// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package winsock provides API definitions for accessing
//wsock32.dll.
package winsock

import (
	. "github.com/tHinqa/outside"
	T "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/wsock32"
)

type (
	timeval       T.Fake_type_Fix_me
	servent       T.Fake_type_Fix_me
	protoent      T.Fake_type_Fix_me
	u_int         T.Fake_type_Fix_me
	u_long        T.Fake_type_Fix_me
	u_short       T.Fake_type_Fix_me
	hostent       T.Fake_type_Fix_me
	fd_set        T.Fake_type_Fix_me
	sockaddr      T.Fake_type_Fix_me
	long          T.Fake_type_Fix_me
	in_addr       T.Fake_type_Fix_me
	unsigned_long T.Fake_type_Fix_me
	char          byte
)

var (
	WSAFDIsSet func(T.SOCKET, *fd_set) int

	Accept func(
		s T.SOCKET,
		addr *sockaddr,
		addrlen *int) T.SOCKET

	Bind func(
		s T.SOCKET,
		addr *sockaddr,
		namelen int) int

	Closesocket func(s T.SOCKET) int

	Connect func(
		s T.SOCKET,
		name *sockaddr,
		namelen int) int

	Ioctlsocket func(
		s T.SOCKET,
		cmd long,
		argp *u_long) int

	Getpeername func(
		s T.SOCKET,
		name *sockaddr,
		namelen *int) int

	Getsockname func(
		s T.SOCKET,
		name *sockaddr,
		namelen *int) int

	Getsockopt func(
		s T.SOCKET,
		level int,
		optname int,
		optval *char,
		optlen *int) int

	Htonl func(hostlong u_long) u_long

	Htons func(hostshort u_short) u_short

	Inet_addr func(cp *char) unsigned_long

	Inet_ntoa func(in in_addr) *char

	Listen func(s T.SOCKET, backlog int) int

	Ntohl func(netlong u_long) u_long

	Ntohs func(netshort u_short) u_short

	Recv func(
		s T.SOCKET,
		buf *char,
		len int,
		flags int) int

	Recvfrom func(
		s T.SOCKET,
		buf *char,
		len int,
		flags int,
		from *sockaddr,
		fromlen *int) int

	Select func(
		fds int,
		readfds *fd_set,
		writefds *fd_set,
		exceptfds *fd_set,
		timeout *timeval) int

	Send func(
		s T.SOCKET,
		buf *char,
		len int,
		flags int) int

	Sendto func(
		s T.SOCKET,
		buf *char,
		len int,
		flags int,
		to *sockaddr,
		tolen int) int

	Setsockopt func(
		s T.SOCKET,
		level int,
		optname int,
		optval *char,
		optlen int) int

	Shutdown func(
		s T.SOCKET,
		how int) int

	Socket func(
		af int,
		typ int,
		protocol int) T.SOCKET

	Gethostbyaddr func(
		addr *char,
		len int,
		typ int) *hostent

	Gethostbyname func(
		name *char) *hostent

	Gethostname func(
		name *char,
		namelen int) int

	Getservbyport func(
		port int,
		proto *char) *servent

	Getservbyname func(
		name *char,
		proto *char) *servent

	Getprotobynumber func(
		proto int) *protoent

	Getprotobyname func(
		name *char) *protoent

	WSAStartup func(
		VersionRequired T.WORD,
		WSAData *T.WSADATA) int

	WSACleanup func() int

	WSASetLastError func(
		iError int)

	WSAGetLastError func() int

	WSAIsBlocking func() T.BOOL

	WSAUnhookBlockingHook func() int

	WSASetBlockingHook func(
		BlockFunc T.FARPROC) T.FARPROC

	WSACancelBlockingCall func() int

	WSAAsyncGetServByName func(
		Wnd T.HWND,
		Msg u_int,
		name *char,
		proto *char,
		buf *char,
		buflen int) T.HANDLE

	WSAAsyncGetServByPort func(
		Wnd T.HWND,
		Msg u_int,
		port int,
		proto *char,
		buf *char,
		buflen int) T.HANDLE

	WSAAsyncGetProtoByName func(
		Wnd T.HWND,
		Msg u_int,
		name *char,
		buf *char,
		buflen int) T.HANDLE

	WSAAsyncGetProtoByNumber func(
		Wnd T.HWND,
		Msg u_int,
		number int,
		buf *char,
		buflen int) T.HANDLE

	WSAAsyncGetHostByName func(
		Wnd T.HWND,
		Msg u_int,
		name *char,
		buf *char,
		buflen int) T.HANDLE

	WSAAsyncGetHostByAddr func(
		Wnd T.HWND,
		Msg u_int,
		addr *char,
		len int,
		typ int,
		buf *char,
		buflen int) T.HANDLE

	WSACancelAsyncRequest func(hAsyncTaskHandle T.HANDLE) int

	WSAAsyncSelect func(
		s T.SOCKET,
		Wnd T.HWND,
		Msg u_int,
		lEvent long) int

	WSARecvEx func(
		s T.SOCKET,
		buf *char,
		len int,
		flags *int) int

	TransmitFile func(
		Socket T.SOCKET,
		File T.HANDLE,
		NumberOfBytesToWrite T.DWORD,
		NumberOfBytesPerSend T.DWORD,
		Overlapped *T.OVERLAPPED,
		TransmitBuffers *T.TRANSMIT_FILE_BUFFERS,
		Reserved T.DWORD) T.BOOL

	AcceptEx func(
		ListenSocket T.SOCKET,
		AcceptSocket T.SOCKET,
		OutputBuffer *T.VOID,
		ReceiveDataLength T.DWORD,
		LocalAddressLength T.DWORD,
		RemoteAddressLength T.DWORD,
		BytesReceived *T.DWORD,
		Overlapped *T.OVERLAPPED) T.BOOL

	GetAcceptExSockaddrs func(
		OutputBuffer *T.VOID,
		ReceiveDataLength T.DWORD,
		LocalAddressLength T.DWORD,
		RemoteAddressLength T.DWORD,
		LocalSockaddr **sockaddr,
		LocalSockaddrLength *T.INT,
		RemoteSockaddr **sockaddr,
		RemoteSockaddrLength *T.INT)
)

var WinSockApis = Apis{
	{"accept", &Accept},
	{"AcceptEx", &AcceptEx},
	{"bind", &Bind},
	{"closesocket", &Closesocket},
	{"connect", &Connect},
	{"GetAcceptExSockaddrs", &GetAcceptExSockaddrs},
	{"gethostbyaddr", &Gethostbyaddr},
	{"gethostbyname", &Gethostbyname},
	{"gethostname", &Gethostname},
	{"getpeername", &Getpeername},
	{"getprotobyname", &Getprotobyname},
	{"getprotobynumber", &Getprotobynumber},
	{"getservbyname", &Getservbyname},
	{"getservbyport", &Getservbyport},
	{"getsockname", &Getsockname},
	{"getsockopt", &Getsockopt},
	{"htonl", &Htonl},
	{"htons", &Htons},
	{"inet_addr", &Inet_addr},
	{"inet_ntoa", &Inet_ntoa},
	{"ioctlsocket", &Ioctlsocket},
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
	{"TransmitFile", &TransmitFile},
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
	{"__WSAFDIsSet", &WSAFDIsSet},
	{"WSAGetLastError", &WSAGetLastError},
	{"WSAIsBlocking", &WSAIsBlocking},
	{"WSARecvEx", &WSARecvEx},
	{"WSASetBlockingHook", &WSASetBlockingHook},
	{"WSASetLastError", &WSASetLastError},
	{"WSAStartup", &WSAStartup},
	{"WSAUnhookBlockingHook", &WSAUnhookBlockingHook},
}
