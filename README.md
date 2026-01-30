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
