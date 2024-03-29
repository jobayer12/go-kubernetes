package namespace

import "github.com/gin-gonic/gin"

type Route struct {
	controller Controller
}

func NewNamespaceRoute(controller Controller) Route {
	return Route{controller}
}

func (r *Route) Route(router *gin.RouterGroup) {
	router.GET("", r.controller.ListNamespace)
}
