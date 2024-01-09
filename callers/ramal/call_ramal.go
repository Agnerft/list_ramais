package ramal

import (
	"fmt"
	"log"
	"net/http"
	"os/user"
	"path/filepath"

	"github.com/makesystem/list_ramais/controller/ramal"
	"github.com/makesystem/list_ramais/execute"
	"github.com/makesystem/list_ramais/util"
)

// HandleRamais é o manipulador para a rota /ramais
func HandleRamais(w http.ResponseWriter, r *http.Request) {

	// Fazer a requisição para obter os dados JSON
	ramais, err := ramal.RequestJsonRamal("link" + "/status_central")
	if err != nil {
		http.Error(w, "Erro ao obter os ramais", http.StatusInternalServerError)
		return
	}

	fmt.Println(ramais)

	// util.RenderTemplate(w, ramais, "main/assets/ramais.html")

}

func HandleSelecionarRamal(w http.ResponseWriter, r *http.Request) {
	// Obter o SIP da query string
	sipSelecionado := r.URL.Query().Get("sip")

	// Implementar a lógica para lidar com o SIP selecionado (por exemplo, armazenar em uma variável global)
	fmt.Println("SIP Selecionado:", sipSelecionado)

	// Responder ao cliente
	w.WriteHeader(http.StatusOK)

	// Obter o diretório do usuário
	usr, err := user.Current()
	if err != nil {
		log.Fatal("Erro ao obter o diretório do usuário:", err)
	}

	url := "https://www.microsip.org/download/MicroSIP-3.21.3.exe"

	destDown := filepath.Join(usr.HomeDir, "AppData", "Local", "MicroSIP", "MicroSIP-3.21.3.exe")

	destFile := filepath.Join(usr.HomeDir, "AppData", "Roaming", "MicroSIP", "microsip.ini")

	err = util.AdicionarConfiguracao(destFile)
	if err != nil {
		log.Fatal("Erro ao inserir a Parte.", err)
	}

	err = execute.DownloadMicrosip(url, destDown)
	if err != nil {
		log.Fatal("Erro ao baixar o Arquivo.", err)
	}

	err = util.ReadFile(destFile, "accountId=1", 2)
	if err != nil {
		log.Fatal("Erro para modificar o AccountId.", err)
	}

	// err = util.ReadFile(destFile, "videoBitrate=256", 24)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o videoBitrate.", err)
	// }

	// err = util.ReadFile(destFile, filepath.Join(usr.HomeDir, "Desktop"), 33)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o recordingPath.", err)
	// }

	// err = util.ReadFile(destFile, "recordingFormat=mp3", 34)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o recordingFormat.", err)
	// }

	// err = util.ReadFile(destFile, "autoAnswer=all", 38)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o autoAnswer.", err)
	// }

	// err = util.ReadFile(destFile, "denyIncoming=", 44)
	// if err != nil {
	// 	log.Fatal("Erro para modificar o denyIncoming.", err)
	// }

	util.RenderTemplate(w, sipSelecionado, "main/assets/ramalSelecionado.html")
}
