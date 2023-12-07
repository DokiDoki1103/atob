package pull

import "github.com/codfrm/cago/server/mux"

type DockerPullRequest struct {
	mux.Meta   `path:"/docker/pull" method:"POST"`
	SourceCode string `json:"imageName" binding:"required,max=10485760" label:"镜像名称"`
}

type DockerPullResponse struct {
}

type GitPullRequest struct {
	mux.Meta   `path:"/git/pull" method:"POST"`
	SourceCode string `json:"sourceCode" binding:"required,max=10485760" label:"源码的下载地址"`
}

type GitPullResponse struct {
}
