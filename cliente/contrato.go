package cliente

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/Nextc3/notificacao-covid-blockchain/consulta"
	"github.com/Nextc3/notificacao-covid-blockchain/entidade"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type Contrato struct {
	contrato *gateway.Contract
}

func (c *Contrato) SetContrato(g *gateway.Contract) {
	c.contrato = g
}

// CriarNotificacao /*
//CriarNotificacao transforma em um JSON e depois em uma String e envia com SubmitTransaction
func (c *Contrato) CriarNotificacao(n entidade.Notificacao) {
	log.Println("--> Transação de Submit: CriarNotificacao, cria ativos do tipo Notificacao")
	contract := c.contrato

	resultado, err := contract.EvaluateTransaction("getUltimoId")
	if err != nil {
		log.Println("erro em buscar o último Id")
	}
	aux, _ := strconv.Atoi(string(resultado))
	aux++
	n.Id = aux

	nEmBytes, _ := json.Marshal(n)
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

//ConsultarNotificacao retorna uma notificação se encontrada, caso contrário devolve notificação vazia
func (c *Contrato) ConsultarNotificacao(registrar bool, id int) (entidade.Notificacao, error) {

	log.Println("--> Transação ConsultarNotificacao, função retorna um ativo")
	var nEmBytes []byte
	var n entidade.Notificacao
	var err error
	idString := strconv.Itoa(id)
	contract := c.contrato
	deuCerto, err := strconv.ParseBool(c.ExisteNotificacao(registrar, id))
	if err != nil {
		log.Fatalf("Erro em converter String para Bool e saber se existe notificação")
	}
	if !deuCerto {
		log.Println("Não foi encontrado Notificação buscada. Retornando Notificação vazia")
		return c.obterVazia(), nil
	}
	if registrar {
		fmt.Println("Você escolheu registrar transação")
		nEmBytes, err = contract.SubmitTransaction("consultarNotificacao", idString)
	} else {
		fmt.Println("Você escolheu não registrar transação")
		nEmBytes, err = contract.EvaluateTransaction("consultarNotificacao", idString)
	}

	if err != nil {
		log.Fatalf("Falhou em Transação consultarNotificacao : %v\n", err)
	}
	_ = json.Unmarshal(nEmBytes, &n)
	return n, err
}
func (c *Contrato) ObterTodasNotificacoes(registrar bool) ([]consulta.ResultadoConsulta, error) {

	contract := c.contrato
	log.Println("--> Contrato: Transação ObterTodasNotificacoes, função que retorna todos os ativos na ledger")
	var resultEmBytes []byte
	var result []consulta.ResultadoConsulta
	var err error
	if registrar {
		log.Println("Consultando contratointeligente")
		fmt.Println(c)
		fmt.Println(c.contrato)
		resultEmBytes, err = contract.SubmitTransaction("ObterTodasNotificacoes")
	} else {
		resultEmBytes, err = contract.EvaluateTransaction("obterTodasNotificacoes")
	}
	log.Println("Obteve resultado do contrato inteligente")

	if err != nil {
		log.Fatalf("Falhou a getTodosOis transação: %v", err)
	}

	_ = json.Unmarshal(resultEmBytes, &result)

	return result, nil

}
func (c *Contrato) obterVazia() entidade.Notificacao {
	return entidade.Notificacao{}

}
func (c *Contrato) InitLedger() {
	contract := c.contrato
	log.Println("--> Transação de Submit: InitLedger, função cria o conjunto inicial de ativos no razão. Para estudo")
	log.Println("Por enquanto está vazia")
	_, err := contract.SubmitTransaction("initLedger")
	if err != nil {

		log.Fatalf("Falhou em InitLedger SUBMIT (altera estado da ledger) %v", err)
	}
}
