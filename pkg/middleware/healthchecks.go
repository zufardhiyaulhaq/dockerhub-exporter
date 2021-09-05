package middleware

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/zufardhiyaulhaq/dockerhub-exporter/pkg/client"
)

func StatusHandler(client client.KubernetesClient) http.Handler {
	ok, err := client.GetStatus()
	if ok {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte(err.Error()))
		log.Error(err)
	})
}
