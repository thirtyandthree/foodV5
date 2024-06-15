package router

import "github.com/google/wire"

var WireRouter = wire.NewSet(
	wire.Struct(new(Router), "*"),
)
