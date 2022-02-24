package fsnotify

import (
	"github.com/rock-go/rock/lua"
	"github.com/rock-go/rock/xbase"
	"reflect"
)

var (
	xEnv   *xbase.EnvT
	typeof = reflect.TypeOf((*watch)(nil)).String()
)


/*

*/

func newLuaFsnotify(L *lua.LState) int {
	cfg := newConfig(L)
	proc := L.NewProc(cfg.name, typeof)
	if proc.IsNil() {
		proc.Set(newWatch(cfg))
	} else {
		w := proc.Data.(*watch)
		xEnv.Free(w.cfg.co)
		w.cfg = cfg
	}

	L.Push(proc)
	return 1
}

func LuaInjectApi(env *xbase.EnvT) {
	xEnv = env
	env.Set("fsnotify", lua.NewFunction(newLuaFsnotify))
}
