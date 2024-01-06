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

func (s *Schedule) Get(ctx *fasthttp.RequestCtx) {
	data, err := s.ScheduleOps.List(ctx.Request.Body())
	if err != nil || data == nil {
		ctx.Err()
		return
	}
	ctx.SetBody(data.([]byte))
	ctx.Done()
}

func (s *Schedule) Create(ctx *fasthttp.RequestCtx) {
	data, err := s.ScheduleOps.Create(ctx.Request.Body())
	if err != nil || data == nil {
		ctx.Err()
		return
	}
	ctx.SetBody(data.([]byte))
	ctx.Done()
}

func (s *Schedule) Update(ctx *fasthttp.RequestCtx) {
	data, err := s.ScheduleOps.Create(ctx.Request.Body())
	if err != nil || data == nil {
		ctx.Err()
		return
	}
	ctx.SetBody(data.([]byte))
	ctx.Done()
}

func (s *Schedule) Delete(ctx *fasthttp.RequestCtx) {
	err := s.ScheduleOps.Delete(ctx.Request.Body())
	if err != nil {
		ctx.Err()
		return
	}
	ctx.Done()
}
