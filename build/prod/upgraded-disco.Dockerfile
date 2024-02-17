FROM golang:1.20-alpine as builder
WORKDIR /app
COPY ${base_dir}/ /app
RUN go build -o app ./cmd/main.go

FROM alpine
ARG CONFIG_SOURCE
ARG CONFIG_PATH
COPY --from=builder /app/app .
COPY --from=builder /app/${CONFIG_PATH}/${CONFIG_SOURCE} .
COPY --from=builder /app/plugins .
ENV CONFIG_SOURCE=${CONFIG_SOURCE}
ENV CONFIG_PATH=${CONFIG_PATH}
RUN ./app
