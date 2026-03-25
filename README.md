# Categories Microservice

API REST para gerenciamento de categorias, construída em Go com arquitetura limpa, PostgreSQL e documentação Swagger interativa.

---

## O que a aplicação faz

- Cria, lista, busca, atualiza e deleta categorias
- Cada categoria possui **nome**, **descrição** e **teste**
- Valida que o nome tenha mais de 5 caracteres
- Impede nomes duplicados
- Gera documentação interativa via Swagger UI em `/swagger/index.html`

---

## Tecnologias

| Pacote | Função |
|---|---|
| [gin-gonic](https://github.com/gin-gonic/gin) | Web server / roteamento HTTP |
| [gorm](https://gorm.io) | ORM para PostgreSQL |
| [swaggo/swag](https://github.com/swaggo/swag) | Geração de documentação Swagger |
| [godotenv](https://github.com/joho/godotenv) | Carregamento de variáveis de ambiente |

---

## Pré-requisitos

- [Docker](https://www.docker.com/) instalado
- [Docker Compose](https://docs.docker.com/compose/) instalado

---

## Como rodar localmente

**1. Clone o repositório**
```bash
git clone https://github.com/Gabstuxo/BancodeDadosEmGO.git
cd BancodeDadosEmGO
```

**2. Crie o arquivo de ambiente**
```bash
cp .env.example .env
```

**3. Suba os containers**
```bash
docker-compose up -d --build
```

Isso inicia dois containers:
- `categories-db` — PostgreSQL na porta `5432`
- `categories-api` — API Go na porta `8080`

**4. Inicie a API dentro do container**
```bash
docker exec categories-api sh -c "echo 'ENVIRONMENT=local' > /app/.env"
docker exec -d categories-api sh -c "cd /app && go run cmd/api/main.go"
```

Aguarde ~10 segundos e verifique:
```bash
curl http://localhost:8080/healthy
# {"success":true}
```

**5. Acesse o Swagger**

Abra no navegador: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## Endpoints

| Método | Rota | Descrição |
|---|---|---|
| GET | `/healthy` | Health check |
| GET | `/categories` | Lista todas as categorias |
| GET | `/categories/:id` | Busca categoria por ID |
| POST | `/categories` | Cria uma nova categoria |
| PATCH | `/categories/:id` | Atualiza uma categoria |
| DELETE | `/categories/:id` | Deleta uma categoria |

### Exemplo — criar categoria
```bash
curl -X POST http://localhost:8080/categories \
  -H "Content-Type: application/json" \
  -d '{"name": "Tecnologia", "description": "Tudo sobre tech", "teste": "valor"}'
```

---

## Rodar os testes

```bash
docker exec categories-api sh -c "cd /app && go test ./..."

# Com detalhes
docker exec categories-api sh -c "cd /app && go test -v ./..."

# Com cobertura
docker exec categories-api sh -c "cd /app && go test -cover ./..."
```

---

## Arquitetura

```
cmd/api/
  main.go                  # bootstrap: .env, DB, rotas
  routes/categories/       # handlers Gin (entrada e saída HTTP)

internal/categories/
  use-cases/               # regras de negócio
  repository/              # acesso ao banco via GORM
  models/                  # struct Category + validações

pkg/
  error/                   # coletor de erros de validação
  utils/                   # utilitários (ex: StringToUint)

docs/                      # gerado pelo swag init (Swagger)
```

---

## Troubleshooting

**Verificar conexão com o banco:**
```bash
docker exec -it categories-db bash
su postgres
psql "dbname=meetup host=localhost user=postgres password=password123 port=5432"
\l    # lista os bancos
exit
exit
exit
```

**Regenerar documentação Swagger após mudanças:**
```bash
docker exec categories-api sh -c "cd /app && swag init -g cmd/api/main.go"
```
