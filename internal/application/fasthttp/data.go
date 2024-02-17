package fasthttp

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	logger "github.com/shashank-priyadarshi/utilities/logger/ports"
	"github.com/valyala/fasthttp"
)

type Data struct {
	ports.DataOps
	logger.Logger
}

func (d *Data) Chess(ctx *fasthttp.RequestCtx) {

	data, err := d.DataOps.Chess()
	if err != nil || data == nil {
		ctx.Err()
		return
	}

	ctx.SetBody(data.([]byte))
	ctx.Done()
}

func (d *Data) GitHub(ctx *fasthttp.RequestCtx) {

	data, err := d.DataOps.GitHub()
	if err != nil || data == nil {
		ctx.Err()
		return
	}

	ctx.SetBody(data.([]byte))
	ctx.Done()
}
