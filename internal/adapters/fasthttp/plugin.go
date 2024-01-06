package fasthttp

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"github.com/valyala/fasthttp"
)

type Plugins struct {
	ports.PluginOps
	logger.Logger
}

func (p *Plugins) Get(ctx *fasthttp.RequestCtx) {
	data, err := p.PluginOps.List()
	if err != nil || data == nil {
		ctx.Err()
		return
	}
	ctx.SetBody(data.([]byte))
	ctx.Done()
}

func (p *Plugins) Update(ctx *fasthttp.RequestCtx) {
	err := p.PluginOps.Update(ctx.Request.Body())
	if err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
}

func (p *Plugins) Install(ctx *fasthttp.RequestCtx) {
	err := p.PluginOps.Install(ctx.Request.Body())
	if err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
}

func (p *Plugins) Trigger(ctx *fasthttp.RequestCtx) {
	err := p.PluginOps.Trigger(ctx.Request.Body())
	if err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
}

func (p *Plugins) Delete(ctx *fasthttp.RequestCtx) {
	err := p.PluginOps.Trigger(ctx.Request.Body())
	if err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
}
