package main

func encode(text string, shift rune) string {
	offset := rune(26)

	runes := []rune(text)
	
	for index, char := range runes {
		if char >= 'a'+shift && char <= 'z' ||
			char >= 'A'+shift && char <= 'Z' {
			char = char - shift
		} else if char >= 'a' && char < 'a'+shift ||
			char >= 'A' && char < 'A'+shift {
			char = char - shift + offset
		}
		runes[index] = char
	}

	return string(runes)
}

func decode(text string, shift rune) string {
	offset := rune(26)

	runes := []rune(text)
	
	for index, char := range runes {
		if char >= 'a'-shift && char <= 'z' ||
			char >= 'A'-shift && char <= 'Z' {
			char = char + shift
		} else if char >= 'a' && char < 'a'+shift ||
			char >= 'A' && char < 'A'+shift {
			char = char + shift - offset
		}
		runes[index] = char
	}

	return string(runes)
}

func main(){
	encoded := encode("i like lynxes", 3)
	println("  encoded: " + encoded)
	decoded := decode(encoded, 3)
	println("  decoded: " + decoded)
}