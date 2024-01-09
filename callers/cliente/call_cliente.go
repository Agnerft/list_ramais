package cliente

import (
	"fmt"
	"log"
	"net/http"

	"github.com/makesystem/list_ramais/controller/cliente"
	"github.com/makesystem/list_ramais/util"
)

var url_padrao = "https://root:agner102030@basesip.makesystem.com.br/clientes?documento="

func HandleClient(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Renderizar o formulário HTML
		util.RenderTemplate(w, nil, "main/assets/clientes.html")
		return

	}

	if r.Method == http.MethodPost {
		// Obter a informação do formulário
		cnpj := r.FormValue("cnpj")

		// fmt.Println(url_padrao + cnpj)

		cliente, err := cliente.RequestJsonCliente(url_padrao + cnpj)
		if err != nil {
			log.Fatal("Erro ao buscar os ramais.", err)
		}

		// for _, itensCliente := range cliente {
		// 	fmt.Println(itensCliente.Link)
		// }

		fmt.Println(cliente)
		fmt.Println("ta aqui")

		util.RenderTemplate(w, cliente, "main/assets/clientes.html")

		// Faça algo com a informação (por exemplo, imprimir no console)
		// fmt.Println("Informação recebida:", cnpj)

		// fmt.Println("Ta aqui?")

		return

	}

	// Método não suportado
	http.Error(w, "Método HTTP não suportado", http.StatusMethodNotAllowed)
}
