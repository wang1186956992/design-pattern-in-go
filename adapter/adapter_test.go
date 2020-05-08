package main

import (
	"testing"
)

func TestAdapter(t *testing.T) {
	p := Phone{
		V5V: &VoltageAdapter{},
	}
	p.Charing()
}
