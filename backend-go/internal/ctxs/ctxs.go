package ctxs

import "context"

func WithValue(parent context.Context, ctx map[ContextKey]interface{}) context.Context {
	return context.WithValue(parent, theCtx, ctx)
}

func GetContexts(ctx context.Context) map[ContextKey]interface{} {
	if val, ok := ctx.Value(theCtx).(map[ContextKey]interface{}); ok {
		return val
	} else {
		return map[ContextKey]interface{}{}
	}
}
