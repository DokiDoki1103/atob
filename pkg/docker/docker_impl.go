package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/archive"
	"io"
	"os"
)

func (d docker) Build(ctx context.Context, path string, imageName string, imageTag string, noCache bool) error {
	tar, err := archive.TarWithOptions(path, &archive.TarOptions{})
	if err != nil {
		return err
	}
	// 构建 Docker 镜像
	res, err := d.dockerClient.ImageBuild(ctx, tar, types.ImageBuildOptions{
		NoCache:    noCache,
		Tags:       []string{fmt.Sprintf("%s:%s", imageName, imageTag)},
		Dockerfile: "Dockerfile",
	})
	defer res.Body.Close()

	_, err = io.Copy(os.Stdout, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func (d docker) Pull(ctx context.Context, imageName, username, password string) (io.ReadCloser, error) {
	fmt.Printf("Pulling image %s\n", imageName, d.dockerClient)
	return d.dockerClient.ImagePull(ctx, imageName, types.ImagePullOptions{})
}
