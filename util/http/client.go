package http

import (
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

var NoRedirectClient *resty.Client
var RestyClient = NewRestyClient()
var UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
var DefaultTimeout = time.Second * 10

func init() {
	NoRedirectClient = resty.New().SetRedirectPolicy(
		resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}),
	)
	NoRedirectClient.SetHeader("user-agent", UserAgent)
}

func NewRestyClient() *resty.Client {
	return resty.New().
		SetHeader("user-agent", UserAgent).
		SetRetryCount(3).
		SetTimeout(DefaultTimeout)
}
