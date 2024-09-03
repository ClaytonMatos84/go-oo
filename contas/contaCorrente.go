package contas

import "oo/clientes"

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito < 0 {
		return "Valor do depósito deve ser maior do que zero", c.saldo
	}

	c.saldo += valorDeposito
	return "Depósito realizado com sucesso", c.saldo
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor > c.saldo {
		return false
	}

	c.Sacar(valor)
	contaDestino.Depositar(valor)
	return true
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
