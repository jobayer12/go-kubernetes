package deployment

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

type ListResponse struct {
	v1.DeploymentList `json:",inline"`
}

type GetDeploymentResponse struct {
	v1.Deployment `json:",inline"`
}

type DeleteDeploymentResponse struct {
	bool `json:",inline"`
}

type ScaleDeploymentResponse struct {
	autoscalingv1.Scale `json:",inline"`
}

type K8sClient struct {
	Client kubernetes.Interface
}

type Controller struct {
	*K8sClient
}

func NewDeploymentController(kubeConfig *K8sClient) Controller {
	return Controller{K8sClient: kubeConfig}
}

// ListDeployment godoc
// @Summary			Get the List of default namespace deployment.
// @Description		Return list of deployment.
// @Tags			deployment
// @Router			/apis/apps/v1/{namespace}/deployments [get]
// @Param 			namespace path string true "Namespace" default(default)
// @Response		200 {array} ListResponse
// @Produce			application/json
func (dc *Controller) ListDeployment(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	deployments, err := dc.Client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, deployments)
}

// GetDeployment godoc
// @Summary			Get deployment by name.
// @Description		Return deployment.
// @Tags			deployment
// @Router			/apis/apps/v1/{namespace}/deployments/{name} [get]
// @Param 			namespace path string true "Namespace" default(default)
// @Param 			name path string true "Deployment name"
// @Response		200 {object} GetDeploymentResponse
// @Produce			application/json
func (dc *Controller) GetDeployment(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	result, err := dc.Client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, result)
}

// DeleteDeployment
// @Summary			Delete deployment
// @Tags			deployment
// @Router			/apis/apps/v1/{namespace}/deployments/{name} [delete]
// @Param 			namespace path string true "Namespace" default(default)
// @Param 			name path string true "Deployment name"
// @response     	default {boolean}  boolean true
// @Produce			application/json
func (dc *Controller) DeleteDeployment(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	err := dc.Client.AppsV1().Deployments(namespace).Delete(context.Background(), name, metav1.DeleteOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, true)
}

// ReadDeploymentScale
// @Summary			Scale deployment
// @Tags			deployment
// @Router			/apis/apps/v1/{namespace}/deployments/{name}/scale [get]
// @Param 			namespace path string true "Namespace" default(default)
// @Param 			name path string true "Deployment name"
// @Response		200 {object} ScaleDeploymentResponse
// @Produce			application/json
func (dc *Controller) ReadDeploymentScale(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	scaleObj, err := dc.Client.AppsV1().Deployments(namespace).GetScale(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, scaleObj)
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
func (dc *Controller) UpdateDeploymentReplica(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	replicaParam, err := strconv.ParseInt(ctx.Param("replica"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	scaleObj, err := dc.Client.AppsV1().Deployments(namespace).GetScale(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	replica := int32(replicaParam)
	sd := *scaleObj
	if sd.Spec.Replicas == replica || replica < 0 {
		ctx.JSON(http.StatusBadRequest, "No changes applied")
		return
	}
	sd.Spec.Replicas = replica
	scaleDeployment, err := dc.Client.AppsV1().Deployments(namespace).UpdateScale(context.Background(), name, &sd, metav1.UpdateOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, scaleDeployment)
}
