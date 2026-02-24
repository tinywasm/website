//go:build wasm

package main

import (
	"github.com/tinywasm/dom"
	"github.com/tinywasm/fmt"
	"github.com/tinywasm/site"
	"github.com/tinywasm/website/web/ui"
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
	return ui.CSS
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
