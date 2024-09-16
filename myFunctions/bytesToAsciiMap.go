package myfunctions

func BytesToAsciiMap(style []byte) map[int]string {
	chars := make(map[int]string)
	line := 1
	next := 9
	char := ""
	nbrChar := 32
	for i := 1; i < len(style) - 1; i++ {
		if i < len(style) - 2 {
		 	if style[i] == '\n' {
				line++
			} else if line == next+1 {
				next += 9
				chars[nbrChar] = char[:len(char)-2]
				nbrChar++
				char = ""
			}
			char += string(style[i])
		} else {
			char += string(style[i])
			chars[nbrChar] = char
		}
	}
	return chars
}
