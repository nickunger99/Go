package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia)

	fmt.Println(contaDaSilvia.sacar(200))

	fmt.Println(contaDaSilvia)

	// contaDoNicolas := ContaCorrente{"Nicolas",
	// 	589, 223456, 23232.22}

	// fmt.Println(contaDoNicolas)

	// var contaDaCris *ContaCorrente
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.titular = "Cris"
	// contaDaCris.saldo = 500

	// var contaDaCris2 *ContaCorrente
	// contaDaCris2 = new(ContaCorrente)
	// contaDaCris2.titular = "Cris"
	// contaDaCris2.saldo = 500

	// fmt.Println(*contaDaCris)
	// fmt.Println(&contaDaCris)
	// fmt.Println(&contaDaCris2)
	// fmt.Println(*contaDaCris == *contaDaCris2)
	// fmt.Println(contaDaCris == contaDaCris2)
}
