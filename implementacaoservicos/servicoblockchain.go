package implementacaoservicos



/*

type InterfaceServico struct {
	ObterTodos () ([]*Notificacao, error)
	Obter(Id int) (*Notificacao, error)
	Salvar(n *Notificacao) error
}

*/

type Servico struct {
	Cliente *ClienteBlockchain
}

func NewServico(c *ClienteBlockchain) *Servico {
	return &Servico{
		Cliente: c,

	}
}

func (s *Servico) ObterTodos()([]*Notificacao, error) {
	var result []*Notificacao

	

}
