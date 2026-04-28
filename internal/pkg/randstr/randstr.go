package randstr

import "math/rand/v2"


const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var rng = rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64()))

func Generate(n int) string {
	b := make([]byte, n)

	for i := range b {
		b[i] = alphabet[rng.IntN(len(alphabet))]
	}
	return string(b)
}