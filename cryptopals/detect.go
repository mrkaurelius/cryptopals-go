package cryptopals

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func DetectSingleByteXor(filepath string) {
	// TODO burada kaldim dosyadan hex encoding decoding te problem olabilir, test caselerini yaz
	content, _ := ioutil.ReadFile(filepath)
	lines := bytes.Split(content, []byte("\n"))

	var bestScore float64
	var bestText string

	// println(line)
	for _, line := range lines {
		for i := 65; i < 91; i++ {
			r := rune(i)
			possiblePlainText := XorSingleByte(line, r)

			// fmt.Println(string(possiblePlainText))

			score := Englishness(string(possiblePlainText))
			// fmt.Printf("Text %s, score %f\n", string(possiblePlainText), score)

			if score > bestScore {
				bestScore = score
				bestText = string(possiblePlainText)
				fmt.Printf("Text %s, score %f\n", string(possiblePlainText), score)
			}
		}

	}

	fmt.Println(bestText, bestScore)
}

// From https://github.com/Lukasa/cryptopals/blob/master/cryptopals/challenge_one/three.py, zor oyunu bozar

func Englishness(plainText string) float64 {
	score := 0.0
	plainText = strings.ToLower(plainText)
	totalCharacters := len(plainText)

	counter := make(map[rune]int, 0)

	for _, v := range plainText {
		counter[v] = counter[v] + 1
	}

	for i, v := range counter {
		charFreq := float64(v) / float64(totalCharacters)
		charEngFreq := englishDist[i]
		// fmt.Println(charEngFreq)
		charScore := math.Sqrt(charFreq * charEngFreq)
		score += charScore
	}

	return score
}

var englishDist = map[rune]float64{
	'a': 0.08167, 'b': 0.01492, 'c': 0.02782,
	'd': 0.04253, 'e': 0.12702, 'f': 0.02228,
	'g': 0.02015, 'h': 0.06094, 'i': 0.06966,
	'j': 0.00153, 'k': 0.00772, 'l': 0.04025,
	'm': 0.02406, 'n': 0.06749, 'o': 0.07507,
	'p': 0.01929, 'q': 0.00095, 'r': 0.05987,
	's': 0.06327, 't': 0.09056, 'u': 0.02758,
	'v': 0.00978, 'w': 0.02360, 'x': 0.00150,
	'y': 0.01974, 'z': 0.00074}
