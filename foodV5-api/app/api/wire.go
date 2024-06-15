package api

import "github.com/google/wire"

var WireController = wire.NewSet(
	wire.Struct(new(AccountApi), "*"),
)
