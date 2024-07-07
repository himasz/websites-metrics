FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/config/json /root/config/json

RUN apk --no-cache add ca-certificates

EXPOSE 8080

CMD ["./main"]
