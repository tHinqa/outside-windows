## *outside-windows*: Win32 API definitions for the Go language *(PRERELEASE)*

#### Fire off quick questions to [@tHinqa](http://twitter.com/tHinqa) on Twitter

#### Covered by the same licence conditions as Go is

### For use with [*outside*](https://github.com/tHinqa/outside)

### Based on naming and definitions in the Windows SDK

To use, include:

```go

		import . "github.com/tHinqa/outside-windows/types"

```

and

```go

		import windows "github.com/tHinqa/outside-windows"
```

or any combination of

```go

		import "github.com/tHinqa/outside-windows/winbase"
		import "github.com/tHinqa/outside-windows/winuser"
		import "github.com/tHinqa/outside-windows/commctrl"
		import "github.com/tHinqa/outside-windows/dde"
		import "github.com/tHinqa/outside-windows/mmsystem"
		import "github.com/tHinqa/outside-windows/nb30"
		import "github.com/tHinqa/outside-windows/ole2"
		import "github.com/tHinqa/outside-windows/shellapi"
		import "github.com/tHinqa/outside-windows/wincon"
		import "github.com/tHinqa/outside-windows/wingdi"
		import "github.com/tHinqa/outside-windows/winnetwk"
		import "github.com/tHinqa/outside-windows/winnls"
		import "github.com/tHinqa/outside-windows/winreg"
		import "github.com/tHinqa/outside-windows/winsock"
		import "github.com/tHinqa/outside-windows/winsock2"
		import "github.com/tHinqa/outside-windows/winsvc"
		import "github.com/tHinqa/outside-windows/types"
```
