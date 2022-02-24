package fsnotify

import (
	"github.com/go-playground/validator/v10"
	"github.com/rock-go/rock/auxlib"
	"github.com/rock-go/rock/lua"
	"github.com/rock-go/rock/pipe"
)

type config struct {
	name string `validate:"required"`

	path   []string
	pipe   []pipe.Pipe
	onErr  pipe.Pipe
	co     *lua.LState
}

func newConfig(L *lua.LState) *config {
	tab := L.CheckTable(1)
	cfg := &config{
		co: xEnv.Clone(L),
	}

	tab.Range(func(key string, val lua.LValue) {
		switch key {
		case "name":
			cfg.name = auxlib.CheckProcName(val, L)
		case "path":
			switch val.Type() {
			case lua.LTString:
				cfg.path = []string{val.String()}
			case lua.LTTable:
				cfg.path = auxlib.LTab2SS(val.(*lua.LTable))
			default:
				//todo
			}
		default:
			//todo
		}
	})

	if err := cfg.valid(); err != nil {
		L.RaiseError("%v", err)
		return nil
	}
	return cfg
}

func (cfg *config) valid() error {
	valid := validator.New()
	return valid.Struct(cfg)
}
