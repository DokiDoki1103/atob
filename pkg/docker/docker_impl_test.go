package docker

import (
	"context"
	"github.com/codfrm/cago/pkg/logger"
	"testing"
)

func Test_docker_Build(t *testing.T) {
	ctx := context.Background()
	c := &docker{}
	err := c.Start(ctx, nil)
	if err != nil {
		logger.Ctx(ctx).Error(err.Error())
		return
	}

	c.Build(ctx, "/Users/zhangxiaoyuan/IdeaProjects/node-zhihuishu", "zhihuishu", "latest", false)
}

func Test_docker_Pull(t *testing.T) {
	ctx := context.Background()
	c := &docker{}
	err := c.Start(ctx, nil)
	if err != nil {
		logger.Ctx(ctx).Error(err.Error())
		return
	}

	c.Pull(ctx, "nginx", "", "latest")
}
