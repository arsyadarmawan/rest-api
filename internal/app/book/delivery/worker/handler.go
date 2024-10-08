package worker

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"rest-api/internal/app/book/usecase"
)

func MakeDeleteRepository(book usecase.Book) func(context.Context, *asynq.Task) error {
	return func(ctx context.Context, t *asynq.Task) error {
		var id string
		if jsonErr := json.Unmarshal(t.Payload(), id); jsonErr != nil {
			return jsonErr
		}
		if err := book.Delete(ctx, id); err != nil {
			return err
		}
		return nil
	}
}
