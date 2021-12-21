package cryptopals

import (
	"encoding/hex"
	"testing"
)

func TestEncodeB64(t *testing.T) {
	inputString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	inputBytes, err := hex.DecodeString(inputString)
	if err != nil {
		t.Error("Hex decoding error")
	}

	encodedString := EncodeB64(inputBytes)
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	// fmt.Println(encodedString)

	if encodedString != want {
		t.Errorf("got %q want %q", encodedString, want)
	}

	inputString = "6b617261626f6761"
	inputBytes, err = hex.DecodeString(string(inputString))

	if err != nil {
		t.Error("Hex decoding error")
	}

	encodedString = EncodeB64(inputBytes)
	want = "a2FyYWJvZ2E="
	// fmt.Println(encodedString)

	if encodedString != want {
		t.Errorf("got %q want %q", encodedString, want)
	}

	inputString = "6b617261"
	inputBytes, err = hex.DecodeString(string(inputString))

	if err != nil {
		t.Error("Hex decoding error")
	}

	encodedString = EncodeB64(inputBytes)
	want = "a2FyYQ=="
	//fmt.Println(encodedString)

	if encodedString != want {
		t.Errorf("got %q want %q", encodedString, want)
	}
}
