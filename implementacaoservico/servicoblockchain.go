package implementacaoservico

import (
	"fmt"
	"github.com/Nextc3/notificacao-covid-blockchain/cliente"
	"github.com/Nextc3/notificacao-covid-blockchain/consulta"
	"github.com/Nextc3/notificacao-covid-blockchain/entidade"
	"log"
)

/*

type Service interface {
	ObterTodos() ([]*entidade.Notificacao, error)
	Obter(id int) (entidade.Notificacao, error)
	Salvar(n entidade.Notificacao)  error
}

*/

type Servico struct {
	Contrato *cliente.Contrato
}

func NewService(c *cliente.Contrato) *Servico {
	return &Servico{
		Contrato: c,
	}
}

func (s *Servico) ObterTodos() ([]*entidade.Notificacao, error) {
	log.Println("Obtendo todas as Notificações")
	var result []consulta.ResultadoConsulta
	var ns []*entidade.Notificacao

	var err error
	result, err = s.Contrato.ObterTodasNotificacoes(true)
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

//Obter responde uma entidade solicitada, caso não exista retorna uma Notificação vazia
func (s *Servico) Obter(id int) (entidade.Notificacao, error) {
	log.Println("Serviço implemetado: Obtendo Notificação")

	n, err := s.Contrato.ConsultarNotificacao(true, id)

	return n, err

}
func (s *Servico) Salvar(n entidade.Notificacao) error {
	log.Println("Salvando notificação")
	s.Contrato.CriarNotificacao(n)

	return nil

}
