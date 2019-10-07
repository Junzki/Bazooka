package config

import "github.com/gin-gonic/gin"

type Route struct {
	Method  string
	Route   string
	Handler gin.HandlerFunc
}

type RouteSet struct {
	Prefix     string
	Routes     []Route
	Middleware []gin.HandlerFunc
}

// Bind binds middleware and routes to gin.RouterGroup
func (s *RouteSet) Bind(r gin.IRoutes) {
	if nil != s.Middleware && 0 < len(s.Middleware) {
		r = r.Use(s.Middleware...)
	}

	for _, route := range s.Routes {
		r.Handle(route.Method, route.Route, route.Handler)
	}
}
