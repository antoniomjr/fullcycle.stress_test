# Usar a imagem base do Go
FROM golang:1.18

# Definir o diretório de trabalho
WORKDIR /app

# Copiar o arquivo main.go para o diretório de trabalho
COPY main.go .

# Compilar o código Go
RUN go build -o loadtester main.go

# Definir o comando de entrada
ENTRYPOINT ["./loadtester"]
