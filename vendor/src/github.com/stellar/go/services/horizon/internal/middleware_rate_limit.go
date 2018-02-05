package horizon

import (
	"net/http"

	"github.com/zenazn/goji/web"
)

func (web *Web) RateLimitMiddleware(c *web.C, next http.Handler) http.Handler {
	return web.rateLimiter.Throttle(next)
}
