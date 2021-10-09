package main

import (
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	count, _ := strconv.Atoi(os.Args[1])
	var phrases []string
	for i := 0; i < count; i++ {
		phrases = append(phrases, generate())
	}

	if len(os.Args) <= 2 || os.Args[2] != "-noprint" {
		for _, str := range phrases {
			println(str)
		}
	}
}

func generate() string {
	phrase := pickRandomly(phrases)
	working := strings.Builder{}
	skip1 := false
	for i, c := range phrase {
		if skip1 {
			skip1 = false
			continue
		}
		if c == '%' {
			replacements := replacers["%"+string(phrase[i+1])]
			working.WriteString(pickRandomly(replacements))
			skip1 = true
		} else {
			working.WriteRune(c)
		}
	}

	return working.String()
}

func pickRandomly(list []string) string {
	return list[int(math.Floor(float64(rand.Float32())*float64(len(list))))]
}
