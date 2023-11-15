package controller

import "github.com/gin-gonic/gin"

type DeploymentRouteController struct {
	deploymentController DeploymentController
}

func NewRouteDeploymentController(deploymentController DeploymentController) DeploymentRouteController {
	return DeploymentRouteController{deploymentController}
}

func (dc *DeploymentRouteController) DeploymentRoute(rg *gin.RouterGroup) {
	router := rg.Group(":namespace/deployments")
	router.GET("", dc.deploymentController.ListDeployment)
}
