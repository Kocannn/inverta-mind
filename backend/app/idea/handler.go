package idea

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Kocannn/self-dunking-ai/domain"
	"github.com/Kocannn/self-dunking-ai/pkg/ollama"
	"github.com/Kocannn/self-dunking-ai/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type (
	handler struct {
		usecase domain.IdeaUsecase
	}
)

// GetIdeas implements domain.IdeaHandler.
func (h *handler) GetIdea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("error converting id to int: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
			Data:    nil,
		}, w)
		return
	}

	data, err := h.usecase.GetIdea(r.Context(), idInt)

	utils.Response(domain.HttpResponse{
		Code:    http.StatusOK,
		Message: "Idea retrieved successfully",
		Data:    data,
	}, w)

}

// SubmitIdeaStream implements domain.IdeaHandler.
func (h *handler) SubmitIdeaStream(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error reading request body: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Error reading request body",
		}, w)
		return
	}

	dataBuffer := domain.SubmitIdeaRequest{}

	if err := json.Unmarshal(bodyBytes, &dataBuffer); err != nil {
		logrus.Errorf("error unmarshalling request body: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusBadRequest,
			Message: "Error parsing request body",
			Data:    nil,
		}, w)
		return
	}

	logrus.Infof("Received idea submission: %s", dataBuffer.Idea)

	now := time.Now()
	submitRequest := domain.SubmitIdeaRequest{
		Idea:      dataBuffer.Idea,
		CreatedAt: &now,
	}

	createdIdea, err := h.usecase.SubmitIdeaStream(r.Context(), submitRequest)
	if err != nil {
		logrus.Errorf("error saving idea stream: %v", err)
		utils.Response(domain.HttpResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error saving idea",
			Data:    nil,
		}, w)
		return
	}

	utils.Response(domain.HttpResponse{
		Code:    http.StatusOK,
		Message: "Idea submitted successfully",
		Data:    createdIdea,
	}, w)
}

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

// Streaming versions of handlers
func (h *handler) StreamSubmitIdea(w http.ResponseWriter, r *http.Request) {
	// Set proper CORS headers for SSE
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle OPTIONS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Errorf("error parsing ID: %v", err)
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	idea, err := h.usecase.GetIdea(r.Context(), idInt)
	if err != nil {
		logrus.Errorf("error getting idea: %v", err)
		http.Error(w, "Error retrieving idea", http.StatusInternalServerError)
		return
	}

	promptSystem := &domain.Message{
		Role:    "system",
		Content: domain.PROMPT_CRITIC,
	}

	promptUser := &domain.Message{
		Role:    "user",
		Content: idea.Idea,
	}

	messages := []*domain.Message{promptSystem, promptUser}

	// Use streaming response
	if err := ollama.StreamPrompt(w, messages); err != nil {
		logrus.Errorf("error streaming idea: %v", err)
		return
	}
}

func (h *handler) StreamDefendIdea(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error reading request body: %v", err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	dataBuffer := domain.Idea{}
	if err := json.Unmarshal(bodyBytes, &dataBuffer); err != nil {
		logrus.Errorf("error unmarshalling request body: %v", err)
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	promptSystem := &domain.Message{
		Role:    "system",
		Content: domain.PROMPT_DEFEND,
	}

	promptUser := &domain.Message{
		Role:    "user",
		Content: dataBuffer.Critique,
	}

	messages := []*domain.Message{promptSystem, promptUser}

	// Use streaming response
	if err := ollama.StreamPrompt(w, messages); err != nil {
		logrus.Errorf("error streaming idea: %v", err)
		return
	}
}

func (h *handler) StreamImproveIdea(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		logrus.Errorf("error reading request body: %v", err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	dataBuffer := domain.Idea{}
	if err := json.Unmarshal(bodyBytes, &dataBuffer); err != nil {
		logrus.Errorf("error unmarshalling request body: %v", err)
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	promptSystem := &domain.Message{
		Role:    "system",
		Content: domain.PROMPT_IMPROVE,
	}

	promptUser := &domain.Message{
		Role:    "user",
		Content: dataBuffer.Critique,
	}

	messages := []*domain.Message{promptSystem, promptUser}

	// Use streaming response
	if err := ollama.StreamPrompt(w, messages); err != nil {
		logrus.Errorf("error streaming idea: %v", err)
		return
	}
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
