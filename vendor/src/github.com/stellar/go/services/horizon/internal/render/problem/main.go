package problem

import (
	"net/http"

	"github.com/stellar/go/services/horizon/internal/context/requestid"
	"github.com/stellar/go/support/render/problem"
	"golang.org/x/net/context"
)

// Inflate expands a problem with contextal information, including setting basic info.
// At present it adds the request's id as the problem's Instance, if available.
func Inflate(ctx context.Context, p *problem.P) {
	problem.Inflate(p)
	p.Instance = requestid.FromContext(ctx)
}

// Well-known and reused problems below:
var (
	// RateLimitExceeded is a well-known problem type.  Use it as a shortcut
	// in your actions.
	RateLimitExceeded = problem.P{
		Type:   "rate_limit_exceeded",
		Title:  "Rate limit exceeded",
		Status: 429,
		Detail: "The rate limit for the requesting IP address is over its alloted " +
			"limit.  The allowed limit and requests left per time period are " +
			"communicated to clients via the http response headers 'X-RateLimit-*' " +
			"headers.",
	}

	// NotImplemented is a well-known problem type.  Use it as a shortcut
	// in your actions.
	NotImplemented = problem.P{
		Type:   "not_implemented",
		Title:  "Resource Not Yet Implemented",
		Status: http.StatusNotFound,
		Detail: "While the requested URL is expected to eventually point to a " +
			"valid resource, the work to implement the resource has not yet " +
			"been completed.",
	}

	// NotAcceptable is a well-known problem type.  Use it as a shortcut
	// in your actions.
	NotAcceptable = problem.P{
		Type: "not_acceptable",
		Title: "An acceptable response content-type could not be provided for " +
			"this request",
		Status: http.StatusNotAcceptable,
	}

	// ServerOverCapacity is a well-known problem type.  Use it as a shortcut
	// in your actions.
	ServerOverCapacity = problem.P{
		Type:   "server_over_capacity",
		Title:  "Server Over Capacity",
		Status: http.StatusServiceUnavailable,
		Detail: "This horizon server is currently overloaded.  Please wait for " +
			"several minutes before trying your request again.",
	}

	// Timeout is a well-known problem type.  Use it as a shortcut
	// in your actions.
	Timeout = problem.P{
		Type:   "timeout",
		Title:  "Timeout",
		Status: http.StatusGatewayTimeout,
		Detail: "Your request timed out before completing.  Please try your " +
			"request again.",
	}

	// UnsupportedMediaType is a well-known problem type.  Use it as a shortcut
	// in your actions.
	UnsupportedMediaType = problem.P{
		Type:   "unsupported_media_type",
		Title:  "Unsupported Media Type",
		Status: http.StatusUnsupportedMediaType,
		Detail: "The request has an unsupported content type. Presently, the " +
			"only supported content type is application/x-www-form-urlencoded.",
	}

	// BeforeHistory is a well-known problem type.  Use it as a shortcut
	// in your actions.
	BeforeHistory = problem.P{
		Type:   "before_history",
		Title:  "Data Requested Is Before Recorded History",
		Status: http.StatusGone,
		Detail: "This horizon instance is configured to only track a " +
			"portion of the stellar network's latest history. This request " +
			"is asking for results prior to the recorded history known to " +
			"this horizon instance.",
	}

	// StaleHistory is a well-known problem type.  Use it as a shortcut
	// in your actions.
	StaleHistory = problem.P{
		Type:   "stale_history",
		Title:  "Historical DB Is Too Stale",
		Status: http.StatusServiceUnavailable,
		Detail: "This horizon instance is configured to reject client requests " +
			"when it can determine that the history database is lagging too far " +
			"behind the connected instance of stellar-core.  If you operate this " +
			"server, please ensure that the ingestion system is properly running.",
	}
)
