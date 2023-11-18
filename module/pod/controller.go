package pod

import (
	"context"
	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

type ListPodResponse struct {
	v1.PodList `json:",inline"`
}

type GetPodResponse struct {
	v1.Pod `json:",inline"`
}

type K8sClient struct {
	Client kubernetes.Interface
}

type Controller struct {
	*K8sClient
}

func NewPodController(k8sClient *K8sClient) Controller {
	return Controller{K8sClient: k8sClient}
}

// ListPod
// @Summary			Get the List of Pod.
// @Description		Return list of Pod.
// @Tags			pod
// @Router			/api/v1/namespaces/{namespace}/pods [get]
// @Param 			namespace path string true "Namespace" default(default)
// @Response		200 {array} ListPodResponse
// @Produce			application/json
func (p *Controller) ListPod(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	pods, err := p.Client.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, pods)
}

// GetPod
// @Summary			Get Pod.
// @Description		Return Pod.
// @Tags			pod
// @Router			/api/v1/namespaces/{namespace}/pods/{podName} [get]
// @Param 			namespace path string true "Namespace" default(default)
// @Param 			podName path string true "Pod name"
// @Response		200 {object} GetPodResponse
// @Produce			application/json
func (p *Controller) GetPod(ctx *gin.Context) {
	namespace := ctx.Param("namespace")
	name := ctx.Param("podName")
	pods, err := p.Client.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, pods)
}
