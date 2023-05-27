package entidade

type Notificacao struct {
	Id                  int    `json:"id"`
	Cpf                 string `json:"cpf"`
	Nome                string `json:"nome"`
	DataNascimento      string `json:"dataNascimento"`
	Sexo                string `json:"sexo"`
	Localidade          string `json:"localidade"`
	Doenca              string `json:"doença"`
	DataInicioSintomas  string `json:"dataInicioSintomas"`
	DataDiagnostico     string `json:"dataDiagnostico"`
	DataNotificacao     string `json:"dataNotificacao"`
	InformacoesClinicas string `json:"informacoesClinicas"`

	/*
		IdNotificador             int    `json:"idNotificador"`
		CpfNotificador            string `json:"cpfNotificador"`
		NomeNotificador           string `json:"nomeNotificador"`
		OcupacaoNotificador       string `json:"ocupacaoNotificador"`
		RegistroNotificador       string `json:"registroNotificador"`
		EhProfissionalDeSaude     bool   `json:"ehProfissionalDeSaude"`
		EhProfissionalDeSeguranca bool   `json:"ehProfissionalDeSeguranca"`

		Ocupacao                  string `json:"ocupacao"`

		DataNascimento            string `json:"dataNascimento"`
		Sexo                      string `json:"sexo"`
		Raca                      string `json:"raca"`
		PovoTradicional           bool   `json:"povoTradicional"`
		Cep                       string `json:"cep"`
		Logradouro                string `json:"logradouro"`
		NumeroEndereco            string `json:"numeroEndereco"`
		Complemento               string `json:"complemento"`
		Bairro                    string `json:"bairro"`
		Estado                    string `json:"estado"`
		Municipio                 string `json:"municipio"`
		Telefone                  string `json:"telefone"`
		Email                     string `json:"email"`
		Estrategia                string `json:"estrategia"` //Estrategia estado da pessoa com sintomas ou não
		LocalizacaoTeste          string `json:"localizacaoTeste"`
		DataNotificacao           string `json:"dataNotificacao"`
		Sintomas                  string `json:"sintomas"`
		Condicoes                 string `json:"condicoes"`
		Vacinas                   string `json:"vacinas"`
		TipoDeTeste               string `json:"tipoDeTeste"`
		EstadoDoTeste             string `json:"estadoDoTeste"`
		DataDaColeta              string `json:"dataDaColeta"`
		Lote                      string `json:"lote"`
		Fabricante                string `json:"fabricante"`

	*/
}

/*
	BRANCA = iota + 1
	PRETA
	PARDA
	AMARELA
	INDIGENA
	IGNORADO
*/

/*
		Diagnóstico assistencial (sintomático)
		Busca ativa de assintomático
		Triagem de população específica


	SINTOMATICO   = 1
	ASSINTOMATICO = 2
	TRIAGEM       = 3
*/

//Localização do teste
/*/
	UNIDADEDESAUDE = iota + 1
	LOCALDETRABALHO
	AEROPORTO
	DROGARIA
	ESCOLA
	COMUNIDADE
	OUTROSLOCALIZACAOTESTE
)
*/

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
