package build_svc

import (
	"context"

	api "github.com/DokiDoki1103/atob/internal/api/build"
)

type BuildSvc interface {
	// Build -
	Build(ctx context.Context, req *api.BuildRequest) (*api.BuildResponse, error)
}

type buildSvc struct {
}

var defaultBuild = &buildSvc{}

func Build() BuildSvc {
	return defaultBuild
}

// Build -
func (b *buildSvc) Build(ctx context.Context, req *api.BuildRequest) (*api.BuildResponse, error) {
	return nil, nil
}
