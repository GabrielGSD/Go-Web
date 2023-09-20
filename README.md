# Go-Web
Este é um projeto de exemplo em Go que demonstra como criar um aplicativo da web simples usando o pacote `net/http`, o roteador `gorilla/mux` e o acesso a um banco de dados MySQL com o pacote `github.com/go-sql-driver/mysql`.

## Pré-requisitos

- **Go instalado**: Certifique-se de que você tenha o Go instalado em seu sistema. Você pode baixá-lo em [https://golang.org/dl/](https://golang.org/dl/).

- **MySQL**: Você deve ter um servidor MySQL em execução e ter criado um banco de dados chamado "go_course". O aplicativo assume que você tem um usuário "root" com senha "root" configurado para acessar o banco de dados. Certifique-se de ajustar as configurações de conexão no código se necessário.

## Configuração

1. Clone o repositório para o seu sistema local:

   ```bash
   git clone https://github.com/seu-usuario/go-web-app.git
   cd go-web-app

2. Baixe as dependências do Go usando o comando go get:

   ```bash
   go get

3. Execute o aplicativo:

   ```bash
   go run main.go
   
O aplicativo será executado na porta 8080.

## Uso

- Acesse o aplicativo em [http://localhost:8080](http://localhost:8080).

- A página inicial lista todos os posts disponíveis no banco de dados.

- A outra rota permite a busca de um post específico com base em seu ID. Para acessar a página de visualização de um post, basta navegar até http://localhost:8080/{id}/view, onde {id} é o número de identificação único do post desejado.

