package main

import (
	"cryptopals/detect"
	"cryptopals/xor"
	"encoding/hex"
	"fmt"
	"math"
)

func main() {
	inpHex := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inpBytes, _ := hex.DecodeString(inpHex)

	key, msg, score := detectSChXOR(inpBytes)
	fmt.Println(key, msg, score)

	// content, _ := ioutil.ReadFile("./4.txt")
	// lines := bytes.Split(content, []byte("\n"))
	// var bestScore float64 = math.MaxFloat64
	// var bestPlaintext string

	// for _, line := range lines {
	// 	_, plaintext, score := detectSChXOR(line)

	// 	if score < 200 {
	// 	fmt.Println(plaintext, score)
	// 	}

	// 	if score < bestScore {
	// 		bestScore = score
	// 		bestPlaintext = plaintext
	// 	}
	// }

	// fmt.Println(bestPlaintext, bestScore)
}

func detectSChXOR(inp []byte) (rune, string, float64) {
	alphab := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	var bRune rune
	var bMsg string
	var bScore float64 = math.MaxFloat64

	for i := 0; i < len(alphab); i++ {
		c := rune(alphab[i])
		pptext := xor.XORSingleByte(inp, c)
		score := detect.ScoreText(string(pptext))

		if score > bScore {
			// fmt.Println(string(pptext))
			// fmt.Println(score)

			bScore = score
			bMsg = string(pptext)
			bRune = c
		}

	}

	return bRune, bMsg, bScore
}
