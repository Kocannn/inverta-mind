package idea

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Kocannn/self-dunking-ai/domain"
	"github.com/Kocannn/self-dunking-ai/utils"
	"github.com/sirupsen/logrus"
)

type (
	handler struct {
		usecase domain.IdeaUsecase
	}
)

// DefendIdea implements domain.IdeaHandler.
func (h *handler) DefendIdea(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error reading request body: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Error reading request body",
			Data:    nil,
		}, w)
		return
	}

	dataBuffer := domain.Idea{}

	if err := json.Unmarshal(bodyBytes, &dataBuffer); err != nil {
		logrus.Errorf("error unmarshalling request body: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Error parsing request body",
			Data:    nil,
		}, w)
		return
	}

	messages, err := h.usecase.DefendIdea(r.Context(), dataBuffer.Critique)
	if err != nil {
		logrus.Errorf("error submitting idea: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error processing idea",
			Data:    nil,
		}, w)
		return
	}

	// Extract only the assistant's message (last message)
	var assistantResponse *domain.Message
	if len(messages) > 0 {
		assistantResponse = messages[len(messages)-1]
	}

	utils.Response(domain.HttpResponse{
		Code:    http.StatusOK,
		Message: "Idea submitted successfully",
		Data:    assistantResponse,
	}, w)
}

// ImproveIdea implements domain.IdeaHandler.
func (h *handler) ImproveIdea(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error reading request body: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Error reading request body",
			Data:    nil,
		}, w)
		return
	}

	dataBuffer := domain.Idea{}

	if err := json.Unmarshal(bodyBytes, &dataBuffer); err != nil {
		logrus.Errorf("error unmarshalling request body: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Error parsing request body",
			Data:    nil,
		}, w)
		return
	}

	messages, err := h.usecase.DefendIdea(r.Context(), dataBuffer.Critique)
	if err != nil {
		logrus.Errorf("error submitting idea: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error processing idea",
			Data:    nil,
		}, w)
		return
	}

	// Extract only the assistant's message (last message)
	var assistantResponse *domain.Message
	if len(messages) > 0 {
		assistantResponse = messages[len(messages)-1]
	}

	utils.Response(domain.HttpResponse{
		Code:    http.StatusOK,
		Message: "Idea submitted successfully",
		Data:    assistantResponse,
	}, w)
}

// SubmitIdea implements domain.IdeaHandler.
func (h *handler) SubmitIdea(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error reading request body: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Error reading request body",
			Data:    nil,
		}, w)
		return
	}

	dataBuffer := domain.Idea{}

	if err := json.Unmarshal(bodyBytes, &dataBuffer); err != nil {
		logrus.Errorf("error unmarshalling request body: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Error parsing request body",
			Data:    nil,
		}, w)
		return
	}

	messages, err := h.usecase.SubmitIdea(r.Context(), dataBuffer.Text)
	if err != nil {
		logrus.Errorf("error submitting idea: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error processing idea",
			Data:    nil,
		}, w)
		return
	}

	// Extract only the assistant's message (last message)
	var assistantResponse *domain.Message
	if len(messages) > 0 {
		assistantResponse = messages[len(messages)-1]
	}

	utils.Response(domain.HttpResponse{
		Code:    http.StatusOK,
		Message: "Idea submitted successfully",
		Data:    assistantResponse,
	}, w)
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
