package main

func scorePlainText(inp []byte) float32 {
	// negative point runes ?
	// normalisedStr := strings.ToLower(string(inp))
	// fmt.Println(normalisedStr)

	var score float32
	for _, ch := range inp {
		score += scoreMap[rune(ch)]
	}
	return score
}

var scoreMap = map[rune]float32{
	'e': 12.00,
	't': 9.00,
	'a': 8.00,
	'i': 8.00,
	'n': 8.00,
	'o': 8.00,
	's': 8.00,
	'h': 6.40,
	'r': 6.20,
	'd': 4.40,
	'l': 4.00,
	'u': 3.40,
	'c': 3.00,
	'm': 3.00,
	'f': 2.50,
	'w': 2.00,
	'y': 2.00,
	'g': 1.70,
	'p': 1.70,
	'b': 1.60,
	'v': 1.20,
	'k': 0.80,
	'q': 0.50,
	'j': 0.40,
	'x': 0.40,
	'z': 0.20,
	'`': -5.00,
}
