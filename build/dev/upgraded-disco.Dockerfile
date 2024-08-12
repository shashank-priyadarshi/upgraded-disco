FROM golang:1.23rc2-alpine as builder
ARG CONFIG_SOURCE
ARG CONFIG_PATH
WORKDIR /app
COPY ${base_dir}/ /app
ENV CONFIG_SOURCE=${CONFIG_SOURCE}
ENV CONFIG_PATH=${CONFIG_PATH}
RUN go install github.com/cosmtrek/air@latest
CMD air -c .air.toml
