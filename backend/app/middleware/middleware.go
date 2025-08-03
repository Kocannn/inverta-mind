package middleware

import (
	"github.com/Kocannn/self-dunking-ai/domain"
	"github.com/hammer-code/lms-be/pkg/jwt"
	"go.opentelemetry.io/otel"
)

type Middleware struct {
	Jwt jwt.JWT
	// UserRepo domain.UserRepository
}

var (
	tracer = otel.Tracer("Start Trace")
)

func InitMiddleware(jwt jwt.JWT) domain.Middleware {
	return &Middleware{
		Jwt: jwt,
		// UserRepo: userRepo,
	}
}
