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

// TODO
func (a *Account) RegisterUser(ctx *fasthttp.RequestCtx) {

}

func (a *Account) Login(ctx *fasthttp.RequestCtx) {

}

func (a *Account) ResetPassword(ctx *fasthttp.RequestCtx) {

}

func (a *Account) DeleteUser(ctx *fasthttp.RequestCtx) {

}
