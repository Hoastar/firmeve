package http

import (
	"github.com/firmeve/firmeve/context"
	"github.com/firmeve/firmeve/kernel/contract"
	logging "github.com/firmeve/firmeve/logger"
	testing2 "github.com/firmeve/firmeve/testing"
	"github.com/kataras/iris/core/errors"
	"net/http"
	"testing"
)

func TestRecovery(t *testing.T) {
	firmeve := testing2.ApplicationDefault(new(logging.Provider))
	req := testing2.NewMockRequest(http.MethodPost, "/?query=queryValue", "").Request
	req.Header.Set(`Content-Type`, contract.HttpMimePlain)
	req.ParseMultipartForm(32 << 20)

	c := context.NewContext(
		firmeve,
		NewHttp(firmeve, req, testing2.NewMockResponseWriter().(contract.HttpWrapResponseWriter)),
		func(c contract.Context) {
			panic(errors.New(`testing error`))
		},
	)

	Recovery(c)
}
