package service

import "github.com/google/wire"

var WireCommonService = wire.NewSet(
	wire.Struct(new(InviteService), "*"),
)
