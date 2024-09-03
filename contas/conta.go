package contas

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarConta(conta verificarConta, valor float64) {
	conta.Sacar(valor)
}
