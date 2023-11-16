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
	router.GET(":name", dc.deploymentController.GetDeployment)
	router.DELETE(":name", dc.deploymentController.DeleteDeployment)
	router.PUT(":name/:replica", dc.deploymentController.UpdateDeploymentReplica)
	router.GET(":name/scale", dc.deploymentController.ReadDeploymentScale)
}
