package errors

import (
	"errors"

	"github.com/audi-skripsi/event_bus_migrator/pkg/dto"
)

var (
	ErrBadRequest     = errors.New("bad request")
	ErrInternalServer = errors.New("internal server error")
)

var errorMapping = map[error]dto.ErrorResponse{
	ErrBadRequest:     {Code: 400, Message: ErrBadRequest.Error()},
	ErrInternalServer: {Code: 500, Message: ErrInternalServer.Error()},
}

func GetErrorResponse(er error) (errRes dto.ErrorResponse) {
	errRes, found := errorMapping[er]
	if !found {
		errRes = errorMapping[ErrInternalServer]
	}
	return
}
