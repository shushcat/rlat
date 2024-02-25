package stemmer

// TODO: Stem per-text rather than per-word to reuse the `env`.

import (
	english "github.com/snowballstem/snowball/go/algorithms"
	snowball "github.com/snowballstem/snowball/go"
)

func Stem(word string) (string, string) {
	env := snowball.NewEnv(word)
	english.Stem(env)
	return env.Current(), word
}
