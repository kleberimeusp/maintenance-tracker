# API de contabilização das tarefas, manutenção realizadas durante um dia de trabalho.

(MODELO ORGANIZACIONAL DE PROJETO)

(APLICADO EM TODOS OS CLIENTE, MODELO DE README ONDE TRABALHEI EM CLIENTES INTERNACIONAIS) 

(TRADUZIDO PARA ESSE TESTE, AQUI EM INGLÊS)

Sistema com finalidade dO técnico realiza tarefas e só pode ver, criar ou atualizar suas próprias
tarefas realizadas.
O gerente pode ver tarefas de todos os técnicos, excluí-las e deve ser
notificado quando algum técnico realiza uma tarefa..

## Principais Tecnologias Usadas - API (Microsserviços Backend)

![Git](https://img.shields.io/badge/Git-2.40.0-%237159c1?style=for-the-badge&logo=git)
![Insomnia](https://img.shields.io/badge/insonmia-2023.1.0-%237159c1?style=for-the-badge&logo=insomnia)
![GoLang](https://img.shields.io/badge/GoLang-1.22.5-%237159c1?style=for-the-badge&logo=go)
![Maven](https://img.shields.io/badge/Maven-3.8-%237159c1?style=for-the-badge&logo=Apache%20Maven)
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
![Maven](https://img.shields.io/badge/Maven-3.8-%237159c1?style=for-the-badge&logo=Apache%20Maven)
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
$ git clone https://github.com/swordhealth-interviewe/maintenance-tracker.git
```

## Upload database MySQL em Docker

Acesse o container do banco de dados e importe o banco de dados.

```shell
docker exec -it db bash
mysql -u root -proot
exit;
mysql -uroot -proot maintenance-tracker < /db/maintenance-tracker.sql
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

## Arquitetura de Referência (Empresa maintenance-tracker)

- Link do material [clique aqui](#)
