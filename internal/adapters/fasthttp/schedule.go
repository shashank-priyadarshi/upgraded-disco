package fasthttp

import (
	"github.com/shashank-priyadarshi/upgraded-disco/internal/ports"
	"github.com/shashank-priyadarshi/upgraded-disco/utils/logger"
	"github.com/valyala/fasthttp"
)

type Schedule struct {
	ports.ScheduleOps
	logger.Logger
}

// TODO
func (s *Schedule) List(ctx *fasthttp.RequestCtx) {

}

func (s *Schedule) Create(ctx *fasthttp.RequestCtx) {

}

func (s *Schedule) Delete(ctx *fasthttp.RequestCtx) {

}
