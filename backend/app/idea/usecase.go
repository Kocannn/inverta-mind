package idea

import (
	"context"

	"github.com/Kocannn/self-dunking-ai/domain"
	pkgDB "github.com/Kocannn/self-dunking-ai/pkg/database"
	"github.com/Kocannn/self-dunking-ai/pkg/ollama"
	"github.com/sirupsen/logrus"
)

type (
	usecase struct {
		dbTx pkgDB.DatabaseTransaction
		repo domain.IdeaRepository
	}
)

// SubmitIdea implements domain.IdeaUsecase.
func (u *usecase) SubmitIdea(ctx context.Context, idea string) ([]*domain.Message, error) {
	var messages []*domain.Message

	promptSystem := &domain.Message{
		Role:    "system",
		Content: domain.PROMPT_CRITIC,
	}

	promptUser := &domain.Message{
		Role:    "user",
		Content: idea,
	}

	messages = append(messages, promptSystem, promptUser)

	response, err := ollama.PostPrompt(messages)
	if err != nil {
		logrus.Errorf("error posting prompt: %v", err)
		return nil, err // Return error to caller
	}

	// Check if response is not nil and has messages
	if response != nil && len(response.Messages) > 0 {
		assistantMessage := &domain.Message{
			Role:    "assistant",
			Content: response.Messages[0].Content,
		}
		messages = append(messages, assistantMessage)
	} else {
		// Handle empty response case
		assistantMessage := &domain.Message{
			Role:    "assistant",
			Content: "I'm sorry, I couldn't process your idea at this time.",
		}
		messages = append(messages, assistantMessage)
	}

	return messages, nil
}

var (
	uc *usecase
)

func NewIdeaUsecase(dbTx pkgDB.DatabaseTransaction, repo domain.IdeaRepository) domain.IdeaUsecase {
	if uc == nil {
		uc = &usecase{
			dbTx: dbTx,
			repo: repo,
		}
	}
	return uc
}
