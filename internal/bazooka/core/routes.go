package core

import (
	"bazooka/internal/bazooka/api"
	"bazooka/internal/bazooka/api/article"
	"github.com/gin-gonic/gin"
	"net/http"

	routeConfig "bazooka/internal/pkg/config"
)

var PublicRouteSet = routeConfig.RouteSet{
	Prefix: "/",
	Routes: []routeConfig.Route{
		{http.MethodGet, "/healthz", gin.HandlersChain{api.HealthzHandler}},
		article.SubmitArticle,
	},
}
