package stockproducts

import "context"

type Usecase struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewUsecase(ctx context.Context) *Usecase {
	cctx, cancel := context.WithCancel(ctx)
	return &Usecase{
		ctx:    cctx,
		cancel: cancel,
	}
}
