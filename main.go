package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	for _, cep := range os.Args[1:] {
		req, erro := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if erro != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", erro)
		}
		defer req.Body.Close()

		res, erro := io.ReadAll(req.Body)
		if erro != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler a requisição: %v\n", erro)
		}

		var dados CEP

		erro = json.Unmarshal(res, &dados)
		if erro != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v", erro)
		}

		fmt.Println(dados)
	}
}
