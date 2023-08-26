FROM golang:1.20 as builder
# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
# Copy app files
COPY . .
# Build app
RUN go build -o app ./

FROM alpine:3.15 as production
## Add certificates
RUN apk update
RUN apk add --no-cache ca-certificates
RUN apk add git
## Copy built binary from builder
COPY --from=builder app .
## Exec built binary
CMD ./app
