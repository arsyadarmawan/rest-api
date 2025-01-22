package worker

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/sirupsen/logrus"

	//"github.com/sirupsen/logrus"
	"rest-api/internal/app/book/usecase"
)

func MakeDeleteRepository(book usecase.Book) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {
		var id string
		if jsonErr := json.Unmarshal(t.Payload(), &id); jsonErr != nil {
			return jsonErr
		}
		err := book.Update(ctx, id)
		if err != nil {
			logrus.WithContext(ctx).Errorf("error deleting book: %v", err)
			return err
		}

		return nil
	}
}
