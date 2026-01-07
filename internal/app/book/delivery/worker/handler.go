package worker

import (
	"context"
	"encoding/json"
	error2 "github.com/arsyadarmawan/rest-api/internal/pkg/error"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"

	//"github.com/sirupsen/logrus"
	"github.com/arsyadarmawan/rest-api/internal/app/book/usecase"
)

func MakeDeleteRepository(book usecase.Book) func(context.Context, *asynq.Task) *error2.AppError {
	return func(ctx context.Context, t *asynq.Task) *error2.AppError {
		var id string
		if jsonErr := json.Unmarshal(t.Payload(), &id); jsonErr != nil {
			return error2.BadRequestError()
		}
		err := book.Update(ctx, id)
		if err != nil {
			logrus.WithContext(ctx).Errorf("error deleting book: %v", err)
			return err
		}

		return nil
	}
}
