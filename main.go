package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaThiago := ContaCorrente{
		titular:       "Thiago",
		numeroAgencia: 123,
		numeroConta:   123456,
		saldo:         357582.89,
	}
	fmt.Println(contaThiago)

}
