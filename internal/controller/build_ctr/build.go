package build_ctr

import (
	"context"

	api "github.com/DokiDoki1103/atob/internal/api/build"
	"github.com/DokiDoki1103/atob/internal/service/build_svc"
)

type Build struct {
}

func NewBuild() *Build {
	return &Build{}
}

// Build -
func (b *Build) Build(ctx context.Context, req *api.BuildRequest) (*api.BuildResponse, error) {
	return build_svc.Build().Build(ctx, req)
}
