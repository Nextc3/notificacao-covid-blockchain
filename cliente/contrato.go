package cliente

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type Contrato struct {
	contrato *gateway.Contract
}

func (c *Contrato) setContrato(g *gateway.Contract) {
	c.contrato = g
}

/*
//Alterar Notificacao está fora do escopo do trabalho e do objetivo do blockchain, mas caso seja
necessário deixei esse atualização e remoção da notificações da blockchain
func (c *Contrato) atualizarNotificacao(id string, novaPessoa string) {
	log.Println("--> Transação de Submit: ChangeOiPessoa OI1, transfere para um novo dono Val Bandeira")
	contract := c.contrato
	_, err := contract.SubmitTransaction("ChangeOiPessoa", id, novaPessoa)
	if err != nil {
		log.Fatalf("Falhou em ChangeOiPessoa Transação de Submit: %v", err)
	}
}

*/
func (c *Contrato) CriarNotificacao(n *Notificacao) {
	log.Println("--> Transação de Submit: CriarNotificacao, cria ativos do tipo Notificacao")
	contract := c.contrato
	nEmBytes,_ := json.Marshal(n)
	nString := string(nEmBytes)
	result, err := contract.SubmitTransaction("criarNotificacao", nString)
	if err != nil {
		log.Fatalf("Falhou a transação de Criar Notificacao SUBMIT (altera estado da ledger) transação: %v", err)
	}
	log.Println(string(result))
}
func (c *Contrato) ExisteNotificacao(registrar bool, id int) string {
	//a variável registrar diz se quero registrar na ledger a consulta 
	idString := strconv.Itoa(id)
	contrato := c.contrato
	var result []byte
	var err error
	log.Println("--> Transação Evaluate e Submit: ExisteNotificacao, função que retorna um boleano se achou o ativo na ledger")

	if registrar {
		
		log.Println("Registrando transação de existeNotificacao na ledger")
		
		result, err = contrato.SubmitTransaction("existeNotificacao", idString)
	} else {
		log.Println("NÃO registrando transação de existeNotificacao na ledger")

		result, err = contrato.EvaluateTransaction("existeNotificacao", idString)
	}

	if err != nil {
		log.Fatalf("Falhou em ExisteNotificacao Transação Evaluate: %v\n", err)
	}
	return string(result)
}
func (c *Contrato) ConsultarNotificacao(registrar bool, id int) *Notificacao {

	log.Println("--> Transação ConsultarNotificacao, função retorna um ativo")
	var n *Notificacao
	var err error
	idString := strconv.Itoa(id)
	contract := c.contrato
	if registrar {
		fmt.Println("Você escolheu registrar transação")
		n, err = contract.SubmitTransaction("consultarNotificacao", idString)
	} else {
		fmt.Println("Você escolheu não registrar transação")
		n, err = contract.EvaluateTransaction("consultarNotificacao", idString)
	}

	if err != nil {
		log.Fatalf("Falhou em Transação consultarNotificacao : %v\n", err)
	}
	return n
}
func (c *Contrato) ObterTodasNotificacoes(registrar bool) ([]*ResultadoConsulta, error) {
	
		contract := c.contrato
		log.Println("--> Transação ObterTodasNotificacoes, função que retorna todos os ativos na ledger")
		var result []*ResultadoConsulta
		var err error
		if registrar {
			result, err = contract.SubmitTransaction("obterTodasNotificacoes")
		} else {
			result, err = contract.EvaluateTransaction("obterTodasNotificacoes")
		}
	
		if err != nil {
			log.Fatalf("Falhou a getTodosOis transação: %v", err)
		}

		return result,nil 

}
func (c *Contrato) initLedger() {
	contract := c.contrato
	log.Println("--> Transação de Submit: InitLedger, função cria o conjunto inicial de ativos no razão. Para estudo")
	_, err := contract.SubmitTransaction("initLedger")
	if err != nil {

		log.Fatalf("Falhou em InitLedger SUBMIT (altera estado da ledger) %v", err)
	}
}
//Funções de configuração 

/* main
log.Println("============ minha primeira aplicação em golang ============")

	var conexao Conexao
	var contrato Contrato
	contrato.setContrato(conexao.iniciarConexao())
	defer conexao.fecharConexao()
	contrato.initLedger()

	/*
		Saudacao:  saudacao,
		Despedida: despedida,
		Oidenovo:  oidenovo,
		Pessoa:	pessoa
	*/
	//log.Println("============ fim da minha primeira aplicação em golang ============")

*/

