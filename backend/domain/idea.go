package domain

import (
	"context"
	"net/http"
)

type Idea struct {
	Id               string `json:"id"`
	UserId           string `json:"user_id"`
	Text             string `json:"text"`
	Feedback         string `json:"feedback"`          // feedback dari reviewer
	ScoreOriginaly   int    `json:"score_originaly"`   // skor orisinal
	ScoreScalability int    `json:"score_scalability"` // skor skalabilitas
	ScoreFeasibility int    `json:"score_feasibility"` //skor kelayakan
	CreatedAt        string `json:"created_at"`        //tanggal pembuatan
}

type IdeaHandler interface {
	SubmitIdea(w http.ResponseWriter, r *http.Request)
}

type IdeaUsecase interface {
	SubmitIdea(ctx context.Context, idea string) ([]*Message, error)
}

type IdeaRepository interface {
	SubmitIdea(ctx context.Context, idea string) error
}
