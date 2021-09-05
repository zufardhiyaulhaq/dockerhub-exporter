package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/zufardhiyaulhaq/dockerhub-exporter/pkg/client"
)

const DefaultMetricsValue float64 = 1

type DockerhubImageCollector struct {
	dockerHubDeploymentImageMetrics *prometheus.Desc
	dockerHubDaemonSetImageMetrics  *prometheus.Desc
	client                          client.KubernetesClient
}

func (collector *DockerhubImageCollector) Describe(channel chan<- *prometheus.Desc) {
	channel <- collector.dockerHubDeploymentImageMetrics
	channel <- collector.dockerHubDaemonSetImageMetrics
}

func (collector *DockerhubImageCollector) Collect(channel chan<- prometheus.Metric) {
	deployments := collector.client.GetDockerhubDeployments()

	for _, deployment := range deployments {
		channel <- prometheus.MustNewConstMetric(collector.dockerHubDeploymentImageMetrics, prometheus.CounterValue, DefaultMetricsValue, deployment.Namespace, deployment.Name, deployment.ContainerName, deployment.Image)
	}

	daemonSets := collector.client.GetDockerhubDaemonSets()

	for _, daemonSet := range daemonSets {
		channel <- prometheus.MustNewConstMetric(collector.dockerHubDaemonSetImageMetrics, prometheus.CounterValue, DefaultMetricsValue, daemonSet.Namespace, daemonSet.Name, daemonSet.ContainerName, daemonSet.Image)
	}
}

func NewDockerhubImageCollector(client client.KubernetesClient) *DockerhubImageCollector {
	return &DockerhubImageCollector{
		dockerHubDeploymentImageMetrics: prometheus.NewDesc("deployment_dockerhub_image",
			"metrics about dockerhub deployment",
			[]string{"namespace", "name", "container_name", "image"}, nil,
		),
		dockerHubDaemonSetImageMetrics: prometheus.NewDesc("daemonset_dockerhub_image",
			"metrics about dockerhub daemonset",
			[]string{"namespace", "name", "container_name", "image"}, nil,
		),
		client: client,
	}
}
