package worker

import (
	"github.com/arsyadarmawan/rest-api/internal/app/book/usecase"
	"github.com/arsyadarmawan/rest-api/internal/pkg/commonval"
	"github.com/hibiken/asynq"
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
