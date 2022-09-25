package main

import (
	"context"
	"fmt"

	l "github.com/SamHennessy/hlive"
	"github.com/SamHennessy/hlive/hlivekit"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func app(ctxWails context.Context) *l.Page {
	env := runtime.Environment(ctxWails)
	assetsPath := "assets"
	dev := env.BuildType == "dev"
	if dev {
		assetsPath = "./src/assets"
	}

	page := l.NewPage()
	page.DOM().Head().Add(l.T("script", l.Attrs{"src": "./src/index.js", "type": "module"}))
	if !dev {
		page.DOM().Head().Add(l.T("link", l.Attrs{"rel": "stylesheet", "href": "/assets/css/index.css"}))
	}

	inputVal := l.Box("")
	greetMsg := l.Box("Please enter your name below 👇")

	page.DOM().Body().Add(
		l.T("div", l.Attrs{"id": "app"},
			l.T("img", l.Attrs{"id": "logo", "src": assetsPath + "/images/logo-universal.png"}, l.Class("logo")),
			l.T("div", l.Attrs{"id": "result"}, l.Class("result"), greetMsg),
			l.T("div", l.Attrs{"id": "input"}, l.Class("input-box"),
				l.C("form", l.PreventDefault(),
					l.On("submit", func(_ context.Context, _ l.Event) {
						greetMsg.Set(fmt.Sprintf("Hello %s, It's show time!", inputVal.Get()))
					}),
					l.C("input", l.Attrs{"id": "name", "type": "text", "autocomplete": "off"}, l.Class("input"),
						hlivekit.Focus(),
						l.OnOnce("focus", func(ctx context.Context, e l.Event) {
							c, ok := e.Binding.Component.(l.Adder)
							if ok {
								hlivekit.FocusRemove(c)
							}
						}),
						l.On("input", func(_ context.Context, e l.Event) {
							inputVal.Set(e.Value)
						}),
					),
					l.T("button", l.Class("btn"), "Greet"),
				),
			),
		),
	)

	return page
}
