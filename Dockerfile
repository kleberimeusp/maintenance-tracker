# Etapa de build
FROM golang:1.20-alpine AS builder

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o go.mod e o go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixar todas as dependências
RUN go mod download

# Copiar o restante do código-fonte do projeto
COPY . .

# Construir o binário
RUN go build -o main ./cmd/main.go

# Etapa de produção
FROM alpine:latest

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /root/

# Copiar o binário construído da etapa de build
COPY --from=builder /app/main .

# Copiar a documentação Swagger gerada
COPY --from=builder /app/docs ./docs

# Expor a porta que o aplicativo vai rodar
EXPOSE 8080

# Comando para executar o aplicativo
CMD ["./main"]
