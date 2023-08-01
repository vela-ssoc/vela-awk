package awk

import (
	"github.com/vela-ssoc/vela-kit/lua"
	"strings"
)

type awk struct {
	data []string
	sep  string
	n    int
}

func (a awk) String() string                         { return strings.Join(a.data, ",") }
func (a awk) Type() lua.LValueType                   { return lua.LTObject }
func (a awk) AssertFloat64() (float64, bool)         { return 0, false }
func (a awk) AssertString() (string, bool)           { return "", false }
func (a awk) AssertFunction() (*lua.LFunction, bool) { return nil, false }
func (a awk) Peek() lua.LValue                       { return a }

func (a awk) Meta(L *lua.LState, key lua.LValue) lua.LValue {
	id := lua.CheckNumber(L, key)
	n := len(a.data)
	if n == 0 {
		return lua.LNil
	}

	var idx int
	if id < 0 {
		idx = n + int(id)
	} else {
		idx = int(id) - 1
	}

	if idx < 0 || idx > n-1 {
		return lua.LNil
	}
	return lua.S2L(a.data[idx])
}

func (a awk) Index(L *lua.LState, key string) lua.LValue {
	switch key {
	case "size":
		return lua.LInt(len(a.data))
	}
	return lua.LNil
}

func newAwk(data, sep string, n int) awk {
	var a []string
	if sep == "" {
		sep = " "
	}

	if n <= 0 {
		a = strings.Split(data, sep)
	} else {
		a = strings.SplitN(data, sep, n)
	}
	return awk{data: a, sep: sep, n: n}
}
