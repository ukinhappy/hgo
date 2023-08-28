package hgo

import (
	"testing"
)

func TestHgo(t *testing.T) {
	hgo := New()
	hgo.Init()
	hgo.Run()
}
