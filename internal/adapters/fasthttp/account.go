package fasthttp

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"github.com/valyala/fasthttp"
)

type Account struct {
	ports.AccountOps
	logger.Logger
}

func (a *Account) Register(ctx *fasthttp.RequestCtx) {
	if err := a.AccountOps.RegisterUser(ctx.Request.Body()); err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
}

func (a *Account) Login(ctx *fasthttp.RequestCtx) {
	var data interface{}
	var err error
	if data, err = a.AccountOps.Login(ctx.Request.Body()); err != nil || data == nil {
		ctx.Err()
		return
	}
	ctx.SetBody(data.([]byte))
	ctx.Done()
}

func (a *Account) Update(ctx *fasthttp.RequestCtx) {
	if err := a.AccountOps.ResetPassword(ctx.Request.Body()); err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
}

func (a *Account) Delete(ctx *fasthttp.RequestCtx) {
	if err := a.AccountOps.DeleteUser(ctx.Request.Body()); err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
}
