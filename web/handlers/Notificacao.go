package handlers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Nextc3/notificacao-covid-blockchain/entidade"
	"github.com/Nextc3/notificacao-covid-blockchain/interfaceservico"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

//a função recebe como terceiro parâmetro a interface
//ou seja, ela pode receber qualquer coisa que implemente a interface
//isso é muito útil para escrevermos testes, ou podermos substituir toda a
//implementação da regra de negócios
func CriarNotificacaoHandlers(r *mux.Router, n *negroni.Negroni, meuservico interfaceservico.Iservico) {
	r.Handle("/notificacao", n.With(
		negroni.Wrap(obterTodasNotificacoes(meuservico))),
	).Methods("GET", "OPTIONS")

	r.Handle("/notificacao/{id}", n.With(
		negroni.Wrap(obterNotificacao(meuservico)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/notificacao", n.With(
		negroni.Wrap(salvarNotificacao(meuservico)),
	)).Methods("POST", "OPTIONS")
}

func salvarNotificacao(meuservico interfaceservico.Iservico) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//adicionar em midleware
		w.Header().Set("Content-Type", "application/json")

		//vamos pegar os dados enviados pelo usuário via body
		var notificacao entidade.Notificacao
		err := json.NewDecoder(r.Body).Decode(&notificacao)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}
		//@TODO precisamos validar os dados antes de salvar na base de dados. Fazer posteriormente
		err = meuservico.Salvar(&notificacao)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

/*
Para testar:
curl http://localhost:portaHttp/notificacao/1
*/
func obterNotificacao(meuservico interfaceservico.Iservico) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Código duplicado em todos o handlers. Posteriormente mudar usando midleware
		w.Header().Set("Content-Type", "application/json")

		//vamos pegar o ID da URL
		//na definição do protocolo http, os parâmetros são enviados no formato de texto
		//por isso precisamos converter em int

		//obter variáveis na requisição
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(formatJSONError(err.Error()))
			return
		}
		notificacao, err := meuservico.Obter(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(formatJSONError(err.Error()))
			return
		}
		//vamos converter o resultado em JSON e gerar a resposta
		err = json.NewEncoder(w).Encode(notificacao)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(formatJSONError("Erro converter em JSON"))
			return
		}
	})
}

/*
Para testar:
curl -H 'Accept: application/json' http://localhost:portaUsada/notificacao
*/
func obterTodasNotificacoes(meuservico interfaceservico.Iservico) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//analisa o que o usuário requisitou via headers
		switch r.Header.Get("Accept") {
		case "application/json":
			obterTodasNotificacoesJSON(w, meuservico)
		default:
			obterTodasNotificacoesHTML(w, meuservico)
		}

	})
}

func obterTodasNotificacoesJSON(w http.ResponseWriter, meuservico interfaceservico.Iservico) {
	w.Header().Set("Content-Type", "application/json")
	todos, err := meuservico.ObterTodos()
	if err != nil {
		//passa um erro como resposta. Sinaliza também no cabeçalho
		w.Write(formatJSONError(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//vamos converter o resultado em JSON e gerar a resposta o response
	//http.ResponseWriter implementa interface ResponseWriter assim como io.Writer	
	err = json.NewEncoder(w).Encode(todos)
	if err != nil {
		w.Write(formatJSONError("Erro convertendo em JSON"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func obterTodasNotificacoesHTML(w http.ResponseWriter, meuservico interfaceservico.Iservico) {
	//Setando cabeçalho
	w.Header().Set("Content-Type", "text/html")
	ts, err := template.ParseFiles(
		"./web/templates/header.html",
		"./web/templates/index.html",
		"./web/templates/rodape.html")
	if err != nil {
		http.Error(w, "Erro na criação de templates"+err.Error(), http.StatusInternalServerError)
		return
	}
	todos, err := meuservico.ObterTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	/*
		dado:= tipo Variavel

	*/
	dado := struct {
		//tipo
		Titulo       string
		Notificacoes []*entidade.Notificacao
	}{
		//Variavel
		Titulo:       "Notificações",
		Notificacoes: todos,
	}
	err = ts.Lookup("index.html").ExecuteTemplate(w, "index", dado)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
