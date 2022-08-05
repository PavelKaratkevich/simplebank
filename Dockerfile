# build stage
FROM golang:1.19.0-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-386.tar.gz | tar xvz 

# run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder app/main .
COPY --from=builder /app/migrate ./migrate 
COPY --from=builder app/app.env .
COPY --from=builder /app/start.sh .
COPY --from=builder /app/db/migration  ./migration

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]