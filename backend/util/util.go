package util

import (
	"time"
	"math/rand"
	"regexp"
	"unicode/utf8"
	"math"
)

var randSrc = rand.NewSource(time.Now().UnixNano())

const (
	rs5LetterIdxBits = 6
	rs6Letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rs6LettersInt = "0123456789"
	rs6LetterIdxBits = 6
	rs6LetterIdxMask = 1 <<rs5LetterIdxBits - 1
	rs6LetterIdxMax = 63 / rs5LetterIdxBits
)

func RandString6(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), rs6LetterIdxMax
	for i := n-1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rs6LetterIdxMax
		}
		idx := int(cache & rs6LetterIdxMask)
		if idx < len(rs6Letters) {
			b[i] = rs6Letters[idx]
			i--
		}
		cache >>= rs6LetterIdxBits
		remain--
	}
	return string(b)
}

func RandString6OnlyInt(n int) string {
	b := make([]byte, n)
	cache, remain := randSrc.Int63(), rs6LetterIdxMax
	for i := n-1; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), rs6LetterIdxMax
		}
		idx := int(cache & rs6LetterIdxMask)
		if idx < len(rs6LettersInt) {
			b[i] = rs6LettersInt[idx]
			i--
		}
		cache >>= rs6LetterIdxBits
		remain--
	}
	return string(b)
}

func IsMultiByte(str string) bool {
	count := 0
	for i, _ :=  range str {
		count = i
	}
	count++

	return count != len(str)
}

func IsPasswordString(password string) bool {
	rep := regexp.MustCompile("[a-zA-Z0-9!-/:-@[-`{-~]{8,20}$")
	return rep.MatchString(password)
}

func InsertIndention(s string, lineNum float64) (string, int) {
	lineCount := 1
	if lineNum < 1 {
		return s, lineCount
	}
	ls := len(s)
	r := make([]byte, 0, ls + ((int(lineNum) + 1) * 2))
	lineLen := int(math.Trunc(float64(ls) / lineNum))
	oneLineLen := lineLen

	for i, w := 0, 0; i < ls; i += w {
		_, w = utf8.DecodeRuneInString(s[i:])
		if i >= lineLen {
			lineCount++
			lineLen += oneLineLen
			r = append(r, ([]byte("\n"))...)
		}
		r = append(r, s[i:i+w]...)
	}
	return string(r), lineCount
}

func IncludeUint(arr []uint, target uint) bool {
	for _, val := range arr{
		if target == val {
			return true
		}
	}

	return false
}