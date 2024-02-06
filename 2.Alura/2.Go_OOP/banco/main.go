package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoNicolas := ContaCorrente{"Nicolas",
		589, 223456, 23232.22}

	fmt.Println(contaDoNicolas)

}
