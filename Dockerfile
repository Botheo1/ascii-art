FROM golang:1.26.3 AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp .

FROM alpine:latest
LABEL version="1.0.0"
WORKDIR /app
COPY --from=builder /app/myapp .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/style ./style
EXPOSE 8080
CMD ["./myapp"]