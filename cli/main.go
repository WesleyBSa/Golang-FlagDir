package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	channel := flag.String("channelName", "WbsaTestes", "Nome do teste")
	flag.Parse()

	// Criar um diretório com o nome do canal
	err := os.Mkdir(*channel, 0755)
	if err != nil {
		fmt.Println("Erro ao criar o diretório:", err)
		return
	}

	fmt.Println("Diretório criado:", *channel)
}
