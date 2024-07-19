package formatter

import (
	"math/rand"
)

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
