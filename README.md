# Projeto Filmes Web em Go

Este projeto exibe os filmes da API `localhost:7111/movies` com visual moderno.

## Passo a passo

### 1. Instalar o Go
- Baixe e instale em https://go.dev/dl/
- Adicione o Go ao PATH se necessário.
- Verifique com:
  ```
  go version
  ```

### 2. Instalar dependências
No terminal, na pasta do projeto:
```
go mod init webgo
go get github.com/joho/godotenv
```

### 3. Configurar variáveis de ambiente
Edite o arquivo `.env` conforme exemplo:
```
API_URL=http://localhost:7111
APP_PORT=8113
```

### 4. Executar o projeto
No terminal:
```
go run main.go
```
Acesse: http://localhost:8113

---

## Estrutura
- `main.go`: Código principal
- `templates/`: HTML
- `static/`: CSS
- `.env`: Variáveis de ambiente
- `.gitignore`: Arquivos ignorados

---

## Dicas
- Para instalar novos pacotes: `go get <pacote>`
- Para build: `go build`

## Tela
Essa é a aparência do app, convido você a criar a página de detalhes do filme.
<img src="https://personalizetudo.com.br/assets/images/frontend.png" alt="drawing" width="600"/>
