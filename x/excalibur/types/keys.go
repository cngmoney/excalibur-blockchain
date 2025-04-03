package types

const (
	// ModuleName defines the module name
	ModuleName = "excalibur"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_excalibur"
)

var (
	ParamsKey = []byte("p_excalibur")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
