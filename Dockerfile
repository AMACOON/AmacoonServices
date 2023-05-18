# Imagem base com Go pré-instalado
FROM golang:1.20 AS builder

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar arquivos do diretório atual para o diretório de trabalho do container
COPY . .

# Compilar o projeto Go dentro do container
RUN go mod tidy
RUN go build -o main cmd/server/main.go

# Usar uma imagem menor para a etapa final
FROM gcr.io/distroless/base

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar o binário compilado e quaisquer outros arquivos necessários
COPY --from=builder /app/main /app/main
COPY --from=builder /app/config.local.yaml /app/config.local.yaml
COPY --from=builder /app/config.prod.yaml /app/config.prod.yaml

# Definir a porta que a aplicação irá utilizar
EXPOSE 8080

# Executar a aplicação compilada
CMD ["/app/main"]

