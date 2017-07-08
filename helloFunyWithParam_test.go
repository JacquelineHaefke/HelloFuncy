package main

import "testing"

func TestSayHello(t *testing.T) {
	if Hello("Ted") != "Hello Ted!" {
		t.Error("Unerwartetes Ergebnis.")
	}
}
