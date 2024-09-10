package keeper

import (
	"github.com/kiarash-naderi/myapp/x/abslayer/types"
)

var _ types.QueryServer = Keeper{}
