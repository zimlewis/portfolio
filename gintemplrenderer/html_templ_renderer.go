package gintemplrenderer

import (
	"context"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin/render"
)

type HTMLTemplRenderer struct {
	FallbackHTMLRenderer render.HTMLRender
}

func (r *HTMLTemplRenderer) Instance(s string, d any) render.Render{
	templData, ok := d.(templ.Component)
	if !ok {
		if r.FallbackHTMLRenderer != nil { return r.FallbackHTMLRenderer.Instance(s, d) }
	}


	return &Renderer {
		Ctx: context.Background(),
		Status: -1,
		Component: templData,
	}
}
