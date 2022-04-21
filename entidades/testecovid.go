package entidades

type TesteCovid struct {
	Id            int               `json:"id"`
	TipoDeTeste   string            `json:"tipoDeTeste"`
	EstadoDoTeste TipoEstadoDoTeste `json:"estadoDoTeste"`
	DataDaColeta  string            `json:"dataDaColeta"`
	Resultado     TipoResultado     `json:"resultado"`
	Lote          string            `json:"lote"`
	Fabricante    string            `json:"fabricante"`
}

//estado do teste
const (
	SOLICITADO = iota + 1
	COLETADO
	CONCLUIDO
	EXAMENAOSOLICITADO
)

type TipoEstadoDoTeste uint8

//Resultado do teste
const (
	REAGENTE = iota + 1
	NAOREAGENTE
	INCONCLUSIVO
)

type TipoResultado uint8
