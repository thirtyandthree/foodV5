package app

import "github.com/google/wire"

var WireApp = wire.NewSet(
	wire.Struct(new(Application), "*"),
)
