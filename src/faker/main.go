package faker

import (
	"math/rand"
	"time"
)

var CHARACTERS []string = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var SYMBOLS []string = []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "+", "=", "{", "}", "[", "]", "|", ":", ";", "\"", "'", "<", ">", ",", ".", "?", "/", "~", "`", " "}
var source rand.Source = rand.NewSource(time.Now().UnixNano())
var random = rand.New(source)

func Boolean(percentage float64) bool {
	return random.Float64() <= percentage
}

func NumberBetween(from int, to int) int {
	return random.Intn(to-from+1) + from
}

func Character() string {
	position := NumberBetween(0, len(CHARACTERS)-1)
	return CHARACTERS[position]
}

func Symbol() string {
	position := NumberBetween(0, len(SYMBOLS)-1)
	return SYMBOLS[position]
}
