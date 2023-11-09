package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"net/http"
)

type K8sClient struct {
	Client kubernetes.Interface
}

type DeploymentController struct {
	*K8sClient
}

func NewDeploymentController(kubeConfig *K8sClient) DeploymentController {
	return DeploymentController{K8sClient: kubeConfig}
}

func (dc *DeploymentController) ListDeployment(ctx *gin.Context) {
	deployments, err := dc.Client.AppsV1().Deployments(apiv1.NamespaceDefault).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"data": deployments})
}
