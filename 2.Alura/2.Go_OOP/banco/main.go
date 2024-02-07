package main

import (
	"Go/2.Alura/2.Go_OOP/banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(200)
	PagarBoleto(&contaDaLuisa, 150)
	
	fmt.Println(contaDaLuisa.ObterSaldo())

	// contaExemplo := contas.ContaCorrente{}
	// contaExemplo.Depositar(100)

	// fmt.Println(contaExemplo.ObterSaldo())

	// clienteBruno := clientes.Titular{
	// 	Nome:      "Bruno",
	// 	CPF:       "123.111.123.12",
	// 	Profissao: "Desenvolvedor"}
	// contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 123, NumeroConta: 123456, Saldo: 100}

	// fmt.Println(contaDoBruno)

	// contaDaSilvia := contas.ContaCorrente{Titular: "Silvia", Saldo: 300}
	// contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	// status := contaDaSilvia.Transferir(200, &contaDoGustavo)

	// fmt.Println(status)
	// fmt.Println(contaDaSilvia)
	// fmt.Println(contaDoGustavo)

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
