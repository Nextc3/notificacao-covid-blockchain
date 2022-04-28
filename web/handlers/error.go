package handlers

import (
	"encoding/json"
)

//esta função gera um erro no formato esperado
//pelo http.ResponseWriter e em um formato
//json
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