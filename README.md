# notificacao-covid-blockchain
Trabalho de conclusão de curso com objetivo construir um sistema de notificação de COVID19 simples utilizando Hyperledger Fabric 2.2

O aplicativo foi construído para ser executado em ambiente linux Ubuntu 20.04.2 LTS arquitetura 64


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





Estrutura Web pensada usando como referência [Curso de Web com Go](https://github.com/eminetto/pos-web-go) junto com explicações dos vídeos [Curso de Web com Go no Youtube](https://www.youtube.com/playlist?list=PL0qudqr7_CuStQUsf2vtHXMxOp5gl_ENc) 
Agradecimentos ao Elton Minetto. 