package entidade

type Notificacao struct {
	Id                        int                  `json:"id"`
	CidadaoNotificador        Notificador          `json:"idNotificador"`
	TemCPF                    bool                 `json:"temCpf"`
	EhProfissionalDeSaude     bool                 `json:"ehProfissionalDeSaude"`
	EhProfissionalDeSeguranca bool                 `json:"ehProfissionalDeSeguranca"`
	Cpf                       string               `json:"cpf"`
	Ocupacao                  string               `json:"ocupacao"`
	Nome                      string               `json:"nome"`
	DataNascimento            string               `json:"dataNascimento"`
	Sexo                      TipoGenero           `json:"sexo"`
	Raca                      TipoRaca             `json:"raca"`
	PovoTradicional           bool                 `json:"povoTradicional"`
	Cep                       string               `json:"cep"`
	Logradouro                string               `json:"logradouro"`
	NumeroEndereco            string               `json:"numeroEndereco"`
	Complemento               string               `json:"complemento"`
	Bairro                    string               `json:"bairro"`
	Estado                    string               `json:"estado"`
	Municipio                 string               `json:"municipio"`
	Telefone                  string               `json:"telefone"`
	Email                     string               `json:"email"`
	Estrategia                TipoEstrategia       `json:"estrategia"` //Estrategia estado da pessoa com sintomas ou não
	LocalizacaoTeste          TipoLocalizacaoTeste `json:"localizacaoTeste"`
	DataNotificacao           string               `json:"dataNotificacao"`
	Sintomas                  map[string]bool      `json:"sintomas"`
	Condicoes                 map[string]bool      `json:"condicoes"`
	Vacinas                   map[string]bool      `json:"vacinas"`
	Teste                     []TesteCovid         `json:"testeCovid"`
	Contatos                  []ContatoNonitorado  `json:"contatos"`
}

const (
	MASCULINO = false
	FEMININO  = true
)

type TipoGenero bool

const (
	BRANCA = iota + 1
	PRETA
	PARDA
	AMARELA
	INDIGENA
	IGNORADO
)

type TipoRaca uint8

const (
	/*
		Diagnóstico assistencial (sintomático)
		Busca ativa de assintomático
		Triagem de população específica

	*/
	SINTOMATICO   = 1
	ASSINTOMATICO = 2
	TRIAGEM       = 3
)

type TipoEstrategia uint8

//Localização do teste
const (
	UNIDADEDESAUDE = iota + 1
	LOCALDETRABALHO
	AEROPORTO
	DROGARIA
	ESCOLA
	COMUNIDADE
	OUTROSLOCALIZACAOTESTE
)

type TipoLocalizacaoTeste uint8

//tipo de teste

/*
Sintomas:
	"assintomatico": false,
	"coriza": false,
	"disturbiosOlfativos": false,
	"disturbiosGustativos": false,
	"dorDeCabeca":false,
	"tosse": false,
	"febre": false,
	"dispneia": false,
	"dorDeGarganta": false,
	"outros": false,

*/

/*
Possíveis condições:
Doenças respiratórias crônicas descompensadas
Doenças cardíacas crônicas
Diabetes
Doenças renais crõnicas em estágio avançado (graus 3, 4 ou 5)
Imunossupressão
Gestante
Portador de doenças cromossômicas ou estado de fragilidade
Puérpera (até 45 dias do parto)
Obsesidade
Outros
*/
