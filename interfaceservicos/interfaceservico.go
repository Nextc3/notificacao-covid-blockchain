package interfaceservicos

import "entidades/Notificacao"




type InterfaceServico struct {
	ObterTodos () ([]*Notificacao, error)
	Obter(Id int) (*Notificacao, error)
	Salvar(n *Notificacao) error
}

