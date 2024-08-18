package Split

func SplitWithNewlines(str string) []string {
	word := ""
	var argLettersDivided []string
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '\n' {
			if word != "" {
				argLettersDivided = append(argLettersDivided, word)
				word = ""
			}
			if count != 0 {
				argLettersDivided = append(argLettersDivided, "\\n")
			}
			count++

		} else {
			word += string(str[i])
		}
	}
	if word != "" {
		argLettersDivided = append(argLettersDivided, word)
	}
	return argLettersDivided
}