FROM golang:1.22.0-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN GOOS=linux go build -ldflags '-w -s' -a -o /app/application cmd/main.go

FROM alpine
WORKDIR /app
RUN apk --update --no-cache add libc6-compat
COPY --from=builder /app/application /app/app
COPY .env /app/.env

CMD ["/app/app",".env"]
