package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	defaultError = "unknown-error"
)

type ErrorInfo struct {
	Status  int
	Message string
}

var errMap = map[string] ErrorInfo {
	"required-field-missing": {http.StatusBadRequest, "Required field missing."},
	"unknown-error": {http.StatusNotAcceptable, "Unknown error."},
}


func Abort(c *gin.Context, e string, m interface{}) {
	i, ok := errMap[e]
	if ! ok {
		e = defaultError
		i, _ = errMap[e]
	}

	switch v := m.(type) {
	case string:
		i.Message = v
	case []byte:
		i.Message = string(v)
	case error:
		i.Message = v.Error()
	}

	c.AbortWithStatusJSON(i.Status, gin.H{
		"error_code": e,
		"message": i.Message,
	})
}
