package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaThiago := contas.ContaCorrente{
		Titular:       "Thiago",
		NumeroAgencia: 123,
		NumeroConta:   123456,
		Saldo:         500,
	}

	contaPamela := contas.ContaCorrente{
		Titular:       "Pamela",
		NumeroAgencia: 321,
		NumeroConta:   654321,
		Saldo:         500,
	}
	fmt.Println(contaThiago)
	fmt.Println(contaPamela)

	contaPamela.Transferir(300, &contaThiago)

	fmt.Println(contaThiago)
	fmt.Println(contaPamela)

}
