package randutil

var (
	LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	UppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Letters          = LowercaseLetters + UppercaseLetters
	Digits           = "0123456789"
	Hex              = Digits + "abcdef"
)

func RandomString(n int, charset string) string {
	candidates := []rune(charset)
	candidatesCount := len(candidates)
	s := make([]rune, n)
	for i := range s {
		s[i] = candidates[rng.Intn(candidatesCount)]
	}
	return string(s)
}
