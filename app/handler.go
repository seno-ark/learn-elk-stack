package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := IndexUsecase(ctx, randomInt(0, 11))
	if err != nil {
		status, message := ErrStatusMessage(err)

		JsonResp(w, status, map[string]any{
			"message": message,
		})
		return
	}

	JsonResp(w, http.StatusOK, map[string]any{
		"message": http.StatusText(http.StatusOK),
	})
}

func IndexUsecase(ctx context.Context, num int) error {

	// !(num%10) = panic
	// !(num%9) = error 500
	// !(num%8) = error 404
	// else		= succes 200

	if num%10 == 0 {
		panic("unknown error")

	} else if num%9 == 0 {
		Logger.Error(
			"Failed",
			slog.Any("reqID", ctx.Value(ContextRequestID)),
			slog.Any("err", errors.New("failed to ***")),
		)
		return ErrInternalServer

	} else if num%8 == 0 {
		return ErrNotFound
	}

	Logger.Info(
		"Success",
		"ReqID", ctx.Value(ContextRequestID),
	)
	return nil
}
