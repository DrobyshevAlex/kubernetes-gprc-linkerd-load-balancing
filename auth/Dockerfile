FROM --platform=linux/amd64 golang:alpine as builder

RUN go install github.com/githubnemo/CompileDaemon@latest

WORKDIR /app

ADD auth .

RUN go mod download

RUN go build -o bin/app ./cmd/main.go

FROM --platform=linux/amd64 alpine:latest
WORKDIR /app

COPY --from=builder /app/bin /app

CMD [ "/app/app" ]
