package consulta

type ResultadoConsulta struct {
	Chave string `json:"chave"`
	Ativo *Notificacao
}
