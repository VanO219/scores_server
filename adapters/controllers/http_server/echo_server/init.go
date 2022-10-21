package echo_server

import (
	"VanO/scores_server/model/scoresd"
	"context"
)

// Handler Структура контроллера
type Handler struct {
	ctx context.Context
	wbs scoresd.ScoresD
}

func (h *Handler) GetStudents() (err error) {
	return
}

func (h *Handler) GetGroups() (err error) {
	return
}
