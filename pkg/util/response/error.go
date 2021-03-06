package response

import (
	//"codeid-boiler/pkg/log"
	//"codeid-boiler/pkg/util/date"
	//"context"
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"
	//"os"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Meta Meta `json:"meta"`
	//Error string `json:"error"`
	Error interface{} `json:"error"`
}

type Error struct {
	Response     errorResponse `json:"response"`
	Code         int           `json:"code"`
	ErrorMessage error
}

const (
	E_DUPLICATE            = "duplicate"
	E_NOT_FOUND            = "not_found"
	E_UNPROCESSABLE_ENTITY = "unprocessable_entity"
	E_UNAUTHORIZED         = "unauthorized"
	E_BAD_REQUEST          = "bad_request"
	E_SERVER_ERROR         = "server_error"
)

type errorConstant struct {
	ForeignKey          Error
	Duplicate           Error
	NotFound            Error
	RouteNotFound       Error
	UnprocessableEntity Error
	Unauthorized        Error
	BadRequest          Error
	Validation          Error
	InternalServerError Error

	NoFileUpload          Error
	OpenFileErr           Error
	UploadFileSrcError    Error
	UploadFileCreateError Error
	UploadFileDestError   Error
	UploadFileError       Error
	SheetFileXLSXSErr     Error
	NotXLSXFileError      Error
}

var ErrorConstant errorConstant = errorConstant{
	OpenFileErr: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Can not open file",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
	SheetFileXLSXSErr: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "File sheetname error",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
	NoFileUpload: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "No files to upload",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
	UploadFileSrcError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Failed to open uploaded file",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
	UploadFileCreateError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Failed to create uploaded file",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
	UploadFileDestError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Failed destination uploaded file",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
	UploadFileError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Failed to upload file",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
	NotXLSXFileError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Document is not a xlsx file",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
	ForeignKey: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Foreign key constraint fails",
			},
			Error: E_DUPLICATE,
		},
		Code: http.StatusConflict,
	},
	Duplicate: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Created value already exists",
			},
			Error: E_DUPLICATE,
		},
		Code: http.StatusConflict,
	},
	NotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Data not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	RouteNotFound: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Route not found",
			},
			Error: E_NOT_FOUND,
		},
		Code: http.StatusNotFound,
	},
	UnprocessableEntity: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Invalid parameters or payload",
			},
			Error: E_UNPROCESSABLE_ENTITY,
		},
		Code: http.StatusUnprocessableEntity,
	},
	Unauthorized: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Unauthorized, please login",
			},
			Error: E_UNAUTHORIZED,
		},
		Code: http.StatusUnauthorized,
	},
	BadRequest: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Bad Request",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	Validation: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Invalid parameters or payload",
			},
			Error: E_BAD_REQUEST,
		},
		Code: http.StatusBadRequest,
	},
	InternalServerError: Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: "Something bad happened",
			},
			Error: E_SERVER_ERROR,
		},
		Code: http.StatusInternalServerError,
	},
}

func ErrorBuilder(res *Error, message error) *Error {
	res.ErrorMessage = message
	return res
}

func CustomErrorBuilder(code int, err string, message string) *Error {
	return &Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: message,
			},
			Error: err,
		},
		Code: code,
	}
}

func CustomErrorBuilderWithData(code int, data interface{}, message string) *Error {
	return &Error{
		Response: errorResponse{
			Meta: Meta{
				Success: false,
				Message: message,
			},
			Error: data,
		},
		Code: code,
	}
}

func ErrorResponse(err error) *Error {
	re, ok := err.(*Error)
	if ok {
		return re
	} else {
		return ErrorBuilder(&ErrorConstant.InternalServerError, err)
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code %d", e.Code)
}

func (e *Error) ParseToError() error {
	return e
}

func (e *Error) Send(c echo.Context) error {
	var errorMessage string
	if e.ErrorMessage != nil {
		errorMessage = fmt.Sprintf("%+v", errors.WithStack(e.ErrorMessage))
	}
	logrus.Error(errorMessage)

	/*
		body, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			logrus.Warn("error read body, message : ", e.Error())
		}

		bHeader, err := json.Marshal(c.Request().Header)
		if err != nil {
			logrus.Warn("error read header, message : ", e.Error())
		}


		go func() {
			retries := 3
			logError := log.LogError{
				ID:           shortid.MustGenerate(),
				Header:       string(bHeader),
				Body:         string(body),
				URL:          c.Request().URL.Path,
				HttpMethod:   c.Request().Method,
				ErrorMessage: errorMessage,
				Level:        "Error",
				AppName:      os.Getenv("APP"),
				Version:      os.Getenv("VERSION"),
				Env:          os.Getenv("ENV"),
				CreatedAt:    *date.DateTodayLocal(),
			}
			for i := 0; i < retries; i++ {
				err := log.InsertErrorLog(context.Background(), &logError)
				if err == nil {
					break
				}
			}
		}()
	*/

	return c.JSON(e.Code, e.Response)
}
