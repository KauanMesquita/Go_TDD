package main

import (
	"reflect"
	"testing"
)

func TestSomaTudo(t *testing.T) {

	recebido := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(recebido, esperado) {
		t.Errorf("recebido %v esperado %v", recebido, esperado)
	}
}
