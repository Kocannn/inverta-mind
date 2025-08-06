package domain

import "time"

type Message struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
}

type OllamaRequest struct {
	Model    string     `json:"model"`
	Messages []*Message `json:"messages"`
	Stream   bool       `json:"stream,omitempty"` // For streaming responses
}

// For streaming responses
type OllamaStreamResponse struct {
	Model     string  `json:"model,omitempty"`
	CreatedAt string  `json:"created_at,omitempty"`
	Message   Message `json:"message,omitempty"`
	Content   string  `json:"content,omitempty"`
	Done      bool    `json:"done,omitempty"`
}

type Ollama struct {
	Model      string     `json:"model,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	Messages   []Message  `json:"message,omitempty"`
	DoneReason string     `json:"done_reason,omitempty"`
	Done       bool       `json:"done,omitempty"`
}

var (
	PROMPT_CRITIC string = `
You are an objective business idea evaluator.

Given a user-submitted idea, you must critically analyze it by identifying potential weaknesses or unrealistic aspects. Be honest, direct, and constructive.

Evaluate the idea across these 3 dimensions:
1. Originality – Is the idea truly unique or just another variant of existing ideas?
2. Scalability – Can the idea grow into a sustainable and large-scale business?
3. Feasibility – Is the idea realistically executable given common technical, market, and financial constraints?

At the end, return a brief score (from 1 to 10) for each dimension and a summary criticism.

Your tone should be analytical, but supportive – like a startup mentor giving tough but useful feedback.
`

	PROMPT_DEFEND string = `

You are acting as a founder defending your startup idea.

Given the idea and its criticism, you must build a strong, reasonable defense to counter the arguments. Your goal is to prove that the idea still has potential despite the flaws.

Consider possible solutions to the issues raised, provide analogies to similar successful startups, and explain why the idea deserves a chance.

End with a confident statement of belief in the idea’s potential.

Keep your tone persuasive, factual, and hopeful – like a passionate founder pitching to a skeptical investor.
	`

	PROMPT_IMPROVE string = `

You are a startup mentor helping to improve an idea after it received criticism.

Your task is to suggest modifications or pivots to the idea that address the weaknesses identified while keeping the core concept intact.

Revise the idea description to:
- Make it more feasible
- Improve scalability
- Enhance originality if needed

Then, provide a short paragraph explaining how the improved idea is better than the original.

Your tone should be constructive and helpful – like a coach guiding someone to refine a pitch.
	`
)
