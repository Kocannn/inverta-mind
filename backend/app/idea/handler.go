package idea

import (
	"github.com/Kocannn/self-dunking-ai/domain"
	"net/http"
)

type (
	handler struct {
		usecase domain.IdeaUsecase
	}
)

// SubmitIdea implements domain.IdeaHandler.
func (h *handler) SubmitIdea(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

var (
	handlr *handler
)

func NewIdeaHandler(usecase domain.IdeaUsecase) domain.IdeaHandler {
	if handlr == nil {
		handlr = &handler{
			usecase,
		}
	}
	return handlr
}
