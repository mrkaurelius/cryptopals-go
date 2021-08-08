package base64

import "strings"

var base64Table string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func EncodeB64(bytes []byte) string {
	// TODO error handling
	var sb strings.Builder
	i := 0

	for i = 0; i < len(bytes)-2; i += 3 {
		b64_1 := uint(bytes[i] >> 2)
		sb.WriteByte(base64Table[b64_1])
		b64_2 := ((bytes[i] & 0b00000011) << 4) ^ ((bytes[i+1] & 0b11110000) >> 4)
		sb.WriteByte(base64Table[b64_2])
		b64_3 := ((bytes[i+1] & 0b00001111) << 2) ^ ((bytes[i+2] & 0b11000000) >> 6)
		sb.WriteByte(base64Table[b64_3])
		b64_4 := (bytes[i+2] & 0b00111111)
		sb.WriteByte(base64Table[b64_4])
	}

	if len(bytes)%3 == 0 {
		return sb.String()
	}

	if len(bytes)%3 == 1 {
		b64_1 := uint(bytes[i] >> 2)
		sb.WriteByte(base64Table[b64_1])
		b64_2 := ((bytes[i] & 0b00000011) << 4) ^ ((0 & 0b11110000) >> 4)
		sb.WriteByte(base64Table[b64_2])
		sb.WriteString("==")
		return sb.String()
	}

	if len(bytes)%3 == 2 {
		b64_1 := uint(bytes[i] >> 2)
		sb.WriteByte(base64Table[b64_1])
		b64_2 := ((bytes[i] & 0b00000011) << 4) ^ ((bytes[i+1] & 0b11110000) >> 4)
		sb.WriteByte(base64Table[b64_2])
		b64_3 := ((bytes[i+1] & 0b00001111) << 2) ^ ((0 & 0b11000000) >> 6)
		sb.WriteByte(base64Table[b64_3])
		sb.WriteString("=")
		return sb.String()
	}

	// bad
	return ""
}

// TODO
func DecodeB64(base64String *string) []byte {
	return make([]byte, 0)
}
