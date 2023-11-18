package namespace

import "github.com/gin-gonic/gin"

type Route struct {
	controller Controller
}

func NewNamespaceRoute(controller Controller) Route {
	return Route{controller}
}

func (r *Route) NamespaceRoute(rg *gin.RouterGroup) {
	router := rg.Group("namespaces")

	router.GET("", r.controller.ListNamespace)
	router.GET(":name", r.controller.GetNamespace)
}
