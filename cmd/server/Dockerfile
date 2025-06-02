FROM golang:1.24-alpine

RUN apk update && apk add --no-cache \
    postgresql-client \
    bash
    
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server/main.go

EXPOSE 50051

CMD ["./server"]