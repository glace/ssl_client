package ssl

type CompressionMethods struct { // 1 byte of Length, followed by up to 2^8-1 bytes of data
	length  uint8
	methods []CompressionMethod
}

func NewCompressionMethods(methods []CompressionMethod) CompressionMethods {
	return CompressionMethods{uint8(len(methods)), methods}
}

/*
	Returns the total size in bytes of this struct
*/
func (methods CompressionMethods) GetSize() int {
	return 1 + int(methods.length)
}

func (methods CompressionMethods) SerializeInto(buf []byte) {
	copy(buf[0:1], []byte{methods.length})

	var start = 1

	for _, method := range methods.methods {
		var end = start + 1

		copy(buf[start:end], []byte{byte(method)})

		start = end
	}
}

func (methods CompressionMethods) Serialize() []byte {
	obj := make([]byte, methods.GetSize())
	methods.SerializeInto(obj)
	return obj
}

func DeserializeCompressionMethods(buf []byte) (CompressionMethods, int) {
	var suites []CompressionMethod

	bufferPosition := 0
	methodsLength := uint8(buf[0])

	bufferPosition += 1

	for i := methodsLength; i > 0; i-- {
		suites = append(suites, NewCompressionMethod(buf[bufferPosition]))
		bufferPosition += 1
	}

	return NewCompressionMethods(suites), bufferPosition
}
