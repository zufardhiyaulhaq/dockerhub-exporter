package client

import (
	"context"
	"flag"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/distribution/distribution/reference"
	"github.com/zufardhiyaulhaq/dockerhub-exporter/pkg/model"
	"github.com/zufardhiyaulhaq/dockerhub-exporter/pkg/settings"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type KubernetesClient struct {
	Client   *kubernetes.Clientset
	Settings settings.Settings
}

func (c *KubernetesClient) Start() {
	var config *rest.Config
	var err error

	if c.Settings.UseServiceAccount {
		log.Info("Using serviceaccount credential")
		config, err = rest.InClusterConfig()
	} else {
		log.Info("Using kubeconfig file credential")
		var kubeConfig *string
		if home := homedir.HomeDir(); home != "" {
			kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		}
		flag.Parse()
		config, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
	}

	if err != nil {
		log.Errorln(err)
	}

	c.Client, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Errorln(err)
	}
}

func (c *KubernetesClient) GetDockerhubDeployments() []model.DeploymentInfo {
	var deploymentData []model.DeploymentInfo

	deploymentsAppsV1, err := c.Client.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Errorln("cannot find deployment")
	}

	for _, deployment := range deploymentsAppsV1.Items {
		for _, container := range deployment.Spec.Template.Spec.Containers {
			repo, err := reference.Parse(container.Image)
			if err != nil {
				log.Errorln("cannot parse container image")
				continue
			}

			if named, ok := repo.(reference.Named); ok {
				domain := reference.Domain(named)

				if !contains(c.Settings.ExcludedRegistry, domain) {
					deploymentData = append(deploymentData, model.DeploymentInfo{
						Name:          deployment.Name,
						Namespace:     deployment.Namespace,
						ContainerName: container.Name,
						Image:         container.Image,
					})
				}
			}
		}
	}

	return deploymentData
}

func (c *KubernetesClient) GetDockerhubDaemonSets() []model.DaemonSetInfo {
	var daemonSetData []model.DaemonSetInfo

	daemonSetAppsV1, err := c.Client.AppsV1().DaemonSets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Errorln("cannot find deployment")
	}

	for _, daemonset := range daemonSetAppsV1.Items {
		for _, container := range daemonset.Spec.Template.Spec.Containers {
			repo, err := reference.Parse(container.Image)
			if err != nil {
				log.Errorln("cannot parse container image")
				continue
			}

			if named, ok := repo.(reference.Named); ok {
				domain := reference.Domain(named)

				if !contains(c.Settings.ExcludedRegistry, domain) {
					daemonSetData = append(daemonSetData, model.DaemonSetInfo{
						Name:          daemonset.Name,
						Namespace:     daemonset.Namespace,
						ContainerName: container.Name,
						Image:         container.Image,
					})
				}
			}
		}
	}

	return daemonSetData
}

func (c *KubernetesClient) GetStatus() (bool, error) {
	version, err := c.Client.ServerVersion()
	if err != nil {
		return false, err
	}

	log.Println("Kubernetes version: " + version.String())
	return true, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
