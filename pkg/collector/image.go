package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/zufardhiyaulhaq/dockerhub-exporter/pkg/client"
)

const DefaultMetricsValue float64 = 1

type DockerhubImageCollector struct {
	dockerhubImageMetrics *prometheus.Desc
	client                client.KubernetesClient
}

func (collector *DockerhubImageCollector) Describe(channel chan<- *prometheus.Desc) {
	channel <- collector.dockerhubImageMetrics
}

func (collector *DockerhubImageCollector) Collect(channel chan<- prometheus.Metric) {
	deployments := collector.client.GetDockerhubDeployments()

	for _, deployment := range deployments {
		channel <- prometheus.MustNewConstMetric(collector.dockerhubImageMetrics, prometheus.CounterValue, DefaultMetricsValue, deployment.Namespace, deployment.Name, deployment.ContainerName, deployment.Image)
	}
}

func NewDockerhubImageCollector(client client.KubernetesClient) *DockerhubImageCollector {
	return &DockerhubImageCollector{
		dockerhubImageMetrics: prometheus.NewDesc("deployment_dockerhub_image",
			"metrics about dockerhub deployment",
			[]string{"namespace", "name", "container_name", "image"}, nil,
		),
		client: client,
	}
}
