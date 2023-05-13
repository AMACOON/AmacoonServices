# # Imagem base com Go pré-instalado
# FROM golang:1.20

# # Definir o diretório de trabalho dentro do container
# WORKDIR /app

# # Copiar arquivos do diretório atual para o diretório de trabalho do container
# COPY . .

# # Compilar o projeto Go dentro do container
# RUN go mod tidy
# RUN go build -o main cmd/server/main.go

# # Definir a porta que a aplicação irá utilizar
# EXPOSE 8080

# # Executar a aplicação compilada
# CMD [ "./main" ]



# FROM golang:1.20 AS builder
# WORKDIR /app
# COPY . .
# RUN go mod tidy
# RUN go build -o main cmd/server/main.go
# FROM gcr.io/distroless/base-debian10
# WORKDIR /
# COPY --from=builder /app/main .
# EXPOSE 8080
# ENTRYPOINT ["/main"]


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
COPY --from=builder /app/config.yaml /app/config.yaml

# Definir a porta que a aplicação irá utilizar
EXPOSE 8080

# Executar a aplicação compilada
CMD ["/app/main"]

