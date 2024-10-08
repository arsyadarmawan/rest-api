package worker

import (
	"github.com/hibiken/asynq"
	"rest-api/internal/app/book/usecase"
	"rest-api/internal/pkg/commonval"
)

type BookRegistryWorkerOpts struct {
	Book usecase.Book
}

type BookRegistryWorker struct {
	Opts BookRegistryWorkerOpts
}

func NewBookRegistryWorker(opts BookRegistryWorkerOpts) *BookRegistryWorker {
	return &BookRegistryWorker{Opts: opts}
}

func (b BookRegistryWorker) RegisterRoutesTo(mux *asynq.ServeMux) {
	mux.HandleFunc(commonval.BookWorkerAsynq, MakeDeleteRepository(b.Opts.Book))
}
