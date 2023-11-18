package pod

import "github.com/gin-gonic/gin"

type Route struct {
	controller Controller
}

func NewPodRoute(controller Controller) Route {
	return Route{controller}
}

func (r *Route) Route(router *gin.RouterGroup) {
	router.GET("", r.controller.ListPod)
	router.GET(":podName", r.controller.GetPod)
}
