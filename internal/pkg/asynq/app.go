package asynq

import (
	"context"
	"github.com/arsyadarmawan/asynq-distributed-task/runner"
	"github.com/arsyadarmawan/rest-api/internal/app/book/delivery/worker"
	"github.com/arsyadarmawan/rest-api/internal/app/book/repository/repositoryimpl"
	"github.com/arsyadarmawan/rest-api/internal/app/book/usecase/usecaseimpl"
	"github.com/arsyadarmawan/rest-api/internal/pkg/mongo"
	"github.com/hibiken/asynq"
)

func InitServeMuxAsynq() {
	ctx := context.Background()
	mux := asynq.NewServeMux()

	run := runner.NewAsynqRunner(runner.AsynqRunnerOpts{
		AsynqServer: AsynqServer.Server,
		Mux:         mux,
	})
	bookServeMux(mux)
	if err := run.Run(ctx); err != nil {
		panic("wrong serve mux")
	}
}

func bookServeMux(r *asynq.ServeMux) {
	config := mongo.MongoConfig
	bookRepository := repositoryimpl.NewBookRepository(repositoryimpl.BookRepositoryOpts{
		DB: config,
	})
	bookUsecaseimpl := usecaseimpl.NewBookImpl(usecaseimpl.BookOpts{
		Repository: bookRepository,
	})

	bookRoute := worker.NewBookRegistryWorker(worker.BookRegistryWorkerOpts{
		bookUsecaseimpl,
	})
	bookRoute.RegisterRoutesTo(r)
}
