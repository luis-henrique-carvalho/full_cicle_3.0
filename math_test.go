package main

import "testing"

func TestSomar(t *testing.T) {
	resultado := Somar(2, 3)
	esperado := 5
	if resultado != esperado {
		t.Errorf("Somar(2, 3) = %d; esperado %d", resultado, esperado)
	}
}
