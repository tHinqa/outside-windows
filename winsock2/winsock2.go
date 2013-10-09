package winsock2

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/ws2_32"
)

var (
	Accept func(
		s SOCKET,
		addr *SOCKADDR,
		addrlen *int) SOCKET
	Bind func(
		s SOCKET,
		name *SOCKADDR,
		namelen int) int
	Closesocket func(
		s SOCKET) int
	Connect func(
		s SOCKET,
		name *SOCKADDR,
		namelen int) int
	Ioctlsocket func(
		s SOCKET,
		cmd LONG,
		argp *U_long) int
	Getpeername func(
		s SOCKET,
		name *SOCKADDR,
		namelen *int) int
	Getsockname func(
		s SOCKET,
		name *SOCKADDR,
		namelen *int) int
	Getsockopt func(
		s SOCKET,
		level int,
		optname int,
		optval *Char,
		optlen *int) int
	Htonl func(
		hostlong U_long) U_long
	Htons func(
		hostshort U_short) U_short
	Inet_addr func(
		cp *Char) Unsigned_long
	Inet_ntoa func(
		in IN_ADDR) *Char
	Listen func(
		s SOCKET,
		backlog int) int
	Ntohl func(
		netlong U_long) U_long
	Ntohs func(
		netshort U_short) U_short
	Recv func(
		s SOCKET,
		buf *Char,
		len int,
		flags int) int
	Recvfrom func(
		s SOCKET,
		buf *Char,
		len int,
		flags int,
		from *SOCKADDR,
		fromlen *int) int
	Select func(
		nfds int,
		readfds *FD_SET,
		writefds *FD_SET,
		exceptfds *FD_SET,
		timeout *TIMEVAL) int
	Send func(
		s SOCKET,
		buf *Char,
		len int,
		flags int) int
	Sendto func(
		s SOCKET,
		buf *Char,
		len int,
		flags int,
		to *SOCKADDR,
		tolen int) int
	Setsockopt func(
		s SOCKET,
		level int,
		optname int,
		optval *Char,
		optlen int) int
	Shutdown func(
		s SOCKET,
		how int) int
	Socket func(
		af int,
		typ int,
		protocol int) SOCKET
	Gethostbyaddr func(
		addr *Char,
		leng int,
		typ int) *HOSTENT
	Gethostbyname func(
		name *Char) *HOSTENT
	Gethostname func(
		name *Char,
		namelen int) int
	Getservbyport func(
		port int,
		proto *Char) *SERVENT
	Getservbyname func(
		name *Char,
		proto *Char) *SERVENT
	Getprotobynumber func(
		number int) *PROTOENT
	Getprotobyname func(
		name *Char) *PROTOENT
	WSAStartup func(
		VersionRequested WORD,
		WSAData *WSADATA) int
	WSACleanup      func() int
	WSASetLastError func(
		Error int)
	WSAGetLastError       func() int
	WSAIsBlocking         func() BOOL
	WSAUnhookBlockingHook func() int
	WSASetBlockingHook    func(BlockFunc FARPROC) FARPROC
	WSACancelBlockingCall func() int
	WSAAsyncGetServByName func(
		Wnd HWND,
		Msg U_int,
		name *Char,
		proto *Char,
		buf *Char,
		buflen int) HANDLE
	WSAAsyncGetServByPort func(
		Wnd HWND,
		Msg U_int,
		port int,
		proto *Char,
		buf *Char,
		buflen int) HANDLE
	WSAAsyncGetProtoByName func(
		Wnd HWND,
		Msg U_int,
		name *Char,
		buf *Char,
		buflen int) HANDLE
	WSAAsyncGetProtoByNumber func(
		Wnd HWND,
		Msg U_int,
		number int,
		buf *Char,
		buflen int) HANDLE
	WSAAsyncGetHostByName func(
		Wnd HWND,
		Msg U_int,
		name *Char,
		buf *Char,
		buflen int) HANDLE
	WSAAsyncGetHostByAddr func(
		Wnd HWND,
		Msg U_int,
		addr *Char,
		len int,
		typ int,
		buf *Char,
		buflen int) HANDLE
	WSACancelAsyncRequest func(
		AsyncTaskHandle HANDLE) int
	WSAAsyncSelect func(
		s SOCKET,
		Wnd HWND,
		Msg U_int,
		lEvent LONG) int
	WSAAccept func(
		s SOCKET,
		addr *SOCKADDR,
		addrlen *INT,
		fnCondition *CONDITIONPROC,
		CallbackData DWORD_PTR) SOCKET
	WSACloseEvent func(
		Event WSAEVENT) BOOL
	WSAConnect func(
		s SOCKET,
		name *SOCKADDR,
		namelen int,
		CallerData *WSABUF,
		CalleeData *WSABUF,
		SQOS *QOS,
		GQOS *QOS) int
	WSACreateEvent      func() WSAEVENT
	WSADuplicateSocketA func(
		s SOCKET,
		ProcessId DWORD,
		ProtocolInfo *WSAPROTOCOL_INFOA) int
	WSADuplicateSocketW func(
		s SOCKET,
		ProcessId DWORD,
		ProtocolInfo *WSAPROTOCOL_INFOW) int
	WSAEnumNetworkEvents func(
		s SOCKET,
		EventObject WSAEVENT,
		NetworkEvents *WSANETWORKEVENTS) int
	WSAEnumProtocolsA func(
		Protocols *INT,
		ProtocolBuffer *WSAPROTOCOL_INFOA,
		BufferLength *DWORD) int
	WSAEnumProtocolsW func(
		Protocols *INT,
		ProtocolBuffer *WSAPROTOCOL_INFOW,
		BufferLength *DWORD) int
	WSAEventSelect func(
		s SOCKET,
		EventObject WSAEVENT,
		NetworkEvents LONG) int
	WSAGetOverlappedResult func(
		s SOCKET,
		Overlapped *WSAOVERLAPPED,
		Transfer *DWORD,
		fWait BOOL,
		Flags *DWORD) BOOL
	WSAGetQOSByName func(
		s SOCKET,
		QOSName *WSABUF,
		QOS *QOS) BOOL
	WSAHtonl func(
		s SOCKET,
		hostlong U_long,
		netlong *U_long) int
	WSAHtons func(
		s SOCKET,
		hostshort U_short,
		netshort *U_short) int
	//TODO(t):Callback indirection
	WSAIoctl func(
		s SOCKET,
		IoControlCode DWORD,
		InBuffer *VOID,
		sInBuffer DWORD,
		OutBuffer *VOID,
		sOutBuffer DWORD,
		BytesReturned *DWORD,
		Overlapped *WSAOVERLAPPED,
		CompletionRoutine *WSAOVERLAPPED_COMPLETION_ROUTINE) int
	WSAJoinLeaf func(
		s SOCKET,
		name *SOCKADDR,
		namelen int,
		CallerData *WSABUF,
		CalleeData *WSABUF,
		SQOS *QOS,
		GQOS *QOS,
		Flags DWORD) SOCKET
	WSANtohl func(
		s SOCKET,
		netlong U_long,
		hostlong *U_long) int
	WSANtohs func(
		s SOCKET,
		netshort U_short,
		hostshort *U_short) int
	//TODO(t):Callback indirection
	WSARecv func(
		s SOCKET,
		Buffers *WSABUF,
		BufferCount DWORD,
		NumberOfBytesRecvd *DWORD,
		Flags *DWORD,
		Overlapped *WSAOVERLAPPED,
		CompletionRoutine *WSAOVERLAPPED_COMPLETION_ROUTINE) int
	WSARecvDisconnect func(
		s SOCKET,
		InboundDisconnectData *WSABUF) int
	//TODO(t):Callback indirection
	WSARecvFrom func(
		s SOCKET,
		Buffers *WSABUF,
		BufferCount DWORD,
		NumberOfBytesRecvd *DWORD,
		Flags *DWORD,
		From *SOCKADDR,
		Fromlen *INT,
		Overlapped *WSAOVERLAPPED,
		CompletionRoutine *WSAOVERLAPPED_COMPLETION_ROUTINE) int
	WSAResetEvent func(
		Event WSAEVENT) BOOL
	//TODO(t):Callback indirection
	WSASend func(
		s SOCKET,
		Buffers *WSABUF,
		BufferCount DWORD,
		NumberOfBytesSent *DWORD,
		Flags DWORD,
		Overlapped *WSAOVERLAPPED,
		CompletionRoutine *WSAOVERLAPPED_COMPLETION_ROUTINE) int
	WSASendDisconnect func(
		s SOCKET,
		OutboundDisconnectData *WSABUF) int
	//TODO(t):Callback indirection
	WSASendTo func(
		s SOCKET,
		Buffers *WSABUF,
		BufferCount DWORD,
		NumberOfBytesSent *DWORD,
		Flags DWORD,
		To *SOCKADDR,
		iTolen int,
		Overlapped *WSAOVERLAPPED,
		CompletionRoutine *WSAOVERLAPPED_COMPLETION_ROUTINE) int
	WSASetEvent func(
		Event WSAEVENT) BOOL
	WSASocketA func(
		af int,
		typ int,
		protocol int,
		ProtocolInfo *WSAPROTOCOL_INFOA,
		g GROUP,
		Flags DWORD) SOCKET
	WSASocketW func(
		af int,
		typ int,
		protocol int,
		ProtocolInfo *WSAPROTOCOL_INFOW,
		g GROUP,
		Flags DWORD) SOCKET
	WSAWaitForMultipleEvents func(
		nEvents DWORD,
		Events *WSAEVENT,
		WaitAll BOOL,
		Timeout DWORD,
		Alertable BOOL) DWORD
	WSAAddressToStringA func(
		Address *SOCKADDR,
		AddressLength DWORD,
		ProtocolInfo *WSAPROTOCOL_INFOA,
		AddressString OVString,
		AddressStringLength *DWORD) INT
	WSAAddressToStringW func(
		Address *SOCKADDR,
		AddressLength DWORD,
		ProtocolInfo *WSAPROTOCOL_INFOW,
		AddressString OVString,
		AddressStringLength *DWORD) INT
	WSAStringToAddressA func(
		AddressString VString,
		AddressFamily INT,
		ProtocolInfo *WSAPROTOCOL_INFOA,
		Address *SOCKADDR,
		AddressLength *INT) INT
	WSAStringToAddressW func(
		AddressString VString,
		AddressFamily INT,
		ProtocolInfo *WSAPROTOCOL_INFOW,
		Address *SOCKADDR,
		AddressLength *INT) INT
	WSALookupServiceBegin func(
		Restrictions *WSAQUERYSET,
		ControlFlags DWORD,
		Lookup *HANDLE) INT
	WSALookupServiceNext func(
		Lookup HANDLE,
		ControlFlags DWORD,
		BufferLength *DWORD,
		Results *WSAQUERYSET) INT
	WSANSPIoctl func(
		Lookup HANDLE,
		ControlCode DWORD,
		InBuffer *VOID,
		sInBuffer DWORD,
		OutBuffer *VOID,
		sOutBuffer DWORD,
		BytesReturned *DWORD,
		Completion *WSACOMPLETION) INT
	WSALookupServiceEnd func(
		Lookup HANDLE) INT
	WSAInstallServiceClass func(
		ServiceClassInfo *WSASERVICECLASSINFO) INT
	WSARemoveServiceClass func(
		ServiceClassId *GUID) INT
	WSAGetServiceClassInfo func(
		ProviderId *GUID,
		ServiceClassId *GUID,
		BufSize *DWORD,
		ServiceClassInfo *WSASERVICECLASSINFO) INT
	WSAEnumNameSpaceProviders func(
		BufferLength *DWORD,
		Buffer *WSANAMESPACE_INFO) INT
	WSAGetServiceClassNameByClassId func(
		ServiceClassId *GUID,
		ServiceClassName OVString,
		BufferLength *DWORD) INT
	WSASetService func(
		RegInfo *WSAQUERYSET,
		essoperation WSAESETSERVICEOP,
		ControlFlags DWORD) INT
	//TODO(t):Callback indirection
	WSAProviderConfigChange func(
		NotificationHandle *HANDLE,
		Overlapped *WSAOVERLAPPED,
		CompletionRoutine *WSAOVERLAPPED_COMPLETION_ROUTINE) INT
)

var WinSock2ANSIApis = Apis{
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

var WinSock2UnicodeApis = Apis{
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

var WinSock2Apis = Apis{
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
