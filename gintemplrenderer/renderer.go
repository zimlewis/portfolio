package gintemplrenderer

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

type Renderer struct {
	Ctx context.Context
	Status int
	Component templ.Component
}

func New(c context.Context, status int, component templ.Component) *Renderer {
	return &Renderer{
		Ctx: c,
		Status: status,
		Component: component,
	}
}


func (renderer Renderer) WriteContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}

func (renderer Renderer) Render(w http.ResponseWriter) error {
	renderer.WriteContentType(w)

	if renderer.Status != -1 {
		w.WriteHeader(renderer.Status)
	}

	if renderer.Component != nil {
		return renderer.Component.Render(renderer.Ctx, w)
	}

	return nil
}
