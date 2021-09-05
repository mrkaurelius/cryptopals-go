package detect

import (
	"testing"
)

func TestDetect(t *testing.T) {
	score := ScoreText("Here's some basic JS code to calculate")
	want := 0.2
	if want != score {
		t.Error("err")
	}
}
