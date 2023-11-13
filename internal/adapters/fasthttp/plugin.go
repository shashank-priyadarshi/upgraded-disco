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

func (p *Plugins) List(ctx *fasthttp.RequestCtx) {

}

func (p *Plugins) Update(ctx *fasthttp.RequestCtx) {

}

func (p *Plugins) Install(ctx *fasthttp.RequestCtx) {

}

func (p *Plugins) Trigger(ctx *fasthttp.RequestCtx) {

}
