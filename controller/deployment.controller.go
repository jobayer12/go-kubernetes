package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"net/http"
)

type DeploymentListResponse struct {
	Status int64             `json:"status"`
	Err    string            `json:"err"`
	Data   v1.DeploymentList `json:"data"`
}

type DeploymentByNameResponse struct {
	Status int64         `json:"status"`
	Err    string        `json:"err"`
	Data   v1.Deployment `json:"data"`
}

type K8sClient struct {
	Client kubernetes.Interface
}

type DeploymentController struct {
	*K8sClient
}

func NewDeploymentController(kubeConfig *K8sClient) DeploymentController {
	return DeploymentController{K8sClient: kubeConfig}
}

// ListDeployment godoc
// @Summary			Get the List of default namespace deployment.
// @Description		Return list of deployment.
// @Tags			deployment
// @Router			/apis/apps/v1/{namespace}/deployments [get]
// @Param 			namespace path string true "Namespace"
// @Response		200 {object} DeploymentListResponse
// @Produce			application/json
func (dc *DeploymentController) ListDeployment(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	deployments, err := dc.Client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": make([]v1.DeploymentList, 0), "status": http.StatusBadRequest, "err": err})
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": deployments, "status": http.StatusOK, "err": nil})
}

// GetDeployment godoc
// @Summary			Get deployment by name.
// @Description		Return deployment.
// @Tags			deployment
// @Router			/apis/apps/v1/:namespace/deployments/:name [get]
// @Response		200 {object} DeploymentByNameResponse
// @Produce			application/json
//func (dc *DeploymentController) GetDeployment(ctx *gin.Context) {
//	result, err := dc.Client.AppsV1().Deployments(apiv1.NamespaceDefault).Get(context.TODO(), "testingapi", metav1.GetOptions{})
//
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "status": http.StatusBadRequest, "err": err})
//		log.Fatal(err)
//	}
//	ctx.JSON(http.StatusOK, gin.H{"data": result, "status": http.StatusOK, "err": nil})
//}
