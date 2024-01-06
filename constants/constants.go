package constants

import "net/http"

type HTTPStatusCode int

const (
	HTTPStatusOK       HTTPStatusCode = http.StatusOK
	HTTPStatusCreated  HTTPStatusCode = http.StatusCreated
	HTTPStatusAccepted HTTPStatusCode = http.StatusAccepted

	HTTPStatusBadRequest         HTTPStatusCode = http.StatusBadRequest
	HTTPStatusUnauthorized       HTTPStatusCode = http.StatusUnauthorized
	HTTPStatusForbidden          HTTPStatusCode = http.StatusForbidden
	HTTPStatusNotFound           HTTPStatusCode = http.StatusNotFound
	HTTPStatusPreconditionFailed HTTPStatusCode = http.StatusPreconditionFailed

	HTTPStatusInternalServerError HTTPStatusCode = http.StatusInternalServerError
	HTTPStatusBadGateway          HTTPStatusCode = http.StatusBadGateway
	HTTPStatusServiceUnavailable  HTTPStatusCode = http.StatusServiceUnavailable
)

const (
	DB_MARIADB = "Maria"
	DB_MONGODB = "Mongo"
	DB_REDIS   = "Redis"
)
