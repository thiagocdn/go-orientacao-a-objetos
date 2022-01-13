package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaThiago := contas.ContaPoupanca{}

	contaThiago.Depositar(100)
	contaThiago.Depositar(200)
	contaThiago.Sacar(150)

	fmt.Println(contaThiago.ObterSaldo())

	PagarBoleto(&contaThiago, 60)

	fmt.Println(contaThiago.ObterSaldo())

}
