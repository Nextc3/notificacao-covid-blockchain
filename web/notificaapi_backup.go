package main

import (
	"fmt"
	"github.com/Nextc3/notificacao-covid-blockchain/cliente"
	"github.com/Nextc3/notificacao-covid-blockchain/implementacaoservico"
	"github.com/Nextc3/notificacao-covid-blockchain/web/handlers"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	//Pra executar go run
	//Main só pra costurar e compor coisas necessárias pra camada de negócio
	var conex cliente.Conexao
	var contra cliente.Contrato
	contratoGateway, gw := conex.IniciarConexao()
	defer gw.Close()

	if contratoGateway == nil && gw == nil {
		log.Fatalf("Falha em começar uma conexão. No método principal")
	}
	fmt.Println(contratoGateway)
	contra.SetContrato(contratoGateway)

	meuservico := implementacaoservico.NewService(contra)

	//roteador pra fazer controle de rotas
	roteador := mux.NewRouter()
	//Negroni: middlewares - código que vai ser executado em todas as requests.
	//Empilhados para serem usados quando quiser em várias requisições
	//aqui podemos colocar logs, inclusão e validação de cabeçalhos, etc
	ngroni := negroni.New(
		negroni.NewLogger(),
	)
	//handlers
	handlers.CriarNotificacaoHandlers(roteador, ngroni, &meuservico)

	/*
		retorna um handler que atende solicitações HTTP com o conteúdo do sistema de
		 arquivos enraizado na raiz.Como um caso especial, o servidor de arquivos
		 retornado redireciona qualquer solicitação que termine em "/index.html"
		  para o mesmo caminho, sem o "index.html" final.
		  Para usar a implementação do sistema de arquivos do sistema operacional, é usado http.Dir:
	*/
	//fileServer := http.FileServer(http.Dir("./web/static"))
	/*o método PathPrefix registra uma nova rota ("/static/")
	Na nova rota criada setado um Http Handler (manipulador de requisições que a respondem)
	o http handler em questão é retornado pela combinação de vários handlers feitos pelo negroni
	com o método With(). O método With recebe conversão feita do Http Handlers em Negroni Handlers
	pelo metódo Wrap(Com ele posso chamar funções e passar qualquer coisas pra essas funções
		como por exemplo negroni.Wrap(funçãoQualquer(parâmetro))).
	 O metódo Wrap recebe um Handler de StripPrefix
	que atende a solicitações HTTP removendo o prefixo fornecido do caminho da URL da
	solicitação  e invocando o manipulador handler h(no caso fileServer).
	Methods:  adiciona um conexão para métodos HTTP. Caso não seja colocado retorna 404
	*/
	//
	/*
		roteador.PathPrefix("/static/").Handler(ngroni.With(
			negroni.Wrap(http.StripPrefix("/static/", fileServer)),
		)).Methods("GET", "OPTIONS")

	*/
	//Para ser verificado de tempos em tempos nas clouds
	roteador.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// used to health check, will return 200
	})

	//Tudo que vier da raiz vou tratar com o roteador criado
	http.Handle("/", roteador)
	//criando um servidor http
	//Usa goroutines pra cada requisição que chegar
	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":8080",                                        //porta que o servidor http está setado
		Handler:      http.DefaultServeMux,                           //raiz criada no http.Handle()
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile), //log como saída de erro padrão no terminal
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
