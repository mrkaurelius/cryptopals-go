package cryptopals

import (
	"encoding/hex"
	"fmt"
	"os"
)

func XorString(inp1, inp2 *string) []byte {
	inp1Bytes, err := hex.DecodeString(*inp1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}

	inp2Bytes, err := hex.DecodeString(*inp2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
	}

	return XorBuffer(inp1Bytes, inp2Bytes)
}

func XorBuffer(buffer1, buffer2 []byte) []byte {
	xorBuffer := make([]byte, len(buffer1))
	for i, v := range buffer1 {
		xorBuffer[i] = v ^ buffer2[i]
	}
	return xorBuffer
}

func XorSingleByte(buffer []byte, key uint8) []byte {
	xoredBytes := make([]byte, len(buffer))
	for i, _ := range buffer {
		xoredBytes[i] = key ^ uint8(buffer[i])
	}
	return xoredBytes
}
