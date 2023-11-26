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
	data, err := d.DataOps.GetGraphData()
	if err != nil || data == nil {
		ctx.Err()
		return
	}
	ctx.SetBody(data.([]byte))
	ctx.Done()
}

func (d *Data) GetGitHubData(ctx *fasthttp.RequestCtx) {
	data, err := d.DataOps.GetGitHubData()
	if err != nil || data == nil {
		ctx.Err()
		return
	}
	ctx.SetBody(data.([]byte))
	ctx.Done()
}
