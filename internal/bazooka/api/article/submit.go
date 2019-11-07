package article

import (
	"bazooka/internal/bazooka/api"
	"bazooka/internal/pkg/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ctxKey = "request-args"
)

var SubmitArticle = config.Route{
	Method: http.MethodPost,
	Route:  "/post",
	Handlers: gin.HandlersChain{
		parseRequest,
		submit,
	},
}

func parseRequest(c *gin.Context) {
	var (
		err error = nil
		arg       = submitRequest{}
	)

	err = c.ShouldBindJSON(&arg)
	if nil != err {
		api.Abort(c, "required-field-missing", err)
		return
	}

	c.Set(ctxKey, &arg)
}

func submit(c *gin.Context) {
	ptr, ok := c.Get(ctxKey)
	if !ok {
		api.Abort(c, "required-field-missing", nil)
		return
	}

	arg := ptr.(*submitRequest)

	c.JSON(http.StatusOK, arg)
}
