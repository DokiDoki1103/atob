package pull_ctr

import (
	"context"
	api "github.com/DokiDoki1103/atob/internal/api/pull"
	"github.com/DokiDoki1103/atob/internal/service/pull_svc"
	"github.com/DokiDoki1103/atob/pkg/docker"
	"github.com/codfrm/cago/pkg/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
)

type Pull struct {
}

func NewPull() *Pull {
	return &Pull{}
}

// DockerPull TODO
func (p *Pull) DockerPull(c *gin.Context) {
	c.SSEvent("pull", "Pulling image")
	c.Writer.Flush()
	pull, err := docker.Default().Pull(context.Background(), "nginx", "", "")

	if err != nil {
		c.SSEvent("error", "Error pulling image")
		c.Writer.Flush()
		logger.Ctx(c).Error("Error pulling image", zap.Error(err))
		return
	}
	defer func(pull io.ReadCloser) {
		err := pull.Close()
		if err != nil {
			c.SSEvent("error", "Error pulling image")
			c.Writer.Flush()
			logger.Ctx(c).Error("Error closing ReadCloser", zap.Error(err))
		}
	}(pull)

	buf := make([]byte, 1024)
	for {
		n, err := pull.Read(buf)
		if err != nil {
			if err != io.EOF {
				c.SSEvent("error", "Error reading from ReadCloser")
				c.Writer.Flush()
				logger.Ctx(c).Error("Error reading from ReadCloser", zap.Error(err))
			}
			break
		}
		data := buf[:n]
		c.SSEvent("message", string(data))
		c.Writer.Flush()
	}
	c.SSEvent("message", "Pulling image finished")
	c.Writer.Flush()
	logger.Ctx(c).Info("Pulling image finished")
}

// GitPull TODO
func (p *Pull) GitPull(ctx context.Context, req *api.GitPullRequest) (*api.GitPullResponse, error) {
	return pull_svc.Pull().GitPull(ctx, req)
}
