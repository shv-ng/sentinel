FROM golang:1.24-alpine
WORKDIR /app

# Copy the proto module first to the correct relative path
COPY ../proto/ ../proto/
COPY ../.env .

# Copy ingestor files
COPY ingestor/go.mod ingestor/go.sum ./
RUN go mod download

COPY ingestor/ ./
RUN go build -o ingestor

CMD ["./ingestor"]
