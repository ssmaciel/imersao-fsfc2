# Definindo imagem de base
FROM golang:1.16

# Definindo diretório de trabalho
WORKDIR /go/src/imersao-fsfc2

# Copiando diretório para a imagem
# Como é o WORKDIR poderia também ser: ADD . .
ADD . .

# Executando comando de build
RUN GOOS=linux go build

EXPOSE 8000

# Definindo comando de inicialização
# Ececuta binário
CMD ./desafio01
