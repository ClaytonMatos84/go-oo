package main

import (
	"fmt"
	"oo/clientes"
	"oo/contas"
)

func main() {
	// cc := contas.ContaCorrente{Titular: clientes.Titular{Nome: "José"}, NumeroAgencia: 1, NumeroConta: 123, Saldo: 55.5}
	// fmt.Println(cc)

	// cc2 := contas.ContaCorrente{clientes.Titular{Nome: "Maria"}, 1, 456, 100.5}
	// fmt.Println(cc2)

	// var cc3 *ContaCorrente
	cc3 := new(contas.ContaCorrente)
	cc3.Titular = clientes.Titular{Nome: "João", CPF: "123.456.789-10", Profissao: "Dev"}
	cc3.NumeroAgencia = 1
	cc3.Depositar(75.5)
	fmt.Println(*cc3)

	cc4 := new(contas.ContaCorrente)
	cc4.Titular = clientes.Titular{Nome: "Bia", CPF: "987.456.789-10", Profissao: "UX"}
	cc4.NumeroAgencia = 1
	cc4.Depositar(90.0)
	fmt.Println(*cc4)

	// fmt.Println(cc3.Sacar(40.0))
	// fmt.Println(cc3.saldo)
	// msg, saldo := cc3.Depositar(-100)
	// fmt.Println(msg, saldo)

	status := cc4.Transferir(30.5, cc3)
	fmt.Println(status)
	fmt.Println(cc3.ObterSaldo())
	fmt.Println(cc4.ObterSaldo())

	contas.PagarConta(cc4, 10)
	fmt.Println(cc4.ObterSaldo())
}
