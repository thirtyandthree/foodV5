package repo

import "github.com/google/wire"

var WireCommonRepo = wire.NewSet(
	wire.Struct(new(Repo), "*"),
	wire.Struct(new(UserRepo), "*"),
	wire.Struct(new(ConfigRepo), "*"),
	wire.Struct(new(IntegralRepo), "*"),
	wire.Struct(new(InviteRepo), "*"),
)
