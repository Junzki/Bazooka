package routes

import (
	"net/http"

	"bazooka/internal/pkg/config"
)

var PublicRouteSet = config.RouteSet{
	Prefix: "/",
	Routes: []config.Route{
		{http.MethodGet, "/healthz", HealthzHandler},
	},
}
