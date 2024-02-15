# Use a imagem base do golang
FROM golang:latest AS build

# Define a pasta de trabalho do container
WORKDIR /go/src/api-rinha

# Copie os arquivos do projeto para a pasta de trabalho
COPY . .

# Baixa as dependencias
RUN go mod download
# Compile o código
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/main cmd/main.go

# Use uma imagem base minimalista, como alpine, para a imagem final
FROM alpine:latest

# Copie o binário compilado para a pasta de trabalho do container
COPY --from=build /go/bin/main /

# Defina qual é o comando para iniciar o container
CMD ["/main"]