package middlewares

import (
	"context"
	"github.com/labstack/echo/v4"
)

type (
	I18nMiddlewareInterface interface {
		HandlerError(next echo.HandlerFunc) echo.HandlerFunc
	}

	i18nMiddleware struct{}
)

func NewI18nMiddleware() I18nMiddlewareInterface {
	return &i18nMiddleware{}
}

func (i *i18nMiddleware) HandlerError(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		lang := c.Request().Header.Get("X-Language")

		c.SetRequest(c.Request().WithContext(setDictLangInContext(c.Request().Context(), lang)))
		return next(c)
	}
}

func setDictLangInContext(ctx context.Context, dictLang string) context.Context {
	return context.WithValue(ctx, "lang", dictLang)
}
