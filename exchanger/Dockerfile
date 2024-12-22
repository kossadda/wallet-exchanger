FROM golang:1.23.3-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download && go build -o main cmd/main.go

FROM alpine:3.18

COPY --from=builder /app/main /main
COPY --from=builder /app/wait-for-it.sh /wait-for-it.sh
COPY --from=builder /app/config /config

RUN chmod +x /wait-for-it.sh

EXPOSE 44044

CMD ["/wait-for-it.sh", "db", "5432", "--", "/main"]
