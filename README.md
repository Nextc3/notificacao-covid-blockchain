# notificacao-covid-blockchain
Trabalho de conclusão de curso com objetivo construir um sistema de notificação de COVID19 simples utilizando Hyperledger Fabric 2.2.x

Instale Golang, Git, cURL, Docker e Docker Compose

Prefira versão do Golang 1.15 e 1.16 e prefira instalar manualmente sem ser pelo gerenciador de pacotes da distribuição Linux.


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
export FABRIC_RAIZ=$HOME/go/src/github.com/Nextc3/fabric-samples
export PATH=$PATH:$HOME/go/src/github.com/Nextc3/fabric-samples/bin
export FABRIC_CFG_PATH=$FABRIC_RAIZ/config

```
No seu diretório de go/src/github.com/SeuUsuarioNoGithub execute:
`go get -r github.com/urfave/negroni`
`go get -r github.com/gorilla/mux`







Estrutura Web pensada usando como referência [Curso de Web com Go](https://github.com/eminetto/pos-web-go) junto com explicações dos vídeos [Curso de Web com Go no Youtube](https://www.youtube.com/playlist?list=PL0qudqr7_CuStQUsf2vtHXMxOp5gl_ENc) 
Agradecimentos ao Elton Minetto. 
