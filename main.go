package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteThiago := clientes.Titular{
		Nome: "Thiago", CPF: "333.888.000-00", Profissao: "Desenvolvedor",
	}
	contaThiago := contas.ContaCorrente{
		Titular:       clienteThiago,
		NumeroAgencia: 123,
		NumeroConta:   123456,
		Saldo:         595548.45,
	}

	fmt.Println(contaThiago)

}
