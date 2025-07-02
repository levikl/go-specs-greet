package webserver

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"github.com/levikl/go-specs-greet/adapters/webserver/internal/pages"
)

type Driver struct {
	baseURL string
	browser *rod.Browser
}

func NewDriver(baseURL string) (*Driver, func() error) {
	l := launcher.MustNewManaged("")
	browser := rod.New().Client(l.MustClient()).MustConnect()
	return &Driver{baseURL: baseURL, browser: browser}, browser.Close
}

func (d Driver) Greet(name string) (string, error) {
	var (
		page      = d.browser.MustPage(d.baseURL)
		replyPage = pages.Reply{Page: page}
		formPage  = pages.Form{Page: page}
	)

	if err := formPage.Greet(name); err != nil {
		return "", err
	}

	return replyPage.ReadReply()
}

func (d Driver) Curse(name string) (string, error) {
	var (
		page      = d.browser.MustPage(d.baseURL)
		replyPage = pages.Reply{Page: page}
		formPage  = pages.Form{Page: page}
	)

	if err := formPage.Curse(name); err != nil {
		return "", err
	}

	return replyPage.ReadReply()
}
