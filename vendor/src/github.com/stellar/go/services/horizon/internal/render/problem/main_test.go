package problem

import (
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stellar/go/services/horizon/internal/context/requestid"
	"github.com/stellar/go/support/render/problem"
	"golang.org/x/net/context"
)

func TestProblemPackage(t *testing.T) {
	ctx := context.Background()

	testRender := func(ctx context.Context, p interface{}) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		problem.Render(ctx, w, p)
		return w
	}

	Convey("Common Problems", t, func() {
		Convey("NotFound", func() {
			w := testRender(ctx, problem.NotFound)
			So(w.Code, ShouldEqual, 404)
			t.Log(w.Body.String())
		})

		Convey("RateLimitExceeded", func() {
			w := testRender(ctx, RateLimitExceeded)
			So(w.Code, ShouldEqual, 429)
			t.Log(w.Body.String())
		})
	})

	Convey("problem.Inflate", t, func() {
		Convey("sets Instance to the request id based upon the context", func() {
			ctx2 := requestid.Context(ctx, "2")
			p := problem.P{}
			Inflate(ctx2, &p)

			So(p.Instance, ShouldEqual, "2")

			// when no request id is set, instance should be ""
			Inflate(ctx, &p)
			So(p.Instance, ShouldEqual, "")
		})
	})
}
