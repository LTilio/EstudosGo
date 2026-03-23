package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile() {
	file, err := os.ReadFile("./24-read-write-files/nomes.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo: ", err)
	}
	fmt.Println(string(file))
}

func lineReadFile() {
	file, err := os.Open("./24-read-write-files/nomes.txt")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo: ", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println("....................................")
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo: ", err)
		return
	}

}

func writeFile() {
	data := []string{"Snorlax", "Chikorita"}
	content := strings.Join(data, ";")

	// os.WriteFile grava bytes brutos ([]byte), não []string. strings.Join junta cada string
	// do slice com "\n" no meio → "Snorlax\nChikorita" (uma linha por nome, como no .txt).
	// []byte(...) converte essa string para UTF-8 em memória, que é o que o arquivo guarda.
	//
	// 0644 é permissão em octal (o zero à esquerda em Go = literal octal, não “64 em decimal”).
	// Três dígitos = três grupos (dono, grupo, outros).
	// Cada dígito 0–7 soma bits de leitura(4)+gravação(2)+execução(1):
	err := os.WriteFile("./24-read-write-files/pokemons.txt", []byte(content), 0644)

	if err != nil {
		fmt.Println("Erro ao escrever o arquivo: ", err)
		return
	}
	fmt.Println("Arquivo escrito corretamente")
}

func writeAllreadyExistsFile() {
	// os.OpenFile abre (ou cria) um arquivo para manipulação manual:
	// - O caminho "./24-read-write-files/pokemons.txt" indica o arquivo a abrir.
	// - os.O_APPEND faz com que tudo escrito seja acrescentado no final do arquivo, sem apagar o que já existe.
	// - os.O_WRONLY abre somente para escrita (não leitura).
	// - 0644 define as permissões: leitura e escrita para o dono, leitura para grupo e outros.
	file, err := os.OpenFile("./24-read-write-files/pokemons.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo para escrita: ", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(";Meowth")
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo: ", err)
		return
	}

	fmt.Println("Arquivo escrito corretamente")
}

func main() {

	fmt.Println("Ready and write files")
	fmt.Println("------------------------------------")
	fmt.Println("leitura de arquivo completo: ")
	readFile()
	fmt.Println("------------------------------------")
	fmt.Println("leitura de arquivo linha por linha: ")
	lineReadFile()
	fmt.Println("------------------------------------")
	fmt.Println("Escrevendo arquivo: ")
	writeFile()
	fmt.Println("------------------------------------")
	fmt.Println("Escrevendo arquivo já existente: ")
	writeAllreadyExistsFile()

}
