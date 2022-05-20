package cliente

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type Conexao struct {
	gateway *gateway.Gateway
}

func setarDiscovery() {
	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("Erro em setar DISCOVERY_AS_LOCALHOST como variável de ambiente: %v", err)
	}
	fmt.Println("êxito em setarDiscovery")

}

func (c *Conexao) IniciarConexao() (*gateway.Contract, *gateway.Gateway) {
	fmt.Println("Iniciando conexão")
	fmt.Println("Setar discovery")
	setarDiscovery()
	fmt.Println("Obtendo Wallet")
	wallet, err := getWallet()
	fmt.Println("Credenciando Wallet")
	err = credenciarWallet(wallet)

	if err != nil {
		log.Fatalf("Falhou em credenciar a Wallet %v", err)
	}

	fmt.Println("Obtendo Caminho da Conexão Org1 em Yaml")
	caminho := getCaminhoConexaoOrg1Yaml()
	fmt.Println("Obtendo Gateway")
	gw := getGateway(caminho, wallet)

	c.gateway = gw
	//defer gw.Close()

	fmt.Println("Obtendo a rede com o nome do canal do gateway")
	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Falhou em pegar a network: %v", err)
	}
	fmt.Println("Obtendo contrato inteligente da rede/canal")
	contrato := network.GetContract("contratointeligente")
	fmt.Println("Retornando contrato inteligente")
	return contrato, gw

}

func (c *Conexao) FecharConexao() {
	fmt.Println("Fechando conexão do gateway")
	c.gateway.Close()
}

func getGateway(caminho string, wallet *gateway.Wallet) *gateway.Gateway {
	fmt.Println("Estabelecendo uma conexão no gateway")
	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(caminho))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		log.Fatalf("Falhou em conectar com o gateway: %v", err)
	}
	return gw
}
func getCaminhoConexaoOrg1Yaml() string {
	caminho := filepath.Join(
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)
	return caminho
}
func credenciarWallet(wallet *gateway.Wallet) error {
	fmt.Println("Credenciar Wallet")
	if !wallet.Exists("appUser") {
		err := populateWallet(wallet)
		if err != nil {
			log.Fatalf("Falhou em colocar credenciais na wallet: %v", err)
		}
	}

	fmt.Println("Credenciou Wallet")
	return nil
}
func getWallet() (*gateway.Wallet, error) {
	fmt.Println("Obtendo Wallet")
	wallet, err := gateway.NewFileSystemWallet("gateway")
	if err != nil {
		log.Fatalf("Falhou em criar wallet: %v", err)
	}
	fmt.Println("Função de obter wallet finalizada com sucesso")
	return wallet, err
}
func populateWallet(wallet *gateway.Wallet) error {
	log.Println("Método crucial para funcionamento: populando Wallet")
	log.Println("============ Populating wallet ============")
	log.Println("Obtendo credenciais")
	credPath := filepath.Join(
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "User1@org1.example.com-cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}
	fmt.Println("pegou as credenciais em de cert.pem")
	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	fmt.Println("pegou a chave")
	if len(files) != 1 {
		return fmt.Errorf("a pasta keystore deve conter um arquivo")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	fmt.Println("pegou arquivo da chave")
	fmt.Println(string(files[0].Name()))
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	log.Println("Obtendo a identidade de Org1 junto com a chave")
	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))
	fmt.Println("Registrando na Wallet a identidade de Org1 e usuário")
	return wallet.Put("appUser", identity)
}
