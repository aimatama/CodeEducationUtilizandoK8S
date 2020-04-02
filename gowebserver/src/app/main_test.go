package main

import "testing"

func TestGreeting(t *testing.T) {

	var mensagemOK string
	var mensagem string

	mensagemOK = "<b>Code.education Rocks!</b>"
	mensagem = "Code.education Rocks!"

	got := greeting(mensagem)
	if got != mensagemOK {
		t.Fatalf("O resultado esperado deveria ter sido '%s' porém o retorno foi igual a '%s'.", mensagemOK, got)
	}
}
