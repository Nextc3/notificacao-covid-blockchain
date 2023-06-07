package main

import (
	"encoding/json"
	"fmt"
	"github.com/Nextc3/notificacao-covid-blockchain/consulta"
	"github.com/Nextc3/notificacao-covid-blockchain/entidade"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"log"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ContratoInteligente struct {
	contractapi.Contract
}

var ultimoId int

//gerarObjetoTeste retorna dois objetos com valores de teste sendo que o segundo pode ser ignorado

func (c *ContratoInteligente) GetUltimoId() int {
	return ultimoId
}

//Para saber a scruct que está sendo utilizada, por favor, veja no pacote Entidades
//InitLedger já inicia a ledger com duas notificações
/*
func (c *ContratoInteligente) InitLedger(contexto contractapi.TransactionContextInterface) error {
	//método inicial. Normalmente para inserir ativos de testes

	n1, n2 := c.gerarObjetoTeste()

	noti1, _ := json.Marshal(n1)
	noti2, _ := json.Marshal(n2)
	err := c.CriarNotificacao(contexto, string(noti2))
	if err != nil {
		return err
	}

	return c.CriarNotificacao(contexto, string(noti1))

}
*/
//Cria notificação
func (c *ContratoInteligente) CriarNotificacao(contexto contractapi.TransactionContextInterface, notificacao string) error {

	notificacaoEmBytes := []byte(notificacao)
	var n entidade.Notificacao
	_ = json.Unmarshal(notificacaoEmBytes, &n)

	//Chave do estado é Notificacao + Id da notificacao
	//Cuidado para não salvar uma Notificação com mesmo Id pois são utilizados para salvar na ledger
	err := contexto.GetStub().PutState("Notificacao"+strconv.Itoa(n.Id), notificacaoEmBytes)
	if err != nil {
		log.Fatalf("Erro ao salvar na ledger %s", err)
	}

	return err
}

//retorna uma notificacao
func (c *ContratoInteligente) ConsultarNotificacao(contexto contractapi.TransactionContextInterface, idNotificacao string) (*entidade.Notificacao, error) {
	notificacaoEmBytes, err := contexto.GetStub().GetState("Notificacao" + idNotificacao)

	if err != nil {
		return nil, fmt.Errorf("Falha em consultar em Notificação na Ledger com GetState %s", err.Error())
	}

	if notificacaoEmBytes == nil {
		return nil, fmt.Errorf("Notificacao%s não existe", idNotificacao)
	}

	notificacao := new(entidade.Notificacao)
	_ = json.Unmarshal(notificacaoEmBytes, notificacao)

	return notificacao, nil
}

//Consulta se Notificação existe
func (s *ContratoInteligente) ExisteNotificacao(contexto contractapi.TransactionContextInterface, idNotificacao string) (bool, error) {
	notificacaoEmBytes, err := contexto.GetStub().GetState("Notificacao" + idNotificacao)
	if err != nil {
		return false, fmt.Errorf("falhou em consultar a existência da Notificacao: %v", err)
	}

	return notificacaoEmBytes != nil, nil
}
func (c *ContratoInteligente) ObterTodasNotificacoes(contexto contractapi.TransactionContextInterface) ([]*consulta.ResultadoConsulta, error) {
	startKey := ""
	endKey := ""
	// GetStateByRange retorna um iterador de intervalo sobre um conjunto de chaves no
	// ledger. O iterador pode ser usado para iterar sobre todas as chaves
	// entre startKey (inclusive) e endKey (exclusivo).
	// No entanto, se o número de chaves entre startKey e endKey for maior que
	// totalQueryLimit (definido em core.yaml), esse iterador não poderá ser usado
	// para buscar todas as chaves (os resultados serão limitados pelo totalQueryLimit).
	// As chaves são retornadas pelo iterador em ordem lexical. Observe
	// ​​que startKey e endKey podem ser uma string vazia, o que implica um intervalo ilimitado
	// consulta no início ou no fim.
	// Chame Close() no objeto StateQueryIteratorInterface retornado quando terminar.
	// A consulta é executada novamente durante a fase de validação para garantir que o conjunto de resultados
	// não foi alterado desde o endosso da transação (leituras fantasma detectadas).

	resultadoIteracao, err := contexto.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer func(resultadoIteracao shim.StateQueryIteratorInterface) {
		err := resultadoIteracao.Close()
		if err != nil {

		}
	}(resultadoIteracao)

	results := []*consulta.ResultadoConsulta{}
	ultimoId = 0
	for resultadoIteracao.HasNext() {
		ultimoId++
		queryResponse, err := resultadoIteracao.Next()

		if err != nil {
			return nil, err
		}

		notificacao := new(entidade.Notificacao)
		_ = json.Unmarshal(queryResponse.Value, &notificacao)

		queryResult := consulta.ResultadoConsulta{Chave: queryResponse.Key, Ativo: notificacao}
		/*fmt.Println("Chave :")
		fmt.Println(queryResponse.Key)
		fmt.Println("Notificação sendo adicionado aos resultados:")
		fmt.Println(notificacao)*/
		results = append(results, &queryResult)
		/*fmt.Println("Notificação adicionada:")
		fmt.Println(notificacao.Id)*/
	}

	return results, nil
}

//Caso tivesse um main nesse arquivo

func main() {
	// See chaincode.env.example
	/*
		config := ServerConfig{
			CCID:    os.Getenv("CHAINCODE_ID"),
			Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
		} */

	contrato, err := contractapi.NewChaincode(new(ContratoInteligente))

	if err != nil {
		fmt.Printf("Erro em criar contrato inteligente contrato: %s", err.Error())
		return
	}

	/*
		server := &shim.ChaincodeServer{
			CCID:    config.CCID,
			Address: config.Address,
			CC:      chaincode,
			TLSProps: shim.TLSProperties{
				Disabled: true,
			},
		}
		if err := server.Start(); err != nil {
			fmt.Printf("Erro em estartar helloworld chaincode: %s", err.Error())
		}
	*/

	if err := contrato.Start(); err != nil {
		fmt.Printf("Erro em criar contratointeligente chaincode: %s", err.Error())
	}

}

//Caso seja necessário criação de um método de alteração das notificações
//Nessa caso de exemplo um ativo chamado Oi é usada para armazenar
/* Mudar notificação está fora do escopo desse trabalho
o método abaixo ainda tá com corpo de teste
// ChangeOiPessoa atualiza o campo Pessoa da Oi com id fornecido no estado mundial
func (s *SmartContract) ChangeOiPessoa(ctx contractapi.TransactionContextInterface, oiNumber string, newPessoa string) error {
	oi, err := s.QueryOi(ctx, oiNumber)

	if err != nil {
		return err
	}

	oi.Pessoa = newPessoa

	oiAsBytes, _ := json.Marshal(oi)

	return ctx.GetStub().PutState(oiNumber, oiAsBytes)
}
*/
//Deletar um ativo não faz parte do escopo desse trabalho
//ainda com escrita de teste
/*
func (s *SmartContract) DeleteOi(contexto contractapi.TransactionContextInterface, oiNumber string) error {
	exists, err := s.ExisteOi(contexto, oiNumber)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("o ativo %s não ecsiste", oiNumber)
	}

	return contexto.GetStub().DelState(oiNumber)

}
*/
