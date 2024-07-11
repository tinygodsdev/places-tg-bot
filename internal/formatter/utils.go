package formatter

import (
	"fmt"
	"math"
	"strings"
	"unicode"

	"math/rand"
)

func Bold(s string) string {
	return fmt.Sprintf("<b>%s</b>", s)
}

func Underline(s string) string {
	return fmt.Sprintf("<u>%s</u>", s)
}

func Italic(s string) string {
	return fmt.Sprintf("<i>%s</i>", s)
}

func Capitalize(s string) string {
	if s == "" {
		return ""
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func RandomHappyEmoji() string {
	happyEmojis := []string{
		happyEmoji,
		heartsFaceEmoji,
		satisfiedEmoji,
		loveFaceEmoji,
		happyCatEmoji,
		partyEmoji,
	}

	return happyEmojis[rand.Intn(len(happyEmojis))]
}

func FormatLargeNumber(num float64) string {
	absNum := math.Abs(num)
	var divisor float64
	var suffix string

	switch {
	case absNum >= 1e12:
		divisor = 1e12
		suffix = "T"
	case absNum >= 1e9:
		divisor = 1e9
		suffix = "B"
	case absNum >= 1e6:
		divisor = 1e6
		suffix = "M"
	case absNum >= 1e3:
		divisor = 1e3
		suffix = "K"
	default:
		divisor = 1
		suffix = ""
	}

	formattedNumber := num / divisor
	if divisor > 1 {
		return fmt.Sprintf("%.0f%s", formattedNumber, suffix)
	}
	return fmt.Sprintf("%.0f", formattedNumber)
}
