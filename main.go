package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jobayer12/go-kubernetes/docs"
	"github.com/jobayer12/go-kubernetes/module/deployment"
	"github.com/jobayer12/go-kubernetes/module/namespace"
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
	DeploymentController      deployment.Controller
	DeploymentRouteController deployment.Route

	NamespaceController namespace.Controller
	NamespaceRoute      namespace.Route
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
	DeploymentController = deployment.NewDeploymentController((*deployment.K8sClient)(client))
	DeploymentRouteController = deployment.NewDeploymentRoute(DeploymentController)

	NamespaceController = namespace.NewNamespaceController((*namespace.K8sClient)(client))
	NamespaceRoute = namespace.NewNamespaceRoute(NamespaceController)

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

	apiAppsV1RouteGroup := server.Group("/apis/apps/v1")
	DeploymentRouteController.DeploymentRoute(apiAppsV1RouteGroup)

	apiV1Route := server.Group("/api/v1")
	NamespaceRoute.NamespaceRoute(apiV1Route)

	log.Fatal(server.Run(":8080"))
}
