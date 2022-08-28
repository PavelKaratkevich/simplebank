# build stage
FROM golang:1.19.0-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/app.env .
COPY --from=builder /app/start.sh .
COPY --from=builder /app/db/migration  ./db/migration
RUN chmod +x ./start.sh

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]