package entidades

type Notificador struct {
	Id             int    `json:"id"`
	Email          string `json:"email"`
	Cpf            string `json:"cpf"`
	DataNascimento string `json:"dataNascimento"`
	Nome           string `json:"nome"`
	NomeDaMae      string `json:"nomeDaMae"`
	Estado         string `json:"estado"`
	Municipio      string `json:"municipio"`
	Telefone       string `json:"telefone"`
	Ocupacao       string `json:"ocupacao"`
}
