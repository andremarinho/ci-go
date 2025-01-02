package main

import "testing"

func Testsoma(t *testing.T) {

	resultado := soma(10, 10)

    esperado := 20
    if resultado != esperado {
        t.Errorf("Soma(10, 10) = %d; esperado %d", resultado, esperado)
    }
}