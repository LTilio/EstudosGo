package main

import (
	"fmt"
	"strings"
)

func main() {

	name := " Leandro Tilio "

	fmt.Println("--------------------------------")

	fmt.Println("Original:", name)
	fmt.Println("Len:", len(name))
	fmt.Println("Sem espaços:", strings.TrimSpace(name))
	fmt.Println("Maiusculo:", strings.ToUpper(name))
	fmt.Println("Minusculo:", strings.ToLower(name))
	fmt.Println("Substituir:", strings.ReplaceAll(name, "Tilio", "Muniz"))
	fmt.Println("Split: ", strings.Split(name, " "))
	fmt.Println("Join: ", strings.Join([]string{"Leandro", "Tilio"}, " "))
	fmt.Println("Contains:", strings.Contains(name, "Tilio"))
	fmt.Println("HasPrefix:", strings.HasPrefix(name, " Leandro"))
	fmt.Println("HasSuffix:", strings.HasSuffix(name, "Tilio "))
	fmt.Println("Replace:", strings.Replace(name, "Tilio", "Muniz", 1))
	fmt.Println("Fields (whitespace, sem vazios):", strings.Fields("  Leandro   Tilio  "))
	fmt.Println("Count 'i' em Tilio:", strings.Count(strings.TrimSpace(name), "i"))
	fmt.Println("Index 'Tilio':", strings.Index(strings.TrimSpace(name), "Tilio"))
	fmt.Println("EqualFold:", strings.EqualFold("leandro", "LEANDRO"))
	fmt.Println("Repeat:", strings.Repeat("Go! ", 3))
	before, after, found := strings.Cut("chave:valor", ":")
	fmt.Println("Cut:", "antes=", before, "depois=", after, "achou=", found)
	fmt.Println("--------------------------------")
	fmt.Println()
	fmt.Println("Formatação de string (Printf → escreve no terminal)")
	fmt.Println("--------------------------------")
	fmt.Printf("O nome é %s e a idade é %d\n", strings.TrimSpace(name), 33)
	fmt.Println("--------------------------------")
	fmt.Println("Formatação com Sprintf (retorna string; você decide onde usar)")
	fmt.Println("--------------------------------")
	msg := fmt.Sprintf("O nome é %s e a idade é %d", strings.TrimSpace(name), 33)
	fmt.Println(msg)
	fmt.Println("--------------------------------")
	fmt.Println("Outros verbos comuns em format strings:")
	pi := 3.14159
	ok := true
	fmt.Printf("%%v (default): %v | %%T (tipo): %T | %%q (string com aspas): %q\n", pi, pi, name)
	fmt.Printf("%%f float: %f | com casas: %.2f | bool: %t | base 2: %b | hex: %#x\n", pi, pi, ok, 42, 42)
	fmt.Println("--------------------------------")

}
