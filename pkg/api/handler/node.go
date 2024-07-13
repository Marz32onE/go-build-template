package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// GetNodeResources handles the GET request for node resources
func GetNodeResources(c echo.Context) error {
	kubeConfigPath := os.Getenv("KUBE_CONFIG_PATH")
	if kubeConfigPath == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "KUBE_CONFIG_PATH environment variable is required"})
	}

	// Get the kubeconfig file and context from the query parameters
	contextName := c.QueryParam("context")
	if contextName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "kubeconfig and context query parameters are required"})
	}

	// Load kubeconfig
	config, err := clientcmd.LoadFromFile(kubeConfigPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Set the context
	config.CurrentContext = contextName

	// Create REST config
	clientConfig, err := clientcmd.NewDefaultClientConfig(*config, &clientcmd.ConfigOverrides{}).ClientConfig()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Get node resources
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	nodeResources := make([]map[string]interface{}, len(nodes.Items))
	for i, node := range nodes.Items {
		cpu := node.Status.Allocatable["cpu"]
		memory := node.Status.Allocatable["memory"]
		nodeResources[i] = map[string]interface{}{
			"name":   node.Name,
			"cpu":    cpu.String(),
			"memory": memory.String(),
		}
	}
	log.Logger.Debug().Interface("nodeResources", nodeResources).Msg("Response from GetNodeResources")
	return c.JSON(http.StatusOK, nodeResources)
}
