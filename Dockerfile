# Imagem base com Go pré-instalado
FROM golang:1.19

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Baixar o código do repositório
RUN git clone https://github.com/scuba13/AmacoonServices.git .

# Copiar arquivos do diretório atual para o diretório de trabalho do container
COPY . .

# Compilar o projeto Go dentro do container
RUN go mod tidy
RUN go build -o main cmd/server/main.go

# Definir a porta que a aplicação irá utilizar
EXPOSE 8080

# Executar a aplicação compilada
CMD [ "./main" ]
