package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//	Reverse("cat") => "tac"
//	Reverse("alphabet") => "tebahpla"
func Reverse(word string) string {
	/*
		How I did it:
		r := []rune(word)
		for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)

		How the teacher did it:
		var res string
		for i := len(word) - 1; i >= 0; i-- {
			res += string(word[i])
		}
		return res*/

	var res string

	for _, r := range word {
		res = string(r) + res
	}

	return res

}
