package main

import (
	"fmt"
	"silicon/hardware"
)

func main() {
	var lote []hardware.Memoria
	for {
		var opcao int
		fmt.Println("1. Cadastrar Memória")
		fmt.Println("2. Listar e Testar Lote")
		fmt.Println("9. Sair")
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
			fmt.Print("Digite a Latencia: ")
			fmt.Scan(&l)
			fmt.Print("Digite a Voltagem: ")
			fmt.Scan(&v)

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

		if opcao == 9 {
			fmt.Println("Saindo...")
			break
		}
	}
}