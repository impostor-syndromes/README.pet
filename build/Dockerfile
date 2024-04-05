FROM golang:1.22.2-alpine3.19

WORKDIR /app

COPY . .

RUN apk update && \
    go mod tidy

CMD ["go", "run", "cmd/README.pet/main.go"]