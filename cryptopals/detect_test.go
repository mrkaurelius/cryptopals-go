package cryptopals_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/mrkaurelius/cryptopals-go/v2/cryptopals"
)

func TestDetectSingleByteXor(t *testing.T) {
	cryptopals.DetectSingleByteXor("../assets/4.txt")
}

func TestXorSingleByte(t *testing.T) {
	cipherText := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	cipherTextBytes, _ := hex.DecodeString(cipherText)

	var bestScore float64 = 0.0
	var bestText string

	for i := 65; i < 91; i++ {
		r := rune(i)
		possiblePlainText := cryptopals.XorSingleByte(cipherTextBytes, r)
		fmt.Println(string(possiblePlainText))

		score := cryptopals.Englishness(string(possiblePlainText))
		fmt.Printf("Text %s, score %f\n", string(possiblePlainText), score)

		if score > bestScore {
			bestScore = score
			bestText = string(possiblePlainText)
		}
	}

	fmt.Println("best text ", bestText, "score ", bestScore)
}

func TestEnglishness(t *testing.T) {
	inp := "hello cruel world"
	score := cryptopals.Englishness(inp)
	fmt.Printf("score %f\n", score)

	inp = "afsd fjaqpw efqnpqoidqw"
	score = cryptopals.Englishness(inp)
	fmt.Printf("score %f\n", score)
}
