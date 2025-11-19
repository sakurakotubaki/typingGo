package mylib

import "testing"

func TestHuman(t *testing.T) {
	if Debug {
		t.Skip("Skip Reason")
	}
	p := Person{Name: "Icchy Doe", Age: 30}
	if p.Name != "Icchy Doe" {
		t.Errorf("Expected Name 'Icchy Doe', got %s", p.Name)
	}
	if p.Age != 30 {
		t.Errorf("Expected Age 30, got %d", p.Age)
	}
}
