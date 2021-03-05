package config

type CSP string

const (
	CONTROL_PLANE = "control-plane"
	WORKER        = "worker"

	CSP_AWS CSP = "aws"
	CSP_GCP CSP = "gcp"
)
