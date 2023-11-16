package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
	"strconv"
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

type DeleteDeploymentResponse struct {
	Status int64  `json:"status"`
	Err    string `json:"err"`
	Data   bool   `json:"data"`
}

type ScaleDeploymentResponse struct {
	Status int64               `json:"status"`
	Err    string              `json:"err"`
	Data   autoscalingv1.Scale `json:"Data"`
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
// @Param 			namespace path string true "Namespace" default(default)
// @Response		200 {object} DeploymentListResponse
// @Produce			application/json
func (dc *DeploymentController) ListDeployment(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	deployments, err := dc.Client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": make([]v1.DeploymentList, 0), "status": http.StatusBadRequest, "err": err})
	}
	ctx.JSON(http.StatusOK, gin.H{"data": deployments, "status": http.StatusOK, "err": nil})
}

// GetDeployment godoc
// @Summary			Get deployment by name.
// @Description		Return deployment.
// @Tags			deployment
// @Router			/apis/apps/v1/{namespace}/deployments/{name} [get]
// @Param 			namespace path string true "Namespace" default(default)
// @Param 			name path string true "Deployment name"
// @Response		200 {object} DeploymentByNameResponse
// @Produce			application/json
func (dc *DeploymentController) GetDeployment(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	result, err := dc.Client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "status": http.StatusBadRequest, "err": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result, "status": http.StatusOK, "err": err})
}

// DeleteDeployment
// @Summary			Delete deployment
// @Tags			deployment
// @Router			/apis/apps/v1/{namespace}/deployments/{name} [delete]
// @Param 			namespace path string true "Namespace" default(default)
// @Param 			name path string true "Deployment name"
// @Response		200 {object} DeleteDeploymentResponse
// @Produce			application/json
func (dc *DeploymentController) DeleteDeployment(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	err := dc.Client.AppsV1().Deployments(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": false, "status": http.StatusBadRequest, "err": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": true, "status": http.StatusOK, "err": err})
}

// ReadDeploymentScale
// @Summary			Scale deployment
// @Tags			deployment
// @Router			/apis/apps/v1/{namespace}/deployments/{name}/scale [get]
// @Param 			namespace path string true "Namespace" default(default)
// @Param 			name path string true "Deployment name"
// @Response		200 {object} ScaleDeploymentResponse
// @Produce			application/json
func (dc *DeploymentController) ReadDeploymentScale(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	scaleObj, err := dc.Client.AppsV1().Deployments(namespace).GetScale(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "status": http.StatusBadRequest, "err": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": scaleObj, "status": http.StatusOK, "err": nil})
}

// UpdateDeploymentReplica
// @Summary			Update Deployment Replica
// @Tags			deployment
// @Router			/apis/apps/v1/{namespace}/deployments/{name}/{replica} [put]
// @Param 			namespace path string true "Namespace" default(default)
// @Param 			name path string true "Deployment name"
// @Param 			replica path string true "Replica"
// @Response		200 {object} ScaleDeploymentResponse
// @Produce			application/json
func (dc *DeploymentController) UpdateDeploymentReplica(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	replicaParam, err := strconv.ParseInt(ctx.Param("replica"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "status": http.StatusBadRequest, "err": err})
		return
	}
	scaleObj, err := dc.Client.AppsV1().Deployments(namespace).GetScale(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "status": http.StatusBadRequest, "err": err})
		return
	}
	replica := int32(replicaParam)
	sd := *scaleObj
	if sd.Spec.Replicas == replica || replica < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "status": http.StatusBadRequest, "err": "No changes applied"})
		return
	}
	sd.Spec.Replicas = replica
	scaleDeployment, err := dc.Client.AppsV1().Deployments(namespace).UpdateScale(context.Background(), name, &sd, metav1.UpdateOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "status": http.StatusBadRequest, "err": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": scaleDeployment, "status": http.StatusOK, "err": nil})
}
