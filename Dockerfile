# FROM golang:1.25.1 AS builder

# # Define o diretório de trabalho para a sua aplicação
# WORKDIR /app
# COPY . .
# # Baixa as dependências e gera o go.sum
# RUN go mod tidy
# # Constrói o executável da sua aplicação Go
# RUN go build -o main .


# # O segundo estágio (imagem final, mais leve)
# FROM debian:bookworm-slim
# WORKDIR /app
# COPY --from=builder /app/main .
# # Define o comando para rodar a aplicação
# CMD ["./main"]


FROM golang:1.25.1 AS builder

# Se sua imagem for golang:1.21 (sem o sufixo -alpine)
RUN apt-get update && apt-get install -y ca-certificates git && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY . .
# comando para instalar as dependencias do projeto
RUN go mod tidy 
RUN go build -o main .

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]