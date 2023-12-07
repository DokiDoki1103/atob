package docker

import (
	"context"
	"github.com/codfrm/cago"
	"github.com/codfrm/cago/configs"
	"github.com/docker/docker/client"
	"io"
)

type DockerInterface interface {
	Build(ctx context.Context, path string, imageName string, imageTag string, cache bool) error
	Pull(ctx context.Context, imageName, username, password string) (io.ReadCloser, error)
}

var defaultDocker *docker

type docker struct {
	dockerClient *client.Client
}

func (d *docker) Start(ctx context.Context, cfg *configs.Config) error {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	d.dockerClient = cli
	return nil
}

func (d *docker) CloseHandle() {
	d.dockerClient.Close()
}

func Docker() cago.Component {
	defaultDocker = &docker{}
	return defaultDocker
}

func Default() DockerInterface {
	return defaultDocker
}
