FROM golang:1.18 as builder
# Define build env
ENV GOOS linux
ENV CGO_ENABLED 0
# Add a work directory
WORKDIR /app
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Build app
RUN go build -o app ./

FROM alpine:3.15 as production
ENV SQL_URI=""
ENV MONGO_URI=""
ENV DB_NAME=""
ENV SERVER_PORT=""
ENV TODO_API_PORT=""
ENV GH_INTEGRATION_ORIGIN=""
ENV GITHUB_TOKEN=""
ENV GITHUB_USERNAME=""
ENV ALLOWED_ORIGIN=""
ENV SECRET_KEY=""
ENV BIO=""
ENV GITHUB=""
ENV TODOS=""
ENV GRAPH=""
ENV SCHEDULE=""

# Add certificates
RUN apk update
RUN apk add --no-cache ca-certificates
RUN apk add git
# Copy built binary from builder
COPY --from=builder app .
# Expose port
EXPOSE 10000
# Exec built binary
CMD ./app
