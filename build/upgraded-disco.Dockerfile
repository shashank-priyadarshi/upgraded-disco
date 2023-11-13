FROM golang:1.20-alpine as builder
COPY . .
WORKDIR .
RUN go build -o app ./cmd/main.go

FROM alpine
COPY --from=builder /app .
COPY --from=builder /config.yaml .
COPY --from=builder /plugins .
RUN ./app
