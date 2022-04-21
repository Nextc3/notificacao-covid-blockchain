package main

import (
	"encoding/json"
	"fmt"
	
)


type MinhaStructure struct {
	Key SegundaStrutura  `json:"key"`

}

type SegundaStrutura struct {
	Nome string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Sexo bool `json:"sexo"`

}

func main() {
	novo := MinhaStructure{
		Key: SegundaStrutura{
			Nome: "Caio",
			Sobrenome: "Costa",
			Sexo: true,
		},
	}
	fmt.Println("Estrutura intecta")
	fmt.Println(novo)
	
	
	emFormatoJson,_ := json.Marshal(novo)
	fmt.Println("Json em string")
	fmt.Println(string(emFormatoJson))

	novo2 := string(emFormatoJson)
	var novo3 MinhaStructure
	
	
	_ = json.Unmarshal([]byte(novo2),&novo3)
	fmt.Println("String Json")
	fmt.Println(string([]byte(novo2)))
	fmt.Println("Estrutura")
	fmt.Println(novo3)



	





}