FROM golang:1.22.6 AS builder

WORKDIR /app

COPY . .
RUN go mod download

COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o myapp ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/.env .

EXPOSE 6666

CMD [ "./myapp" ]