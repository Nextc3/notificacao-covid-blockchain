package entidades

type TesteCovid struct {
	Id            int    `json:"id"`
	TipoDeTeste   string `json:"tipoDeTeste"`
	EstadoDoTeste uint8  `json:"estadoDoTeste"`
	DataDaColeta  string `json:"dataDaColeta"`
	Resultado     uint8  `json:"resultado"`
	Lote          string `json:"lote"`
	Fabricante    string `json:"fabricante"`
}

//estado do teste
const (
	SOLICITADO = iota + 1
	COLETADO
	CONCLUIDO
	EXAMENAOSOLICITADO
)

//Resultado do teste
const (
	REAGENTE = iota + 1
	NAOREAGENTE
	INCONCLUSIVO
)
