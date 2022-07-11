package main

import "testing"

func TestFoo(t *testing.T) {
	t.Log("success test")
}

func TestBar(t *testing.T) {
	t.Error("fail test")
}
