package interfaceservico

import (
	"github.com/Nextc3/notificacao-covid-blockchain/entidade"
)

type Service interface {
	ObterTodos() ([]*entidade.Notificacao, error)
	Obter(id int) (entidade.Notificacao, error)
	Salvar(n entidade.Notificacao) error
}
