package asynq

import (
	"context"
	"github.com/arsyadarmawan/asynq-distributed-task/runner"
	"github.com/hibiken/asynq"
	"rest-api/internal/app/book/delivery/worker"
	"rest-api/internal/app/book/repository/repositoryimpl"
	"rest-api/internal/app/book/usecase/usecaseimpl"
	"rest-api/internal/pkg/mongo"
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
