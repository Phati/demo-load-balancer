FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main
FROM alpine:latest AS production
COPY --from=builder /app/migrations/ /app/migrations/
COPY --from=builder /app/main /app/main
EXPOSE 8080
CMD [ "/app/main" ]
