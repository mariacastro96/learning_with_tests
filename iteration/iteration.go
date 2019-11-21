package iteration

// Repeat takes a character and repeats it 5 times
func Repeat(char string) string {
	var repeatedChar string

	for i := 0; i < 5; i++ {
		repeatedChar += char
	}

	return repeatedChar
}
