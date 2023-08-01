package awk

import (
	"github.com/vela-ssoc/vela-kit/vela"
	"github.com/vela-ssoc/vela-kit/lua"
)

var xEnv vela.Environment

func newAwkL(L *lua.LState) int {
	raw := L.CheckString(1)
	sep := L.IsString(2)
	n := L.IsInt(3)
	L.Push(newAwk(raw, sep, n))
	return 1
}

func WithEnv(env vela.Environment) {
	xEnv = env
	xEnv.Set("awk", lua.NewExport("vela.awk.export", lua.WithFunc(newAwkL)))
}
