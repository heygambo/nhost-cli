package cmd

import (
	"context"
	"io/fs"
	"net/http"
	"plugin"

	"github.com/docker/docker/client"
	"github.com/mrinalwahal/cli/hasura"
	"github.com/mrinalwahal/cli/nhost"
)

type (

	// Container service
	Container struct {
		Image               string
		Name                string
		Command             []string
		Entrypoint          string
		Environment         map[string]interface{}
		Ports               []string
		Restart             string
		User                string
		Volumes             []string
		DependsOn           []string `yaml:"depends_on,omitempty"`
		EnvFile             []string `yaml:"env_file,omitempty"`
		Build               map[string]string
		HealthCheckEndpoint string
	}

	ExecResult struct {
		StdOut   string
		StdErr   string
		ExitCode int
	}

	Function struct {
		Route        string
		File         fs.FileInfo
		Path         string
		Handler      func(http.ResponseWriter, *http.Request)
		Base         string
		Build        string
		ServerConfig string
		Plugin       *plugin.Plugin
	}

	Environment struct {
		Name    string
		Active  bool
		Cancel  context.CancelFunc
		Port    int
		HTTP    *http.Client
		Hasura  *hasura.Client
		Docker  *client.Client
		Config  nhost.Configuration
		Context context.Context
		Network string
	}
)
