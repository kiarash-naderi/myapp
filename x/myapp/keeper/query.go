package keeper

import (
	"github.com/kiarash-naderi/myapp/x/myapp/types"
)

var _ types.QueryServer = Keeper{}
