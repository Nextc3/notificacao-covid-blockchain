package main

import (
	"github.com/Nextc3/notificacao-covid-blockchain/cliente"
	"github.com/Nextc3/notificacao-covid-blockchain/implementacaoservico"
	"github.com/Nextc3/notificacao-covid-blockchain/web/handlers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var c *cliente.ClienteBlockchain
	c.Contrato.SetContrato(c.Conexao.IniciarConexao())
	defer c.Conexao.FecharConexao()
	c.Contrato.InitLedger()
	meuservico := implementacaoservico.NewServico(c)

	r := mux.NewRouter()
	//middlewares - código que vai ser executado em todas as requests
	//aqui podemos colocar logs, inclusão e validação de cabeçalhos, etc
	n := negroni.New(
		negroni.NewLogger(),
	)
	//handlers
	handlers.CriarNotificacaoHandlers(r, n, meuservico)

	//static files
	fileServer := http.FileServer(http.Dir("./web/static"))
	r.PathPrefix("/static/").Handler(n.With(
		negroni.Wrap(http.StripPrefix("/static/", fileServer)),
	)).Methods("GET", "OPTIONS")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// used to health check, will return 200
	})

	http.Handle("/", r)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + os.Getenv("HTTP_PORT"),
		Handler:      http.DefaultServeMux,
		ErrorLog:     log.New(os.Stderr, "logger: ", log.Lshortfile),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
