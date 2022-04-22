package entidade

type ContatoNonitorado struct {
	Id                int                 `json:"id"`
	Nome              string              `json:"nome"`
	Estado            string              `json:"estado"`
	Municipio         string              `json:"municipio"`
	Cpf               string              `json:"cpf"`
	Telefone1         string              `json:"telefone1"`
	Telefone2         string              `json:"telefone2"`
	DataUltimoContato string              `json:"dataUltimoContato"`
	RelacaoComOCaso   TipoRelacaoComOCaso `json:"relacaoComOCaso"`
}

//Relações com a pessoa com suspeita de covid

const (
	DOMICILIAR = iota + 1
	ESCOLAR
	EVENTOSOCIAL
	FAMILIAR
	LABORAL
	OUTROSRELACAO
)

type TipoRelacaoComOCaso uint8
