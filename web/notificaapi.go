package main

import (
	"encoding/json"
	"fmt"
	"github.com/Nextc3/notificacao-covid-blockchain/cliente"
	"github.com/Nextc3/notificacao-covid-blockchain/entidade"
	"github.com/Nextc3/notificacao-covid-blockchain/implementacaoservico"
	"github.com/Nextc3/notificacao-covid-blockchain/servico"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
func formatJSONError(mensagem string) []byte {
	appErro := struct {
		Mensagem string `json:"mensagem"`
	}{
		mensagem,
	}
	response, err := json.Marshal(appErro)
	if err != nil {
		return []byte(err.Error())
	}
	return response
}

func NotificacaoHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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

	//-------
	//

	sid := strings.TrimPrefix(r.URL.Path, "/notificacao/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		obterNotificacao(w, r, id, &meuservico)
	case r.Method == "GET":
		obterTodasNotificacoes(w, r, &meuservico)
	case r.Method == "POST":
		salvarNotificacao(w, r, &meuservico)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func obterNotificacao(w http.ResponseWriter, r *http.Request, id int, meuservico servico.Service) {
	//enableCors(&w)

	var n entidade.Notificacao
	n, err := meuservico.Obter(id)
	if err != nil {
		w.Write([]byte("Não encontrada Notificação"))
		w.WriteHeader(http.StatusNotFound)
		w.Write(formatJSONError(err.Error()))
		return
	}
	json, _ := json.Marshal(n)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}
func obterTodasNotificacoes(w http.ResponseWriter, r *http.Request, meuservico servico.Service) {
	todos, err := meuservico.ObterTodos()
	if err != nil {
		//passa um erro como resposta. Sinaliza também no cabeçalho
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		return
	}

	json, _ := json.Marshal(todos)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))

}
func salvarNotificacao(w http.ResponseWriter, r *http.Request, meuservico servico.Service) {
	w.Header().Set("Content-Type", "application/json")

	//vamos pegar os dados enviados pelo usuário via body
	var notificacao entidade.Notificacao
	err := json.NewDecoder(r.Body).Decode(&notificacao)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(formatJSONError(err.Error()))
		return
	}

	err = meuservico.Salvar(notificacao)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(formatJSONError(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func main() {

	http.HandleFunc("/notificacao/", NotificacaoHandler)
	log.Println("Executando meu NotificaAPI repaginado...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
