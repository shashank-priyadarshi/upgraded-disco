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

// TODO
func (g *GraphQL) GraphQL(ctx *fasthttp.RequestCtx) {

}
