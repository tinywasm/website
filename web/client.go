//go:build wasm

package main

import (
	"github.com/tinywasm/dom"
	"github.com/tinywasm/fmt"
	"github.com/tinywasm/site"
)

// UnderConstruction is the single page for tinywasm.app while the site is being built.
type UnderConstruction struct {
	id string
}

func (u *UnderConstruction) HandlerName() string { return "home" }
func (u *UnderConstruction) ModuleTitle() string { return "tinywasm.app" }

func (u *UnderConstruction) GetID() string             { return u.id }
func (u *UnderConstruction) SetID(id string)           { u.id = id }
func (u *UnderConstruction) Children() []dom.Component { return nil }

func (u *UnderConstruction) RenderHTML() string {
	return dom.Main().Class("uc-page").
		Add(
			dom.Div().Class("uc-card").
				Add(
					dom.H1().Class("uc-title").Text("tinywasm.app"),
					dom.P().Class("uc-sub").Text("Under Construction"),
					dom.P().Class("uc-msg").Text("We're building something great. Check back soon."),
				),
		).
		RenderHTML()
}

func (u *UnderConstruction) RenderCSS() string {
	return `
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }

body {
	background: #0d1117;
	color: #e6edf3;
	font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
	min-height: 100vh;
	display: flex;
	align-items: center;
	justify-content: center;
}

.uc-page {
	display: flex;
	align-items: center;
	justify-content: center;
	width: 100%;
	min-height: 100vh;
}

.uc-card {
	text-align: center;
	padding: 3rem 2rem;
	border: 1px solid #30363d;
	border-radius: 12px;
	background: #161b22;
	max-width: 480px;
	width: 90%;
}

.uc-title {
	font-size: 2rem;
	font-weight: 700;
	color: #58a6ff;
	letter-spacing: -0.02em;
	margin-bottom: 0.75rem;
}

.uc-sub {
	font-size: 1.1rem;
	font-weight: 600;
	color: #8b949e;
	margin-bottom: 1rem;
}

.uc-msg {
	font-size: 0.95rem;
	color: #6e7681;
	line-height: 1.6;
}
`
}

func main() {
	if err := site.RegisterHandlers(&UnderConstruction{}); err != nil {
		fmt.Println("site: register error:", err)
		return
	}
	if err := site.Mount("body"); err != nil {
		fmt.Println("site: mount error:", err)
	}
}
