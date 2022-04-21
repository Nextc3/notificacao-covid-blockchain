package cliente

import (
	"fmt"
	"log"
	"os"
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

func (c *Conexao) iniciarConexao() *gateway.Contract {
	setarDiscovery()
	wallet, err := getWallet()
	err = credenciarWallet(wallet, err)

	if err != nil {
		log.Fatalf("Falhou em credenciar a Wallet %v", err)
	}

	ccpPath := getCaminhoConnectionOrg1Yaml()
	gw := getGateway(ccpPath, wallet)
	c.gateway = gw
	//defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Falhou em pegar a network: %v", err)
	}

	contrato := network.GetContract("helloworld")

	return contrato

}

func (c *Conexao) fecharConexao() {
	c.gateway.Close()
}

func getGateway(caminho string, wallet *gateway.Wallet) *gateway.Gateway {
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
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)
	return caminho
}
func credenciarWallet(wallet *gateway.Wallet, err error) error {
	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("Falhou em colocar credenciais na wallet: %v", err)
		}
	}

	fmt.Println("Credenciou Wallet")
	return err
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
