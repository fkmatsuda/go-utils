package fkstr

// CountFormatParams counts number of expected params on a format string
func CountFormatParams(s string) int {

	countParams := 0
	end := len(s)

	for i := 0; i < end; {
		if s[i] == '%' && (i+1) < end {
			if s[i+1] == '%' {
				i++
			} else {
				countParams++
			}
		}
		i++
	}

	return countParams
}
