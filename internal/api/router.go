package api

import (
	"context"
	"github.com/DokiDoki1103/atob/internal/controller/pull_ctr"
	"github.com/codfrm/cago/server/mux"
)

func Router(ctx context.Context, root *mux.Router) error {
	r := root.Group("/atob/v1")
	pull := pull_ctr.NewPull()
	r.GET("/docker/pull", pull.DockerPull)

	return nil
}
