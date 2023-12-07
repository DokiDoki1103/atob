package build

import "github.com/codfrm/cago/server/mux"

// BuildRequest -
type BuildRequest struct {
	mux.Meta       `path:"/build" method:"POST"`
	SourceCodePath string `json:"sourceCodePath" binding:"required,max=10485760" label:"源码的路径"`
}

// BuildResponse -
type BuildResponse struct {
}
