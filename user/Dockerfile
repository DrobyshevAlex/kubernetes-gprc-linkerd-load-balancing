FROM golang:alpine as builder

WORKDIR /app

ADD . .

RUN go mod download

RUN go build -o bin/app ./cmd/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/bin /app

CMD [ "./app" ]
