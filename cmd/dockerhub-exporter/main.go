package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zufardhiyaulhaq/dockerhub-exporter/pkg/client"
	"github.com/zufardhiyaulhaq/dockerhub-exporter/pkg/collector"
	"github.com/zufardhiyaulhaq/dockerhub-exporter/pkg/middleware"
	"github.com/zufardhiyaulhaq/dockerhub-exporter/pkg/settings"
)

func main() {
	log.Println("Starting dockerhub-exporter")
	settings := settings.NewSettings()
	log.Println("excluded repository")
	log.Println(settings.ExcludedRegistry)

	client := client.KubernetesClient{
		Settings: settings,
	}
	client.Start()

	DockerhubImageCollector := collector.NewDockerhubImageCollector(client)
	prometheus.MustRegister(DockerhubImageCollector)

	router := http.NewServeMux()
	router.Handle("/metrics", promhttp.Handler())
	router.Handle("/healthz", middleware.StatusHandler(client))
	router.Handle("/readyz", middleware.StatusHandler(client))

	err := http.ListenAndServe(":9125", router)
	log.Fatal(err)
}
