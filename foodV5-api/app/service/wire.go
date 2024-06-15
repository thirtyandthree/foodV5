package service

import "github.com/google/wire"

var WireService = wire.NewSet(
	wire.Struct(new(AccountService), "*"),
)
