package http

import "github.com/siroj100/hikarie-islamy/internal/usecase"

type (
	EchoHttpHandler struct {
		ucase usecase.IslamyUseCase
	}
)
