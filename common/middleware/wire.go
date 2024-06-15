package middleware

import "github.com/google/wire"

var WireMiddleware = wire.NewSet(
	wire.Struct(new(Middleware), "*"),
)
