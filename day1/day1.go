package day1

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const input = "inputs/day1"

var digit_words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var replacement_word string
var words [][]byte

func Part1() uint64 {
	words = readInput()
	var sum uint64
	alpha := regexp.MustCompile("[[:alpha:]]")
	for _, word := range words {
		processed := alpha.ReplaceAll(word, []byte(""))
		digits := []byte{processed[0], processed[len(processed)-1]}
		number, err := strconv.ParseUint(string(digits), 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		sum += number
	}

	return sum
}

func Part2() uint64 {
	words = readInput()
	var sum uint64
	alpha := regexp.MustCompile("[[:alpha:]]")
	for _, word := range words {
		fd, fdIdx := firstDigitWord(word)
		ld, ldIdx := lastDigitWord(word)
		if fdIdx != -1 {
			word[fdIdx] = fd[0]
		}
		if ldIdx != -1 {
			word[ldIdx] = ld[0]
		}
		word = alpha.ReplaceAll(word, []byte(""))
		digits := []byte{word[0], word[len(word)-1]}
		number, err := strconv.ParseUint(string(digits), 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
		sum += number
	}

	return sum
}

func readInput() [][]byte {
	content, err := os.ReadFile(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(-1)
	}
	return bytes.Fields(content)
}

func firstDigitWord(word []byte) ([]byte, int) {
	index := len(word)

	for i, digit_word := range digit_words {
		new_index := bytes.Index(word, []byte(digit_word))
		if new_index < index && new_index >= 0 {
			index = new_index
			replacement_word = strconv.FormatInt(int64(i+1), 10)
		}
	}
	if index == len(word) {
		index = -1
	}

	return []byte(replacement_word), index
}

func lastDigitWord(word []byte) ([]byte, int) {
	index := -1

	for i, digit_word := range digit_words {
		new_index := bytes.LastIndex(word, []byte(digit_word))
		if new_index > index && new_index >= 0 {
			index = new_index
			replacement_word = strconv.FormatInt(int64(i+1), 10)
		}
	}

	return []byte(replacement_word), index
}
