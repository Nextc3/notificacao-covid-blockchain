package consulta

import "github.com/Nextc3/notificacao-covid-blockchain/entidade"

type ResultadoConsulta struct {
	Chave string `json:"chave"`
	Ativo *entidade.Notificacao `json:"ativo"`
}
