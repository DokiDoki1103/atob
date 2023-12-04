package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()
	tar, err := archive.TarWithOptions("/Users", &archive.TarOptions{})
	if err != nil {
		return
	}

	// 构建 Docker 镜像
	res, err := cli.ImageBuild(ctx, tar, types.ImageBuildOptions{
		NoCache:    true,
		Tags:       []string{"test:v1"},
		Dockerfile: "Dockerfile",
	})
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}

}
