package winsock

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/wsock32"
)

type (
	timeval       Fake_type_Fix_me
	servent       Fake_type_Fix_me
	protoent      Fake_type_Fix_me
	u_int         Fake_type_Fix_me
	u_long        Fake_type_Fix_me
	u_short       Fake_type_Fix_me
	hostent       Fake_type_Fix_me
	fd_set        Fake_type_Fix_me
	sockaddr      Fake_type_Fix_me
	long          Fake_type_Fix_me
	in_addr       Fake_type_Fix_me
	unsigned_long Fake_type_Fix_me
	char          byte
)

var (
	WSAFDIsSet func(SOCKET, *fd_set) int

	Accept func(
		s SOCKET,
		addr *sockaddr,
		addrlen *int) SOCKET

	Bind func(
		s SOCKET,
		addr *sockaddr,
		namelen int) int

	Closesocket func(s SOCKET) int

	Connect func(
		s SOCKET,
		name *sockaddr,
		namelen int) int

	Ioctlsocket func(
		s SOCKET,
		cmd long,
		argp *u_long) int

	Getpeername func(
		s SOCKET,
		name *sockaddr,
		namelen *int) int

	Getsockname func(
		s SOCKET,
		name *sockaddr,
		namelen *int) int

	Getsockopt func(
		s SOCKET,
		level int,
		optname int,
		optval *char,
		optlen *int) int

	Htonl func(hostlong u_long) u_long

	Htons func(hostshort u_short) u_short

	Inet_addr func(cp *char) unsigned_long

	Inet_ntoa func(in in_addr) *char

	Listen func(s SOCKET, backlog int) int

	Ntohl func(netlong u_long) u_long

	Ntohs func(netshort u_short) u_short

	Recv func(
		s SOCKET,
		buf *char,
		len int,
		flags int) int

	Recvfrom func(
		s SOCKET,
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
		s SOCKET,
		buf *char,
		len int,
		flags int) int

	Sendto func(
		s SOCKET,
		buf *char,
		len int,
		flags int,
		to *sockaddr,
		tolen int) int

	Setsockopt func(
		s SOCKET,
		level int,
		optname int,
		optval *char,
		optlen int) int

	Shutdown func(
		s SOCKET,
		how int) int

	Socket func(
		af int,
		typ int,
		protocol int) SOCKET

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
		VersionRequired WORD,
		WSAData *WSADATA) int

	WSACleanup func() int

	WSASetLastError func(
		iError int)

	WSAGetLastError func() int

	WSAIsBlocking func() BOOL

	WSAUnhookBlockingHook func() int

	WSASetBlockingHook func(
		BlockFunc FARPROC) FARPROC

	WSACancelBlockingCall func() int

	WSAAsyncGetServByName func(
		Wnd HWND,
		Msg u_int,
		name *char,
		proto *char,
		buf *char,
		buflen int) HANDLE

	WSAAsyncGetServByPort func(
		Wnd HWND,
		Msg u_int,
		port int,
		proto *char,
		buf *char,
		buflen int) HANDLE

	WSAAsyncGetProtoByName func(
		Wnd HWND,
		Msg u_int,
		name *char,
		buf *char,
		buflen int) HANDLE

	WSAAsyncGetProtoByNumber func(
		Wnd HWND,
		Msg u_int,
		number int,
		buf *char,
		buflen int) HANDLE

	WSAAsyncGetHostByName func(
		Wnd HWND,
		Msg u_int,
		name *char,
		buf *char,
		buflen int) HANDLE

	WSAAsyncGetHostByAddr func(
		Wnd HWND,
		Msg u_int,
		addr *char,
		len int,
		typ int,
		buf *char,
		buflen int) HANDLE

	WSACancelAsyncRequest func(hAsyncTaskHandle HANDLE) int

	WSAAsyncSelect func(
		s SOCKET,
		Wnd HWND,
		Msg u_int,
		lEvent long) int

	WSARecvEx func(
		s SOCKET,
		buf *char,
		len int,
		flags *int) int

	TransmitFile func(
		Socket SOCKET,
		File HANDLE,
		NumberOfBytesToWrite DWORD,
		NumberOfBytesPerSend DWORD,
		Overlapped *OVERLAPPED,
		TransmitBuffers *TRANSMIT_FILE_BUFFERS,
		Reserved DWORD) BOOL

	AcceptEx func(
		ListenSocket SOCKET,
		AcceptSocket SOCKET,
		OutputBuffer *VOID,
		ReceiveDataLength DWORD,
		LocalAddressLength DWORD,
		RemoteAddressLength DWORD,
		BytesReceived *DWORD,
		Overlapped *OVERLAPPED) BOOL

	GetAcceptExSockaddrs func(
		OutputBuffer *VOID,
		ReceiveDataLength DWORD,
		LocalAddressLength DWORD,
		RemoteAddressLength DWORD,
		LocalSockaddr **sockaddr,
		LocalSockaddrLength *INT,
		RemoteSockaddr **sockaddr,
		RemoteSockaddrLength *INT)
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
