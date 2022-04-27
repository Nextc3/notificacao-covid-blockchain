package handlers

import (
	"github.com/Nextc3/notificacao-covid-blockchain/interfaceservico"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func CriarNotificacaoHandlers(r *mux.Router, n *negroni.Negroni, meuservico interfaceservico.Iservico) {

}
