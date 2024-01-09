### Fase de criação (:

Se tiver o Go instalado

Só entrar fazer o pull e entar na pasta list_ramais

#### Dentro da pasta só execurar o comando

`{ go run main/main.go } `

PS: Está rodando na porta :8080 local, verificar se ela está Up, mas pode trocar a porta no arquivo main/main.go

log.Fatal(http.ListenAndServe(":8080", nil))

Só colocar a porta que tiver Up.

