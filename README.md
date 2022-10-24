## Star Wars API to Game

Este projeto tem o objetivo de fornecer uma API Rest com operações de consulta, inserção e carregamento para os planetas da saga Star Wars.

## Começando

Essas instruções fornecerão uma cópia do projeto que pode ser executada em sua máquina local para fins de desenvolvimento e teste. Consulte o item de implantação para obter notas sobre como implantar o projeto em um sistema ativo.

## Pré-requisitos

Este pacote foi criado com go 1.19 e tudo que você precisa é a biblioteca go padrão.
O projeto ate o momento possui somente 3 dependências

- mock - para geração dos mocks nos testes unitários
- pq - driver de conexão para o banco postgres
- sqlx - para manipulação das queres na camada `repository`

## Variáveis

Para facilitar a execução da aplicação, o projeto possui um arquivo `.env` para declarar as variáveis de ambiente

`HOST` - Endereço do host que a aplicação usará.<br/>
`PORT` - Número da porta onde a aplicação escutará as requisições http.<br/>
`DEBUG` - Variável para definir se uma API produzirá logs para stdout.<br/>
`STORAGE` - Endereço do banco de dados.<br/>

## Observação

A maior parte das operações descritas nessa documentação como, instalação, teste e etc, possuem tarefas definidas no arquivo `Makefile`. Facilitando no processo de execução e mantenabilidade do projeto.

## Instalação

Isto é o que você precisa para instalar o aplicativo a partir do código-fonte:

```bash
    git clone https://github.com/paraizofelipe/star-planet.git 
```

Para construir a versão do docker, você pode utilizar o CLI `docker-compose`, com o comando:

```bash
	make dk-deploy
```

## Executando os testes

Até eu terminar este README não há tantos testes de unidade escritos.

Você pode executar testes assim:

```bash
	make test
```

Para rodar a API local em sua workstation, você pode executar o comando:

```bash
    STORAGE=postgres://star:planet@localhost:5432/star-planet?sslmode=disable DEBUG=true HOST=0.0.0.0 PORT=3000 make run
```

Lembrando que todas a variáveis de ambiente podem ser consultadas no arquivo `.env` e até mesmo utilizada em conjunto com o [direnv](https://direnv.net/)

## API

### Load

Carrega um planeta não existente na base de dados, consultando outra API [swapi.dev](https://swapi.dev/). Caso o planeta já exista, o processo de load sera ignorado.

```bash
curl -i -X POST "http://localhost:300/api/planets/1"
```

### FindByBy

Busca um planeta na base de dados pelo seu ID.

```bash
curl -i -X GET "http://localhost:300/api/planets/id/1"
```

### FindByName

Busca um planeta na base de dados pelo seu nome.

```bash
curl -i -X GET "http://localhost:300/api/planets/name/Tatooine"
```

### Remove

Remove um planet da base de dados pelo seu ID.

```bash
curl -i -X DELETE "http://localhost:300/api/planets/id/1"
```

### List

Lista todos os planetas existentes na base de dados.

```bash
curl -i -X GET "http://localhost:300/api/planets/"
```
