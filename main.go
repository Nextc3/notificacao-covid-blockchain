package main

import (
	"github.com/Nextc3/notificacao-covid-blockchain/cliente"
	"github.com/Nextc3/notificacao-covid-blockchain/implementacaoservico"
)

func main() {
	var cliente *cliente.ClienteBlockchain
	cliente.Contrato.SetContrato(cliente.Conexao.IniciarConexao())
	defer cliente.Conexao.FecharConexao()
	cliente.Contrato.InitLedger()
	meuservico := implementacaoservico.NewServico(cliente)
	



}