FROM golang:1.16.5 as builder
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

FROM alpine:3.14 as production
ENV SERVER_PORT=""
ENV TODO_API_PORT=""
ENV MONGO_URI=""
ENV GITHUB_TOKEN=""
ENV DB_NAME=""
ENV COLLECTION_GITHUBDATA=""
ENV COLLECTION_TODOS=""
ENV COLLECTION_BIODATA=""
ENV COLLECTION_SCHEDULE=""
ENV COLLECTION_ARTICLES=""
ENV GITHUB_DATA=""
ENV ISSUE_DATA=""
ENV GH_INTEGRATION_ORIGIN=""
ENV ALLOWED_ORIGIN=""
# Add certificates
RUN apk add --no-cache ca-certificates
# Copy built binary from builder
COPY --from=builder app .
# Expose port
EXPOSE 10000
# Exec built binary
CMD ./app
