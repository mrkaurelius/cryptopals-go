package detect

import "math"

func ScoreText(str string) float64 {
	score := getChi2(str)
	// fmt.Println(score)
	return score
}

// TODO
// https://github.com/Lukasa/cryptopals/blob/master/cryptopals/challenge_one/three.py

// https://crypto.stackexchange.com/questions/30209/developing-algorithm-for-detecting-plain-text-via-frequency-analysis
// yemedi
func getChi2(str string) float64 {
	var count [26]int // zero inited
	ignored := 0
	// fmt.Println(count)

	for i := 0; i < len(str); i++ {
		c := str[i]
		// c := int(str[i])
		// fmt.Printf("%d %c\n", c, c)

		if c >= 65 && c <= 90 {
			count[c-65]++ // uppercase A-Z
		} else if c >= 97 && c <= 122 {
			count[c-97]++ // lowercase a-z
		} else if c >= 32 && c <= 126 {
			ignored++ // numbers and punct.
		} else if c == 9 || c == 10 || c == 13 {
			ignored++ // TAB, CR, LF
		} else {
			// fmt.Println("non print")
			// ignored++
			return math.Inf(+1) // not printable ASCII = impossible(?)
		}
	}

	// fmt.Println(count)
	chi2 := 0.0
	len := len(str) - ignored

	for i := 0; i < 26; i++ {
		observed := count[i]
		expected := float64(len) * english_freq[i]
		difference := float64(observed) - expected
		chi2 += difference * difference / expected
	}

	return chi2
}

var english_freq = []float64{
	0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015,
	0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749,
	0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758,
	0.00978, 0.02360, 0.00150, 0.01974, 0.00074}
