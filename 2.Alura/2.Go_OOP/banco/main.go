package main

import (
	"Go/2.Alura/2.Go_OOP/banco/contas"
	"fmt"
)

func main() {
	contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	status := contaDaSilvia.Transferir(200, &contaDoGustavo)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)

	// contaDaSilvia := ContaCorrente{}
	// contaDaSilvia.titular = "Silvia"
	// contaDaSilvia.saldo = 500

	// fmt.Println(contaDaSilvia)

	// fmt.Println(contaDaSilvia.sacar(200))

	// fmt.Println(contaDaSilvia)

	// status, valor := contaDaSilvia.depositar(2000)
	// fmt.Println(status, valor)

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
