package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadFile(filePath, novoValor string, numeroLinha int) error {

	// Ler o arquivo INI existente
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("Erro ao abrir o arquivo: %v", err)
	}
	defer file.Close()

	// Estatísticas sobre o arquivo
	// fileInfo, err := file.Stat()
	// if err != nil {
	// 	log.Fatal("Erro ao obter informações sobre o arquivo:", err)
	// }

	// // Leia o conteúdo do arquivo
	// data := make([]byte, fileInfo.Size())
	// n, err := file.Read(data)
	// if err != nil {
	// 	log.Fatal("Erro ao ler o arquivo:", err)
	// }

	// // Imprima o conteúdo lido
	// fmt.Printf("Conteúdo do arquivo:\n%s\n", data[:n])

	// // Contar as linhas no conteúdo
	// numLinhas := contarLinhasNoConteudo(data[:n])

	// Imprimir o número de linhas
	// fmt.Printf("Número de linhas no conteúdo do arquivo: %d\n", numLinhas)

	//TESTE
	fmt.Println("Ta aqui?")
	fmt.Println(novoValor)
	conteudo, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao ler o conteúdo do arquivo: %v", err)
	}

	// Converte o conteúdo para uma string
	conteudoArquivo := string(conteudo)

	linhas := strings.Split(conteudoArquivo, "\n")

	if numeroLinha > 0 && numeroLinha < len(linhas) {
		linhas[numeroLinha-1] = novoValor
	}

	novoConteudoArquivo := strings.Join(linhas, "\n")

	_, err = file.Seek(0, 0)
	if err != nil {
		log.Fatalf("Erro ao mover o ponteiro: %v", err)

	}

	_, err = file.WriteString(novoConteudoArquivo)
	if err != nil {
		log.Fatalf("Erro ao salvar novo conteudo: %v", err)
	}

	err = file.Truncate(int64(len(novoConteudoArquivo)))
	if err != nil {
		log.Fatalf("Erro ao truncar: %v", err)
	}
	// // Criar um scanner para ler o conteúdo do arquivo linha por linha
	// fmt.Println(novoConteudoArquivo)

	return nil
}

func contarLinhasNoConteudo(data []byte) int {
	// Inicializar o contador de linhas
	numLinhas := 0

	// Criar um scanner para contar as linhas no conteúdo
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		numLinhas++
	}

	// Verificar por erros durante o scanner
	if err := scanner.Err(); err != nil {
		log.Fatal("Erro ao contar as linhas no conteúdo:", err)
	}

	return numLinhas
}

func AdicionarConfiguracao(destFile string) error {
	// Abrir o arquivo em modo de escrita, criando-o se não existir
	file, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	conteudo, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Erro ao ler o conteúdo do arquivo2: %v", err)
	}

	valor := contarLinhasNoConteudo(conteudo)

	// Conteúdo a ser adicionado
	novoConteudo := `
[Account1]
label=7801
server=jrsolution.gvctelecom.com.br:5131
proxy=jrsolution.gvctelecom.com.br:5131
domain=jrsolution.gvctelecom.com.br:5131
username=7801
password=4d76270cf0a782e201b1a2de17c6fe7ca32f49e783a964a2
authID=7801
displayName=
dialingPrefix=
dialPlan=
hideCID=0
voicemailNumber=
transport=
publicAddr=
SRTP=
registerRefresh=300
keepAlive=15
publish=0
ICE=0
allowRewrite=0
disableSessionTimer=0
`

	if valor != 105 {
		// Escrever o novo conteúdo no arquivo
		_, err = file.WriteString(novoConteudo)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Não vai precisar")
	}
	return nil
}
