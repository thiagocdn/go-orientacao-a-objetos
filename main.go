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

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar((valorTransferencia))
		return true
	} else {
		return false
	}

}

func main() {
	contaThiago := ContaCorrente{
		titular:       "Thiago",
		numeroAgencia: 123,
		numeroConta:   123456,
		saldo:         500,
	}

	contaPamela := ContaCorrente{
		titular:       "Pamela",
		numeroAgencia: 321,
		numeroConta:   654321,
		saldo:         500,
	}
	fmt.Println(contaThiago)
	fmt.Println(contaPamela)

	contaPamela.Transferir(300, &contaThiago)

	fmt.Println(contaThiago)
	fmt.Println(contaPamela)

}
