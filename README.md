# Desafio Clean Architecture

Olá devs!
Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:
- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

Configurações para execução:
- DB_DRIVER=mysql
- DB_HOST=localhost
- DB_PORT=3306
- DB_USER=root
- DB_PASSWORD=root
- DB_NAME=orders
- WEB_SERVER_PORT=:8000
- GRPC_SERVER_PORT=50051
- GRAPHQL_SERVER_PORT=8080
---
## Como executar
Após fazer o clone do repositório através do git, execute os seguintes passos:
1. Abra o projeto no Visual Studio Code - VSCode

2. Abra um terminal no VSCode, e acesse o diretório raiz do projeto: /clean-arch e execute o comando ```docker compose up -d```.<br>
Este comando deixará a infra estrtura pronta para uso.

3. Abra outro terminal no VSCode, acesse o diretório /clean-arch/cmd/ordersystem e execute o comando ```go run main.go wire_gen.go```<br>
Este comando deixará os serviços em execução nas seguintes portas:
    * Web Server - porta:8000
    * gRPC Server - porta:50051
    * GraphQL Server - porta:8080

Como requisito é necessário ter o git e o docker instalado, ambos pronto para serem executados no computador, assim como demonstrado durante o curso.

## Como testar

### Teste Web Server
- Dentro do VSCode, acesse o Explorer e navegue até o diretório /clean-arch/api
    * para inserir registro: clique no arquivo create_order.http e em seguida clique em Send Request
    * para listar registros: clique no arquivo list_orders.http e em seguida clique em Send Request

Como requisito é necessário ter a extensão REST Client instalada no VSCode, a mesma informada durante o curso.

### Teste gRPC
- Abra um novo terminal no VSCode, acesse o diretório raiz do projeto, /clean-arch, e digite o comando ```evans -r repl``` para abrir o gRPC Client ensinado durante o curso.<br>
No mesmo terminal, digite os comandos e pressione \<ENTER>:
    * ```package pb```
    * ```service OrderService```

- Para criar registro: digite o comando, ```call CreateOrder```, pressione \<ENTER>,  e informe o dado a cada interação com o client gRPC
- Para listar registros: digite o comand, ```call ListOrders```, pressione \<ENTER>

Como requisito é necessário ter o gRPC Client, Evans, ensinado durante o curso instalado.

### Teste GraphQL
- Abra um web browser, e digite a URL: http://localhost:8080/. Cole o conteúdo abaixo dentro do playground que será aberto:
```
mutation creteOrder{
  createOrder(input:{id: "a", Price: 1000.0, Tax: 1.5}) {
    id
    Price
    Tax
    FinalPrice
  } 
}

query orders{
  orders {
		id
    Price
    Tax
    FinalPrice
  }
}
```
- Dentro do playground haverá um botão para rodar a instrução acima, e ao clicar neste botão, escolha:
    * para criar registro: createOrder
    * para listar registros: orders

# Nota
**Ao realizar o teste de inserção de registro, o campo *id* não poderá conter valores duplicados, visto que resultará em falha, pois, é chave primária.**
> **Todos os requisitios listados bem como o modo de operação seguiram o que foi ministrado no curso de pós gradução da FullCycle, Pós GO Expert.**