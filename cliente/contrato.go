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
func (c *Contrato) criarNotificacao(n *Notificacao) {
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
func (c *Contrato) existeNotificacao(registrar bool, id int) string {
	//a variável registrar diz se quero registrar na ledger a consulta 
	idString := strconv.Itoa(id)
	contrato := c.contrato
	var result []byte
	var err error
	log.Println("--> Transação Evaluate e Submit: ExisteNotificacao, função que retorna um boleano se achou o ativo na ledger")

	if registrar {
		result, err = contrato.SubmitTransaction("existeNotificacao", idString)
	} else {
		result, err = contrato.EvaluateTransaction("existeNotificacao", idString)
	}

	if err != nil {
		log.Fatalf("Falhou em ExisteNotificacao Transação Evaluate: %v\n", err)
	}
	return string(result)
}
func (c *Contrato) consultarNotificacao(registrar bool, id int) string {

	log.Println("--> Transação ConsultarNotificacao, função retorna um ativo")
	var result []byte
	var err error
	idString := strconv.Itoa(id)
	contract := c.contrato
	if registrar {
		fmt.Println("Você escolheu registrar transação")
		result, err = contract.SubmitTransaction("QueryOi", idString)
	} else {
		fmt.Println("Você escolheu não registrar transação")
		result, err = contract.EvaluateTransaction("QueryOi", idString)
	}

	if err != nil {
		log.Fatalf("Falhou em Transação consultarNotificacao : %v\n", err)
	}
	return string(result)
}

