package model

type DeploymentInfo struct {
	Name          string
	Namespace     string
	ContainerName string
	Image         string
}

type DaemonSetInfo struct {
	Name          string
	Namespace     string
	ContainerName string
	Image         string
}
