package handlers

import (
	"github.com/ErezLevip/fx-medium-example/cache"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
	"go.uber.org/zap"
)

type MeaningOfLife struct {
	cache  cache.MeaningOfLifeCache
	logger *zap.Logger
}

func NewMeaningOfLifeHandler(cache cache.MeaningOfLifeCache, logger *zap.Logger) *MeaningOfLife {
	return &MeaningOfLife{
		cache:  cache,
		logger: logger,
	}
}

func (mol *MeaningOfLife) Handle(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	res, err := mol.cache.LoadOrStore(func() (string, error) {
		return "42", nil
	})
	if err != nil {
		ctx.Error(err.Error(), 500)
		return
	}
	ctx.SetBody([]byte(res))
}
