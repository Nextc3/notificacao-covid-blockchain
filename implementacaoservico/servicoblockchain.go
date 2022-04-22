package implementacaoservico

import (
	"fmt"
	"github.com/Nextc3/notificacao-covid-blockchain/cliente"
	"github.com/Nextc3/notificacao-covid-blockchain/consulta"
	"github.com/Nextc3/notificacao-covid-blockchain/entidade"
	"log"
)

/*

type InterfaceServico struct {
	ObterTodos () ([]*Notificacao, error)
	Obter(Id int) (*Notificacao, error)
	Salvar(n *Notificacao) error
}

*/

type Servico struct {
	Cliente *cliente.ClienteBlockchain
}

func NewServico(c *cliente.ClienteBlockchain) *Servico {
	return &Servico{
		Cliente: c,
	}
}

func (s *Servico) ObterTodos() ([]*entidade.Notificacao, error) {
	log.Println("Obtendo todas as Notificações")
	var result []*consulta.ResultadoConsulta
	var ns []*entidade.Notificacao

	var err error
	result, err = s.Cliente.Contrato.ObterTodasNotificacoes(true)
	if err != nil {
		log.Fatalf("Falhou em obter todas as Notificações: %v", err)
		return nil, err
	}
	for _, value := range result {
		ns = append(ns, value.Ativo)
		log.Println("Adicionando Notificações ao resultado")
		fmt.Printf("Notificação %s adicionada", value.Chave)
	}
	return ns, nil

}
func (s *Servico) Obter(Id int) (*entidade.Notificacao, error) {
	log.Println("Obtendo Notificação")
	var n *entidade.Notificacao

	n, err := s.Cliente.Contrato.ConsultarNotificacao(true, Id)
	if err != nil {
		log.Println("Erro em obter Notificação")
		return nil, err
	}

	return n, nil

}
func (s *Servico) Salvar(n *entidade.Notificacao) error {
	log.Println("Salvando notificação")
	s.Cliente.Contrato.CriarNotificacao(n)

	return nil

}
