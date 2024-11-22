package module01

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//	DecToBase(14, 16) => "E"
//	DecToBase(14, 2) => "1110"
func DecToBase(dec, base int) string {
	const charset = "0123456789ABCDEF"
	var res string
	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec = dec / base
	}
	return res
}

/*func DecToBase(dec, base int) string {
	var rd []int
	for dec > 0 {
		n := dec % base
		rd = append(rd, n)
		dec /= base
	}

	for i, j := 0, len(rd)-1; i < j; i, j = i+1, j-1 {
		rd[i], rd[j] = rd[j], rd[i]
	}

	symbols := "0123456789ABCDEF"
	var result string

	for _, digits := range rd {
		result += string(symbols[digits])
	}

	return result
}*/
