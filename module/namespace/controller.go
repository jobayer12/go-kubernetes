package namespace

import (
	"context"
	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
)

type K8sClient struct {
	Client kubernetes.Interface
}

type Controller struct {
	*K8sClient
}

func NewNamespaceController(k8sClient *K8sClient) Controller {
	return Controller{K8sClient: k8sClient}
}

// ListNamespace
// @Summary			Get the List of namespace.
// @Description		Return list of namespace.
// @Tags			namespace
// @Router			/api/v1/namespaces [get]
// @Response		200 {object} v1.NamespaceList
// @Produce			application/json
func (ns *Controller) ListNamespace(ctx *gin.Context) {
	namespaces, err := ns.Client.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, namespaces)
}
