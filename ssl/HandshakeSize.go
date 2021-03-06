package ssl

type HandshakeSize [3]byte

func NewHandshakeSize(num int) HandshakeSize {
	// Attempts to convert an int to an unsigned three-byte slice
	return [3]byte{
		uint8(num >> 16),
		uint8(num >> 8),
		uint8(num)}
}

func (num HandshakeSize) GetValue() uint {
	return uint(num[0] << 16) + uint(num[1] << 8) + uint(num[2])
}

func (num HandshakeSize) GetSize() int {
	return 3
}

func (num HandshakeSize) Serialize() []byte {
	return num[0:3]
}

func (num HandshakeSize) SerializeInto(buf []byte) {
	copy(buf[0:num.GetSize()], num.Serialize())
}

func DeserializeHandshakeSize(buf []byte) (HandshakeSize, int) {
	return HandshakeSize{ buf[0], buf[1], buf[2] }, 3
}