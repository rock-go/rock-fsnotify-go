package fsnotify

import (
	"github.com/fsnotify/fsnotify"
	"github.com/rock-go/rock/lua"
)

type event fsnotify.Event

func (ev event) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "op":
		return lua.S2L(ev.Op.String())
	case "name":
		return lua.S2L(ev.Name)

	case "create":
		return lua.LBool(ev.Op&fsnotify.Create == fsnotify.Create)

	case "write":
		return lua.LBool(ev.Op&fsnotify.Write == fsnotify.Write)

	case "remove":
		return lua.LBool(ev.Op&fsnotify.Remove == fsnotify.Remove)

	case "rename":
		return lua.LBool(ev.Op&fsnotify.Rename == fsnotify.Rename)

	case "chmod":
		return lua.LBool(ev.Op&fsnotify.Chmod == fsnotify.Chmod)

	}
	return lua.LNil
}
