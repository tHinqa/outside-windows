package nb30

//TODO(t): on 7/8 there is a netbios.dll with ep Netbios too

import (
	. "github.com/tHinqa/outside"
	. "github.com/tHinqa/outside-windows/types"
	_ "github.com/tHinqa/outside/win32/netapi32"
)

var (
	Netbios func(pncb *NCB) UCHAR
)

var Nb30Apis = Apis{{"Netbios", &Netbios}}
