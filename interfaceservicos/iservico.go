package interfaceservicos

import (
	"github.com/Nextc3/notificacao-covid-blockchain/entidade"
)

type Iservico interface {
	ObterTodos() ([]*entidade.Notificacao, error)
	Obter(Id int) (*entidade.Notificacao, error)
	Salvar(n *entidade.Notificacao) error
}
