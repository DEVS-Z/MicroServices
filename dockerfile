# Etapa de construcción
FROM golang:1.25rc2-bookworm AS builder

# Argumento para arquitectura (por defecto: arm64)
# ARG TARGETARCH=arm64
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=${TARGETARCH}

WORKDIR /app

# Copiar solo go.mod y go.sum primero para mejor cache
COPY go.mod go.sum ./
RUN go mod download

# Ahora el resto del código
COPY . .

# Compilar el binario
RUN go build -ldflags="-s -w" -o server ./cmd/rest/main.go

# Etapa de ejecución (imagen mínima)
FROM debian:bookworm-slim

# Instalar sólo lo necesario
RUN apt-get update && \
    apt-get install -y --no-install-recommends postgresql-client && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Copiar el binario desde la etapa anterior
COPY --from=builder /app/server .

# Exponer el puerto
EXPOSE 8080

# Comando por defecto
CMD ["./server"]