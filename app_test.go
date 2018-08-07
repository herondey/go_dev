package main

import (
	"testing"
)

func TestAppName(t *testing.T) {
	expect := "kuso appli"
	actual := AppName()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
