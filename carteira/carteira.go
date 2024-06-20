package main

import "fmt"

type Carteira struct {
	saldo int
}

func (c *Carteira) Depositar(quantidade int) {
	fmt.Printf("o endereço do saldo no Depositar é %v \n", &c.saldo)
	c.saldo += quantidade
}

func (c *Carteira) Saldo() int {
	return c.saldo
}
