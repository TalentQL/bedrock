package injection

import (
	"github.com/TalentQL/bedrock/pkg/util"
	"github.com/gin-gonic/gin"
)

const (
	userContextKey = "user_context"
)

func SetUser(c *gin.Context, v util.Object) {
	c.Set(userContextKey, v)
}

func GetUser(c *gin.Context) util.Object {
	u, ok := c.Get(userContextKey)
	if !ok {
		return util.Object{}
	}

	v := u.(util.Object)
	return v
}

func MustGetUser(c *gin.Context) util.Object {
	u := c.MustGet(userContextKey)

	v := u.(util.Object)
	return v
}
