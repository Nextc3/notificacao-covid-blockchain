# notificacao-covid-blockchain
Trabalho de conclusão de curso com objetivo construir um sistema de notificação de COVID19 simples utilizando Hyperledger Fabric 2.2.x

Instale Golang, Git, cURL, Docker e Docker Compose

Prefira versão do Golang 1.15 e 1.16 e prefira instalar manualmente sem ser pelo gerenciador de pacotes da distribuição Linux.
Caso ocorra erros utilize o snap.


O aplicativo foi construído para ser executado em ambiente linux Ubuntu 20.04.2 LTS arquitetura 64

Se você está trabalhando com git em Windows (mesmo que seja com interação com máquina virtual Linux ou pretende levar código para Linux posteriormente), precisa adequar retorno de linha e entendimento de diretórios com os seguintes comandos no git:

```
git config --global core.autocrlf false
git config --global core.longpaths true

```
Caso tenha problemas com permissões do Docker isso pode ser resolvido com:

```
sudo setfacl -m "g:docker:rw" /var/run/docker.sock
$sudo addgroup --system docker
$sudo adduser $USER docker
$newgrp docker
```


No seu /home/nomeUsuario

Edite o arquivo .bashrc e coloque as seguintes linhas alterando adequadamente seu caso:
```
export GOPATH=$HOME/go
export GOBIN=$GOPATH/bin
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:/usr/local/go/bin
export FABRIC_RAIZ=$HOME/go/src/github.com/SeuUsuarioNoGitHub/fabric-samples
export PATH=$PATH:$HOME/go/src/github.com/SeuUsuarioNoGitHub/fabric-samples/bin
export FABRIC_CFG_PATH=$FABRIC_RAIZ/config
export HTTP_PORT=8080
```
No seu diretório de go/src/github.com/SeuUsuarioNoGithub execute:

`go get -r github.com/urfave/negroni`

`go get -r github.com/gorilla/mux`

Instale o Fabric Samples 2.2

Clone esse repositório abaixo e coloque o conteúdo dele dentro do diretório de instalação do Fabric sobrescrevendo os diretórios e arquivos existentes
[https://github.com/Nextc3/fabric-samples-2.2](https://github.com/Nextc3/fabric-samples-2.2)

Uma vez dentro desse diretório criado, clone esse repositório notificacao-covid-blockchain

O diretório test-network deve está no mesmo nível de notificacao-covid-blockchain

Entre em notificacao-covid-blockchain.
Execute o arquivo meexecuteumavezantes.sh (Este arquivo já insere ativos default para serem testados. 2 ativos)
Entre dentro do diretório web
Execute o comando

`go run notificacaoapi.go`

ou

`go build noticacaoapi.go`
e execute o arquivo criado

Os endpoints são:

GET http://seuip:8080/notificacao (Lista todas as notificações armazenadas)

GET http://seuip:8080/notificacao/{id} (Traz uma notificação específica com id)

POST http://seuip:8080/notificacao (Passando um JSON é cadastrado uma notificação)

Dentro test-network pode ser testado interações com o chaincode contratointeligente.go pelos executáveis .sh





Estrutura Web pensada usando como referência [Curso de Web com Go](https://github.com/eminetto/pos-web-go) junto com explicações dos vídeos [Curso de Web com Go no Youtube](https://www.youtube.com/playlist?list=PL0qudqr7_CuStQUsf2vtHXMxOp5gl_ENc) 
Agradecimentos ao Elton Minetto. 
