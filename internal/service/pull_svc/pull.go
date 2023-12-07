package pull_svc

import (
	"context"

	api "github.com/DokiDoki1103/atob/internal/api/pull"
)

type PullSvc interface {
	// DockerPull TODO
	DockerPull(ctx context.Context, req *api.DockerPullRequest) (*api.DockerPullResponse, error)
	// GitPull TODO
	GitPull(ctx context.Context, req *api.GitPullRequest) (*api.GitPullResponse, error)
}

type pullSvc struct {
}

var defaultPull = &pullSvc{}

func Pull() PullSvc {
	return defaultPull
}

// DockerPull TODO
func (p *pullSvc) DockerPull(ctx context.Context, req *api.DockerPullRequest) (*api.DockerPullResponse, error) {

	return nil, nil
}

// GitPull TODO
func (p *pullSvc) GitPull(ctx context.Context, req *api.GitPullRequest) (*api.GitPullResponse, error) {
	return nil, nil
}
