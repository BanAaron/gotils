package gotils

// LeftPad prefixes X number of spaces to the 'input' string
func LeftPad(input string, spaces uint) (result string) {
	pad := ""
	var i uint
	for i = 0; i < spaces; i++ {
		pad += " "
	}
	result = pad + input
	return
}
