# API de contabilização das tarefas, manutenção realizadas durante um dia de trabalho.

(MODELO ORGANIZACIONAL DE PROJETO)

(APLICADO EM TODOS OS CLIENTE, MODELO DE README ONDE TRABALHEI EM CLIENTES INTERNACIONAIS) 

(TRADUZIDO PARA ESSE TESTE, AQUI EM PORTUGUÊS)

Sistema com finalidade do profissional técnico realizar tarefas e só pode ver, criar ou atualizar suas próprias tarefas realizadas.
O gerente pode ver tarefas de todos os técnicos, excluí-las e deve ser notificado quando algum técnico realiza uma tarefa..

## Principais Tecnologias Usadas - API (Microsserviços Backend)

![Git](https://img.shields.io/badge/Git-2.40.0-%237159c1?style=for-the-badge&logo=git)
![Insomnia](https://img.shields.io/badge/insonmia-2023.1.0-%237159c1?style=for-the-badge&logo=insomnia)
![GoLang](https://img.shields.io/badge/GoLang-1.22.5-%237159c1?style=for-the-badge&logo=go)
![Notepad++](https://img.shields.io/badge/notepad++-8.0-%237159c1?style=for-the-badge&logo=notepadplusplus)
![Postman](https://img.shields.io/badge/postman-10.13.10-%237159c1?style=for-the-badge&logo=postman)
![VSCode](https://img.shields.io/badge/vscode-1.77.3-%237159c1?style=for-the-badge&logo=visualstudiocode)
![Dbeaver](https://img.shields.io/badge/Dbeaver-22.2.3-%237159c1?style=for-the-badge&logo=Dbeaver)
![Docker](https://img.shields.io/badge/Docker-4.1.8-%237159c1?style=for-the-badge&logo=docker)
![Test Runner](https://img.shields.io/badge/Test%20Runner-1.1.0-%237159c1?style=for-the-badge&logo=tesrunner)

---

## Pré-requisitos Ambiente 

![Git](https://img.shields.io/badge/Git-2.40.0-%237159c1?style=for-the-badge&logo=git)
![Insomnia](https://img.shields.io/badge/insonmia-2023.1.0-%237159c1?style=for-the-badge&logo=insomnia)
![GoLang](https://img.shields.io/badge/GoLang-1.22.5-%237159c1?style=for-the-badge&logo=go)
![Notepad++](https://img.shields.io/badge/notepad++-8.0-%237159c1?style=for-the-badge&logo=notepadplusplus)
![Postman](https://img.shields.io/badge/postman-10.13.10-%237159c1?style=for-the-badge&logo=postman)
![VSCode](https://img.shields.io/badge/vscode-1.77.3-%237159c1?style=for-the-badge&logo=visualstudiocode)
![Dbeaver](https://img.shields.io/badge/Dbeaver-22.2.3-%237159c1?style=for-the-badge&logo=Dbeaver)
![Docker](https://img.shields.io/badge/Docker-4.1.8-%237159c1?style=for-the-badge&logo=docker)
![Test Runner](https://img.shields.io/badge/Test%20Runner-1.1.0-%237159c1?style=for-the-badge&logo=tesrunner)

---

## Pré-requisitos - VDI

- [Dbeaver](#)
- [Docker](#)
- [VSCode](#)
- [GoLang](#)
- [Notepad++](https://github.com/notepad-plus-plus/notepad-plus-plus/releases/download/v8.5.2/npp.8.5.2.Installer.x64.exe)
- [Test Runner](#)

## Configuração de Ambiente 

- Baixar o código da aplicação no diretório de trabalho de escolha própria

```sh
$ git clone https://github.com/swordhealth-interviewer/maintenance-tracker2.git
```

## Upload database MySQL em Docker

Atualização das packages:

```shell
$ go mod tidy   
```

Execute o comando abaixo para instalar o golang-migrate:

```shell
$ go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Aplicar as Migrações:

```shell
$ migrate -path migrations -database "mysql://root:root@tcp(localhost:3306)/maintenance_tasks" up
```

No PowerShell, adicione as dependências necessárias:

```shell
$ go get github.com/go-sql-driver/mysql
$ go get github.com/gorilla/mux
$ go get github.com/jmoiron/sqlx
$ go get github.com/rs/zerolog
$ go get github.com/go-playground/validator/v10
$ go get github.com/go-playground/universal-translator
$ go get github.com/golang-jwt/jwt
$ go get github.com/xhit/go-simple-mail/v2
```

Instale a ferramenta swag:

```shell
$ go get -u github.com/swaggo/swag/cmd/swag
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
$ go install github.com/swaggo/swag/cmd/swag@latest
```

Por padrão, o GOPATH está definido em $HOME/go, mas você pode verificá-lo executando:

```shell
$ go env GOPATH
```

No PowerShell, adicione GOPATH/bin ao PATH:

```shell
$env:GOPATH = (go env GOPATH)
$env:PATH += ";$env:GOPATH\bin"
```

Verificar versão Swagger.

```shell
$ swag --version
```

Gerar Swagger.

```shell
$ swag init -g cmd/main.go
```

Acesse o container do banco de dados e importe o banco de dados.

```shell
$ docker-compose up --build
```

Deploy aplicação local.

```shell
$ go run cmd/main.go
```

URL Swagger.

```shell
http://localhost:8080/swagger-ui/index.html
```

## Modo de acesso ao sistema/módulo (urls)
O acesso ao sistema pode ser efetuado por meio dos links abaixo:
- Localhost: [Swagger máquina local](http://localhost:8080/swagger-ui/index.html)
  ![img.png](img.png)
- Desenvolvimento: [XXX](x)
- Homologação: [XXX](x)
- QA: [XXX](x)
- Produção: [XXX](x)

---

## Arquitetura de Referência (Empresa maintenance-tracker2)

- Link do material [clique aqui](#)