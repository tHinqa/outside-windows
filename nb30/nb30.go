// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

//Package netbios provides the API definition for accessing
//netbios in netapi32.dll.
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
