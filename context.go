package grapher

import (
	"context"
	"net/http"
)

type reqKey struct{}

func ContextWithRequest(ctx context.Context, req *http.Request) context.Context {
	return context.WithValue(ctx, reqKey{}, req)
}

func RequestFromContext(ctx context.Context) *http.Request {
	return ctx.Value(reqKey{}).(*http.Request)
}
