# Task: Extract CSS to web/ui/ssr.go

## Objective

Move the inline CSS currently in `web/client.go` (`RenderCSS()` method) into a
dedicated file `web/ui/ssr.go`, then update `client.go` to reference it.

## Current State

`web/client.go` contains an `UnderConstruction` component whose `RenderCSS()`
method returns a long inline CSS string literal. The HTML and CSS live in the
same file, which violates Single Responsibility.

## Required Changes

### 1. Create `web/ui/ssr.go`

- Package: `package ui`
- Export a single string constant `CSS` with the exact CSS content currently
  returned by `RenderCSS()` in `web/client.go` (do not modify the CSS rules)

```go
package ui

// CSS is the stylesheet for the tinywasm.app website.
const CSS = `...` // move the exact content from RenderCSS()
```

### 2. Update `web/client.go`

- Import `github.com/tinywasm/website/web/ui`
- Change `RenderCSS()` to return `ui.CSS`

```go
import "github.com/tinywasm/website/web/ui"

func (u *UnderConstruction) RenderCSS() string {
    return ui.CSS
}
```

## Constraints

- Do NOT modify the CSS rules — copy them verbatim
- Do NOT change the HTML structure in `RenderHTML()`
- The build tag `//go:build wasm` stays only in `client.go`
- `web/ui/ssr.go` has NO build tag (it is plain Go, usable from any context)
- Keep `web/client.go` as the entry point (`package main`, `func main()`)

## Verification

```
cd web && go vet ./...   # no errors
```
The component must still satisfy the `dom.Component` interface after the change.
