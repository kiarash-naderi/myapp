package types

const (
	// ModuleName defines the module name
	ModuleName = "myapp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_myapp"
)

var (
	ParamsKey = []byte("p_myapp")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
