package main

import (
	"fmt"
	"silicon/hardware" 
	"encoding/json"
	"os"
)

func main() {
	var lote []hardware.Memoria
	arquivo, err := os.ReadFile("estoque.json")

    if err == nil {
        json.Unmarshal(arquivo, &lote)
        fmt.Println("Backup carregado! Itens recuperados: ", len(lote))
    } else {
        fmt.Println("Nenhum backup anterior, Começando do zero.")
    }
	for {
		var opcao int
		fmt.Println("1. Cadastrar Memória")
		fmt.Println("2. Listar e Testar Lote")
		fmt.Println("3. Importar para JSON")
		fmt.Println("4. Sair")
		fmt.Print("Escolha: ")
		fmt.Scan(&opcao)

		if opcao == 1 {
			var m string
			var f int
			var l int
			var v float64
			fmt.Print("Digite a Marca: ")
			fmt.Scan(&m)
			fmt.Print("Digite a Frequencia: ")
			fmt.Scan(&f)

			if f <= 1800 {
			    fmt.Println("Frequência não pode ser menor que 1800mhz")
				continue }
			if f > 6000 {
				fmt.Println("Nenhuma DDR4 chega nisso.") 
				continue }

			fmt.Print("Digite a Latencia: ")
			fmt.Scan(&l)

			if l <= 10 {
				fmt.Println("Essa Latencia não existe.")
				continue }
			if l > 22 {
				fmt.Println("Essa Latencia não existe.")
				continue }	

			fmt.Print("Digite a Voltagem: ")
			fmt.Scan(&v)

            if v >= 2.0 {
				fmt.Println("Essa Voltagem não existe.")
			    continue }
			if v <= 0.8 {
				fmt.Println("Essa Voltagem não existe.")
				continue}
			
			lote = append(lote, hardware.Memoria{
				Marca:      m,
				Frequencia: f,
				Latencia:   l,
				Voltagem:   v,
			})
			fmt.Println("Cadastrado!")
		}

		if opcao == 2 {
			for _, item := range lote {
				if item.EhBdie() {
					fmt.Printf("ACHEI OURO! %s é B-Die (CL %d)\n", item.Marca, item.Latencia)
				} else {
					fmt.Printf("Lixo! %s é genérica.\n", item.Marca)
				}
			}
		}
        if opcao == 3 {
			fmt.Println("Converter para JSON.")
			bytes, err := json.MarshalIndent(lote, "", "  ")
			errArquivo := os.WriteFile("estoque.json", bytes, 0644)
			if err != nil {
                fmt.Println("Erro ao salvar arquivo:", errArquivo)
            } else {
                fmt.Println("Sucesso! Arquivo 'estoque.json' criado.")
			}
		}
		if opcao == 4 {
			fmt.Println("Saindo...")
			break
		}
	}
}