package _const

import "testing"

const c1 = 2

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

func TestConst(t *testing.T) {
	t.Log(c1)
	t.Log(Monday, Tuesday, Wednesday)
}
