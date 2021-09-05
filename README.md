# dockerhub-exporter

dockerhub-exporter, export deployment with dockerhub image information to prometheus

![Version: 1.1.0](https://img.shields.io/badge/Version-1.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.1.0](https://img.shields.io/badge/AppVersion-1.1.0-informational?style=flat-square) [![made with Go](https://img.shields.io/badge/made%20with-Go-brightgreen)](http://golang.org)

## Installing

To install the chart with the release name `my-release`:

```console
helm repo add zufardhiyaulhaq https://charts.zufardhiyaulhaq.com/
helm install my-release zufardhiyaulhaq/dockerhub-exporter --values values.yaml
```

### Example Metrics
```
# HELP deployment_dockerhub_image metrics about dockerhub deployment
# TYPE deployment_dockerhub_image counter
deployment_dockerhub_image{container_name="external-dns",image="docker.io/bitnami/external-dns:0.9.0-debian-10-r0",name="external-dns-digitalocean",namespace="infrastructure"} 1
deployment_dockerhub_image{container_name="grafana",image="grafana/grafana:7.2.1",name="prometheus-grafana",namespace="infrastructure"} 1
deployment_dockerhub_image{container_name="grafana-sc-dashboard",image="kiwigrid/k8s-sidecar:1.1.0",name="prometheus-grafana",namespace="infrastructure"} 1
deployment_dockerhub_image{container_name="grpc-testing",image="zufardhiyaulhaq/grpc-testing",name="grpc-testing",namespace="default"} 1
deployment_dockerhub_image{container_name="community-operator-container",image="cloudnativeid/community-operator:0.0.6",name="community-operator-deployment",namespace="nodejs-community"} 1

# HELP daemonset_dockerhub_image metrics about dockerhub daemonset
# TYPE daemonset_dockerhub_image counter
daemonset_dockerhub_image{container_name="csi-cephfsplugin",image="quay.io/cephcsi/cephcsi:v3.2.0",name="csi-cephfsplugin",namespace="infrastructure"} 1
daemonset_dockerhub_image{container_name="csi-rbdplugin",image="quay.io/cephcsi/cephcsi:v3.2.0",name="csi-rbdplugin",namespace="infrastructure"} 1
daemonset_dockerhub_image{container_name="driver-registrar",image="k8s.gcr.io/sig-storage/csi-node-driver-registrar:v2.0.1",name="csi-cephfsplugin",namespace="infrastructure"} 1
```

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| excludedRepository | string | `"ghcr.io,quay.io,k8s.gcr.io,gcr.io"` |  |
| image.name | string | `"zufardhiyaulhaq/dockerhub-exporter"` |  |
| image.tag | string | `"v1.1.0"` |  |
| pullPolicy | string | `"Always"` |  |
| serviceMonitor.enabled | bool | `true` |  |

