package domain

import (
	"context"
	"net/http"
	"time"
)

type Idea struct {
	Id               string `json:"id"`
	UserId           string `json:"user_id"`
	Text             string `json:"text"`
	Critique         string `json:"critique"`          // Add this field for the critique
	Feedback         string `json:"feedback"`          // feedback dari reviewer
	ScoreOriginaly   int    `json:"score_originaly"`   // skor orisinal
	ScoreScalability int    `json:"score_scalability"` // skor skalabilitas
	ScoreFeasibility int    `json:"score_feasibility"` //skor kelayakan
	CreatedAt        string `json:"created_at"`        //tanggal pembuatan
}

type SubmitIdeaRequest struct {
	Id        int        `json:"id" gorm:"primary_key auto_increment"`
	Idea      string     `json:"idea"`
	CreatedAt *time.Time `json:"created_at" gorm:"not null" default:"CURRENT_TIMESTAMP"`
}

type IdeaHandler interface {
	SubmitIdea(w http.ResponseWriter, r *http.Request)
	StreamSubmitIdea(w http.ResponseWriter, r *http.Request)
	StreamDefendIdea(w http.ResponseWriter, r *http.Request)
	StreamImproveIdea(w http.ResponseWriter, r *http.Request)
	DefendIdea(w http.ResponseWriter, r *http.Request)
	ImproveIdea(w http.ResponseWriter, r *http.Request)
	SubmitIdeaStream(w http.ResponseWriter, r *http.Request)
}

type IdeaUsecase interface {
	SubmitIdea(ctx context.Context, idea string) ([]*Message, error)
	DefendIdea(ctx context.Context, critique string) ([]*Message, error)
	ImproveIdea(ctx context.Context, critique string) ([]*Message, error)
	SubmitIdeaStream(ctx context.Context, idea SubmitIdeaRequest) (SubmitIdeaRequest, error)
}

type IdeaRepository interface {
	SubmitIdea(ctx context.Context, idea string) error
	SubmitIdeaStream(ctx context.Context, idea SubmitIdeaRequest) (SubmitIdeaRequest, error)
}
