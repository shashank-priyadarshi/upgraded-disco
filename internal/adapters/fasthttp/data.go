package fasthttp

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"github.com/valyala/fasthttp"
)

type Data struct {
	ports.DataOps
	logger.Logger
}

func (d *Data) GetGraphData(ctx *fasthttp.RequestCtx) {

}

func (d *Data) GetGitHubData(ctx *fasthttp.RequestCtx) {

}
