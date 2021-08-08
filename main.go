package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("./4.txt")
	if err != nil {
		//Do something
	}

	lines := bytes.Split(content, []byte("\n"))
	bestScore := 0
	var bestPlaintext string

	for _, line := range lines {
		_, plaintext, score := TrySingleCharXOR(line)
		fmt.Println(plaintext, score)
		if score > bestScore {
			bestScore = score
			bestPlaintext = plaintext
		}
	}

	fmt.Println(bestScore, bestPlaintext)

	inp := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inpBytes, _ := hex.DecodeString(inp)
	fmt.Println(inpBytes)

	key, plaintext, score := TrySingleCharXOR(inpBytes)
	fmt.Println(key, plaintext, score)
}

// return top 3 or 5 ?
func TrySingleCharXOR(inp []byte) (rune, string, int) {
	alphaBeta := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var topScore float32 = 0
	var bestPlaintext string
	var key rune

	var XORResult []byte
	for _, v := range alphaBeta {
		XORResult = XORSingleByte(inp, v)
		score := scorePlainText(XORResult)
		if score > topScore {
			bestPlaintext = string(XORResult)
			topScore = score
			key = v
		}
	}

	// fmt.Printf("Topscore: %f, index: %d\n", topScore, topScoreIndex)
	return key, bestPlaintext, int(topScore)
}
