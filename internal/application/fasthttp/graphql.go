package fasthttp

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"github.com/valyala/fasthttp"
)

type GraphQL struct {
	ports.GraphQLOps
	logger.Logger
}

func (g *GraphQL) GraphQL(ctx *fasthttp.RequestCtx) {
	data, err := g.GraphQLOps.GraphQL(ctx.Request.Body())
	if err != nil || data == nil {
		ctx.Err()
		return
	}
	ctx.SetBody(data.([]byte))
	ctx.Done()
}
