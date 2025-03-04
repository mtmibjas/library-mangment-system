package middleware

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

type RequestLimiter struct {
	Limiter *rate.Limiter
}

func NewRequestLimiter(rps int) *RequestLimiter {
	limiter := rate.NewLimiter(rate.Every(time.Second), rps)
	return &RequestLimiter{Limiter: limiter}
}

func (rl *RequestLimiter) RequestLimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !rl.Limiter.Allow() {
			return c.JSON(http.StatusTooManyRequests, map[string]string{
				"error": "Too many requests, please try again later.",
			})
		}
		return next(c)
	}
}
