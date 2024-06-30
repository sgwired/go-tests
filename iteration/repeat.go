package iteration

// Repeat takes a character and number and returns the charcter printed
// that number of times
func Repeat(character string, number int) string {
	var repeated string
	for i := 0; i < number; i++ {
		repeated += character
	}
	return repeated
}
