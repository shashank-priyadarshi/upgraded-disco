package fasthttp

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"github.com/valyala/fasthttp"
)

type Plugins struct {
	ports.PluginOps
	log logger.Logger
}

func (p *Plugins) Install(ctx *fasthttp.RequestCtx) {
	err := p.PluginOps.Install(ctx.Request.Body())
	if err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
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
	err := p.PluginOps.Upgrade(ctx.Request.Body())
	if err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
}

func (p *Plugins) Trigger(ctx *fasthttp.RequestCtx) {
	var pluginID string
	pluginID = getPathParts(string(ctx.URI().Path()))[2]

	p.log.Infof("PluginID: %s", pluginID)
	err := p.PluginOps.Trigger(pluginID)
	if err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
}

func (p *Plugins) Delete(ctx *fasthttp.RequestCtx) {
	var pluginID string
	pluginID = getPathParts(string(ctx.URI().Path()))[2]
	err := p.PluginOps.Uninstall(pluginID)
	if err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
}
