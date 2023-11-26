FROM golang:1.20-alpine as builder
ARG CONFIG_SOURCE
ARG CONFIG_PATH
WORKDIR /app
COPY ${base_dir}/ /app
ENV CONFIG_SOURCE=${CONFIG_SOURCE}
ENV CONFIG_PATH=${CONFIG_PATH}
RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest
CMD air -c .air.toml
