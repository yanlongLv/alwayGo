package middleware

import "context"

//Handler ...
type Handler func(ctx context.Context, req interface{}) (interface{}, error)

//MiddleWare ...
type MiddleWare func(Handler) Handler

//Chain ...
func Chain(outer MiddleWare, others ...MiddleWare) MiddleWare {
	return func(next Handler) Handler {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}
