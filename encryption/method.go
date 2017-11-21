package encryption

type EMethod uint8

var (
	None EMethod = 0
	Xor  EMethod = 1 << 1
)
