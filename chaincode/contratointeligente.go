package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/Nextc3/notificacao-covid-blockchain/consulta"
	"github.com/Nextc3/notificacao-covid-blockchain/entidade"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type ContratoInteligente struct {
	contractapi.Contract
}

//Para saber a scruct que está sendo utilizada, por favor, veja no pacote Entidades

func (c *ContratoInteligente) InitLedger(contexto contractapi.TransactionContextInterface) error {
	//método inicial. Normalmente para inserir ativos de testes

	return nil
}

//Cria notificação
func (c *ContratoInteligente) CriarNotificacao(contexto contractapi.TransactionContextInterface, notificacao string) error {

	notificacaoEmBytes := []byte(notificacao)
	var n entidade.Notificacao
	_ = json.Unmarshal(notificacaoEmBytes, &n)

	//Chave do estado é Notificacao + Id da notificacao

	return contexto.GetStub().PutState("Notificacao"+strconv.Itoa(n.Id), notificacaoEmBytes)
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

//Consulta se Oi existe
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
	defer resultadoIteracao.Close()

	results := []*consulta.ResultadoConsulta{}

	for resultadoIteracao.HasNext() {
		queryResponse, err := resultadoIteracao.Next()

		if err != nil {
			return nil, err
		}

		notificacao := new(entidade.Notificacao)
		_ = json.Unmarshal(queryResponse.Value, &notificacao)

		queryResult := consulta.ResultadoConsulta{Chave: queryResponse.Key, Ativo: notificacao}
		results = append(results, &queryResult)
	}

	return results, nil
}

//Caso tivesse um main nesse arquivo
/*
func main() {
	// See chaincode.env.example
	/*
		config := ServerConfig{
			CCID:    os.Getenv("CHAINCODE_ID"),
			Address: os.Getenv("CHAINCODE_SERVER_ADDRESS"),
		} */
/*
	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Erro em criar helloworld chaincode: %s", err.Error())
		return
	}
*/
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
/*
		if err := chaincode.Start(); err != nil {
			fmt.Printf("Erro em criar helloworld chaincode: %s", err.Error())
		}
	}

*/
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
