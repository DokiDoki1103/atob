package main

import (
	"context"
	"github.com/DokiDoki1103/atob/internal/api"
	"github.com/DokiDoki1103/atob/pkg/docker"
	"github.com/codfrm/cago"
	"github.com/codfrm/cago/configs"
	"github.com/codfrm/cago/pkg/component"
	"github.com/codfrm/cago/server/mux"
	"log"
)

func main() {
	ctx := context.Background()
	cfg, err := configs.NewConfig("simple")
	if err != nil {
		log.Fatalf("load config err: %v", err)
	}

	err = cago.New(ctx, cfg).
		Registry(component.Core()).
		Registry(docker.Docker()).
		RegistryCancel(mux.HTTP(api.Router)).
		Start()

	if err != nil {
		log.Fatalf("start err: %v", err)
		return
	}
}
