FROM golang:1.22.2-alpine3.19

WORKDIR /app

COPY api .

RUN ls

RUN apk update && apk --no-cache add git && \
    go mod tidy && \
    go install github.com/cosmtrek/air@v1.49.0

CMD ["air", "-c", ".air.toml"]