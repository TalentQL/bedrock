package injection

import (
	"github.com/TalentQL/bedrock/pkg/util"
	"github.com/gin-gonic/gin"
)

const (
	workspaceContextKey = "workspace_context"
)

func SetWorkspace(c *gin.Context, v util.Object) {
	c.Set(workspaceContextKey, v)
}

func GetWorkspace(c *gin.Context) util.Object {
	w, ok := c.Get(workspaceContextKey)
	if !ok {
		return util.Object{}
	}

	v := w.(util.Object)
	return v
}

func MustGetWorkspace(c *gin.Context) util.Object {
	w := c.MustGet(workspaceContextKey)

	v := w.(util.Object)
	return v
}
