package main

func test(a string) bool {
	// Remove all spaces and all non-alphanumeric characters
	cleaned := ""
	for _, ch := range a {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			cleaned += string(ch)
		}
	}

	// convert to lowercase
	lowered := ""
	for _, ch := range cleaned {
		if ch >= 'A' && ch <= 'Z' {
			lowered += string(ch + ('a' - 'A'))
		} else {
			lowered += string(ch)
		}
	}

	// reverse the string
	reversed := ""
	for i := len(lowered) - 1; i >= 0; i-- {
		reversed += string(lowered[i])
	}

	return lowered == reversed

}

func main() {
	// Example usage
	s := "A man, a plan, a canal: Panama"
	result := test(s)
	if result {
		println("true")
	} else {
		println("false")
	}
}
