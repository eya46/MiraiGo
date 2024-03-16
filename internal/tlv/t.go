package tlv

import "github.com/eya46/MiraiGo/binary"

func T(tag uint16, value []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(tag)
		w.WriteBytesShort(value)
	})
}
