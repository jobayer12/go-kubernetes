package deployment

import "github.com/gin-gonic/gin"

type Route struct {
	controller Controller
}

func NewDeploymentRoute(deploymentController Controller) Route {
	return Route{deploymentController}
}

func (r *Route) DeploymentRoute(rg *gin.RouterGroup) {
	router := rg.Group(":namespace/deployments")
	router.GET("", r.controller.ListDeployment)
	router.GET(":name", r.controller.GetDeployment)
	router.DELETE(":name", r.controller.DeleteDeployment)
	router.PUT(":name/:replica", r.controller.UpdateDeploymentReplica)
	router.GET(":name/scale", r.controller.ReadDeploymentScale)
}
