package http

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/siroj100/hikarie-islamy/internal/ctxs"
	"github.com/siroj100/hikarie-islamy/internal/usecase"
)

func ContextEcho(c echo.Context) context.Context {
	return ctxs.WithValue(c.Request().Context(), map[ctxs.ContextKey]interface{}{
		//ctxs.JwtUser: c.Get(ctxs.EchoUserToken),
		//ctxs.AppUser: c.Get(ctxs.EchoDbAppUser),
	})
}

func NewEchoHttpHandler(ucase usecase.IslamyUseCase) EchoHttpHandler {
	return EchoHttpHandler{ucase: ucase}
}
