package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Deposito inválido", c.saldo
	}
}

func main() {
	contaThiago := ContaCorrente{
		titular:       "Thiago",
		numeroAgencia: 123,
		numeroConta:   123456,
		saldo:         500,
	}
	fmt.Println(contaThiago.saldo)

	fmt.Println(contaThiago.Depositar(200))

	fmt.Println(contaThiago.saldo)

}
