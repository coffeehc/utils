// binary project binary.go
package binary

import "encoding/binary"
import "bytes"

func IntToBytes(i interface{}) []byte {
	var b bytes.Buffer
	binary.Write(&b, binary.BigEndian, i)
	return b.Bytes()
}

func BytesToInt16(v []byte) uint16 {
	return binary.BigEndian.Uint16(v)
}

func BytesToInt32(v []byte) uint32 {
	return binary.BigEndian.Uint32(v)
}

func BytesToInt64(v []byte) uint64 {
	return binary.BigEndian.Uint64(v)
}
