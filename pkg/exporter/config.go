package exporter

import (
	"github.com/opsgenie/kubernetes-event-exporter/pkg/sinks"
)

// Config allows configuration
type Config struct {
	// Route is the top route that the events will match
	// TODO: There is currently a tight coupling with route and config, but not with receiver config and sink so
	// TODO: I am not sure what to do here.
	LogLevel  string                 `yaml:"logLevel"`
	Route     Route                  `yaml:"route"`
	Receivers []sinks.ReceiverConfig `yaml:"receivers"`
	ClusterName string               `yaml:"clusterName,omitempty"`
}

func (c *Config) Validate() error {
	// No duplicate receivers
	// Receivers individually
	// Routers recursive
	return nil
}
