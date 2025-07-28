FROM golang:1.24-alpine
WORKDIR /app

# Copy the proto module first to the correct relative path
COPY ../proto/ ../proto/
COPY ../.env .

# Copy ingestor files
COPY api_gateway/go.mod api_gateway/go.sum ./
RUN go mod download

COPY api_gateway/ ./
RUN go build -o api_gateway

CMD ["./api_gateway"]
