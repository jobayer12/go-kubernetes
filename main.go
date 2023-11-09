package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jobayer12/go-kubernetes/controller"
	_ "github.com/jobayer12/go-kubernetes/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

type K8sClient struct {
	Client kubernetes.Interface
}

var (
	server                    *gin.Engine
	DeploymentController      controller.DeploymentController
	DeploymentRouteController controller.DeploymentRouteController
)

func getK8sClient() *K8sClient {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error getting user home dir: %v\n", err)
		os.Exit(1)
	}
	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatal(err)
	}
	client, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatal(err)
	}
	return &K8sClient{
		Client: client,
	}
}

func init() {
	client := getK8sClient()
	DeploymentController = controller.NewDeploymentController((*controller.K8sClient)(client))
	DeploymentRouteController = controller.NewRouteDeploymentController(DeploymentController)

	server = gin.Default()

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// @title Kubernetes API
// @version 1.0
// @description List of kubernetes API
// @host localhost:8080
// @BasePath /
func main() {
	server.ForwardedByClientIP = true
	server.ForwardedByClientIP = true
	err := server.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal(err)
	}

	route := server.Group("/apis/apps/v1")
	DeploymentRouteController.DeploymentRoute(route)

	log.Fatal(server.Run(":8080"))
}
